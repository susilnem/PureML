package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	authmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/auth/middlewares"
	"github.com/PureMLHQ/PureML/packages/purebackend/core"
	coreservice "github.com/PureMLHQ/PureML/packages/purebackend/core/apis/service"
	"github.com/PureMLHQ/PureML/packages/purebackend/core/models"
	userorgmodels "github.com/PureMLHQ/PureML/packages/purebackend/user_org/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// BindUserApi registers the admin api endpoints and the corresponding handlers.
func BindUserApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	userGroup := rg.Group("/user")
	userGroup.GET("/profile", api.DefaultHandler(GetProfile), authmiddlewares.RequireAuthContext)
	userGroup.GET("/profile/:userHandle", api.DefaultHandler(GetProfileByHandle))
	userGroup.POST("/profile", api.DefaultHandler(UpdateProfile), authmiddlewares.RequireAuthContext)
	userGroup.POST("/signup", api.DefaultHandler(UserSignUp))
	userGroup.POST("/login", api.DefaultHandler(UserLogin))
	userGroup.POST("/verify-email", api.DefaultHandler(UserVerifyEmail))
	userGroup.POST("/resend-verification", api.DefaultHandler(UserResendVerification))
	userGroup.POST("/forgot-password", api.DefaultHandler(UserForgotPassword))
	userGroup.POST("/verify-reset-password", api.DefaultHandler(UserVerifyResetPassword))
	userGroup.POST("/reset-password", api.DefaultHandler(UserResetPassword))

	userGroup.POST("/create-session", api.DefaultHandler(CreateSession))
	userGroup.POST("/session-token", api.DefaultHandler(GetSessionToken))
	userGroup.POST("/verify-session", api.DefaultHandler(VerifySession), authmiddlewares.RequireAuthContext)

	userGroup.GET("/tokens", api.DefaultHandler(GetTokens), authmiddlewares.RequireAuthContext)
	userGroup.POST("/create-token", api.DefaultHandler(CreateToken), authmiddlewares.RequireAuthContext)
	userGroup.DELETE("/delete-token/:tokenUUID", api.DefaultHandler(DeleteToken), authmiddlewares.RequireAuthContext)
}

