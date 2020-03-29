package catalog

import (
	"encoding/json"
	"fmt"
	"github.com/qinhong2015/bingContent/client"
	"net/http"
)

const ApiPath = "catalogs"

type Catalog struct {
	/**
	An ID that uniquely identifies the catalog in the store.
	This field is read-only; do not set this field.
	*/
	Id int `json:"id"`
	/**
	A Boolean value that determines whether the catalog is the store's default catalog.
	Is true if the catalog is the store's default catalog; otherwise, false.
	When you create a store, you get a default catalog that products are written to if you do not specify another catalog.
	This field is read-only; do not set this field.
	*/
	IsDefault bool `json:"isDefault"`
	/**
	A Boolean value that determines whether Microsoft may publish products from the catalog.
	Set to true if Microsoft may publish products from the catalog; otherwise, set it to false.
	You may update this field.
	*/
	IsPublishingEnabled bool `json:"isPublishingEnabled"`
	/**
	The market where products in the catalog are published to. The following are the possible markets that you may specify.
	*/
	Market string `json:"market"`
	/**
	The name of the store. The name may contain a maximum of 70 characters.
	You may update this field.
	*/
	Name string `json:"name"`
}

type Collection struct {
	Catalogs []Catalog `json:"catalogs"`
}

type Resource struct {
	Client *client.BingClient
}

//Use to get a catalog from the store.
func (r *Resource) Get(id int) (*Catalog, error) {
	path := fmt.Sprintf("%s/%d", ApiPath, id)
	queries := map[string]string{
		"alt": "json",
	}

	request := &client.BingRequest{
		Path:    path,
		Method:  http.MethodGet,
		Queries: queries,
	}

	bodyBytes, err := r.Client.Request(request)

	if err != nil {
		return nil, err
	}

	var catalog = new(Catalog)
	err = json.Unmarshal(bodyBytes, &catalog)
	if err != nil {
		return nil, err
	}

	return catalog, nil
}

//Use to get a list of catalogs from the store.
func (r *Resource) GetAll() (*Collection, error) {
	queries := map[string]string{
		"alt": "json",
	}

	request := &client.BingRequest{
		Path:    ApiPath,
		Method:  http.MethodGet,
		Queries: queries,
	}

	bodyBytes, err := r.Client.Request(request)

	if err != nil {
		return nil, err
	}

	var catalogCollection = new(Collection)
	err = json.Unmarshal(bodyBytes, &catalogCollection)
	if err != nil {
		return nil, err
	}

	return catalogCollection, nil
}

//Use to add a catalog to the store. To add a catalog, its name must be unique.
//You may add a maximum of 100 catalogs to a store.
func (r *Resource) Create(catalog *Catalog) (*Catalog, error) {
	queries := map[string]string{
		"alt": "json",
	}

	create := struct {
		IsPublishingEnabled bool `json:"isPublishingEnabled"`
		Market string `json:"market"`
		Name string `json:"name"`
	}{
		IsPublishingEnabled: catalog.IsPublishingEnabled,
		Market: catalog.Market,
		Name: catalog.Name,
	}

	body, err := json.Marshal(create)
	if err != nil {
		return nil, err
	}

	request := &client.BingRequest{
		Path:    ApiPath,
		Method:  http.MethodPost,
		Queries: queries,
		Body:    body,
	}

	bodyBytes, err := r.Client.Request(request)

	if err != nil {
		return nil, err
	}

	var catalogResponse = new(Catalog)
	err = json.Unmarshal(bodyBytes, &catalogResponse)
	if err != nil {
		return nil, err
	}

	return catalogResponse, nil
}

//Use to update a catalog in the store.
//The only fields you may update are the name and isPublishingEnabled fields, and you must specify both.
func (r *Resource) Update(catalog *Catalog) (*Catalog, error) {
	if catalog.Id == 0 {
		err := &client.InvalidArgumentError{
			Message: "catalog id is required",
		}
		return nil, err
	}

	queries := map[string]string{
		"alt": "json",
	}

	update := struct {
		IsPublishingEnabled bool `json:"isPublishingEnabled"`
		Name string `json:"name"`
	}{
		IsPublishingEnabled: catalog.IsPublishingEnabled,
		Name: catalog.Name,
	}

	body, err := json.Marshal(update)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%d", ApiPath, catalog.Id)

	request := &client.BingRequest{
		Path:    path,
		Method:  http.MethodPut,
		Queries: queries,
		Body:    body,
	}

	bodyBytes, err := r.Client.Request(request)

	if err != nil {
		return nil, err
	}

	var catalogResponse = new(Catalog)
	err = json.Unmarshal(bodyBytes, &catalogResponse)
	if err != nil {
		return nil, err
	}

	return catalogResponse, nil
}

//Use to delete a catalog from the store.
func (r *Resource) Delete(id int) error {
	path := fmt.Sprintf("%s/%d", ApiPath, id)
	queries := map[string]string{
		"alt": "json",
	}

	request := &client.BingRequest{
		Path:    path,
		Method:  http.MethodDelete,
		Queries: queries,
	}

	_, err := r.Client.Request(request)

	if err != nil {
		return err
	}

	return nil
}
