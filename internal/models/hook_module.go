package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HookModuleMgr struct {
	*_BaseMgr
}

// HookModuleMgr open func
func HookModuleMgr(db *gorm.DB) *_HookModuleMgr {
	if db == nil {
		panic(fmt.Errorf("HookModuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HookModuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_hook_module"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HookModuleMgr) GetTableName() string {
	return "eg_hook_module"
}

// Get 获取
func (obj *_HookModuleMgr) Get() (result HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HookModuleMgr) Gets() (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取
func (obj *_HookModuleMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDShop id_shop获取
func (obj *_HookModuleMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDHook id_hook获取
func (obj *_HookModuleMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

// WithPosition position获取
func (obj *_HookModuleMgr) WithPosition(position uint8) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// GetByOption 功能选项模式获取
func (obj *_HookModuleMgr) GetByOption(opts ...Option) (result HookModule, err error) {
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
func (obj *_HookModuleMgr) GetByOptions(opts ...Option) (results []*HookModule, err error) {
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
func (obj *_HookModuleMgr) GetFromIDModule(idModule uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_HookModuleMgr) GetBatchFromIDModule(idModules []uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_HookModuleMgr) GetFromIDShop(idShop uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_HookModuleMgr) GetBatchFromIDShop(idShops []uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDHook 通过id_hook获取内容
func (obj *_HookModuleMgr) GetFromIDHook(idHook uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}

// GetBatchFromIDHook 批量唯一主键查找
func (obj *_HookModuleMgr) GetBatchFromIDHook(idHooks []uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_HookModuleMgr) GetFromPosition(position uint8) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_HookModuleMgr) GetBatchFromPosition(positions []uint8) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_HookModuleMgr) FetchByPrimaryKey(idModule uint32, idShop uint32, idHook uint32) (result HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_hook = ?", idModule, idShop, idHook).Find(&result).Error

	return
}

// FetchIndexByIDModule  获取多个内容
func (obj *_HookModuleMgr) FetchIndexByIDModule(idModule uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// FetchIndexByPosition  获取多个内容
func (obj *_HookModuleMgr) FetchIndexByPosition(idShop uint32, position uint8) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND position = ?", idShop, position).Find(&results).Error

	return
}

// FetchIndexByIDHook  获取多个内容
func (obj *_HookModuleMgr) FetchIndexByIDHook(idHook uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}