// UserSignUp godoc
//
//	@Summary		User sign up.
//	@Description	User sign up with email, name, handle and password. The user will receive an email with a verification link if enabled from backend.
//	@Description	Response status code 202 means that the user has been created but the email verification is pending.
//	@Description	Response status code 200 means that the user has been created and no verification of email is needed.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/signup [post]
//	@Param			user	body	models.UserSignupRequest	true	"User details"
func (api *Api) UserSignUp(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	var emailData string
	if email == nil {
		emailData = ""
	} else {
		emailData = email.(string)
	}
	if emailData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is required")
	}
	if addr, ok := coreservice.ValidateMailAddress(emailData); ok {
		emailData = addr
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
	}
	handle := request.GetParsedBodyAttribute("handle")
	var handleData string
	if handle == nil {
		handleData = ""
	} else {
		handleData = handle.(string)
	}
	if handleData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Handle is required")
	}
	name := request.GetParsedBodyAttribute("name")
	if name == nil {
		name = ""
	}
	bio := request.GetParsedBodyAttribute("bio")
	if bio == nil {
		bio = ""
	}
	avatar := request.GetParsedBodyAttribute("avatar")
	if avatar == nil {
		avatar = ""
	}
	password := request.GetParsedBodyAttribute("password")
	var passwordData string
	if password == nil {
		passwordData = ""
	} else {
		passwordData = password.(string)
	}
	if passwordData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Password is required")
	}
	dbuser, err := api.app.Dao().GetUserByEmail(emailData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if dbuser != nil {
		return models.NewErrorResponse(http.StatusConflict, "User with email already exists")
	}
	dbuserhandle, err := api.app.Dao().GetUserByHandle(handleData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if dbuserhandle != nil {
		return models.NewErrorResponse(http.StatusConflict, "User with handle already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordData), bcrypt.DefaultCost)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	user, err := api.app.Dao().CreateUser(name.(string), emailData, handleData, bio.(string), avatar.(string), string(hashedPassword), !api.app.Settings().MailService.Enabled)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if api.app.Settings().MailService.Enabled {
		// Send verification email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid":  user.UUID,
			"email": user.Email,
			"exp":   time.Now().Unix() + api.app.Settings().MailVerifificationAuthToken.Duration,
		})
		signedString, err := token.SignedString([]byte(api.app.Settings().MailVerifificationAuthToken.Secret))
		if err != nil {
			panic(err)
		}
		verifyLink := api.app.Settings().Site.BaseURL + "/verify-email?token=" + signedString
		emailTemplate := coreservice.BaseEmailTemplate("Email Verification",
			`<div style="margin: 0px auto; max-width: 600px">
			<table
			  align="center"
			  border="0"
			  cellpadding="0"
			  cellspacing="0"
			  role="presentation"
			  style="width: 100%"
			>
			  <tbody>
				<tr>
				  <td
					style="
					  direction: ltr;
					  font-size: 0px;
					  padding: 0px;
					  text-align: center;
					"
				  >
					<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
					<div
					  class="mj-column-per-100 mj-outlook-group-fix"
					  style="
						font-size: 0px;
						text-align: left;
						direction: ltr;
						display: inline-block;
						vertical-align: top;
						width: 100%;
					  "
					>
					  <table
						border="0"
						cellpadding="0"
						cellspacing="0"
						role="presentation"
						style="vertical-align: top"
						width="100%"
					  >
						<tbody>
						  <tr>
							<td
							  align="center"
							  style="
								font-size: 0px;
								padding: 10px 25px;
								word-break: break-word;
							  "
							>
							  <div
								style="
								  font-family: Roboto, Helvetica, sans-serif;
								  font-size: 16px;
								  font-weight: 300;
								  line-height: 24px;
								  text-align: center;
								  color: #616161;
								"
							  >
								Hi `+user.Handle+`, click on the button below to
								verify your email.
							  </div>
							</td>
						  </tr>
						</tbody>
					  </table>
					</div>
					<!--[if mso | IE]></td></tr></table><![endif]-->
				  </td>
				</tr>
			  </tbody>
			</table>
		  </div>
		  <!--[if mso | IE]></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="" style="width:600px;" width="600" ><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->
		  <div style="margin: 0px auto; max-width: 600px">
			<table
			  align="center"
			  border="0"
			  cellpadding="0"
			  cellspacing="0"
			  role="presentation"
			  style="width: 100%"
			>
			  <tbody>
				<tr>
				  <td
					style="
					  direction: ltr;
					  font-size: 0px;
					  padding: 0px;
					  padding-top: 36px;
					  text-align: center;
					"
				  >
					<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
					<div
					  class="mj-column-per-100 mj-outlook-group-fix"
					  style="
						font-size: 0px;
						text-align: left;
						direction: ltr;
						display: inline-block;
						vertical-align: top;
						width: 100%;
					  "
					>
					  <table
						border="0"
						cellpadding="0"
						cellspacing="0"
						role="presentation"
						style="vertical-align: top"
						width="100%"
					  >
						<tbody>
						  <tr>
							<td
							  align="center"
							  vertical-align="middle"
							  style="
								font-size: 0px;
								padding: 10px 25px;
								word-break: break-word;
							  "
							>
							  <table
								border="0"
								cellpadding="0"
								cellspacing="0"
								role="presentation"
								style="
								  border-collapse: separate;
								  line-height: 100%;
								"
							  >
								<tr>
								  <td
									align="center"
									bgcolor="#191f4d"
									role="presentation"
									style="
									  border: none;
									  border-radius: 3px;
									  cursor: auto;
									  mso-padding-alt: 10px 25px;
									  background: #191f4d;
									"
									valign="middle"
								  >
								  	<a href="`+verifyLink+`" target="_blank">
										<p
										style="
											display: inline-block;
											background: #191f4d;
											color: #ffffff;
											font-family: Roboto, Helvetica, sans-serif;
											font-size: 13px;
											font-weight: normal;
											line-height: 120%;
											margin: 0;
											text-decoration: none;
											text-transform: none;
											padding: 10px 25px;
											mso-padding-alt: 0px;
											border-radius: 3px;
										"
										>
										Verify my email
										</p>
									</a>
								  </td>
								</tr>
							  </table>
							</td>
						  </tr>
						</tbody>
					  </table>
					</div>
					<!--[if mso | IE]></td></tr></table><![endif]-->
				  </td>
				</tr>
			  </tbody>
			</table>
		  </div>
		  <!--[if mso | IE]></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="" style="width:600px;" width="600" ><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->
		  <div style="margin: 0px auto; max-width: 600px">
			<table
			  align="center"
			  border="0"
			  cellpadding="0"
			  cellspacing="0"
			  role="presentation"
			  style="width: 100%"
			>
			  <tbody>
				<tr>
				  <td
					style="
					  direction: ltr;
					  font-size: 0px;
					  padding: 0px;
					  padding-top: 24px;
					  text-align: center;
					"
				  >
					<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
					<div
					  class="mj-column-per-100 mj-outlook-group-fix"
					  style="
						font-size: 0px;
						text-align: left;
						direction: ltr;
						display: inline-block;
						vertical-align: top;
						width: 100%;
					  "
					>
					  <table
						border="0"
						cellpadding="0"
						cellspacing="0"
						role="presentation"
						style="vertical-align: top"
						width="100%"
					  >
						<tbody>
						  <tr>
							<td
							  align="center"
							  style="
								font-size: 0px;
								padding: 10px 25px;
								word-break: break-word;
							  "
							>
							  <p
								style="
								  border-top: solid 1px #e0e0e0;
								  font-size: 1px;
								  margin: 0px auto;
								  width: 100%;
								"
							  ></p>
							  <!--[if mso | IE
								]><table
								  align="center"
								  border="0"
								  cellpadding="0"
								  cellspacing="0"
								  style="
									border-top: solid 1px #e0e0e0;
									font-size: 1px;
									margin: 0px auto;
									width: 550px;
								  "
								  role="presentation"
								  width="550px"
								>
								  <tr>
									<td style="height: 0; line-height: 0">
									  &nbsp;
									</td>
								  </tr>
								</table><!
							  [endif]-->
							</td>
						  </tr>
						</tbody>
					  </table>
					</div>
					<!--[if mso | IE]></td></tr></table><![endif]-->
				  </td>
				</tr>
			  </tbody>
			</table>
		  </div>`,
		)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		subject := "Verify your email address | PureML"
		err = api.app.SendMail(user.Email, subject, emailTemplate)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		return models.NewDataResponse(http.StatusAccepted, user, "User created. Please verify your email address")
	}
	return models.NewDataResponse(http.StatusOK, user, "User created")
}

