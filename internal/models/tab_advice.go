package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TabAdviceMgr struct {
	*_BaseMgr
}

// TabAdviceMgr open func
func TabAdviceMgr(db *gorm.DB) *_TabAdviceMgr {
	if db == nil {
		panic(fmt.Errorf("TabAdviceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TabAdviceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tab_advice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TabAdviceMgr) GetTableName() string {
	return "ps_tab_advice"
}

// Get 获取
func (obj *_TabAdviceMgr) Get() (result TabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TabAdviceMgr) Gets() (results []*TabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTab id_tab获取
func (obj *_TabAdviceMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

// WithIDAdvice id_advice获取
func (obj *_TabAdviceMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}

// GetByOption 功能选项模式获取
func (obj *_TabAdviceMgr) GetByOption(opts ...Option) (result TabAdvice, err error) {
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
func (obj *_TabAdviceMgr) GetByOptions(opts ...Option) (results []*TabAdvice, err error) {
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

// GetFromIDTab 通过id_tab获取内容
func (obj *_TabAdviceMgr) GetFromIDTab(idTab int) (results []*TabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

// GetBatchFromIDTab 批量唯一主键查找
func (obj *_TabAdviceMgr) GetBatchFromIDTab(idTabs []int) (results []*TabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

// GetFromIDAdvice 通过id_advice获取内容
func (obj *_TabAdviceMgr) GetFromIDAdvice(idAdvice int) (results []*TabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&results).Error

	return
}

// GetBatchFromIDAdvice 批量唯一主键查找
func (obj *_TabAdviceMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*TabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TabAdviceMgr) FetchByPrimaryKey(idTab int, idAdvice int) (result TabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ? AND id_advice = ?", idTab, idAdvice).Find(&result).Error

	return
}
