package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredIndexableAttributeGroupMgr struct {
	*_BaseMgr
}

func LayeredIndexableAttributeGroupMgr(db *gorm.DB) *_LayeredIndexableAttributeGroupMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredIndexableAttributeGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredIndexableAttributeGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_indexable_attribute_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredIndexableAttributeGroupMgr) GetTableName() string {
	return "ps_layered_indexable_attribute_group"
}

func (obj *_LayeredIndexableAttributeGroupMgr) Get() (result LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupMgr) Gets() (results []*LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LayeredIndexableAttributeGroupMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

func (obj *_LayeredIndexableAttributeGroupMgr) WithIndexable(indexable bool) Option {
	return optionFunc(func(o *options) { o.query["indexable"] = indexable })
}

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


func (obj *_LayeredIndexableAttributeGroupMgr) GetFromIDAttributeGroup(idAttributeGroup int) (result LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupMgr) GetFromIndexable(indexable bool) (results []*LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable = ?", indexable).Find(&results).Error

	return
}

func (obj *_LayeredIndexableAttributeGroupMgr) GetBatchFromIndexable(indexables []bool) (results []*LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable IN (?)", indexables).Find(&results).Error

	return
}


func (obj *_LayeredIndexableAttributeGroupMgr) FetchByPrimaryKey(idAttributeGroup int) (result LayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error

	return
}
