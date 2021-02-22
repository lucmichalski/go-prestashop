package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/bearco.csv
// Catalog Separator: 	
// Catalog excerpt: #1

// ===================================================================================== \\
id: 31805291593821,
title: New era unisexe t-shirt "all teams" nba rouge xxs,
description: T-shirt imprimé NBA; Imprimé avec les noms des équipes de la NBA; Drapeau New Era brodé sur la manche gauche; Col ras du cou; 100% coton tissé;,
link: https://imx.bearco.fr/?P5114D74598317S1U841325101V4,
image_link: https://cdn.shopify.com/s/files/1/0072/7865/7625/products/homme-all-teams-nba-new-era-7.png?v=1582034576,
availability: in stock,
price: 32.00 EUR,
sale_price: 27.04 EUR,
brand: New Era,
gtin: 194089028199,
age_group: adult,
color: Red,
gender: unisex,
size: XXS,
item_group_id: 4496868278365,
shipping: FR:::3.04 EUR,
condition: new,
product_category: t-shirt
\\ ===================================================================================== //

*/

type Bearco struct {
	AgeGroup        string `json:"age_group" struct2map:"key:age_group"`
	Availability    string `json:"availability" struct2map:"key:availability"`
	Brand           string `json:"brand" struct2map:"key:brand"`
	Color           string `json:"color" struct2map:"key:color"`
	Condition       string `json:"condition" struct2map:"key:condition"`
	Description     string `json:"description" struct2map:"key:description"`
	Gender          string `json:"gender" struct2map:"key:gender"`
	Gtin            string `json:"gtin" struct2map:"key:gtin"`
	ID              string `json:"id" struct2map:"key:id"`
	ImageLink       string `json:"image_link" struct2map:"key:image_link"`
	ItemGroupID     string `json:"item_group_id" struct2map:"key:item_group_id"`
	Link            string `json:"link" struct2map:"key:link"`
	Price           string `json:"price" struct2map:"key:price"`
	ProductCategory string `json:"product_category" struct2map:"key:product_category"`
	SalePrice       string `json:"sale_price" struct2map:"key:sale_price"`
	Shipping        string `json:"shipping" struct2map:"key:shipping"`
	Size            string `json:"size" struct2map:"key:size"`
	Title           string `json:"title" struct2map:"key:title"`
}

/*
func (o *Bearco) Convert() ([]string, error) {

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
