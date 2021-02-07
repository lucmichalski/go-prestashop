package langs

import (
	"encoding/xml"
)

// Profile was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type ProfileLang struct {
	XMLName xml.Name `xml:"entity_profile"`
	Text    string   `xml:",chardata"`
	Profile []struct {
		Text string `xml:",chardata" struct2map:"key:text"`
		ID   string `xml:"id,attr" struct2map:"key:id"`
		Name string `xml:"name,attr" struct2map:"key:name"`
	} `xml:"profile"`
}
