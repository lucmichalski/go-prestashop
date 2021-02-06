package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentCriterionProductMgr struct {
	*_BaseMgr
}

// ProductCommentCriterionProductMgr open func
func ProductCommentCriterionProductMgr(db *gorm.DB) *_ProductCommentCriterionProductMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentCriterionProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentCriterionProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment_criterion_product"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductCommentCriterionProductMgr) GetTableName() string {
	return "ps_product_comment_criterion_product"
}

// Get 获取
func (obj *_ProductCommentCriterionProductMgr) Get() (result ProductCommentCriterionProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductCommentCriterionProductMgr) Gets() (results []*ProductCommentCriterionProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取
func (obj *_ProductCommentCriterionProductMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDProductCommentCriterion id_product_comment_criterion获取
func (obj *_ProductCommentCriterionProductMgr) WithIDProductCommentCriterion(idProductCommentCriterion uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

// GetByOption 功能选项模式获取
func (obj *_ProductCommentCriterionProductMgr) GetByOption(opts ...Option) (result ProductCommentCriterionProduct, err error) {
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
func (obj *_ProductCommentCriterionProductMgr) GetByOptions(opts ...Option) (results []*ProductCommentCriterionProduct, err error) {
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

// GetFromIDProduct 通过id_product获取内容
func (obj *_ProductCommentCriterionProductMgr) GetFromIDProduct(idProduct uint32) (results []*ProductCommentCriterionProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_ProductCommentCriterionProductMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductCommentCriterionProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDProductCommentCriterion 通过id_product_comment_criterion获取内容
func (obj *_ProductCommentCriterionProductMgr) GetFromIDProductCommentCriterion(idProductCommentCriterion uint32) (results []*ProductCommentCriterionProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error

	return
}

// GetBatchFromIDProductCommentCriterion 批量唯一主键查找
func (obj *_ProductCommentCriterionProductMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []uint32) (results []*ProductCommentCriterionProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductCommentCriterionProductMgr) FetchByPrimaryKey(idProduct uint32, idProductCommentCriterion uint32) (result ProductCommentCriterionProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_product_comment_criterion = ?", idProduct, idProductCommentCriterion).Find(&result).Error

	return
}

// FetchIndexByIDProductCommentCriterion  获取多个内容
func (obj *_ProductCommentCriterionProductMgr) FetchIndexByIDProductCommentCriterion(idProductCommentCriterion uint32) (results []*ProductCommentCriterionProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error

	return
}
