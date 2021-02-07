package langs

import (
	"encoding/xml"
)

// Feature was generated 2021-02-06 11:43:01 by evolutive on eg-cdn.gsi-network.com.
type FeatureLang struct {
	XMLName xml.Name `xml:"entity_feature"`
	Text    string   `xml:",chardata"`
	Feature []struct {
		Text string `xml:",chardata" struct2map:"-"`
		ID   string `xml:"id,attr" struct2map:"key:id"`
		Name string `xml:"name,attr" struct2map:"key:name"`
	} `xml:"feature"`
}
