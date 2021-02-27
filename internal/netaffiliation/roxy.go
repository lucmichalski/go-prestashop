package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/roxy.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
Color ID: ARJX303486wbb3,
EAN: 3613376342123,
Name: Roxy PT BEACH CLASSICS FIX TRI - Blanc - Roxy,
Unique ID: ARJX303486wbb3,
Current Price: 32.99,
Was Price: 32.99,
Category: Swim > Haut de bikini,
URL: https://action.metaffiliation.com/trk.php?mclic=P447F3459832A7S1UC41281b070877V4,
Image URL: http://images.boardriders.com/globalGrey/roxy-products/all/default/large/arjx303486_roxy,w_wbb3_frt1.jpg,
MPN: 3613376342123,
Brand: Roxy,
Description: Roxy PT BEACH CLASSICS FIX TRI , Swim > Haut de bikini,
Shipping Cost: 0,
Ecotax: ,
Short Description: ,
Stock Indicator: 1,
Availability: ,
Performance Indicator: ,
Discount Indicator: 2,
Novelty Indicator: ,
Size: L,M,S,XL,XS
\\ ===================================================================================== //

*/

type Roxy struct {
	Availability         string `json:"availability" struct2map:"key:availability"`
	Brand                string `json:"brand" struct2map:"key:brand"`
	Category             string `json:"category" struct2map:"key:category"`
	ColorID              string `json:"color_id" struct2map:"key:color_id"`
	CurrentPrice         string `json:"current_price" struct2map:"key:current_price"`
	Description          string `json:"description" struct2map:"key:description"`
	DiscountIndicator    string `json:"discount_indicator" struct2map:"key:discount_indicator"`
	Ean                  string `json:"ean" struct2map:"key:ean"`
	Ecotax               string `json:"ecotax" struct2map:"key:ecotax"`
	ImageURL             string `json:"image_url" struct2map:"key:image_url"`
	Mpn                  string `json:"mpn" struct2map:"key:mpn"`
	Name                 string `json:"name" struct2map:"key:name"`
	NoveltyIndicator     string `json:"novelty_indicator" struct2map:"key:novelty_indicator"`
	PerformanceIndicator string `json:"performance_indicator" struct2map:"key:performance_indicator"`
	ShippingCost         string `json:"shipping_cost" struct2map:"key:shipping_cost"`
	ShortDescription     string `json:"short_description" struct2map:"key:short_description"`
	Size                 string `json:"size" struct2map:"key:size"`
	StockIndicator       string `json:"stock_indicator" struct2map:"key:stock_indicator"`
	UniqueID             string `json:"unique_id" struct2map:"key:unique_id"`
	URL                  string `json:"url" struct2map:"key:url"`
	WasPrice             string `json:"was_price" struct2map:"key:was_price"`
}

/*
func (o *Roxy) Convert() ([]string, error) {

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
