package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WarehouseShopMgr struct {
	*_BaseMgr
}

func WarehouseShopMgr(db *gorm.DB) *_WarehouseShopMgr {
	if db == nil {
		panic(fmt.Errorf("WarehouseShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WarehouseShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_warehouse_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_WarehouseShopMgr) GetTableName() string {
	return "ps_warehouse_shop"
}

func (obj *_WarehouseShopMgr) Get() (result WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_WarehouseShopMgr) Gets() (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_WarehouseShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_WarehouseShopMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

func (obj *_WarehouseShopMgr) GetByOption(opts ...Option) (result WarehouseShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_WarehouseShopMgr) GetByOptions(opts ...Option) (results []*WarehouseShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_WarehouseShopMgr) GetFromIDShop(idShop uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_WarehouseShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_WarehouseShopMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_WarehouseShopMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}


func (obj *_WarehouseShopMgr) FetchByPrimaryKey(idShop uint32, idWarehouse uint32) (result WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_warehouse = ?", idShop, idWarehouse).Find(&result).Error

	return
}

func (obj *_WarehouseShopMgr) FetchIndexByIDShop(idShop uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_WarehouseShopMgr) FetchIndexByIDWarehouse(idWarehouse uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}
