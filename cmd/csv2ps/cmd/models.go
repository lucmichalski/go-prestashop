package cmd

type CatalogMap struct {
	Name      string   `yaml:"name"`
	Feed      string   `yaml:"feed"`
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
	Attributes       []string `yaml:"attributes"`
	Features         []string `yaml:"features"`
}

type Feature struct {
	IDFeature      uint
	IDFeatureValue uint
}
