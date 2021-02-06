package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _RequestSQLMgr struct {
	*_BaseMgr
}

func RequestSQLMgr(db *gorm.DB) *_RequestSQLMgr {
	if db == nil {
		panic(fmt.Errorf("RequestSQLMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RequestSQLMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_request_sql"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_RequestSQLMgr) GetTableName() string {
	return "ps_request_sql"
}

func (obj *_RequestSQLMgr) Get() (result RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_RequestSQLMgr) Gets() (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_RequestSQLMgr) WithIDRequestSQL(idRequestSQL int) Option {
	return optionFunc(func(o *options) { o.query["id_request_sql"] = idRequestSQL })
}

func (obj *_RequestSQLMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_RequestSQLMgr) WithSQL(sql string) Option {
	return optionFunc(func(o *options) { o.query["sql"] = sql })
}

func (obj *_RequestSQLMgr) GetByOption(opts ...Option) (result RequestSQL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_RequestSQLMgr) GetByOptions(opts ...Option) (results []*RequestSQL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_RequestSQLMgr) GetFromIDRequestSQL(idRequestSQL int) (result RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_request_sql = ?", idRequestSQL).Find(&result).Error

	return
}

func (obj *_RequestSQLMgr) GetBatchFromIDRequestSQL(idRequestSQLs []int) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_request_sql IN (?)", idRequestSQLs).Find(&results).Error

	return
}

func (obj *_RequestSQLMgr) GetFromName(name string) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_RequestSQLMgr) GetBatchFromName(names []string) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_RequestSQLMgr) GetFromSQL(sql string) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sql = ?", sql).Find(&results).Error

	return
}

func (obj *_RequestSQLMgr) GetBatchFromSQL(sqls []string) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sql IN (?)", sqls).Find(&results).Error

	return
}


func (obj *_RequestSQLMgr) FetchByPrimaryKey(idRequestSQL int) (result RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_request_sql = ?", idRequestSQL).Find(&result).Error

	return
}
