package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableAttributeGroupMgr struct {
	*_BaseMgr
}

// LayeredIndexableAttributeGroupMgr open func
func LayeredIndexableAttributeGroupMgr(db *gorm.DB) *_LayeredIndexableAttributeGroupMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableAttributeGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableAttributeGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_indexable_attribute_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredIndexableAttributeGroupMgr) GetTableName() string {
	return "eg_layered_indexable_attribute_group"
}

// Get 获取
func (obj *_LayeredIndexableAttributeGroupMgr) Get() (result LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredIndexableAttributeGroupMgr) Gets() (results []*LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeGroup id_attribute_group获取
func (obj *_LayeredIndexableAttributeGroupMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIndexable indexable获取
func (obj *_LayeredIndexableAttributeGroupMgr) WithIndexable(indexable bool) Option {
	return optionFunc(func(o *options) { o.query["indexable"] = indexable })
}

// GetByOption 功能选项模式获取
func (obj *_LayeredIndexableAttributeGroupMgr) GetByOption(opts ...Option) (result LayeredIndexableAttributeGroup, err error) {
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
func (obj *_LayeredIndexableAttributeGroupMgr) GetByOptions(opts ...Option) (results []*LayeredIndexableAttributeGroup, err error) {
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

// GetFromIDAttributeGroup 通过id_attribute_group获取内容
func (obj *_LayeredIndexableAttributeGroupMgr) GetFromIDAttributeGroup(idAttributeGroup int) (result LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error

	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找
func (obj *_LayeredIndexableAttributeGroupMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

// GetFromIndexable 通过indexable获取内容
func (obj *_LayeredIndexableAttributeGroupMgr) GetFromIndexable(indexable bool) (results []*LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable = ?", indexable).Find(&results).Error

	return
}

// GetBatchFromIndexable 批量唯一主键查找
func (obj *_LayeredIndexableAttributeGroupMgr) GetBatchFromIndexable(indexables []bool) (results []*LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable IN (?)", indexables).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredIndexableAttributeGroupMgr) FetchByPrimaryKey(idAttributeGroup int) (result LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error

	return
}
