package service

import "net/mail"

func ValidateMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func BaseEmailTemplate(title, body string) string {
	return `
	<!DOCTYPE html>
	<html
	  xmlns="http://www.w3.org/1999/xhtml"
	  xmlns:v="urn:schemas-microsoft-com:vml"
	  xmlns:o="urn:schemas-microsoft-com:office:office"
	>
	  <head>
		<title>` + title + `</title>
		<!--[if !mso]><!-->
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<!--<![endif]-->
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<style type="text/css">
		  #outlook a {
			padding: 0;
		  }
	
		  body {
			margin: 0;
			padding: 0;
			-webkit-text-size-adjust: 100%;
			-ms-text-size-adjust: 100%;
		  }
	
		  table,
		  td {
			border-collapse: collapse;
			mso-table-lspace: 0pt;
			mso-table-rspace: 0pt;
		  }
	
		  img {
			border: 0;
			height: auto;
			line-height: 100%;
			outline: none;
			text-decoration: none;
			-ms-interpolation-mode: bicubic;
		  }
	
		  p {
			display: block;
			margin: 13px 0;
		  }
		</style>
		<!--[if mso]>
		  <noscript>
			<xml>
			  <o:OfficeDocumentSettings>
				<o:AllowPNG />
				<o:PixelsPerInch>96</o:PixelsPerInch>
			  </o:OfficeDocumentSettings>
			</xml>
		  </noscript>
		<![endif]-->
		<!--[if lte mso 11]>
		  <style type="text/css">
			.mj-outlook-group-fix {
			  width: 100% !important;
			}
		  </style>
		<![endif]-->
		<!--[if !mso]><!-->
		<link
		  href="https://fonts.googleapis.com/css?family=Roboto:300,500"
		  rel="stylesheet"
		  type="text/css"
		/>
		<style type="text/css">
		  @import url(https://fonts.googleapis.com/css?family=Roboto:300,500);
		</style>
		<!--<![endif]-->
		<style type="text/css">
		  @media only screen and (min-width: 480px) {
			.mj-column-per-30 {
			  width: 30% !important;
			  max-width: 30%;
			}
	
			.mj-column-per-50 {
			  width: 50% !important;
			  max-width: 50%;
			}
	
			.mj-column-per-100 {
			  width: 100% !important;
			  max-width: 100%;
			}
	
			.mj-column-per-65 {
			  width: 65% !important;
			  max-width: 65%;
			}
	
			.mj-column-per-35 {
			  width: 35% !important;
			  max-width: 35%;
			}
		  }
		</style>
		<style media="screen and (min-width:480px)">
		  .moz-text-html .mj-column-per-30 {
			width: 30% !important;
			max-width: 30%;
		  }
	
		  .moz-text-html .mj-column-per-50 {
			width: 50% !important;
			max-width: 50%;
		  }
	
		  .moz-text-html .mj-column-per-100 {
			width: 100% !important;
			max-width: 100%;
		  }
	
		  .moz-text-html .mj-column-per-65 {
			width: 65% !important;
			max-width: 65%;
		  }
	
		  .moz-text-html .mj-column-per-35 {
			width: 35% !important;
			max-width: 35%;
		  }
		</style>
		<style type="text/css">
		  @media only screen and (max-width: 480px) {
			table.mj-full-width-mobile {
			  width: 100% !important;
			}
	
			td.mj-full-width-mobile {
			  width: auto !important;
			}
		  }
		</style>
	  </head>
	
	  <body style="word-spacing: normal; display: flex; justify-content: center">
		<div style="width: 70%; margin-top: 16px; margin-bottom: 16px">
		  <div
			style="
			  background-color: #ffffff;
			  border: 1px solid #e2e8f0;
			  border-radius: 8px;
			  margin-bottom: 32px;
			  padding: 16px;
			"
		  >
			<!--[if mso | IE]><table align="center" border="0" cellpadding="0" cellspacing="0" class="" style="width:600px;" width="600" ><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]-->
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
						padding-top: 32px;
						text-align: center;
					  "
					>
					  <!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:180px;" ><![endif]-->
					  <div
						class="mj-column-per-30 mj-outlook-group-fix"
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
								<table
								  border="0"
								  cellpadding="0"
								  cellspacing="0"
								  role="presentation"
								  style="
									border-collapse: collapse;
									border-spacing: 0px;
								  "
								>
								  <tbody>
									<tr>
									  <td style="width: 130px">
										<a
										  href="https://pureml.com"
										  target="_blank"
										>
										  <img
											height="auto"
											src="https://i.imgur.com/zpnB7l7.png"
											style="
											  border: 0;
											  display: block;
											  outline: none;
											  text-decoration: none;
											  height: auto;
											  width: 100%;
											  font-size: 13px;
											"
											width="130"
										  />
										</a>
									  </td>
									</tr>
								  </tbody>
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
						text-align: center;
					  "
					>
					  <!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:300px;" ><![endif]-->
					  <div
						class="mj-column-per-50 mj-outlook-group-fix"
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
								  padding-top: 30px;
								  word-break: break-word;
								"
							  >
								<div
								  style="
									font-family: Roboto, Helvetica, sans-serif;
									font-size: 28px;
									font-weight: 500;
									line-height: 24px;
									text-align: center;
									color: #616161;
								  "
								>` + title + `</div>
								</td>
							  </tr>
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
									  border-top: solid 2px #616161;
									  font-size: 1px;
									  margin: 0px auto;
									  width: 45%;
									"
								  ></p>
								  <!--[if mso | IE
									]><table
									  align="center"
									  border="0"
									  cellpadding="0"
									  cellspacing="0"
									  style="
										border-top: solid 2px #616161;
										font-size: 1px;
										margin: 0px auto;
										width: 112.5px;
									  "
									  role="presentation"
									  width="112.5px"
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
			  </div>` + body + `
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
								  align="left"
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
									  text-align: left;
									  color: #616161;
									"
								  >
									<p>
									  If you did not make this request, then
									  <i>please ignore this mail.</i>
									</p>
									<p>
									  If you run into any issues you can drop a
									  message on any of our social media platforms or
									  reach out to us on
									  <a href="mailto:contact.pureml@gmail.com"
										>contact.pureml@gmail.com</a
									  >. We look forward to serving you and your
									  business.
									</p>
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
								  align="left"
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
									  text-align: left;
									  color: #616161;
									"
								  >
									<p
									  style="
										font-weight: 400;
										color: #475569;
										line-height: 9px;
									  "
									>
									  Regards
									</p>
									<p
									  style="
										font-weight: 400;
										color: #475569;
										line-height: 9px;
									  "
									>
									  Team PureML
									</p>
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
			</div>
			<div style="background-color: #f8fafc; border-radius: 8px">
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
								  align="left"
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
									  text-align: left;
									  color: #616161;
									"
								  >
									<p>
									  You recieved this email because you just
									  requested for new password. If it's not you
									  contact us.
									</p>
									<p>©️ 2022 PureML. Inc | Texas, USA</p>
									<p>
									  Please visit our
									  <a href="https://pureml.com">website</a> for
									  more support.
									</p>
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
						  text-align: center;
						"
					  >
						<!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:390px;" ><![endif]-->
						<div
						  class="mj-column-per-65 mj-outlook-group-fix"
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
								  align="left"
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
									  border-collapse: collapse;
									  border-spacing: 0px;
									"
								  >
									<tbody>
									  <tr>
										<td style="width: 100px">
										  <a
											href="https://pureml.com"
											target="_blank"
										  >
											<img
											  height="auto"
											  src="https://i.imgur.com/zpnB7l7.png"
											  style="
												border: 0;
												display: block;
												outline: none;
												text-decoration: none;
												height: auto;
												width: 100%;
												font-size: 13px;
											  "
											  width="100"
											/>
										  </a>
										</td>
									  </tr>
									</tbody>
								  </table>
								</td>
							  </tr>
							</tbody>
						  </table>
						</div>
						<!--[if mso | IE]></td><td class="" style="vertical-align:top;width:210px;" ><![endif]-->
						<div
						  class="mj-column-per-35 mj-outlook-group-fix"
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
							width="100%"
						  >
							<tbody>
							  <tr>
								<td style="vertical-align: top; padding-top: 3px">
								  <table
									border="0"
									cellpadding="0"
									cellspacing="0"
									role="presentation"
									style=""
									width="100%"
								  >
									<tbody>
									  <tr>
										<td
										  align="left"
										  style="
											font-size: 0px;
											padding: 10px 25px;
											word-break: break-word;
										  "
										>
										  <table
											cellpadding="0"
											cellspacing="0"
											width="100%"
											border="0"
											style="
											  color: #000000;
											  font-family: Roboto, Helvetica,
												sans-serif;
											  font-size: 13px;
											  line-height: 22px;
											  table-layout: auto;
											  width: 100%;
											  border: none;
											"
										  >
											<tr
											  style="list-style: none; line-height: 1"
											>
											  <td>
												<a href="https://github.com/puremlhq">
												  <img
													width="20"
													src="https://i.imgur.com/8HnK5SF.png"
												  />
												</a>
											  </td>
											  <td>
												<a
												  href="https://twitter.com/puremlHQ"
												>
												  <img
													width="20"
													src="https://i.imgur.com/wVP2kYe.png"
												  />
												</a>
											  </td>
											  <td>
												<a
												  href="https://www.linkedin.com/company/pureml-inc/"
												>
												  <img
													width="18"
													src="https://i.imgur.com/JVZRF61.png"
												  />
												</a>
											  </td>
											  <td>
												<a
												  href="https://discord.gg/xNUHt9yguJ"
												>
												  <img
													width="20"
													src="https://i.imgur.com/uNRvJWW.png"
												  />
												</a>
											  </td>
											</tr>
										  </table>
										</td>
									  </tr>
									</tbody>
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
			  <!--[if mso | IE]></td></tr></table><![endif]-->
			</div>
		  </div>
		</body>
	  </html>	  
	`
}
