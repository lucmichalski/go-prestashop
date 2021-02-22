package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/maison-ft.csv
// Catalog Separator: ,
// Catalog excerpt: #1

// ===================================================================================== \\
ID: 5791560073374,
Maximum Price: 59.00,
Product Description:  Infos 
 Conseils de taille
 Entretien 


🍃Bruno est un sweatshirt 100% coton bio | 350 g/m² tissé à Porto puis brodé à Avignon. 🚲 Il a parcouru en moyenne 1 580 kilomètres avant d'arriver dans votre armoire. La petite histoire Parce que nous aimons revendiquer notre côté Frenchy [et vous aussi peut être?], voici Bruno pour les petits, et sa broderie French en bouclette, la simplicité et l'élégance à la Française.

 Nos sweats enfant taillent normalement et sont légèrement cintrés. . Prenez la taille de votre enfant pour une utilisation immédiate. Prenez la taille au-dessus pour que le sweat lui aille plus longtemps. Entre deux tailles, on vous conseille la taille au-dessus. ↕️ Louane mesure 122 cm pour 22kg et porte du 6 ans. ↕️ César mesure 116 cm pour 20kg et porte du 6 ans.



Bruno est un sweatshirt 100% coton bio, il aime l'entretien délicat. Lavage On vous conseille de le laver à 30 degrés, retourné. Repassage A repasser sur l’envers avec Amour.  

,
Published Scope: web,
Tags: BIO, Blanc, Bleu, Coton, eco-responsable, France, Français, French, Frenchie, Frenchy, sweat en coton bio, sweatshirt, t-shirt bio enfant, vêtement responsable,
Total Sales: 236.00,
Created At: 2020-10-07T10:19:03,
Inventory Quantity: 13,
Minimum Compare At Price: 59.00,
Product Type: T-shirt BIO enfant,
Seo Description: 1,
Template Suffix: sweatshirt,
Total Tax: 37.36,
URL: https://why.maisonft.com/?P5114B34598317S1U84131df01V4,
Handle: sweatshirt-bruno-enfant-coton-bio-1,
Maximum Compare At Price: 59.00,
Minimum Price: 59.00,
Published At: 2020-10-07T10:19:03,
Seo Title: ,
Title: Sweatshirt Bruno Enfant - Coton Bio,
Total Units Sold: 4,
Vendor: Maison FT,
Updated At: 2020-11-10T10:50:55,
Visible: 1,
Image Image Contents: ,
Image Image Name: 
\\ ===================================================================================== //

*/

type MaisonFt struct {
	CreatedAt             string `json:"created_at" struct2map:"key:created_at"`
	Handle                string `json:"handle" struct2map:"key:handle"`
	ID                    string `json:"id" struct2map:"key:id"`
	ImageImageContents    string `json:"image_image_contents" struct2map:"key:image_image_contents"`
	ImageImageName        string `json:"image_image_name" struct2map:"key:image_image_name"`
	InventoryQuantity     string `json:"inventory_quantity" struct2map:"key:inventory_quantity"`
	MaximumCompareAtPrice string `json:"maximum_compare_at_price" struct2map:"key:maximum_compare_at_price"`
	MaximumPrice          string `json:"maximum_price" struct2map:"key:maximum_price"`
	MinimumCompareAtPrice string `json:"minimum_compare_at_price" struct2map:"key:minimum_compare_at_price"`
	MinimumPrice          string `json:"minimum_price" struct2map:"key:minimum_price"`
	ProductDescription    string `json:"product_description" struct2map:"key:product_description"`
	ProductType           string `json:"product_type" struct2map:"key:product_type"`
	PublishedAt           string `json:"published_at" struct2map:"key:published_at"`
	PublishedScope        string `json:"published_scope" struct2map:"key:published_scope"`
	SeoDescription        string `json:"seo_description" struct2map:"key:seo_description"`
	SeoTitle              string `json:"seo_title" struct2map:"key:seo_title"`
	Tags                  string `json:"tags" struct2map:"key:tags"`
	TemplateSuffix        string `json:"template_suffix" struct2map:"key:template_suffix"`
	Title                 string `json:"title" struct2map:"key:title"`
	TotalSales            string `json:"total_sales" struct2map:"key:total_sales"`
	TotalTax              string `json:"total_tax" struct2map:"key:total_tax"`
	TotalUnitsSold        string `json:"total_units_sold" struct2map:"key:total_units_sold"`
	UpdatedAt             string `json:"updated_at" struct2map:"key:updated_at"`
	URL                   string `json:"url" struct2map:"key:url"`
	Vendor                string `json:"vendor" struct2map:"key:vendor"`
	Visible               string `json:"visible" struct2map:"key:visible"`
}

/*
func (o *MaisonFt) Convert() ([]string, error) {

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
