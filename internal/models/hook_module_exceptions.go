package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HookModuleExceptionsMgr struct {
	*_BaseMgr
}

func HookModuleExceptionsMgr(db *gorm.DB) *_HookModuleExceptionsMgr {
	if db == nil {
		panic(fmt.Errorf("HookModuleExceptionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HookModuleExceptionsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_hook_module_exceptions"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_HookModuleExceptionsMgr) GetTableName() string {
	return "ps_hook_module_exceptions"
}

func (obj *_HookModuleExceptionsMgr) Get() (result HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_HookModuleExceptionsMgr) Gets() (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_HookModuleExceptionsMgr) WithIDHookModuleExceptions(idHookModuleExceptions uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook_module_exceptions"] = idHookModuleExceptions })
}

func (obj *_HookModuleExceptionsMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_HookModuleExceptionsMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_HookModuleExceptionsMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

func (obj *_HookModuleExceptionsMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

func (obj *_HookModuleExceptionsMgr) GetByOption(opts ...Option) (result HookModuleExceptions, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetByOptions(opts ...Option) (results []*HookModuleExceptions, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_HookModuleExceptionsMgr) GetFromIDHookModuleExceptions(idHookModuleExceptions uint32) (result HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions = ?", idHookModuleExceptions).Find(&result).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetBatchFromIDHookModuleExceptions(idHookModuleExceptionss []uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions IN (?)", idHookModuleExceptionss).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetFromIDShop(idShop uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetBatchFromIDShop(idShops []uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetFromIDModule(idModule uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetBatchFromIDModule(idModules []uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetFromIDHook(idHook uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetBatchFromIDHook(idHooks []uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetFromFileName(fileName string) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) GetBatchFromFileName(fileNames []string) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error

	return
}


func (obj *_HookModuleExceptionsMgr) FetchByPrimaryKey(idHookModuleExceptions uint32) (result HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions = ?", idHookModuleExceptions).Find(&result).Error

	return
}

func (obj *_HookModuleExceptionsMgr) FetchIndexByIDModule(idModule uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_HookModuleExceptionsMgr) FetchIndexByIDHook(idHook uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}
