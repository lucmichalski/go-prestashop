package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCountryTaxMgr struct {
	*_BaseMgr
}

func ProductCountryTaxMgr(db *gorm.DB) *_ProductCountryTaxMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCountryTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCountryTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_country_tax"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductCountryTaxMgr) GetTableName() string {
	return "ps_product_country_tax"
}

func (obj *_ProductCountryTaxMgr) Get() (result ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductCountryTaxMgr) Gets() (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ProductCountryTaxMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductCountryTaxMgr) WithIDCountry(idCountry int) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_ProductCountryTaxMgr) WithIDTax(idTax int) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

func (obj *_ProductCountryTaxMgr) GetByOption(opts ...Option) (result ProductCountryTax, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductCountryTaxMgr) GetByOptions(opts ...Option) (results []*ProductCountryTax, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ProductCountryTaxMgr) GetFromIDProduct(idProduct int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductCountryTaxMgr) GetBatchFromIDProduct(idProducts []int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductCountryTaxMgr) GetFromIDCountry(idCountry int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_ProductCountryTaxMgr) GetBatchFromIDCountry(idCountrys []int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_ProductCountryTaxMgr) GetFromIDTax(idTax int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}

func (obj *_ProductCountryTaxMgr) GetBatchFromIDTax(idTaxs []int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

func (obj *_ProductCountryTaxMgr) FetchByPrimaryKey(idProduct int, idCountry int) (result ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_country = ?", idProduct, idCountry).Find(&result).Error

	return
}
