package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/boutique-arthur.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
id: FECFILDEC-ROUGE-43-46,
title: Chaussettes fil d'ecosse rouge,
link: https://jrt.boutique-arthur.com/?P4FF454598317S1UB412d5909360V4,
price: 15.00,
description: Chaussette unie 39/42 / 43/46 – 100 % coton égyptien, fil d'EcosseJE ME JETTE A LOT: 3 Chaussettes achetées, la 4ème OFFERTE Ajoutez 4 chaussettes à votre panier pour bénéficier automatiquement de la gratuité sur le 4ème article,
condition: new,
gtin: 3369692658805,
brand: Arthur,
mpn: MPN_FECFILDEC-ROUGE-43-46,
image_link: https://cdn.shopify.com/s/files/1/0015/3088/8245/products/FECFILDEC-CARM-1_0c2b39dd-169f-461c-bab6-b0aa1e11d36e.jpg,
product_type: Chaussettes Unies homme,
availability: in stock,
shipping: 3,
shipping_weight: 0,
google_product_category: Vêtements et accessoires > Vêtements > Sous-vêtements et chaussettes > Chaussettes,
adwords_grouping: Chaussettes Unies homme,
adwords_labels: FECFILDEC-ROUGE-43-46,
gender: male,
age_group: adult,
color: ROUGE,
size: 43/46,
item_group_id: 1355395268661,
sale_price: 13.50,
adwords_redirect: ,
identifier_exists: ,
sale_price_effective_date: 2021-01-20T08:00+0100/2021-03-02T20:00+0100,
tax: ,
custom_label_0: ,
custom_label_1: ,
custom_label_2: ,
custom_label_3: ,
custom_label_4: ,
adult: ,
promotion_id: ,
shipping_length: ,
shipping_width: ,
shipping_height: ,
shipping_label: ,
additional_image_link: https://cdn.shopify.com/s/files/1/0015/3088/8245/products/FECFILDEC-CARM-1_0c2b39dd-169f-461c-bab6-b0aa1e11d36e.jpg,
mobile_link: ,
is_bundle: FALSE,
material: ,
display_ads_link: ,
display_ads_title: ,
excluded_destination: ,
display_ads_value: ,
pattern: ,
installment: ,
loyalty_points: ,
size_type: ,
size_system: ,
cross_sellers_product_id: ,
multipack: ,
availability_date: ,
unit_pricing_measure: ,
unit_pricing_base_measure: ,
display_ads_id: ,
min_handling_time: ,
max_handling_time: ,
sell_on_google_quantity: 2,
energy_efficiency_class: ,
min_energy_efficiency_class: ,
max_energy_efficiency_class: ,
product_fee: ,
product_detail: ,
consumer_datasheet: ,
return_address_label: ,
return_policy_label: ,
google_funded_promotion_eligibility: ,
signature_required: ,
return_rule_label: ,
product_highlight: 
\\ ===================================================================================== //

*/

