package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleShopMgr struct {
	*_BaseMgr
}

// ModuleShopMgr open func
func ModuleShopMgr(db *gorm.DB) *_ModuleShopMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ModuleShopMgr) GetTableName() string {
	return "ps_module_shop"
}

// Get 获取
func (obj *_ModuleShopMgr) Get() (result ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ModuleShopMgr) Gets() (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取
func (obj *_ModuleShopMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDShop id_shop获取
func (obj *_ModuleShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithEnableDevice enable_device获取
func (obj *_ModuleShopMgr) WithEnableDevice(enableDevice bool) Option {
	return optionFunc(func(o *options) { o.query["enable_device"] = enableDevice })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDModule 通过id_module获取内容
func (obj *_ModuleShopMgr) GetFromIDModule(idModule uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_ModuleShopMgr) GetBatchFromIDModule(idModules []uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ModuleShopMgr) GetFromIDShop(idShop uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ModuleShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromEnableDevice 通过enable_device获取内容
func (obj *_ModuleShopMgr) GetFromEnableDevice(enableDevice bool) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enable_device = ?", enableDevice).Find(&results).Error

	return
}

// GetBatchFromEnableDevice 批量唯一主键查找
func (obj *_ModuleShopMgr) GetBatchFromEnableDevice(enableDevices []bool) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enable_device IN (?)", enableDevices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ModuleShopMgr) FetchByPrimaryKey(idModule uint32, idShop uint32) (result ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ?", idModule, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_ModuleShopMgr) FetchIndexByIDShop(idShop uint32) (results []*ModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
