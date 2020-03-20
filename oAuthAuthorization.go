package bingContent

import "golang.org/x/oauth2"

type OAuthAuthorization struct {
	ClientId, RedirectUri, Environment, ClientSecret, State string
	RequireLiveConnect bool
	OAuthTokens *oauth2.Token
}

func (o *OAuthAuthorization) WithClientId(clientId string) *OAuthAuthorization {
	o.ClientId = clientId
	return o
}

func (o *OAuthAuthorization) WithClientSecret(clientSecret string) *OAuthAuthorization {
	o.ClientSecret = clientSecret
	return o
}

func (o *OAuthAuthorization) WithRedirectUri(redirectUrl string) *OAuthAuthorization {
	o.RedirectUri = redirectUrl
	return o
}

func (o *OAuthAuthorization) WithRefreshToken(refreshToken string) *OAuthAuthorization {
	token := &oauth2.Token {
		RefreshToken: refreshToken,
	}

	o.OAuthTokens = token
	return o
}

func (o *OAuthAuthorization) WithOAuthTokens(oAuthTokens *oauth2.Token) *OAuthAuthorization {
	o.OAuthTokens = oAuthTokens
	return o
}

func (o *OAuthAuthorization) WithEnvironment(environment string) *OAuthAuthorization {
	o.Environment = environment
	//o.RedirectUri = GetRedirectUrl(environment, o.RequireLiveConnect);
	return o
}

func (o *OAuthAuthorization) WithRequireLiveConnect(requireLiveConnect bool) *OAuthAuthorization {
	o.RequireLiveConnect = requireLiveConnect
	//o.RedirectUri = GetRedirectUrl(o.Environment, requireLiveConnect);
	return o;
}