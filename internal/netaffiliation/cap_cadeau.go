package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/cap-cadeau.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
id: 2126,
title: Formation auto-massage en ligne,
description: L'auto-massage aide à évacuer les tensions, détendre les muscles, stimuler la circulation sanguine. Il contribue par ailleurs à un bon équilibre dans votre corps et votre esprit. Dans cette formation de 90 minutes, vous apprendrez de nombreuses techniques et points importants pour le pratiquer dans votre quotidien. Vous bénéficierez par ailleurs d'un conseil personnalisé pour l'utilisation des huiles.,
availability: in stock,
condition: new,
price: 75.00 EUR,
link: https://nec.capcadeau.com/?P5113E34598315S1UB4131af01360V4,
image_link: https://www.capcadeau.com/media/cache/og_filter/place/58/4304-img-3892.jpg,
brand: Sneha Bien-être
\\ ===================================================================================== //

*/

type CapCadeau struct {
	Availability string `json:"availability" struct2map:"key:availability"`
	Brand        string `json:"brand" struct2map:"key:brand"`
	Condition    string `json:"condition" struct2map:"key:condition"`
	Description  string `json:"description" struct2map:"key:description"`
	ID           string `json:"id" struct2map:"key:id"`
	ImageLink    string `json:"image_link" struct2map:"key:image_link"`
	Link         string `json:"link" struct2map:"key:link"`
	Price        string `json:"price" struct2map:"key:price"`
	Title        string `json:"title" struct2map:"key:title"`
}

/*
func (o *CapCadeau) Convert() ([]string, error) {

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
