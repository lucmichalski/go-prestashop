package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/ripauste.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
Nom: Carte Cadeau 100,
R_f_rence Interne: 76,
 Montant (TTC) :  100,00   ,
 Montant (TTC) :  100,00   ,
Cat_gorie Fil d'ariane: Accueil>E-BOUTIQUE>CARTES CADEAUX,
URL de la page: https://exp.ripauste.fr/?P5114134598315S1UA4131c50170V4,
URL Image: https://ripauste.fr/1678/carte-cadeau-100.jpg,
R_f_rence fabricant: RIP76,
Nom de la marque: RIPAUSTE,
Description: Carte Cadeau d'un valeur de 100Â Vous cherchez une ide cadeau pour Elle, et vous savez qu'une Pochette en Cuir Ripauste, une ceinture, un porte-monnaie... Elle Va Adorer...Mais vous hsitez sur la couleur, le modle, la taille...N'hsitez plus et Offrez lui une Carte Cadeau d'une valeur de 100Â On s'occupe de tout, vous recevez la Carte Cadeau par email ou par courrier (  prciser dans les remarques lors de la confirmation de commande), Elle n'aura plus qu' choisir son accessoire de mode prfr parmi toute la collection RipausteLa Carte Cadeau est valable 6 mois.
\\ ===================================================================================== //

*/

type Ripauste struct {
	CatGorieFildAriane string `json:"cat_gorie_fil_d'ariane" struct2map:"key:cat_gorie_fil_d'ariane"`
	Description        string `json:"description" struct2map:"key:description"`
	MontantTtc         string `json:"montant_ttc" struct2map:"key:montant_ttc"`
	Nom                string `json:"nom" struct2map:"key:nom"`
	NomDeLaMarque      string `json:"nom_de_la_marque" struct2map:"key:nom_de_la_marque"`
	RfRenceFabricant   string `json:"r_f_rence_fabricant" struct2map:"key:r_f_rence_fabricant"`
	RfRenceInterne     string `json:"r_f_rence_interne" struct2map:"key:r_f_rence_interne"`
	URLDeLaPage        string `json:"url_de_la_page" struct2map:"key:url_de_la_page"`
	URLImage           string `json:"url_image" struct2map:"key:url_image"`
}

/*
func (o *Ripauste) Convert() ([]string, error) {

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
