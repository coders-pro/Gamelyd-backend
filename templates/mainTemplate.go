package templates

import (
	"strings"
)

func PasswordChanged(name string) string {

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
                                          <h1 style="color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;">Password Changed</h1>
                                          <span
                                              style="display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;"></span>
                                          <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                                             Hi [[name]]
                                          </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	This is to inform you that your password was recently changed in our plartform.
									   	  </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	If your are not the one that initiated the request please quickly reach out to us, becuase your account might have been compromised
									   	  </p>

                         <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">Thank You</p>

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
                          &copy; <strong>www.gamelyd.co</strong>
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

func CreateTournament(name string) string {

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
                                          <h1 style="color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;">Tournament Created</h1>
                                          <span
                                              style="display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;"></span>
                                          <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                                             Hi [[name]]
                                          </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	Thank you for choosing Gamelyd LTD. we are here to serve you more
									   	  </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	Your tournament was successfully created, you can reach out to us on support@gamelyd.com if you are facing challanges with your tournament
									   	  </p>

                         <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">Thank You</p>

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
                          &copy; <strong>www.gamelyd.co</strong>
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


  func RegisterTournament(name string, tournamentName string, teamName string, date string, id string) string {

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
                                          <h1 style="color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;">Registered</h1>
                                          <span
                                              style="display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;"></span>
                                          <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                                             Hi [[name]]
                                          </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	Your account was just registered for [[tournyName]] tournament under [[teamName]]
                        Team, which should hold on [[date]]. you can visit the tournament page for more details.
									   	  </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	Note: you can cancel your registration by clicking the unregister button in the tournament page
									   	  </p>

                         <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">Thank You</p>

                         <a href="https://gamelyd.co/tournament/view/[[id]]"
                          style="background:#0BC0B4;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;">
                          View Tournament</a>

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
                          &copy; <strong>www.gamelyd.co</strong>
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

	text = strings.Replace(text, "[[tournyName]]", tournamentName, n)
	text = strings.Replace(text, "[[teamName]]", teamName, n)
	text = strings.Replace(text, "[[date]]", date, n)
	text = strings.Replace(text, "[[id]]", id, n)
	return text
  }


  func DrawTournament(name string, tournamentName string, id string) string {

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
                                          <h1 style="color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;">Tournament Draw</h1>
                                          <span
                                              style="display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;"></span>
                                          <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                                             Hi [[name]]
                                          </p>
										  <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
										  	A tournament you registered for has been drawn for next stage you can visit the tournament page to view statistics and your next match
									   	  </p>


                         <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">Thank You</p>

                         <a href="https://gamelyd.co/tournament/view/[[id]]"
                          style="background:#0BC0B4;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;">
                          View Tournament</a>

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
                          &copy; <strong>www.gamelyd.co</strong>
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

  </html>
` 
   
	replace := "[[name]]"
	newValue := name
	n := 1
	text = strings.Replace(text, replace, newValue, n)

	text = strings.Replace(text, "[[tournyName]]", tournamentName, n)
	text = strings.Replace(text, "[[id]]", id, n)
	return text
  }


  func AddTime(name string, tournamentName string, id string, time string, date string) string {

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
                                            <h1 style="color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;">Match Time</h1>
                                            <span
                                                style="display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;"></span>
                                            <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                                               Hi [[name]]
                                            </p>
                        <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                          Your next match has been scheduled for [[date]] [[time]], you can visit the tournament page to view all statistics.
                           </p>
  
  
                           <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">Thank You</p>
  
                           <a href="https://gamelyd.co/tournament/view/[[id]]"
                            style="background:#0BC0B4;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;">
                            View Tournament</a>
  
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
                            &copy; <strong>www.gamelyd.co</strong>
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
  
    </html>
  ` 
     
    replace := "[[name]]"
    newValue := name
    n := 1
    text = strings.Replace(text, replace, newValue, n)
  
    text = strings.Replace(text, "[[tournyName]]", tournamentName, n)
    text = strings.Replace(text, "[[id]]", id, n)
    text = strings.Replace(text, "[[time]]", time, n)
    text = strings.Replace(text, "[[date]]", date, n)
    return text
    }

    func AddLink(name string, tournamentName string, id string) string {

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
                                              <h1 style="color:#1e1e2d; font-weight:500; margin:0;font-size:32px;font-family:'Rubik',sans-serif;">Match Link</h1>
                                              <span
                                                  style="display:inline-block; vertical-align:middle; margin:29px 0 26px; border-bottom:1px solid #cecece; width:100px;"></span>
                                              <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                                                 Hi [[name]]
                                              </p>
                          <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">
                            Link to your next match has been added to draw, you can visit the tournament page to join match.
                             </p>
    
    
                             <p style="color:#455056; font-size:15px;line-height:24px; margin:0;">Thank You</p>
    
                             <a href="https://gamelyd.co/tournament/view/[[id]]"
                              style="background:#0BC0B4;text-decoration:none !important; font-weight:500; margin-top:35px; color:#fff;text-transform:uppercase; font-size:14px;padding:10px 24px;display:inline-block;border-radius:50px;">
                              View Tournament</a>
    
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
                              &copy; <strong>www.gamelyd.co</strong>
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
    
      </html>
    ` 
       
      replace := "[[name]]"
      newValue := name
      n := 1
      text = strings.Replace(text, replace, newValue, n)
    
      text = strings.Replace(text, "[[tournyName]]", tournamentName, n)
      text = strings.Replace(text, "[[id]]", id, n)

      return text
      }


  
