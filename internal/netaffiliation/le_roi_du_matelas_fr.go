package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/le-roi-du-matelas-fr.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
id: 10004,
title: Matelas Altaïr - 90x190 (Livré roulé),
description: Avantages du matelas en latex Alta&iuml;r de Click 'n Sleep L'&acirc;me de ce matelas roul&eacute; est constitu&eacute;e d'une mousse latex de tr&egrave;s haute densit&eacute; afin de garantir une tenue exceptionnelle nuit apr&egrave;s nuit. Le latex offre un accueil plut&ocirc;t &eacute;quilibr&eacute; et &eacute;pouse vos formes sous l'effet de la chaleur corporelle pour un confort enveloppant. Mati&egrave;re naturellement hypoallerg&eacute;nique et mieux a&eacute;r&eacute;e qu'une mousse classique, le matelas Alta&iuml;r offre un excellent confort et une tr&egrave;s bonne dur&eacute;e de vie. Il est destin&eacute; aux personnes de faible ou de moyenne corpulence ne souffrant pas excessivement de la chaleur durant le sommeil. Composition du matelas en latex Alta&iuml;r Le matelas latex Alta&iuml;r est compos&eacute; d'une &acirc;me de 14 cm de latex d'une densit&eacute; de 76kg/m&sup3; pour une &eacute;paisseur de matelas totale de 17,5cm. Son coutil est compos&eacute; &agrave; 100% polyester piqu&eacute; sur 200gr/m&sup2; ouate polyester et 7mm de mousse confort. Choix du sommier pour le matelas en latex Alta&iuml;r Pour cette technologie de matelas roul&eacute; en latex, nous recommandons d'utiliser un sommier &agrave; lattes ou un sommier tapissier.,
google_product_category: Arborescence Globale > Matelas  > Matelas 1 personne > Matelas 90x190,
product_type: Arborescence Globale > Matelas  > Matelas 1 personne > Matelas 90x190,
link: https://action.metaffiliation.com/trk.php?mclic=P4D4434598316S1UB41248205140V4,
mobile_link: ,
image_link: http://www.leroidumatelas.fr/media/catalog/product/m/a/matelas-latex-altair_6.jpg,
additional_image_link: http://www.leroidumatelas.fr/media/catalog/product/m/a/matelas-latex-altair-1_6.jpg,
condition: new,
availability: in stock,
availability_date: ,
price: 259,
sale_price: ,
sale_price_effective_date: ,
gtin: 3000000270011,
mpn: ,
brand: Click 'n Sleep,
identifier_exists: ,
item_group_id: ,
color: ,
gender: ,
age_group: ,
material: ,
pattern: ,
size: ,
size_type: ,
size_system: ,
shipping: FR:::35,
shipping_weight: ,
shipping_label: ,
multipack: ,
is_bundle: ,
adult: ,
adwords_redirect: ,
custom_label_0: ,
custom_label_1: ,
custom_label_2: ,
custom_label_3: ,
custom_label_4: ,
excluded_destination: ,
expiration_date: ,
unit_pricing_measure: ,
unit_pricing_base_measure: ,
energy_efficiency_class: ,
promotion_id: ,
shipping_length: ,
shipping_width: ,
shipping_height: ,
max_handling_time: ,
min_handling_time: ,
min_energy_efficiency_class: ,
max_energy_efficiency_class: 
\\ ===================================================================================== //

*/

