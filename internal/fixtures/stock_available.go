package fixtures

import (
	"encoding/xml"
)

// StockAvailable was generated 2021-02-06 11:40:20 by evolutive on eg-cdn.gsi-network.com.
type StockAvailable struct {
	XMLName xml.Name `xml:"stock_available"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text    string `xml:",chardata"`
		Primary string `xml:"primary,attr"`
		Class   string `xml:"class,attr"`
		Field   []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text           string `xml:",chardata"`
		StockAvailable []struct {
			Text               string `xml:",chardata"`
			IDProductAttribute string `xml:"id_product_attribute,attr"`
			IDProduct          string `xml:"id_product,attr"`
			IDStockAvailable   string `xml:"id_stock_available,attr"`
			Quantity           string `xml:"quantity,attr"`
			IDShop             string `xml:"id_shop,attr"`
			IDShopGroup        string `xml:"id_shop_group,attr"`
			DependsOnStock     string `xml:"depends_on_stock,attr"`
			OutOfStock         string `xml:"out_of_stock,attr"`
		} `xml:"stock_available"`
	} `xml:"entities"`
}
