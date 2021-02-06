package fixtures

import (
	"encoding/xml"
)

// CarrierGroup was generated 2021-02-06 11:40:20 by evolutive on eg-cdn.gsi-network.com.
type CarrierGroup struct {
	XMLName xml.Name `xml:"carrier_group"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text    string `xml:",chardata"`
		Primary string `xml:"primary,attr"`
		Sql     string `xml:"sql,attr"`
		Field   []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text         string `xml:",chardata"`
		CarrierGroup []struct {
			Text      string `xml:",chardata"`
			IDCarrier string `xml:"id_carrier,attr"`
			IDGroup   string `xml:"id_group,attr"`
			ID        string `xml:"id,attr"`
		} `xml:"carrier_group"`
	} `xml:"entities"`
}
