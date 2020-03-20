package catalog

import "bingContent"

type Catalog struct {
	/**
	An ID that uniquely identifies the catalog in the store.
	This field is read-only; do not set this field.
	*/
	Id                  int64  `json:"id"`
	/**
	A Boolean value that determines whether the catalog is the store's default catalog.
	Is true if the catalog is the store's default catalog; otherwise, false.
	When you create a store, you get a default catalog that products are written to if you do not specify another catalog.
	This field is read-only; do not set this field.
	*/
	IsDefault           bool   `json:"isDefault"`
	/**
	A Boolean value that determines whether Microsoft may publish products from the catalog.
	Set to true if Microsoft may publish products from the catalog; otherwise, set it to false.
	You may update this field.
	*/
	IsPublishingEnabled bool   `json:"isPublishingEnabled"`
	/**
	The market where products in the catalog are published to. The following are the possible markets that you may specify.
	*/
	Market              string `json:"market"`
	/**
	The name of the store. The name may contain a maximum of 70 characters.
	You may update this field.
	*/
	Name                string `json:"name"`
}

type Resource struct {
	Client *bingContent.BingClient
	MerchantId string
}

func (r *Resource) Get(id int64) {

}