// UserLogin godoc
//
//	@Summary		User login.
//	@Description	User login with email and password.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/login [post]
//	@Param			org	body	models.UserLoginRequest	true	"User details"
func (api *Api) UserLogin(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	var emailData string
	if email == nil {
		emailData = ""
	} else {
		emailData = email.(string)
	}
	handle := request.GetParsedBodyAttribute("handle")
	var handleData string
	if handle == nil {
		handleData = ""
	} else {
		handleData = handle.(string)
	}
	if emailData == "" && handleData == "" {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Email or handle is required")
	}
	if emailData != "" {
		if addr, ok := coreservice.ValidateMailAddress(emailData); ok {
			emailData = addr
		} else {
			return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
		}
	}
	password := request.GetParsedBodyAttribute("password")
	var passwordData string
	if password == nil {
		passwordData = ""
	} else {
		passwordData = password.(string)
	}
	if passwordData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Password is required")
	}
	var user *userorgmodels.UserResponse
	var err error
	if email != nil {
		user, err = api.app.Dao().GetSecureUserByEmail(emailData)
	} else {
		user, err = api.app.Dao().GetSecureUserByHandle(handleData)
	}
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordData))
	if err != nil {
		return models.NewDataResponse(http.StatusUnauthorized, nil, "Invalid credentials")
	}
	if !user.IsVerified {
		return models.NewDataResponse(http.StatusUnauthorized, nil, "User email is not verified")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":   user.UUID,
		"email":  user.Email,
		"handle": user.Handle,
	})
	signedString, err := token.SignedString([]byte(api.app.Settings().AdminAuthToken.Secret))
	if err != nil {
		panic(err)
	}
	data := []map[string]string{
		{
			"email":       user.Email,
			"accessToken": signedString,
		},
	}
	return models.NewDataResponse(http.StatusAccepted, data, "User logged in")
}