type BoutiqueArthur struct {
	AdditionalImageLink              string `json:"additional_image_link" struct2map:"key:additional_image_link"`
	Adult                            string `json:"adult" struct2map:"key:adult"`
	AdwordsGrouping                  string `json:"adwords_grouping" struct2map:"key:adwords_grouping"`
	AdwordsLabels                    string `json:"adwords_labels" struct2map:"key:adwords_labels"`
	AdwordsRedirect                  string `json:"adwords_redirect" struct2map:"key:adwords_redirect"`
	AgeGroup                         string `json:"age_group" struct2map:"key:age_group"`
	Availability                     string `json:"availability" struct2map:"key:availability"`
	AvailabilityDate                 string `json:"availability_date" struct2map:"key:availability_date"`
	Brand                            string `json:"brand" struct2map:"key:brand"`
	Color                            string `json:"color" struct2map:"key:color"`
	Condition                        string `json:"condition" struct2map:"key:condition"`
	ConsumerDatasheet                string `json:"consumer_datasheet" struct2map:"key:consumer_datasheet"`
	CrossSellersProductID            string `json:"cross_sellers_product_id" struct2map:"key:cross_sellers_product_id"`
	CustomLabel0                     string `json:"custom_label_0" struct2map:"key:custom_label_0"`
	CustomLabel1                     string `json:"custom_label_1" struct2map:"key:custom_label_1"`
	CustomLabel2                     string `json:"custom_label_2" struct2map:"key:custom_label_2"`
	CustomLabel3                     string `json:"custom_label_3" struct2map:"key:custom_label_3"`
	CustomLabel4                     string `json:"custom_label_4" struct2map:"key:custom_label_4"`
	Description                      string `json:"description" struct2map:"key:description"`
	DisplayAdsID                     string `json:"display_ads_id" struct2map:"key:display_ads_id"`
	DisplayAdsLink                   string `json:"display_ads_link" struct2map:"key:display_ads_link"`
	DisplayAdsTitle                  string `json:"display_ads_title" struct2map:"key:display_ads_title"`
	DisplayAdsValue                  string `json:"display_ads_value" struct2map:"key:display_ads_value"`
	EnergyEfficiencyClass            string `json:"energy_efficiency_class" struct2map:"key:energy_efficiency_class"`
	ExcludedDestination              string `json:"excluded_destination" struct2map:"key:excluded_destination"`
	Gender                           string `json:"gender" struct2map:"key:gender"`
	GoogleFundedPromotionEligibility string `json:"google_funded_promotion_eligibility" struct2map:"key:google_funded_promotion_eligibility"`
	GoogleProductCategory            string `json:"google_product_category" struct2map:"key:google_product_category"`
	Gtin                             string `json:"gtin" struct2map:"key:gtin"`
	ID                               string `json:"id" struct2map:"key:id"`
	IdentifierExists                 string `json:"identifier_exists" struct2map:"key:identifier_exists"`
	ImageLink                        string `json:"image_link" struct2map:"key:image_link"`
	Installment                      string `json:"installment" struct2map:"key:installment"`
	IsBundle                         string `json:"is_bundle" struct2map:"key:is_bundle"`
	ItemGroupID                      string `json:"item_group_id" struct2map:"key:item_group_id"`
	Link                             string `json:"link" struct2map:"key:link"`
	LoyaltyPoints                    string `json:"loyalty_points" struct2map:"key:loyalty_points"`
	Material                         string `json:"material" struct2map:"key:material"`
	MaxEnergyEfficiencyClass         string `json:"max_energy_efficiency_class" struct2map:"key:max_energy_efficiency_class"`
	MaxHandlingTime                  string `json:"max_handling_time" struct2map:"key:max_handling_time"`
	MinEnergyEfficiencyClass         string `json:"min_energy_efficiency_class" struct2map:"key:min_energy_efficiency_class"`
	MinHandlingTime                  string `json:"min_handling_time" struct2map:"key:min_handling_time"`
	MobileLink                       string `json:"mobile_link" struct2map:"key:mobile_link"`
	Mpn                              string `json:"mpn" struct2map:"key:mpn"`
	Multipack                        string `json:"multipack" struct2map:"key:multipack"`
	Pattern                          string `json:"pattern" struct2map:"key:pattern"`
	Price                            string `json:"price" struct2map:"key:price"`
	ProductDetail                    string `json:"product_detail" struct2map:"key:product_detail"`
	ProductFee                       string `json:"product_fee" struct2map:"key:product_fee"`
	ProductHighlight                 string `json:"product_highlight" struct2map:"key:product_highlight"`
	ProductType                      string `json:"product_type" struct2map:"key:product_type"`
	PromotionID                      string `json:"promotion_id" struct2map:"key:promotion_id"`
	ReturnAddressLabel               string `json:"return_address_label" struct2map:"key:return_address_label"`
	ReturnPolicyLabel                string `json:"return_policy_label" struct2map:"key:return_policy_label"`
	ReturnRuleLabel                  string `json:"return_rule_label" struct2map:"key:return_rule_label"`
	SalePrice                        string `json:"sale_price" struct2map:"key:sale_price"`
	SalePriceEffectiveDate           string `json:"sale_price_effective_date" struct2map:"key:sale_price_effective_date"`
	SellOnGoogleQuantity             string `json:"sell_on_google_quantity" struct2map:"key:sell_on_google_quantity"`
	Shipping                         string `json:"shipping" struct2map:"key:shipping"`
	ShippingHeight                   string `json:"shipping_height" struct2map:"key:shipping_height"`
	ShippingLabel                    string `json:"shipping_label" struct2map:"key:shipping_label"`
	ShippingLength                   string `json:"shipping_length" struct2map:"key:shipping_length"`
	ShippingWeight                   string `json:"shipping_weight" struct2map:"key:shipping_weight"`
	ShippingWidth                    string `json:"shipping_width" struct2map:"key:shipping_width"`
	SignatureRequired                string `json:"signature_required" struct2map:"key:signature_required"`
	Size                             string `json:"size" struct2map:"key:size"`
	SizeSystem                       string `json:"size_system" struct2map:"key:size_system"`
	SizeType                         string `json:"size_type" struct2map:"key:size_type"`
	Tax                              string `json:"tax" struct2map:"key:tax"`
	Title                            string `json:"title" struct2map:"key:title"`
	UnitPricingBaseMeasure           string `json:"unit_pricing_base_measure" struct2map:"key:unit_pricing_base_measure"`
	UnitPricingMeasure               string `json:"unit_pricing_measure" struct2map:"key:unit_pricing_measure"`
}

/*
func (o *BoutiqueArthur) Convert() ([]string, error) {

	mp := make(map[string]string, 0) 

	mp["product_name"] = o.ProductName
	mp["reference"] = o.Reference

	categoryId, err := findOrCreateCategoryByNam(db, dbTablePrefix, o.Category)
	if err != nil {
		return nil, err
	}
	mp["category_id"] = fmt.Sprintf("%d", categoryId)

	defaultCategory, err := findOrCreateCategoryByName(db, dbTablePrefix, o.DefaultCategory)
	if err != nil {
		return nil, err
	}
	mp["default_category"] = fmt.Sprintf("%d", defaultCategory)

	mp["short_description"] = o.ShortDescription
	mp["description"] = o.Description
	mp["price"] = fo.Price
	mp["quantity"] = o.Quantity
	mp["image_ref"] = o.ImageRef

	// go ast ?
	// m := structs.Name(o)
	// pp.Println(m)
	// for _, feature := range o.Features {}
	mp["feature_name"] = o.FeatureNames
	mp["feature_values"] = o.FeatureValues

	row =  []string{
		mp["product_name"], 
		mp["reference"], 
		mp["category_id"],
		mp["default_category"],
		mp["short_description"],
		mp["description"],
		mp["price"],
		mp["quantity"],
		mp["image_ref"],
		mp["feature_name"],
		mp["feature_values"],
	}

	return row, nil
}
*/
