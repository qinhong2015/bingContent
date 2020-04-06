package product

import (
	"encoding/json"
	"fmt"
	"github.com/qinhong2015/bingContent/client"
	"net/http"
)

const ApiPath = "products"

type Product struct {
	//The URLs of additional images of the product that can be used in the product ad. To specify multiple images
	//MMC does not use the additional images; this field is included for Google compatibility.
	AdditionalImageLinks []string `json:"additionalImageLinks"`
	//A Boolean value that determines whether the item is an adult product.
	//Set to true if the item's target market is adults. The default is false.
	//Note that adult products are not supported and will be rejected.
	Adult string `json:"adult"`
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
	//The intended destinations of the product.
	Destinations []Destination `json:"destinations"`
	//The energy efficiency class as defined in EU directive 2010/30/EU.
	//https://www.eea.europa.eu/policy-documents/directive-2010-30-eu
	EnergyEfficiencyClass string `json:"energyEfficiencyClass"`
	//The UTC date and time that specifies when the product will expire
	//If you do not specify an expiration date, the product expires 30 days from the date and time that you add or update the product
	//Use this field to specify an expiration date that's less than 30 days from today
	ExpirationDate string `json:"expirationDate"`
	//The gender that the product targets.
	Gender string `json:"gender"`
	//The product category that the product is found in.
	//You may specify either a category string (for example, Animals & Pet Supplies > Pet supplies > Bird Supplies)
	//a category ID (for example, 3). For a category string, the list of subcategories is delimited by the greater than symbol ('>').
	//The field is limited to 255 characters.
	//https://advertiseonbing.blob.core.windows.net/blob/bingads/media/library/docs/bing-category-taxonomy-us.txt
	GoogleProductCategory string `json:"googleProductCategory"`
	//The Global Trade Item Number (GTIN) assigned by the manufacturer.
	Gtin string `json:"gtin"`
	//The fully qualified product id.
	//The ID is a composite of channel, contentLanguage, targetCountry, and offerId. The ID is case sensitive.
	Id string `json:"id"`
	//A Boolean value that determines whether the product offer specifies the gtin, mpn, or brand identifiers.
	IdentifierExists string `json:"identifierExists"`
	//The URL to an image of the product that can be used in the product ad. The URL is limited to 1,000 characters.
	ImageLink string `json:"imageLink"`
	//A Boolean value that determines whether the product is a merchant-defined bundle.
	IsBundle string `json:"isBundle"`
	//An ID that can be used to group all variants of the same product.
	ItemGroupId string `json:"item_group_id"`
	//The object's type. This field is set to content#product
	Kind string `json:"kind"`
	//The URL to the product's page on your website.
	//The URL is limited to 2,000 characters and may use the HTTP or HTTPS protocol.
	//The domain must match the store's domain.
	Link string `json:"link"`
	//A URL to a mobile-optimized version of the webpage that contains information about the product
	MobileLink string `json:"mobile_link"`
	//The number of identical products being sold as a single unit (for example, 4 flashlights).
	//When setting the price, it must be the total price of the multipack.
	Multipack int `json:"multipack"`
	//The manufacturer part number (MPN) of the product.
	//If the manufacturer assigns an MPN, you must specify it.
	//The MPN is limited to 70 characters.
	Mpn string `json:"mpn"`
	//The user-defined ID of the product being offered.
	//The offer ID is case insensitive and must be unique within a catalog, and is limited to a maximum of 50 characters.
	OfferId string `json:"offerId"`
	//A Boolean value that determines whether the product is only available for purchase online.
	//The value is true if the product is available only online. The default is false.
	OnlineOnly string `json:"onlineOnly"`
	//The pattern or graphic print of the product (for example, plaid). The pattern is limited to 100 characters.
	Pattern string `json:"pattern"`
	//The product's price. Specify the price in the currency of the target country.
	Price Price `json:"price"`
	//The advertiser-defined product category, which may be different from googleProductCategory.
	//For example, Animals & Pet Supplies > Pet supplies > Bird Supplies > Veterinary.
	ProductType string `json:"productType"`
	//A comma-delimited list of IDs that identify promotions in your Promotions feed.
	//You may specify a maximum of 10 promotion IDs.
	PromotionId string `json:"promotionId"`
	//The item's sale price. The sale price must be in the range 0.01 (1 cent) through 10000000.00 (10 million).
	SalePrice Price `json:"salePrice"`
	//The sale's UTC start and end date. Specify the dates in ISO 8601 format.
	//For example, 2016-04-05T08:00-08:00/2016-04-10T19:30-08:00
	SalePriceEffectiveDate string `json:"salePriceEffectiveDate"`
	//The name of the merchant that's selling the product. Used only by aggregators to identify the merchant.
	//Aggregators are third party sites that behave on behalf of individual merchants.
	SellerName string `json:"sellerName"`
	//The price to ship the product based on location.
	Shipping Shipping `json:"shipping"`
	//The shipping label.
	ShippingLabel string `json:"shippingLabel"`
	//The weight of the product. The weight is used for shipping purposes.
	ShippingWeight ShippingWeight `json:"shippingWeight"`
	//The available sizes of the product.
	//The size value is user-defined but should be based on your target country.
	Sizes []string `json:"sizes"`
	//The measurement system used to size the product.
	//For example, shoes may be sized using the US system or the UK system.
	SizeSystem string `json:"sizeSystem"`
	//The cut of the product.
	SizeType string `json:"sizeType"`
	//The two-letter ISO 3166 country code of the target country (the country where you want to advertise the product).
	//The country must match the market specified by the catalog.
	TargetCountry string `json:"targetCountry"`
	//The tax information of the product.
	//MMC does not use this field; it's included for Google compatibility.
	Taxes []Tax `json:"taxes"`
	//The product's title (for example, Women's Shoes). The title may not include promotional text.
	//The title is limited to a maximum of 150 characters, and may include any Unicode character.
	Title string `json:"title"`
	//The product’s base measure for pricing (for example, 100ml means the price is calculated based on a 100ml unit).
	UnitPricingBaseMeasure UnitPricing `json:"unitPricingBaseMeasure"`
	//The measure and dimension of product as it is sold.
	UnitPricingMeasure UnitPricing `json:"unitPricingMeasure"`
	//The read-only list of intended destinations that have passed validation.
	//MMC does not use this field; it's included for Google compatibility.
	ValidatedDestinations []string `json:"validatedDestinations"`
	//A list of warnings about issues with the product offer.
	Warnings []Warning `json:"warnings"`
}