// UserVerifyEmail godoc
//
//	@Summary		User verify email.
//	@Description	User can verify email by providing verification token.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/verify-email [post]
//	@Param			org	body	models.UserVerifyTokenRequest	true	"Verification token"
func (api *Api) UserVerifyEmail(request *models.Request) *models.Response {
	request.ParseJsonBody()
	token := request.GetParsedBodyAttribute("token")
	var tokenData string
	if token == nil {
		tokenData = ""
	} else {
		tokenData = token.(string)
	}
	if tokenData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Token is required")
	}
	parsedToken, err := jwt.Parse(tokenData, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing algorithm %v", t.Method.Alg())
		}
		return []byte(api.app.Settings().MailVerifificationAuthToken.Secret), nil
	})
	if err != nil {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Could not parse verification token")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userUUID := uuid.Must(uuid.FromString(claims["uuid"].(string)))
		expireTime := int64(claims["exp"].(float64))
		if expireTime < time.Now().Unix() {
			return models.NewDataResponse(http.StatusBadRequest, nil, "Verification token expired")
		}
		user, err := api.app.Dao().GetUserByUUID(userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		if user == nil {
			return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
		}
		if user.IsVerified {
			return models.NewDataResponse(http.StatusBadRequest, nil, "User email is already verified")
		}
		err = api.app.Dao().VerifyUserEmail(userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
	} else {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Invalid verification token")
	}
	response := models.NewDataResponse(http.StatusOK, nil, "User verified")
	return response
}

// UserResendVerification godoc
//
//	@Summary		User resend verification.
//	@Description	User can resend verification email by providing email id.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/resend-verification [post]
//	@Param			data	body	models.UserEmailRequest	true	"Email id"
func (api *Api) UserResendVerification(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	var emailData string
	if email == nil {
		emailData = ""
	} else {
		emailData = email.(string)
	}
	if emailData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is required")
	}
	if emailData != "" {
		if addr, ok := coreservice.ValidateMailAddress(emailData); ok {
			emailData = addr
		} else {
			return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
		}
	}
	user, err := api.app.Dao().GetUserByEmail(emailData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
	}
	if user.IsVerified {
		return models.NewDataResponse(http.StatusBadRequest, nil, "User email is already verified")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":  user.UUID,
		"email": user.Email,
		"exp":   time.Now().Unix() + api.app.Settings().MailVerifificationAuthToken.Duration,
	})
	signedString, err := token.SignedString([]byte(api.app.Settings().MailVerifificationAuthToken.Secret))
	if err != nil {
		panic(err)
	}
	verifyLink := api.app.Settings().Site.BaseURL + "/verify-email?token=" + signedString
	emailTemplate := coreservice.BaseEmailTemplate("Email Verification",
		`<div style="margin: 0px auto; max-width: 600px">
			<table
			  align="center"
			  border="0"
			  cellpadding="0"
			  cellspacing="0"
			  role="presentation"
			  style="width: 100%"
			>
			  <tbody>
				<tr>
				  <td
					style="
					  direction: ltr;
					  font-size: 0px;
					  padding: 0px;
					  text-align: center;
					"
				  >
					<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
					<div
					  class="mj-column-per-100 mj-outlook-group-fix"
					  style="
						font-size: 0px;
						text-align: left;
						direction: ltr;
						display: inline-block;
						vertical-align: top;
						width: 100%;
					  "
					>
					  <table
						border="0"
						cellpadding="0"
						cellspacing="0"
						role="presentation"
						style="vertical-align: top"
						width="100%"
					  >
						<tbody>
						  <tr>
							<td
							  align="center"
							  style="
								font-size: 0px;
								padding: 10px 25px;
								word-break: break-word;
							  "
							>
							  <div
								style="
								  font-family: Roboto, Helvetica, sans-serif;
								  font-size: 16px;
								  font-weight: 300;
								  line-height: 24px;
								  text-align: center;
								  color: #616161;
								"
							  >
								Hi `+user.Handle+`, click on the button below to
								verify your email.
							  </div>
							</td>
						  </tr>
						</tbody>
					  </table>
					</div>
					<!--[if mso | IE]></td></tr></table><![endif]-->
				  </td>
				</tr>
			  </tbody>
			</table>
		  </div>
		  <!--[if mso | IE]></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="" style="width:600px;" width="600" ><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->
		  <div style="margin: 0px auto; max-width: 600px">
			<table
			  align="center"
			  border="0"
			  cellpadding="0"
			  cellspacing="0"
			  role="presentation"
			  style="width: 100%"
			>
			  <tbody>
				<tr>
				  <td
					style="
					  direction: ltr;
					  font-size: 0px;
					  padding: 0px;
					  padding-top: 36px;
					  text-align: center;
					"
				  >
					<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
					<div
					  class="mj-column-per-100 mj-outlook-group-fix"
					  style="
						font-size: 0px;
						text-align: left;
						direction: ltr;
						display: inline-block;
						vertical-align: top;
						width: 100%;
					  "
					>
					  <table
						border="0"
						cellpadding="0"
						cellspacing="0"
						role="presentation"
						style="vertical-align: top"
						width="100%"
					  >
						<tbody>
						  <tr>
							<td
							  align="center"
							  vertical-align="middle"
							  style="
								font-size: 0px;
								padding: 10px 25px;
								word-break: break-word;
							  "
							>
							  <table
								border="0"
								cellpadding="0"
								cellspacing="0"
								role="presentation"
								style="
								  border-collapse: separate;
								  line-height: 100%;
								"
							  >
								<tr>
								  <td
									align="center"
									bgcolor="#191f4d"
									role="presentation"
									style="
									  border: none;
									  border-radius: 3px;
									  cursor: auto;
									  mso-padding-alt: 10px 25px;
									  background: #191f4d;
									"
									valign="middle"
								  >
								  	<a href="`+verifyLink+`" target="_blank">
										<p
										style="
											display: inline-block;
											background: #191f4d;
											color: #ffffff;
											font-family: Roboto, Helvetica, sans-serif;
											font-size: 13px;
											font-weight: normal;
											line-height: 120%;
											margin: 0;
											text-decoration: none;
											text-transform: none;
											padding: 10px 25px;
											mso-padding-alt: 0px;
											border-radius: 3px;
										"
										>
										Verify my email
										</p>
									</a>
								  </td>
								</tr>
							  </table>
							</td>
						  </tr>
						</tbody>
					  </table>
					</div>
					<!--[if mso | IE]></td></tr></table><![endif]-->
				  </td>
				</tr>
			  </tbody>
			</table>
		  </div>
		  <!--[if mso | IE]></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="" style="width:600px;" width="600" ><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->
		  <div style="margin: 0px auto; max-width: 600px">
			<table
			  align="center"
			  border="0"
			  cellpadding="0"
			  cellspacing="0"
			  role="presentation"
			  style="width: 100%"
			>
			  <tbody>
				<tr>
				  <td
					style="
					  direction: ltr;
					  font-size: 0px;
					  padding: 0px;
					  padding-top: 24px;
					  text-align: center;
					"
				  >
					<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
					<div
					  class="mj-column-per-100 mj-outlook-group-fix"
					  style="
						font-size: 0px;
						text-align: left;
						direction: ltr;
						display: inline-block;
						vertical-align: top;
						width: 100%;
					  "
					>
					  <table
						border="0"
						cellpadding="0"
						cellspacing="0"
						role="presentation"
						style="vertical-align: top"
						width="100%"
					  >
						<tbody>
						  <tr>
							<td
							  align="center"
							  style="
								font-size: 0px;
								padding: 10px 25px;
								word-break: break-word;
							  "
							>
							  <p
								style="
								  border-top: solid 1px #e0e0e0;
								  font-size: 1px;
								  margin: 0px auto;
								  width: 100%;
								"
							  ></p>
							  <!--[if mso | IE
								]><table
								  align="center"
								  border="0"
								  cellpadding="0"
								  cellspacing="0"
								  style="
									border-top: solid 1px #e0e0e0;
									font-size: 1px;
									margin: 0px auto;
									width: 550px;
								  "
								  role="presentation"
								  width="550px"
								>
								  <tr>
									<td style="height: 0; line-height: 0">
									  &nbsp;
									</td>
								  </tr>
								</table><!
							  [endif]-->
							</td>
						  </tr>
						</tbody>
					  </table>
					</div>
					<!--[if mso | IE]></td></tr></table><![endif]-->
				  </td>
				</tr>
			  </tbody>
			</table>
		  </div>`,
	)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	subject := "Verify your email address | PureML"
	err = api.app.SendMail(user.Email, subject, emailTemplate)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, nil, "Verification email sent")
	return response
}

