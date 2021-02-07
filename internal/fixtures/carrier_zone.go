package fixtures

import (
	"encoding/xml"
)

// CarrierZone was generated 2021-02-06 11:37:24 by evolutive on eg-cdn.gsi-network.com.
type CarrierZone struct {
	XMLName xml.Name `xml:"entity_carrier_zone"`
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
		Text        string `xml:",chardata"`
		CarrierZone []struct {
			Text      string `xml:",chardata"`
			IDCarrier string `xml:"id_carrier,attr"`
			IDZone    string `xml:"id_zone,attr"`
			ID        string `xml:"id,attr"`
		} `xml:"carrier_zone"`
	} `xml:"entities"`
}
