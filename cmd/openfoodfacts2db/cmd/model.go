package cmd

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var Tables = []interface{}{
	&OpenFoodFact{},
}

type Product struct {
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

type Attribute struct {
}

type Feature struct {
	IngredientsFromPalmOil              string `json:"ingredients_from_palm_oil,omitempty" csv:"ingredients_from_palm_oil,omitempty"`
	IngredientsFromPalmOilTags          string `json:"ingredients_from_palm_oil_tags,omitempty" csv:"ingredients_from_palm_oil_tags,omitempty"`
	IngredientsText                     string `json:"ingredients_text,omitempty" csv:"ingredients_text,omitempty"`
	IngredientsThatMayBeFromPalmOil     string `json:"ingredients_that_may_be_from_palm_oil,omitempty" csv:"ingredients_that_may_be_from_palm_oil,omitempty"`
	IngredientsThatMayBeFromPalmOilN    string `json:"ingredients_that_may_be_from_palm_oil_n,omitempty" csv:"ingredients_that_may_be_from_palm_oil_n,omitempty"`
	IngredientsThatMayBeFromPalmOilTags string `json:"ingredients_that_may_be_from_palm_oil_tags,omitempty" csv:"ingredients_that_may_be_from_palm_oil_tags,omitempty"`
	IngredientsFromPalmOilN             string `json:"ingredients_from_palm_oil_n,omitempty" csv:"ingredients_from_palm_oil_n,omitempty"`
	AlphaLinolenicAcid100g              string `gorm:"column:alpha-linolenic-acid_100g;type:varchar(64)" json:"-alpha-linolenic-acid_100g,omitempty" csv:"-alpha-linolenic-acid_100g,omitempty"`
	ArachidicAcid100g                   string `gorm:"column:arachidic-acid_100g;type:varchar(64)" json:"-arachidic-acid_100g,omitempty" csv:"-arachidic-acid_100g,omitempty"`
	ArachidonicAcid100g                 string `gorm:"column:arachidonic-acid_100g;type:varchar(64)" json:"-arachidonic-acid_100g,omitempty" csv:"-arachidonic-acid_100g,omitempty"`
	BehenicAcid100g                     string `gorm:"column:behenic-acid_100g;type:varchar(64)" json:"-behenic-acid_100g,omitempty" csv:"-behenic-acid_100g,omitempty"`
	ButyricAcid100g                     string `gorm:"column:butyric-acid_100g;type:varchar(64)" json:"-butyric-acid_100g,omitempty" csv:"-butyric-acid_100g,omitempty"`
	CapricAcid100g                      string `gorm:"column:capric-acid_100g;type:varchar(64)" json:"-capric-acid_100g,omitempty" csv:"-capric-acid_100g,omitempty"`
	CaproicAcid100g                     string `gorm:"column:caproic-acid_100g;type:varchar(64)" json:"-caproic-acid_100g,omitempty" csv:"-caproic-acid_100g,omitempty"`
	CaprylicAcid100g                    string `gorm:"column:caprylic-acid_100g;type:varchar(64)" json:"-caprylic-acid_100g,omitempty" csv:"-caprylic-acid_100g,omitempty"`
	CeroticAcid100g                     string `gorm:"column:cerotic-acid_100g;type:varchar(64)" json:"-cerotic-acid_100g,omitempty" csv:"-cerotic-acid_100g,omitempty"`
	DihomoGammaLinolenicAcid100g        string `gorm:"column:dihomo-gamma-linolenic-acid_100g;type:varchar(64)" json:"-dihomo-gamma-linolenic-acid_100g,omitempty" csv:"-dihomo-gamma-linolenic-acid_100g,omitempty"`
	DocosahexaenoicAcid100g             string `gorm:"column:docosahexaenoic-acid_100g;type:varchar(64)" json:"-docosahexaenoic-acid_100g,omitempty" csv:"-docosahexaenoic-acid_100g,omitempty"`
	EicosapentaenoicAcid100g            string `gorm:"column:eicosapentaenoic-acid_100g;type:varchar(64)" json:"-eicosapentaenoic-acid_100g,omitempty" csv:"-eicosapentaenoic-acid_100g,omitempty"`
	ElaidicAcid100g                     string `gorm:"column:elaidic-acid_100g;type:varchar(64)" json:"-elaidic-acid_100g,omitempty" csv:"-elaidic-acid_100g,omitempty"`
	ErucicAcid100g                      string `gorm:"column:erucic-acid_100g;type:varchar(64)" json:"-erucic-acid_100g,omitempty" csv:"-erucic-acid_100g,omitempty"`
	Fructose100g                        string `gorm:"column:fructose_100g;type:varchar(64)" json:"-fructose_100g,omitempty" csv:"-fructose_100g,omitempty"`
	GammaLinolenicAcid100g              string `gorm:"column:gamma-linolenic-acid_100g;type:varchar(64)" json:"-gamma-linolenic-acid_100g,omitempty" csv:"-gamma-linolenic-acid_100g,omitempty"`
	Glucose100g                         string `gorm:"column:glucose_100g;type:varchar(64)" json:"-glucose_100g,omitempty" csv:"-glucose_100g,omitempty"`
	GondoicAcid100g                     string `gorm:"column:gondoic-acid_100g;type:varchar(64)" json:"-gondoic-acid_100g,omitempty" csv:"-gondoic-acid_100g,omitempty"`
	InsolubleFiber100g                  string `gorm:"column:insoluble-fiber_100g;type:varchar(64)" json:"-insoluble-fiber_100g,omitempty" csv:"-insoluble-fiber_100g,omitempty"`
	Lactose100g                         string `gorm:"column:lactose_100g;type:varchar(64)" json:"-lactose_100g,omitempty" csv:"-lactose_100g,omitempty"`
	LauricAcid100g                      string `gorm:"column:lauric-acid_100g;type:varchar(64)" json:"-lauric-acid_100g,omitempty" csv:"-lauric-acid_100g,omitempty"`
	LignocericAcid100g                  string `gorm:"column:lignoceric-acid_100g;type:varchar(64)" json:"-lignoceric-acid_100g,omitempty" csv:"-lignoceric-acid_100g,omitempty"`
	LinoleicAcid100g                    string `gorm:"column:linoleic-acid_100g;type:varchar(64)" json:"-linoleic-acid_100g,omitempty" csv:"-linoleic-acid_100g,omitempty"`
	Maltodextrins100g                   string `gorm:"column:maltodextrins_100g;type:varchar(64)" json:"-maltodextrins_100g,omitempty" csv:"-maltodextrins_100g,omitempty"`
	Maltose100g                         string `gorm:"column:maltose_100g;type:varchar(64)" json:"-maltose_100g,omitempty" csv:"-maltose_100g,omitempty"`
	MeadAcid100g                        string `gorm:"column:mead-acid_100g;type:varchar(64)" json:"-mead-acid_100g,omitempty" csv:"-mead-acid_100g,omitempty"`
	MelissicAcid100g                    string `gorm:"column:melissic-acid_100g;type:varchar(64)" json:"-melissic-acid_100g,omitempty" csv:"-melissic-acid_100g,omitempty"`
	MontanicAcid100g                    string `gorm:"column:montanic-acid_100g;type:varchar(64)" json:"-montanic-acid_100g,omitempty" csv:"-montanic-acid_100g,omitempty"`
	MyristicAcid100g                    string `gorm:"column:myristic-acid_100g;type:varchar(64)" json:"-myristic-acid_100g,omitempty" csv:"-myristic-acid_100g,omitempty"`
	PalmiticAcid100g                    string `gorm:"column:palmitic-acid_100g;type:varchar(64)" json:"-palmitic-acid_100g,omitempty" csv:"-palmitic-acid_100g,omitempty"`
	NervonicAcid100g                    string `gorm:"column:nervonic-acid_100g;type:varchar(64)" json:"-nervonic-acid_100g,omitempty" csv:"-nervonic-acid_100g,omitempty"`
	OleicAcid100g                       string `gorm:"column:oleic-acid_100g;type:varchar(64)" json:"-oleic-acid_100g,omitempty" csv:"-oleic-acid_100g,omitempty"`
	SolubleFiber100g                    string `gorm:"column:soluble-fiber_100g;type:varchar(64)" json:"-soluble-fiber_100g,omitempty" csv:"-soluble-fiber_100g,omitempty"`
	StearicAcid100g                     string `gorm:"column:stearic-acid_100g;type:varchar(64)" json:"-stearic-acid_100g,omitempty" csv:"-stearic-acid_100g,omitempty"`
	Sucrose100g                         string `gorm:"column:sucrose_100g;type:varchar(64)" json:"-sucrose_100g,omitempty" csv:"-sucrose_100g,omitempty"`
	Alcohol100g                         string `gorm:"column:alcohol_100g;type:varchar(64)" json:"alcohol_100g,omitempty" csv:"alcohol_100g,omitempty"`
	BetaCarotene100g                    string `gorm:"column:beta-carotene_100g;type:varchar(64)" json:"beta-carotene_100g,omitempty" csv:"beta-carotene_100g,omitempty"`
	BetaGlucan100g                      string `gorm:"column:beta-glucan_100g;type:varchar(64)" json:"beta-glucan_100g,omitempty" csv:"beta-glucan_100g,omitempty"`
	Bicarbonate100g                     string `gorm:"column:bicarbonate_100g;type:varchar(64)" json:"bicarbonate_100g,omitempty" csv:"bicarbonate_100g,omitempty"`
	Biotin100g                          string `gorm:"column:biotin_100g;type:varchar(64)" json:"biotin_100g,omitempty" csv:"biotin_100g,omitempty"`
	Caffeine100g                        string `gorm:"column:caffeine_100g;type:varchar(64)" json:"caffeine_100g,omitempty" csv:"caffeine_100g,omitempty"`
	Calcium100g                         string `gorm:"column:calcium_100g;type:varchar(64)" json:"calcium_100g,omitempty" csv:"calcium_100g,omitempty"`
	Carbohydrates100g                   string `gorm:"column:carbohydrates_100g;type:varchar(64)" json:"carbohydrates_100g,omitempty" csv:"carbohydrates_100g,omitempty"`
	CarbonFootprintFromMeatOrFish100g   string `gorm:"column:carbon-footprint-from-meat-or-fish_100g;type:varchar(64)" json:"carbon-footprint-from-meat-or-fish_100g,omitempty" csv:"carbon-footprint-from-meat-or-fish_100g,omitempty"`
	CarbonFootprint100g                 string `gorm:"column:carbon-footprint_100g;type:varchar(64)" json:"carbon-footprint_100g,omitempty" csv:"carbon-footprint_100g,omitempty"`
	Carnitine100g                       string `gorm:"column:carnitine_100g;type:varchar(64)" json:"carnitine_100g,omitempty" csv:"carnitine_100g,omitempty"`
	Casein100g                          string `gorm:"column:casein_100g;type:varchar(64)" json:"casein_100g,omitempty" csv:"casein_100g,omitempty"`
	Chloride100g                        string `gorm:"column:chloride_100g;type:varchar(64)" json:"chloride_100g,omitempty" csv:"chloride_100g,omitempty"`
	Chlorophyl100g                      string `gorm:"column:chlorophyl_100g;type:varchar(64)" json:"chlorophyl_100g,omitempty" csv:"chlorophyl_100g,omitempty"`
	Cholesterol100g                     string `gorm:"column:cholesterol_100g;type:varchar(64)" json:"cholesterol_100g,omitempty" csv:"cholesterol_100g,omitempty"`
	Choline100g                         string `gorm:"column:choline_100g;type:varchar(64)" json:"choline_100g,omitempty" csv:"choline_100g,omitempty"`
	Chromium100g                        string `gorm:"column:chromium_100g;type:varchar(64)" json:"chromium_100g,omitempty" csv:"chromium_100g,omitempty"`
	Cocoa100g                           string `gorm:"column:cocoa_100g;type:varchar(64)" json:"cocoa_100g,omitempty" csv:"cocoa_100g,omitempty"`
	CollagenMeatProteinRatio100g        string `gorm:"column:collagen-meat-protein-ratio_100g;type:varchar(64)" json:"collagen-meat-protein-ratio_100g,omitempty" csv:"collagen-meat-protein-ratio_100g,omitempty"`
	Copper100g                          string `gorm:"column:copper_100g;type:varchar(64)" json:"copper_100g,omitempty" csv:"copper_100g,omitempty"`
	EnergyFromFat100g                   string `gorm:"column:energy-from-fat_100g;type:varchar(64)" json:"energy-from-fat_100g,omitempty" csv:"energy-from-fat_100g,omitempty"`
	EnergyKcal100g                      string `gorm:"column:energy-kcal_100g;type:varchar(64)" json:"energy-kcal_100g,omitempty" csv:"energy-kcal_100g,omitempty"`
	EnergyKj100g                        string `gorm:"column:energy-kj_100g;type:varchar(64)" json:"energy-kj_100g,omitempty" csv:"energy-kj_100g,omitempty"`
	Energy100g                          string `gorm:"column:energy_100g;type:varchar(64)" json:"energy_100g,omitempty" csv:"energy_100g,omitempty"`
	Fat100g                             string `gorm:"column:fat_100g;type:varchar(64)" json:"fat_100g,omitempty" csv:"fat_100g,omitempty"`
	Fiber100g                           string `gorm:"column:fiber_100g;type:varchar(64)" json:"fiber_100g,omitempty" csv:"fiber_100g,omitempty"`
	FirstPackagingCodeGeo               string `gorm:"column:first_packaging_code_geo;type:varchar(64)" json:"first_packaging_code_geo,omitempty" csv:"first_packaging_code_geo,omitempty"`
	Fluoride100g                        string `gorm:"column:fluoride_100g;type:varchar(64)" json:"fluoride_100g,omitempty" csv:"fluoride_100g,omitempty"`
	Folates100g                         string `gorm:"column:folates_100g;type:varchar(64)" json:"folates_100g,omitempty" csv:"folates_100g,omitempty"`
	FruitsVegetablesNutsDried100g       string `gorm:"column:fruits-vegetables-nuts-dried_100g;type:varchar(64)" json:"fruits-vegetables-nuts-dried_100g,omitempty" csv:"fruits-vegetables-nuts-dried_100g,omitempty"`
	FruitsVegetablesNutsEstimate100g    string `gorm:"column:fruits-vegetables-nuts-estimate_100g;type:varchar(64)" json:"fruits-vegetables-nuts-estimate_100g,omitempty" csv:"fruits-vegetables-nuts-estimate_100g,omitempty"`
	FruitsVegetablesNuts100g            string `gorm:"column:fruits-vegetables-nuts_100g;type:varchar(64)" json:"fruits-vegetables-nuts_100g,omitempty" csv:"fruits-vegetables-nuts_100g,omitempty"`
	GlycemicIndex100g                   string `gorm:"column:glycemic-index_100g;type:varchar(64)" json:"glycemic-index_100g,omitempty" csv:"glycemic-index_100g,omitempty"`
	Inositol100g                        string `gorm:"column:inositol_100g;type:varchar(64)" json:"inositol_100g,omitempty" csv:"inositol_100g,omitempty"`
	Iodine100g                          string `gorm:"column:iodine_100g;type:varchar(64)" json:"iodine_100g,omitempty" csv:"iodine_100g,omitempty"`
	Iron100g                            string `gorm:"column:iron_100g;type:varchar(64)" json:"iron_100g,omitempty" csv:"iron_100g,omitempty"`
	Magnesium100g                       string `gorm:"column:magnesium_100g;type:varchar(64)" json:"magnesium_100g,omitempty" csv:"magnesium_100g,omitempty"`
	Manganese100g                       string `gorm:"column:manganese_100g;type:varchar(64)" json:"manganese_100g,omitempty" csv:"manganese_100g,omitempty"`
	Molybdenum100g                      string `gorm:"column:molybdenum_100g;type:varchar(64)" json:"molybdenum_100g,omitempty" csv:"molybdenum_100g,omitempty"`
	MonounsaturatedFat100g              string `gorm:"column:monounsaturated-fat_100g;type:varchar(64)" json:"monounsaturated-fat_100g,omitempty" csv:"monounsaturated-fat_100g,omitempty"`
	Nucleotides100g                     string `gorm:"column:nucleotides_100g;type:varchar(64)" json:"nucleotides_100g,omitempty" csv:"nucleotides_100g,omitempty"`
	NutritionScoreFr100g                string `gorm:"column:nutrition-score-fr_100g;type:varchar(64)" json:"nutrition-score-fr_100g,omitempty" csv:"nutrition-score-fr_100g,omitempty"`
	NutritionScoreUk100g                string `gorm:"column:nutrition-score-uk_100g;type:varchar(64)" json:"nutrition-score-uk_100g,omitempty" csv:"nutrition-score-uk_100g,omitempty"`
	Omega3Fat100g                       string `gorm:"column:omega-3-fat_100g;type:varchar(64)" json:"omega-3-fat_100g,omitempty" csv:"omega-3-fat_100g,omitempty"`
	Omega6Fat100g                       string `gorm:"column:omega-6-fat_100g;type:varchar(64)" json:"omega-6-fat_100g,omitempty" csv:"omega-6-fat_100g,omitempty"`
	Omega9Fat100g                       string `gorm:"column:omega-9-fat_100g;type:varchar(64)" json:"omega-9-fat_100g,omitempty" csv:"omega-9-fat_100g,omitempty"`
	PantothenicAcid100g                 string `gorm:"column:pantothenic-acid_100g;type:varchar(64)" json:"pantothenic-acid_100g,omitempty" csv:"pantothenic-acid_100g,omitempty"`
	Ph100g                              string `gorm:"column:ph_100g;type:varchar(64)" json:"ph_100g,omitempty" csv:"ph_100g,omitempty"`
	Phosphorus100g                      string `gorm:"column:phosphorus_100g;type:varchar(64)" json:"phosphorus_100g,omitempty" csv:"phosphorus_100g,omitempty"`
	Phylloquinone100g                   string `gorm:"column:phylloquinone_100g;type:varchar(64)" json:"phylloquinone_100g,omitempty" csv:"phylloquinone_100g,omitempty"`
	PnnsGroups1                         string `gorm:"column:pnns_groups_1;type:varchar(64)" json:"pnns_groups_1,omitempty" csv:"pnns_groups_1,omitempty"`
	PnnsGroups2                         string `gorm:"column:pnns_groups_2;type:varchar(64)" json:"pnns_groups_2,omitempty" csv:"pnns_groups_2,omitempty"`
	Polyols100g                         string `gorm:"column:polyols_100g;type:varchar(64)" json:"polyols_100g,omitempty" csv:"polyols_100g,omitempty"`
	PolyunsaturatedFat100g              string `gorm:"column:polyunsaturated-fat_100g;type:varchar(64)" json:"polyunsaturated-fat_100g,omitempty" csv:"polyunsaturated-fat_100g,omitempty"`
	Potassium100g                       string `gorm:"column:potassium_100g;type:varchar(64)" json:"potassium_100g,omitempty" csv:"potassium_100g,omitempty"`
	Proteins100g                        string `gorm:"column:proteins_100g;type:varchar(64)" json:"proteins_100g,omitempty" csv:"proteins_100g,omitempty"`
	Salt100g                            string `gorm:"column:salt_100g;type:varchar(64)" json:"salt_100g,omitempty" csv:"salt_100g,omitempty"`
	SaturatedFat100g                    string `gorm:"column:saturated-fat_100g;type:varchar(64)"  json:"saturated-fat_100g,omitempty" csv:"saturated-fat_100g,omitempty"`
	Selenium100g                        string `gorm:"column:selenium_100g;type:varchar(64)" json:"selenium_100g,omitempty" csv:"selenium_100g,omitempty"`
	SerumProteins100g                   string `gorm:"column:serum-proteins_100g;type:varchar(64)" json:"serum-proteins_100g,omitempty" csv:"serum-proteins_100g,omitempty"`
	Silica100g                          string `gorm:"column:silica_100g;type:varchar(64)" json:"silica_100g,omitempty" csv:"silica_100g,omitempty"`
	Sodium100g                          string `gorm:"column:sodium_100g;type:varchar(64)" json:"sodium_100g,omitempty" csv:"sodium_100g,omitempty"`
	Starch100g                          string `gorm:"column:starch_100g;type:varchar(64)" json:"starch_100g,omitempty" csv:"starch_100g,omitempty"`
	Sugars100g                          string `gorm:"column:sugars_100g;type:varchar(64)" json:"sugars_100g,omitempty" csv:"sugars_100g,omitempty"`
	Taurine100g                         string `gorm:"column:taurine_100g;type:varchar(64)" json:"taurine_100g,omitempty" csv:"taurine_100g,omitempty"`
	TransFat100g                        string `gorm:"column:trans-fat_100g;type:varchar(64)" json:"trans-fat_100g,omitempty" csv:"trans-fat_100g,omitempty"`
	VitaminA100g                        string `gorm:"column:vitamin-a_100g;type:varchar(64)" json:"vitamin-a_100g,omitempty" csv:"vitamin-a_100g,omitempty"`
	VitaminB12100g                      string `gorm:"column:vitamin-b12_100g;type:varchar(64)" json:"vitamin-b12_100g,omitempty" csv:"vitamin-b12_100g,omitempty"`
	VitaminB1100g                       string `gorm:"column:vitamin-b1_100g;type:varchar(64)" json:"vitamin-b1_100g,omitempty" csv:"vitamin-b1_100g,omitempty"`
	VitaminB2100g                       string `gorm:"column:vitamin-b2_100g;type:varchar(64)" json:"vitamin-b2_100g,omitempty" csv:"vitamin-b2_100g,omitempty"`
	VitaminB6100g                       string `gorm:"column:vitamin-b6_100g;type:varchar(64)" json:"vitamin-b6_100g,omitempty" csv:"vitamin-b6_100g,omitempty"`
	VitaminB9100g                       string `gorm:"column:vitamin-b9_100g;type:varchar(64)" json:"vitamin-b9_100g,omitempty" csv:"vitamin-b9_100g,omitempty"`
	VitaminC100g                        string `gorm:"column:vitamin-c_100g;type:varchar(64)" json:"vitamin-c_100g,omitempty" csv:"vitamin-c_100g,omitempty"`
	VitaminD100g                        string `gorm:"column:vitamin-d_100g;type:varchar(64)" json:"vitamin-d_100g,omitempty" csv:"vitamin-d_100g,omitempty"`
	VitaminE100g                        string `gorm:"column:vitamin-e_100g;type:varchar(64)" json:"vitamin-e_100g,omitempty" csv:"vitamin-e_100g,omitempty"`
	VitaminK100g                        string `gorm:"column:vitamin-k_100g;type:varchar(64)" json:"vitamin-k_100g,omitempty" csv:"vitamin-k_100g,omitempty"`
	VitaminPp100g                       string `gorm:"column:vitamin-pp_100g;type:varchar(64)" json:"vitamin-pp_100g,omitempty" csv:"vitamin-pp_100g,omitempty"`
	WaterHardness100g                   string `gorm:"column:water-hardness_100g;type:varchar(64)" json:"water-hardness_100g,omitempty" csv:"water-hardness_100g,omitempty"`
	Zinc100g                            string `gorm:"column:zinc_100g;type:varchar(64)"json:"zinc_100g,omitempty" csv:"zinc_100g,omitempty"`
}

type OpenFoodFact struct {
	gorm.Model
	Additives                           string         `json:"additives,omitempty" csv:"additives,omitempty"`
	AdditivesN                          string         `json:"additives_n,omitempty" csv:"additives_n,omitempty"`
	AdditivesTags                       string         `json:"additives_tags,omitempty" csv:"additives_tags,omitempty"`
	Allergens                           string         `json:"allergens,omitempty" csv:"allergens,omitempty"`
	AllergensEn                         string         `json:"allergens_en,omitempty" csv:"allergens_en,omitempty"`
	AbbreviatedProductName              string         `json:"abbreviated_product_name,omitempty" csv:"abbreviated_product_name,omitempty"`
	AdditivesEn                         string         `json:"additives_en,omitempty" csv:"additives_en,omitempty"`
	BrandOwner                          string         `json:"brand_owner,omitempty" csv:"brand_owner,omitempty"`
	Brands                              string         `json:"brands,omitempty" csv:"brands,omitempty"`
	BrandsTags                          string         `json:"brands_tags,omitempty" csv:"brands_tags,omitempty"`
	Categories                          string         `json:"categories,omitempty" csv:"categories,omitempty"`
	CategoriesEn                        string         `json:"categories_en,omitempty" csv:"categories_en,omitempty"`
	CategoriesTags                      string         `json:"categories_tags,omitempty" csv:"categories_tags,omitempty"`
	Cities                              string         `json:"cities,omitempty" csv:"cities,omitempty"`
	CitiesTags                          string         `json:"cities_tags,omitempty" csv:"cities_tags,omitempty"`
	Countries                           string         `json:"countries,omitempty" csv:"countries,omitempty"`
	CountriesEn                         string         `json:"countries_en,omitempty" csv:"countries_en,omitempty"`
	CountriesTags                       string         `json:"countries_tags,omitempty" csv:"countries_tags,omitempty"`
	CreatedDatetime                     datatypes.Date `json:"created_datetime,omitempty" csv:"created_datetime,omitempty"`
	CreatedT                            time.Time      `json:"created_t,omitempty" csv:"created_t,omitempty"`
	Creator                             string         `json:"creator,omitempty" csv:"creator,omitempty"`
	EmbCodes                            string         `json:"emb_codes,omitempty" csv:"emb_codes,omitempty"`
	EmbCodesTags                        string         `json:"emb_codes_tags,omitempty" csv:"emb_codes_tags,omitempty"`
	Code                                string         `gorm:"uniqueIndex:code_idx;" json:"code,omitempty" csv:"code,omitempty"`
	GenericName                         string         `json:"generic_name,omitempty" csv:"generic_name,omitempty"`
	ImageIngredientsSmallURL            string         `json:"image_ingredients_small_url,omitempty" csv:"image_ingredients_small_url,omitempty"`
	ImageIngredientsURL                 string         `json:"image_ingredients_url,omitempty" csv:"image_ingredients_url,omitempty"`
	ImageNutritionSmallURL              string         `json:"image_nutrition_small_url,omitempty" csv:"image_nutrition_small_url,omitempty"`
	ImageNutritionURL                   string         `json:"image_nutrition_url,omitempty" csv:"image_nutrition_url,omitempty"`
	ImageSmallURL                       string         `json:"image_small_url,omitempty" csv:"image_small_url,omitempty"`
	ImageURL                            string         `json:"image_url,omitempty" csv:"image_url,omitempty"`
	IngredientsFromPalmOil              string         `json:"ingredients_from_palm_oil,omitempty" csv:"ingredients_from_palm_oil,omitempty"`
	IngredientsFromPalmOilTags          string         `json:"ingredients_from_palm_oil_tags,omitempty" csv:"ingredients_from_palm_oil_tags,omitempty"`
	IngredientsText                     string         `json:"ingredients_text,omitempty" csv:"ingredients_text,omitempty"`
	IngredientsThatMayBeFromPalmOil     string         `json:"ingredients_that_may_be_from_palm_oil,omitempty" csv:"ingredients_that_may_be_from_palm_oil,omitempty"`
	IngredientsThatMayBeFromPalmOilN    string         `json:"ingredients_that_may_be_from_palm_oil_n,omitempty" csv:"ingredients_that_may_be_from_palm_oil_n,omitempty"`
	IngredientsThatMayBeFromPalmOilTags string         `json:"ingredients_that_may_be_from_palm_oil_tags,omitempty" csv:"ingredients_that_may_be_from_palm_oil_tags,omitempty"`
	IngredientsFromPalmOilN             string         `json:"ingredients_from_palm_oil_n,omitempty" csv:"ingredients_from_palm_oil_n,omitempty"`
	Labels                              string         `json:"labels,omitempty" csv:"labels,omitempty"`
	LabelsEn                            string         `json:"labels_en,omitempty" csv:"labels_en,omitempty"`
	LabelsTags                          string         `json:"labels_tags,omitempty" csv:"labels_tags,omitempty"`
	LastModifiedDatetime                time.Time      `json:"last_modified_datetime,omitempty" csv:"last_modified_datetime,omitempty"`
	LastModifiedT                       time.Time      `json:"last_modified_t,omitempty" csv:"last_modified_t,omitempty"`
	NoNutriments                        string         `json:"no_nutriments,omitempty" csv:"no_nutriments,omitempty"`
	NovaGroup                           string         `json:"nova_group,omitempty" csv:"nova_group,omitempty"`
	MainCategory                        string         `json:"main_category,omitempty" csv:"main_category,omitempty"`
	MainCategoryEn                      string         `json:"main_category_en,omitempty" csv:"main_category_en,omitempty"`
	PurchasePlaces                      string         `json:"purchase_places,omitempty" csv:"purchase_places,omitempty"`
	ManufacturingPlaces                 string         `json:"manufacturing_places,omitempty" csv:"manufacturing_places,omitempty"`
	ManufacturingPlacesTags             string         `json:"manufacturing_places_tags,omitempty" csv:"manufacturing_places_tags,omitempty"`
	NutriscoreGrade                     string         `json:"nutriscore_grade,omitempty" csv:"nutriscore_grade,omitempty"`
	NutriscoreScore                     string         `json:"nutriscore_score,omitempty" csv:"nutriscore_score,omitempty"`
	Origins                             string         `json:"origins,omitempty" csv:"origins,omitempty"`
	OriginsEn                           string         `json:"origins_en,omitempty" csv:"origins_en,omitempty"`
	OriginsTags                         string         `json:"origins_tags,omitempty" csv:"origins_tags,omitempty"`
	Packaging                           string         `json:"packaging,omitempty" csv:"packaging,omitempty"`
	PackagingTags                       string         `json:"packaging_tags,omitempty" csv:"packaging_tags,omitempty"`
	PackagingText                       string         `json:"packaging_text,omitempty" csv:"packaging_text,omitempty"`
	ProductName                         string         `json:"product_name,omitempty" csv:"product_name,omitempty"`
	Quantity                            string         `json:"quantity,omitempty" csv:"quantity,omitempty"`
	ServingQuantity                     string         `json:"serving_quantity,omitempty" csv:"serving_quantity,omitempty"`
	ServingSize                         string         `json:"serving_size,omitempty" csv:"serving_size,omitempty"`
	States                              string         `json:"states,omitempty" csv:"states,omitempty"`
	StatesEn                            string         `json:"states_en,omitempty" csv:"states_en,omitempty"`
	StatesTags                          string         `json:"states_tags,omitempty" csv:"states_tags,omitempty"`
	Stores                              string         `json:"stores,omitempty" csv:"stores,omitempty"`
	URL                                 string         `json:"url,omitempty" csv:"url,omitempty"`
	Traces                              string         `json:"traces,omitempty" csv:"traces,omitempty"`
	TracesEn                            string         `json:"traces_en,omitempty" csv:"traces_en,omitempty"`
	TracesTags                          string         `json:"traces_tags,omitempty" csv:"traces_tags,omitempty"`
	AlphaLinolenicAcid100g              string         `gorm:"column:alpha-linolenic-acid_100g;type:varchar(64)" json:"-alpha-linolenic-acid_100g,omitempty" csv:"-alpha-linolenic-acid_100g,omitempty"`
	ArachidicAcid100g                   string         `gorm:"column:arachidic-acid_100g;type:varchar(64)" json:"-arachidic-acid_100g,omitempty" csv:"-arachidic-acid_100g,omitempty"`
	ArachidonicAcid100g                 string         `gorm:"column:arachidonic-acid_100g;type:varchar(64)" json:"-arachidonic-acid_100g,omitempty" csv:"-arachidonic-acid_100g,omitempty"`
	BehenicAcid100g                     string         `gorm:"column:behenic-acid_100g;type:varchar(64)" json:"-behenic-acid_100g,omitempty" csv:"-behenic-acid_100g,omitempty"`
	ButyricAcid100g                     string         `gorm:"column:butyric-acid_100g;type:varchar(64)" json:"-butyric-acid_100g,omitempty" csv:"-butyric-acid_100g,omitempty"`
	CapricAcid100g                      string         `gorm:"column:capric-acid_100g;type:varchar(64)" json:"-capric-acid_100g,omitempty" csv:"-capric-acid_100g,omitempty"`
	CaproicAcid100g                     string         `gorm:"column:caproic-acid_100g;type:varchar(64)" json:"-caproic-acid_100g,omitempty" csv:"-caproic-acid_100g,omitempty"`
	CaprylicAcid100g                    string         `gorm:"column:caprylic-acid_100g;type:varchar(64)" json:"-caprylic-acid_100g,omitempty" csv:"-caprylic-acid_100g,omitempty"`
	CeroticAcid100g                     string         `gorm:"column:cerotic-acid_100g;type:varchar(64)" json:"-cerotic-acid_100g,omitempty" csv:"-cerotic-acid_100g,omitempty"`
	DihomoGammaLinolenicAcid100g        string         `gorm:"column:dihomo-gamma-linolenic-acid_100g;type:varchar(64)" json:"-dihomo-gamma-linolenic-acid_100g,omitempty" csv:"-dihomo-gamma-linolenic-acid_100g,omitempty"`
	DocosahexaenoicAcid100g             string         `gorm:"column:docosahexaenoic-acid_100g;type:varchar(64)" json:"-docosahexaenoic-acid_100g,omitempty" csv:"-docosahexaenoic-acid_100g,omitempty"`
	EicosapentaenoicAcid100g            string         `gorm:"column:eicosapentaenoic-acid_100g;type:varchar(64)" json:"-eicosapentaenoic-acid_100g,omitempty" csv:"-eicosapentaenoic-acid_100g,omitempty"`
	ElaidicAcid100g                     string         `gorm:"column:elaidic-acid_100g;type:varchar(64)" json:"-elaidic-acid_100g,omitempty" csv:"-elaidic-acid_100g,omitempty"`
	ErucicAcid100g                      string         `gorm:"column:erucic-acid_100g;type:varchar(64)" json:"-erucic-acid_100g,omitempty" csv:"-erucic-acid_100g,omitempty"`
	Fructose100g                        string         `gorm:"column:fructose_100g;type:varchar(64)" json:"-fructose_100g,omitempty" csv:"-fructose_100g,omitempty"`
	GammaLinolenicAcid100g              string         `gorm:"column:gamma-linolenic-acid_100g;type:varchar(64)" json:"-gamma-linolenic-acid_100g,omitempty" csv:"-gamma-linolenic-acid_100g,omitempty"`
	Glucose100g                         string         `gorm:"column:glucose_100g;type:varchar(64)" json:"-glucose_100g,omitempty" csv:"-glucose_100g,omitempty"`
	GondoicAcid100g                     string         `gorm:"column:gondoic-acid_100g;type:varchar(64)" json:"-gondoic-acid_100g,omitempty" csv:"-gondoic-acid_100g,omitempty"`
	InsolubleFiber100g                  string         `gorm:"column:insoluble-fiber_100g;type:varchar(64)" json:"-insoluble-fiber_100g,omitempty" csv:"-insoluble-fiber_100g,omitempty"`
	Lactose100g                         string         `gorm:"column:lactose_100g;type:varchar(64)" json:"-lactose_100g,omitempty" csv:"-lactose_100g,omitempty"`
	LauricAcid100g                      string         `gorm:"column:lauric-acid_100g;type:varchar(64)" json:"-lauric-acid_100g,omitempty" csv:"-lauric-acid_100g,omitempty"`
	LignocericAcid100g                  string         `gorm:"column:lignoceric-acid_100g;type:varchar(64)" json:"-lignoceric-acid_100g,omitempty" csv:"-lignoceric-acid_100g,omitempty"`
	LinoleicAcid100g                    string         `gorm:"column:linoleic-acid_100g;type:varchar(64)" json:"-linoleic-acid_100g,omitempty" csv:"-linoleic-acid_100g,omitempty"`
	Maltodextrins100g                   string         `gorm:"column:maltodextrins_100g;type:varchar(64)" json:"-maltodextrins_100g,omitempty" csv:"-maltodextrins_100g,omitempty"`
	Maltose100g                         string         `gorm:"column:maltose_100g;type:varchar(64)" json:"-maltose_100g,omitempty" csv:"-maltose_100g,omitempty"`
	MeadAcid100g                        string         `gorm:"column:mead-acid_100g;type:varchar(64)" json:"-mead-acid_100g,omitempty" csv:"-mead-acid_100g,omitempty"`
	MelissicAcid100g                    string         `gorm:"column:melissic-acid_100g;type:varchar(64)" json:"-melissic-acid_100g,omitempty" csv:"-melissic-acid_100g,omitempty"`
	MontanicAcid100g                    string         `gorm:"column:montanic-acid_100g;type:varchar(64)" json:"-montanic-acid_100g,omitempty" csv:"-montanic-acid_100g,omitempty"`
	MyristicAcid100g                    string         `gorm:"column:myristic-acid_100g;type:varchar(64)" json:"-myristic-acid_100g,omitempty" csv:"-myristic-acid_100g,omitempty"`
	PalmiticAcid100g                    string         `gorm:"column:palmitic-acid_100g;type:varchar(64)" json:"-palmitic-acid_100g,omitempty" csv:"-palmitic-acid_100g,omitempty"`
	NervonicAcid100g                    string         `gorm:"column:nervonic-acid_100g;type:varchar(64)" json:"-nervonic-acid_100g,omitempty" csv:"-nervonic-acid_100g,omitempty"`
	OleicAcid100g                       string         `gorm:"column:oleic-acid_100g;type:varchar(64)" json:"-oleic-acid_100g,omitempty" csv:"-oleic-acid_100g,omitempty"`
	SolubleFiber100g                    string         `gorm:"column:soluble-fiber_100g;type:varchar(64)" json:"-soluble-fiber_100g,omitempty" csv:"-soluble-fiber_100g,omitempty"`
	StearicAcid100g                     string         `gorm:"column:stearic-acid_100g;type:varchar(64)" json:"-stearic-acid_100g,omitempty" csv:"-stearic-acid_100g,omitempty"`
	Sucrose100g                         string         `gorm:"column:sucrose_100g;type:varchar(64)" json:"-sucrose_100g,omitempty" csv:"-sucrose_100g,omitempty"`
	Alcohol100g                         string         `gorm:"column:alcohol_100g;type:varchar(64)" json:"alcohol_100g,omitempty" csv:"alcohol_100g,omitempty"`
	BetaCarotene100g                    string         `gorm:"column:beta-carotene_100g;type:varchar(64)" json:"beta-carotene_100g,omitempty" csv:"beta-carotene_100g,omitempty"`
	BetaGlucan100g                      string         `gorm:"column:beta-glucan_100g;type:varchar(64)" json:"beta-glucan_100g,omitempty" csv:"beta-glucan_100g,omitempty"`
	Bicarbonate100g                     string         `gorm:"column:bicarbonate_100g;type:varchar(64)" json:"bicarbonate_100g,omitempty" csv:"bicarbonate_100g,omitempty"`
	Biotin100g                          string         `gorm:"column:biotin_100g;type:varchar(64)" json:"biotin_100g,omitempty" csv:"biotin_100g,omitempty"`
	Caffeine100g                        string         `gorm:"column:caffeine_100g;type:varchar(64)" json:"caffeine_100g,omitempty" csv:"caffeine_100g,omitempty"`
	Calcium100g                         string         `gorm:"column:calcium_100g;type:varchar(64)" json:"calcium_100g,omitempty" csv:"calcium_100g,omitempty"`
	Carbohydrates100g                   string         `gorm:"column:carbohydrates_100g;type:varchar(64)" json:"carbohydrates_100g,omitempty" csv:"carbohydrates_100g,omitempty"`
	CarbonFootprintFromMeatOrFish100g   string         `gorm:"column:carbon-footprint-from-meat-or-fish_100g;type:varchar(64)" json:"carbon-footprint-from-meat-or-fish_100g,omitempty" csv:"carbon-footprint-from-meat-or-fish_100g,omitempty"`
	CarbonFootprint100g                 string         `gorm:"column:carbon-footprint_100g;type:varchar(64)" json:"carbon-footprint_100g,omitempty" csv:"carbon-footprint_100g,omitempty"`
	Carnitine100g                       string         `gorm:"column:carnitine_100g;type:varchar(64)" json:"carnitine_100g,omitempty" csv:"carnitine_100g,omitempty"`
	Casein100g                          string         `gorm:"column:casein_100g;type:varchar(64)" json:"casein_100g,omitempty" csv:"casein_100g,omitempty"`
	Chloride100g                        string         `gorm:"column:chloride_100g;type:varchar(64)" json:"chloride_100g,omitempty" csv:"chloride_100g,omitempty"`
	Chlorophyl100g                      string         `gorm:"column:chlorophyl_100g;type:varchar(64)" json:"chlorophyl_100g,omitempty" csv:"chlorophyl_100g,omitempty"`
	Cholesterol100g                     string         `gorm:"column:cholesterol_100g;type:varchar(64)" json:"cholesterol_100g,omitempty" csv:"cholesterol_100g,omitempty"`
	Choline100g                         string         `gorm:"column:choline_100g;type:varchar(64)" json:"choline_100g,omitempty" csv:"choline_100g,omitempty"`
	Chromium100g                        string         `gorm:"column:chromium_100g;type:varchar(64)" json:"chromium_100g,omitempty" csv:"chromium_100g,omitempty"`
	Cocoa100g                           string         `gorm:"column:cocoa_100g;type:varchar(64)" json:"cocoa_100g,omitempty" csv:"cocoa_100g,omitempty"`
	CollagenMeatProteinRatio100g        string         `gorm:"column:collagen-meat-protein-ratio_100g;type:varchar(64)" json:"collagen-meat-protein-ratio_100g,omitempty" csv:"collagen-meat-protein-ratio_100g,omitempty"`
	Copper100g                          string         `gorm:"column:copper_100g;type:varchar(64)" json:"copper_100g,omitempty" csv:"copper_100g,omitempty"`
	EnergyFromFat100g                   string         `gorm:"column:energy-from-fat_100g;type:varchar(64)" json:"energy-from-fat_100g,omitempty" csv:"energy-from-fat_100g,omitempty"`
	EnergyKcal100g                      string         `gorm:"column:energy-kcal_100g;type:varchar(64)" json:"energy-kcal_100g,omitempty" csv:"energy-kcal_100g,omitempty"`
	EnergyKj100g                        string         `gorm:"column:energy-kj_100g;type:varchar(64)" json:"energy-kj_100g,omitempty" csv:"energy-kj_100g,omitempty"`
	Energy100g                          string         `gorm:"column:energy_100g;type:varchar(64)" json:"energy_100g,omitempty" csv:"energy_100g,omitempty"`
	Fat100g                             string         `gorm:"column:fat_100g;type:varchar(64)" json:"fat_100g,omitempty" csv:"fat_100g,omitempty"`
	Fiber100g                           string         `gorm:"column:fiber_100g;type:varchar(64)" json:"fiber_100g,omitempty" csv:"fiber_100g,omitempty"`
	FirstPackagingCodeGeo               string         `gorm:"column:first_packaging_code_geo;type:varchar(64)" json:"first_packaging_code_geo,omitempty" csv:"first_packaging_code_geo,omitempty"`
	Fluoride100g                        string         `gorm:"column:fluoride_100g;type:varchar(64)" json:"fluoride_100g,omitempty" csv:"fluoride_100g,omitempty"`
	Folates100g                         string         `gorm:"column:folates_100g;type:varchar(64)" json:"folates_100g,omitempty" csv:"folates_100g,omitempty"`
	FruitsVegetablesNutsDried100g       string         `gorm:"column:fruits-vegetables-nuts-dried_100g;type:varchar(64)" json:"fruits-vegetables-nuts-dried_100g,omitempty" csv:"fruits-vegetables-nuts-dried_100g,omitempty"`
	FruitsVegetablesNutsEstimate100g    string         `gorm:"column:fruits-vegetables-nuts-estimate_100g;type:varchar(64)" json:"fruits-vegetables-nuts-estimate_100g,omitempty" csv:"fruits-vegetables-nuts-estimate_100g,omitempty"`
	FruitsVegetablesNuts100g            string         `gorm:"column:fruits-vegetables-nuts_100g;type:varchar(64)" json:"fruits-vegetables-nuts_100g,omitempty" csv:"fruits-vegetables-nuts_100g,omitempty"`
	GlycemicIndex100g                   string         `gorm:"column:glycemic-index_100g;type:varchar(64)" json:"glycemic-index_100g,omitempty" csv:"glycemic-index_100g,omitempty"`
	Inositol100g                        string         `gorm:"column:inositol_100g;type:varchar(64)" json:"inositol_100g,omitempty" csv:"inositol_100g,omitempty"`
	Iodine100g                          string         `gorm:"column:iodine_100g;type:varchar(64)" json:"iodine_100g,omitempty" csv:"iodine_100g,omitempty"`
	Iron100g                            string         `gorm:"column:iron_100g;type:varchar(64)" json:"iron_100g,omitempty" csv:"iron_100g,omitempty"`
	Magnesium100g                       string         `gorm:"column:magnesium_100g;type:varchar(64)" json:"magnesium_100g,omitempty" csv:"magnesium_100g,omitempty"`
	Manganese100g                       string         `gorm:"column:manganese_100g;type:varchar(64)" json:"manganese_100g,omitempty" csv:"manganese_100g,omitempty"`
	Molybdenum100g                      string         `gorm:"column:molybdenum_100g;type:varchar(64)" json:"molybdenum_100g,omitempty" csv:"molybdenum_100g,omitempty"`
	MonounsaturatedFat100g              string         `gorm:"column:monounsaturated-fat_100g;type:varchar(64)" json:"monounsaturated-fat_100g,omitempty" csv:"monounsaturated-fat_100g,omitempty"`
	Nucleotides100g                     string         `gorm:"column:nucleotides_100g;type:varchar(64)" json:"nucleotides_100g,omitempty" csv:"nucleotides_100g,omitempty"`
	NutritionScoreFr100g                string         `gorm:"column:nutrition-score-fr_100g;type:varchar(64)" json:"nutrition-score-fr_100g,omitempty" csv:"nutrition-score-fr_100g,omitempty"`
	NutritionScoreUk100g                string         `gorm:"column:nutrition-score-uk_100g;type:varchar(64)" json:"nutrition-score-uk_100g,omitempty" csv:"nutrition-score-uk_100g,omitempty"`
	Omega3Fat100g                       string         `gorm:"column:omega-3-fat_100g;type:varchar(64)" json:"omega-3-fat_100g,omitempty" csv:"omega-3-fat_100g,omitempty"`
	Omega6Fat100g                       string         `gorm:"column:omega-6-fat_100g;type:varchar(64)" json:"omega-6-fat_100g,omitempty" csv:"omega-6-fat_100g,omitempty"`
	Omega9Fat100g                       string         `gorm:"column:omega-9-fat_100g;type:varchar(64)" json:"omega-9-fat_100g,omitempty" csv:"omega-9-fat_100g,omitempty"`
	PantothenicAcid100g                 string         `gorm:"column:pantothenic-acid_100g;type:varchar(64)" json:"pantothenic-acid_100g,omitempty" csv:"pantothenic-acid_100g,omitempty"`
	Ph100g                              string         `gorm:"column:ph_100g;type:varchar(64)" json:"ph_100g,omitempty" csv:"ph_100g,omitempty"`
	Phosphorus100g                      string         `gorm:"column:phosphorus_100g;type:varchar(64)" json:"phosphorus_100g,omitempty" csv:"phosphorus_100g,omitempty"`
	Phylloquinone100g                   string         `gorm:"column:phylloquinone_100g;type:varchar(64)" json:"phylloquinone_100g,omitempty" csv:"phylloquinone_100g,omitempty"`
	PnnsGroups1                         string         `gorm:"column:pnns_groups_1;type:varchar(64)" json:"pnns_groups_1,omitempty" csv:"pnns_groups_1,omitempty"`
	PnnsGroups2                         string         `gorm:"column:pnns_groups_2;type:varchar(64)" json:"pnns_groups_2,omitempty" csv:"pnns_groups_2,omitempty"`
	Polyols100g                         string         `gorm:"column:polyols_100g;type:varchar(64)" json:"polyols_100g,omitempty" csv:"polyols_100g,omitempty"`
	PolyunsaturatedFat100g              string         `gorm:"column:polyunsaturated-fat_100g;type:varchar(64)" json:"polyunsaturated-fat_100g,omitempty" csv:"polyunsaturated-fat_100g,omitempty"`
	Potassium100g                       string         `gorm:"column:potassium_100g;type:varchar(64)" json:"potassium_100g,omitempty" csv:"potassium_100g,omitempty"`
	Proteins100g                        string         `gorm:"column:proteins_100g;type:varchar(64)" json:"proteins_100g,omitempty" csv:"proteins_100g,omitempty"`
	Salt100g                            string         `gorm:"column:salt_100g;type:varchar(64)" json:"salt_100g,omitempty" csv:"salt_100g,omitempty"`
	SaturatedFat100g                    string         `gorm:"column:saturated-fat_100g;type:varchar(64)"  json:"saturated-fat_100g,omitempty" csv:"saturated-fat_100g,omitempty"`
	Selenium100g                        string         `gorm:"column:selenium_100g;type:varchar(64)" json:"selenium_100g,omitempty" csv:"selenium_100g,omitempty"`
	SerumProteins100g                   string         `gorm:"column:serum-proteins_100g;type:varchar(64)" json:"serum-proteins_100g,omitempty" csv:"serum-proteins_100g,omitempty"`
	Silica100g                          string         `gorm:"column:silica_100g;type:varchar(64)" json:"silica_100g,omitempty" csv:"silica_100g,omitempty"`
	Sodium100g                          string         `gorm:"column:sodium_100g;type:varchar(64)" json:"sodium_100g,omitempty" csv:"sodium_100g,omitempty"`
	Starch100g                          string         `gorm:"column:starch_100g;type:varchar(64)" json:"starch_100g,omitempty" csv:"starch_100g,omitempty"`
	Sugars100g                          string         `gorm:"column:sugars_100g;type:varchar(64)" json:"sugars_100g,omitempty" csv:"sugars_100g,omitempty"`
	Taurine100g                         string         `gorm:"column:taurine_100g;type:varchar(64)" json:"taurine_100g,omitempty" csv:"taurine_100g,omitempty"`
	TransFat100g                        string         `gorm:"column:trans-fat_100g;type:varchar(64)" json:"trans-fat_100g,omitempty" csv:"trans-fat_100g,omitempty"`
	VitaminA100g                        string         `gorm:"column:vitamin-a_100g;type:varchar(64)" json:"vitamin-a_100g,omitempty" csv:"vitamin-a_100g,omitempty"`
	VitaminB12100g                      string         `gorm:"column:vitamin-b12_100g;type:varchar(64)" json:"vitamin-b12_100g,omitempty" csv:"vitamin-b12_100g,omitempty"`
	VitaminB1100g                       string         `gorm:"column:vitamin-b1_100g;type:varchar(64)" json:"vitamin-b1_100g,omitempty" csv:"vitamin-b1_100g,omitempty"`
	VitaminB2100g                       string         `gorm:"column:vitamin-b2_100g;type:varchar(64)" json:"vitamin-b2_100g,omitempty" csv:"vitamin-b2_100g,omitempty"`
	VitaminB6100g                       string         `gorm:"column:vitamin-b6_100g;type:varchar(64)" json:"vitamin-b6_100g,omitempty" csv:"vitamin-b6_100g,omitempty"`
	VitaminB9100g                       string         `gorm:"column:vitamin-b9_100g;type:varchar(64)" json:"vitamin-b9_100g,omitempty" csv:"vitamin-b9_100g,omitempty"`
	VitaminC100g                        string         `gorm:"column:vitamin-c_100g;type:varchar(64)" json:"vitamin-c_100g,omitempty" csv:"vitamin-c_100g,omitempty"`
	VitaminD100g                        string         `gorm:"column:vitamin-d_100g;type:varchar(64)" json:"vitamin-d_100g,omitempty" csv:"vitamin-d_100g,omitempty"`
	VitaminE100g                        string         `gorm:"column:vitamin-e_100g;type:varchar(64)" json:"vitamin-e_100g,omitempty" csv:"vitamin-e_100g,omitempty"`
	VitaminK100g                        string         `gorm:"column:vitamin-k_100g;type:varchar(64)" json:"vitamin-k_100g,omitempty" csv:"vitamin-k_100g,omitempty"`
	VitaminPp100g                       string         `gorm:"column:vitamin-pp_100g;type:varchar(64)" json:"vitamin-pp_100g,omitempty" csv:"vitamin-pp_100g,omitempty"`
	WaterHardness100g                   string         `gorm:"column:water-hardness_100g;type:varchar(64)" json:"water-hardness_100g,omitempty" csv:"water-hardness_100g,omitempty"`
	Zinc100g                            string         `gorm:"column:zinc_100g;type:varchar(64)"json:"zinc_100g,omitempty" csv:"zinc_100g,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (OpenFoodFact) TableName() string {
	return "openfoodfact_tmp"
}
