package langs

import (
	"encoding/xml"
)

// Image was generated 2021-02-06 11:42:59 by evolutive on eg-cdn.gsi-network.com.
type ImageLang struct {
	XMLName xml.Name `xml:"entity_image"`
	Text    string   `xml:",chardata"`
	Image   []struct {
		IDLang string `xml:",chardata" struct2map:"key:id_lang"`
		ID     string `xml:"id,attr" struct2map:"key:id_image"`
		Legend string `xml:"legend,attr" struct2map:"key:legend"`
	} `xml:"image"`
}
