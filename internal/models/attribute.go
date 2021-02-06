package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttributeMgr struct {
	*_BaseMgr
}

// AttributeMgr open func
func AttributeMgr(db *gorm.DB) *_AttributeMgr {
	if db == nil {
		panic(fmt.Errorf("AttributeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttributeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_attribute"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AttributeMgr) GetTableName() string {
	return "ps_attribute"
}

// Get 获取
func (obj *_AttributeMgr) Get() (result Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AttributeMgr) Gets() (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttribute id_attribute获取
func (obj *_AttributeMgr) WithIDAttribute(idAttribute int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

// WithIDAttributeGroup id_attribute_group获取
func (obj *_AttributeMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithColor color获取
func (obj *_AttributeMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

// WithPosition position获取
func (obj *_AttributeMgr) WithPosition(position int) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// GetByOption 功能选项模式获取
func (obj *_AttributeMgr) GetByOption(opts ...Option) (result Attribute, err error) {
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
func (obj *_AttributeMgr) GetByOptions(opts ...Option) (results []*Attribute, err error) {
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

// GetFromIDAttribute 通过id_attribute获取内容
func (obj *_AttributeMgr) GetFromIDAttribute(idAttribute int) (result Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&result).Error

	return
}

// GetBatchFromIDAttribute 批量唯一主键查找
func (obj *_AttributeMgr) GetBatchFromIDAttribute(idAttributes []int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error

	return
}

// GetFromIDAttributeGroup 通过id_attribute_group获取内容
func (obj *_AttributeMgr) GetFromIDAttributeGroup(idAttributeGroup int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找
func (obj *_AttributeMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

// GetFromColor 通过color获取内容
func (obj *_AttributeMgr) GetFromColor(color string) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error

	return
}

// GetBatchFromColor 批量唯一主键查找
func (obj *_AttributeMgr) GetBatchFromColor(colors []string) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_AttributeMgr) GetFromPosition(position int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_AttributeMgr) GetBatchFromPosition(positions []int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AttributeMgr) FetchByPrimaryKey(idAttribute int) (result Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&result).Error

	return
}

// FetchIndexByAttributeGroup  获取多个内容
func (obj *_AttributeMgr) FetchIndexByAttributeGroup(idAttributeGroup int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}
