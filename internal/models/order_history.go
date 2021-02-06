package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderHistoryMgr struct {
	*_BaseMgr
}

// OrderHistoryMgr open func
func OrderHistoryMgr(db *gorm.DB) *_OrderHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("OrderHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderHistoryMgr) GetTableName() string {
	return "eg_order_history"
}

// Get 获取
func (obj *_OrderHistoryMgr) Get() (result OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderHistoryMgr) Gets() (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderHistory id_order_history获取
func (obj *_OrderHistoryMgr) WithIDOrderHistory(idOrderHistory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_history"] = idOrderHistory })
}

// WithIDEmployee id_employee获取
func (obj *_OrderHistoryMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDOrder id_order获取
func (obj *_OrderHistoryMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithIDOrderState id_order_state获取
func (obj *_OrderHistoryMgr) WithIDOrderState(idOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_state"] = idOrderState })
}

// WithDateAdd date_add获取
func (obj *_OrderHistoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDOrderHistory 通过id_order_history获取内容
func (obj *_OrderHistoryMgr) GetFromIDOrderHistory(idOrderHistory uint32) (result OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history = ?", idOrderHistory).Find(&result).Error

	return
}

// GetBatchFromIDOrderHistory 批量唯一主键查找
func (obj *_OrderHistoryMgr) GetBatchFromIDOrderHistory(idOrderHistorys []uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history IN (?)", idOrderHistorys).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_OrderHistoryMgr) GetFromIDEmployee(idEmployee uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_OrderHistoryMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_OrderHistoryMgr) GetFromIDOrder(idOrder uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_OrderHistoryMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromIDOrderState 通过id_order_state获取内容
func (obj *_OrderHistoryMgr) GetFromIDOrderState(idOrderState uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&results).Error

	return
}

// GetBatchFromIDOrderState 批量唯一主键查找
func (obj *_OrderHistoryMgr) GetBatchFromIDOrderState(idOrderStates []uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state IN (?)", idOrderStates).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_OrderHistoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_OrderHistoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderHistoryMgr) FetchByPrimaryKey(idOrderHistory uint32) (result OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history = ?", idOrderHistory).Find(&result).Error

	return
}

// FetchIndexByIDEmployee  获取多个内容
func (obj *_OrderHistoryMgr) FetchIndexByIDEmployee(idEmployee uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// FetchIndexByOrderHistoryOrder  获取多个内容
func (obj *_OrderHistoryMgr) FetchIndexByOrderHistoryOrder(idOrder uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// FetchIndexByIDOrderState  获取多个内容
func (obj *_OrderHistoryMgr) FetchIndexByIDOrderState(idOrderState uint32) (results []*OrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&results).Error

	return
}
