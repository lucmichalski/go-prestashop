package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderReturnStateMgr struct {
	*_BaseMgr
}

func OrderReturnStateMgr(db *gorm.DB) *_OrderReturnStateMgr {
	if db == nil {
		panic(fmt.Errorf("OrderReturnStateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderReturnStateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_return_state"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderReturnStateMgr) GetTableName() string {
	return "ps_order_return_state"
}

func (obj *_OrderReturnStateMgr) Get() (result OrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderReturnStateMgr) Gets() (results []*OrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderReturnStateMgr) WithIDOrderReturnState(idOrderReturnState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return_state"] = idOrderReturnState })
}

func (obj *_OrderReturnStateMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

func (obj *_OrderReturnStateMgr) GetByOption(opts ...Option) (result OrderReturnState, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderReturnStateMgr) GetByOptions(opts ...Option) (results []*OrderReturnState, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_OrderReturnStateMgr) GetFromIDOrderReturnState(idOrderReturnState uint32) (result OrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ?", idOrderReturnState).Find(&result).Error

	return
}

func (obj *_OrderReturnStateMgr) GetBatchFromIDOrderReturnState(idOrderReturnStates []uint32) (results []*OrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state IN (?)", idOrderReturnStates).Find(&results).Error

	return
}

func (obj *_OrderReturnStateMgr) GetFromColor(color string) (results []*OrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error

	return
}

func (obj *_OrderReturnStateMgr) GetBatchFromColor(colors []string) (results []*OrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error

	return
}


func (obj *_OrderReturnStateMgr) FetchByPrimaryKey(idOrderReturnState uint32) (result OrderReturnState, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ?", idOrderReturnState).Find(&result).Error

	return
}
