package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttributeMgr struct {
	*_BaseMgr
}

func AttributeMgr(db *gorm.DB) *_AttributeMgr {
	if db == nil {
		panic(fmt.Errorf("AttributeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttributeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_attribute"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AttributeMgr) GetTableName() string {
	return "ps_attribute"
}

func (obj *_AttributeMgr) Get() (result Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AttributeMgr) Gets() (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_AttributeMgr) WithIDAttribute(idAttribute int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

func (obj *_AttributeMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

func (obj *_AttributeMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

func (obj *_AttributeMgr) WithPosition(position int) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

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

func (obj *_AttributeMgr) GetFromIDAttribute(idAttribute int) (result Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&result).Error

	return
}

func (obj *_AttributeMgr) GetBatchFromIDAttribute(idAttributes []int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error

	return
}

func (obj *_AttributeMgr) GetFromIDAttributeGroup(idAttributeGroup int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

func (obj *_AttributeMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

func (obj *_AttributeMgr) GetFromColor(color string) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error

	return
}

func (obj *_AttributeMgr) GetBatchFromColor(colors []string) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error

	return
}

func (obj *_AttributeMgr) GetFromPosition(position int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_AttributeMgr) GetBatchFromPosition(positions []int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_AttributeMgr) FetchByPrimaryKey(idAttribute int) (result Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&result).Error

	return
}

func (obj *_AttributeMgr) FetchIndexByAttributeGroup(idAttributeGroup int) (results []*Attribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}
