package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCarrierMgr struct {
	*_BaseMgr
}

func ProductCarrierMgr(db *gorm.DB) *_ProductCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_carrier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductCarrierMgr) GetTableName() string {
	return "ps_product_carrier"
}

func (obj *_ProductCarrierMgr) Get() (result ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductCarrierMgr) Gets() (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductCarrierMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductCarrierMgr) WithIDCarrierReference(idCarrierReference uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier_reference"] = idCarrierReference })
}

func (obj *_ProductCarrierMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ProductCarrierMgr) GetByOption(opts ...Option) (result ProductCarrier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductCarrierMgr) GetByOptions(opts ...Option) (results []*ProductCarrier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductCarrierMgr) GetFromIDProduct(idProduct uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductCarrierMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductCarrierMgr) GetFromIDCarrierReference(idCarrierReference uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier_reference = ?", idCarrierReference).Find(&results).Error

	return
}

func (obj *_ProductCarrierMgr) GetBatchFromIDCarrierReference(idCarrierReferences []uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier_reference IN (?)", idCarrierReferences).Find(&results).Error

	return
}

func (obj *_ProductCarrierMgr) GetFromIDShop(idShop uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ProductCarrierMgr) GetBatchFromIDShop(idShops []uint32) (results []*ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_ProductCarrierMgr) FetchByPrimaryKey(idProduct uint32, idCarrierReference uint32, idShop uint32) (result ProductCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_carrier_reference = ? AND id_shop = ?", idProduct, idCarrierReference, idShop).Find(&result).Error

	return
}
