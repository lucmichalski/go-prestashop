package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WebserviceAccountMgr struct {
	*_BaseMgr
}

// WebserviceAccountMgr open func
func WebserviceAccountMgr(db *gorm.DB) *_WebserviceAccountMgr {
	if db == nil {
		panic(fmt.Errorf("WebserviceAccountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WebserviceAccountMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_webservice_account"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WebserviceAccountMgr) GetTableName() string {
	return "eg_webservice_account"
}

// Get 获取
func (obj *_WebserviceAccountMgr) Get() (result WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WebserviceAccountMgr) Gets() (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWebserviceAccount id_webservice_account获取
func (obj *_WebserviceAccountMgr) WithIDWebserviceAccount(idWebserviceAccount int) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_account"] = idWebserviceAccount })
}

// WithKey key获取
func (obj *_WebserviceAccountMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithDescription description获取
func (obj *_WebserviceAccountMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithClassName class_name获取
func (obj *_WebserviceAccountMgr) WithClassName(className string) Option {
	return optionFunc(func(o *options) { o.query["class_name"] = className })
}

// WithIsModule is_module获取
func (obj *_WebserviceAccountMgr) WithIsModule(isModule int8) Option {
	return optionFunc(func(o *options) { o.query["is_module"] = isModule })
}

// WithModuleName module_name获取
func (obj *_WebserviceAccountMgr) WithModuleName(moduleName string) Option {
	return optionFunc(func(o *options) { o.query["module_name"] = moduleName })
}

// WithActive active获取
func (obj *_WebserviceAccountMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// GetByOption 功能选项模式获取
func (obj *_WebserviceAccountMgr) GetByOption(opts ...Option) (result WebserviceAccount, err error) {
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
func (obj *_WebserviceAccountMgr) GetByOptions(opts ...Option) (results []*WebserviceAccount, err error) {
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

// GetFromIDWebserviceAccount 通过id_webservice_account获取内容
func (obj *_WebserviceAccountMgr) GetFromIDWebserviceAccount(idWebserviceAccount int) (result WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&result).Error

	return
}

// GetBatchFromIDWebserviceAccount 批量唯一主键查找
func (obj *_WebserviceAccountMgr) GetBatchFromIDWebserviceAccount(idWebserviceAccounts []int) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account IN (?)", idWebserviceAccounts).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容
func (obj *_WebserviceAccountMgr) GetFromKey(key string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key = ?", key).Find(&results).Error

	return
}

// GetBatchFromKey 批量唯一主键查找
func (obj *_WebserviceAccountMgr) GetBatchFromKey(keys []string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key IN (?)", keys).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_WebserviceAccountMgr) GetFromDescription(description string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_WebserviceAccountMgr) GetBatchFromDescription(descriptions []string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromClassName 通过class_name获取内容
func (obj *_WebserviceAccountMgr) GetFromClassName(className string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name = ?", className).Find(&results).Error

	return
}

// GetBatchFromClassName 批量唯一主键查找
func (obj *_WebserviceAccountMgr) GetBatchFromClassName(classNames []string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name IN (?)", classNames).Find(&results).Error

	return
}

// GetFromIsModule 通过is_module获取内容
func (obj *_WebserviceAccountMgr) GetFromIsModule(isModule int8) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module = ?", isModule).Find(&results).Error

	return
}

// GetBatchFromIsModule 批量唯一主键查找
func (obj *_WebserviceAccountMgr) GetBatchFromIsModule(isModules []int8) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module IN (?)", isModules).Find(&results).Error

	return
}

// GetFromModuleName 通过module_name获取内容
func (obj *_WebserviceAccountMgr) GetFromModuleName(moduleName string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name = ?", moduleName).Find(&results).Error

	return
}

// GetBatchFromModuleName 批量唯一主键查找
func (obj *_WebserviceAccountMgr) GetBatchFromModuleName(moduleNames []string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name IN (?)", moduleNames).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_WebserviceAccountMgr) GetFromActive(active int8) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_WebserviceAccountMgr) GetBatchFromActive(actives []int8) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_WebserviceAccountMgr) FetchByPrimaryKey(idWebserviceAccount int) (result WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&result).Error

	return
}

// FetchIndexByKey  获取多个内容
func (obj *_WebserviceAccountMgr) FetchIndexByKey(key string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key = ?", key).Find(&results).Error

	return
}
