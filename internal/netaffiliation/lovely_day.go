package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/lovely-day.csv
// Catalog Separator: 	
// Catalog excerpt: #1

// ===================================================================================== \\
id: 1394,
title: Bracelet chaine boules motif triangle bordeaux en laiton doré - Lovely Day Bijoux,
brand: Lovely Day Bijoux,
description: Bracelet chaine boules motif triangle bordeaux en laiton doré,
price: 52.00 EUR,
sale_price: 52.00 EUR,
link: https://jaz.lovelydaybijoux.com/?P51140F4598315S1UB4131b501847V4,
image_link: https://www.lovelydaybijoux.com/media/cache/google_feed_product/2020/11/4901-3g4a0734.jpg,
condition: new,
availability: in stock,
google_product_category: ,
identifier_exists: no,
gtin: ,
shipping_weight: 0 kgs,
custom_label_0: ,
custom_label_1: ,
custom_label_2: ,
custom_label_3: LOD020918-BRP,
item_group_id: 857,
color: Bordeaux,
size: ,
material: Laiton doré,
gender: female,
age_group: adult
\\ ===================================================================================== //

*/

type LovelyDay struct {
	AgeGroup              string `json:"age_group" struct2map:"key:age_group"`
	Availability          string `json:"availability" struct2map:"key:availability"`
	Brand                 string `json:"brand" struct2map:"key:brand"`
	Color                 string `json:"color" struct2map:"key:color"`
	Condition             string `json:"condition" struct2map:"key:condition"`
	CustomLabel0          string `json:"custom_label_0" struct2map:"key:custom_label_0"`
	CustomLabel1          string `json:"custom_label_1" struct2map:"key:custom_label_1"`
	CustomLabel2          string `json:"custom_label_2" struct2map:"key:custom_label_2"`
	CustomLabel3          string `json:"custom_label_3" struct2map:"key:custom_label_3"`
	Description           string `json:"description" struct2map:"key:description"`
	Gender                string `json:"gender" struct2map:"key:gender"`
	GoogleProductCategory string `json:"google_product_category" struct2map:"key:google_product_category"`
	Gtin                  string `json:"gtin" struct2map:"key:gtin"`
	ID                    string `json:"id" struct2map:"key:id"`
	IdentifierExists      string `json:"identifier_exists" struct2map:"key:identifier_exists"`
	ImageLink             string `json:"image_link" struct2map:"key:image_link"`
	ItemGroupID           string `json:"item_group_id" struct2map:"key:item_group_id"`
	Link                  string `json:"link" struct2map:"key:link"`
	Material              string `json:"material" struct2map:"key:material"`
	Price                 string `json:"price" struct2map:"key:price"`
	SalePrice             string `json:"sale_price" struct2map:"key:sale_price"`
	ShippingWeight        string `json:"shipping_weight" struct2map:"key:shipping_weight"`
	Size                  string `json:"size" struct2map:"key:size"`
	Title                 string `json:"title" struct2map:"key:title"`
}

/*
func (o *LovelyDay) Convert() ([]string, error) {

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
