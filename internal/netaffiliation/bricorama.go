package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/bricorama.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
sku: 1931106,
product_id: 4,
qty: 5,
status: Enabled,
category: Catalogue Web Fr,
category_url: http://bo.bricorama.fr/scheduler.php/catalog/category/view/s/catalogue-web-fr/id/1019/,
category_sub_1: Outillage & construction,
category_url_sub_1: http://bo.bricorama.fr/scheduler.php/outillage-construction.html,
category_sub_2: Electricité & domotique,
category_url_sub_2: http://bo.bricorama.fr/scheduler.php/outillage-construction/electricite-domotique.html,
category_sub_3: Alarme & vidéosurveillance,
category_url_sub_3: http://bo.bricorama.fr/scheduler.php/outillage-construction/electricite-domotique/alarme-videosurveillance.html,
category_sub_4: ,
category_url_sub_4: ,
category_sub_5: ,
category_url_sub_5: ,
category_breadcrumb: Catalogue Web Fr > Outillage & construction > Electricité & domotique > Alarme & vidéosurveillance,
price_ttc: 372,
price_before_discount: 372,
discount_amount: 0,
discount_percent: 0,
start_date_discount: ,
end_date_discount: ,
shipping_delay: 4,
shipping_name: Chronopost,
shipping_price: 13.02,
image_url_1: https://bo.bricorama.fr/media/catalog/product/5/7/57800c9704be2d2f0181b25cd4f766972cc18180_3432863006856.jpg,
image_url_2: ,
image_url_3: ,
image_url_4: ,
image_url_5: ,
product_url: https://action.metaffiliation.com/trk.php?mclic=P4600545983219S1UD4129cd0277046V4,
name: Sirène flash extérieure DIAGRAL DIAG50AAX,
description: La sirène-flash permet : • de dissuader l'intrus (en modulation intérieure), • d'alerter le voisinage (en modulation extérieure), • de localiser facilement l'alarme (flash clignotant). En cas de tentative d'arrachement, la sirène se déclenche, puis par l'intermédiaire de la centrale, déclenche l'ensemble des moyens d'alarme.La sirène signale à la centrale son défaut de piles.En cas de détection incendie, la sirène se déclenche pour une durée de 5 mn (modulation incendie) Durée de la sirène: 20 à 180s. Programmation directement sur la centrale. Garantie 5 ans.,
short_description: Autonomie 4 ans en usage courant,
parent_id: ,
product_type: simple,
product_variation: ,
image_default: https://bo.bricorama.fr/media/catalog/product/5/7/57800c9704be2d2f0181b25cd4f766972cc18180_3432863006856.jpg,
child_name: Sirène flash extérieure DIAGRAL DIAG50AAX,
weight: 2.6000,
color: ,
gallery: ,
grp_marche: 01173117,
ean: 3432863006856,
vente_sur_stock: 1,
marque: DIAGRAL,
couleur: ,
statut_sap: ,
puissance_utilisation_cont_w: ,
puissance_maximum_allumage_w: ,
diametre_mm: ,
longueur_mm: ,
poids_kg: ,
vitesse_vide: ,
temperature_maximum: ,
autonomie_mn: ,
puissance_w: ,
capacite_aspiration: ,
longueur_flexible_aspiration: ,
hauteur_coupe: ,
largeur_cm: ,
surface_toile: ,
forme_toile: ,
matiere_toile: ,
largeur_toile: ,
longueur_toile: ,
charge_maximum_kg: ,
ecartement_maximum_presse_mm: ,
epaisseur_plateau_mm: ,
hauteur_max_etabli_cm: ,
largeur_plateau_m: ,
longueur_etabli_cm: ,
profondeur_etabli_cm: ,
caisson_rangement: ,
matiere_principale: ,
hauteur_cm: ,
livre_soutils: ,
nombre_rang: ,
roues_pivotante: ,
nombre_compart: ,
diametre_cm: ,
type_peinture: ,
diametre_largeur: ,
diametre_pneu: ,
profondeur_cm: ,
contenance_kg: ,
poids_past: ,
dimension_cm: ,
profondeur_inter: ,
largeur_inter: ,
largeur_ext: ,
montage: ,
dimension: ,
mcid: MC-2809,
shop_list: ,
cat_rdc: ACCESS-61,
rakuten: 
\\ ===================================================================================== //

*/

