package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WebservicePermissionMgr struct {
	*_BaseMgr
}

func WebservicePermissionMgr(db *gorm.DB) *_WebservicePermissionMgr {
	if db == nil {
		panic(fmt.Errorf("WebservicePermissionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WebservicePermissionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_webservice_permission"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_WebservicePermissionMgr) GetTableName() string {
	return "ps_webservice_permission"
}

func (obj *_WebservicePermissionMgr) Get() (result WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_WebservicePermissionMgr) Gets() (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_WebservicePermissionMgr) WithIDWebservicePermission(idWebservicePermission int) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_permission"] = idWebservicePermission })
}

func (obj *_WebservicePermissionMgr) WithResource(resource string) Option {
	return optionFunc(func(o *options) { o.query["resource"] = resource })
}

func (obj *_WebservicePermissionMgr) WithMethod(method string) Option {
	return optionFunc(func(o *options) { o.query["method"] = method })
}

func (obj *_WebservicePermissionMgr) WithIDWebserviceAccount(idWebserviceAccount int) Option {
	return optionFunc(func(o *options) { o.query["id_webservice_account"] = idWebserviceAccount })
}

func (obj *_WebservicePermissionMgr) GetByOption(opts ...Option) (result WebservicePermission, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_WebservicePermissionMgr) GetByOptions(opts ...Option) (results []*WebservicePermission, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_WebservicePermissionMgr) GetFromIDWebservicePermission(idWebservicePermission int) (result WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_permission = ?", idWebservicePermission).Find(&result).Error

	return
}

func (obj *_WebservicePermissionMgr) GetBatchFromIDWebservicePermission(idWebservicePermissions []int) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_permission IN (?)", idWebservicePermissions).Find(&results).Error

	return
}

func (obj *_WebservicePermissionMgr) GetFromResource(resource string) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resource = ?", resource).Find(&results).Error

	return
}

func (obj *_WebservicePermissionMgr) GetBatchFromResource(resources []string) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resource IN (?)", resources).Find(&results).Error

	return
}

func (obj *_WebservicePermissionMgr) GetFromMethod(method string) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method = ?", method).Find(&results).Error

	return
}

func (obj *_WebservicePermissionMgr) GetBatchFromMethod(methods []string) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method IN (?)", methods).Find(&results).Error

	return
}

func (obj *_WebservicePermissionMgr) GetFromIDWebserviceAccount(idWebserviceAccount int) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&results).Error

	return
}

func (obj *_WebservicePermissionMgr) GetBatchFromIDWebserviceAccount(idWebserviceAccounts []int) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account IN (?)", idWebserviceAccounts).Find(&results).Error

	return
}


func (obj *_WebservicePermissionMgr) FetchByPrimaryKey(idWebservicePermission int) (result WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_permission = ?", idWebservicePermission).Find(&result).Error

	return
}

func (obj *_WebservicePermissionMgr) FetchUniqueIndexByResource2(resource string, method string, idWebserviceAccount int) (result WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resource = ? AND method = ? AND id_webservice_account = ?", resource, method, idWebserviceAccount).Find(&result).Error

	return
}

func (obj *_WebservicePermissionMgr) FetchIndexByResource(resource string) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("resource = ?", resource).Find(&results).Error

	return
}

func (obj *_WebservicePermissionMgr) FetchIndexByMethod(method string) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method = ?", method).Find(&results).Error

	return
}

func (obj *_WebservicePermissionMgr) FetchIndexByIDWebserviceAccount(idWebserviceAccount int) (results []*WebservicePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_webservice_account = ?", idWebserviceAccount).Find(&results).Error

	return
}
