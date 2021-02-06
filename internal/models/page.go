package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PageMgr struct {
	*_BaseMgr
}

// PageMgr open func
func PageMgr(db *gorm.DB) *_PageMgr {
	if db == nil {
		panic(fmt.Errorf("PageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_page"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PageMgr) GetTableName() string {
	return "ps_page"
}

// Get 获取
func (obj *_PageMgr) Get() (result Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PageMgr) Gets() (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPage id_page获取
func (obj *_PageMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

// WithIDPageType id_page_type获取
func (obj *_PageMgr) WithIDPageType(idPageType uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page_type"] = idPageType })
}

// WithIDObject id_object获取
func (obj *_PageMgr) WithIDObject(idObject uint32) Option {
	return optionFunc(func(o *options) { o.query["id_object"] = idObject })
}

// GetByOption 功能选项模式获取
func (obj *_PageMgr) GetByOption(opts ...Option) (result Page, err error) {
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
func (obj *_PageMgr) GetByOptions(opts ...Option) (results []*Page, err error) {
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

// GetFromIDPage 通过id_page获取内容
func (obj *_PageMgr) GetFromIDPage(idPage uint32) (result Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&result).Error

	return
}

// GetBatchFromIDPage 批量唯一主键查找
func (obj *_PageMgr) GetBatchFromIDPage(idPages []uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error

	return
}

// GetFromIDPageType 通过id_page_type获取内容
func (obj *_PageMgr) GetFromIDPageType(idPageType uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&results).Error

	return
}

// GetBatchFromIDPageType 批量唯一主键查找
func (obj *_PageMgr) GetBatchFromIDPageType(idPageTypes []uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type IN (?)", idPageTypes).Find(&results).Error

	return
}

// GetFromIDObject 通过id_object获取内容
func (obj *_PageMgr) GetFromIDObject(idObject uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object = ?", idObject).Find(&results).Error

	return
}

// GetBatchFromIDObject 批量唯一主键查找
func (obj *_PageMgr) GetBatchFromIDObject(idObjects []uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object IN (?)", idObjects).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PageMgr) FetchByPrimaryKey(idPage uint32) (result Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&result).Error

	return
}

// FetchIndexByIDPageType  获取多个内容
func (obj *_PageMgr) FetchIndexByIDPageType(idPageType uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&results).Error

	return
}

// FetchIndexByIDObject  获取多个内容
func (obj *_PageMgr) FetchIndexByIDObject(idObject uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object = ?", idObject).Find(&results).Error

	return
}
