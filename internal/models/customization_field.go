package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomizationFieldMgr struct {
	*_BaseMgr
}

func CustomizationFieldMgr(db *gorm.DB) *_CustomizationFieldMgr {
	if db == nil {
		panic(fmt.Errorf("CustomizationFieldMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomizationFieldMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customization_field"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomizationFieldMgr) GetTableName() string {
	return "ps_customization_field"
}

func (obj *_CustomizationFieldMgr) Get() (result CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomizationFieldMgr) Gets() (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CustomizationFieldMgr) WithIDCustomizationField(idCustomizationField uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization_field"] = idCustomizationField })
}

func (obj *_CustomizationFieldMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_CustomizationFieldMgr) WithType(_type bool) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

func (obj *_CustomizationFieldMgr) WithRequired(required bool) Option {
	return optionFunc(func(o *options) { o.query["required"] = required })
}

func (obj *_CustomizationFieldMgr) WithIsModule(isModule bool) Option {
	return optionFunc(func(o *options) { o.query["is_module"] = isModule })
}

func (obj *_CustomizationFieldMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

func (obj *_CustomizationFieldMgr) GetByOption(opts ...Option) (result CustomizationField, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CustomizationFieldMgr) GetByOptions(opts ...Option) (results []*CustomizationField, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CustomizationFieldMgr) GetFromIDCustomizationField(idCustomizationField uint32) (result CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field = ?", idCustomizationField).Find(&result).Error

	return
}

func (obj *_CustomizationFieldMgr) GetBatchFromIDCustomizationField(idCustomizationFields []uint32) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field IN (?)", idCustomizationFields).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetFromIDProduct(idProduct uint32) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetFromType(_type bool) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetBatchFromType(_types []bool) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetFromRequired(required bool) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("required = ?", required).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetBatchFromRequired(requireds []bool) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("required IN (?)", requireds).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetFromIsModule(isModule bool) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module = ?", isModule).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetBatchFromIsModule(isModules []bool) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module IN (?)", isModules).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetFromIsDeleted(isDeleted bool) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_deleted = ?", isDeleted).Find(&results).Error

	return
}

func (obj *_CustomizationFieldMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_deleted IN (?)", isDeleteds).Find(&results).Error

	return
}


func (obj *_CustomizationFieldMgr) FetchByPrimaryKey(idCustomizationField uint32) (result CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field = ?", idCustomizationField).Find(&result).Error

	return
}

func (obj *_CustomizationFieldMgr) FetchIndexByIDProduct(idProduct uint32) (results []*CustomizationField, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}