type LeRoiDuMatelasFr struct {
	AdditionalImageLink      string `json:"additional_image_link" struct2map:"key:additional_image_link"`
	Adult                    string `json:"adult" struct2map:"key:adult"`
	AdwordsRedirect          string `json:"adwords_redirect" struct2map:"key:adwords_redirect"`
	AgeGroup                 string `json:"age_group" struct2map:"key:age_group"`
	Availability             string `json:"availability" struct2map:"key:availability"`
	AvailabilityDate         string `json:"availability_date" struct2map:"key:availability_date"`
	Brand                    string `json:"brand" struct2map:"key:brand"`
	Color                    string `json:"color" struct2map:"key:color"`
	Condition                string `json:"condition" struct2map:"key:condition"`
	CustomLabel0             string `json:"custom_label_0" struct2map:"key:custom_label_0"`
	CustomLabel1             string `json:"custom_label_1" struct2map:"key:custom_label_1"`
	CustomLabel2             string `json:"custom_label_2" struct2map:"key:custom_label_2"`
	CustomLabel3             string `json:"custom_label_3" struct2map:"key:custom_label_3"`
	CustomLabel4             string `json:"custom_label_4" struct2map:"key:custom_label_4"`
	Description              string `json:"description" struct2map:"key:description"`
	EnergyEfficiencyClass    string `json:"energy_efficiency_class" struct2map:"key:energy_efficiency_class"`
	ExcludedDestination      string `json:"excluded_destination" struct2map:"key:excluded_destination"`
	ExpirationDate           string `json:"expiration_date" struct2map:"key:expiration_date"`
	Gender                   string `json:"gender" struct2map:"key:gender"`
	GoogleProductCategory    string `json:"google_product_category" struct2map:"key:google_product_category"`
	Gtin                     string `json:"gtin" struct2map:"key:gtin"`
	ID                       string `json:"id" struct2map:"key:id"`
	IdentifierExists         string `json:"identifier_exists" struct2map:"key:identifier_exists"`
	ImageLink                string `json:"image_link" struct2map:"key:image_link"`
	IsBundle                 string `json:"is_bundle" struct2map:"key:is_bundle"`
	ItemGroupID              string `json:"item_group_id" struct2map:"key:item_group_id"`
	Link                     string `json:"link" struct2map:"key:link"`
	Material                 string `json:"material" struct2map:"key:material"`
	MaxEnergyEfficiencyClass string `json:"max_energy_efficiency_class" struct2map:"key:max_energy_efficiency_class"`
	MaxHandlingTime          string `json:"max_handling_time" struct2map:"key:max_handling_time"`
	MinEnergyEfficiencyClass string `json:"min_energy_efficiency_class" struct2map:"key:min_energy_efficiency_class"`
	MinHandlingTime          string `json:"min_handling_time" struct2map:"key:min_handling_time"`
	MobileLink               string `json:"mobile_link" struct2map:"key:mobile_link"`
	Mpn                      string `json:"mpn" struct2map:"key:mpn"`
	Multipack                string `json:"multipack" struct2map:"key:multipack"`
	Pattern                  string `json:"pattern" struct2map:"key:pattern"`
	Price                    string `json:"price" struct2map:"key:price"`
	ProductType              string `json:"product_type" struct2map:"key:product_type"`
	PromotionID              string `json:"promotion_id" struct2map:"key:promotion_id"`
	SalePrice                string `json:"sale_price" struct2map:"key:sale_price"`
	SalePriceEffectiveDate   string `json:"sale_price_effective_date" struct2map:"key:sale_price_effective_date"`
	Shipping                 string `json:"shipping" struct2map:"key:shipping"`
	ShippingHeight           string `json:"shipping_height" struct2map:"key:shipping_height"`
	ShippingLabel            string `json:"shipping_label" struct2map:"key:shipping_label"`
	ShippingLength           string `json:"shipping_length" struct2map:"key:shipping_length"`
	ShippingWeight           string `json:"shipping_weight" struct2map:"key:shipping_weight"`
	ShippingWidth            string `json:"shipping_width" struct2map:"key:shipping_width"`
	Size                     string `json:"size" struct2map:"key:size"`
	SizeSystem               string `json:"size_system" struct2map:"key:size_system"`
	SizeType                 string `json:"size_type" struct2map:"key:size_type"`
	Title                    string `json:"title" struct2map:"key:title"`
	UnitPricingBaseMeasure   string `json:"unit_pricing_base_measure" struct2map:"key:unit_pricing_base_measure"`
	UnitPricingMeasure       string `json:"unit_pricing_measure" struct2map:"key:unit_pricing_measure"`
}

/*
func (o *LeRoiDuMatelasFr) Convert() ([]string, error) {

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
