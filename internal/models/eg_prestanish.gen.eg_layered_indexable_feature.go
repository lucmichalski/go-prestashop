package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLayeredIndexableFeatureMgr struct {
	*_BaseMgr
}

// EgLayeredIndexableFeatureMgr open func
func EgLayeredIndexableFeatureMgr(db *gorm.DB) *_EgLayeredIndexableFeatureMgr {
	if db == nil {
		panic(fmt.Errorf("EgLayeredIndexableFeatureMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLayeredIndexableFeatureMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_indexable_feature"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLayeredIndexableFeatureMgr) GetTableName() string {
	return "eg_layered_indexable_feature"
}

// Get 获取
func (obj *_EgLayeredIndexableFeatureMgr) Get() (result EgLayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLayeredIndexableFeatureMgr) Gets() (results []*EgLayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeature id_feature获取 
func (obj *_EgLayeredIndexableFeatureMgr) WithIDFeature(idFeature int) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

// WithIndexable indexable获取 
func (obj *_EgLayeredIndexableFeatureMgr) WithIndexable(indexable bool) Option {
	return optionFunc(func(o *options) { o.query["indexable"] = indexable })
}


// GetByOption 功能选项模式获取
func (obj *_EgLayeredIndexableFeatureMgr) GetByOption(opts ...Option) (result EgLayeredIndexableFeature, err error) {
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
func (obj *_EgLayeredIndexableFeatureMgr) GetByOptions(opts ...Option) (results []*EgLayeredIndexableFeature, err error) {
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
func (obj *_EgLayeredIndexableFeatureMgr)  GetFromIDFeature(idFeature int) (result EgLayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&result).Error
	
	return
}

// GetBatchFromIDFeature 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureMgr) GetBatchFromIDFeature(idFeatures []int) (results []*EgLayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error
	
	return
}
 
// GetFromIndexable 通过indexable获取内容  
func (obj *_EgLayeredIndexableFeatureMgr) GetFromIndexable(indexable bool) (results []*EgLayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable = ?", indexable).Find(&results).Error
	
	return
}

// GetBatchFromIndexable 批量唯一主键查找 
func (obj *_EgLayeredIndexableFeatureMgr) GetBatchFromIndexable(indexables []bool) (results []*EgLayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable IN (?)", indexables).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLayeredIndexableFeatureMgr) FetchByPrimaryKey(idFeature int ) (result EgLayeredIndexableFeature, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&result).Error
	
	return
}
 

 

	

