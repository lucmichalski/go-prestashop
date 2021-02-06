package fixtures

import (
	"encoding/xml"
)

// Carrier was generated 2021-02-06 11:17:19 by evolutive on eg-cdn.gsi-network.com.
type Carrier struct {
	XMLName xml.Name `xml:"carrier"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Class string `xml:"class,attr"`
		Sql   string `xml:"sql,attr"`
		Field []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text    string `xml:",chardata"`
		Carrier []struct {
			Text             string `xml:",chardata"`
			ID               string `xml:"id,attr"`
			IDReference      string `xml:"id_reference,attr"`
			Active           string `xml:"active,attr"`
			ShippingHandling string `xml:"shipping_handling,attr"`
			RangeBehaviour   string `xml:"range_behaviour,attr"`
			IsFree           string `xml:"is_free,attr"`
			ShippingExternal string `xml:"shipping_external,attr"`
			NeedRange        string `xml:"need_range,attr"`
			ShippingMethod   string `xml:"shipping_method,attr"`
			MaxWidth         string `xml:"max_width,attr"`
			MaxHeight        string `xml:"max_height,attr"`
			MaxDepth         string `xml:"max_depth,attr"`
			MaxWeight        string `xml:"max_weight,attr"`
			Grade            string `xml:"grade,attr"`
			Name             string `xml:"name,attr"`
			URL              string `xml:"url,attr"`
			IDTaxRulesGroup  string `xml:"id_tax_rules_group,attr"`
			RangeBehavior    string `xml:"range_behavior,attr"`
		} `xml:"carrier"`
	} `xml:"entities"`
}
