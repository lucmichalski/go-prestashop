package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgSmartyCacheMgr struct {
	*_BaseMgr
}

// EgSmartyCacheMgr open func
func EgSmartyCacheMgr(db *gorm.DB) *_EgSmartyCacheMgr {
	if db == nil {
		panic(fmt.Errorf("EgSmartyCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSmartyCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_smarty_cache"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSmartyCacheMgr) GetTableName() string {
	return "eg_smarty_cache"
}

// Get 获取
func (obj *_EgSmartyCacheMgr) Get() (result EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSmartyCacheMgr) Gets() (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSmartyCache id_smarty_cache获取 
func (obj *_EgSmartyCacheMgr) WithIDSmartyCache(idSmartyCache string) Option {
	return optionFunc(func(o *options) { o.query["id_smarty_cache"] = idSmartyCache })
}

// WithName name获取 
func (obj *_EgSmartyCacheMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCacheID cache_id获取 
func (obj *_EgSmartyCacheMgr) WithCacheID(cacheID string) Option {
	return optionFunc(func(o *options) { o.query["cache_id"] = cacheID })
}

// WithModified modified获取 
func (obj *_EgSmartyCacheMgr) WithModified(modified time.Time) Option {
	return optionFunc(func(o *options) { o.query["modified"] = modified })
}

// WithContent content获取 
func (obj *_EgSmartyCacheMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}


// GetByOption 功能选项模式获取
func (obj *_EgSmartyCacheMgr) GetByOption(opts ...Option) (result EgSmartyCache, err error) {
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
func (obj *_EgSmartyCacheMgr) GetByOptions(opts ...Option) (results []*EgSmartyCache, err error) {
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


// GetFromIDSmartyCache 通过id_smarty_cache获取内容  
func (obj *_EgSmartyCacheMgr)  GetFromIDSmartyCache(idSmartyCache string) (result EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache = ?", idSmartyCache).Find(&result).Error
	
	return
}

// GetBatchFromIDSmartyCache 批量唯一主键查找 
func (obj *_EgSmartyCacheMgr) GetBatchFromIDSmartyCache(idSmartyCaches []string) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache IN (?)", idSmartyCaches).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgSmartyCacheMgr) GetFromName(name string) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgSmartyCacheMgr) GetBatchFromName(names []string) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromCacheID 通过cache_id获取内容  
func (obj *_EgSmartyCacheMgr) GetFromCacheID(cacheID string) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error
	
	return
}

// GetBatchFromCacheID 批量唯一主键查找 
func (obj *_EgSmartyCacheMgr) GetBatchFromCacheID(cacheIDs []string) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id IN (?)", cacheIDs).Find(&results).Error
	
	return
}
 
// GetFromModified 通过modified获取内容  
func (obj *_EgSmartyCacheMgr) GetFromModified(modified time.Time) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified = ?", modified).Find(&results).Error
	
	return
}

// GetBatchFromModified 批量唯一主键查找 
func (obj *_EgSmartyCacheMgr) GetBatchFromModified(modifieds []time.Time) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified IN (?)", modifieds).Find(&results).Error
	
	return
}
 
// GetFromContent 通过content获取内容  
func (obj *_EgSmartyCacheMgr) GetFromContent(content string) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error
	
	return
}

// GetBatchFromContent 批量唯一主键查找 
func (obj *_EgSmartyCacheMgr) GetBatchFromContent(contents []string) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSmartyCacheMgr) FetchByPrimaryKey(idSmartyCache string ) (result EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_smarty_cache = ?", idSmartyCache).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByName  获取多个内容
 func (obj *_EgSmartyCacheMgr) FetchIndexByName(name string ) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}
 
 // FetchIndexByCacheID  获取多个内容
 func (obj *_EgSmartyCacheMgr) FetchIndexByCacheID(cacheID string ) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cache_id = ?", cacheID).Find(&results).Error
	
	return
}
 
 // FetchIndexByModified  获取多个内容
 func (obj *_EgSmartyCacheMgr) FetchIndexByModified(modified time.Time ) (results []*EgSmartyCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("modified = ?", modified).Find(&results).Error
	
	return
}
 

	