// UserForgotPassword godoc
//
//	@Summary		User forgot password.
//	@Description	User can reset password by providing email id to send reset password link.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/forgot-password [post]
//	@Param			org	body	models.UserEmailRequest	true	"User email"
func (api *Api) UserForgotPassword(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	user, err := api.app.Dao().GetUserByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User with given email not found")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":  user.UUID,
		"email": user.Email,
		"exp":   time.Now().Unix() + api.app.Settings().PasswordResetAuthToken.Duration,
	})
	signedString, err := token.SignedString([]byte(api.app.Settings().PasswordResetAuthToken.Secret))
	if err != nil {
		panic(err)
	}
	verifyLink := api.app.Settings().Site.BaseURL + "/reset-password?token=" + signedString
	emailTemplate := coreservice.BaseEmailTemplate("Reset your password",
		`<div style="margin: 0px auto; max-width: 600px">
		<table
		  align="center"
		  border="0"
		  cellpadding="0"
		  cellspacing="0"
		  role="presentation"
		  style="width: 100%"
		>
		  <tbody>
			<tr>
			  <td
				style="
				  direction: ltr;
				  font-size: 0px;
				  padding: 0px;
				  text-align: center;
				"
			  >
				<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
				<div
				  class="mj-column-per-100 mj-outlook-group-fix"
				  style="
					font-size: 0px;
					text-align: left;
					direction: ltr;
					display: inline-block;
					vertical-align: top;
					width: 100%;
				  "
				>
				  <table
					border="0"
					cellpadding="0"
					cellspacing="0"
					role="presentation"
					style="vertical-align: top"
					width="100%"
				  >
					<tbody>
					  <tr>
						<td
						  align="center"
						  style="
							font-size: 0px;
							padding: 10px 25px;
							word-break: break-word;
						  "
						>
						  <div
							style="
							  font-family: Roboto, Helvetica, sans-serif;
							  font-size: 16px;
							  font-weight: 300;
							  line-height: 24px;
							  text-align: center;
							  color: #616161;
							"
						  >
							Hi `+user.Handle+`, click on the button below to
							reset your password.
						  </div>
						</td>
					  </tr>
					</tbody>
				  </table>
				</div>
				<!--[if mso | IE]></td></tr></table><![endif]-->
			  </td>
			</tr>
		  </tbody>
		</table>
	  </div>
	  <!--[if mso | IE]></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="" style="width:600px;" width="600" ><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->
	  <div style="margin: 0px auto; max-width: 600px">
		<table
		  align="center"
		  border="0"
		  cellpadding="0"
		  cellspacing="0"
		  role="presentation"
		  style="width: 100%"
		>
		  <tbody>
			<tr>
			  <td
				style="
				  direction: ltr;
				  font-size: 0px;
				  padding: 0px;
				  padding-top: 36px;
				  text-align: center;
				"
			  >
				<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
				<div
				  class="mj-column-per-100 mj-outlook-group-fix"
				  style="
					font-size: 0px;
					text-align: left;
					direction: ltr;
					display: inline-block;
					vertical-align: top;
					width: 100%;
				  "
				>
				  <table
					border="0"
					cellpadding="0"
					cellspacing="0"
					role="presentation"
					style="vertical-align: top"
					width="100%"
				  >
					<tbody>
					  <tr>
						<td
						  align="center"
						  vertical-align="middle"
						  style="
							font-size: 0px;
							padding: 10px 25px;
							word-break: break-word;
						  "
						>
						  <table
							border="0"
							cellpadding="0"
							cellspacing="0"
							role="presentation"
							style="
							  border-collapse: separate;
							  line-height: 100%;
							"
						  >
							<tr>
							  <td
								align="center"
								bgcolor="#191f4d"
								role="presentation"
								style="
								  border: none;
								  border-radius: 3px;
								  cursor: auto;
								  mso-padding-alt: 10px 25px;
								  background: #191f4d;
								"
								valign="middle"
							  >
							  	<a href="` + verifyLink + `" target="_blank" >
									<p
									style="
										display: inline-block;
										background: #191f4d;
										color: #ffffff;
										font-family: Roboto, Helvetica, sans-serif;
										font-size: 13px;
										font-weight: normal;
										line-height: 120%;
										margin: 0;
										text-decoration: none;
										text-transform: none;
										padding: 10px 25px;
										mso-padding-alt: 0px;
										border-radius: 3px;
									"
									>
									Reset my password
									</p>
								</a>
							  </td>
							</tr>
						  </table>
						</td>
					  </tr>
					</tbody>
				  </table>
				</div>
				<!--[if mso | IE]></td></tr></table><![endif]-->
			  </td>
			</tr>
		  </tbody>
		</table>
	  </div>
	  <!--[if mso | IE]></td></tr></table><table align="center" border="0" cellpadding="0" cellspacing="0" class="" style="width:600px;" width="600" ><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->
	  <div style="margin: 0px auto; max-width: 600px">
		<table
		  align="center"
		  border="0"
		  cellpadding="0"
		  cellspacing="0"
		  role="presentation"
		  style="width: 100%"
		>
		  <tbody>
			<tr>
			  <td
				style="
				  direction: ltr;
				  font-size: 0px;
				  padding: 0px;
				  padding-top: 24px;
				  text-align: center;
				"
			  >
				<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]-->
				<div
				  class="mj-column-per-100 mj-outlook-group-fix"
				  style="
					font-size: 0px;
					text-align: left;
					direction: ltr;
					display: inline-block;
					vertical-align: top;
					width: 100%;
				  "
				>
				  <table
					border="0"
					cellpadding="0"
					cellspacing="0"
					role="presentation"
					style="vertical-align: top"
					width="100%"
				  >
					<tbody>
					  <tr>
						<td
						  align="center"
						  style="
							font-size: 0px;
							padding: 10px 25px;
							word-break: break-word;
						  "
						>
						  <p
							style="
							  border-top: solid 1px #e0e0e0;
							  font-size: 1px;
							  margin: 0px auto;
							  width: 100%;
							"
						  ></p>
						  <!--[if mso | IE
							]><table
							  align="center"
							  border="0"
							  cellpadding="0"
							  cellspacing="0"
							  style="
								border-top: solid 1px #e0e0e0;
								font-size: 1px;
								margin: 0px auto;
								width: 550px;
							  "
							  role="presentation"
							  width="550px"
							>
							  <tr>
								<td style="height: 0; line-height: 0">
								  &nbsp;
								</td>
							  </tr>
							</table><!
						  [endif]-->
						</td>
					  </tr>
					</tbody>
				  </table>
				</div>
				<!--[if mso | IE]></td></tr></table><![endif]-->
			  </td>
			</tr>
		  </tbody>
		</table>
	  </div>`,
	)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	subject := "Reset your password | PureML"
	err = api.app.SendMail(user.Email, subject, emailTemplate)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Reset password link sent to your mail")
}

