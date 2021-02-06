package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductAttributeImageMgr struct {
	*_BaseMgr
}

func ProductAttributeImageMgr(db *gorm.DB) *_ProductAttributeImageMgr {
	if db == nil {
		panic(fmt.Errorf("ProductAttributeImageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductAttributeImageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_attribute_image"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductAttributeImageMgr) GetTableName() string {
	return "ps_product_attribute_image"
}

func (obj *_ProductAttributeImageMgr) Get() (result ProductAttributeImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductAttributeImageMgr) Gets() (results []*ProductAttributeImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductAttributeImageMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

func (obj *_ProductAttributeImageMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

func (obj *_ProductAttributeImageMgr) GetByOption(opts ...Option) (result ProductAttributeImage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductAttributeImageMgr) GetByOptions(opts ...Option) (results []*ProductAttributeImage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductAttributeImageMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*ProductAttributeImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

func (obj *_ProductAttributeImageMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*ProductAttributeImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

func (obj *_ProductAttributeImageMgr) GetFromIDImage(idImage uint32) (results []*ProductAttributeImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error

	return
}

func (obj *_ProductAttributeImageMgr) GetBatchFromIDImage(idImages []uint32) (results []*ProductAttributeImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error

	return
}


func (obj *_ProductAttributeImageMgr) FetchByPrimaryKey(idProductAttribute uint32, idImage uint32) (result ProductAttributeImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_image = ?", idProductAttribute, idImage).Find(&result).Error

	return
}

func (obj *_ProductAttributeImageMgr) FetchIndexByIDImage(idImage uint32) (results []*ProductAttributeImage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error

	return
}
