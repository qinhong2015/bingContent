package bingContent

type OAuthRequestParameters struct {
	//Your application's registered client identifier.
	ClientId string
	//Your application's registered client secret.
	//This parameter is required with OAuthWebAuthCodeGrant requests.
	ClientSecret string
	//The URI where you want the authorization response to be redirected.
	RedirectUri string
	//The authorization grant param name.
	//For example the grant param name could be 'refresh_token' or 'authorization_code'.
	GrantType string
	//The authorization grant param name.
	//For example the grant param name could be 'refresh_token' or 'code'.
	GrantParamName string
	//The value depends on the $GrantType and $GrantParamName.
	//For example if $GrantType and $GrantParamName are both set to 'refresh_token',
	//the value is equal to the last known and valid refresh token.
	GrantValue string
}

func (o *OAuthRequestParameters) WithClientId(clientId string) *OAuthRequestParameters {
	o.ClientId = clientId
	return o
}

func (o *OAuthRequestParameters) WithClientSecret(clientSecret string) *OAuthRequestParameters {
	o.ClientSecret = clientSecret
	return o
}

func (o *OAuthRequestParameters) WithRedirectUri(redirectUri string) *OAuthRequestParameters {
	o.RedirectUri = redirectUri
	return o
}

func (o *OAuthRequestParameters) WithGrantType(grantType string) *OAuthRequestParameters {
	o.GrantType = grantType
	return o
}

func (o *OAuthRequestParameters) WithGrantParamName(grantParamName string) *OAuthRequestParameters {
	o.GrantParamName = grantParamName
	return o
}

func (o *OAuthRequestParameters) WithGrantValue(grantValue string) *OAuthRequestParameters {
	o.GrantValue = grantValue
	return o
}