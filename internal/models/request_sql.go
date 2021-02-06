package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _RequestSQLMgr struct {
	*_BaseMgr
}

// RequestSQLMgr open func
func RequestSQLMgr(db *gorm.DB) *_RequestSQLMgr {
	if db == nil {
		panic(fmt.Errorf("RequestSQLMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RequestSQLMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_request_sql"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RequestSQLMgr) GetTableName() string {
	return "ps_request_sql"
}

// Get 获取
func (obj *_RequestSQLMgr) Get() (result RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RequestSQLMgr) Gets() (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDRequestSQL id_request_sql获取
func (obj *_RequestSQLMgr) WithIDRequestSQL(idRequestSQL int) Option {
	return optionFunc(func(o *options) { o.query["id_request_sql"] = idRequestSQL })
}

// WithName name获取
func (obj *_RequestSQLMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSQL sql获取
func (obj *_RequestSQLMgr) WithSQL(sql string) Option {
	return optionFunc(func(o *options) { o.query["sql"] = sql })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDRequestSQL 通过id_request_sql获取内容
func (obj *_RequestSQLMgr) GetFromIDRequestSQL(idRequestSQL int) (result RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_request_sql = ?", idRequestSQL).Find(&result).Error

	return
}

// GetBatchFromIDRequestSQL 批量唯一主键查找
func (obj *_RequestSQLMgr) GetBatchFromIDRequestSQL(idRequestSQLs []int) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_request_sql IN (?)", idRequestSQLs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_RequestSQLMgr) GetFromName(name string) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_RequestSQLMgr) GetBatchFromName(names []string) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromSQL 通过sql获取内容
func (obj *_RequestSQLMgr) GetFromSQL(sql string) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sql = ?", sql).Find(&results).Error

	return
}

// GetBatchFromSQL 批量唯一主键查找
func (obj *_RequestSQLMgr) GetBatchFromSQL(sqls []string) (results []*RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sql IN (?)", sqls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_RequestSQLMgr) FetchByPrimaryKey(idRequestSQL int) (result RequestSQL, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_request_sql = ?", idRequestSQL).Find(&result).Error

	return
}
