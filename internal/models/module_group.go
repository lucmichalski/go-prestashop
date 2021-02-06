package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleGroupMgr struct {
	*_BaseMgr
}

// ModuleGroupMgr open func
func ModuleGroupMgr(db *gorm.DB) *_ModuleGroupMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_module_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ModuleGroupMgr) GetTableName() string {
	return "eg_module_group"
}

// Get 获取
func (obj *_ModuleGroupMgr) Get() (result ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ModuleGroupMgr) Gets() (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取
func (obj *_ModuleGroupMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDShop id_shop获取
func (obj *_ModuleGroupMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDGroup id_group获取
func (obj *_ModuleGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// GetByOption 功能选项模式获取
func (obj *_ModuleGroupMgr) GetByOption(opts ...Option) (result ModuleGroup, err error) {
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
func (obj *_ModuleGroupMgr) GetByOptions(opts ...Option) (results []*ModuleGroup, err error) {
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
func (obj *_ModuleGroupMgr) GetFromIDModule(idModule uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_ModuleGroupMgr) GetBatchFromIDModule(idModules []uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_ModuleGroupMgr) GetFromIDShop(idShop uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ModuleGroupMgr) GetBatchFromIDShop(idShops []uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDGroup 通过id_group获取内容
func (obj *_ModuleGroupMgr) GetFromIDGroup(idGroup uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_ModuleGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ModuleGroupMgr) FetchByPrimaryKey(idModule uint32, idShop uint32, idGroup uint32) (result ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_group = ?", idModule, idShop, idGroup).Find(&result).Error

	return
}
