package fixtures

import (
	"encoding/xml"
)

// Delivery was generated 2021-02-06 11:40:25 by evolutive on eg-cdn.gsi-network.com.
type Delivery struct {
	XMLName xml.Name `xml:"entity_delivery"`
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
		Text     string `xml:",chardata"`
		Delivery []struct {
			Text          string `xml:",chardata"`
			IDCarrier     string `xml:"id_carrier,attr"`
			IDRangePrice  string `xml:"id_range_price,attr"`
			IDRangeWeight string `xml:"id_range_weight,attr"`
			IDZone        string `xml:"id_zone,attr"`
			ID            string `xml:"id,attr"`
			Price         string `xml:"price,attr"`
			IDShop        string `xml:"id_shop,attr"`
			IDShopGroup   string `xml:"id_shop_group,attr"`
		} `xml:"delivery"`
	} `xml:"entities"`
}
