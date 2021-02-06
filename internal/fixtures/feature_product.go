package fixtures

import (
	"encoding/xml"
)

// FeatureProduct was generated 2021-02-06 11:16:04 by evolutive on eg-cdn.gsi-network.com.
type FeatureProduct struct {
	XMLName xml.Name `xml:"feature_product"`
	Text    string   `xml:",chardata"`
	Fields  struct {
		Text    string `xml:",chardata"`
		Primary string `xml:"primary,attr"`
		Field   []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Relation string `xml:"relation,attr"`
		} `xml:"field"`
	} `xml:"fields"`
	Entities struct {
		Text           string `xml:",chardata"`
		FeatureProduct []struct {
			Text           string `xml:",chardata"`
			IDProduct      string `xml:"id_product,attr"`
			IDFeatureValue string `xml:"id_feature_value,attr"`
			IDFeature      string `xml:"id_feature,attr"`
		} `xml:"feature_product"`
	} `xml:"entities"`
}
