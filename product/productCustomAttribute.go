package product

type CustomAttribute struct {
	//Gets or sets the attribute's name.
	Name string `json:"name"`
	//Gets or sets the attribute's type. The following are the possible values.
	Type string `json:"type"`
	//Gets or sets the attribute's unit of measure. Used for values of type INT and FLOAT only.
	Unit string `json:"unit"`
	//Gets or sets the attribute's value.
	Value string `json:"value"`
}