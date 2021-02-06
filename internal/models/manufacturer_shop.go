package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ManufacturerShopMgr struct {
	*_BaseMgr
}

// ManufacturerShopMgr open func
func ManufacturerShopMgr(db *gorm.DB) *_ManufacturerShopMgr {
	if db == nil {
		panic(fmt.Errorf("ManufacturerShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ManufacturerShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_manufacturer_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ManufacturerShopMgr) GetTableName() string {
	return "eg_manufacturer_shop"
}

// Get 获取
func (obj *_ManufacturerShopMgr) Get() (result ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ManufacturerShopMgr) Gets() (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDManufacturer id_manufacturer获取
func (obj *_ManufacturerShopMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

// WithIDShop id_shop获取
func (obj *_ManufacturerShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_ManufacturerShopMgr) GetByOption(opts ...Option) (result ManufacturerShop, err error) {
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
func (obj *_ManufacturerShopMgr) GetByOptions(opts ...Option) (results []*ManufacturerShop, err error) {
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

// GetFromIDManufacturer 通过id_manufacturer获取内容
func (obj *_ManufacturerShopMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error

	return
}

// GetBatchFromIDManufacturer 批量唯一主键查找
func (obj *_ManufacturerShopMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ManufacturerShopMgr) GetFromIDShop(idShop uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ManufacturerShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ManufacturerShopMgr) FetchByPrimaryKey(idManufacturer uint32, idShop uint32) (result ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ? AND id_shop = ?", idManufacturer, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_ManufacturerShopMgr) FetchIndexByIDShop(idShop uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
