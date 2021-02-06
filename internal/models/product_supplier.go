package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductSupplierMgr struct {
	*_BaseMgr
}

func ProductSupplierMgr(db *gorm.DB) *_ProductSupplierMgr {
	if db == nil {
		panic(fmt.Errorf("ProductSupplierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductSupplierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_supplier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductSupplierMgr) GetTableName() string {
	return "ps_product_supplier"
}

func (obj *_ProductSupplierMgr) Get() (result ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductSupplierMgr) Gets() (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductSupplierMgr) WithIDProductSupplier(idProductSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_supplier"] = idProductSupplier })
}

func (obj *_ProductSupplierMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductSupplierMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_ProductSupplierMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

func (obj *_ProductSupplierMgr) WithProductSupplierReference(productSupplierReference string) Option {
	return optionFunc(func(o *options) { o.query["product_supplier_reference"] = productSupplierReference })
}

func (obj *_ProductSupplierMgr) WithProductSupplierPriceTe(productSupplierPriceTe float64) Option {
	return optionFunc(func(o *options) { o.query["product_supplier_price_te"] = productSupplierPriceTe })
}

func (obj *_ProductSupplierMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_ProductSupplierMgr) GetByOption(opts ...Option) (result ProductSupplier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductSupplierMgr) GetByOptions(opts ...Option) (results []*ProductSupplier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductSupplierMgr) GetFromIDProductSupplier(idProductSupplier uint32) (result ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_supplier = ?", idProductSupplier).Find(&result).Error

	return
}

func (obj *_ProductSupplierMgr) GetBatchFromIDProductSupplier(idProductSuppliers []uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_supplier IN (?)", idProductSuppliers).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetFromIDProduct(idProduct uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetFromIDSupplier(idSupplier uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetFromProductSupplierReference(productSupplierReference string) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_reference = ?", productSupplierReference).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetBatchFromProductSupplierReference(productSupplierReferences []string) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_reference IN (?)", productSupplierReferences).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetFromProductSupplierPriceTe(productSupplierPriceTe float64) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_price_te = ?", productSupplierPriceTe).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetBatchFromProductSupplierPriceTe(productSupplierPriceTes []float64) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_price_te IN (?)", productSupplierPriceTes).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetFromIDCurrency(idCurrency uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_ProductSupplierMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}


func (obj *_ProductSupplierMgr) FetchByPrimaryKey(idProductSupplier uint32) (result ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_supplier = ?", idProductSupplier).Find(&result).Error

	return
}

func (obj *_ProductSupplierMgr) FetchUniqueIndexByIDProduct(idProduct uint32, idProductAttribute uint32, idSupplier uint32) (result ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_product_attribute = ? AND id_supplier = ?", idProduct, idProductAttribute, idSupplier).Find(&result).Error

	return
}

func (obj *_ProductSupplierMgr) FetchIndexByIDSupplier(idProduct uint32, idSupplier uint32) (results []*ProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_supplier = ?", idProduct, idSupplier).Find(&results).Error

	return
}
