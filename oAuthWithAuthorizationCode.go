package bingContent

import (
	"golang.org/x/oauth2"
)

type OAuthWithAuthorizationCode struct {
	OAuthAuthorization *OAuthAuthorization
}

func (o *OAuthWithAuthorizationCode) RequestOAuthTokensByRefreshToken(refreshToken string) (*oauth2.Token, error) {
	OAuthRequestParameters := new(OAuthRequestParameters)
	OAuthRequestParameters.WithClientId(o.OAuthAuthorization.ClientId)
	OAuthRequestParameters.WithClientSecret(o.OAuthAuthorization.ClientSecret)
	OAuthRequestParameters.WithGrantType("refresh_token")
	OAuthRequestParameters.WithGrantParamName("refresh_token")
	OAuthRequestParameters.WithGrantValue(refreshToken)
	OAuthTokens, err := GetAccessTokens(OAuthRequestParameters, o.OAuthAuthorization.Environment, o.OAuthAuthorization.RequireLiveConnect);
	return OAuthTokens, err;
}

func NewOAuthWithAuthorizationCode() *OAuthWithAuthorizationCode {
	authorizationCode := new(OAuthWithAuthorizationCode)
	authorization := new(OAuthAuthorization)
	//default environment to production
	authorization.WithEnvironment(EnvProduction)
	authorization.WithRequireLiveConnect(false)
	authorizationCode.OAuthAuthorization = authorization
	return authorizationCode
}
