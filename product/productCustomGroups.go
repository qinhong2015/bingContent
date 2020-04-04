package product

type CustomGroup struct {
	//Gets or sets the name of the group.
	Name string `json:"name"`
	//Gets or sets the attributes for the group.
	Attributes CustomAttribute `json:"attributes"`
}