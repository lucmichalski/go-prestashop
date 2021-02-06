package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConnectionsSourceMgr struct {
	*_BaseMgr
}

// ConnectionsSourceMgr open func
func ConnectionsSourceMgr(db *gorm.DB) *_ConnectionsSourceMgr {
	if db == nil {
		panic(fmt.Errorf("ConnectionsSourceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConnectionsSourceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_connections_source"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConnectionsSourceMgr) GetTableName() string {
	return "eg_connections_source"
}

// Get 获取
func (obj *_ConnectionsSourceMgr) Get() (result ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConnectionsSourceMgr) Gets() (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConnectionsSource id_connections_source获取
func (obj *_ConnectionsSourceMgr) WithIDConnectionsSource(idConnectionsSource uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections_source"] = idConnectionsSource })
}

// WithIDConnections id_connections获取
func (obj *_ConnectionsSourceMgr) WithIDConnections(idConnections uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections"] = idConnections })
}

// WithHTTPReferer http_referer获取
func (obj *_ConnectionsSourceMgr) WithHTTPReferer(httpReferer string) Option {
	return optionFunc(func(o *options) { o.query["http_referer"] = httpReferer })
}

// WithRequestURI request_uri获取
func (obj *_ConnectionsSourceMgr) WithRequestURI(requestURI string) Option {
	return optionFunc(func(o *options) { o.query["request_uri"] = requestURI })
}

// WithKeywords keywords获取
func (obj *_ConnectionsSourceMgr) WithKeywords(keywords string) Option {
	return optionFunc(func(o *options) { o.query["keywords"] = keywords })
}

// WithDateAdd date_add获取
func (obj *_ConnectionsSourceMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
func (obj *_ConnectionsSourceMgr) GetByOption(opts ...Option) (result ConnectionsSource, err error) {
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
func (obj *_ConnectionsSourceMgr) GetByOptions(opts ...Option) (results []*ConnectionsSource, err error) {
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
func (obj *_ConnectionsSourceMgr) GetFromIDConnectionsSource(idConnectionsSource uint32) (result ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ?", idConnectionsSource).Find(&result).Error

	return
}

// GetBatchFromIDConnectionsSource 批量唯一主键查找
func (obj *_ConnectionsSourceMgr) GetBatchFromIDConnectionsSource(idConnectionsSources []uint32) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source IN (?)", idConnectionsSources).Find(&results).Error

	return
}

// GetFromIDConnections 通过id_connections获取内容
func (obj *_ConnectionsSourceMgr) GetFromIDConnections(idConnections uint32) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&results).Error

	return
}

// GetBatchFromIDConnections 批量唯一主键查找
func (obj *_ConnectionsSourceMgr) GetBatchFromIDConnections(idConnectionss []uint32) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections IN (?)", idConnectionss).Find(&results).Error

	return
}

// GetFromHTTPReferer 通过http_referer获取内容
func (obj *_ConnectionsSourceMgr) GetFromHTTPReferer(httpReferer string) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error

	return
}

// GetBatchFromHTTPReferer 批量唯一主键查找
func (obj *_ConnectionsSourceMgr) GetBatchFromHTTPReferer(httpReferers []string) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer IN (?)", httpReferers).Find(&results).Error

	return
}

// GetFromRequestURI 通过request_uri获取内容
func (obj *_ConnectionsSourceMgr) GetFromRequestURI(requestURI string) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri = ?", requestURI).Find(&results).Error

	return
}

// GetBatchFromRequestURI 批量唯一主键查找
func (obj *_ConnectionsSourceMgr) GetBatchFromRequestURI(requestURIs []string) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri IN (?)", requestURIs).Find(&results).Error

	return
}

// GetFromKeywords 通过keywords获取内容
func (obj *_ConnectionsSourceMgr) GetFromKeywords(keywords string) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords = ?", keywords).Find(&results).Error

	return
}

// GetBatchFromKeywords 批量唯一主键查找
func (obj *_ConnectionsSourceMgr) GetBatchFromKeywords(keywordss []string) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords IN (?)", keywordss).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_ConnectionsSourceMgr) GetFromDateAdd(dateAdd time.Time) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_ConnectionsSourceMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ConnectionsSourceMgr) FetchByPrimaryKey(idConnectionsSource uint32) (result ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections_source = ?", idConnectionsSource).Find(&result).Error

	return
}

// FetchIndexByConnections  获取多个内容
func (obj *_ConnectionsSourceMgr) FetchIndexByConnections(idConnections uint32) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&results).Error

	return
}

// FetchIndexByHTTPReferer  获取多个内容
func (obj *_ConnectionsSourceMgr) FetchIndexByHTTPReferer(httpReferer string) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer = ?", httpReferer).Find(&results).Error

	return
}

// FetchIndexByRequestURI  获取多个内容
func (obj *_ConnectionsSourceMgr) FetchIndexByRequestURI(requestURI string) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri = ?", requestURI).Find(&results).Error

	return
}

// FetchIndexByOrderby  获取多个内容
func (obj *_ConnectionsSourceMgr) FetchIndexByOrderby(dateAdd time.Time) (results []*ConnectionsSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}
