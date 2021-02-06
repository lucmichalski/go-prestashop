package fixtures

import (
	"encoding/xml"
)

// FeatureValue was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type FeatureValue struct {
	XMLName      xml.Name `xml:"feature_value"`
	Text         string   `xml:",chardata"`
	FeatureValue []struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Value string `xml:"value,attr"`
	} `xml:"feature_value"`
}
