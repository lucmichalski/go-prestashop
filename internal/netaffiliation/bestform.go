package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/bestform.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
EAN or ISBN: BL24446_042,
name of the product: Soutien-gorge emboîtant avec armatures Skin Stockholm Bestform skin,
internal reference: BL24446_042,
current price: 32,
crossed price:  32,
product category: Bestform > Lingerie > Soutien-gorge > Avec armatures,
product page URL: https://action.metaffiliation.com/trk.php?mclic=P44CB145983225S1UB41284b05092V4,
big image: https://www.bestform-lingerie.fr/media/catalog/product/5/a/5a813a3c888a6b3298e3cbe665325741774fa805_BL24446_042_WEB_0.jpg,
manufacturer reference: ,
brand: Bestform,
description: STOCKHOLM sera votre indispensable invisible. Enfin des dessous discrets ! Soutien-gorge coque pour une poitrine sublimée et parfaitement maintenue ! Son mélange de maille et de rayures ajourées lui donne une allure graphique et moderne. Discret sous les vêtements en coloris skin.,
product availability: in stock,
condition: Enabled,
shipping costs: 0,
ecotaxe: ,
warranty: ,
stock indicator: ,
performance indicator: ,
discount indicator: ,
Saison: ,
Collection: STOCKHOLM
\\ ===================================================================================== //

*/

type Bestform struct {
	BigImage              string `json:"big_image" struct2map:"key:big_image"`
	Brand                 string `json:"brand" struct2map:"key:brand"`
	Collection            string `json:"collection" struct2map:"key:collection"`
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
	Saison                string `json:"saison" struct2map:"key:saison"`
	ShippingCosts         string `json:"shipping_costs" struct2map:"key:shipping_costs"`
	StockIndicator        string `json:"stock_indicator" struct2map:"key:stock_indicator"`
	Warranty              string `json:"warranty" struct2map:"key:warranty"`
}

/*
func (o *Bestform) Convert() ([]string, error) {

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
