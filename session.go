package bingContent

import (
	"github.com/qinhong2015/bingContent/catalog"
	bingClient "github.com/qinhong2015/bingContent/client"
	"github.com/qinhong2015/bingContent/product"
)

type Session struct {
	config *bingClient.Config
	client *bingClient.BingClient
}

func NewSession(c *bingClient.Config) *Session{
	session := new(Session)
	session.config = c
	return session
}

func (s *Session) GetClient() (*bingClient.BingClient, error){
	if s.client == nil {
		s.client = &bingClient.BingClient{
			Config: s.config,
		}
	}
	return s.client, nil
}

func (s *Session) GetCatalogResource() (*catalog.Resource, error){
	client, err := s.GetClient()
	if err != nil {
		return nil, err
	}

	resource := &catalog.Resource{
		Client: client,
	}

	return resource, nil
}

func (s *Session) GetProductResource() (*product.Resource, error){
	client, err := s.GetClient()
	if err != nil {
		return nil, err
	}

	resource := &product.Resource{
		Client: client,
	}

	return resource, nil
}