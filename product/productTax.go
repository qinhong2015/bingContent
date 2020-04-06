package product

type Tax struct {
	//Gets or sets the country whose tax rate applies. Uses the two-letter ISO 3166 country code.
	Country string `json:"country"`
	//Gets or sets the ID of the geographic location where the item is being shipped to.
	//https://docs.microsoft.com/en-us/advertising/guides/geographical-location-codes?view=bingads-13
	LocationId int `json:"locationId"`
	//Gets or sets the postal code or postal code range of the location where the item is being shipped to.
	PostalCode string `json:"postalCode"`
	//Gets or sets the percentage tax rate to apply to the price of the item.
	//To specify a 5% rate, set this field to 5. To specify a 9.8% rate, set this field to 9.8.
	Rate float64 `json:"rate"`
	//Gets or sets a geographical region whose tax rate applies.
	Region string `json:"region"`
	//Gets or sets a Boolean value that determines whether to apply the tax to the cost of shipping.
	//Set to true if tax is charged on shipping.
	TaxShip string `json:"taxShip"`
}