// UserVerifyResetPassword godoc
//
//	@Summary		User reset password verify token.
//	@Description	User can verify token to view reset password form.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/verify-reset-password [post]
//	@Param			org	body	models.UserVerifyTokenRequest	true	"Verification token"
func (api *Api) UserVerifyResetPassword(request *models.Request) *models.Response {
	request.ParseJsonBody()
	token := request.GetParsedBodyAttribute("token")
	var tokenData string
	if token == nil {
		tokenData = ""
	} else {
		tokenData = token.(string)
	}
	if tokenData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Token is required")
	}
	parsedToken, err := jwt.Parse(tokenData, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing algorithm %v", t.Method.Alg())
		}
		return []byte(api.app.Settings().PasswordResetAuthToken.Secret), nil
	})
	if err != nil {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Could not parse verification token")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userUUID := uuid.Must(uuid.FromString(claims["uuid"].(string)))
		expireTime := int64(claims["exp"].(float64))
		if expireTime < time.Now().Unix() {
			return models.NewDataResponse(http.StatusBadRequest, nil, "Password Reset token expired")
		}
		user, err := api.app.Dao().GetSecureUserByUUID(userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		if user == nil {
			return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
		}
	} else {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Invalid verification token")
	}
	return models.NewDataResponse(http.StatusOK, nil, "Password reset token valid")
}

// UserResetPassword godoc
//
//	@Summary		User reset password.
//	@Description	User can reset password by providing old password and new password.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/reset-password [post]
//	@Param			org	body	models.UserResetPasswordRequest	true	"User email"
func (api *Api) UserResetPassword(request *models.Request) *models.Response {
	request.ParseJsonBody()
	token := request.GetParsedBodyAttribute("token")
	var tokenData string
	if token == nil {
		tokenData = ""
	} else {
		tokenData = token.(string)
	}
	if tokenData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Token is required")
	}
	parsedToken, err := jwt.Parse(tokenData, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing algorithm %v", t.Method.Alg())
		}
		return []byte(api.app.Settings().PasswordResetAuthToken.Secret), nil
	})
	if err != nil {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Could not parse verification token")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userUUID := uuid.Must(uuid.FromString(claims["uuid"].(string)))
		expireTime := int64(claims["exp"].(float64))
		if expireTime < time.Now().Unix() {
			return models.NewDataResponse(http.StatusBadRequest, nil, "Password Reset token expired")
		}
		user, err := api.app.Dao().GetSecureUserByUUID(userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		if user == nil {
			return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
		}
		newPassword := request.GetParsedBodyAttribute("new_password")
		var newPasswordData string
		if newPassword == nil {
			newPasswordData = ""
		} else {
			newPasswordData = newPassword.(string)
		}
		if newPasswordData == "" {
			return models.NewErrorResponse(http.StatusBadRequest, "New password is required")
		}
		//TODO: Validate password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPasswordData), bcrypt.DefaultCost)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		err = api.app.Dao().UpdateUserPassword(userUUID, string(hashedPassword))
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
	} else {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Invalid verification token")
	}
	return models.NewDataResponse(http.StatusOK, nil, "Password reset successfully")
}

