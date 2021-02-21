package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/dresslily-fr.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
"Universal reference": 475112101,
Product Name: Ceintures Femme Ceinture Élastique Evidée Bouclée,
Internal reference: 8372198,
Current Price: 5.95,
Crossed Price: 7.83,
Product category: Accessoires > Accessoires Femme > Ceintures,
Product page URL: https://action.metaffiliation.com/trk.php?mclic=P4D181459831FS1UA4126630250V4,
URL of the big image: https://gloimg.drlcdn.com/L/pdm-product-pic/Clothing/2020/11/11/source-img/20201111091317_5fab3aadc2545.jpg,
Name of brand: Dresslily,
Discription of the product: Dresslily  Type: Pour Femmes Groupe: Adulte Style: à la Mode Matériel de Ceinture: PU Type de Motif: Couleur Unie Longueur de Ceinture: 80CM Largeur de Ceinture: 7CM Poids: 0,0630kg Liste d'emballage: 1 x Ceinture ,
Shipping cost: ,
Stock indicator: 1,
New indicator: 1,
Color: Noir,
Size: ONE SIZE
\\ ===================================================================================== //

*/

type DresslilyFr struct {
	Color                   string `json:"color" struct2map:"key:color"`
	CrossedPrice            string `json:"crossed_price" struct2map:"key:crossed_price"`
	CurrentPrice            string `json:"current_price" struct2map:"key:current_price"`
	DiscriptionOfTheProduct string `json:"discription_of_the_product" struct2map:"key:discription_of_the_product"`
	InternalReference       string `json:"internal_reference" struct2map:"key:internal_reference"`
	NameOfBrand             string `json:"name_of_brand" struct2map:"key:name_of_brand"`
	NewIndicator            string `json:"new_indicator" struct2map:"key:new_indicator"`
	ProductCategory         string `json:"product_category" struct2map:"key:product_category"`
	ProductName             string `json:"product_name" struct2map:"key:product_name"`
	ProductPageURL          string `json:"product_page_url" struct2map:"key:product_page_url"`
	ShippingCost            string `json:"shipping_cost" struct2map:"key:shipping_cost"`
	Size                    string `json:"size" struct2map:"key:size"`
	StockIndicator          string `json:"stock_indicator" struct2map:"key:stock_indicator"`
	UniversalReference      string `json:"universal_reference" struct2map:"key:universal_reference"`
	URLOfTheBigImage        string `json:"url_of_the_big_image" struct2map:"key:url_of_the_big_image"`
}

/*
func (o *DresslilyFr) Convert() ([]string, error) {

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
