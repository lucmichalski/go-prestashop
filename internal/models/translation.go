package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TranslationMgr struct {
	*_BaseMgr
}

// TranslationMgr open func
func TranslationMgr(db *gorm.DB) *_TranslationMgr {
	if db == nil {
		panic(fmt.Errorf("TranslationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TranslationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_translation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TranslationMgr) GetTableName() string {
	return "ps_translation"
}

// Get 获取
func (obj *_TranslationMgr) Get() (result Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TranslationMgr) Gets() (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTranslation id_translation获取
func (obj *_TranslationMgr) WithIDTranslation(idTranslation int) Option {
	return optionFunc(func(o *options) { o.query["id_translation"] = idTranslation })
}

// WithIDLang id_lang获取
func (obj *_TranslationMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithKey key获取
func (obj *_TranslationMgr) WithKey(key string) Option {
	return optionFunc(func(o *options) { o.query["key"] = key })
}

// WithTranslation translation获取
func (obj *_TranslationMgr) WithTranslation(translation string) Option {
	return optionFunc(func(o *options) { o.query["translation"] = translation })
}

// WithDomain domain获取
func (obj *_TranslationMgr) WithDomain(domain string) Option {
	return optionFunc(func(o *options) { o.query["domain"] = domain })
}

// WithTheme theme获取
func (obj *_TranslationMgr) WithTheme(theme string) Option {
	return optionFunc(func(o *options) { o.query["theme"] = theme })
}

// GetByOption 功能选项模式获取
func (obj *_TranslationMgr) GetByOption(opts ...Option) (result Translation, err error) {
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
func (obj *_TranslationMgr) GetByOptions(opts ...Option) (results []*Translation, err error) {
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

// GetFromIDTranslation 通过id_translation获取内容
func (obj *_TranslationMgr) GetFromIDTranslation(idTranslation int) (result Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_translation = ?", idTranslation).Find(&result).Error

	return
}

// GetBatchFromIDTranslation 批量唯一主键查找
func (obj *_TranslationMgr) GetBatchFromIDTranslation(idTranslations []int) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_translation IN (?)", idTranslations).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_TranslationMgr) GetFromIDLang(idLang int) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_TranslationMgr) GetBatchFromIDLang(idLangs []int) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromKey 通过key获取内容
func (obj *_TranslationMgr) GetFromKey(key string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key = ?", key).Find(&results).Error

	return
}

// GetBatchFromKey 批量唯一主键查找
func (obj *_TranslationMgr) GetBatchFromKey(keys []string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("key IN (?)", keys).Find(&results).Error

	return
}

// GetFromTranslation 通过translation获取内容
func (obj *_TranslationMgr) GetFromTranslation(translation string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("translation = ?", translation).Find(&results).Error

	return
}

// GetBatchFromTranslation 批量唯一主键查找
func (obj *_TranslationMgr) GetBatchFromTranslation(translations []string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("translation IN (?)", translations).Find(&results).Error

	return
}

// GetFromDomain 通过domain获取内容
func (obj *_TranslationMgr) GetFromDomain(domain string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain = ?", domain).Find(&results).Error

	return
}

// GetBatchFromDomain 批量唯一主键查找
func (obj *_TranslationMgr) GetBatchFromDomain(domains []string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain IN (?)", domains).Find(&results).Error

	return
}

// GetFromTheme 通过theme获取内容
func (obj *_TranslationMgr) GetFromTheme(theme string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("theme = ?", theme).Find(&results).Error

	return
}

// GetBatchFromTheme 批量唯一主键查找
func (obj *_TranslationMgr) GetBatchFromTheme(themes []string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("theme IN (?)", themes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TranslationMgr) FetchByPrimaryKey(idTranslation int) (result Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_translation = ?", idTranslation).Find(&result).Error

	return
}

// FetchIndexByIDX3B44757BBA299860  获取多个内容
func (obj *_TranslationMgr) FetchIndexByIDX3B44757BBA299860(idLang int) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// FetchIndexByKey  获取多个内容
func (obj *_TranslationMgr) FetchIndexByKey(domain string) (results []*Translation, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("domain = ?", domain).Find(&results).Error

	return
}
