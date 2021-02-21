package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/zaful-fr.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
"Universal reference": 465916001,
Product Name: Emma MacDonald x ZAFUL Maillot de Bain Bikini Animal Imprimé Plissé à Bretelle S Léopard,
Internal reference: 899045,
Current Price: 14.98,
Crossed Price: 26.37,
Product category: Maillots de Bain > Bikinis > Ensembles de bikini,
Product page URL: https://action.metaffiliation.com/trk.php?mclic=P4DE33459831BS1UE4125f501697045V4,
URL of the big image: https://gloimg.zafcdn.com/zaful/pdm-product-pic/Clothing/2020/06/17/source-img/20200617181507_21504.jpg,
Name of brand: ZAFUL,
Discription of the product: Zaful  Couvert dans un imprimé léopard à carillon avec l'ambiance décontractée, l'ensemble en deux parties comprend un haut triangulaire mémoires et cravate côté, deux caractéristiques smockée détail qui épouse votre corps et créer une finition dimensionnelle, de texture. Style: Sexy Ty,
Shipping cost: ,
Stock indicator: 1,
New indicator: 1,
Color: Léopard,
Size: S
\\ ===================================================================================== //

*/

type ZafulFr struct {
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
func (o *ZafulFr) Convert() ([]string, error) {

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
