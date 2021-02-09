package langs

import (
	"encoding/xml"
)

// FeatureValue was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type FeatureValueLang struct {
	XMLName      xml.Name `xml:"entity_feature_value"`
	Text         string   `xml:",chardata"`
	FeatureValue []struct {
		IDLang string `xml:",chardata" struct2map:"key:id_lang"`
		ID     string `xml:"id,attr" struct2map:"key:id_feature_value"`
		Value  string `xml:"value,attr" struct2map:"key:value"`
	} `xml:"feature_value"`
}
