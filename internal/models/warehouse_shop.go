package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WarehouseShopMgr struct {
	*_BaseMgr
}

// WarehouseShopMgr open func
func WarehouseShopMgr(db *gorm.DB) *_WarehouseShopMgr {
	if db == nil {
		panic(fmt.Errorf("WarehouseShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WarehouseShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_warehouse_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WarehouseShopMgr) GetTableName() string {
	return "ps_warehouse_shop"
}

// Get 获取
func (obj *_WarehouseShopMgr) Get() (result WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WarehouseShopMgr) Gets() (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDShop id_shop获取
func (obj *_WarehouseShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDWarehouse id_warehouse获取
func (obj *_WarehouseShopMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDShop 通过id_shop获取内容
func (obj *_WarehouseShopMgr) GetFromIDShop(idShop uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_WarehouseShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDWarehouse 通过id_warehouse获取内容
func (obj *_WarehouseShopMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找
func (obj *_WarehouseShopMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WarehouseShopMgr) FetchByPrimaryKey(idShop uint32, idWarehouse uint32) (result WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_warehouse = ?", idShop, idWarehouse).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_WarehouseShopMgr) FetchIndexByIDShop(idShop uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// FetchIndexByIDWarehouse  获取多个内容
func (obj *_WarehouseShopMgr) FetchIndexByIDWarehouse(idWarehouse uint32) (results []*WarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}
