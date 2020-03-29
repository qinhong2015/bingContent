package client

import (
	"bytes"
	"fmt"
	"github.com/qinhong2015/bingContent/bingOAuth2"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

//bing content api base url
const BaseUrl = "https://content.api.bingads.microsoft.com/shopping";

//bing content api version
const ApiVersion = "v9.1"

//bing content api path
const ApiMerchantCenterPath = "bmc"

type Config struct {
	//Your application's registered client identifier.
	ClientId string
	//Your application's registered client secret.
	//This parameter is required with OAuthWebAuthCodeGrant requests.
	ClientSecret string
	//The URI where you want the authorization response to be redirected.
	Environment string
	//The client application's developer access token. Each request must include this header. For information about getting a token
	DeveloperToken string
	RequireLiveConnect bool
	RefreshToken string
	MerchantId string
	//The customer ID of the customer whose store you manage.
	CustomerId string
	//The account ID of any of the accounts that you manage on behalf of the customer
	AdsId string
}

type BingClient struct {
	OAuthTokens *oauth2.Token
	client *http.Client
	Config *Config
}

type BingRequest struct {
	Path string
	Method string
	Body []byte
	Queries map[string]string
}

func NewConfig() *Config {
	config := new(Config)
	config.Environment = bingOAuth2.EnvProduction
	config.RequireLiveConnect = false
	return config
}

func (c *BingClient) GetClient() *http.Client{
	if c.client == nil {
		c.client = &http.Client{}
	}
	return c.client
}

func (c *BingClient) RefreshAccessToken() (*oauth2.Token, error){
	if c.Config.RefreshToken == "" {
		err := &InvalidArgumentError{
			Message: "Refresh token is required to refresh access token",
		}
		return nil, err
	}

	oAuthWithAuthorizationCode := bingOAuth2.NewOAuthWithAuthorizationCode()
	oAuthWithAuthorizationCode.OAuthAuthorization.ClientSecret = c.Config.ClientSecret
	oAuthWithAuthorizationCode.OAuthAuthorization.ClientId = c.Config.ClientId
	oAuthWithAuthorizationCode.OAuthAuthorization.SetEnvironment(c.Config.Environment)
	if c.Config.RequireLiveConnect {
		oAuthWithAuthorizationCode.OAuthAuthorization.SetRequireLiveConnect(c.Config.RequireLiveConnect)
	}
	return oAuthWithAuthorizationCode.RequestOAuthTokensByRefreshToken(c.Config.RefreshToken)
}

func (c *BingClient) Request(request *BingRequest) ([] byte, error){
	endPoint := fmt.Sprintf(
		"%s/%s/%s/%s/%s",
		BaseUrl,
		ApiVersion,
		ApiMerchantCenterPath,
		c.Config.MerchantId,
		request.Path,
	)

	req, err := http.NewRequest(request.Method, endPoint, bytes.NewReader(request.Body))
	if err != nil {
		return nil, err
	}

	if c.OAuthTokens == nil || !c.OAuthTokens.Valid() {
		c.OAuthTokens, err = c.RefreshAccessToken()
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Accept", "application/json")

	if c.Config.CustomerId != "" {
		req.Header.Add("CustomerId", c.Config.CustomerId)
	}

	if c.Config.AdsId != "" {
		req.Header.Add("CustomerAccountId", c.Config.AdsId)
	}

	req.Header.Add("AuthenticationToken", c.OAuthTokens.AccessToken)
	req.Header.Add("DeveloperToken", c.Config.DeveloperToken)

	q := req.URL.Query()

	for key, value := range request.Queries {
		q.Add(key, value)
	}

	resp, err := c.GetClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		err = &UnauthorizedError{
			Response: resp,
			Body:     bodyBytes,
		}

		return nil, err
	}

	if resp.StatusCode >= 300 {
		err = &ApiError{
			Response: resp,
			Body:     bodyBytes,
		}

		return nil, err
	}

	return bodyBytes, nil
}