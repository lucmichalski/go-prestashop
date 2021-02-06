package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductAttributeCombinationMgr struct {
	*_BaseMgr
}

// ProductAttributeCombinationMgr open func
func ProductAttributeCombinationMgr(db *gorm.DB) *_ProductAttributeCombinationMgr {
	if db == nil {
		panic(fmt.Errorf("ProductAttributeCombinationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductAttributeCombinationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_attribute_combination"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductAttributeCombinationMgr) GetTableName() string {
	return "ps_product_attribute_combination"
}

// Get 获取
func (obj *_ProductAttributeCombinationMgr) Get() (result ProductAttributeCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductAttributeCombinationMgr) Gets() (results []*ProductAttributeCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttribute id_attribute获取
func (obj *_ProductAttributeCombinationMgr) WithIDAttribute(idAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

// WithIDProductAttribute id_product_attribute获取
func (obj *_ProductAttributeCombinationMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// GetByOption 功能选项模式获取
func (obj *_ProductAttributeCombinationMgr) GetByOption(opts ...Option) (result ProductAttributeCombination, err error) {
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
func (obj *_ProductAttributeCombinationMgr) GetByOptions(opts ...Option) (results []*ProductAttributeCombination, err error) {
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
func (obj *_ProductAttributeCombinationMgr) GetFromIDAttribute(idAttribute uint32) (results []*ProductAttributeCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error

	return
}

// GetBatchFromIDAttribute 批量唯一主键查找
func (obj *_ProductAttributeCombinationMgr) GetBatchFromIDAttribute(idAttributes []uint32) (results []*ProductAttributeCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error

	return
}

// GetFromIDProductAttribute 通过id_product_attribute获取内容
func (obj *_ProductAttributeCombinationMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*ProductAttributeCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找
func (obj *_ProductAttributeCombinationMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*ProductAttributeCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductAttributeCombinationMgr) FetchByPrimaryKey(idAttribute uint32, idProductAttribute uint32) (result ProductAttributeCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_product_attribute = ?", idAttribute, idProductAttribute).Find(&result).Error

	return
}

// FetchIndexByIDProductAttribute  获取多个内容
func (obj *_ProductAttributeCombinationMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32) (results []*ProductAttributeCombination, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error

	return
}
