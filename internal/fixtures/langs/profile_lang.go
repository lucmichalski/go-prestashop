package langs

import (
	"encoding/xml"
)

// Profile was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type ProfileLang struct {
	XMLName xml.Name `xml:"entity_profile"`
	Text    string   `xml:",chardata"`
	Profile []struct {
		IDLang string `xml:",chardata" struct2map:"key:id_lang"`
		ID     string `xml:"id,attr" struct2map:"key:id_profile"`
		Name   string `xml:"name,attr" struct2map:"key:name"`
	} `xml:"profile"`
}
