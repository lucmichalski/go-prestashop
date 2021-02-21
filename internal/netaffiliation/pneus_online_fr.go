package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/pneus-online-fr.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
id: 1214,
marque: Continental,
modele: Conti-WinterContact TS 800,
largeur: 175,
hauteur: 65,
diametre: 13,
vitesse: T,
charge: 80,
saison: Winter,
runflat: non,
renforce: non,
vehicule: car,
prix: 71,76,
image: https://img2-pneusonlinesuiss.netdna-ssl.com/produit/pneu-auto/120/CONTINENTAL_TS800.jpg,
url: https://action.metaffiliation.com/trk.php?mclic=P4455F4598321BS1UD4127430697703V4,
ean: 4019238314441,
rechape: non,
note economie (eu label): E,
note adherance (eu label): C,
note bruit (eu label): B,
bruit dB (eu label): 71,
: 
\\ ===================================================================================== //

*/

type PneusOnlineFr struct {
	BruitdbEuLabel       string `json:"bruit_d_b_eu_label" struct2map:"key:bruit_d_b_eu_label"`
	Charge               string `json:"charge" struct2map:"key:charge"`
	Diametre             string `json:"diametre" struct2map:"key:diametre"`
	Ean                  string `json:"ean" struct2map:"key:ean"`
	Hauteur              string `json:"hauteur" struct2map:"key:hauteur"`
	ID                   string `json:"id" struct2map:"key:id"`
	Image                string `json:"image" struct2map:"key:image"`
	Largeur              string `json:"largeur" struct2map:"key:largeur"`
	Marque               string `json:"marque" struct2map:"key:marque"`
	Modele               string `json:"modele" struct2map:"key:modele"`
	NoteAdheranceEuLabel string `json:"note_adherance_eu_label" struct2map:"key:note_adherance_eu_label"`
	NoteBruitEuLabel     string `json:"note_bruit_eu_label" struct2map:"key:note_bruit_eu_label"`
	NoteEconomieEuLabel  string `json:"note_economie_eu_label" struct2map:"key:note_economie_eu_label"`
	Prix                 string `json:"prix" struct2map:"key:prix"`
	Rechape              string `json:"rechape" struct2map:"key:rechape"`
	Renforce             string `json:"renforce" struct2map:"key:renforce"`
	Runflat              string `json:"runflat" struct2map:"key:runflat"`
	Saison               string `json:"saison" struct2map:"key:saison"`
	URL                  string `json:"url" struct2map:"key:url"`
	Vehicule             string `json:"vehicule" struct2map:"key:vehicule"`
	Vitesse              string `json:"vitesse" struct2map:"key:vitesse"`
}

/*
func (o *PneusOnlineFr) Convert() ([]string, error) {

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
