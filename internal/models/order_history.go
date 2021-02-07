package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderHistoryMgr struct {
	*_BaseMgr
}

func OrderHistoryMgr(db *gorm.DB) *_OrderHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("OrderHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderHistoryMgr) GetTableName() string {
	return "ps_order_history"
}

func (obj *_OrderHistoryMgr) Get() (result OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderHistoryMgr) Gets() (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) WithIDOrderHistory(idOrderHistory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_history"] = idOrderHistory })
}

func (obj *_OrderHistoryMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_OrderHistoryMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

func (obj *_OrderHistoryMgr) WithIDOrderState(idOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_state"] = idOrderState })
}

func (obj *_OrderHistoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_OrderHistoryMgr) GetByOption(opts ...Option) (result OrderHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderHistoryMgr) GetByOptions(opts ...Option) (results []*OrderHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetFromIDOrderHistory(idOrderHistory uint32) (result OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history = ?", idOrderHistory).Find(&result).Error

	return
}

func (obj *_OrderHistoryMgr) GetBatchFromIDOrderHistory(idOrderHistorys []uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history IN (?)", idOrderHistorys).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetFromIDEmployee(idEmployee uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetFromIDOrder(idOrder uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetFromIDOrderState(idOrderState uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetBatchFromIDOrderState(idOrderStates []uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state IN (?)", idOrderStates).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) FetchByPrimaryKey(idOrderHistory uint32) (result OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history = ?", idOrderHistory).Find(&result).Error

	return
}

func (obj *_OrderHistoryMgr) FetchIndexByIDEmployee(idEmployee uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) FetchIndexByOrderHistoryOrder(idOrder uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_OrderHistoryMgr) FetchIndexByIDOrderState(idOrderState uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&results).Error

	return
}
