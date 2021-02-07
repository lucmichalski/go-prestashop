package fixtures

import (
	"encoding/xml"
)

// OrderCarrier was generated 2021-02-06 11:40:27 by evolutive on eg-cdn.gsi-network.com.
type OrderCarrier struct {
	XMLName xml.Name `xml:"entity_order_carrier"`
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
		Text         string `xml:",chardata"`
		OrderCarrier []struct {
			Text                string `xml:",chardata"`
			IDOrder             string `xml:"id_order,attr"`
			IDCarrier           string `xml:"id_carrier,attr"`
			Weight              string `xml:"weight,attr"`
			IDOrderInvoice      string `xml:"id_order_invoice,attr"`
			ShippingCostTaxExcl string `xml:"shipping_cost_tax_excl,attr"`
			ShippingCostTaxIncl string `xml:"shipping_cost_tax_incl,attr"`
			TrackingNumber      string `xml:"tracking_number,attr"`
		} `xml:"order_carrier"`
	} `xml:"entities"`
}
