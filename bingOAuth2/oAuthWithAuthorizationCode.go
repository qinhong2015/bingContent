package bingOAuth2

import (
	"golang.org/x/oauth2"
)

type OAuthWithAuthorizationCode struct {
	OAuthAuthorization *OAuthAuthorization
}

func (o *OAuthWithAuthorizationCode) RequestOAuthTokensByRefreshToken(refreshToken string) (*oauth2.Token, error) {
	OAuthRequestParameters := &OAuthRequestParameters {
		ClientId: o.OAuthAuthorization.ClientId,
		ClientSecret: o.OAuthAuthorization.ClientSecret,
		GrantType: "refresh_token",
		GrantParamName: "refresh_token",
		GrantValue: refreshToken,
	}
	OAuthTokens, err := GetAccessTokens(
		OAuthRequestParameters,
		o.OAuthAuthorization.GetEnvironment(),
		o.OAuthAuthorization.IsRequireLiveConnect(),
	);
	return OAuthTokens, err;
}

func NewOAuthWithAuthorizationCode() *OAuthWithAuthorizationCode {
	authorizationCode := new(OAuthWithAuthorizationCode)
	authorization := new(OAuthAuthorization)
	//default environment to production
	authorization.SetEnvironment(EnvProduction)
	//default require live connect to false
	authorization.SetRequireLiveConnect(false)
	authorizationCode.OAuthAuthorization = authorization
	return authorizationCode
}
