package fixtures

import (
	"encoding/xml"
)

// Store was generated 2021-02-06 11:40:35 by evolutive on eg-cdn.gsi-network.com.
type Store struct {
	XMLName xml.Name `xml:"store"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Class string `xml:"class,attr"`
		Image string `xml:"image,attr"`
		Field []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text  string `xml:",chardata"`
		Store []struct {
			Text      string `xml:",chardata"`
			IDCountry string `xml:"id_country,attr"`
			IDState   string `xml:"id_state,attr"`
			ID        string `xml:"id,attr"`
			City      string `xml:"city,attr"`
			Postcode  string `xml:"postcode,attr"`
			Latitude  string `xml:"latitude,attr"`
			Longitude string `xml:"longitude,attr"`
			Email     string `xml:"email,attr"`
			Phone     string `xml:"phone,attr"`
			Fax       string `xml:"fax,attr"`
			Active    string `xml:"active,attr"`
		} `xml:"store"`
	} `xml:"entities"`
}
