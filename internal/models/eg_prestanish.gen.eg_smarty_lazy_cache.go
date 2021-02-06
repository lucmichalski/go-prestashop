package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgSmartyLazyCacheMgr struct {
	*_BaseMgr
}

// EgSmartyLazyCacheMgr open func
func EgSmartyLazyCacheMgr(db *gorm.DB) *_EgSmartyLazyCacheMgr {
	if db == nil {
		panic(fmt.Errorf("EgSmartyLazyCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSmartyLazyCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_smarty_lazy_cache"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSmartyLazyCacheMgr) GetTableName() string {
	return "eg_smarty_lazy_cache"
}

// Get 获取
func (obj *_EgSmartyLazyCacheMgr) Get() (result EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSmartyLazyCacheMgr) Gets() (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithTemplateHash template_hash获取 
func (obj *_EgSmartyLazyCacheMgr) WithTemplateHash(templateHash string) Option {
	return optionFunc(func(o *options) { o.query["template_hash"] = templateHash })
}

// WithCacheID cache_id获取 
func (obj *_EgSmartyLazyCacheMgr) WithCacheID(cacheID string) Option {
	return optionFunc(func(o *options) { o.query["cache_id"] = cacheID })
}

// WithCompileID compile_id获取 
func (obj *_EgSmartyLazyCacheMgr) WithCompileID(compileID string) Option {
	return optionFunc(func(o *options) { o.query["compile_id"] = compileID })
}

// WithFilepath filepath获取 
func (obj *_EgSmartyLazyCacheMgr) WithFilepath(filepath string) Option {
	return optionFunc(func(o *options) { o.query["filepath"] = filepath })
}

// WithLastUpdate last_update获取 
func (obj *_EgSmartyLazyCacheMgr) WithLastUpdate(lastUpdate time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_update"] = lastUpdate })
}


// GetByOption 功能选项模式获取
func (obj *_EgSmartyLazyCacheMgr) GetByOption(opts ...Option) (result EgSmartyLazyCache, err error) {
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
func (obj *_EgSmartyLazyCacheMgr) GetByOptions(opts ...Option) (results []*EgSmartyLazyCache, err error) {
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


// GetFromTemplateHash 通过template_hash获取内容  
func (obj *_EgSmartyLazyCacheMgr) GetFromTemplateHash(templateHash string) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash = ?", templateHash).Find(&results).Error
	
	return
}

// GetBatchFromTemplateHash 批量唯一主键查找 
func (obj *_EgSmartyLazyCacheMgr) GetBatchFromTemplateHash(templateHashs []string) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash IN (?)", templateHashs).Find(&results).Error
	
	return
}
 
// GetFromCacheID 通过cache_id获取内容  
func (obj *_EgSmartyLazyCacheMgr) GetFromCacheID(cacheID string) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error
	
	return
}

// GetBatchFromCacheID 批量唯一主键查找 
func (obj *_EgSmartyLazyCacheMgr) GetBatchFromCacheID(cacheIDs []string) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id IN (?)", cacheIDs).Find(&results).Error
	
	return
}
 
// GetFromCompileID 通过compile_id获取内容  
func (obj *_EgSmartyLazyCacheMgr) GetFromCompileID(compileID string) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("compile_id = ?", compileID).Find(&results).Error
	
	return
}

// GetBatchFromCompileID 批量唯一主键查找 
func (obj *_EgSmartyLazyCacheMgr) GetBatchFromCompileID(compileIDs []string) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("compile_id IN (?)", compileIDs).Find(&results).Error
	
	return
}
 
// GetFromFilepath 通过filepath获取内容  
func (obj *_EgSmartyLazyCacheMgr) GetFromFilepath(filepath string) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filepath = ?", filepath).Find(&results).Error
	
	return
}

// GetBatchFromFilepath 批量唯一主键查找 
func (obj *_EgSmartyLazyCacheMgr) GetBatchFromFilepath(filepaths []string) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filepath IN (?)", filepaths).Find(&results).Error
	
	return
}
 
// GetFromLastUpdate 通过last_update获取内容  
func (obj *_EgSmartyLazyCacheMgr) GetFromLastUpdate(lastUpdate time.Time) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_update = ?", lastUpdate).Find(&results).Error
	
	return
}

// GetBatchFromLastUpdate 批量唯一主键查找 
func (obj *_EgSmartyLazyCacheMgr) GetBatchFromLastUpdate(lastUpdates []time.Time) (results []*EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_update IN (?)", lastUpdates).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSmartyLazyCacheMgr) FetchByPrimaryKey(templateHash string ,cacheID string ,compileID string ) (result EgSmartyLazyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template_hash = ? AND cache_id = ? AND compile_id = ?", templateHash , cacheID , compileID).Find(&result).Error
	
	return
}
 

 

	

