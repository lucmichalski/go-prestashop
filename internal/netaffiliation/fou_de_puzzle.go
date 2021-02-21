package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/fou-de-puzzle.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
prod_number: 1,
prod_name: Puzzle Observation - Princesse,
prod_price: 13.95,
bp: 13.95,
bt: ,
prod_url: https://action.metaffiliation.com/trk.php?mclic=P4AD3545983221S1UC41290b082408V4,
category: Puzzles - Princes et Princesses,
prod_description: Puzzle Observation - Princesse 54 Pièces  Djeco,
prod_description_long: Puzzle Observation - Princesse 54 Pièces  Djeco,
img_small: https://data.fou-de-puzzle.com/m22/p1/p1_80x80.jpg,
img_medium: https://data.fou-de-puzzle.com/m22/p1/p1.jpg,
img_large: https://data.fou-de-puzzle.com/m22/p1/p1.jpg,
prod_ean: ,
ShippingAndHandlingCost: 4.95,
nom_marque: Djeco,
image2: ,
image3: ,
prod_ean: 
\\ ===================================================================================== //

*/

type FouDePuzzle struct {
	Bp                      string `json:"bp" struct2map:"key:bp"`
	Bt                      string `json:"bt" struct2map:"key:bt"`
	Category                string `json:"category" struct2map:"key:category"`
	Image2                  string `json:"image_2" struct2map:"key:image_2"`
	Image3                  string `json:"image_3" struct2map:"key:image_3"`
	ImgLarge                string `json:"img_large" struct2map:"key:img_large"`
	ImgMedium               string `json:"img_medium" struct2map:"key:img_medium"`
	ImgSmall                string `json:"img_small" struct2map:"key:img_small"`
	NomMarque               string `json:"nom_marque" struct2map:"key:nom_marque"`
	ProdDescription         string `json:"prod_description" struct2map:"key:prod_description"`
	ProdDescriptionLong     string `json:"prod_description_long" struct2map:"key:prod_description_long"`
	ProdEan                 string `json:"prod_ean" struct2map:"key:prod_ean"`
	ProdName                string `json:"prod_name" struct2map:"key:prod_name"`
	ProdNumber              string `json:"prod_number" struct2map:"key:prod_number"`
	ProdPrice               string `json:"prod_price" struct2map:"key:prod_price"`
	ProdURL                 string `json:"prod_url" struct2map:"key:prod_url"`
	ShippingAndHandlingCost string `json:"shipping_and_handling_cost" struct2map:"key:shipping_and_handling_cost"`
}

/*
func (o *FouDePuzzle) Convert() ([]string, error) {

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
