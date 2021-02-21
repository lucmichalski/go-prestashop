package netaffiliation

import (
	"encoding/xml"
)

// Feeds was generated 2021-02-18 14:19:59 by evolutive on eg-cdn.gsi-network.com.
type Feeds struct {
	XMLName  xml.Name `xml:"feeds"`
	Text     string   `xml:",chardata"`
	Campaign []struct {
		Text         string `xml:",chardata"`
		Name         string `xml:"name"`
		ProductFeeds struct {
			Text        string `xml:",chardata"`
			ProductFeed []struct {
				Text    string `xml:",chardata"`
				Name    string `xml:"name,attr"`
				Site    string `xml:"site,attr"`
				Version string `xml:"version,attr"`
			} `xml:"product_feed"`
		} `xml:"product_feeds"`
	} `xml:"campaign"`
}
