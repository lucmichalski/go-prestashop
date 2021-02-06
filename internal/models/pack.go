package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PackMgr struct {
	*_BaseMgr
}

func PackMgr(db *gorm.DB) *_PackMgr {
	if db == nil {
		panic(fmt.Errorf("PackMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PackMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_pack"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PackMgr) GetTableName() string {
	return "ps_pack"
}

func (obj *_PackMgr) Get() (result Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PackMgr) Gets() (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PackMgr) WithIDProductPack(idProductPack uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_pack"] = idProductPack })
}

func (obj *_PackMgr) WithIDProductItem(idProductItem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_item"] = idProductItem })
}

func (obj *_PackMgr) WithIDProductAttributeItem(idProductAttributeItem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute_item"] = idProductAttributeItem })
}

func (obj *_PackMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

func (obj *_PackMgr) GetByOption(opts ...Option) (result Pack, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_PackMgr) GetByOptions(opts ...Option) (results []*Pack, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_PackMgr) GetFromIDProductPack(idProductPack uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack = ?", idProductPack).Find(&results).Error

	return
}

func (obj *_PackMgr) GetBatchFromIDProductPack(idProductPacks []uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack IN (?)", idProductPacks).Find(&results).Error

	return
}

func (obj *_PackMgr) GetFromIDProductItem(idProductItem uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item = ?", idProductItem).Find(&results).Error

	return
}

func (obj *_PackMgr) GetBatchFromIDProductItem(idProductItems []uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item IN (?)", idProductItems).Find(&results).Error

	return
}

func (obj *_PackMgr) GetFromIDProductAttributeItem(idProductAttributeItem uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute_item = ?", idProductAttributeItem).Find(&results).Error

	return
}

func (obj *_PackMgr) GetBatchFromIDProductAttributeItem(idProductAttributeItems []uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute_item IN (?)", idProductAttributeItems).Find(&results).Error

	return
}

func (obj *_PackMgr) GetFromQuantity(quantity uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

func (obj *_PackMgr) GetBatchFromQuantity(quantitys []uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}


func (obj *_PackMgr) FetchByPrimaryKey(idProductPack uint32, idProductItem uint32, idProductAttributeItem uint32) (result Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack = ? AND id_product_item = ? AND id_product_attribute_item = ?", idProductPack, idProductItem, idProductAttributeItem).Find(&result).Error

	return
}

func (obj *_PackMgr) FetchIndexByProductItem(idProductItem uint32, idProductAttributeItem uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item = ? AND id_product_attribute_item = ?", idProductItem, idProductAttributeItem).Find(&results).Error

	return
}
