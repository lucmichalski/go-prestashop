package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgFeatureLangMgr struct {
	*_BaseMgr
}

// EgFeatureLangMgr open func
func EgFeatureLangMgr(db *gorm.DB) *_EgFeatureLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgFeatureLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgFeatureLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_feature_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgFeatureLangMgr) GetTableName() string {
	return "eg_feature_lang"
}

// Get 获取
func (obj *_EgFeatureLangMgr) Get() (result EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgFeatureLangMgr) Gets() (results []*EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeature id_feature获取 
func (obj *_EgFeatureLangMgr) WithIDFeature(idFeature uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

// WithIDLang id_lang获取 
func (obj *_EgFeatureLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgFeatureLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgFeatureLangMgr) GetByOption(opts ...Option) (result EgFeatureLang, err error) {
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
func (obj *_EgFeatureLangMgr) GetByOptions(opts ...Option) (results []*EgFeatureLang, err error) {
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
func (obj *_EgFeatureLangMgr) GetFromIDFeature(idFeature uint32) (results []*EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error
	
	return
}

// GetBatchFromIDFeature 批量唯一主键查找 
func (obj *_EgFeatureLangMgr) GetBatchFromIDFeature(idFeatures []uint32) (results []*EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgFeatureLangMgr) GetFromIDLang(idLang uint32) (results []*EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgFeatureLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgFeatureLangMgr) GetFromName(name string) (results []*EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgFeatureLangMgr) GetBatchFromName(names []string) (results []*EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgFeatureLangMgr) FetchByPrimaryKey(idFeature uint32 ,idLang uint32 ) (result EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ? AND id_lang = ?", idFeature , idLang).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDLang  获取多个内容
 func (obj *_EgFeatureLangMgr) FetchIndexByIDLang(idLang uint32 ,name string ) (results []*EgFeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ? AND name = ?", idLang , name).Find(&results).Error
	
	return
}
 

	
