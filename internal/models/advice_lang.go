package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AdviceLangMgr struct {
	*_BaseMgr
}

// AdviceLangMgr open func
func AdviceLangMgr(db *gorm.DB) *_AdviceLangMgr {
	if db == nil {
		panic(fmt.Errorf("AdviceLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdviceLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_advice_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdviceLangMgr) GetTableName() string {
	return "eg_advice_lang"
}

// Get 获取
func (obj *_AdviceLangMgr) Get() (result AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AdviceLangMgr) Gets() (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAdvice id_advice获取
func (obj *_AdviceLangMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}

// WithIDLang id_lang获取
func (obj *_AdviceLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithHTML html获取
func (obj *_AdviceLangMgr) WithHTML(html string) Option {
	return optionFunc(func(o *options) { o.query["html"] = html })
}

// GetByOption 功能选项模式获取
func (obj *_AdviceLangMgr) GetByOption(opts ...Option) (result AdviceLang, err error) {
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
func (obj *_AdviceLangMgr) GetByOptions(opts ...Option) (results []*AdviceLang, err error) {
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

// GetFromIDAdvice 通过id_advice获取内容
func (obj *_AdviceLangMgr) GetFromIDAdvice(idAdvice int) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&results).Error

	return
}

// GetBatchFromIDAdvice 批量唯一主键查找
func (obj *_AdviceLangMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_AdviceLangMgr) GetFromIDLang(idLang int) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_AdviceLangMgr) GetBatchFromIDLang(idLangs []int) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromHTML 通过html获取内容
func (obj *_AdviceLangMgr) GetFromHTML(html string) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("html = ?", html).Find(&results).Error

	return
}

// GetBatchFromHTML 批量唯一主键查找
func (obj *_AdviceLangMgr) GetBatchFromHTML(htmls []string) (results []*AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("html IN (?)", htmls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AdviceLangMgr) FetchByPrimaryKey(idAdvice int, idLang int) (result AdviceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ? AND id_lang = ?", idAdvice, idLang).Find(&result).Error

	return
}
