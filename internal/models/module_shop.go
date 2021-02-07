package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleShopMgr struct {
	*_BaseMgr
}

func ModuleShopMgr(db *gorm.DB) *_ModuleShopMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ModuleShopMgr) GetTableName() string {
	return "ps_module_shop"
}

func (obj *_ModuleShopMgr) Get() (result ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ModuleShopMgr) Gets() (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ModuleShopMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_ModuleShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ModuleShopMgr) WithEnableDevice(enableDevice bool) Option {
	return optionFunc(func(o *options) { o.query["enable_device"] = enableDevice })
}

func (obj *_ModuleShopMgr) GetByOption(opts ...Option) (result ModuleShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ModuleShopMgr) GetByOptions(opts ...Option) (results []*ModuleShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ModuleShopMgr) GetFromIDModule(idModule uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_ModuleShopMgr) GetBatchFromIDModule(idModules []uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_ModuleShopMgr) GetFromIDShop(idShop uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ModuleShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ModuleShopMgr) GetFromEnableDevice(enableDevice bool) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enable_device = ?", enableDevice).Find(&results).Error

	return
}

func (obj *_ModuleShopMgr) GetBatchFromEnableDevice(enableDevices []bool) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enable_device IN (?)", enableDevices).Find(&results).Error

	return
}

func (obj *_ModuleShopMgr) FetchByPrimaryKey(idModule uint32, idShop uint32) (result ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ?", idModule, idShop).Find(&result).Error

	return
}

func (obj *_ModuleShopMgr) FetchIndexByIDShop(idShop uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
