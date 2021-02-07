package langs

import (
	"encoding/xml"
)

// Store was generated 2021-02-06 11:43:01 by evolutive on eg-cdn.gsi-network.com.
type StoreLang struct {
	XMLName xml.Name `xml:"entity_store"`
	Text    string   `xml:",chardata"`
	Store   []struct {
		Text     string `xml:",chardata" struct2map:"key:text"`
		ID       string `xml:"id,attr" struct2map:"key:id"`
		Name     string `xml:"name,attr" struct2map:"key:name"`
		Address1 string `xml:"address1,attr" struct2map:"key:address1"`
		Address2 string `xml:"address2,attr" struct2map:"key:address2"`
		Hours    string `xml:"hours,attr" struct2map:"key:hours"`
	} `xml:"store"`
}
