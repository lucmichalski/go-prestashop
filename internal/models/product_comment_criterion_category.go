package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentCriterionCategoryMgr struct {
	*_BaseMgr
}

func ProductCommentCriterionCategoryMgr(db *gorm.DB) *_ProductCommentCriterionCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentCriterionCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentCriterionCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment_criterion_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductCommentCriterionCategoryMgr) GetTableName() string {
	return "ps_product_comment_criterion_category"
}

func (obj *_ProductCommentCriterionCategoryMgr) Get() (result ProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) Gets() (results []*ProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) WithIDProductCommentCriterion(idProductCommentCriterion uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

func (obj *_ProductCommentCriterionCategoryMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

func (obj *_ProductCommentCriterionCategoryMgr) GetByOption(opts ...Option) (result ProductCommentCriterionCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) GetByOptions(opts ...Option) (results []*ProductCommentCriterionCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) GetFromIDProductCommentCriterion(idProductCommentCriterion uint32) (results []*ProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []uint32) (results []*ProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) GetFromIDCategory(idCategory uint32) (results []*ProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*ProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) FetchByPrimaryKey(idProductCommentCriterion uint32, idCategory uint32) (result ProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ? AND id_category = ?", idProductCommentCriterion, idCategory).Find(&result).Error

	return
}

func (obj *_ProductCommentCriterionCategoryMgr) FetchIndexByIDCategory(idCategory uint32) (results []*ProductCommentCriterionCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}
