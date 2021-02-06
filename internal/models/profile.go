package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProfileMgr struct {
	*_BaseMgr
}

// ProfileMgr open func
func ProfileMgr(db *gorm.DB) *_ProfileMgr {
	if db == nil {
		panic(fmt.Errorf("ProfileMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProfileMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_profile"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProfileMgr) GetTableName() string {
	return "ps_profile"
}

// Get 获取
func (obj *_ProfileMgr) Get() (result Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProfileMgr) Gets() (results []*Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProfile id_profile获取
func (obj *_ProfileMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}

// GetByOption 功能选项模式获取
func (obj *_ProfileMgr) GetByOption(opts ...Option) (result Profile, err error) {
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
func (obj *_ProfileMgr) GetByOptions(opts ...Option) (results []*Profile, err error) {
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
func (obj *_ProfileMgr) GetFromIDProfile(idProfile uint32) (result Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&result).Error

	return
}

// GetBatchFromIDProfile 批量唯一主键查找
func (obj *_ProfileMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProfileMgr) FetchByPrimaryKey(idProfile uint32) (result Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&result).Error

	return
}
