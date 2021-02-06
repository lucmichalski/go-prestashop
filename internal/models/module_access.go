package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleAccessMgr struct {
	*_BaseMgr
}

func ModuleAccessMgr(db *gorm.DB) *_ModuleAccessMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleAccessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleAccessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_access"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ModuleAccessMgr) GetTableName() string {
	return "ps_module_access"
}

func (obj *_ModuleAccessMgr) Get() (result ModuleAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ModuleAccessMgr) Gets() (results []*ModuleAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ModuleAccessMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}

func (obj *_ModuleAccessMgr) WithIDAuthorizationRole(idAuthorizationRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_authorization_role"] = idAuthorizationRole })
}

func (obj *_ModuleAccessMgr) GetByOption(opts ...Option) (result ModuleAccess, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ModuleAccessMgr) GetByOptions(opts ...Option) (results []*ModuleAccess, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ModuleAccessMgr) GetFromIDProfile(idProfile uint32) (results []*ModuleAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&results).Error

	return
}

func (obj *_ModuleAccessMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*ModuleAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error

	return
}

func (obj *_ModuleAccessMgr) GetFromIDAuthorizationRole(idAuthorizationRole uint32) (results []*ModuleAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role = ?", idAuthorizationRole).Find(&results).Error

	return
}

func (obj *_ModuleAccessMgr) GetBatchFromIDAuthorizationRole(idAuthorizationRoles []uint32) (results []*ModuleAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role IN (?)", idAuthorizationRoles).Find(&results).Error

	return
}


func (obj *_ModuleAccessMgr) FetchByPrimaryKey(idProfile uint32, idAuthorizationRole uint32) (result ModuleAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ? AND id_authorization_role = ?", idProfile, idAuthorizationRole).Find(&result).Error

	return
}
