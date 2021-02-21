package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/craftine.csv
// Catalog Separator: 	
// Catalog excerpt: #1

// ===================================================================================== \\
ean: 3663945029764,
title: Billes de polystyrène - Sac de 20 litres,
sku: TX_BILLPOLY_20L,
price: 9,50,
original_price: 9,50,
category: Tissus > Divers,
url: https://action.metaffiliation.com/trk.php?mclic=P4CB054598321BS1UB41283701335V4,
image_link: https://www.craftine.com/media/catalog/product/t/h/thumb_img_1132_1024.jpg,
brand: Craftine,
description: Composition : Polystyr&egrave;ne - Poids : 200 g/m² - Laize (largeur) : 150 cm - Motif : Divers,
availability: 1
\\ ===================================================================================== //

*/

type Craftine struct {
	Availability  string `json:"availability" struct2map:"key:availability"`
	Brand         string `json:"brand" struct2map:"key:brand"`
	Category      string `json:"category" struct2map:"key:category"`
	Description   string `json:"description" struct2map:"key:description"`
	Ean           string `json:"ean" struct2map:"key:ean"`
	ImageLink     string `json:"image_link" struct2map:"key:image_link"`
	OriginalPrice string `json:"original_price" struct2map:"key:original_price"`
	Price         string `json:"price" struct2map:"key:price"`
	Sku           string `json:"sku" struct2map:"key:sku"`
	Title         string `json:"title" struct2map:"key:title"`
	URL           string `json:"url" struct2map:"key:url"`
}

/*
func (o *Craftine) Convert() ([]string, error) {

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
