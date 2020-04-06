package product

type Warning struct {
	//For internal use only.
	Domain string `json:"domain"`
	//A description of the warning.
	Message string `json:"message"`
	//The reason why the offer generated a warning.
	//For example, you did not provide an identifier (gtin, mpn, or brand)
	//when the manufacturer is known to have assigned them.
	Reason string `json:"reason"`
}