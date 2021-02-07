package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AuthorizationRoleMgr struct {
	*_BaseMgr
}

func AuthorizationRoleMgr(db *gorm.DB) *_AuthorizationRoleMgr {
	if db == nil {
		panic(fmt.Errorf("AuthorizationRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AuthorizationRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_authorization_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AuthorizationRoleMgr) GetTableName() string {
	return "ps_authorization_role"
}

func (obj *_AuthorizationRoleMgr) Get() (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AuthorizationRoleMgr) Gets() (results []*AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_AuthorizationRoleMgr) WithIDAuthorizationRole(idAuthorizationRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_authorization_role"] = idAuthorizationRole })
}

func (obj *_AuthorizationRoleMgr) WithSlug(slug string) Option {
	return optionFunc(func(o *options) { o.query["slug"] = slug })
}

func (obj *_AuthorizationRoleMgr) GetByOption(opts ...Option) (result AuthorizationRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AuthorizationRoleMgr) GetByOptions(opts ...Option) (results []*AuthorizationRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_AuthorizationRoleMgr) GetFromIDAuthorizationRole(idAuthorizationRole uint32) (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role = ?", idAuthorizationRole).Find(&result).Error

	return
}

func (obj *_AuthorizationRoleMgr) GetBatchFromIDAuthorizationRole(idAuthorizationRoles []uint32) (results []*AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role IN (?)", idAuthorizationRoles).Find(&results).Error

	return
}

func (obj *_AuthorizationRoleMgr) GetFromSlug(slug string) (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug = ?", slug).Find(&result).Error

	return
}

func (obj *_AuthorizationRoleMgr) GetBatchFromSlug(slugs []string) (results []*AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug IN (?)", slugs).Find(&results).Error

	return
}

func (obj *_AuthorizationRoleMgr) FetchByPrimaryKey(idAuthorizationRole uint32) (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role = ?", idAuthorizationRole).Find(&result).Error

	return
}

func (obj *_AuthorizationRoleMgr) FetchUniqueBySlug(slug string) (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug = ?", slug).Find(&result).Error

	return
}
