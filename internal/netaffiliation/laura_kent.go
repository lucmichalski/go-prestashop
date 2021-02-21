package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/laura-kent.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
name: Maillot de bain 1 pièce Maritim Noir::Turquoise ,
model number: 295937,
brand: Maritim,
sku: 00002/10X,
price: 49.99,
shipping costs: 5.95,
currency: EUR,
category: Femme > Vêtements > Maillots de bain > Maillots de bain 1 pièce,
description: Ce maillot de bain pratique fait partie des pièces indispensables pour une garde-robe complète! Il a des coques souples. Le maillot de bain ne comporte pas d'armatures.,
url: https://action.metaffiliation.com/trk.php?mclic=P46A8945983223S1UE412a3501563538V4,
ean: 4055705407346,
availability: auf Lager,
medium image url: https://ct-res.cloudinary.com/f_auto%2Cq_auto:good%2Cw_700/images/c25c28615c7194f80679d8a3f035db49.jpg,
large image url: https://ct-res.cloudinary.com/f_auto%2Cq_auto:good%2Cw_700/images/9daf7a40d84b3a386bd76369253923a4.jpg
\\ ===================================================================================== //

*/

type LauraKent struct {
	Availability   string `json:"availability" struct2map:"key:availability"`
	Brand          string `json:"brand" struct2map:"key:brand"`
	Category       string `json:"category" struct2map:"key:category"`
	Currency       string `json:"currency" struct2map:"key:currency"`
	Description    string `json:"description" struct2map:"key:description"`
	Ean            string `json:"ean" struct2map:"key:ean"`
	LargeImageURL  string `json:"large_image_url" struct2map:"key:large_image_url"`
	MediumImageURL string `json:"medium_image_url" struct2map:"key:medium_image_url"`
	ModelNumber    string `json:"model_number" struct2map:"key:model_number"`
	Name           string `json:"name" struct2map:"key:name"`
	Price          string `json:"price" struct2map:"key:price"`
	ShippingCosts  string `json:"shipping_costs" struct2map:"key:shipping_costs"`
	Sku            string `json:"sku" struct2map:"key:sku"`
	URL            string `json:"url" struct2map:"key:url"`
}

/*
func (o *LauraKent) Convert() ([]string, error) {

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
