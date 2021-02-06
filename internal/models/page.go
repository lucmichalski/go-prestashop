package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PageMgr struct {
	*_BaseMgr
}

func PageMgr(db *gorm.DB) *_PageMgr {
	if db == nil {
		panic(fmt.Errorf("PageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_page"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PageMgr) GetTableName() string {
	return "ps_page"
}

func (obj *_PageMgr) Get() (result Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PageMgr) Gets() (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PageMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

func (obj *_PageMgr) WithIDPageType(idPageType uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page_type"] = idPageType })
}

func (obj *_PageMgr) WithIDObject(idObject uint32) Option {
	return optionFunc(func(o *options) { o.query["id_object"] = idObject })
}

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


func (obj *_PageMgr) GetFromIDPage(idPage uint32) (result Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&result).Error

	return
}

func (obj *_PageMgr) GetBatchFromIDPage(idPages []uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error

	return
}

func (obj *_PageMgr) GetFromIDPageType(idPageType uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&results).Error

	return
}

func (obj *_PageMgr) GetBatchFromIDPageType(idPageTypes []uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type IN (?)", idPageTypes).Find(&results).Error

	return
}

func (obj *_PageMgr) GetFromIDObject(idObject uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object = ?", idObject).Find(&results).Error

	return
}

func (obj *_PageMgr) GetBatchFromIDObject(idObjects []uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object IN (?)", idObjects).Find(&results).Error

	return
}


func (obj *_PageMgr) FetchByPrimaryKey(idPage uint32) (result Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&result).Error

	return
}

func (obj *_PageMgr) FetchIndexByIDPageType(idPageType uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&results).Error

	return
}

func (obj *_PageMgr) FetchIndexByIDObject(idObject uint32) (results []*Page, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_object = ?", idObject).Find(&results).Error

	return
}
