package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AccessMgr struct {
	*_BaseMgr
}

// AccessMgr open func
func AccessMgr(db *gorm.DB) *_AccessMgr {
	if db == nil {
		panic(fmt.Errorf("AccessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AccessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_access"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AccessMgr) GetTableName() string {
	return "ps_access"
}

// Get 获取
func (obj *_AccessMgr) Get() (result Access, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AccessMgr) Gets() (results []*Access, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProfile id_profile获取
func (obj *_AccessMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}

// WithIDAuthorizationRole id_authorization_role获取
func (obj *_AccessMgr) WithIDAuthorizationRole(idAuthorizationRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_authorization_role"] = idAuthorizationRole })
}

// GetByOption 功能选项模式获取
func (obj *_AccessMgr) GetByOption(opts ...Option) (result Access, err error) {
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
func (obj *_AccessMgr) GetByOptions(opts ...Option) (results []*Access, err error) {
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

// GetFromIDProfile 通过id_profile获取内容
func (obj *_AccessMgr) GetFromIDProfile(idProfile uint32) (results []*Access, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&results).Error

	return
}

// GetBatchFromIDProfile 批量唯一主键查找
func (obj *_AccessMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*Access, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error

	return
}

// GetFromIDAuthorizationRole 通过id_authorization_role获取内容
func (obj *_AccessMgr) GetFromIDAuthorizationRole(idAuthorizationRole uint32) (results []*Access, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role = ?", idAuthorizationRole).Find(&results).Error

	return
}

// GetBatchFromIDAuthorizationRole 批量唯一主键查找
func (obj *_AccessMgr) GetBatchFromIDAuthorizationRole(idAuthorizationRoles []uint32) (results []*Access, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_authorization_role IN (?)", idAuthorizationRoles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AccessMgr) FetchByPrimaryKey(idProfile uint32, idAuthorizationRole uint32) (result Access, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ? AND id_authorization_role = ?", idProfile, idAuthorizationRole).Find(&result).Error

	return
}
