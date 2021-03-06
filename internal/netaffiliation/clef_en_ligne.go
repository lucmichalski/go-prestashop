package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/clef-en-ligne.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
nom: Tac ENIQ MIFARE,
référence interne: 5051.00518,
prix actuel: 24.00,
catégorie: Accueil > Cylindres de porte > Cylindres électroniques > Cylindres DOM,
URL de la page produit: https://action.metaffiliation.com/trk.php?mclic=P4C989459831FS1UA4128790603V4,
modèle: ,
marque: ,
référence universelle: ,
référence constructeur: ,
frais de ports: 8.00,
descriptif: ,
garantie: ,
indicateur de stock: 1,
disponibilité: ,
indicateur de performance: ,
indicateur promotion: ,
indicateur nouveauté: ,
URL image petite: ,
URL image moyenne: ,
URL image grande: https://www.clef-en-ligne.com/4455-large_default/tac-els-protector.jpg,
écotaxe: 0.00,
Matière: 
\\ ===================================================================================== //

*/

type ClefEnLigne struct {
	Categorie               string `json:"categorie" struct2map:"key:categorie"`
	Descriptif              string `json:"descriptif" struct2map:"key:descriptif"`
	Disponibilite           string `json:"disponibilite" struct2map:"key:disponibilite"`
	Ecotaxe                 string `json:"ecotaxe" struct2map:"key:ecotaxe"`
	FraisDePorts            string `json:"frais_de_ports" struct2map:"key:frais_de_ports"`
	Garantie                string `json:"garantie" struct2map:"key:garantie"`
	IndicateurDePerformance string `json:"indicateur_de_performance" struct2map:"key:indicateur_de_performance"`
	IndicateurDeStock       string `json:"indicateur_de_stock" struct2map:"key:indicateur_de_stock"`
	IndicateurNouveaute     string `json:"indicateur_nouveaute" struct2map:"key:indicateur_nouveaute"`
	IndicateurPromotion     string `json:"indicateur_promotion" struct2map:"key:indicateur_promotion"`
	Marque                  string `json:"marque" struct2map:"key:marque"`
	Matiere                 string `json:"matiere" struct2map:"key:matiere"`
	Modele                  string `json:"modele" struct2map:"key:modele"`
	Nom                     string `json:"nom" struct2map:"key:nom"`
	PrixActuel              string `json:"prix_actuel" struct2map:"key:prix_actuel"`
	ReferenceConstructeur   string `json:"reference_constructeur" struct2map:"key:reference_constructeur"`
	ReferenceInterne        string `json:"reference_interne" struct2map:"key:reference_interne"`
	ReferenceUniverselle    string `json:"reference_universelle" struct2map:"key:reference_universelle"`
	URLDeLaPageProduit      string `json:"url_de_la_page_produit" struct2map:"key:url_de_la_page_produit"`
	URLImageGrande          string `json:"url_image_grande" struct2map:"key:url_image_grande"`
	URLImageMoyenne         string `json:"url_image_moyenne" struct2map:"key:url_image_moyenne"`
	URLImagePetite          string `json:"url_image_petite" struct2map:"key:url_image_petite"`
}

/*
func (o *ClefEnLigne) Convert() ([]string, error) {

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
