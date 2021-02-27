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

type WkMpSellerProduct struct {
	IdMpProduct             int // primary_key
	IdSeller                int
	IdPsProduct             int
	IdPsShop                int
	IdCategory              int
	Price                   float64
	WholesalePrice          float64
	Unity                   string
	UnitPrice               float64
	IdTaxRulesGroup         int
	OnSale                  bool
	AdditionalShippingCost  float64
	Quantity                int
	MinimalQuantity         int
	LowStockThreshold       int
	LowStockAlert           bool
	Active                  bool
	StatusBeforeDeactivate  bool
	ShowCondition           bool
	Condition               string
	AvailableForOrder       bool
	ShowPrice               bool
	OnlineOnly              bool
	Visibility              string
	AdminAssigned           bool
	Width                   float64
	Height                  float64
	Depth                   float64
	Weight                  float64
	Reference               string
	Ean13                   string
	Upc                     string
	Isbn                    string
	OutOfStock              int
	AvailableDate           time.Time
	PsIdCarrierReference    string
	AdminApproved           bool
	AdditionalDeliveryTimes bool
	DateAdd                 time.Time
	DateUpd                 time.Time
	CsvRequestNo            string
}

type WkMpSellerProductCategory struct {
	IdMpCategoryProduct int // primary_key
	IdCategory          int
	IdSellerProduct     int
	IsDefault           bool
}

type WkMpSellerProductImage struct {
	IdMpProductImage       int // primary_key
	SellerProductId        int
	SellerProductImageName string
	IdPsImage              int
	Position               int
	Cover                  bool
	Active                 bool
}

type WkMpSellerProductLang struct {
	IdMpProduct      int // primary_key
	IdLang           int
	ProductName      string
	ShortDescription string
	Description      string
	AvailableNow     string
	AvailableLater   string
	MetaTitle        string
	MetaDescription  string
	LinkRewrite      string
	DeliveryInStock  string
	DeliveryOutStock string
}
