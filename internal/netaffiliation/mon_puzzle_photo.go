package netaffiliation

import (
	"fmt"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
)

/*

// Catalog Path: catalogs/mon-puzzle-photo.csv
// Catalog Separator: ;
// Catalog excerpt: #1

// ===================================================================================== \\
Id: puzzle-100,
Name: Puzzle photo 100 pièces,
URL: https://action.metaffiliation.com/trk.php?mclic=P4B90D45983215S1U9412a2b020V4,
Price: 24.99,
Formatted Price: 24,99 €,
Description: Souhaitez-vous en tant que débutant d'abord faire connaissance avec nos puzzles photo? Ou souhaitez-vous offrir un cadeau original à une personne qui ne trouve pas toujours le loisir de faire un puzzle? Alors nous vous conseillons d’opter pour un de nos plus petits puzzles photo premium avec boîte. Ce puzzle donne de la vie à vos photos ou images personnelles en plaçant seulement 100 pièces. Ce puzzle photo de 100 pièces initiera chacun au monde des puzzles photo. De plus, vous reconnaîtrez dans ce puzzle photo prémium notre haute qualité habituelle à un prix avantageux.,
Image: https://www.monpuzzlephoto.fr/images/default/monpuzzlephoto_fr_common/products/puzzle-100/1/pi-max-800x600.jpg?v=20190918-01,
Category: Toys & Games > Games,
Shipping: FR::Standard:4.99,
Availability: in stock,
Condition: new
\\ ===================================================================================== //

*/

type MonPuzzlePhoto struct {
	Availability   string `json:"availability" struct2map:"key:availability"`
	Category       string `json:"category" struct2map:"key:category"`
	Condition      string `json:"condition" struct2map:"key:condition"`
	Description    string `json:"description" struct2map:"key:description"`
	FormattedPrice string `json:"formatted_price" struct2map:"key:formatted_price"`
	ID             string `json:"id" struct2map:"key:id"`
	Image          string `json:"image" struct2map:"key:image"`
	Name           string `json:"name" struct2map:"key:name"`
	Price          string `json:"price" struct2map:"key:price"`
	Shipping       string `json:"shipping" struct2map:"key:shipping"`
	URL            string `json:"url" struct2map:"key:url"`
}

/*
func (o *MonPuzzlePhoto) Convert() ([]string, error) {

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
