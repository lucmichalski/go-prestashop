package cmd

import (
	"time"
)

type CatalogMap struct {
	Name      string   `yaml:"name"`
	Feed      string   `yaml:"feed"`
	SellerId  int      `yaml:"seller_id"`
	Separator string   `yaml:"separator"`
	Fields    []string `yaml:"fields"`
	Total     int      `yaml:"total"`
	Mapping   struct {
		Update     bool     `yaml:"update" default:"true"` // Either ADD
		LangSuffix []string `yaml:"multi_language_suffix"`
		LangFields []string `yaml:"multi_language_fields"`
		Product    Product  `yaml:"product"`
		Category   struct {
			Name       string `yaml:"name"`
			Breadcrumb string `yaml:"breaddcrumb"`
			Separator  string `yaml:"separator"`
		} `yaml:"category"`
	} `yaml:"mapping"`
}

type Product struct {
	Name             string   `yaml:"name"`
	Reference        string   `yaml:"reference"`
	Ean13            string   `yaml:"ean13"`
	Sku              string   `yaml:"sku"`
	Mpn              string   `yaml:"mpn"`
	Price            string   `yaml:"price"`
	Description      string   `yaml:"description"`
	DescriptionShort string   `yaml:"description_short"`
	Image            string   `yaml:"image"`
	Quantity         string   `yaml:"quantity"`
	Width            string   `yaml:"width"`
	Height           string   `yaml:"height"`
	Weight           string   `yaml:"weight"`
	Shipping         string   `yaml:"shipping"`
	Redirect         string   `yaml:"redirect"`
	Attributes       []string `yaml:"attributes"`
	Features         []string `yaml:"features"`
}

type Feature struct {
	IDFeature      uint
	IDFeatureValue uint
}

type WkMpSeller struct {
	IdSeller                int
	ShopNameUnique          string
	LinkRewrite             string
	SellerFirstname         string
	SellerLastname          string
	BusinessEmail           string
	Phone                   string
	Fax                     string
	Address                 string
	Postcode                string
	City                    string
	IdCountry               int
	IdState                 int
	TaxIdentificationNumber string
	DefaultLang             int
	FacebookId              string
	TwitterId               string
	GoogleId                string
	InstagramId             string
	ProfileImage            string
	ProfileBanner           string
	ShopImage               string
	ShopBanner              string
	Active                  bool
	ShopApproved            bool
	SellerCustomerId        int
	SellerDetailsAccess     string
	DateAdd                 time.Time
	DateUpd                 time.Time
}

type WkMpSellerLang struct {
	IdSeller  int
	IdLang    int
	ShopName  string
	AboutShop string
}
