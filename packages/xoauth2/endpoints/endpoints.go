// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package endpoints provides constants for using OAuth2 to access various services.
package endpoints

import (
	"strings"

	"github.com/2637309949/dolphin/packages/xoauth2"
)

// Amazon is the endpoint for Amazon.
var Amazon = xoauth2.Endpoint{
	AuthURL:  "https://www.amazon.com/ap/oa",
	TokenURL: "https://api.amazon.com/auth/o2/token",
}

// Bitbucket is the endpoint for Bitbucket.
var Bitbucket = xoauth2.Endpoint{
	AuthURL:  "https://bitbucket.org/site/oauth2/authorize",
	TokenURL: "https://bitbucket.org/site/oauth2/access_token",
}

// Cern is the endpoint for CERN.
var Cern = xoauth2.Endpoint{
	AuthURL:  "https://oauth.web.cern.ch/OAuth/Authorize",
	TokenURL: "https://oauth.web.cern.ch/OAuth/Token",
}

// Facebook is the endpoint for Facebook.
var Facebook = xoauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/v3.2/dialog/oauth",
	TokenURL: "https://graph.facebook.com/v3.2/oauth/access_token",
}

// Foursquare is the endpoint for Foursquare.
var Foursquare = xoauth2.Endpoint{
	AuthURL:  "https://foursquare.com/oauth2/authorize",
	TokenURL: "https://foursquare.com/oauth2/access_token",
}

// Fitbit is the endpoint for Fitbit.
var Fitbit = xoauth2.Endpoint{
	AuthURL:  "https://www.fitbit.com/oauth2/authorize",
	TokenURL: "https://api.fitbit.com/oauth2/token",
}

// GitHub is the endpoint for Github.
var GitHub = xoauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

// GitLab is the endpoint for GitLab.
var GitLab = xoauth2.Endpoint{
	AuthURL:  "https://gitlab.com/oauth/authorize",
	TokenURL: "https://gitlab.com/oauth/token",
}

// Google is the endpoint for Google.
var Google = xoauth2.Endpoint{
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://xoauth2.googleapis.com/token",
}

// Heroku is the endpoint for Heroku.
var Heroku = xoauth2.Endpoint{
	AuthURL:  "https://id.heroku.com/oauth/authorize",
	TokenURL: "https://id.heroku.com/oauth/token",
}

// HipChat is the endpoint for HipChat.
var HipChat = xoauth2.Endpoint{
	AuthURL:  "https://www.hipchat.com/users/authorize",
	TokenURL: "https://api.hipchat.com/v2/oauth/token",
}

// Instagram is the endpoint for Instagram.
var Instagram = xoauth2.Endpoint{
	AuthURL:  "https://api.instagram.com/oauth/authorize",
	TokenURL: "https://api.instagram.com/oauth/access_token",
}

// KaKao is the endpoint for KaKao.
var KaKao = xoauth2.Endpoint{
	AuthURL:  "https://kauth.kakao.com/oauth/authorize",
	TokenURL: "https://kauth.kakao.com/oauth/token",
}

// LinkedIn is the endpoint for LinkedIn.
var LinkedIn = xoauth2.Endpoint{
	AuthURL:  "https://www.linkedin.com/oauth/v2/authorization",
	TokenURL: "https://www.linkedin.com/oauth/v2/accessToken",
}

// Mailchimp is the endpoint for Mailchimp.
var Mailchimp = xoauth2.Endpoint{
	AuthURL:  "https://login.mailchimp.com/oauth2/authorize",
	TokenURL: "https://login.mailchimp.com/oauth2/token",
}

// Mailru is the endpoint for Mail.Ru.
var Mailru = xoauth2.Endpoint{
	AuthURL:  "https://o2.mail.ru/login",
	TokenURL: "https://o2.mail.ru/token",
}

// MediaMath is the endpoint for MediaMath.
var MediaMath = xoauth2.Endpoint{
	AuthURL:  "https://api.mediamath.com/oauth2/v1.0/authorize",
	TokenURL: "https://api.mediamath.com/oauth2/v1.0/token",
}

// MediaMathSandbox is the endpoint for MediaMath Sandbox.
var MediaMathSandbox = xoauth2.Endpoint{
	AuthURL:  "https://t1sandbox.mediamath.com/oauth2/v1.0/authorize",
	TokenURL: "https://t1sandbox.mediamath.com/oauth2/v1.0/token",
}

// Microsoft is the endpoint for Microsoft.
var Microsoft = xoauth2.Endpoint{
	AuthURL:  "https://login.live.com/oauth20_authorize.srf",
	TokenURL: "https://login.live.com/oauth20_token.srf",
}

