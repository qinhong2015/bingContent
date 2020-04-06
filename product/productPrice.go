package product

type Price struct {
	//Gets or sets the currency that the price is stated in. Specify the currency using ISO 4217 currency codes.
	Currency string `json:"currency"`
	//Gets or sets the price of the item. Do not include currency symbols such as '$'.
	Value float64 `json:"value"`
}
