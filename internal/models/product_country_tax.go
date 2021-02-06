package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCountryTaxMgr struct {
	*_BaseMgr
}

// ProductCountryTaxMgr open func
func ProductCountryTaxMgr(db *gorm.DB) *_ProductCountryTaxMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCountryTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCountryTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_country_tax"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductCountryTaxMgr) GetTableName() string {
	return "eg_product_country_tax"
}

// Get 获取
func (obj *_ProductCountryTaxMgr) Get() (result ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductCountryTaxMgr) Gets() (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取
func (obj *_ProductCountryTaxMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDCountry id_country获取
func (obj *_ProductCountryTaxMgr) WithIDCountry(idCountry int) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDTax id_tax获取
func (obj *_ProductCountryTaxMgr) WithIDTax(idTax int) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDProduct 通过id_product获取内容
func (obj *_ProductCountryTaxMgr) GetFromIDProduct(idProduct int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_ProductCountryTaxMgr) GetBatchFromIDProduct(idProducts []int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDCountry 通过id_country获取内容
func (obj *_ProductCountryTaxMgr) GetFromIDCountry(idCountry int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

// GetBatchFromIDCountry 批量唯一主键查找
func (obj *_ProductCountryTaxMgr) GetBatchFromIDCountry(idCountrys []int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

// GetFromIDTax 通过id_tax获取内容
func (obj *_ProductCountryTaxMgr) GetFromIDTax(idTax int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}

// GetBatchFromIDTax 批量唯一主键查找
func (obj *_ProductCountryTaxMgr) GetBatchFromIDTax(idTaxs []int) (results []*ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductCountryTaxMgr) FetchByPrimaryKey(idProduct int, idCountry int) (result ProductCountryTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_country = ?", idProduct, idCountry).Find(&result).Error

	return
}
