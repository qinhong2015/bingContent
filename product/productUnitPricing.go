package product

type UnitPricing struct {
	//Gets or sets the unit of measure. For example, oz if the price is per ounce.
	Unit string `json:"unit"`
	//Gets or sets the price per unit.
	Value float64 `json:"value"`
}