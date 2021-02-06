package fixtures

import (
	"encoding/xml"
)

// RangePrice was generated 2021-02-06 11:18:51 by evolutive on eg-cdn.gsi-network.com.
type RangePrice struct {
	XMLName xml.Name `xml:"range_price"`
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
		Text       string `xml:",chardata"`
		RangePrice []struct {
			Text       string `xml:",chardata"`
			IDCarrier  string `xml:"id_carrier,attr"`
			ID         string `xml:"id,attr"`
			Delimiter2 string `xml:"delimiter2,attr"`
			Delimiter1 string `xml:"delimiter1,attr"`
		} `xml:"range_price"`
	} `xml:"entities"`
}