// GetProfile godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get logged in user profile.
//	@Description	Get logged in user profile.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/profile [get]
func (api *Api) GetProfile(request *models.Request) *models.Response {
	userUUID := request.GetUserUUID()
	user, err := api.app.Dao().GetUserProfileByUUID(userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, user, "User profile")
}

// GetProfileByHandle godoc
//
//	@Summary		Get user profile by handle.
//	@Description	Get user profile by handle. Accessible without login.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/profile/{userHandle} [get]
//	@Param			userHandle	path	string	true	"User handle"
func (api *Api) GetProfileByHandle(request *models.Request) *models.Response {
	userHandle := request.GetPathParam("userHandle")
	user, err := api.app.Dao().GetUserByHandle(userHandle)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, user, "Public User profile")
}

// UpdateProfile godoc
//
//	@Security		ApiKeyAuth
//	@Summary		User update profile.
//	@Description	User can update profile by providing name, avatar, bio.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/profile [post]
//	@Param			org	body	models.UserUpdateRequest	true	"User details"
func (api *Api) UpdateProfile(request *models.Request) *models.Response {
	request.ParseJsonBody()
	name := request.GetParsedBodyAttribute("name")
	avatar := request.GetParsedBodyAttribute("avatar")
	bio := request.GetParsedBodyAttribute("bio")
	updatedAttributes := map[string]interface{}{}
	if name != nil {
		updatedAttributes["name"] = name.(string)
		if name.(string) == "" {
			return models.NewErrorResponse(http.StatusBadRequest, "Name cannot be empty")
		}
	}
	if avatar != nil {
		updatedAttributes["avatar"] = avatar.(string)
	}
	if bio != nil {
		updatedAttributes["bio"] = bio.(string)
	}
	email := request.GetUserMail()
	user, err := api.app.Dao().UpdateUser(email, updatedAttributes)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, user, "User profile updated")
}

// TODO: DeleteProfile godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete user profile.
//	@Description	Delete user profile.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/delete [delete]
func (api *Api) DeleteProfile(request *models.Request) *models.Response {
	return nil
}

// CreateSession godoc
//
//	@Summary		Create Session for CLI.
//	@Description	Create empty session flow.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/create-session [post]
//	@Param			org	body	models.CreateSessionRequest	true	"Session details"
func (api *Api) CreateSession(request *models.Request) *models.Response {
	request.ParseJsonBody()
	device := request.GetParsedBodyAttribute("device")
	var deviceData string
	if device == nil {
		deviceData = ""
	} else {
		deviceData = device.(string)
	}
	deviceId := request.GetParsedBodyAttribute("device_id")
	var deviceIdData string
	if deviceId == nil {
		deviceIdData = ""
	} else {
		deviceIdData = deviceId.(string)
	}
	if deviceIdData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Device Id is required")
	}
	deviceLoc := request.GetParsedBodyAttribute("device_location")
	var deviceLocData string
	if deviceLoc == nil {
		deviceLocData = ""
	} else {
		deviceLocData = deviceLoc.(string)
	}
	session, err := api.app.Dao().CreateSession(deviceData, deviceIdData, deviceLocData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, session, "Session created")
}

// GetSessionToken godoc
//
//	@Summary		Get Session Token if approved for CLI.
//	@Description	Get Session Token if approved for CLI.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/session-token [post]
//	@Param			org	body	models.SessionTokenRequest	true	"Session details"
func (api *Api) GetSessionToken(request *models.Request) *models.Response {
	request.ParseJsonBody()
	sessionId := request.GetParsedBodyAttribute("session_id")
	var sessionIdData uuid.UUID
	if sessionId == nil {
		sessionIdData = uuid.Nil
	} else {
		sessionIdData = uuid.FromStringOrNil(sessionId.(string))
	}
	if sessionIdData == uuid.Nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Session Id is required")
	}
	deviceId := request.GetParsedBodyAttribute("device_id")
	var deviceIdData string
	if deviceId == nil {
		deviceIdData = ""
	} else {
		deviceIdData = deviceId.(string)
	}
	if deviceIdData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Device Id is required")
	}
	session, err := api.app.Dao().GetSession(sessionIdData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if session == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Session not found")
	}
	if !session.Approved {
		return models.NewErrorResponse(http.StatusUnauthorized, "Session not approved")
	}
	if session.Invalid {
		return models.NewErrorResponse(http.StatusForbidden, "Session invalid")
	}
	if session.DeviceId != deviceIdData {
		return models.NewErrorResponse(http.StatusForbidden, "Session invalid device")
	}
	userDb, err := api.app.Dao().GetUserByUUID(session.UserUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":   userDb.UUID,
		"email":  userDb.Email,
		"handle": userDb.Handle,
	})
	signedString, err := token.SignedString([]byte(api.app.Settings().AdminAuthToken.Secret))
	if err != nil {
		panic(err)
	}
	// Set session as invalid hence cannot be reused to get token.
	_, err = api.app.Dao().UpdateSession(sessionIdData, userDb.UUID, map[string]interface{}{
		"invalid": true,
	})
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	data := []map[string]string{
		{
			"email":       userDb.Email,
			"accessToken": signedString,
		},
	}
	return models.NewDataResponse(http.StatusOK, data, "Session Token created")
}

