package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentGradeMgr struct {
	*_BaseMgr
}

func ProductCommentGradeMgr(db *gorm.DB) *_ProductCommentGradeMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentGradeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentGradeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment_grade"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductCommentGradeMgr) GetTableName() string {
	return "ps_product_comment_grade"
}

func (obj *_ProductCommentGradeMgr) Get() (result ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductCommentGradeMgr) Gets() (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

func (obj *_ProductCommentGradeMgr) WithIDProductCommentCriterion(idProductCommentCriterion int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

func (obj *_ProductCommentGradeMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

func (obj *_ProductCommentGradeMgr) GetByOption(opts ...Option) (result ProductCommentGrade, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductCommentGradeMgr) GetByOptions(opts ...Option) (results []*ProductCommentGrade, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) GetFromIDProductComment(idProductComment int) (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) GetFromIDProductCommentCriterion(idProductCommentCriterion int) (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []int) (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) GetFromGrade(grade int) (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade = ?", grade).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) GetBatchFromGrade(grades []int) (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade IN (?)", grades).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) FetchByPrimaryKey(idProductComment int, idProductCommentCriterion int) (result ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ? AND id_product_comment_criterion = ?", idProductComment, idProductCommentCriterion).Find(&result).Error

	return
}

func (obj *_ProductCommentGradeMgr) FetchIndexByIDX367426EBACF38A54(idProductComment int) (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}

func (obj *_ProductCommentGradeMgr) FetchIndexByIDX367426EB8375853C(idProductCommentCriterion int) (results []*ProductCommentGrade, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error

	return
}
