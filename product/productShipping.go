package product

type Shipping struct {
	//Gets or sets the two-letter ISO 3166 country code of the country where the item is being shipped to.
	Country string `json:"country"`
	//Gets or sets the location group name.
	LocationGroupName int `json:"locationGroupName"`
	//Gets or sets the ID of the geographic location where the item is being shipped to.
	//https://docs.microsoft.com/en-us/advertising/guides/geographical-location-codes?view=bingads-13
	LocationId string `json:"locationId"`
	//Gets or sets the postal code or postal code range of the location where the item is being shipped to.
	PostalCode string `json:"postalCode"`
	//Gets or sets the fixed price to ship the item to the specified location.
	Price Price `json:"price"`
	//Gets or sets geographical region where the item is being shipped to (for example, zip code).
	Region string `json:"region"`
	//Gets or sets a text description that describes the service class or delivery speed.
	Service string `json:"service"`
}
