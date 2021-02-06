package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgFeatureValueLangMgr struct {
	*_BaseMgr
}

// EgFeatureValueLangMgr open func
func EgFeatureValueLangMgr(db *gorm.DB) *_EgFeatureValueLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgFeatureValueLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgFeatureValueLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_feature_value_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgFeatureValueLangMgr) GetTableName() string {
	return "eg_feature_value_lang"
}

// Get 获取
func (obj *_EgFeatureValueLangMgr) Get() (result EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgFeatureValueLangMgr) Gets() (results []*EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeatureValue id_feature_value获取 
func (obj *_EgFeatureValueLangMgr) WithIDFeatureValue(idFeatureValue uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature_value"] = idFeatureValue })
}

// WithIDLang id_lang获取 
func (obj *_EgFeatureValueLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithValue value获取 
func (obj *_EgFeatureValueLangMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}


// GetByOption 功能选项模式获取
func (obj *_EgFeatureValueLangMgr) GetByOption(opts ...Option) (result EgFeatureValueLang, err error) {
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
func (obj *_EgFeatureValueLangMgr) GetByOptions(opts ...Option) (results []*EgFeatureValueLang, err error) {
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
func (obj *_EgFeatureValueLangMgr) GetFromIDFeatureValue(idFeatureValue uint32) (results []*EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&results).Error
	
	return
}

// GetBatchFromIDFeatureValue 批量唯一主键查找 
func (obj *_EgFeatureValueLangMgr) GetBatchFromIDFeatureValue(idFeatureValues []uint32) (results []*EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value IN (?)", idFeatureValues).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgFeatureValueLangMgr) GetFromIDLang(idLang uint32) (results []*EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgFeatureValueLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_EgFeatureValueLangMgr) GetFromValue(value string) (results []*EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量唯一主键查找 
func (obj *_EgFeatureValueLangMgr) GetBatchFromValue(values []string) (results []*EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgFeatureValueLangMgr) FetchByPrimaryKey(idFeatureValue uint32 ,idLang uint32 ) (result EgFeatureValueLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ? AND id_lang = ?", idFeatureValue , idLang).Find(&result).Error
	
	return
}
 

 

	

