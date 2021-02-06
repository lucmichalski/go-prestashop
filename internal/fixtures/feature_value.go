package fixtures

import (
	"encoding/xml"
)

// FeatureValue was generated 2021-02-06 11:18:51 by evolutive on eg-cdn.gsi-network.com.
type FeatureValue struct {
	XMLName xml.Name `xml:"feature_value"`
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
		Text         string `xml:",chardata"`
		FeatureValue []struct {
			Text      string `xml:",chardata"`
			IDFeature string `xml:"id_feature,attr"`
			ID        string `xml:"id,attr"`
			Custom    string `xml:"custom,attr"`
		} `xml:"feature_value"`
	} `xml:"entities"`
}
