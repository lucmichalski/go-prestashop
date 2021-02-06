package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PackMgr struct {
	*_BaseMgr
}

// PackMgr open func
func PackMgr(db *gorm.DB) *_PackMgr {
	if db == nil {
		panic(fmt.Errorf("PackMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PackMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_pack"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PackMgr) GetTableName() string {
	return "ps_pack"
}

// Get 获取
func (obj *_PackMgr) Get() (result Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PackMgr) Gets() (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductPack id_product_pack获取
func (obj *_PackMgr) WithIDProductPack(idProductPack uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_pack"] = idProductPack })
}

// WithIDProductItem id_product_item获取
func (obj *_PackMgr) WithIDProductItem(idProductItem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_item"] = idProductItem })
}

// WithIDProductAttributeItem id_product_attribute_item获取
func (obj *_PackMgr) WithIDProductAttributeItem(idProductAttributeItem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute_item"] = idProductAttributeItem })
}

// WithQuantity quantity获取
func (obj *_PackMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDProductPack 通过id_product_pack获取内容
func (obj *_PackMgr) GetFromIDProductPack(idProductPack uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack = ?", idProductPack).Find(&results).Error

	return
}

// GetBatchFromIDProductPack 批量唯一主键查找
func (obj *_PackMgr) GetBatchFromIDProductPack(idProductPacks []uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack IN (?)", idProductPacks).Find(&results).Error

	return
}

// GetFromIDProductItem 通过id_product_item获取内容
func (obj *_PackMgr) GetFromIDProductItem(idProductItem uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item = ?", idProductItem).Find(&results).Error

	return
}

// GetBatchFromIDProductItem 批量唯一主键查找
func (obj *_PackMgr) GetBatchFromIDProductItem(idProductItems []uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item IN (?)", idProductItems).Find(&results).Error

	return
}

// GetFromIDProductAttributeItem 通过id_product_attribute_item获取内容
func (obj *_PackMgr) GetFromIDProductAttributeItem(idProductAttributeItem uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute_item = ?", idProductAttributeItem).Find(&results).Error

	return
}

// GetBatchFromIDProductAttributeItem 批量唯一主键查找
func (obj *_PackMgr) GetBatchFromIDProductAttributeItem(idProductAttributeItems []uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute_item IN (?)", idProductAttributeItems).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容
func (obj *_PackMgr) GetFromQuantity(quantity uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量唯一主键查找
func (obj *_PackMgr) GetBatchFromQuantity(quantitys []uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PackMgr) FetchByPrimaryKey(idProductPack uint32, idProductItem uint32, idProductAttributeItem uint32) (result Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack = ? AND id_product_item = ? AND id_product_attribute_item = ?", idProductPack, idProductItem, idProductAttributeItem).Find(&result).Error

	return
}

// FetchIndexByProductItem  获取多个内容
func (obj *_PackMgr) FetchIndexByProductItem(idProductItem uint32, idProductAttributeItem uint32) (results []*Pack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item = ? AND id_product_attribute_item = ?", idProductItem, idProductAttributeItem).Find(&results).Error

	return
}
