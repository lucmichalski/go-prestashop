package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgLayeredIndexableFeatureValueLangValueMgr struct {
	*_BaseMgr
}

// EgLayeredIndexableFeatureValueLangValueMgr open func
func EgLayeredIndexableFeatureValueLangValueMgr(db *gorm.DB) *_EgLayeredIndexableFeatureValueLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("EgLayeredIndexableFeatureValueLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLayeredIndexableFeatureValueLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_indexable_feature_value_lang_value"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetTableName() string {
	return "eg_layered_indexable_feature_value_lang_value"
}

// Get 获取
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) Get() (result EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) Gets() (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeatureValue id_feature_value获取 
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) WithIDFeatureValue(idFeatureValue int) Option {
	return optionFunc(func(o *options) { o.query["id_feature_value"] = idFeatureValue })
}

// WithIDLang id_lang获取 
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithURLName url_name获取 
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

// WithMetaTitle meta_title获取 
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}


// GetByOption 功能选项模式获取
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetByOption(opts ...Option) (result EgLayeredIndexableFeatureValueLangValue, err error) {
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
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetByOptions(opts ...Option) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
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


// GetFromIDFeatureValue 通过id_feature_value获取内容  
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetFromIDFeatureValue(idFeatureValue int) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&results).Error
	
	return
}

// GetBatchFromIDFeatureValue 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetBatchFromIDFeatureValue(idFeatureValues []int) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value IN (?)", idFeatureValues).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetFromIDLang(idLang int) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromURLName 通过url_name获取内容  
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetFromURLName(urlName string) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error
	
	return
}

// GetBatchFromURLName 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error
	
	return
}
 
// GetFromMetaTitle 通过meta_title获取内容  
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error
	
	return
}

// GetBatchFromMetaTitle 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLayeredIndexableFeatureValueLangValueMgr) FetchByPrimaryKey(idFeatureValue int ,idLang int ) (result EgLayeredIndexableFeatureValueLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ? AND id_lang = ?", idFeatureValue , idLang).Find(&result).Error
	
	return
}
 

 

	

