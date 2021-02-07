package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HookModuleMgr struct {
	*_BaseMgr
}

func HookModuleMgr(db *gorm.DB) *_HookModuleMgr {
	if db == nil {
		panic(fmt.Errorf("HookModuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HookModuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_hook_module"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_HookModuleMgr) GetTableName() string {
	return "ps_hook_module"
}

func (obj *_HookModuleMgr) Get() (result HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_HookModuleMgr) Gets() (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_HookModuleMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_HookModuleMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

func (obj *_HookModuleMgr) WithPosition(position uint8) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

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

func (obj *_HookModuleMgr) GetFromIDModule(idModule uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) GetBatchFromIDModule(idModules []uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) GetFromIDShop(idShop uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) GetBatchFromIDShop(idShops []uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) GetFromIDHook(idHook uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) GetBatchFromIDHook(idHooks []uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) GetFromPosition(position uint8) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) GetBatchFromPosition(positions []uint8) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) FetchByPrimaryKey(idModule uint32, idShop uint32, idHook uint32) (result HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_hook = ?", idModule, idShop, idHook).Find(&result).Error

	return
}

func (obj *_HookModuleMgr) FetchIndexByIDModule(idModule uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) FetchIndexByPosition(idShop uint32, position uint8) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND position = ?", idShop, position).Find(&results).Error

	return
}

func (obj *_HookModuleMgr) FetchIndexByIDHook(idHook uint32) (results []*HookModule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}
