package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SupplierShopMgr struct {
	*_BaseMgr
}

func SupplierShopMgr(db *gorm.DB) *_SupplierShopMgr {
	if db == nil {
		panic(fmt.Errorf("SupplierShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplierShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supplier_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SupplierShopMgr) GetTableName() string {
	return "ps_supplier_shop"
}

func (obj *_SupplierShopMgr) Get() (result SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SupplierShopMgr) Gets() (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SupplierShopMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

func (obj *_SupplierShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_SupplierShopMgr) GetByOption(opts ...Option) (result SupplierShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SupplierShopMgr) GetByOptions(opts ...Option) (results []*SupplierShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SupplierShopMgr) GetFromIDSupplier(idSupplier uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

func (obj *_SupplierShopMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

func (obj *_SupplierShopMgr) GetFromIDShop(idShop uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_SupplierShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_SupplierShopMgr) FetchByPrimaryKey(idSupplier uint32, idShop uint32) (result SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ? AND id_shop = ?", idSupplier, idShop).Find(&result).Error

	return
}

func (obj *_SupplierShopMgr) FetchIndexByIDShop(idShop uint32) (results []*SupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
