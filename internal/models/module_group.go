package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleGroupMgr struct {
	*_BaseMgr
}

func ModuleGroupMgr(db *gorm.DB) *_ModuleGroupMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ModuleGroupMgr) GetTableName() string {
	return "ps_module_group"
}

func (obj *_ModuleGroupMgr) Get() (result ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ModuleGroupMgr) Gets() (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ModuleGroupMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_ModuleGroupMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ModuleGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

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

func (obj *_ModuleGroupMgr) GetFromIDModule(idModule uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_ModuleGroupMgr) GetBatchFromIDModule(idModules []uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_ModuleGroupMgr) GetFromIDShop(idShop uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ModuleGroupMgr) GetBatchFromIDShop(idShops []uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ModuleGroupMgr) GetFromIDGroup(idGroup uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_ModuleGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_ModuleGroupMgr) FetchByPrimaryKey(idModule uint32, idShop uint32, idGroup uint32) (result ModuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_group = ?", idModule, idShop, idGroup).Find(&result).Error

	return
}
