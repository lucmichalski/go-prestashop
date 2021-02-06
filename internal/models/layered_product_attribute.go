package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredProductAttributeMgr struct {
	*_BaseMgr
}

func LayeredProductAttributeMgr(db *gorm.DB) *_LayeredProductAttributeMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredProductAttributeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredProductAttributeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_product_attribute"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LayeredProductAttributeMgr) GetTableName() string {
	return "ps_layered_product_attribute"
}

func (obj *_LayeredProductAttributeMgr) Get() (result LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LayeredProductAttributeMgr) Gets() (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LayeredProductAttributeMgr) WithIDAttribute(idAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

func (obj *_LayeredProductAttributeMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_LayeredProductAttributeMgr) WithIDAttributeGroup(idAttributeGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

func (obj *_LayeredProductAttributeMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_LayeredProductAttributeMgr) GetByOption(opts ...Option) (result LayeredProductAttribute, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LayeredProductAttributeMgr) GetByOptions(opts ...Option) (results []*LayeredProductAttribute, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LayeredProductAttributeMgr) GetFromIDAttribute(idAttribute uint32) (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error

	return
}

func (obj *_LayeredProductAttributeMgr) GetBatchFromIDAttribute(idAttributes []uint32) (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error

	return
}

func (obj *_LayeredProductAttributeMgr) GetFromIDProduct(idProduct uint32) (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_LayeredProductAttributeMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_LayeredProductAttributeMgr) GetFromIDAttributeGroup(idAttributeGroup uint32) (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

func (obj *_LayeredProductAttributeMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []uint32) (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

func (obj *_LayeredProductAttributeMgr) GetFromIDShop(idShop uint32) (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_LayeredProductAttributeMgr) GetBatchFromIDShop(idShops []uint32) (results []*LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}


func (obj *_LayeredProductAttributeMgr) FetchByPrimaryKey(idAttribute uint32, idProduct uint32, idShop uint32) (result LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_product = ? AND id_shop = ?", idAttribute, idProduct, idShop).Find(&result).Error

	return
}

func (obj *_LayeredProductAttributeMgr) FetchUniqueIndexByIDAttributeGroup(idAttribute uint32, idProduct uint32, idAttributeGroup uint32, idShop uint32) (result LayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_product = ? AND id_attribute_group = ? AND id_shop = ?", idAttribute, idProduct, idAttributeGroup, idShop).Find(&result).Error

	return
}
