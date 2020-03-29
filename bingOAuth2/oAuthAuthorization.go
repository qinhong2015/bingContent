package bingOAuth2

import "golang.org/x/oauth2"

type OAuthAuthorization struct {
	ClientId, RedirectUri, ClientSecret, State string
	environment string
	requireLiveConnect bool
	OAuthTokens *oauth2.Token
}

func (o *OAuthAuthorization) GetEnvironment() string {
	return o.environment
}

func (o *OAuthAuthorization) SetEnvironment(environment string) *OAuthAuthorization {
	o.environment = environment
	o.RedirectUri = GetRedirectUrl(environment, o.requireLiveConnect);
	return o
}

func (o *OAuthAuthorization) IsRequireLiveConnect() bool {
	return o.requireLiveConnect;
}

func (o *OAuthAuthorization) SetRequireLiveConnect(requireLiveConnect bool) *OAuthAuthorization {
	o.requireLiveConnect = requireLiveConnect
	o.RedirectUri = GetRedirectUrl(o.environment, requireLiveConnect);
	return o;
}