// NokiaHealth is the endpoint for Nokia Health.
var NokiaHealth = xoauth2.Endpoint{
	AuthURL:  "https://account.health.nokia.com/oauth2_user/authorize2",
	TokenURL: "https://account.health.nokia.com/oauth2/token",
}

// Odnoklassniki is the endpoint for Odnoklassniki.
var Odnoklassniki = xoauth2.Endpoint{
	AuthURL:  "https://www.odnoklassniki.ru/oauth/authorize",
	TokenURL: "https://api.odnoklassniki.ru/oauth/token.do",
}

// PayPal is the endpoint for PayPal.
var PayPal = xoauth2.Endpoint{
	AuthURL:  "https://www.paypal.com/webapps/auth/protocol/openidconnect/v1/authorize",
	TokenURL: "https://api.paypal.com/v1/identity/openidconnect/tokenservice",
}

// PayPalSandbox is the endpoint for PayPal Sandbox.
var PayPalSandbox = xoauth2.Endpoint{
	AuthURL:  "https://www.sandbox.paypal.com/webapps/auth/protocol/openidconnect/v1/authorize",
	TokenURL: "https://api.sandbox.paypal.com/v1/identity/openidconnect/tokenservice",
}

// Slack is the endpoint for Slack.
var Slack = xoauth2.Endpoint{
	AuthURL:  "https://slack.com/oauth/authorize",
	TokenURL: "https://slack.com/api/oauth.access",
}

// Spotify is the endpoint for Spotify.
var Spotify = xoauth2.Endpoint{
	AuthURL:  "https://accounts.spotify.com/authorize",
	TokenURL: "https://accounts.spotify.com/api/token",
}

// StackOverflow is the endpoint for Stack Overflow.
var StackOverflow = xoauth2.Endpoint{
	AuthURL:  "https://stackoverflow.com/oauth",
	TokenURL: "https://stackoverflow.com/oauth/access_token",
}

// Twitch is the endpoint for Twitch.
var Twitch = xoauth2.Endpoint{
	AuthURL:  "https://id.twitch.tv/oauth2/authorize",
	TokenURL: "https://id.twitch.tv/oauth2/token",
}

// Uber is the endpoint for Uber.
var Uber = xoauth2.Endpoint{
	AuthURL:  "https://login.uber.com/oauth/v2/authorize",
	TokenURL: "https://login.uber.com/oauth/v2/token",
}

// Vk is the endpoint for Vk.
var Vk = xoauth2.Endpoint{
	AuthURL:  "https://oauth.vk.com/authorize",
	TokenURL: "https://oauth.vk.com/access_token",
}

// Yahoo is the endpoint for Yahoo.
var Yahoo = xoauth2.Endpoint{
	AuthURL:  "https://api.login.yahoo.com/oauth2/request_auth",
	TokenURL: "https://api.login.yahoo.com/oauth2/get_token",
}

// Yandex is the endpoint for Yandex.
var Yandex = xoauth2.Endpoint{
	AuthURL:  "https://oauth.yandex.com/authorize",
	TokenURL: "https://oauth.yandex.com/token",
}

// AzureAD returns a new xoauth2.Endpoint for the given tenant at Azure Active Directory.
// If tenant is empty, it uses the tenant called `common`.
//
// For more information see:
// https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-v2-protocols#endpoints
func AzureAD(tenant string) xoauth2.Endpoint {
	if tenant == "" {
		tenant = "common"
	}
	return xoauth2.Endpoint{
		AuthURL:  "https://login.microsoftonline.com/" + tenant + "/oauth2/v2.0/authorize",
		TokenURL: "https://login.microsoftonline.com/" + tenant + "/oauth2/v2.0/token",
	}
}

// HipChatServer returns a new xoauth2.Endpoint for a HipChat Server instance
// running on the given domain or host.
func HipChatServer(host string) xoauth2.Endpoint {
	return xoauth2.Endpoint{
		AuthURL:  "https://" + host + "/users/authorize",
		TokenURL: "https://" + host + "/v2/oauth/token",
	}
}

// AWSCognito returns a new xoauth2.Endpoint for the supplied AWS Cognito domain which is
// linked to your Cognito User Pool.
//
// Example domain: https://testing.auth.us-east-1.amazoncognito.com
//
// For more information see:
// https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain.html
// https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-userpools-server-contract-reference.html
func AWSCognito(domain string) xoauth2.Endpoint {
	domain = strings.TrimRight(domain, "/")
	return xoauth2.Endpoint{
		AuthURL:  domain + "/oauth2/authorize",
		TokenURL: domain + "/oauth2/token",
	}
}
