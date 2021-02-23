package cmd

import (
	"time"

	"gorm.io/datatypes"
	// "gorm.io/gorm"
)

/*
code	url	creator	created_t	created_datetime	last_modified_t	last_modified_datetime	product_name	abbreviated_product_name	generic_name	quantity	packaging	packaging_tags	packaging_text	brands	brands_tags	categories	categories_tags	categories_en	origins	origins_tags	origins_en	manufacturing_places	manufacturing_places_tags	labels	labels_tags	labels_en	emb_codes	emb_codes_tags	first_packaging_code_geo	cities	cities_tags	purchase_places	stores	countries	countries_tags	countries_en	ingredients_text	allergens	allergens_en	traces	traces_tags	traces_en	serving_size	serving_quantity	no_nutriments	additives_n	additives	additives_tags	additives_eningredients_from_palm_oil_n	ingredients_from_palm_oil	ingredients_from_palm_oil_tags	ingredients_that_may_be_from_palm_oil_n	ingredients_that_may_be_from_palm_oil	ingredients_that_may_be_from_palm_oil_tags	nutriscore_score	nutriscore_grade	nova_group	pnns_groups_1	pnns_groups_2	states	states_tags	states_en	brand_owner	main_category	main_category_en	image_url	image_small_url	image_ingredients_url	image_ingredients_small_url	image_nutrition_url	image_nutrition_small_url	energy-kj_100g	energy-kcal_100g	energy_100g	energy-from-fat_100g	fat_100g	saturated-fat_100g	-butyric-acid_100g	-caproic-acid_100g	-caprylic-acid_100g	-capric-acid_100g	-lauric-acid_100g	-myristic-acid_100g-palmitic-acid_100g	-stearic-acid_100g	-arachidic-acid_100g	-behenic-acid_100g	-lignoceric-acid_100g	-cerotic-acid_100g	-montanic-acid_100g	-melissic-acid_100g	monounsaturated-fat_100g	polyunsaturated-fat_100g	omega-3-fat_100g	-alpha-linolenic-acid_100g	-eicosapentaenoic-acid_100g	-docosahexaenoic-acid_100g	omega-6-fat_100g	-linoleic-acid_100g	-arachidonic-acid_100g	-gamma-linolenic-acid_100g	-dihomo-gamma-linolenic-acid_100g	omega-9-fat_100g	-oleic-acid_100g	-elaidic-acid_100g	-gondoic-acid_100g	-mead-acid_100g	-erucic-acid_100g	-nervonic-acid_100g	trans-fat_100g	cholesterol_100g	carbohydrates_100g	sugars_100g	-sucrose_100g	-glucose_100g	-fructose_100g	-lactose_100g	-maltose_100g	-maltodextrins_100g	starch_100g	polyols_100g	fiber_100g	-soluble-fiber_100g	-insoluble-fiber_100g	proteins_100g	casein_100g	serum-proteins_100g	nucleotides_100g	salt_100g	sodium_100g	alcohol_100g	vitamin-a_100g	beta-carotene_100g	vitamin-d_100g	vitamin-e_100g	vitamin-k_100g	vitamin-c_100g	vitamin-b1_100g	vitamin-b2_100g	vitamin-pp_100g	vitamin-b6_100g	vitamin-b9_100g	folates_100g	vitamin-b12_100g	biotin_100g	pantothenic-acid_100g	silica_100g	bicarbonate_100g	potassium_100g	chloride_100g	calcium_100g	phosphorus_100g	iron_100g	magnesium_100g	zinc_100g	copper_100g	manganese_100g	fluoride_100g	selenium_100g	chromium_100g	molybdenum_100g	iodine_100g	caffeine_100g	taurine_100g	ph_100g	fruits-vegetables-nuts_100g	fruits-vegetables-nuts-dried_100g	fruits-vegetables-nuts-estimate_100g	collagen-meat-protein-ratio_100g	cocoa_100g	chlorophyl_100g	carbon-footprint_100g	carbon-footprint-from-meat-or-fish_100g	nutrition-score-fr_100g	nutrition-score-uk_100g	glycemic-index_100g	water-hardness_100g	choline_100g	phylloquinone_100g	beta-glucan_100g	inositol_100g	carnitine_100g
0000000000017	http://world-en.openfoodfacts.org/product/0000000000017/vitoria-crackers	kiliweb	1529059080	2018-06-15T10:38:00Z	1561463718	2019-06-25T11:55:18Z	Vitória crackers	France	en:france	France																							unknown	unknown	en:to-be-completed, en:nutrition-facts-completed, en:ingredients-to-be-completed, en:expiration-date-to-be-completed, en:packaging-code-to-be-completed, en:characteristics-to-be-completed, en:categories-to-be-completed, en:brands-to-be-completed, en:packaging-to-be-completed, en:quantity-to-be-completed, en:product-name-completed, en:photos-to-be-validated, en:photos-uploaded	en:to-be-completed,en:nutrition-facts-completed,en:ingredients-to-be-completed,en:expiration-date-to-be-completed,en:packaging-code-to-be-completed,en:characteristics-to-be-completed,en:categories-to-be-completed,en:brands-to-be-completed,en:packaging-to-be-completed,en:quantity-to-be-completed,en:product-name-completed,en:photos-to-be-validated,en:photos-uploaded	To be completed,Nutrition facts completed,Ingredients to be completed,Expiration date to be completed,Packaging code to be completed,Characteristics to be completed,Categories to be completed,Brands to be completed,Packaging to be completed,Quantity to be completed,Product name completed,Photos to be validated,Photos uploaded				https://static.openfoodfacts.org/images/products/000/000/000/0017/front_fr.4.400.jpg	https://static.openfoodfacts.org/images/products/000/000/000/0017/front_fr.4.200.jpg	https://static.openfoodfacts.org/images/products/000/000/000/0017/ingredients_fr.9.400.jpg	https://static.openfoodfacts.org/images/products/000/000/000/0017/ingredients_fr.9.200.jpg				375	1569		7	3.08												70.1	15												7.8				1.4	0.56
*/