// VerifySession godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Verify Session and approve if valid for CLI.
//	@Description	Verify Session and approve if valid for CLI.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/verify-session [post]
//	@Param			org	body	models.VerifySessionRequest	true	"Session details"
func (api *Api) VerifySession(request *models.Request) *models.Response {
	request.ParseJsonBody()
	sessionId := request.GetParsedBodyAttribute("session_id")
	var sessionIdData uuid.UUID
	if sessionId == nil {
		sessionIdData = uuid.Nil
	} else {
		sessionIdData = uuid.FromStringOrNil(sessionId.(string))
	}
	if sessionIdData == uuid.Nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Session Id is required")
	}
	session, err := api.app.Dao().GetSession(sessionIdData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if session == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Session not found")
	}
	if session.Approved {
		return models.NewErrorResponse(http.StatusBadRequest, "Session already approved")
	}
	if session.Invalid {
		return models.NewErrorResponse(http.StatusForbidden, "Session invalid")
	}
	userUUID := request.GetUserUUID()
	if session.CreatedAt.Unix() < time.Now().Add(-time.Minute*10).Unix() {
		_, err = api.app.Dao().UpdateSession(sessionIdData, userUUID, map[string]interface{}{
			"invalid": true,
		})
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		return models.NewErrorResponse(http.StatusForbidden, "Session expired")
	}
	_, err = api.app.Dao().UpdateSession(sessionIdData, userUUID, map[string]interface{}{
		"approved": true,
	})
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Session approved")
}

// GetTokens godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get API Tokens of user.
//	@Description	Get API token for user
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/tokens [get]
func (api *Api) GetTokens(request *models.Request) *models.Response {
	userUUID := request.GetUserUUID()
	apiToken, err := api.app.Dao().GetTokens(userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, apiToken, "Tokens fetched")
}

// CreateToken godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create API Token.
//	@Description	Create a new API token for user
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/create-token [post]
//	@Param			token	body	models.CreateTokenRequest	true	"Token details"
func (api *Api) CreateToken(request *models.Request) *models.Response {
	request.ParseJsonBody()
	userUUID := request.GetUserUUID()

	tokenName := request.GetParsedBodyAttribute("name")
	var tokenNameData string
	if tokenName == nil {
		tokenNameData = ""
	} else {
		tokenNameData = tokenName.(string)
	}
	if tokenNameData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Token name is required")
	}

	// Check token name is unique
	dbApiToken, err := api.app.Dao().GetTokenByName(userUUID, tokenNameData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if dbApiToken != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Token name already exists")
	}

	tokenUUID := uuid.NewV4()
	tokenSecretHMAC := hmac.New(sha256.New, []byte(api.app.Settings().APITokenSecretKey))
	tokenSecretHMAC.Write([]byte(tokenUUID.String()))
	tokenSecret := hex.EncodeToString(tokenSecretHMAC.Sum(nil))
	apiToken, err := api.app.Dao().CreateToken(userUUID, tokenUUID, tokenNameData, tokenSecret)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, apiToken, "Token created")
}

// DeleteToken godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete API Token.
//	@Description	Delete given API token of user
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/delete-token/{tokenUUID} [delete]
//	@Param			tokenUUID	path	string	true	"Token UUID"
func (api *Api) DeleteToken(request *models.Request) *models.Response {
	userUUID := request.GetUserUUID()
	tokenUUID := request.GetPathParam("tokenUUID")
	if tokenUUID == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Token UUID is required")
	}
	tokenUUIDData, err := uuid.FromString(tokenUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Invalid Token UUID")
	}
	validateUser, err := api.app.Dao().ValidateUserToken(userUUID, tokenUUIDData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if !validateUser {
		return models.NewErrorResponse(http.StatusUnauthorized, "Unauthorized")
	}
	err = api.app.Dao().DeleteToken(tokenUUIDData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Token deleted")
}

var UserSignUp ServiceFunc = (*Api).UserSignUp
var UserLogin ServiceFunc = (*Api).UserLogin
var UserVerifyEmail ServiceFunc = (*Api).UserVerifyEmail
var UserResendVerification ServiceFunc = (*Api).UserResendVerification
var UserVerifyResetPassword ServiceFunc = (*Api).UserVerifyResetPassword
var UserResetPassword ServiceFunc = (*Api).UserResetPassword
var UserForgotPassword ServiceFunc = (*Api).UserForgotPassword
var GetProfile ServiceFunc = (*Api).GetProfile
var GetProfileByHandle ServiceFunc = (*Api).GetProfileByHandle
var UpdateProfile ServiceFunc = (*Api).UpdateProfile
var DeleteProfile ServiceFunc = (*Api).DeleteProfile

var CreateSession ServiceFunc = (*Api).CreateSession
var GetSessionToken ServiceFunc = (*Api).GetSessionToken
var VerifySession ServiceFunc = (*Api).VerifySession

var GetTokens ServiceFunc = (*Api).GetTokens
var CreateToken ServiceFunc = (*Api).CreateToken
var DeleteToken ServiceFunc = (*Api).DeleteToken
