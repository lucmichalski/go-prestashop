package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _RequiredFieldMgr struct {
	*_BaseMgr
}

func RequiredFieldMgr(db *gorm.DB) *_RequiredFieldMgr {
	if db == nil {
		panic(fmt.Errorf("RequiredFieldMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RequiredFieldMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_required_field"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_RequiredFieldMgr) GetTableName() string {
	return "ps_required_field"
}

func (obj *_RequiredFieldMgr) Get() (result RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_RequiredFieldMgr) Gets() (results []*RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_RequiredFieldMgr) WithIDRequiredField(idRequiredField int) Option {
	return optionFunc(func(o *options) { o.query["id_required_field"] = idRequiredField })
}

func (obj *_RequiredFieldMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

func (obj *_RequiredFieldMgr) WithFieldName(fieldName string) Option {
	return optionFunc(func(o *options) { o.query["field_name"] = fieldName })
}

func (obj *_RequiredFieldMgr) GetByOption(opts ...Option) (result RequiredField, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_RequiredFieldMgr) GetByOptions(opts ...Option) (results []*RequiredField, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_RequiredFieldMgr) GetFromIDRequiredField(idRequiredField int) (result RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_required_field = ?", idRequiredField).Find(&result).Error

	return
}

func (obj *_RequiredFieldMgr) GetBatchFromIDRequiredField(idRequiredFields []int) (results []*RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_required_field IN (?)", idRequiredFields).Find(&results).Error

	return
}

func (obj *_RequiredFieldMgr) GetFromObjectName(objectName string) (results []*RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_name = ?", objectName).Find(&results).Error

	return
}

func (obj *_RequiredFieldMgr) GetBatchFromObjectName(objectNames []string) (results []*RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_name IN (?)", objectNames).Find(&results).Error

	return
}

func (obj *_RequiredFieldMgr) GetFromFieldName(fieldName string) (results []*RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("field_name = ?", fieldName).Find(&results).Error

	return
}

func (obj *_RequiredFieldMgr) GetBatchFromFieldName(fieldNames []string) (results []*RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("field_name IN (?)", fieldNames).Find(&results).Error

	return
}


func (obj *_RequiredFieldMgr) FetchByPrimaryKey(idRequiredField int) (result RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_required_field = ?", idRequiredField).Find(&result).Error

	return
}

func (obj *_RequiredFieldMgr) FetchIndexByObjectName(objectName string) (results []*RequiredField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_name = ?", objectName).Find(&results).Error

	return
}
