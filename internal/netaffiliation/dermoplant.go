package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/dermoplant.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
EAN: 3760113220005,
ID Wordpress: 25,
Nom produit: Crème Lifting Céleste,
Référence interne: CEL0002,
Prix actuel TTC: 32,39€,
Prix barré TTC: 32,39€,
Catégorie du produit: Céleste > Crème Lifting Céleste,
URL: https://wmb.dermoplant.com/?P51158745983211S1U941322d013V4,
URL de l'image: https://dermoplant.com/wp-content/uploads/2013/09/Creme-lifting-celeste.jpg,
Code HAN: ,
Nom de Marque: DERMOPLANT,
Descriptif du produit: Pour une peau hydratée, protégé, lifté et rajeunie. Les principes actifs de sa composition renforcent la qualité du collagène et de l'élastine contenu dans le derme. La teneur équilibrée en huile de germe de blé, en huile d'amande et en huile d'abricot remplace le sébum produit naturellement par la peau, pour un effet émollient: hydrate et protège contre les radicaux libres.,
Montant des FDP: 5,67€,
Montant Eco-taxe: ,
Garantie de base: 0,
Indicateur de stock: >1000,
Indicateur de performance: 19,
Indicateur de promotion: 2,
Indicateur de nouveauté: 0
\\ ===================================================================================== //

*/

type Dermoplant struct {
	CategorieDuProduit      string `json:"categorie_du_produit" struct2map:"key:categorie_du_produit"`
	CodeHan                 string `json:"code_han" struct2map:"key:code_han"`
	DescriptifDuProduit     string `json:"descriptif_du_produit" struct2map:"key:descriptif_du_produit"`
	Ean                     string `json:"ean" struct2map:"key:ean"`
	GarantieDeBase          string `json:"garantie_de_base" struct2map:"key:garantie_de_base"`
	IDWordpress             string `json:"id_wordpress" struct2map:"key:id_wordpress"`
	IndicateurDeNouveaute   string `json:"indicateur_de_nouveaute" struct2map:"key:indicateur_de_nouveaute"`
	IndicateurDePerformance string `json:"indicateur_de_performance" struct2map:"key:indicateur_de_performance"`
	IndicateurDePromotion   string `json:"indicateur_de_promotion" struct2map:"key:indicateur_de_promotion"`
	IndicateurDeStock       string `json:"indicateur_de_stock" struct2map:"key:indicateur_de_stock"`
	MontantDesFdp           string `json:"montant_des_fdp" struct2map:"key:montant_des_fdp"`
	MontantEcoTaxe          string `json:"montant_eco_taxe" struct2map:"key:montant_eco_taxe"`
	NomDeMarque             string `json:"nom_de_marque" struct2map:"key:nom_de_marque"`
	NomProduit              string `json:"nom_produit" struct2map:"key:nom_produit"`
	PrixActuelTtc           string `json:"prix_actuel_ttc" struct2map:"key:prix_actuel_ttc"`
	PrixBarreTtc            string `json:"prix_barre_ttc" struct2map:"key:prix_barre_ttc"`
	ReferenceInterne        string `json:"reference_interne" struct2map:"key:reference_interne"`
	URL                     string `json:"url" struct2map:"key:url"`
	URLDelImage             string `json:"url_de_l'image" struct2map:"key:url_de_l'image"`
}

/*
func (o *Dermoplant) Convert() ([]string, error) {

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
