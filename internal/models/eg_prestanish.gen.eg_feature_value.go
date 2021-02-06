package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgFeatureValueMgr struct {
	*_BaseMgr
}

// EgFeatureValueMgr open func
func EgFeatureValueMgr(db *gorm.DB) *_EgFeatureValueMgr {
	if db == nil {
		panic(fmt.Errorf("EgFeatureValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgFeatureValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_feature_value"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgFeatureValueMgr) GetTableName() string {
	return "eg_feature_value"
}

// Get 获取
func (obj *_EgFeatureValueMgr) Get() (result EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgFeatureValueMgr) Gets() (results []*EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeatureValue id_feature_value获取 
func (obj *_EgFeatureValueMgr) WithIDFeatureValue(idFeatureValue uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature_value"] = idFeatureValue })
}

// WithIDFeature id_feature获取 
func (obj *_EgFeatureValueMgr) WithIDFeature(idFeature uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

// WithCustom custom获取 
func (obj *_EgFeatureValueMgr) WithCustom(custom uint8) Option {
	return optionFunc(func(o *options) { o.query["custom"] = custom })
}


// GetByOption 功能选项模式获取
func (obj *_EgFeatureValueMgr) GetByOption(opts ...Option) (result EgFeatureValue, err error) {
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
func (obj *_EgFeatureValueMgr) GetByOptions(opts ...Option) (results []*EgFeatureValue, err error) {
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
func (obj *_EgFeatureValueMgr)  GetFromIDFeatureValue(idFeatureValue uint32) (result EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&result).Error
	
	return
}

// GetBatchFromIDFeatureValue 批量唯一主键查找 
func (obj *_EgFeatureValueMgr) GetBatchFromIDFeatureValue(idFeatureValues []uint32) (results []*EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value IN (?)", idFeatureValues).Find(&results).Error
	
	return
}
 
// GetFromIDFeature 通过id_feature获取内容  
func (obj *_EgFeatureValueMgr) GetFromIDFeature(idFeature uint32) (results []*EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error
	
	return
}

// GetBatchFromIDFeature 批量唯一主键查找 
func (obj *_EgFeatureValueMgr) GetBatchFromIDFeature(idFeatures []uint32) (results []*EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error
	
	return
}
 
// GetFromCustom 通过custom获取内容  
func (obj *_EgFeatureValueMgr) GetFromCustom(custom uint8) (results []*EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom = ?", custom).Find(&results).Error
	
	return
}

// GetBatchFromCustom 批量唯一主键查找 
func (obj *_EgFeatureValueMgr) GetBatchFromCustom(customs []uint8) (results []*EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom IN (?)", customs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgFeatureValueMgr) FetchByPrimaryKey(idFeatureValue uint32 ) (result EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByFeature  获取多个内容
 func (obj *_EgFeatureValueMgr) FetchIndexByFeature(idFeature uint32 ) (results []*EgFeatureValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error
	
	return
}
 

	

