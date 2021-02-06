package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LangMgr struct {
	*_BaseMgr
}

// LangMgr open func
func LangMgr(db *gorm.DB) *_LangMgr {
	if db == nil {
		panic(fmt.Errorf("LangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LangMgr) GetTableName() string {
	return "ps_lang"
}

// Get 获取
func (obj *_LangMgr) Get() (result Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LangMgr) Gets() (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLang id_lang获取
func (obj *_LangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_LangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithActive active获取
func (obj *_LangMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIsoCode iso_code获取
func (obj *_LangMgr) WithIsoCode(isoCode string) Option {
	return optionFunc(func(o *options) { o.query["iso_code"] = isoCode })
}

// WithLanguageCode language_code获取
func (obj *_LangMgr) WithLanguageCode(languageCode string) Option {
	return optionFunc(func(o *options) { o.query["language_code"] = languageCode })
}

// WithLocale locale获取
func (obj *_LangMgr) WithLocale(locale string) Option {
	return optionFunc(func(o *options) { o.query["locale"] = locale })
}

// WithDateFormatLite date_format_lite获取
func (obj *_LangMgr) WithDateFormatLite(dateFormatLite string) Option {
	return optionFunc(func(o *options) { o.query["date_format_lite"] = dateFormatLite })
}

// WithDateFormatFull date_format_full获取
func (obj *_LangMgr) WithDateFormatFull(dateFormatFull string) Option {
	return optionFunc(func(o *options) { o.query["date_format_full"] = dateFormatFull })
}

// WithIsRtl is_rtl获取
func (obj *_LangMgr) WithIsRtl(isRtl bool) Option {
	return optionFunc(func(o *options) { o.query["is_rtl"] = isRtl })
}

// GetByOption 功能选项模式获取
func (obj *_LangMgr) GetByOption(opts ...Option) (result Lang, err error) {
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
func (obj *_LangMgr) GetByOptions(opts ...Option) (results []*Lang, err error) {
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

// GetFromIDLang 通过id_lang获取内容
func (obj *_LangMgr) GetFromIDLang(idLang int) (result Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&result).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromIDLang(idLangs []int) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_LangMgr) GetFromName(name string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromName(names []string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_LangMgr) GetFromActive(active bool) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromActive(actives []bool) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromIsoCode 通过iso_code获取内容
func (obj *_LangMgr) GetFromIsoCode(isoCode string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error

	return
}

// GetBatchFromIsoCode 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromIsoCode(isoCodes []string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code IN (?)", isoCodes).Find(&results).Error

	return
}

// GetFromLanguageCode 通过language_code获取内容
func (obj *_LangMgr) GetFromLanguageCode(languageCode string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("language_code = ?", languageCode).Find(&results).Error

	return
}

// GetBatchFromLanguageCode 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromLanguageCode(languageCodes []string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("language_code IN (?)", languageCodes).Find(&results).Error

	return
}

// GetFromLocale 通过locale获取内容
func (obj *_LangMgr) GetFromLocale(locale string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("locale = ?", locale).Find(&results).Error

	return
}

// GetBatchFromLocale 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromLocale(locales []string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("locale IN (?)", locales).Find(&results).Error

	return
}

// GetFromDateFormatLite 通过date_format_lite获取内容
func (obj *_LangMgr) GetFromDateFormatLite(dateFormatLite string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_format_lite = ?", dateFormatLite).Find(&results).Error

	return
}

// GetBatchFromDateFormatLite 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromDateFormatLite(dateFormatLites []string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_format_lite IN (?)", dateFormatLites).Find(&results).Error

	return
}

// GetFromDateFormatFull 通过date_format_full获取内容
func (obj *_LangMgr) GetFromDateFormatFull(dateFormatFull string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_format_full = ?", dateFormatFull).Find(&results).Error

	return
}

// GetBatchFromDateFormatFull 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromDateFormatFull(dateFormatFulls []string) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_format_full IN (?)", dateFormatFulls).Find(&results).Error

	return
}

// GetFromIsRtl 通过is_rtl获取内容
func (obj *_LangMgr) GetFromIsRtl(isRtl bool) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_rtl = ?", isRtl).Find(&results).Error

	return
}

// GetBatchFromIsRtl 批量唯一主键查找
func (obj *_LangMgr) GetBatchFromIsRtl(isRtls []bool) (results []*Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_rtl IN (?)", isRtls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LangMgr) FetchByPrimaryKey(idLang int) (result Lang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&result).Error

	return
}
