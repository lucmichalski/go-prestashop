package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/chemin-de-table.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
nom: Timbale Ercuis en métal argenté, modèle Mascottes,
reference: ERC8310,
prix: 190.0000,
categorie: Cadeau bapteme,
url: https://action.metaffiliation.com/trk.php?mclic=P442D34598319S1U841291501V4,
marque: Ercuis,
prix_livraison: 6.90,
descriptif: <p style="text-align: justify;">Grand succ&egrave;s pour ce mod&egrave;le de <strong>Timbale</strong> originale et moderne!!! Vous pouvez faire&nbsp;graver le pr&eacute;nom et la date&nbsp;de bapt&ecirc;me de l'enfant en dessous du motif, le co&ucirc;t de la gravure comprend la gravure du pr&eacute;nom et de la date.&nbsp;<strong>Timbale Ercuis</strong>&nbsp;en m&eacute;tal-argent&eacute;, de hauteur: 7,5 cm.&nbsp;Cette <strong>timbale </strong>est livr&eacute;e dans sa bo&icirc;te verte <strong>Ercuis </strong>et avec son emballage cadeau<strong> Ercuis</strong>.&nbsp;<br /><br /></p>,
en_stock: In Stock,
grande_image: http://www.chemindetable.fr/media/catalog/product/e/r/ercuis-timbale-mascottes-metal-argente.jpg,
moyenne_image: http://www.chemindetable.fr/media/catalog/product/e/r/ercuis-timbale-mascottes-metal-argente.jpg,
petite_image: http://www.chemindetable.fr/media/catalog/product/e/r/ercuis-timbale-mascottes-metal-argente.jpg
\\ ===================================================================================== //

*/

type CheminDeTable struct {
	Categorie     string `json:"categorie" struct2map:"key:categorie"`
	Descriptif    string `json:"descriptif" struct2map:"key:descriptif"`
	EnStock       string `json:"en_stock" struct2map:"key:en_stock"`
	GrandeImage   string `json:"grande_image" struct2map:"key:grande_image"`
	Marque        string `json:"marque" struct2map:"key:marque"`
	MoyenneImage  string `json:"moyenne_image" struct2map:"key:moyenne_image"`
	Nom           string `json:"nom" struct2map:"key:nom"`
	PetiteImage   string `json:"petite_image" struct2map:"key:petite_image"`
	Prix          string `json:"prix" struct2map:"key:prix"`
	PrixLivraison string `json:"prix_livraison" struct2map:"key:prix_livraison"`
	Reference     string `json:"reference" struct2map:"key:reference"`
	URL           string `json:"url" struct2map:"key:url"`
}

/*
func (o *CheminDeTable) Convert() ([]string, error) {

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
