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

func OrderMessageMgr(db *gorm.DB) *_OrderMessageMgr {
	if db == nil {
		panic(fmt.Errorf("OrderMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderMessageMgr) GetTableName() string {
	return "ps_order_message"
}

func (obj *_OrderMessageMgr) Get() (result OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderMessageMgr) Gets() (results []*OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderMessageMgr) WithIDOrderMessage(idOrderMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_message"] = idOrderMessage })
}

func (obj *_OrderMessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

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


func (obj *_OrderMessageMgr) GetFromIDOrderMessage(idOrderMessage uint32) (result OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message = ?", idOrderMessage).Find(&result).Error

	return
}

func (obj *_OrderMessageMgr) GetBatchFromIDOrderMessage(idOrderMessages []uint32) (results []*OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message IN (?)", idOrderMessages).Find(&results).Error

	return
}

func (obj *_OrderMessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_OrderMessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_OrderMessageMgr) FetchByPrimaryKey(idOrderMessage uint32) (result OrderMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message = ?", idOrderMessage).Find(&result).Error

	return
}