type Bricorama struct {
	AutonomieMn                string `json:"autonomie_mn" struct2map:"key:autonomie_mn"`
	CaissonRangement           string `json:"caisson_rangement" struct2map:"key:caisson_rangement"`
	CapaciteAspiration         string `json:"capacite_aspiration" struct2map:"key:capacite_aspiration"`
	CatRdc                     string `json:"cat_rdc" struct2map:"key:cat_rdc"`
	Category                   string `json:"category" struct2map:"key:category"`
	CategoryBreadcrumb         string `json:"category_breadcrumb" struct2map:"key:category_breadcrumb"`
	CategorySub1               string `json:"category_sub_1" struct2map:"key:category_sub_1"`
	CategorySub2               string `json:"category_sub_2" struct2map:"key:category_sub_2"`
	CategorySub3               string `json:"category_sub_3" struct2map:"key:category_sub_3"`
	CategorySub4               string `json:"category_sub_4" struct2map:"key:category_sub_4"`
	CategorySub5               string `json:"category_sub_5" struct2map:"key:category_sub_5"`
	CategoryURL                string `json:"category_url" struct2map:"key:category_url"`
	CategoryURLSub1            string `json:"category_url_sub_1" struct2map:"key:category_url_sub_1"`
	CategoryURLSub2            string `json:"category_url_sub_2" struct2map:"key:category_url_sub_2"`
	CategoryURLSub3            string `json:"category_url_sub_3" struct2map:"key:category_url_sub_3"`
	CategoryURLSub4            string `json:"category_url_sub_4" struct2map:"key:category_url_sub_4"`
	CategoryURLSub5            string `json:"category_url_sub_5" struct2map:"key:category_url_sub_5"`
	ChargeMaximumKg            string `json:"charge_maximum_kg" struct2map:"key:charge_maximum_kg"`
	ChildName                  string `json:"child_name" struct2map:"key:child_name"`
	Color                      string `json:"color" struct2map:"key:color"`
	ContenanceKg               string `json:"contenance_kg" struct2map:"key:contenance_kg"`
	Couleur                    string `json:"couleur" struct2map:"key:couleur"`
	Description                string `json:"description" struct2map:"key:description"`
	DiametreCm                 string `json:"diametre_cm" struct2map:"key:diametre_cm"`
	DiametreLargeur            string `json:"diametre_largeur" struct2map:"key:diametre_largeur"`
	DiametreMm                 string `json:"diametre_mm" struct2map:"key:diametre_mm"`
	DiametrePneu               string `json:"diametre_pneu" struct2map:"key:diametre_pneu"`
	Dimension                  string `json:"dimension" struct2map:"key:dimension"`
	DimensionCm                string `json:"dimension_cm" struct2map:"key:dimension_cm"`
	DiscountAmount             string `json:"discount_amount" struct2map:"key:discount_amount"`
	DiscountPercent            string `json:"discount_percent" struct2map:"key:discount_percent"`
	Ean                        string `json:"ean" struct2map:"key:ean"`
	EcartementMaximumPresseMm  string `json:"ecartement_maximum_presse_mm" struct2map:"key:ecartement_maximum_presse_mm"`
	EndDateDiscount            string `json:"end_date_discount" struct2map:"key:end_date_discount"`
	EpaisseurPlateauMm         string `json:"epaisseur_plateau_mm" struct2map:"key:epaisseur_plateau_mm"`
	FormeToile                 string `json:"forme_toile" struct2map:"key:forme_toile"`
	Gallery                    string `json:"gallery" struct2map:"key:gallery"`
	GrpMarche                  string `json:"grp_marche" struct2map:"key:grp_marche"`
	HauteurCm                  string `json:"hauteur_cm" struct2map:"key:hauteur_cm"`
	HauteurCoupe               string `json:"hauteur_coupe" struct2map:"key:hauteur_coupe"`
	HauteurMaxEtabliCm         string `json:"hauteur_max_etabli_cm" struct2map:"key:hauteur_max_etabli_cm"`
	ImageDefault               string `json:"image_default" struct2map:"key:image_default"`
	ImageURL1                  string `json:"image_url_1" struct2map:"key:image_url_1"`
	ImageURL2                  string `json:"image_url_2" struct2map:"key:image_url_2"`
	ImageURL3                  string `json:"image_url_3" struct2map:"key:image_url_3"`
	ImageURL4                  string `json:"image_url_4" struct2map:"key:image_url_4"`
	ImageURL5                  string `json:"image_url_5" struct2map:"key:image_url_5"`
	LargeurCm                  string `json:"largeur_cm" struct2map:"key:largeur_cm"`
	LargeurExt                 string `json:"largeur_ext" struct2map:"key:largeur_ext"`
	LargeurInter               string `json:"largeur_inter" struct2map:"key:largeur_inter"`
	LargeurPlateaum            string `json:"largeur_plateau_m" struct2map:"key:largeur_plateau_m"`
	LargeurToile               string `json:"largeur_toile" struct2map:"key:largeur_toile"`
	LivreSoutils               string `json:"livre_soutils" struct2map:"key:livre_soutils"`
	LongueurEtabliCm           string `json:"longueur_etabli_cm" struct2map:"key:longueur_etabli_cm"`
	LongueurFlexibleAspiration string `json:"longueur_flexible_aspiration" struct2map:"key:longueur_flexible_aspiration"`
	LongueurMm                 string `json:"longueur_mm" struct2map:"key:longueur_mm"`
	LongueurToile              string `json:"longueur_toile" struct2map:"key:longueur_toile"`
	Marque                     string `json:"marque" struct2map:"key:marque"`
	MatierePrincipale          string `json:"matiere_principale" struct2map:"key:matiere_principale"`
	MatiereToile               string `json:"matiere_toile" struct2map:"key:matiere_toile"`
	Mcid                       string `json:"mcid" struct2map:"key:mcid"`
	Montage                    string `json:"montage" struct2map:"key:montage"`
	Name                       string `json:"name" struct2map:"key:name"`
	NombreCompart              string `json:"nombre_compart" struct2map:"key:nombre_compart"`
	NombreRang                 string `json:"nombre_rang" struct2map:"key:nombre_rang"`
	ParentID                   string `json:"parent_id" struct2map:"key:parent_id"`
	PoidsKg                    string `json:"poids_kg" struct2map:"key:poids_kg"`
	PoidsPast                  string `json:"poids_past" struct2map:"key:poids_past"`
	PriceBeforeDiscount        string `json:"price_before_discount" struct2map:"key:price_before_discount"`
	PriceTtc                   string `json:"price_ttc" struct2map:"key:price_ttc"`
	ProductID                  string `json:"product_id" struct2map:"key:product_id"`
	ProductType                string `json:"product_type" struct2map:"key:product_type"`
	ProductURL                 string `json:"product_url" struct2map:"key:product_url"`
	ProductVariation           string `json:"product_variation" struct2map:"key:product_variation"`
	ProfondeurCm               string `json:"profondeur_cm" struct2map:"key:profondeur_cm"`
	ProfondeurEtabliCm         string `json:"profondeur_etabli_cm" struct2map:"key:profondeur_etabli_cm"`
	ProfondeurInter            string `json:"profondeur_inter" struct2map:"key:profondeur_inter"`
	PuissanceMaximumAllumagew  string `json:"puissance_maximum_allumage_w" struct2map:"key:puissance_maximum_allumage_w"`
	PuissanceUtilisationContw  string `json:"puissance_utilisation_cont_w" struct2map:"key:puissance_utilisation_cont_w"`
	Puissancew                 string `json:"puissance_w" struct2map:"key:puissance_w"`
	Qty                        string `json:"qty" struct2map:"key:qty"`
	Rakuten                    string `json:"rakuten" struct2map:"key:rakuten"`
	RouesPivotante             string `json:"roues_pivotante" struct2map:"key:roues_pivotante"`
	ShippingDelay              string `json:"shipping_delay" struct2map:"key:shipping_delay"`
	ShippingName               string `json:"shipping_name" struct2map:"key:shipping_name"`
	ShippingPrice              string `json:"shipping_price" struct2map:"key:shipping_price"`
	ShopList                   string `json:"shop_list" struct2map:"key:shop_list"`
	ShortDescription           string `json:"short_description" struct2map:"key:short_description"`
	Sku                        string `json:"sku" struct2map:"key:sku"`
	StartDateDiscount          string `json:"start_date_discount" struct2map:"key:start_date_discount"`
	Status                     string `json:"status" struct2map:"key:status"`
	StatutSap                  string `json:"statut_sap" struct2map:"key:statut_sap"`
	SurfaceToile               string `json:"surface_toile" struct2map:"key:surface_toile"`
	TemperatureMaximum         string `json:"temperature_maximum" struct2map:"key:temperature_maximum"`
	TypePeinture               string `json:"type_peinture" struct2map:"key:type_peinture"`
	VenteSurStock              string `json:"vente_sur_stock" struct2map:"key:vente_sur_stock"`
	VitesseVide                string `json:"vitesse_vide" struct2map:"key:vitesse_vide"`
	Weight                     string `json:"weight" struct2map:"key:weight"`
}

