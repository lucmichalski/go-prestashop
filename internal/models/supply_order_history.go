package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SupplyOrderHistoryMgr struct {
	*_BaseMgr
}

func SupplyOrderHistoryMgr(db *gorm.DB) *_SupplyOrderHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("SupplyOrderHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplyOrderHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supply_order_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SupplyOrderHistoryMgr) GetTableName() string {
	return "ps_supply_order_history"
}

func (obj *_SupplyOrderHistoryMgr) Get() (result SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) Gets() (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SupplyOrderHistoryMgr) WithIDSupplyOrderHistory(idSupplyOrderHistory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_history"] = idSupplyOrderHistory })
}

func (obj *_SupplyOrderHistoryMgr) WithIDSupplyOrder(idSupplyOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

func (obj *_SupplyOrderHistoryMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_SupplyOrderHistoryMgr) WithEmployeeLastname(employeeLastname string) Option {
	return optionFunc(func(o *options) { o.query["employee_lastname"] = employeeLastname })
}

func (obj *_SupplyOrderHistoryMgr) WithEmployeeFirstname(employeeFirstname string) Option {
	return optionFunc(func(o *options) { o.query["employee_firstname"] = employeeFirstname })
}

func (obj *_SupplyOrderHistoryMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

func (obj *_SupplyOrderHistoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_SupplyOrderHistoryMgr) GetByOption(opts ...Option) (result SupplyOrderHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetByOptions(opts ...Option) (results []*SupplyOrderHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SupplyOrderHistoryMgr) GetFromIDSupplyOrderHistory(idSupplyOrderHistory uint32) (result SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history = ?", idSupplyOrderHistory).Find(&result).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetBatchFromIDSupplyOrderHistory(idSupplyOrderHistorys []uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history IN (?)", idSupplyOrderHistorys).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetFromIDSupplyOrder(idSupplyOrder uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetFromIDEmployee(idEmployee uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetFromEmployeeLastname(employeeLastname string) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname = ?", employeeLastname).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetBatchFromEmployeeLastname(employeeLastnames []string) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname IN (?)", employeeLastnames).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetFromEmployeeFirstname(employeeFirstname string) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname = ?", employeeFirstname).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetBatchFromEmployeeFirstname(employeeFirstnames []string) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname IN (?)", employeeFirstnames).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetFromIDState(idState uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetBatchFromIDState(idStates []uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_SupplyOrderHistoryMgr) FetchByPrimaryKey(idSupplyOrderHistory uint32) (result SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history = ?", idSupplyOrderHistory).Find(&result).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) FetchIndexByIDSupplyOrder(idSupplyOrder uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) FetchIndexByIDEmployee(idEmployee uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_SupplyOrderHistoryMgr) FetchIndexByIDState(idState uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}
