package fixtures

import (
	"encoding/xml"
)

// SpecificPrice was generated 2021-02-06 11:40:20 by evolutive on eg-cdn.gsi-network.com.
type SpecificPrice struct {
	XMLName xml.Name `xml:"entity_specific_price"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		Class string `xml:"class,attr"`
		Field []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text          string `xml:",chardata"`
		SpecificPrice []struct {
			Text               string `xml:",chardata"`
			IDCart             string `xml:"id_cart,attr"`
			IDProduct          string `xml:"id_product,attr"`
			IDCountry          string `xml:"id_country,attr"`
			IDGroup            string `xml:"id_group,attr"`
			IDCustomer         string `xml:"id_customer,attr"`
			IDProductAttribute string `xml:"id_product_attribute,attr"`
			IDShop             string `xml:"id_shop,attr"`
			IDShopGroup        string `xml:"id_shop_group,attr"`
			IDCurrency         string `xml:"id_currency,attr"`
			FromQuantity       string `xml:"from_quantity,attr"`
			Reduction          string `xml:"reduction,attr"`
			From               string `xml:"from,attr"`
			To                 string `xml:"to,attr"`
			Price              string `xml:"price,attr"`
			ReductionTax       string `xml:"reduction_tax,attr"`
			ReductionType      string `xml:"reduction_type,attr"`
		} `xml:"specific_price"`
	} `xml:"entities"`
}
