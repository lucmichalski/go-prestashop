package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/famille-mary-miel-et-plantes.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
EAN or ISBN: 3760115091542,
name of the product: Pastilles du Soir - 1 boite,
internal reference: 100,
current price: 5.1,
crossed price: ,
product category: ,
product page URL: https://rse.famillemary.fr/?P411D045983217S1UB41296101356V4,
big image: https://www.famillemary.fr/media/catalog/product/p/a/pastilles_du_soir_2.png,
manufacturer reference: ,
brand: Famille Mary,
description: Nos pastilles du soir associent 3 actifs réputés retrouver le sommeil : du miel, de la verveine et de l'huile essentielle d'Oranger. Ces pastilles vous offre une solution naturelle qui vous aidera à tomber plus facilement dans les bras de Morphée. Utilisation : Dégustez les pastilles du soir juste avant le coucher. 1 Boîte de 60g,
product availability: ,
condition: ,
shipping costs: ,
ecotaxe: ,
warranty: ,
stock indicator: ,
performance indicator: ,
discount indicator: 
\\ ===================================================================================== //

*/

type FamilleMaryMielEtPlante struct {
	BigImage              string `json:"big_image" struct2map:"key:big_image"`
	Brand                 string `json:"brand" struct2map:"key:brand"`
	Condition             string `json:"condition" struct2map:"key:condition"`
	CrossedPrice          string `json:"crossed_price" struct2map:"key:crossed_price"`
	CurrentPrice          string `json:"current_price" struct2map:"key:current_price"`
	Description           string `json:"description" struct2map:"key:description"`
	DiscountIndicator     string `json:"discount_indicator" struct2map:"key:discount_indicator"`
	EanOrIsbn             string `json:"ean_or_isbn" struct2map:"key:ean_or_isbn"`
	Ecotaxe               string `json:"ecotaxe" struct2map:"key:ecotaxe"`
	InternalReference     string `json:"internal_reference" struct2map:"key:internal_reference"`
	ManufacturerReference string `json:"manufacturer_reference" struct2map:"key:manufacturer_reference"`
	NameOfTheProduct      string `json:"name_of_the_product" struct2map:"key:name_of_the_product"`
	PerformanceIndicator  string `json:"performance_indicator" struct2map:"key:performance_indicator"`
	ProductAvailability   string `json:"product_availability" struct2map:"key:product_availability"`
	ProductCategory       string `json:"product_category" struct2map:"key:product_category"`
	ProductPageURL        string `json:"product_page_url" struct2map:"key:product_page_url"`
	ShippingCosts         string `json:"shipping_costs" struct2map:"key:shipping_costs"`
	StockIndicator        string `json:"stock_indicator" struct2map:"key:stock_indicator"`
	Warranty              string `json:"warranty" struct2map:"key:warranty"`
}

/*
func (o *FamilleMaryMielEtPlante) Convert() ([]string, error) {

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
