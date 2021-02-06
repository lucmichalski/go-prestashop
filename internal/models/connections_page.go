package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConnectionsPageMgr struct {
	*_BaseMgr
}

func ConnectionsPageMgr(db *gorm.DB) *_ConnectionsPageMgr {
	if db == nil {
		panic(fmt.Errorf("ConnectionsPageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConnectionsPageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_connections_page"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ConnectionsPageMgr) GetTableName() string {
	return "ps_connections_page"
}

func (obj *_ConnectionsPageMgr) Get() (result ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ConnectionsPageMgr) Gets() (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ConnectionsPageMgr) WithIDConnections(idConnections uint32) Option {
	return optionFunc(func(o *options) { o.query["id_connections"] = idConnections })
}

func (obj *_ConnectionsPageMgr) WithIDPage(idPage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_page"] = idPage })
}

func (obj *_ConnectionsPageMgr) WithTimeStart(timeStart time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_start"] = timeStart })
}

func (obj *_ConnectionsPageMgr) WithTimeEnd(timeEnd time.Time) Option {
	return optionFunc(func(o *options) { o.query["time_end"] = timeEnd })
}

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


func (obj *_ConnectionsPageMgr) GetFromIDConnections(idConnections uint32) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ?", idConnections).Find(&results).Error

	return
}

func (obj *_ConnectionsPageMgr) GetBatchFromIDConnections(idConnectionss []uint32) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections IN (?)", idConnectionss).Find(&results).Error

	return
}

func (obj *_ConnectionsPageMgr) GetFromIDPage(idPage uint32) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page = ?", idPage).Find(&results).Error

	return
}

func (obj *_ConnectionsPageMgr) GetBatchFromIDPage(idPages []uint32) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_page IN (?)", idPages).Find(&results).Error

	return
}

func (obj *_ConnectionsPageMgr) GetFromTimeStart(timeStart time.Time) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_start = ?", timeStart).Find(&results).Error

	return
}

func (obj *_ConnectionsPageMgr) GetBatchFromTimeStart(timeStarts []time.Time) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_start IN (?)", timeStarts).Find(&results).Error

	return
}

func (obj *_ConnectionsPageMgr) GetFromTimeEnd(timeEnd time.Time) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_end = ?", timeEnd).Find(&results).Error

	return
}

func (obj *_ConnectionsPageMgr) GetBatchFromTimeEnd(timeEnds []time.Time) (results []*ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_end IN (?)", timeEnds).Find(&results).Error

	return
}


func (obj *_ConnectionsPageMgr) FetchByPrimaryKey(idConnections uint32, idPage uint32, timeStart time.Time) (result ConnectionsPage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_connections = ? AND id_page = ? AND time_start = ?", idConnections, idPage, timeStart).Find(&result).Error

	return
}
