package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgFeatureMgr struct {
	*_BaseMgr
}

// EgFeatureMgr open func
func EgFeatureMgr(db *gorm.DB) *_EgFeatureMgr {
	if db == nil {
		panic(fmt.Errorf("EgFeatureMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgFeatureMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_feature"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgFeatureMgr) GetTableName() string {
	return "eg_feature"
}

// Get 获取
func (obj *_EgFeatureMgr) Get() (result EgFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgFeatureMgr) Gets() (results []*EgFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeature id_feature获取 
func (obj *_EgFeatureMgr) WithIDFeature(idFeature uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

// WithPosition position获取 
func (obj *_EgFeatureMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}


// GetByOption 功能选项模式获取
func (obj *_EgFeatureMgr) GetByOption(opts ...Option) (result EgFeature, err error) {
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
func (obj *_EgFeatureMgr) GetByOptions(opts ...Option) (results []*EgFeature, err error) {
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
func (obj *_EgFeatureMgr)  GetFromIDFeature(idFeature uint32) (result EgFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&result).Error
	
	return
}

// GetBatchFromIDFeature 批量唯一主键查找 
func (obj *_EgFeatureMgr) GetBatchFromIDFeature(idFeatures []uint32) (results []*EgFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgFeatureMgr) GetFromPosition(position uint32) (results []*EgFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgFeatureMgr) GetBatchFromPosition(positions []uint32) (results []*EgFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgFeatureMgr) FetchByPrimaryKey(idFeature uint32 ) (result EgFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&result).Error
	
	return
}
 

 

	

