package example

import (
	"bingContent/catalog"
)

func GetCatalog(id int) (*catalog.Catalog, error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return nil, err
	}
	return catalogResource.Get(id)
}

func GetCatalogs() (*catalog.Collection, error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return nil, err
	}
	return catalogResource.GetAll()
}

func DeleteCatalog(id int) ( error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return err
	}
	return catalogResource.Delete(id)
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

	return catalogResource.Create(c)
}

func UpdateCatalog() (*catalog.Catalog, error) {
	session := InitSessionWithRefreshToken()
	catalogResource, err := session.GetCatalogResource()
	if err != nil {
		return nil, err
	}

	c, _ := GetCatalog(1)
	c.IsPublishingEnabled = false

	return catalogResource.Create(c)
}