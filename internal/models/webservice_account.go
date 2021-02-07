package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WebserviceAccountMgr struct {
	*_BaseMgr
}

func WebserviceAccountMgr(db *gorm.DB) *_WebserviceAccountMgr {
	if db == nil {
		panic(fmt.Errorf("WebserviceAccountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WebserviceAccountMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_webservice_account"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_WebserviceAccountMgr) GetTableName() string {
	return "ps_webservice_account"
}

func (obj *_WebserviceAccountMgr) Get() (result WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_WebserviceAccountMgr) Gets() (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) WithIDWebserviceAccount(idWebserviceAccount int) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_account"] = idWebserviceAccount })
}

func (obj *_WebserviceAccountMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

func (obj *_WebserviceAccountMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_WebserviceAccountMgr) WithClassName(className string) Option {
	return optionFunc(func(o *options) { o.query["class_name"] = className })
}

func (obj *_WebserviceAccountMgr) WithIsModule(isModule int8) Option {
	return optionFunc(func(o *options) { o.query["is_module"] = isModule })
}

func (obj *_WebserviceAccountMgr) WithModuleName(moduleName string) Option {
	return optionFunc(func(o *options) { o.query["module_name"] = moduleName })
}

func (obj *_WebserviceAccountMgr) WithActive(active int8) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

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

func (obj *_WebserviceAccountMgr) GetFromIDWebserviceAccount(idWebserviceAccount int) (result WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&result).Error

	return
}

func (obj *_WebserviceAccountMgr) GetBatchFromIDWebserviceAccount(idWebserviceAccounts []int) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account IN (?)", idWebserviceAccounts).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetFromKey(key string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key = ?", key).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetBatchFromKey(keys []string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key IN (?)", keys).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetFromDescription(description string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetBatchFromDescription(descriptions []string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetFromClassName(className string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name = ?", className).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetBatchFromClassName(classNames []string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name IN (?)", classNames).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetFromIsModule(isModule int8) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module = ?", isModule).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetBatchFromIsModule(isModules []int8) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module IN (?)", isModules).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetFromModuleName(moduleName string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name = ?", moduleName).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetBatchFromModuleName(moduleNames []string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module_name IN (?)", moduleNames).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetFromActive(active int8) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) GetBatchFromActive(actives []int8) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_WebserviceAccountMgr) FetchByPrimaryKey(idWebserviceAccount int) (result WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&result).Error

	return
}

func (obj *_WebserviceAccountMgr) FetchIndexByKey(key string) (results []*WebserviceAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key = ?", key).Find(&results).Error

	return
}
