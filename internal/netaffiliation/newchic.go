package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/newchic.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
product_id: SKU516069,
id: https://action.metaffiliation.com/trk.php?mclic=P4D7014598315S1UF4124cf010025126V4,
product_url: https://www.newchic.com/other-tools-4166/p-https://action.metaffiliation.com/trk.php?mclic=P4D7014598315S1UF4124cf010025126V4.html,
name: 3/32 '' Flame Ceramic Nail Drill Bit Tool Carbide Pedicure,
crossed_price: 13.70 EUR,
current_price: 13.70 EUR,
category: Beauty > Nail Art > Nail Tools > Other Tools,
image: https://imgaz1.chiccdn.com/thumb/large/oaupload/newchic/images/B6/BC/832cc829-03ae-4e26-a8c1-f82673b90065.jpg?s=906x906,
brand: Newchic,
currency: EUR,
description: 3/32 '' Flame Ceramic Nail Drill Bit Tool Carbide Pedicure,
availability: in stock,
condition: new,
shipping_cost: 4.57 EUR
\\ ===================================================================================== //

*/

type Newchic struct {
	Availability string `json:"availability" struct2map:"key:availability"`
	Brand        string `json:"brand" struct2map:"key:brand"`
	Category     string `json:"category" struct2map:"key:category"`
	Condition    string `json:"condition" struct2map:"key:condition"`
	CrossedPrice string `json:"crossed_price" struct2map:"key:crossed_price"`
	Currency     string `json:"currency" struct2map:"key:currency"`
	CurrentPrice string `json:"current_price" struct2map:"key:current_price"`
	Description  string `json:"description" struct2map:"key:description"`
	ID           string `json:"id" struct2map:"key:id"`
	Image        string `json:"image" struct2map:"key:image"`
	Name         string `json:"name" struct2map:"key:name"`
	ProductID    string `json:"product_id" struct2map:"key:product_id"`
	ProductURL   string `json:"product_url" struct2map:"key:product_url"`
	ShippingCost string `json:"shipping_cost" struct2map:"key:shipping_cost"`
}

/*
func (o *Newchic) Convert() ([]string, error) {

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