/*
func (o *Bricorama) Convert() ([]string, error) {

	mp := make(map[string]string, 0) 

	mp["product_name"] = o.ProductName
	mp["reference"] = o.Reference

	categoryId, err := findOrCreateCategoryByNam(db, dbTablePrefix, o.Category)
	if err != nil {
		return nil, err
	}
	mp["category_id"] = fmt.Sprintf("%d", categoryId)

	defaultCategory, err := findOrCreateCategoryByName(db, dbTablePrefix, o.DefaultCategory)
	if err != nil {
		return nil, err
	}
	mp["default_category"] = fmt.Sprintf("%d", defaultCategory)

	mp["short_description"] = o.ShortDescription
	mp["description"] = o.Description
	mp["price"] = fo.Price
	mp["quantity"] = o.Quantity
	mp["image_ref"] = o.ImageRef

	// go ast ?
	// m := structs.Name(o)
	// pp.Println(m)
	// for _, feature := range o.Features {}
	mp["feature_name"] = o.FeatureNames
	mp["feature_values"] = o.FeatureValues

	row =  []string{
		mp["product_name"], 
		mp["reference"], 
		mp["category_id"],
		mp["default_category"],
		mp["short_description"],
		mp["description"],
		mp["price"],
		mp["quantity"],
		mp["image_ref"],
		mp["feature_name"],
		mp["feature_values"],
	}

	return row, nil
}
*/
