package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _MetaMgr struct {
	*_BaseMgr
}

// MetaMgr open func
func MetaMgr(db *gorm.DB) *_MetaMgr {
	if db == nil {
		panic(fmt.Errorf("MetaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MetaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_meta"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MetaMgr) GetTableName() string {
	return "ps_meta"
}

// Get 获取
func (obj *_MetaMgr) Get() (result Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MetaMgr) Gets() (results []*Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMeta id_meta获取
func (obj *_MetaMgr) WithIDMeta(idMeta uint32) Option {
	return optionFunc(func(o *options) { o.query["id_meta"] = idMeta })
}

// WithPage page获取
func (obj *_MetaMgr) WithPage(page string) Option {
	return optionFunc(func(o *options) { o.query["page"] = page })
}

// WithConfigurable configurable获取
func (obj *_MetaMgr) WithConfigurable(configurable bool) Option {
	return optionFunc(func(o *options) { o.query["configurable"] = configurable })
}

// GetByOption 功能选项模式获取
func (obj *_MetaMgr) GetByOption(opts ...Option) (result Meta, err error) {
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
func (obj *_MetaMgr) GetByOptions(opts ...Option) (results []*Meta, err error) {
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

// GetFromIDMeta 通过id_meta获取内容
func (obj *_MetaMgr) GetFromIDMeta(idMeta uint32) (result Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta = ?", idMeta).Find(&result).Error

	return
}

// GetBatchFromIDMeta 批量唯一主键查找
func (obj *_MetaMgr) GetBatchFromIDMeta(idMetas []uint32) (results []*Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta IN (?)", idMetas).Find(&results).Error

	return
}

// GetFromPage 通过page获取内容
func (obj *_MetaMgr) GetFromPage(page string) (result Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("page = ?", page).Find(&result).Error

	return
}

// GetBatchFromPage 批量唯一主键查找
func (obj *_MetaMgr) GetBatchFromPage(pages []string) (results []*Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("page IN (?)", pages).Find(&results).Error

	return
}

// GetFromConfigurable 通过configurable获取内容
func (obj *_MetaMgr) GetFromConfigurable(configurable bool) (results []*Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("configurable = ?", configurable).Find(&results).Error

	return
}

// GetBatchFromConfigurable 批量唯一主键查找
func (obj *_MetaMgr) GetBatchFromConfigurable(configurables []bool) (results []*Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("configurable IN (?)", configurables).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MetaMgr) FetchByPrimaryKey(idMeta uint32) (result Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta = ?", idMeta).Find(&result).Error

	return
}

// FetchUniqueByPage primay or index 获取唯一内容
func (obj *_MetaMgr) FetchUniqueByPage(page string) (result Meta, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("page = ?", page).Find(&result).Error

	return
}
