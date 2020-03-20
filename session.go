package bingContent

import (
	"bingContent/catalog"
	"golang.org/x/oauth2"
)

type config struct {
	//Your application's registered client identifier.
	ClientId string
	//Your application's registered client secret.
	//This parameter is required with OAuthWebAuthCodeGrant requests.
	ClientSecret string
	//The URI where you want the authorization response to be redirected.
	Environment string
	RequireLiveConnect bool
	RefreshToken string
	MerchantId string
}

type Session struct {
	config *config
	client *BingClient
}

func NewSession(c *config) *Session{
	session := new(Session)
	session.config = c
	return session
}

func (s *Session) GetClient() (*BingClient, error){
	if s.client == nil {
		token, err := s.RefreshAccessToken()

		if err == nil {
			return nil, err
		}

		s.client = &BingClient{
			BaseUri:     "",
			Headers:     nil,
			Query:       nil,
			RequestBody: nil,
			OAuthTokens: token,
		}
	}
	//TODO handle expired access token
	return s.client, nil
}

func (s *Session) RefreshAccessToken() (*oauth2.Token, error){
	if s.config.RefreshToken == "" {
		err := &InvalidArgumentError{
			message:"Refresh token is required to refresh access token",
		}
		return nil, err
	}

	oAuthWithAuthorizationCode := NewOAuthWithAuthorizationCode()
	oAuthWithAuthorizationCode.OAuthAuthorization.WithClientSecret(s.config.ClientSecret)
	oAuthWithAuthorizationCode.OAuthAuthorization.WithClientId(s.config.ClientId)
	oAuthWithAuthorizationCode.OAuthAuthorization.WithEnvironment(s.config.Environment)
	if s.config.RequireLiveConnect == true {
		oAuthWithAuthorizationCode.OAuthAuthorization.WithRequireLiveConnect(true)
	}
	return oAuthWithAuthorizationCode.RequestOAuthTokensByRefreshToken(s.config.RefreshToken)
}

func (s *Session) GetCatalogResource() (*catalog.Resource, error){
	client, err := s.GetClient()
	if err == nil {
		return nil, err
	}

	resource := &catalog.Resource{
		Client: client,
		MerchantId: s.config.MerchantId,
	}

	return resource, nil
}