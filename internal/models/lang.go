package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLangMgr struct {
	*_BaseMgr
}

// EgLangMgr open func
func EgLangMgr(db *gorm.DB) *_EgLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLangMgr) GetTableName() string {
	return "eg_lang"
}

// Get 获取
func (obj *_EgLangMgr) Get() (result EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLangMgr) Gets() (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLang id_lang获取 
func (obj *_EgLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithActive active获取 
func (obj *_EgLangMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIsoCode iso_code获取 
func (obj *_EgLangMgr) WithIsoCode(isoCode string) Option {
	return optionFunc(func(o *options) { o.query["iso_code"] = isoCode })
}

// WithLanguageCode language_code获取 
func (obj *_EgLangMgr) WithLanguageCode(languageCode string) Option {
	return optionFunc(func(o *options) { o.query["language_code"] = languageCode })
}

// WithLocale locale获取 
func (obj *_EgLangMgr) WithLocale(locale string) Option {
	return optionFunc(func(o *options) { o.query["locale"] = locale })
}

// WithDateFormatLite date_format_lite获取 
func (obj *_EgLangMgr) WithDateFormatLite(dateFormatLite string) Option {
	return optionFunc(func(o *options) { o.query["date_format_lite"] = dateFormatLite })
}

// WithDateFormatFull date_format_full获取 
func (obj *_EgLangMgr) WithDateFormatFull(dateFormatFull string) Option {
	return optionFunc(func(o *options) { o.query["date_format_full"] = dateFormatFull })
}

// WithIsRtl is_rtl获取 
func (obj *_EgLangMgr) WithIsRtl(isRtl bool) Option {
	return optionFunc(func(o *options) { o.query["is_rtl"] = isRtl })
}


// GetByOption 功能选项模式获取
func (obj *_EgLangMgr) GetByOption(opts ...Option) (result EgLang, err error) {
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
func (obj *_EgLangMgr) GetByOptions(opts ...Option) (results []*EgLang, err error) {
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
func (obj *_EgLangMgr)  GetFromIDLang(idLang int) (result EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&result).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromIDLang(idLangs []int) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgLangMgr) GetFromName(name string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromName(names []string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgLangMgr) GetFromActive(active bool) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromActive(actives []bool) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromIsoCode 通过iso_code获取内容  
func (obj *_EgLangMgr) GetFromIsoCode(isoCode string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error
	
	return
}

// GetBatchFromIsoCode 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromIsoCode(isoCodes []string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code IN (?)", isoCodes).Find(&results).Error
	
	return
}
 
// GetFromLanguageCode 通过language_code获取内容  
func (obj *_EgLangMgr) GetFromLanguageCode(languageCode string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("language_code = ?", languageCode).Find(&results).Error
	
	return
}

// GetBatchFromLanguageCode 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromLanguageCode(languageCodes []string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("language_code IN (?)", languageCodes).Find(&results).Error
	
	return
}
 
// GetFromLocale 通过locale获取内容  
func (obj *_EgLangMgr) GetFromLocale(locale string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("locale = ?", locale).Find(&results).Error
	
	return
}

// GetBatchFromLocale 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromLocale(locales []string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("locale IN (?)", locales).Find(&results).Error
	
	return
}
 
// GetFromDateFormatLite 通过date_format_lite获取内容  
func (obj *_EgLangMgr) GetFromDateFormatLite(dateFormatLite string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_format_lite = ?", dateFormatLite).Find(&results).Error
	
	return
}

// GetBatchFromDateFormatLite 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromDateFormatLite(dateFormatLites []string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_format_lite IN (?)", dateFormatLites).Find(&results).Error
	
	return
}
 
// GetFromDateFormatFull 通过date_format_full获取内容  
func (obj *_EgLangMgr) GetFromDateFormatFull(dateFormatFull string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_format_full = ?", dateFormatFull).Find(&results).Error
	
	return
}

// GetBatchFromDateFormatFull 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromDateFormatFull(dateFormatFulls []string) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_format_full IN (?)", dateFormatFulls).Find(&results).Error
	
	return
}
 
// GetFromIsRtl 通过is_rtl获取内容  
func (obj *_EgLangMgr) GetFromIsRtl(isRtl bool) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_rtl = ?", isRtl).Find(&results).Error
	
	return
}

// GetBatchFromIsRtl 批量唯一主键查找 
func (obj *_EgLangMgr) GetBatchFromIsRtl(isRtls []bool) (results []*EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_rtl IN (?)", isRtls).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLangMgr) FetchByPrimaryKey(idLang int ) (result EgLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&result).Error
	
	return
}
 

 

	

