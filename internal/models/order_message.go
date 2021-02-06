package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderMessageMgr struct {
	*_BaseMgr
}

// OrderMessageMgr open func
func OrderMessageMgr(db *gorm.DB) *_OrderMessageMgr {
	if db == nil {
		panic(fmt.Errorf("OrderMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderMessageMgr) GetTableName() string {
	return "eg_order_message"
}

// Get 获取
func (obj *_OrderMessageMgr) Get() (result OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderMessageMgr) Gets() (results []*OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderMessage id_order_message获取
func (obj *_OrderMessageMgr) WithIDOrderMessage(idOrderMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_message"] = idOrderMessage })
}

// WithDateAdd date_add获取
func (obj *_OrderMessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
func (obj *_OrderMessageMgr) GetByOption(opts ...Option) (result OrderMessage, err error) {
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
func (obj *_OrderMessageMgr) GetByOptions(opts ...Option) (results []*OrderMessage, err error) {
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

// GetFromIDOrderMessage 通过id_order_message获取内容
func (obj *_OrderMessageMgr) GetFromIDOrderMessage(idOrderMessage uint32) (result OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message = ?", idOrderMessage).Find(&result).Error

	return
}

// GetBatchFromIDOrderMessage 批量唯一主键查找
func (obj *_OrderMessageMgr) GetBatchFromIDOrderMessage(idOrderMessages []uint32) (results []*OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message IN (?)", idOrderMessages).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_OrderMessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_OrderMessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderMessageMgr) FetchByPrimaryKey(idOrderMessage uint32) (result OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message = ?", idOrderMessage).Find(&result).Error

	return
}
