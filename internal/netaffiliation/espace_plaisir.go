package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/espace-plaisir.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
EAN: 4032498950013,
Nom: Coupes Menstruelles Fun Cup Taille A,
ID: 9349,
Prix TTC: 34.9,
Prix TTC Final: 34.9,
Categorie: Accueil > Bien-être,
URL: https://action.metaffiliation.com/trk.php?mclic=P4AD4B459831FS1UC412a33015202V4,
Image: https://www.espaceplaisir.fr/img/p/9349-84387.jpg,
REF: 9349-0,
Marque: Fun Factory,
Description: Adoptez les Coupes Menstruelles Fun Cup Taille A de la marque Fun Factory, pour un confort optimal et une hygiène irréprochable pendant vos règles !   La coupe menstruelle constitue une innovation récente qui vient révolutionner le domaine de la protection intime. En effet, elle vous garantit une meilleure hygiène que les tampons et les serviettes hygiéniques et sa dimension pratique s'adaptera à toutes les activités de votre vie quotidienne. Vous pourrez les utiliser en faisant du sport sans aucune gêne et même à la piscine grâce à son étanchéité qui bloquera toute fuite.   Contrairement aux tampons qui peuvent provoquer une sécheresse vaginale ou laissent parfois des résidus de coton pouvant provoquer gênes ou allergies, les Coupes Menstruelles Fun Cup ont été conçues pour respecter votre corps. Elles s,
Transport: 2.90,
Stock: 1
\\ ===================================================================================== //

*/

type EspacePlaisir struct {
	Categorie    string `json:"categorie" struct2map:"key:categorie"`
	Description  string `json:"description" struct2map:"key:description"`
	Ean          string `json:"ean" struct2map:"key:ean"`
	ID           string `json:"id" struct2map:"key:id"`
	Image        string `json:"image" struct2map:"key:image"`
	Marque       string `json:"marque" struct2map:"key:marque"`
	Nom          string `json:"nom" struct2map:"key:nom"`
	PrixTtc      string `json:"prix_ttc" struct2map:"key:prix_ttc"`
	PrixTtcFinal string `json:"prix_ttc_final" struct2map:"key:prix_ttc_final"`
	Ref          string `json:"ref" struct2map:"key:ref"`
	Stock        string `json:"stock" struct2map:"key:stock"`
	Transport    string `json:"transport" struct2map:"key:transport"`
	URL          string `json:"url" struct2map:"key:url"`
}

/*
func (o *EspacePlaisir) Convert() ([]string, error) {

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
