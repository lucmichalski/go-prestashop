package netaffiliation

import (
	"strings"

	"gorm.io/gorm"

	"github.com/lucmichalski/go-prestashop/internal/models"
)

var ( 
	db *gorm.DB
	dbTablePrefix string
	Catalogs = []interface{}{
					
					&Flux2{},
					&Villatech{},
					&PneusOnlineFr{},
					&DcshoesFr{},
					&Bazaravenue{},
					&Bestform{},
					&FannyChaussure{},
					&LentillesMoinsChère{},
					&Bricorama{},
					&LauraKent{},
					&RougierAndPlé{},
					&GsellMaroquinerieEtArticle{},
					&Abix{},
					&TradeDiscount{},
					&FamilleMaryMielEtPlante{},
					&Roxy{},
					&PlanetPuzzle{},
					&Quiksilver{},
					&Meltonic{},
					&KareClick{},
					&FouDePuzzle{},
					&Homemaison{},
					&CheminDeTable{},
					&VeniVici{},
					&EspacePlaisir{},
					&Condomz{},
					&Thalassoissambre{},
					&Vitalco{},
					&MonPuzzlePhoto{},
					&Layole{},
					&Hypnium{},
					&ShopMini{},
					&Craftine{},
					&ClefEnLigne{},
					&Deelux74{},
					&BoutiqueArthur{},
					&LeRoiDuMatelasFr{},
					&Newchic{},
					&OriginalCup{},
					&ZafulFr{},
					&DresslilyFr{},
				}
)

func initDB(client *gorm.DB, tablePrefix string) *gorm.DB {
	db = client
	dbTablePrefix = tablePrefix
	return db
}

func findOrCreateCategoryByName(db *gorm.DB, dbTablePrefix, name string) (*models.CategoryLang, error) {
	var cat models.CategoryLang
	db.Debug().Raw("SELECT * FROM " + dbTablePrefix + "category_lang WHERE name = '" + strings.ToLower(name) + "'").Scan(&cat)
	if cat.IDCategory == 0 {
		db.Debug().Raw("SELECT * FROM eg_lang WHERE iso_code = '" + strings.ToLower(iso_code) + "'").Scan(&lang)
		return &lang, nil // todo. fix this hacky return
	}
	return &cat, nil
}