package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HookModuleExceptionsMgr struct {
	*_BaseMgr
}

// HookModuleExceptionsMgr open func
func HookModuleExceptionsMgr(db *gorm.DB) *_HookModuleExceptionsMgr {
	if db == nil {
		panic(fmt.Errorf("HookModuleExceptionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HookModuleExceptionsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_hook_module_exceptions"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HookModuleExceptionsMgr) GetTableName() string {
	return "ps_hook_module_exceptions"
}

// Get 获取
func (obj *_HookModuleExceptionsMgr) Get() (result HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HookModuleExceptionsMgr) Gets() (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHookModuleExceptions id_hook_module_exceptions获取
func (obj *_HookModuleExceptionsMgr) WithIDHookModuleExceptions(idHookModuleExceptions uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook_module_exceptions"] = idHookModuleExceptions })
}

// WithIDShop id_shop获取
func (obj *_HookModuleExceptionsMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDModule id_module获取
func (obj *_HookModuleExceptionsMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDHook id_hook获取
func (obj *_HookModuleExceptionsMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

// WithFileName file_name获取
func (obj *_HookModuleExceptionsMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDHookModuleExceptions 通过id_hook_module_exceptions获取内容
func (obj *_HookModuleExceptionsMgr) GetFromIDHookModuleExceptions(idHookModuleExceptions uint32) (result HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions = ?", idHookModuleExceptions).Find(&result).Error

	return
}

// GetBatchFromIDHookModuleExceptions 批量唯一主键查找
func (obj *_HookModuleExceptionsMgr) GetBatchFromIDHookModuleExceptions(idHookModuleExceptionss []uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions IN (?)", idHookModuleExceptionss).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_HookModuleExceptionsMgr) GetFromIDShop(idShop uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_HookModuleExceptionsMgr) GetBatchFromIDShop(idShops []uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDModule 通过id_module获取内容
func (obj *_HookModuleExceptionsMgr) GetFromIDModule(idModule uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_HookModuleExceptionsMgr) GetBatchFromIDModule(idModules []uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromIDHook 通过id_hook获取内容
func (obj *_HookModuleExceptionsMgr) GetFromIDHook(idHook uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}

// GetBatchFromIDHook 批量唯一主键查找
func (obj *_HookModuleExceptionsMgr) GetBatchFromIDHook(idHooks []uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error

	return
}

// GetFromFileName 通过file_name获取内容
func (obj *_HookModuleExceptionsMgr) GetFromFileName(fileName string) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error

	return
}

// GetBatchFromFileName 批量唯一主键查找
func (obj *_HookModuleExceptionsMgr) GetBatchFromFileName(fileNames []string) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_HookModuleExceptionsMgr) FetchByPrimaryKey(idHookModuleExceptions uint32) (result HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_module_exceptions = ?", idHookModuleExceptions).Find(&result).Error

	return
}

// FetchIndexByIDModule  获取多个内容
func (obj *_HookModuleExceptionsMgr) FetchIndexByIDModule(idModule uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// FetchIndexByIDHook  获取多个内容
func (obj *_HookModuleExceptionsMgr) FetchIndexByIDHook(idHook uint32) (results []*HookModuleExceptions, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}
