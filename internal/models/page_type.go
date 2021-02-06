package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PageTypeMgr struct {
	*_BaseMgr
}

func PageTypeMgr(db *gorm.DB) *_PageTypeMgr {
	if db == nil {
		panic(fmt.Errorf("PageTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PageTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_page_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PageTypeMgr) GetTableName() string {
	return "ps_page_type"
}

func (obj *_PageTypeMgr) Get() (result PageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PageTypeMgr) Gets() (results []*PageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PageTypeMgr) WithIDPageType(idPageType uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page_type"] = idPageType })
}

func (obj *_PageTypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_PageTypeMgr) GetByOption(opts ...Option) (result PageType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_PageTypeMgr) GetByOptions(opts ...Option) (results []*PageType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_PageTypeMgr) GetFromIDPageType(idPageType uint32) (result PageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&result).Error

	return
}

func (obj *_PageTypeMgr) GetBatchFromIDPageType(idPageTypes []uint32) (results []*PageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type IN (?)", idPageTypes).Find(&results).Error

	return
}

func (obj *_PageTypeMgr) GetFromName(name string) (results []*PageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_PageTypeMgr) GetBatchFromName(names []string) (results []*PageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_PageTypeMgr) FetchByPrimaryKey(idPageType uint32) (result PageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page_type = ?", idPageType).Find(&result).Error

	return
}

func (obj *_PageTypeMgr) FetchIndexByName(name string) (results []*PageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
