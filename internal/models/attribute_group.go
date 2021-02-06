package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttributeGroupMgr struct {
	*_BaseMgr
}

// AttributeGroupMgr open func
func AttributeGroupMgr(db *gorm.DB) *_AttributeGroupMgr {
	if db == nil {
		panic(fmt.Errorf("AttributeGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttributeGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attribute_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AttributeGroupMgr) GetTableName() string {
	return "eg_attribute_group"
}

// Get 获取
func (obj *_AttributeGroupMgr) Get() (result AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AttributeGroupMgr) Gets() (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeGroup id_attribute_group获取
func (obj *_AttributeGroupMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIsColorGroup is_color_group获取
func (obj *_AttributeGroupMgr) WithIsColorGroup(isColorGroup bool) Option {
	return optionFunc(func(o *options) { o.query["is_color_group"] = isColorGroup })
}

// WithGroupType group_type获取
func (obj *_AttributeGroupMgr) WithGroupType(groupType string) Option {
	return optionFunc(func(o *options) { o.query["group_type"] = groupType })
}

// WithPosition position获取
func (obj *_AttributeGroupMgr) WithPosition(position int) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDAttributeGroup 通过id_attribute_group获取内容
func (obj *_AttributeGroupMgr) GetFromIDAttributeGroup(idAttributeGroup int) (result AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error

	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找
func (obj *_AttributeGroupMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

// GetFromIsColorGroup 通过is_color_group获取内容
func (obj *_AttributeGroupMgr) GetFromIsColorGroup(isColorGroup bool) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_color_group = ?", isColorGroup).Find(&results).Error

	return
}

// GetBatchFromIsColorGroup 批量唯一主键查找
func (obj *_AttributeGroupMgr) GetBatchFromIsColorGroup(isColorGroups []bool) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_color_group IN (?)", isColorGroups).Find(&results).Error

	return
}

// GetFromGroupType 通过group_type获取内容
func (obj *_AttributeGroupMgr) GetFromGroupType(groupType string) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_type = ?", groupType).Find(&results).Error

	return
}

// GetBatchFromGroupType 批量唯一主键查找
func (obj *_AttributeGroupMgr) GetBatchFromGroupType(groupTypes []string) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_type IN (?)", groupTypes).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_AttributeGroupMgr) GetFromPosition(position int) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_AttributeGroupMgr) GetBatchFromPosition(positions []int) (results []*AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AttributeGroupMgr) FetchByPrimaryKey(idAttributeGroup int) (result AttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error

	return
}
