package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _InfoMgr struct {
	*_BaseMgr
}

// InfoMgr open func
func InfoMgr(db *gorm.DB) *_InfoMgr {
	if db == nil {
		panic(fmt.Errorf("InfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_InfoMgr) GetTableName() string {
	return "eg_info"
}

// Get 获取
func (obj *_InfoMgr) Get() (result Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_InfoMgr) Gets() (results []*Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDInfo id_info获取
func (obj *_InfoMgr) WithIDInfo(idInfo uint32) Option {
	return optionFunc(func(o *options) { o.query["id_info"] = idInfo })
}

// GetByOption 功能选项模式获取
func (obj *_InfoMgr) GetByOption(opts ...Option) (result Info, err error) {
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
func (obj *_InfoMgr) GetByOptions(opts ...Option) (results []*Info, err error) {
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

// GetFromIDInfo 通过id_info获取内容
func (obj *_InfoMgr) GetFromIDInfo(idInfo uint32) (result Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&result).Error

	return
}

// GetBatchFromIDInfo 批量唯一主键查找
func (obj *_InfoMgr) GetBatchFromIDInfo(idInfos []uint32) (results []*Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info IN (?)", idInfos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_InfoMgr) FetchByPrimaryKey(idInfo uint32) (result Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&result).Error

	return
}
