package product

const ApiPath = "products"

type Product struct {
	//The URLs of additional images of the product that can be used in the product ad. To specify multiple images
	//MMC does not use the additional images; this field is included for Google compatibility.
	AdditionalImageLinks []string `json:"additionalImageLinks"`
	//A Boolean value that determines whether the item is an adult product.
	//Set to true if the item's target market is adults. The default is false.
	//Note that adult products are not supported and will be rejected.
	Adult bool `json:"adult"`
	//A group of items for Cost-per-acquisition (CPA) bidding.
	//MMC does not use the additional images; this field is included for Google compatibility.
	AdwordsGrouping string `json:"adwordsGrouping"`
	//The labels for grouped items (see adwordsGrouping). Applies to Cost-per-click (CPC) only.
	//MMC does not use the additional images; this field is included for Google compatibility.
	AdwordsLabels []string `json:"adwordsLabels"`
	//The URL to use in the product ad. If specified, this URL must redirect to the URL specified in link.
	AdwordsRedirect string `json:"adwordsRedirect"`
	//The target age group of the item.
	AgeGroup string `json:"ageGroup"`
	//The availability status of the product. The following are the possible values.
	Availability string `json:"availability"`
	//The UTC date that a pre-order product will be available for shipping (see the availability field).
	//This field is optional but if you know the date that the pre-ordered product will be available for shipping,
	//you should set this field. Specify the date in ISO 8601 format.
	//NOTE: MMC currently ignores the contents of this field.
	AvailabilityDate string `json:"availabilityDate"`
	//The item's brand, manufacturer, or publisher.
	//The string may contain a maximum of 10 words and 1,000 characters.
	//To ensure the string displays well in the UX, you should limit the brand name to no more than 70 characters.
	Brand string `json:"brand"`
	//The sales channel for the product. The following are the possible case-insensitive values.
	//Because channel is used to create the product ID, you may not change this field after adding the product to the store.
	Channel string `json:"channel"`
	//The product's dominant color. If the color is a blend of colors,
	//you may specify a slash-delimited list of up to 3 colors (for example, red/green/blue).
	Color string `json:"color"`
	//The condition of the product. The following are the possible values.
	//Default is new
	Condition string `json:"condition"`
	//The two-letter ISO 639-1 language code for the product.
	//Because the language is used to create the product ID, you may not change this field after adding the product to the store.
	ContentLanguage string `json:"content_language"`
	//A list of custom attributes used by the merchant.
	CustomAttributes []CustomAttribute `json:"customAttributes"`
	//A list of custom groups used by the merchant.
	CustomGroups []CustomGroup `json:"customGroups"`
	//Custom label 0, which is used to filter products for Microsoft Shopping campaigns.
	//The label is limited to 100 characters.
	CustomLabel0 string `json:"customLabel0"`
	//Custom label 1, which is used to filter products for Microsoft Shopping campaigns.
	//The label is limited to 100 characters.
	CustomLabel1 string `json:"customLabel1"`
	//Custom label 2, which is used to filter products for Microsoft Shopping campaigns.
	//The label is limited to 100 characters.
	CustomLabel2 string `json:"customLabel2"`
	//Custom label 0, which is used to filter products for Microsoft Shopping campaigns.
	//The label is limited to 100 characters.
	CustomLabel3 string `json:"customLabel3"`
	//Custom label 0, which is used to filter products for Microsoft Shopping campaigns.
	//The label is limited to 100 characters.
	CustomLabel4 string `json:"customLabel4"`
	//A description of the product. The description may not include promotional text.
	//The description is limited to a maximum of 10,000 characters, and may include any Unicode characters.
	Description string `json:"description"`
}
