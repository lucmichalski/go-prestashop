package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredPriceIndexMgr struct {
	*_BaseMgr
}

func LayeredPriceIndexMgr(db *gorm.DB) *_LayeredPriceIndexMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredPriceIndexMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredPriceIndexMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_price_index"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredPriceIndexMgr) GetTableName() string {
	return "ps_layered_price_index"
}

func (obj *_LayeredPriceIndexMgr) Get() (result LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredPriceIndexMgr) Gets() (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_LayeredPriceIndexMgr) WithIDCurrency(idCurrency int) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_LayeredPriceIndexMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_LayeredPriceIndexMgr) WithPriceMin(priceMin float64) Option {
	return optionFunc(func(o *options) { o.query["price_min"] = priceMin })
}

func (obj *_LayeredPriceIndexMgr) WithPriceMax(priceMax float64) Option {
	return optionFunc(func(o *options) { o.query["price_max"] = priceMax })
}

func (obj *_LayeredPriceIndexMgr) WithIDCountry(idCountry int) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_LayeredPriceIndexMgr) GetByOption(opts ...Option) (result LayeredPriceIndex, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetByOptions(opts ...Option) (results []*LayeredPriceIndex, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetFromIDProduct(idProduct int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetBatchFromIDProduct(idProducts []int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetFromIDCurrency(idCurrency int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetBatchFromIDCurrency(idCurrencys []int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetFromIDShop(idShop int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetBatchFromIDShop(idShops []int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetFromPriceMin(priceMin float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_min = ?", priceMin).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetBatchFromPriceMin(priceMins []float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_min IN (?)", priceMins).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetFromPriceMax(priceMax float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_max = ?", priceMax).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetBatchFromPriceMax(priceMaxs []float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_max IN (?)", priceMaxs).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetFromIDCountry(idCountry int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) GetBatchFromIDCountry(idCountrys []int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) FetchByPrimaryKey(idProduct int, idCurrency int, idShop int, idCountry int) (result LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_currency = ? AND id_shop = ? AND id_country = ?", idProduct, idCurrency, idShop, idCountry).Find(&result).Error

	return
}

func (obj *_LayeredPriceIndexMgr) FetchIndexByIDCurrency(idCurrency int) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) FetchIndexByPriceMin(priceMin float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_min = ?", priceMin).Find(&results).Error

	return
}

func (obj *_LayeredPriceIndexMgr) FetchIndexByPriceMax(priceMax float64) (results []*LayeredPriceIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_max = ?", priceMax).Find(&results).Error

	return
}
