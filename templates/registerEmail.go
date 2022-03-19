package templates

import (
	"strings"
)

func RegisterEmail(name string) string {

  text := `
  <html lang="en-US">
  <head>
      <meta content="text/html; charset=utf-8" http-equiv="Content-Type" />
      <title>Reset Password Email Template</title>
      <meta name="description" content="Reset Password Email Template.">
      <style type="text/css">
          a:hover {text-decoration: underline !important;}
      </style>
  </head>
  
  <body marginheight="0" topmargin="0" marginwidth="0" style="margin: 0px; background-color: #f2f3f8;" leftmargin="0">
      <table cellspacing="0" border="0" cellpadding="0" width="100%" bgcolor="#f2f3f8"
          style="@import url(https://fonts.googleapis.com/css?family=Rubik:300,400,500,700|Open+Sans:300,400,600,700); font-family: 'Open Sans', sans-serif;">
          <tr>
              <td>
                  <table style="background-color: #f2f3f8; max-width:670px;  margin:0 auto;" width="100%" border="0"
                      align="center" cellpadding="0" cellspacing="0">
                      <tr>
                          <td style="height:80px;">&nbsp;</td>
                      </tr>
                      <tr>
                          <td style="text-align:center;">
                            <a href="https://rakeshmandal.com" title="logo" target="_blank">
                              <img width="200" src="https://gamelyd.co/images/logos/logo.png" title="logo" alt="logo">
                            </a>
                          </td>
                      </tr>
                      <tr>
                          <td style="height:20px;">&nbsp;</td>
                      </tr>
                      <tr>
                          <td>
                              <table width="95%" border="0" align="center" cellpadding="0" cellspacing="0"
                                  style="max-width:670px;background:#fff; border-radius:3px; text-align:center;-webkit-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);-moz-box-shadow:0 6px 18px 0 rgba(0,0,0,.06);box-shadow:0 6px 18px 0 rgba(0,0,0,.06);">
                                  <tr>
                                      <td style="height:40px;">&nbsp;</td>
                                  </tr>
                                  <tr>
                                      <td style="padding:0 35px;">
                                          <h1 style="color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;">Welcome to gamelyd</h1>
                                          <span
                                              style="display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;"></span>
                                          <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                                             Hi [[name]]
                                          </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	This is really exciting: welcome (officially) to your service here, we are so lucky to have you.
									   	  </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	We are here to make sure you get the maximum gaming experience you expect from using this plartform, so dont hesitate to reach out with questions, we would love to hear from you.
									   	  </p>

											<p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
											 If you need support you can email us on support@gamelyd.com, we can talk you through the information you need to get started.
											</p>

											<p style="color:#455056; font-size:15px;line-height:24px; margin:0;">Looking forward to hearing from you,</p>
									   	  </p>
                                      </td>
                                  </tr>
                                  <tr>
                                      <td style="height:40px;">&nbsp;</td>
                                  </tr>
                              </table>
                          </td>
                      <tr>
                          <td style="height:20px;">&nbsp;</td>
                      </tr>
                      <tr>
                      <td style="text-align: center">
                        <p
                          style="
                            font-size: 14px;
                            color: rgba(69, 80, 86, 0.7411764705882353);
                            line-height: 18px;
                            margin: 0 0 0;
                          "
                        >
                          <stron>Gamelyd is an online plartform that helps in organizing game
                            tournaments in diffrent modes, Gamelyd simplifies and makes it easy for
                            people in diffrent locations to meet and play tournaments together.</strong>
                        </p>
                        <p style="
                          font-size: 14px;
                          color: rgba(69, 80, 86, 0.7411764705882353);
                          line-height: 18px;
                          margin: 0 0 0;
                        ">If you did not initiate this email please contact us on contact@gamelyd.com or
                        visit our<a href="https://gamelyd.co"> website</a></p>
                        <div style="
        
                        margin: 20px 0;
                      "></div>
                        <p
                          style="
                            font-size: 14px;
                            color: rgba(69, 80, 86, 0.7411764705882353);
                            line-height: 18px;
                            margin: 0 0 0;
                          "
                        >
                          &copy; <strong>www.gamelyd.com</strong>
                        </p>
                      </td>
                    </tr>
                      <tr>
                          <td style="height:80px;">&nbsp;</td>
                      </tr>
                  </table>
              </td>
          </tr>
      </table>
  </body>
  
  </html>`
 
  replace := "[[name]]"
  newValue := name
  n := 1
  text = strings.Replace(text, replace, newValue, n)
  return text
}
