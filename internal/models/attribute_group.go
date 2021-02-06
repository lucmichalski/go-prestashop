package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttributeGroupMgr struct {
	*_BaseMgr
}

func AttributeGroupMgr(db *gorm.DB) *_AttributeGroupMgr {
	if db == nil {
		panic(fmt.Errorf("AttributeGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttributeGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_attribute_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AttributeGroupMgr) GetTableName() string {
	return "ps_attribute_group"
}

func (obj *_AttributeGroupMgr) Get() (result AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AttributeGroupMgr) Gets() (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_AttributeGroupMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

func (obj *_AttributeGroupMgr) WithIsColorGroup(isColorGroup bool) Option {
	return optionFunc(func(o *options) { o.query["is_color_group"] = isColorGroup })
}

func (obj *_AttributeGroupMgr) WithGroupType(groupType string) Option {
	return optionFunc(func(o *options) { o.query["group_type"] = groupType })
}

func (obj *_AttributeGroupMgr) WithPosition(position int) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_AttributeGroupMgr) GetByOption(opts ...Option) (result AttributeGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AttributeGroupMgr) GetByOptions(opts ...Option) (results []*AttributeGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_AttributeGroupMgr) GetFromIDAttributeGroup(idAttributeGroup int) (result AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error

	return
}

func (obj *_AttributeGroupMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

func (obj *_AttributeGroupMgr) GetFromIsColorGroup(isColorGroup bool) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_color_group = ?", isColorGroup).Find(&results).Error

	return
}

func (obj *_AttributeGroupMgr) GetBatchFromIsColorGroup(isColorGroups []bool) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_color_group IN (?)", isColorGroups).Find(&results).Error

	return
}

func (obj *_AttributeGroupMgr) GetFromGroupType(groupType string) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_type = ?", groupType).Find(&results).Error

	return
}

func (obj *_AttributeGroupMgr) GetBatchFromGroupType(groupTypes []string) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_type IN (?)", groupTypes).Find(&results).Error

	return
}

func (obj *_AttributeGroupMgr) GetFromPosition(position int) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_AttributeGroupMgr) GetBatchFromPosition(positions []int) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}


func (obj *_AttributeGroupMgr) FetchByPrimaryKey(idAttributeGroup int) (result AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error

	return
}