type Collection struct {
	//Gets the object's type. This field is set to content#productsListResponse.
	Kind string `json:"kind"`
	//Gets the token used to get the next page of results.
	NextPageToken string `json:"nextPageToken"`
	//Gets the list of products. If the catalog does not contain any offers, the array is empty.
	Resources []Product `json:"resources"`
}

type Parameter struct {
	//Optional. Use to specify the catalog to insert (update) product offerings into.
	CatalogId string
	//json or xml
	Alt string
	//Optional. Use when debugging your application to test calls.
	//Calls that include this parameter won't affect production data (products are not inserted or deleted);
	//however, the response will contain any errors that the call generates.
	DryRun bool
	//Optional. Use to specify the maximum number of items to return in a List request.
	//The maximum value that you may specify is 250. The default is 25.
	MaxResults int
	//Optional. Use to page through a store's list of products.
	//The token identifies the next page of products to return in a List request.
	//Do not specify this parameter in the first List request.
	StartToken string
}

func (r *Resource) NewParameter() *Parameter {
	p := new(Parameter)
	p.Alt = "json"
	p.DryRun = false
	return p
}

func (r *Resource) getQueries(p *Parameter) map[string]string{
	queries := make(map[string]string)

	queries["alt"] = p.Alt

	if p.CatalogId != "" {
		queries["bmc-catalog-id"] = p.CatalogId
	}

	if p.DryRun {
		queries["dry-run"] = "true"
	}

	if p.MaxResults > 0 {
		queries["max-results"] = fmt.Sprintf("%d", p.MaxResults)
	}

	if p.StartToken != "" {
		queries["start-token"] = p.StartToken
	}

	return queries
}

type Resource struct {
	Client *client.BingClient
}

func (r *Resource) Get(id string, p *Parameter) (*Product, error) {
	path := fmt.Sprintf("%s/%s", ApiPath, id)
	queries := r.getQueries(p)

	request := &client.BingRequest{
		Path:    path,
		Method:  http.MethodGet,
		Queries: queries,
	}

	bodyBytes, err := r.Client.Request(request)

	if err != nil {
		return nil, err
	}

	var product = new(Product)
	err = json.Unmarshal(bodyBytes, &product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *Resource) GetAll(p *Parameter) (*Collection, error) {
	queries := r.getQueries(p)

	request := &client.BingRequest{
		Path:    ApiPath,
		Method:  http.MethodGet,
		Queries: queries,
	}

	bodyBytes, err := r.Client.Request(request)

	if err != nil {
		return nil, err
	}

	var collection = new(Collection)
	err = json.Unmarshal(bodyBytes, &collection)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func (r *Resource) Create(product *Product, p *Parameter) (*Product, error) {
	queries := r.getQueries(p)

	body, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	request := &client.BingRequest{
		Path:    ApiPath,
		Method:  http.MethodPost,
		Queries: queries,
		Body:    body,
	}

	bodyBytes, err := r.Client.Request(request)

	if err != nil {
		return nil, err
	}

	var productResponse = new(Product)
	err = json.Unmarshal(bodyBytes, &productResponse)
	if err != nil {
		return nil, err
	}

	return productResponse, nil
}