package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentCriterionMgr struct {
	*_BaseMgr
}

func ProductCommentCriterionMgr(db *gorm.DB) *_ProductCommentCriterionMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentCriterionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentCriterionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment_criterion"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductCommentCriterionMgr) GetTableName() string {
	return "ps_product_comment_criterion"
}

func (obj *_ProductCommentCriterionMgr) Get() (result ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductCommentCriterionMgr) Gets() (results []*ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionMgr) WithIDProductCommentCriterion(idProductCommentCriterion int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

func (obj *_ProductCommentCriterionMgr) WithIDProductCommentCriterionType(idProductCommentCriterionType int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion_type"] = idProductCommentCriterionType })
}

func (obj *_ProductCommentCriterionMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ProductCommentCriterionMgr) GetByOption(opts ...Option) (result ProductCommentCriterion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductCommentCriterionMgr) GetByOptions(opts ...Option) (results []*ProductCommentCriterion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionMgr) GetFromIDProductCommentCriterion(idProductCommentCriterion int) (result ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&result).Error

	return
}

func (obj *_ProductCommentCriterionMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []int) (results []*ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionMgr) GetFromIDProductCommentCriterionType(idProductCommentCriterionType int) (results []*ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion_type = ?", idProductCommentCriterionType).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionMgr) GetBatchFromIDProductCommentCriterionType(idProductCommentCriterionTypes []int) (results []*ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion_type IN (?)", idProductCommentCriterionTypes).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionMgr) GetFromActive(active bool) (results []*ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionMgr) GetBatchFromActive(actives []bool) (results []*ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_ProductCommentCriterionMgr) FetchByPrimaryKey(idProductCommentCriterion int) (result ProductCommentCriterion, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&result).Error

	return
}
