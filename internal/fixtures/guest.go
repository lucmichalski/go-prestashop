package fixtures

import (
	"encoding/xml"
)

// Guest was generated 2021-02-06 11:38:40 by evolutive on eg-cdn.gsi-network.com.
type Guest struct {
	XMLName xml.Name `xml:"entity_guest"`
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
		Text  string `xml:",chardata"`
		Guest []struct {
			Text              string `xml:",chardata"`
			IDOperatingSystem string `xml:"id_operating_system,attr"`
			IDWebBrowser      string `xml:"id_web_browser,attr"`
			IDCustomer        string `xml:"id_customer,attr"`
			ID                string `xml:"id,attr"`
			Javascript        string `xml:"javascript,attr"`
			AppleQuicktime    string `xml:"apple_quicktime,attr"`
			RealPlayer        string `xml:"real_player,attr"`
			WindowsMedia      string `xml:"windows_media,attr"`
			ScreenResolutionX string `xml:"screen_resolution_x,attr"`
			ScreenResolutionY string `xml:"screen_resolution_y,attr"`
			ScreenColor       string `xml:"screen_color,attr"`
			SunJava           string `xml:"sun_java,attr"`
			AdobeFlash        string `xml:"adobe_flash,attr"`
			AdobeDirector     string `xml:"adobe_director,attr"`
			AcceptLanguage    string `xml:"accept_language,attr"`
		} `xml:"guest"`
	} `xml:"entities"`
}