var Tables = []interface{}{
	&OpenFoodFact{},
}

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
	Name             string `yaml:"name"`
	Reference        string `yaml:"reference"`
	Price            string `yaml:"price"`
	Description      string `yaml:"description"`
	DescriptionShort string `yaml:"description_short"`
	Image            string `yaml:"image"`
	Quantity         string `yaml:"quantity"`
	Redirect         string `yaml:"redirect"`
	// Attributes  []string `yaml:"attributes"`
	Features []string `yaml:"features"`
}

type Feature struct {
	IDFeature      uint
	IDFeatureValue uint
}

type OpenFoodFact struct {
	Additives                string         `json:"additives,omitempty" csv:"additives,omitempty"`
	AdditivesN               string         `json:"additives_n,omitempty" csv:"additives_n,omitempty"`
	AdditivesTags            string         `json:"additives_tags,omitempty" csv:"additives_tags,omitempty"`
	Allergens                string         `json:"allergens,omitempty" csv:"allergens,omitempty"`
	AllergensEn              string         `json:"allergens_en,omitempty" csv:"allergens_en,omitempty"`
	AbbreviatedProductName   string         `json:"abbreviated_product_name,omitempty" csv:"abbreviated_product_name,omitempty"`
	AdditivesEn              string         `json:"additives_en,omitempty" csv:"additives_en,omitempty"`
	BrandOwner               string         `json:"brand_owner,omitempty" csv:"brand_owner,omitempty"`
	Brands                   string         `json:"brands,omitempty" csv:"brands,omitempty"`
	BrandsTags               string         `json:"brands_tags,omitempty" csv:"brands_tags,omitempty"`
	Categories               string         `json:"categories,omitempty" csv:"categories,omitempty"`
	CategoriesEn             string         `json:"categories_en,omitempty" csv:"categories_en,omitempty"`
	CategoriesTags           string         `json:"categories_tags,omitempty" csv:"categories_tags,omitempty"`
	Cities                   string         `json:"cities,omitempty" csv:"cities,omitempty"`
	CitiesTags               string         `json:"cities_tags,omitempty" csv:"cities_tags,omitempty"`
	Countries                string         `json:"countries,omitempty" csv:"countries,omitempty"`
	CountriesEn              string         `json:"countries_en,omitempty" csv:"countries_en,omitempty"`
	CountriesTags            string         `json:"countries_tags,omitempty" csv:"countries_tags,omitempty"`
	CreatedDatetime          datatypes.Date `json:"created_datetime,omitempty" csv:"created_datetime,omitempty"`
	CreatedT                 time.Time      `json:"created_t,omitempty" csv:"created_t,omitempty"`
	Creator                  string         `json:"creator,omitempty" csv:"creator,omitempty"`
	EmbCodes                 string         `json:"emb_codes,omitempty" csv:"emb_codes,omitempty"`
	EmbCodesTags             string         `json:"emb_codes_tags,omitempty" csv:"emb_codes_tags,omitempty"`
	Code                     string         `gorm:"uniqueIndex:code_idx;" json:"code,omitempty" csv:"code,omitempty"`
	GenericName              string         `json:"generic_name,omitempty" csv:"generic_name,omitempty"`
	ImageIngredientsSmallURL string         `json:"image_ingredients_small_url,omitempty" csv:"image_ingredients_small_url,omitempty"`
	ImageIngredientsURL      string         `json:"image_ingredients_url,omitempty" csv:"image_ingredients_url,omitempty"`
	ImageNutritionSmallURL   string         `json:"image_nutrition_small_url,omitempty" csv:"image_nutrition_small_url,omitempty"`
	ImageNutritionURL        string         `json:"image_nutrition_url,omitempty" csv:"image_nutrition_url,omitempty"`
	ImageSmallURL            string         `json:"image_small_url,omitempty" csv:"image_small_url,omitempty"`
	ImageURL                 string         `json:"image_url,omitempty" csv:"image_url,omitempty"`
	Labels                   string         `json:"labels,omitempty" csv:"labels,omitempty"`
	LabelsEn                 string         `json:"labels_en,omitempty" csv:"labels_en,omitempty"`
	LabelsTags               string         `json:"labels_tags,omitempty" csv:"labels_tags,omitempty"`
	LastModifiedDatetime     time.Time      `json:"last_modified_datetime,omitempty" csv:"last_modified_datetime,omitempty"`
	LastModifiedT            time.Time      `json:"last_modified_t,omitempty" csv:"last_modified_t,omitempty"`
	NoNutriments             string         `json:"no_nutriments,omitempty" csv:"no_nutriments,omitempty"`
	NovaGroup                string         `json:"nova_group,omitempty" csv:"nova_group,omitempty"`
	MainCategory             string         `json:"main_category,omitempty" csv:"main_category,omitempty"`
	MainCategoryEn           string         `json:"main_category_en,omitempty" csv:"main_category_en,omitempty"`
	PurchasePlaces           string         `json:"purchase_places,omitempty" csv:"purchase_places,omitempty"`
	ManufacturingPlaces      string         `json:"manufacturing_places,omitempty" csv:"manufacturing_places,omitempty"`
	ManufacturingPlacesTags  string         `json:"manufacturing_places_tags,omitempty" csv:"manufacturing_places_tags,omitempty"`
	NutriscoreGrade          string         `json:"nutriscore_grade,omitempty" csv:"nutriscore_grade,omitempty"`
	NutriscoreScore          string         `json:"nutriscore_score,omitempty" csv:"nutriscore_score,omitempty"`
	Origins                  string         `json:"origins,omitempty" csv:"origins,omitempty"`
	OriginsEn                string         `json:"origins_en,omitempty" csv:"origins_en,omitempty"`
	OriginsTags              string         `json:"origins_tags,omitempty" csv:"origins_tags,omitempty"`
	Packaging                string         `json:"packaging,omitempty" csv:"packaging,omitempty"`
	PackagingTags            string         `json:"packaging_tags,omitempty" csv:"packaging_tags,omitempty"`
	PackagingText            string         `json:"packaging_text,omitempty" csv:"packaging_text,omitempty"`
	ProductName              string         `json:"product_name,omitempty" csv:"product_name,omitempty"`
	Quantity                 string         `json:"quantity,omitempty" csv:"quantity,omitempty"`
	ServingQuantity          string         `json:"serving_quantity,omitempty" csv:"serving_quantity,omitempty"`
	ServingSize              string         `json:"serving_size,omitempty" csv:"serving_size,omitempty"`
	States                   string         `json:"states,omitempty" csv:"states,omitempty"`
	StatesEn                 string         `json:"states_en,omitempty" csv:"states_en,omitempty"`
	StatesTags               string         `json:"states_tags,omitempty" csv:"states_tags,omitempty"`
	Stores                   string         `json:"stores,omitempty" csv:"stores,omitempty"`
	URL                      string         `json:"url,omitempty" csv:"url,omitempty"`
	Traces                   string         `json:"traces,omitempty" csv:"traces,omitempty"`
	TracesEn                 string         `json:"traces_en,omitempty" csv:"traces_en,omitempty"`
	TracesTags               string         `json:"traces_tags,omitempty" csv:"traces_tags,omitempty"`
}
