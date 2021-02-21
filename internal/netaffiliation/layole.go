package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/layole.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
EAN13: 3760037752280,
Nom: Sommelier Laguiole manche galbé, bois de pistachier, 2 mitres inox mat,
Référence du produit: 2SO312IM2AFTINPIS0,
Prix de vente: 199.60,
Prix sans réduction: 199.60,
Noms de toutes les catégories: Couverts de table Laguiole>Sommeliers Laguiole>Sommeliers Laguiole>Manches galbés>Manches galbés>Sommeliers Laguiole>Sommeliers manches galbés (matières naturelles),
URL du produit: https://action.metaffiliation.com/trk.php?mclic=P4BFE945983219S1UB4129bf02764V4,
URL de l'image par défaut: https://www.layole.com/50966/sommelier-laguiole-manche-galbe-bois-de-pistachier-2-mitres-inox-mat.jpg,
Fabricant: ,
Description: <p>Sommelier avec les 3 pièces :</p>  <p>-mèche de tire-bouchon de type professionnelle</p>  <p>-lame avec microdents</p>  <p>-levier avec décapsuleur</p>  <p>Il comporte un <strong>manche galbé avec la croix du berger,</strong> le ressort est <strong>ciselé</strong> et il vous est expédié dans une boîte en bois.</p>
\\ ===================================================================================== //

*/

type Layole struct {
	Description               string `json:"description" struct2map:"key:description"`
	Ean13                     string `json:"ean_13" struct2map:"key:ean_13"`
	Fabricant                 string `json:"fabricant" struct2map:"key:fabricant"`
	Nom                       string `json:"nom" struct2map:"key:nom"`
	NomsDeToutesLesCategories string `json:"noms_de_toutes_les_categories" struct2map:"key:noms_de_toutes_les_categories"`
	PrixDeVente               string `json:"prix_de_vente" struct2map:"key:prix_de_vente"`
	PrixSansReduction         string `json:"prix_sans_reduction" struct2map:"key:prix_sans_reduction"`
	ReferenceDuProduit        string `json:"reference_du_produit" struct2map:"key:reference_du_produit"`
	URLDelImageParDefaut      string `json:"url_de_l'image_par_defaut" struct2map:"key:url_de_l'image_par_defaut"`
	URLDuProduit              string `json:"url_du_produit" struct2map:"key:url_du_produit"`
}

/*
func (o *Layole) Convert() ([]string, error) {

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
