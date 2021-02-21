package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/gsell-maroquinerie-et-articles.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
universal reference: SA65753,
title: Fabriquée dans un polycarbonate Mackrolon métallisé, cette valise 69 cm, combine légèrté et solidité grâce à son design moderne rainuré agrémenté d'angles renforcés.- Taille de la valise avec roues et poignées : L46 x H69 x E27 cm >> 85L- Taille de la valise sans roues et sans poignées : L44 x H65 x E27 cm >> 74LOn trouve à l'intérieur deux compartiments avec deux plateaux zippés et une paire de sangles de maintien.L'extérieur en détail :- 4 doubles roues pour faciliter le roulement,- 2 poignées main,- 1 poignée trolley,- 1 serrure chiffrée TSA,- et 1 porte-adresse au niveau de la poignée trolley.>> Comment coder / décoder ma serrure TSA ?Nos photographies sont retouchées. La retouche d'image n'altère en rien les qualités du produit.,
id: 63284,
current price: 195.0000,
old price: 279.00,
product category: Bagages et maroquinerie > Valises,
product link: https://action.metaffiliation.com/trk.php?mclic=P2E345983223S1UC4128b5015547V4,
image link: https://media.gsell.fr/pub/media/catalog/product/s/a/sa65753~saneo1541.jpg,
manufacturer reference: 5414847565670,
brand: Samsonite,
description: Fabriquée dans un polycarbonate Mackrolon métallisé, cette valise 69 cm, combine légèrté et solidité grâce à son design moderne rainuré agrémenté d'angles renforcés.- Taille de la valise avec roues et poignées : L46 x H69 x E27 cm >> 85L- Taille de la valise sans roues et sans poignées : L44 x H65 x E27 cm >> 74LOn trouve à l'intérieur deux compartiments avec deux plateaux zippés et une paire de sangles de maintien.L'extérieur en détail :- 4 doubles roues pour faciliter le roulement,- 2 poignées main,- 1 poignée trolley,- 1 serrure chiffrée TSA,- et 1 porte-adresse au niveau de la poignée trolley.>> Comment coder / décoder ma serrure TSA ?Nos photographies sont retouchées. La retouche d'image n'altère en rien les qualités du produit.,
product availability: in stock,
condition: new,
shipping costs: 0.00,
ecotaxe: ,
warranty: ,
stock indicator: 3,
performance indicator: ,
discount indicator: ,
novelty indicator: 
\\ ===================================================================================== //

*/

type GsellMaroquinerieEtArticle struct {
	Brand                 string `json:"brand" struct2map:"key:brand"`
	Condition             string `json:"condition" struct2map:"key:condition"`
	CurrentPrice          string `json:"current_price" struct2map:"key:current_price"`
	Description           string `json:"description" struct2map:"key:description"`
	DiscountIndicator     string `json:"discount_indicator" struct2map:"key:discount_indicator"`
	Ecotaxe               string `json:"ecotaxe" struct2map:"key:ecotaxe"`
	ID                    string `json:"id" struct2map:"key:id"`
	ImageLink             string `json:"image_link" struct2map:"key:image_link"`
	ManufacturerReference string `json:"manufacturer_reference" struct2map:"key:manufacturer_reference"`
	NoveltyIndicator      string `json:"novelty_indicator" struct2map:"key:novelty_indicator"`
	OldPrice              string `json:"old_price" struct2map:"key:old_price"`
	PerformanceIndicator  string `json:"performance_indicator" struct2map:"key:performance_indicator"`
	ProductAvailability   string `json:"product_availability" struct2map:"key:product_availability"`
	ProductCategory       string `json:"product_category" struct2map:"key:product_category"`
	ProductLink           string `json:"product_link" struct2map:"key:product_link"`
	ShippingCosts         string `json:"shipping_costs" struct2map:"key:shipping_costs"`
	StockIndicator        string `json:"stock_indicator" struct2map:"key:stock_indicator"`
	Title                 string `json:"title" struct2map:"key:title"`
	UniversalReference    string `json:"universal_reference" struct2map:"key:universal_reference"`
	Warranty              string `json:"warranty" struct2map:"key:warranty"`
}

/*
func (o *GsellMaroquinerieEtArticle) Convert() ([]string, error) {

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
