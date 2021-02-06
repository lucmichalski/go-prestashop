package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgManufacturerLangMgr struct {
	*_BaseMgr
}

// EgManufacturerLangMgr open func
func EgManufacturerLangMgr(db *gorm.DB) *_EgManufacturerLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgManufacturerLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgManufacturerLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_manufacturer_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgManufacturerLangMgr) GetTableName() string {
	return "eg_manufacturer_lang"
}

// Get 获取
func (obj *_EgManufacturerLangMgr) Get() (result EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgManufacturerLangMgr) Gets() (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDManufacturer id_manufacturer获取 
func (obj *_EgManufacturerLangMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

// WithIDLang id_lang获取 
func (obj *_EgManufacturerLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithDescription description获取 
func (obj *_EgManufacturerLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithShortDescription short_description获取 
func (obj *_EgManufacturerLangMgr) WithShortDescription(shortDescription string) Option {
	return optionFunc(func(o *options) { o.query["short_description"] = shortDescription })
}

// WithMetaTitle meta_title获取 
func (obj *_EgManufacturerLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// WithMetaKeywords meta_keywords获取 
func (obj *_EgManufacturerLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取 
func (obj *_EgManufacturerLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}


// GetByOption 功能选项模式获取
func (obj *_EgManufacturerLangMgr) GetByOption(opts ...Option) (result EgManufacturerLang, err error) {
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
func (obj *_EgManufacturerLangMgr) GetByOptions(opts ...Option) (results []*EgManufacturerLang, err error) {
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


// GetFromIDManufacturer 通过id_manufacturer获取内容  
func (obj *_EgManufacturerLangMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error
	
	return
}

// GetBatchFromIDManufacturer 批量唯一主键查找 
func (obj *_EgManufacturerLangMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgManufacturerLangMgr) GetFromIDLang(idLang uint32) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgManufacturerLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgManufacturerLangMgr) GetFromDescription(description string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgManufacturerLangMgr) GetBatchFromDescription(descriptions []string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromShortDescription 通过short_description获取内容  
func (obj *_EgManufacturerLangMgr) GetFromShortDescription(shortDescription string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("short_description = ?", shortDescription).Find(&results).Error
	
	return
}

// GetBatchFromShortDescription 批量唯一主键查找 
func (obj *_EgManufacturerLangMgr) GetBatchFromShortDescription(shortDescriptions []string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("short_description IN (?)", shortDescriptions).Find(&results).Error
	
	return
}
 
// GetFromMetaTitle 通过meta_title获取内容  
func (obj *_EgManufacturerLangMgr) GetFromMetaTitle(metaTitle string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error
	
	return
}

// GetBatchFromMetaTitle 批量唯一主键查找 
func (obj *_EgManufacturerLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error
	
	return
}
 
// GetFromMetaKeywords 通过meta_keywords获取内容  
func (obj *_EgManufacturerLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error
	
	return
}

// GetBatchFromMetaKeywords 批量唯一主键查找 
func (obj *_EgManufacturerLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error
	
	return
}
 
// GetFromMetaDescription 通过meta_description获取内容  
func (obj *_EgManufacturerLangMgr) GetFromMetaDescription(metaDescription string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error
	
	return
}

// GetBatchFromMetaDescription 批量唯一主键查找 
func (obj *_EgManufacturerLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgManufacturerLangMgr) FetchByPrimaryKey(idManufacturer uint32 ,idLang uint32 ) (result EgManufacturerLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ? AND id_lang = ?", idManufacturer , idLang).Find(&result).Error
	
	return
}
 

 

	

