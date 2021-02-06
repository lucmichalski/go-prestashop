package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentCriterionLangMgr struct {
	*_BaseMgr
}

// ProductCommentCriterionLangMgr open func
func ProductCommentCriterionLangMgr(db *gorm.DB) *_ProductCommentCriterionLangMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentCriterionLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentCriterionLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment_criterion_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductCommentCriterionLangMgr) GetTableName() string {
	return "ps_product_comment_criterion_lang"
}

// Get 获取
func (obj *_ProductCommentCriterionLangMgr) Get() (result ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductCommentCriterionLangMgr) Gets() (results []*ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductCommentCriterion id_product_comment_criterion获取
func (obj *_ProductCommentCriterionLangMgr) WithIDProductCommentCriterion(idProductCommentCriterion uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment_criterion"] = idProductCommentCriterion })
}

// WithIDLang id_lang获取
func (obj *_ProductCommentCriterionLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_ProductCommentCriterionLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_ProductCommentCriterionLangMgr) GetByOption(opts ...Option) (result ProductCommentCriterionLang, err error) {
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
func (obj *_ProductCommentCriterionLangMgr) GetByOptions(opts ...Option) (results []*ProductCommentCriterionLang, err error) {
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

// GetFromIDProductCommentCriterion 通过id_product_comment_criterion获取内容
func (obj *_ProductCommentCriterionLangMgr) GetFromIDProductCommentCriterion(idProductCommentCriterion uint32) (results []*ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ?", idProductCommentCriterion).Find(&results).Error

	return
}

// GetBatchFromIDProductCommentCriterion 批量唯一主键查找
func (obj *_ProductCommentCriterionLangMgr) GetBatchFromIDProductCommentCriterion(idProductCommentCriterions []uint32) (results []*ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion IN (?)", idProductCommentCriterions).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_ProductCommentCriterionLangMgr) GetFromIDLang(idLang uint32) (results []*ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_ProductCommentCriterionLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ProductCommentCriterionLangMgr) GetFromName(name string) (results []*ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ProductCommentCriterionLangMgr) GetBatchFromName(names []string) (results []*ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductCommentCriterionLangMgr) FetchByPrimaryKey(idProductCommentCriterion uint32, idLang uint32) (result ProductCommentCriterionLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment_criterion = ? AND id_lang = ?", idProductCommentCriterion, idLang).Find(&result).Error

	return
}
