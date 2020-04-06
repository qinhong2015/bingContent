package product

type ShippingWeight struct {
	//Gets or sets the unit of measure.
	Unit string `json:"unit"`
	//Gets or sets the item's weight, which is used to calculate the item's shipping cost.
	Value string `json:"value"`
}
