package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgReferrerCacheMgr struct {
	*_BaseMgr
}

// EgReferrerCacheMgr open func
func EgReferrerCacheMgr(db *gorm.DB) *_EgReferrerCacheMgr {
	if db == nil {
		panic(fmt.Errorf("EgReferrerCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgReferrerCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_referrer_cache"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgReferrerCacheMgr) GetTableName() string {
	return "eg_referrer_cache"
}

// Get 获取
func (obj *_EgReferrerCacheMgr) Get() (result EgReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgReferrerCacheMgr) Gets() (results []*EgReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConnectionsSource id_connections_source获取 
func (obj *_EgReferrerCacheMgr) WithIDConnectionsSource(idConnectionsSource uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections_source"] = idConnectionsSource })
}

// WithIDReferrer id_referrer获取 
func (obj *_EgReferrerCacheMgr) WithIDReferrer(idReferrer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_referrer"] = idReferrer })
}


// GetByOption 功能选项模式获取
func (obj *_EgReferrerCacheMgr) GetByOption(opts ...Option) (result EgReferrerCache, err error) {
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
func (obj *_EgReferrerCacheMgr) GetByOptions(opts ...Option) (results []*EgReferrerCache, err error) {
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


// GetFromIDConnectionsSource 通过id_connections_source获取内容  
func (obj *_EgReferrerCacheMgr) GetFromIDConnectionsSource(idConnectionsSource uint32) (results []*EgReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ?", idConnectionsSource).Find(&results).Error
	
	return
}

// GetBatchFromIDConnectionsSource 批量唯一主键查找 
func (obj *_EgReferrerCacheMgr) GetBatchFromIDConnectionsSource(idConnectionsSources []uint32) (results []*EgReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source IN (?)", idConnectionsSources).Find(&results).Error
	
	return
}
 
// GetFromIDReferrer 通过id_referrer获取内容  
func (obj *_EgReferrerCacheMgr) GetFromIDReferrer(idReferrer uint32) (results []*EgReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&results).Error
	
	return
}

// GetBatchFromIDReferrer 批量唯一主键查找 
func (obj *_EgReferrerCacheMgr) GetBatchFromIDReferrer(idReferrers []uint32) (results []*EgReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer IN (?)", idReferrers).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgReferrerCacheMgr) FetchByPrimaryKey(idConnectionsSource uint32 ,idReferrer uint32 ) (result EgReferrerCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ? AND id_referrer = ?", idConnectionsSource , idReferrer).Find(&result).Error
	
	return
}
 

 

	

