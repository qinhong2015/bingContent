package bingContent

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type EndPointModel struct {
	RedirectUrl, OAuthTokenUrl, AuthorizationEndpointUrl, Scope string
}

var endPoints = map[string]EndPointModel{
	"ProductionMSIdentityV2": {
		RedirectUrl:              "https://login.microsoftonline.com/common/oauth2/nativeclient",
		OAuthTokenUrl:            "https://login.microsoftonline.com/common/oauth2/v2.0/token",
		AuthorizationEndpointUrl: "https://login.microsoftonline.com/common/oauth2/v2.0/authorize?scope=https://ads.microsoft.com/ads.manage offline_access",
		Scope:                    "https://ads.microsoft.com/ads.manage offline_access"},
	"ProductionLiveConnect": {
		RedirectUrl:              "https://login.live.com/oauth20_desktop.srf",
		OAuthTokenUrl:            "https://login.live.com/oauth20_token.srf",
		AuthorizationEndpointUrl: "https://login.live.com/oauth20_authorize.srf?scope=bingads.manage",
		Scope:                    "bingads.manage"},
	"SandboxLiveConnect": {
		RedirectUrl:              "https://login.live-int.com/oauth20_desktop.srf",
		OAuthTokenUrl:            "https://login.live-int.com/oauth20_token.srff",
		AuthorizationEndpointUrl: "https://login.live-int.com/oauth20_authorize.srf?scope=bingads.manage&prompt=login",
		Scope:                    "bingads.manage"},
}

func GetOAuthEndpointType(environment string, requireLiveConnect bool) string {
	if environment == EnvProduction {
		if requireLiveConnect {
			return ProductionLiveConnect
		} else {
			return ProductionMSIdentityV2
		}
	} else {
		return SandboxLiveConnect
	}
}

func GetRedirectUrl(environment string, requireLiveConnect bool) string {
	endpointType := GetOAuthEndpointType(environment, requireLiveConnect)
	redirectUrl := endPoints[endpointType].RedirectUrl
	return redirectUrl
}

func GetAuthTokenUrl(environment string, requireLiveConnect bool) string {
	endpointType := GetOAuthEndpointType(environment, requireLiveConnect)
	oAuthTokenUrl := endPoints[endpointType].OAuthTokenUrl
	return oAuthTokenUrl
}

func GetAuthorizeUrl(environment string, requireLiveConnect bool) string {
	endpointType := GetOAuthEndpointType(environment, requireLiveConnect)
	authorizationEndpointUrl := endPoints[endpointType].AuthorizationEndpointUrl
	return authorizationEndpointUrl
}

func GetAuthorizationEndpoint(oAuthUrlParameters *OAuthUrlParameters, environment string, requireLiveConnect bool) string {
	endpointType := GetOAuthEndpointType(environment, requireLiveConnect)
	baseUrl := endPoints[endpointType].AuthorizationEndpointUrl
	endPoint := fmt.Sprintf(
		"%s&client_id=%s&response_type=%s&redirect_uri=%s",
		baseUrl,
		oAuthUrlParameters.ClientId,
		oAuthUrlParameters.ResponseType,
		oAuthUrlParameters.RedirectUri)

	if requireLiveConnect {
		endPoint = endPoint + "&state=" + oAuthUrlParameters.State
	}
	return endPoint
}

func GetAccessTokens(oAuthRequestParameters *OAuthRequestParameters, environment string, requireLiveConnect bool) (*oauth2.Token, error) {
	var err error
	var tokenResponse = new(oauth2.Token)

	endpointType := GetOAuthEndpointType(environment, requireLiveConnect)
	scope := endPoints[endpointType].Scope
	tokenUrl := GetAuthTokenUrl(environment, requireLiveConnect)

	var accessTokenExchangeParams = map[string]string{
		"client_id":                           oAuthRequestParameters.ClientId,
		"grant_type":                          oAuthRequestParameters.GrantType,
		oAuthRequestParameters.GrantParamName: oAuthRequestParameters.GrantValue,
		"scope":                               scope,
	}

	if oAuthRequestParameters.RedirectUri != "" {
		accessTokenExchangeParams["redirect_uri"] = oAuthRequestParameters.RedirectUri
	}

	if oAuthRequestParameters.ClientSecret != "" {
		accessTokenExchangeParams["client_secret"] = oAuthRequestParameters.ClientSecret
	}

	resp, err := PostRequest(tokenUrl, accessTokenExchangeParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(bodyBytes, &tokenResponse)
		if err != nil {
			return nil, err
		}
	} else {
		err = &oauth2.RetrieveError{
			Response: resp,
			Body:     bodyBytes,
		}
		return nil, err
	}

	return tokenResponse, nil
}

func PostRequest(endPoint string, accessTokenExchangeParams map[string]string) (*http.Response, error) {
	data := url.Values{}
	for key, val := range accessTokenExchangeParams {
		data.Set(key, val)
	}

	req, err := http.NewRequest("POST", endPoint, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
