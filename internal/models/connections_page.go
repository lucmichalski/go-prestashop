package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConnectionsPageMgr struct {
	*_BaseMgr
}

// ConnectionsPageMgr open func
func ConnectionsPageMgr(db *gorm.DB) *_ConnectionsPageMgr {
	if db == nil {
		panic(fmt.Errorf("ConnectionsPageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConnectionsPageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_connections_page"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConnectionsPageMgr) GetTableName() string {
	return "ps_connections_page"
}

// Get 获取
func (obj *_ConnectionsPageMgr) Get() (result ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConnectionsPageMgr) Gets() (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConnections id_connections获取
func (obj *_ConnectionsPageMgr) WithIDConnections(idConnections uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections"] = idConnections })
}

// WithIDPage id_page获取
func (obj *_ConnectionsPageMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

// WithTimeStart time_start获取
func (obj *_ConnectionsPageMgr) WithTimeStart(timeStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_start"] = timeStart })
}

// WithTimeEnd time_end获取
func (obj *_ConnectionsPageMgr) WithTimeEnd(timeEnd time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_end"] = timeEnd })
}

// GetByOption 功能选项模式获取
func (obj *_ConnectionsPageMgr) GetByOption(opts ...Option) (result ConnectionsPage, err error) {
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
func (obj *_ConnectionsPageMgr) GetByOptions(opts ...Option) (results []*ConnectionsPage, err error) {
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

// GetFromIDConnections 通过id_connections获取内容
func (obj *_ConnectionsPageMgr) GetFromIDConnections(idConnections uint32) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&results).Error

	return
}

// GetBatchFromIDConnections 批量唯一主键查找
func (obj *_ConnectionsPageMgr) GetBatchFromIDConnections(idConnectionss []uint32) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections IN (?)", idConnectionss).Find(&results).Error

	return
}

// GetFromIDPage 通过id_page获取内容
func (obj *_ConnectionsPageMgr) GetFromIDPage(idPage uint32) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error

	return
}

// GetBatchFromIDPage 批量唯一主键查找
func (obj *_ConnectionsPageMgr) GetBatchFromIDPage(idPages []uint32) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error

	return
}

// GetFromTimeStart 通过time_start获取内容
func (obj *_ConnectionsPageMgr) GetFromTimeStart(timeStart time.Time) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_start = ?", timeStart).Find(&results).Error

	return
}

// GetBatchFromTimeStart 批量唯一主键查找
func (obj *_ConnectionsPageMgr) GetBatchFromTimeStart(timeStarts []time.Time) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_start IN (?)", timeStarts).Find(&results).Error

	return
}

// GetFromTimeEnd 通过time_end获取内容
func (obj *_ConnectionsPageMgr) GetFromTimeEnd(timeEnd time.Time) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_end = ?", timeEnd).Find(&results).Error

	return
}

// GetBatchFromTimeEnd 批量唯一主键查找
func (obj *_ConnectionsPageMgr) GetBatchFromTimeEnd(timeEnds []time.Time) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_end IN (?)", timeEnds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ConnectionsPageMgr) FetchByPrimaryKey(idConnections uint32, idPage uint32, timeStart time.Time) (result ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ? AND id_page = ? AND time_start = ?", idConnections, idPage, timeStart).Find(&result).Error

	return
}
