package bingOAuth2

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