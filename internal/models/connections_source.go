package	model	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _EgConnectionsSourceMgr struct {
	*_BaseMgr
}

// EgConnectionsSourceMgr open func
func EgConnectionsSourceMgr(db *gorm.DB) *_EgConnectionsSourceMgr {
	if db == nil {
		panic(fmt.Errorf("EgConnectionsSourceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgConnectionsSourceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_connections_source"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgConnectionsSourceMgr) GetTableName() string {
	return "eg_connections_source"
}

// Get 获取
func (obj *_EgConnectionsSourceMgr) Get() (result EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgConnectionsSourceMgr) Gets() (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConnectionsSource id_connections_source获取 
func (obj *_EgConnectionsSourceMgr) WithIDConnectionsSource(idConnectionsSource uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections_source"] = idConnectionsSource })
}

// WithIDConnections id_connections获取 
func (obj *_EgConnectionsSourceMgr) WithIDConnections(idConnections uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections"] = idConnections })
}

// WithHTTPReferer http_referer获取 
func (obj *_EgConnectionsSourceMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

// WithRequestURI request_uri获取 
func (obj *_EgConnectionsSourceMgr) WithRequestURI(requestURI string) Option {
	return optionFunc(func(o *options) { o.query["request_uri"] = requestURI })
}

// WithKeywords keywords获取 
func (obj *_EgConnectionsSourceMgr) WithKeywords(keywords string) Option {
	return optionFunc(func(o *options) { o.query["keywords"] = keywords })
}

// WithDateAdd date_add获取 
func (obj *_EgConnectionsSourceMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgConnectionsSourceMgr) GetByOption(opts ...Option) (result EgConnectionsSource, err error) {
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
func (obj *_EgConnectionsSourceMgr) GetByOptions(opts ...Option) (results []*EgConnectionsSource, err error) {
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
func (obj *_EgConnectionsSourceMgr)  GetFromIDConnectionsSource(idConnectionsSource uint32) (result EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ?", idConnectionsSource).Find(&result).Error
	
	return
}

// GetBatchFromIDConnectionsSource 批量唯一主键查找 
func (obj *_EgConnectionsSourceMgr) GetBatchFromIDConnectionsSource(idConnectionsSources []uint32) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source IN (?)", idConnectionsSources).Find(&results).Error
	
	return
}
 
// GetFromIDConnections 通过id_connections获取内容  
func (obj *_EgConnectionsSourceMgr) GetFromIDConnections(idConnections uint32) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&results).Error
	
	return
}

// GetBatchFromIDConnections 批量唯一主键查找 
func (obj *_EgConnectionsSourceMgr) GetBatchFromIDConnections(idConnectionss []uint32) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections IN (?)", idConnectionss).Find(&results).Error
	
	return
}
 
// GetFromHTTPReferer 通过http_referer获取内容  
func (obj *_EgConnectionsSourceMgr) GetFromHTTPReferer(httpReferer string) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error
	
	return
}

// GetBatchFromHTTPReferer 批量唯一主键查找 
func (obj *_EgConnectionsSourceMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error
	
	return
}
 
// GetFromRequestURI 通过request_uri获取内容  
func (obj *_EgConnectionsSourceMgr) GetFromRequestURI(requestURI string) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri = ?", requestURI).Find(&results).Error
	
	return
}

// GetBatchFromRequestURI 批量唯一主键查找 
func (obj *_EgConnectionsSourceMgr) GetBatchFromRequestURI(requestURIs []string) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri IN (?)", requestURIs).Find(&results).Error
	
	return
}
 
// GetFromKeywords 通过keywords获取内容  
func (obj *_EgConnectionsSourceMgr) GetFromKeywords(keywords string) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords = ?", keywords).Find(&results).Error
	
	return
}

// GetBatchFromKeywords 批量唯一主键查找 
func (obj *_EgConnectionsSourceMgr) GetBatchFromKeywords(keywordss []string) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords IN (?)", keywordss).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgConnectionsSourceMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgConnectionsSourceMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgConnectionsSourceMgr) FetchByPrimaryKey(idConnectionsSource uint32 ) (result EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ?", idConnectionsSource).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByConnections  获取多个内容
 func (obj *_EgConnectionsSourceMgr) FetchIndexByConnections(idConnections uint32 ) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&results).Error
	
	return
}
 
 // FetchIndexByHTTPReferer  获取多个内容
 func (obj *_EgConnectionsSourceMgr) FetchIndexByHTTPReferer(httpReferer string ) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error
	
	return
}
 
 // FetchIndexByRequestURI  获取多个内容
 func (obj *_EgConnectionsSourceMgr) FetchIndexByRequestURI(requestURI string ) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri = ?", requestURI).Find(&results).Error
	
	return
}
 
 // FetchIndexByOrderby  获取多个内容
 func (obj *_EgConnectionsSourceMgr) FetchIndexByOrderby(dateAdd time.Time ) (results []*EgConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}
 

	

