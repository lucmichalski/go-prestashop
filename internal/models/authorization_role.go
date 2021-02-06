package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AuthorizationRoleMgr struct {
	*_BaseMgr
}

// AuthorizationRoleMgr open func
func AuthorizationRoleMgr(db *gorm.DB) *_AuthorizationRoleMgr {
	if db == nil {
		panic(fmt.Errorf("AuthorizationRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AuthorizationRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_authorization_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AuthorizationRoleMgr) GetTableName() string {
	return "eg_authorization_role"
}

// Get 获取
func (obj *_AuthorizationRoleMgr) Get() (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AuthorizationRoleMgr) Gets() (results []*AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAuthorizationRole id_authorization_role获取
func (obj *_AuthorizationRoleMgr) WithIDAuthorizationRole(idAuthorizationRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_authorization_role"] = idAuthorizationRole })
}

// WithSlug slug获取
func (obj *_AuthorizationRoleMgr) WithSlug(slug string) Option {
	return optionFunc(func(o *options) { o.query["slug"] = slug })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDAuthorizationRole 通过id_authorization_role获取内容
func (obj *_AuthorizationRoleMgr) GetFromIDAuthorizationRole(idAuthorizationRole uint32) (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role = ?", idAuthorizationRole).Find(&result).Error

	return
}

// GetBatchFromIDAuthorizationRole 批量唯一主键查找
func (obj *_AuthorizationRoleMgr) GetBatchFromIDAuthorizationRole(idAuthorizationRoles []uint32) (results []*AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role IN (?)", idAuthorizationRoles).Find(&results).Error

	return
}

// GetFromSlug 通过slug获取内容
func (obj *_AuthorizationRoleMgr) GetFromSlug(slug string) (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug = ?", slug).Find(&result).Error

	return
}

// GetBatchFromSlug 批量唯一主键查找
func (obj *_AuthorizationRoleMgr) GetBatchFromSlug(slugs []string) (results []*AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug IN (?)", slugs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AuthorizationRoleMgr) FetchByPrimaryKey(idAuthorizationRole uint32) (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role = ?", idAuthorizationRole).Find(&result).Error

	return
}

// FetchUniqueBySlug primay or index 获取唯一内容
func (obj *_AuthorizationRoleMgr) FetchUniqueBySlug(slug string) (result AuthorizationRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("slug = ?", slug).Find(&result).Error

	return
}
