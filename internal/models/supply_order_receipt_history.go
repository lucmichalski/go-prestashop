package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SupplyOrderReceiptHistoryMgr struct {
	*_BaseMgr
}

func SupplyOrderReceiptHistoryMgr(db *gorm.DB) *_SupplyOrderReceiptHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("SupplyOrderReceiptHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplyOrderReceiptHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supply_order_receipt_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetTableName() string {
	return "ps_supply_order_receipt_history"
}

func (obj *_SupplyOrderReceiptHistoryMgr) Get() (result SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) Gets() (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SupplyOrderReceiptHistoryMgr) WithIDSupplyOrderReceiptHistory(idSupplyOrderReceiptHistory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_receipt_history"] = idSupplyOrderReceiptHistory })
}

func (obj *_SupplyOrderReceiptHistoryMgr) WithIDSupplyOrderDetail(idSupplyOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_detail"] = idSupplyOrderDetail })
}

func (obj *_SupplyOrderReceiptHistoryMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_SupplyOrderReceiptHistoryMgr) WithEmployeeLastname(employeeLastname string) Option {
	return optionFunc(func(o *options) { o.query["employee_lastname"] = employeeLastname })
}

func (obj *_SupplyOrderReceiptHistoryMgr) WithEmployeeFirstname(employeeFirstname string) Option {
	return optionFunc(func(o *options) { o.query["employee_firstname"] = employeeFirstname })
}

func (obj *_SupplyOrderReceiptHistoryMgr) WithIDSupplyOrderState(idSupplyOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_state"] = idSupplyOrderState })
}

func (obj *_SupplyOrderReceiptHistoryMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

func (obj *_SupplyOrderReceiptHistoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetByOption(opts ...Option) (result SupplyOrderReceiptHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetByOptions(opts ...Option) (results []*SupplyOrderReceiptHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SupplyOrderReceiptHistoryMgr) GetFromIDSupplyOrderReceiptHistory(idSupplyOrderReceiptHistory uint32) (result SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_receipt_history = ?", idSupplyOrderReceiptHistory).Find(&result).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetBatchFromIDSupplyOrderReceiptHistory(idSupplyOrderReceiptHistorys []uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_receipt_history IN (?)", idSupplyOrderReceiptHistorys).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetFromIDSupplyOrderDetail(idSupplyOrderDetail uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail = ?", idSupplyOrderDetail).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetBatchFromIDSupplyOrderDetail(idSupplyOrderDetails []uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail IN (?)", idSupplyOrderDetails).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetFromIDEmployee(idEmployee uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetFromEmployeeLastname(employeeLastname string) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname = ?", employeeLastname).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetBatchFromEmployeeLastname(employeeLastnames []string) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname IN (?)", employeeLastnames).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetFromEmployeeFirstname(employeeFirstname string) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname = ?", employeeFirstname).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetBatchFromEmployeeFirstname(employeeFirstnames []string) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname IN (?)", employeeFirstnames).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetFromIDSupplyOrderState(idSupplyOrderState uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetBatchFromIDSupplyOrderState(idSupplyOrderStates []uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state IN (?)", idSupplyOrderStates).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetFromQuantity(quantity uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetBatchFromQuantity(quantitys []uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_SupplyOrderReceiptHistoryMgr) FetchByPrimaryKey(idSupplyOrderReceiptHistory uint32) (result SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_receipt_history = ?", idSupplyOrderReceiptHistory).Find(&result).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) FetchIndexByIDSupplyOrderDetail(idSupplyOrderDetail uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail = ?", idSupplyOrderDetail).Find(&results).Error

	return
}

func (obj *_SupplyOrderReceiptHistoryMgr) FetchIndexByIDSupplyOrderState(idSupplyOrderState uint32) (results []*SupplyOrderReceiptHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&results).Error

	return
}
