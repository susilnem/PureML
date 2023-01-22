package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

const AuthHeaderName = "Authorization"

func AuthenticateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		authHeaderValue := extractRequestHeader(AuthHeaderName, context)
		if authHeaderValue == "" {
			context.Response().WriteHeader(http.StatusUnauthorized)
			context.Response().Writer.Write([]byte("Authentication Token Required"))
			return nil
		}
		if !strings.HasPrefix(authHeaderValue, "Bearer ") {
			context.Response().WriteHeader(http.StatusUnauthorized)
			context.Response().Writer.Write([]byte("Invalid Authentication Token"))
			return nil
		}
		authHeaderValue = strings.Split(authHeaderValue, " ")[1] //Splitting the bearer part?? yep
		token, _ := jwt.Parse(authHeaderValue, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token signing algorithm %v", t.Method.Alg())
			}
			return config.TokenSigningSecret(), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			context.Set("User", &models.UserClaims{
				UUID:   uuid.Must(uuid.FromString(claims["uuid"].(string))),
				Email:  claims["email"].(string),
				Handle: claims["handle"].(string),
			})
			next(context)
		} else {
			context.Response().WriteHeader(http.StatusForbidden)
			context.Response().Writer.Write([]byte("Invalid Authentication Token"))
		}
		return nil
	}
}
