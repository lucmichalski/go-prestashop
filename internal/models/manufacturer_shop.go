package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ManufacturerShopMgr struct {
	*_BaseMgr
}

func ManufacturerShopMgr(db *gorm.DB) *_ManufacturerShopMgr {
	if db == nil {
		panic(fmt.Errorf("ManufacturerShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ManufacturerShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_manufacturer_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ManufacturerShopMgr) GetTableName() string {
	return "ps_manufacturer_shop"
}

func (obj *_ManufacturerShopMgr) Get() (result ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ManufacturerShopMgr) Gets() (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ManufacturerShopMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

func (obj *_ManufacturerShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

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

func (obj *_ManufacturerShopMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error

	return
}

func (obj *_ManufacturerShopMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error

	return
}

func (obj *_ManufacturerShopMgr) GetFromIDShop(idShop uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ManufacturerShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ManufacturerShopMgr) FetchByPrimaryKey(idManufacturer uint32, idShop uint32) (result ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ? AND id_shop = ?", idManufacturer, idShop).Find(&result).Error

	return
}

func (obj *_ManufacturerShopMgr) FetchIndexByIDShop(idShop uint32) (results []*ManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
