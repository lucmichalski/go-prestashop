package effiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/pieces-et-pneus-flux-pneus.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
additional_image_link: https://piecesetpneus.s3.eu-west-3.amazonaws.com/medias/images/TOYO_C1S_2.JPG,
additional_image_link2: https://piecesetpneus.s3.eu-west-3.amazonaws.com/medias/images/TOYO_C1S_3.JPG,
additional_image_link3: https://piecesetpneus.s3.eu-west-3.amazonaws.com/medias/images/,
additional_image_link4: ,
age_group: ,
availability: ,
availability_date: ,
brand: TOYO,
category: Pneus,
category_level2: ,
category_level3: ,
category_level4: https://piecesetpneus.s3.eu-west-3.amazonaws.com/medias/images/TOYO_C1S_4.JPG,
color: ,
condition: ,
currency: ,
custom_label_0: ,
custom_label_1: ,
custom_label_2: ,
custom_label_3: ,
custom_label_4: ,
description: Le pneu Confort - Silence. Développé spécifiquement pour les attentes des conducteurs de berlines puissantes et grandes routières, Le PROXES C1S est un pneu à basse résistance au roulement ce qui permet de réduire les émissions de CO2 et donc la consommation de carburant. Excellent comportement en aquaplaning et accroche sur route mouillée. Les repères: confort et silence remarquables / faible résistance au roulement / repousse les limites de l'aquaplaning / bénéficie du label écologique 'E-Leaf' Caractéristiques techniques: Freinage sur sol sec: 8/10 Tenue sur sol Humide: 8/10 Freinage sur routes mouillées: C Bruit: 68dB Longévité: 9/10 Stabilité: 8/10 Qualité prix: 9/10 Economies de carburant: F Conduite confort: 10/10 Nombre d'ondes: 1,
ecotax: ,
energy_efficiency_class: ,
expiration_date: ,
gender: ,
google_product_category: ,
gtin: 4981910883074,
id: 11320,
image_link: https://piecesetpneus.s3.eu-west-3.amazonaws.com/medias/images/TOYO_TC1S.JPG,
is_bundle: ,
item_group_id: ,
link: https://track.effiliation.com/servlet/effi.redir?id_compteur=22430515&url=https%3A%2F%2Fwww.piecesetpneus.com%2FF-11314-pneus%2FP-11320-pneu-voiture-toyo-proxes-c1s-205-55r16-94w,
material: ,
mobile_link: ,
mpn: ,
multipack: ,
pattern: ,
pneu_charge: ,
pneu_diametre: 16,
pneu_fueleconomy: ,
pneu_hauteur: ,
pneu_largeur: ,
pneu_modele: ,
pneu_noisedb: ,
pneu_noiselevel: 68,
pneu_rechape: ,
pneu_renforce: ,
pneu_runflat: ,
pneu_saison: Eté,
pneu_vitesse: W (inf à 270 km/h),
pneu_wetadhesion: C (Meilleure note possible: A),
price: 124.34,
price_norebate: 124.34,
rebate_end_date: ,
rebate_start_date: ,
sale_price_effective_date: ,
shipping: ,
shipping_cost: ,
shipping_height: 55,
shipping_length: ,
shipping_time: Départ sous 7 jours,
shipping_weight: ,
shipping_width: 205,
size: ,
size_system: ,
size_type: ,
stock: ,
style: ,
title: Pneu voiture Toyo PROXES C1S 205/55R16 94W,
warranty: 
\\ ===================================================================================== //

*/

type PiecesEtPneusFluxPneu struct {
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
func (o *PiecesEtPneusFluxPneu) Convert() ([]string, error) {

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
