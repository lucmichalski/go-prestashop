package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/kare-click.csv
// Catalog Separator: |
// Catalog excerpt: #1

// ===================================================================================== \\
nom usuel du produit: Suspension Saloon Flowers 5 Kare Design,
reference interne: 30001,
prix TTC du produit: 129,
categorie: Luminaire > Lustre & suspension,
URL: https://tkh.kare-click.fr/?P4A60145983227S1UC412881015464V4,
numero modele produit: ,
marque: Kare Design,
code ean: 4025621300014,
reference constructeur: ,
frais de ports: 8,
descriptif du produit: <p style= text-align: justify; >La réplique du fameux lustre Saloon Flowers en dimension plus modeste et qui s'adaptera facilement aux petits espaces avec ses fils d'1 m de descente. Les formes classiques des abat-jours se marient parfaitement aux motifs colorés, de quoi égayer votre intérieur.</p> <p style= text-align: justify; >Classe d'efficacité A-E, 230 V, 50/60 Hz. Douille E14. Nombre d'ampoules : 5, max. 40 Watts. Ampoules non incluses. Ampoules remplaçables.</p>,
garantie: 0,
indicateur de stock: 1,
disponibilite du produit: ,
indicateur de performance: ,
indicateur de promotion: 2,
indicateur de nouveaute: ,
URL image petite: https://www.kare-click.fr/32551/suspension-saloon-flowers-5-kare-design.jpg,
URL image moyenne: https://www.kare-click.fr/52120/suspension-saloon-flowers-5-kare-design.jpg,
URL image grande: https://www.kare-click.fr/52122/suspension-saloon-flowers-5-kare-design.jpg,
ecotaxe: 0,
prix barre: 
\\ ===================================================================================== //

*/

type KareClick struct {
	Categorie               string `json:"categorie" struct2map:"key:categorie"`
	CodeEan                 string `json:"code_ean" struct2map:"key:code_ean"`
	DescriptifDuProduit     string `json:"descriptif_du_produit" struct2map:"key:descriptif_du_produit"`
	DisponibiliteDuProduit  string `json:"disponibilite_du_produit" struct2map:"key:disponibilite_du_produit"`
	Ecotaxe                 string `json:"ecotaxe" struct2map:"key:ecotaxe"`
	FraisDePorts            string `json:"frais_de_ports" struct2map:"key:frais_de_ports"`
	Garantie                string `json:"garantie" struct2map:"key:garantie"`
	IndicateurDeNouveaute   string `json:"indicateur_de_nouveaute" struct2map:"key:indicateur_de_nouveaute"`
	IndicateurDePerformance string `json:"indicateur_de_performance" struct2map:"key:indicateur_de_performance"`
	IndicateurDePromotion   string `json:"indicateur_de_promotion" struct2map:"key:indicateur_de_promotion"`
	IndicateurDeStock       string `json:"indicateur_de_stock" struct2map:"key:indicateur_de_stock"`
	Marque                  string `json:"marque" struct2map:"key:marque"`
	NomUsuelDuProduit       string `json:"nom_usuel_du_produit" struct2map:"key:nom_usuel_du_produit"`
	NumeroModeleProduit     string `json:"numero_modele_produit" struct2map:"key:numero_modele_produit"`
	PrixBarre               string `json:"prix_barre" struct2map:"key:prix_barre"`
	PrixTtcDuProduit        string `json:"prix_ttc_du_produit" struct2map:"key:prix_ttc_du_produit"`
	ReferenceConstructeur   string `json:"reference_constructeur" struct2map:"key:reference_constructeur"`
	ReferenceInterne        string `json:"reference_interne" struct2map:"key:reference_interne"`
	URL                     string `json:"url" struct2map:"key:url"`
	URLImageGrande          string `json:"url_image_grande" struct2map:"key:url_image_grande"`
	URLImageMoyenne         string `json:"url_image_moyenne" struct2map:"key:url_image_moyenne"`
	URLImagePetite          string `json:"url_image_petite" struct2map:"key:url_image_petite"`
}

/*
func (o *KareClick) Convert() ([]string, error) {

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
