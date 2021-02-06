package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLayeredIndexableFeatureLangValueMgr struct {
	*_BaseMgr
}

// EgLayeredIndexableFeatureLangValueMgr open func
func EgLayeredIndexableFeatureLangValueMgr(db *gorm.DB) *_EgLayeredIndexableFeatureLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("EgLayeredIndexableFeatureLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLayeredIndexableFeatureLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_indexable_feature_lang_value"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetTableName() string {
	return "eg_layered_indexable_feature_lang_value"
}

// Get 获取
func (obj *_EgLayeredIndexableFeatureLangValueMgr) Get() (result EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLayeredIndexableFeatureLangValueMgr) Gets() (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeature id_feature获取 
func (obj *_EgLayeredIndexableFeatureLangValueMgr) WithIDFeature(idFeature int) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

// WithIDLang id_lang获取 
func (obj *_EgLayeredIndexableFeatureLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithURLName url_name获取 
func (obj *_EgLayeredIndexableFeatureLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

// WithMetaTitle meta_title获取 
func (obj *_EgLayeredIndexableFeatureLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}


// GetByOption 功能选项模式获取
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetByOption(opts ...Option) (result EgLayeredIndexableFeatureLangValue, err error) {
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
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetByOptions(opts ...Option) (results []*EgLayeredIndexableFeatureLangValue, err error) {
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


// GetFromIDFeature 通过id_feature获取内容  
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetFromIDFeature(idFeature int) (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error
	
	return
}

// GetBatchFromIDFeature 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetBatchFromIDFeature(idFeatures []int) (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetFromIDLang(idLang int) (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromURLName 通过url_name获取内容  
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetFromURLName(urlName string) (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error
	
	return
}

// GetBatchFromURLName 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error
	
	return
}
 
// GetFromMetaTitle 通过meta_title获取内容  
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error
	
	return
}

// GetBatchFromMetaTitle 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLayeredIndexableFeatureLangValueMgr) FetchByPrimaryKey(idFeature int ,idLang int ) (result EgLayeredIndexableFeatureLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ? AND id_lang = ?", idFeature , idLang).Find(&result).Error
	
	return
}
 

 

	

