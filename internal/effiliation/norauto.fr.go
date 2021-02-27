package effiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/norauto-fr.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
additional_image_link: ,
additional_image_link2: ,
additional_image_link3: ,
additional_image_link4: ,
age_group: ,
availability: en stock,
availability_date: ,
brand: Norauto,
category: Equiper ma voiture,
category_level2: Sécurité - prévention,
category_level3: Cannes antivols - alarmes,
category_level4: Canne antivol - bloque roue,
color: ,
condition: neuf,
currency: EUR,
custom_label_0: 0,
custom_label_1: ,
custom_label_2: ,
custom_label_3: ,
custom_label_4: ,
description: L'antivol frein à main NORAUTO vous permettra de dissuader contre le vol grâce à sa grande visibilité. Il a un design pratique et compact pour plus de simplicité d'installation et d'utilisation. Pour cela, il suffit de fixer l'antivol entre le levier de vitesse et le frein à main de votre véhicule.Vous pouvez ajuster la longueur de l'antivol pour atteindre la longueur désirée. Cet ajustement se fait grâce à la vis brevetée pour un réglage plus simple.L'antivol mesure 34 cm en longueur une fois plié, et lorsqu'il est déplié, il mesure au maximum 36 cm en longueur. Ne convient pas pour les véhicules avec unécartement entre le levier de vitesse et le frein à main inférieur à 20 cm et supérieur à 36 cm. Le levier de vitesse ne doit pas mesurer plus de 3,5 cm de diamètre.L'antivol se ferme grâce à une serrure à clés (fournies avec l'antivol).Cet antivol est fabriqué en Europe avec la technologie la plus avancée.,
ecotax: ,
energy_efficiency_class: ,
expiration_date: ,
gender: ,
google_product_category: ,
gtin: 3501361141311,
id: 2004744,
image_link: https://s3.medias-norauto.fr/images_produits/2004744/650x650/antivol-frein-a-main-norauto--2004744.jpg,
is_bundle: ,
item_group_id: ,
link: https://track.effiliation.com/servlet/effi.redir?id_compteur=22432280&url=https%3A%2F%2Fwww.norauto.fr%2Fp%2Fantivol-frein-a-main-norauto-2004744.html,
material: ,
mobile_link: ,
mpn: ,
multipack: ,
pattern: ,
pneu_charge: ,
pneu_diametre: ,
pneu_fueleconomy: ,
pneu_hauteur: ,
pneu_largeur: ,
pneu_modele: ,
pneu_noisedb: ,
pneu_noiselevel: ,
pneu_rechape: ,
pneu_renforce: ,
pneu_runflat: ,
pneu_saison: ,
pneu_vitesse: ,
pneu_wetadhesion: ,
price: 34.95,
price_norebate: ,
rebate_end_date: ,
rebate_start_date: ,
sale_price_effective_date: ,
shipping: Livraison standard (sous 48h ouvrées après expédition),
shipping_cost: 0,
shipping_height: ,
shipping_length: ,
shipping_time: à partir de 1 heures,
shipping_weight: ,
shipping_width: ,
size: ,
size_system: ,
size_type: ,
stock: ,
style: ,
title: Antivol Frein À Main Norauto,
warranty: 
\\ ===================================================================================== //

*/

type NorautoFr struct {
	AdditionalImageLink    string `json:"additional_image_link" struct2map:"key:additional_image_link"`
	AdditionalImageLink2   string `json:"additional_image_link_2" struct2map:"key:additional_image_link_2"`
	AdditionalImageLink3   string `json:"additional_image_link_3" struct2map:"key:additional_image_link_3"`
	AdditionalImageLink4   string `json:"additional_image_link_4" struct2map:"key:additional_image_link_4"`
	AgeGroup               string `json:"age_group" struct2map:"key:age_group"`
	Availability           string `json:"availability" struct2map:"key:availability"`
	AvailabilityDate       string `json:"availability_date" struct2map:"key:availability_date"`
	Brand                  string `json:"brand" struct2map:"key:brand"`
	Category               string `json:"category" struct2map:"key:category"`
	CategoryLevel2         string `json:"category_level_2" struct2map:"key:category_level_2"`
	CategoryLevel3         string `json:"category_level_3" struct2map:"key:category_level_3"`
	CategoryLevel4         string `json:"category_level_4" struct2map:"key:category_level_4"`
	Color                  string `json:"color" struct2map:"key:color"`
	Condition              string `json:"condition" struct2map:"key:condition"`
	Currency               string `json:"currency" struct2map:"key:currency"`
	CustomLabel0           string `json:"custom_label_0" struct2map:"key:custom_label_0"`
	CustomLabel1           string `json:"custom_label_1" struct2map:"key:custom_label_1"`
	CustomLabel2           string `json:"custom_label_2" struct2map:"key:custom_label_2"`
	CustomLabel3           string `json:"custom_label_3" struct2map:"key:custom_label_3"`
	CustomLabel4           string `json:"custom_label_4" struct2map:"key:custom_label_4"`
	Description            string `json:"description" struct2map:"key:description"`
	Ecotax                 string `json:"ecotax" struct2map:"key:ecotax"`
	EnergyEfficiencyClass  string `json:"energy_efficiency_class" struct2map:"key:energy_efficiency_class"`
	ExpirationDate         string `json:"expiration_date" struct2map:"key:expiration_date"`
	Gender                 string `json:"gender" struct2map:"key:gender"`
	GoogleProductCategory  string `json:"google_product_category" struct2map:"key:google_product_category"`
	Gtin                   string `json:"gtin" struct2map:"key:gtin"`
	ID                     string `json:"id" struct2map:"key:id"`
	ImageLink              string `json:"image_link" struct2map:"key:image_link"`
	IsBundle               string `json:"is_bundle" struct2map:"key:is_bundle"`
	ItemGroupID            string `json:"item_group_id" struct2map:"key:item_group_id"`
	Link                   string `json:"link" struct2map:"key:link"`
	Material               string `json:"material" struct2map:"key:material"`
	MobileLink             string `json:"mobile_link" struct2map:"key:mobile_link"`
	Mpn                    string `json:"mpn" struct2map:"key:mpn"`
	Multipack              string `json:"multipack" struct2map:"key:multipack"`
	Pattern                string `json:"pattern" struct2map:"key:pattern"`
	PneuCharge             string `json:"pneu_charge" struct2map:"key:pneu_charge"`
	PneuDiametre           string `json:"pneu_diametre" struct2map:"key:pneu_diametre"`
	PneuFueleconomy        string `json:"pneu_fueleconomy" struct2map:"key:pneu_fueleconomy"`
	PneuHauteur            string `json:"pneu_hauteur" struct2map:"key:pneu_hauteur"`
	PneuLargeur            string `json:"pneu_largeur" struct2map:"key:pneu_largeur"`
	PneuModele             string `json:"pneu_modele" struct2map:"key:pneu_modele"`
	PneuNoisedb            string `json:"pneu_noisedb" struct2map:"key:pneu_noisedb"`
	PneuNoiselevel         string `json:"pneu_noiselevel" struct2map:"key:pneu_noiselevel"`
	PneuRechape            string `json:"pneu_rechape" struct2map:"key:pneu_rechape"`
	PneuRenforce           string `json:"pneu_renforce" struct2map:"key:pneu_renforce"`
	PneuRunflat            string `json:"pneu_runflat" struct2map:"key:pneu_runflat"`
	PneuSaison             string `json:"pneu_saison" struct2map:"key:pneu_saison"`
	PneuVitesse            string `json:"pneu_vitesse" struct2map:"key:pneu_vitesse"`
	PneuWetadhesion        string `json:"pneu_wetadhesion" struct2map:"key:pneu_wetadhesion"`
	Price                  string `json:"price" struct2map:"key:price"`
	PriceNorebate          string `json:"price_norebate" struct2map:"key:price_norebate"`
	RebateEndDate          string `json:"rebate_end_date" struct2map:"key:rebate_end_date"`
	RebateStartDate        string `json:"rebate_start_date" struct2map:"key:rebate_start_date"`
	SalePriceEffectiveDate string `json:"sale_price_effective_date" struct2map:"key:sale_price_effective_date"`
	Shipping               string `json:"shipping" struct2map:"key:shipping"`
	ShippingCost           string `json:"shipping_cost" struct2map:"key:shipping_cost"`
	ShippingHeight         string `json:"shipping_height" struct2map:"key:shipping_height"`
	ShippingLength         string `json:"shipping_length" struct2map:"key:shipping_length"`
	ShippingTime           string `json:"shipping_time" struct2map:"key:shipping_time"`
	ShippingWeight         string `json:"shipping_weight" struct2map:"key:shipping_weight"`
	ShippingWidth          string `json:"shipping_width" struct2map:"key:shipping_width"`
	Size                   string `json:"size" struct2map:"key:size"`
	SizeSystem             string `json:"size_system" struct2map:"key:size_system"`
	SizeType               string `json:"size_type" struct2map:"key:size_type"`
	Stock                  string `json:"stock" struct2map:"key:stock"`
	Style                  string `json:"style" struct2map:"key:style"`
	Title                  string `json:"title" struct2map:"key:title"`
	Warranty               string `json:"warranty" struct2map:"key:warranty"`
}

/*
func (o *NorautoFr) Convert() ([]string, error) {

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
