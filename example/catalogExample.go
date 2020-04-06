package example

import (
	"github.com/qinhong2015/bingContent/catalog"
)

func GetCatalog(id int) (*catalog.Catalog, error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return nil, err
	}
	return catalogResource.Get(id, catalogResource.NewParameter())
}

func GetCatalogStatus(id int) (*catalog.Status, error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return nil, err
	}
	return catalogResource.GetStatus(id, catalogResource.NewParameter())
}

func GetCatalogs() (*catalog.Collection, error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return nil, err
	}
	return catalogResource.GetAll(catalogResource.NewParameter())
}

func DeleteCatalog(id int) error {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return err
	}
	return catalogResource.Delete(id, catalogResource.NewParameter())
}

func CreateCatalog() (*catalog.Catalog, error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return nil, err
	}

	c := &catalog.Catalog{
		IsPublishingEnabled: true,
		Market:              "en-CA",
		Name:                "Name",
	}

	return catalogResource.Create(c, catalogResource.NewParameter())
}

func UpdateCatalog() (*catalog.Catalog, error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return nil, err
	}

	c, _ := GetCatalog(1)
	c.IsPublishingEnabled = false

	return catalogResource.Create(c, catalogResource.NewParameter())
}
