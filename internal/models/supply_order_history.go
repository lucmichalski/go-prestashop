package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SupplyOrderHistoryMgr struct {
	*_BaseMgr
}

// SupplyOrderHistoryMgr open func
func SupplyOrderHistoryMgr(db *gorm.DB) *_SupplyOrderHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("SupplyOrderHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplyOrderHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supply_order_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SupplyOrderHistoryMgr) GetTableName() string {
	return "ps_supply_order_history"
}

// Get 获取
func (obj *_SupplyOrderHistoryMgr) Get() (result SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SupplyOrderHistoryMgr) Gets() (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplyOrderHistory id_supply_order_history获取
func (obj *_SupplyOrderHistoryMgr) WithIDSupplyOrderHistory(idSupplyOrderHistory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_history"] = idSupplyOrderHistory })
}

// WithIDSupplyOrder id_supply_order获取
func (obj *_SupplyOrderHistoryMgr) WithIDSupplyOrder(idSupplyOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

// WithIDEmployee id_employee获取
func (obj *_SupplyOrderHistoryMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithEmployeeLastname employee_lastname获取
func (obj *_SupplyOrderHistoryMgr) WithEmployeeLastname(employeeLastname string) Option {
	return optionFunc(func(o *options) { o.query["employee_lastname"] = employeeLastname })
}

// WithEmployeeFirstname employee_firstname获取
func (obj *_SupplyOrderHistoryMgr) WithEmployeeFirstname(employeeFirstname string) Option {
	return optionFunc(func(o *options) { o.query["employee_firstname"] = employeeFirstname })
}

// WithIDState id_state获取
func (obj *_SupplyOrderHistoryMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

// WithDateAdd date_add获取
func (obj *_SupplyOrderHistoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDSupplyOrderHistory 通过id_supply_order_history获取内容
func (obj *_SupplyOrderHistoryMgr) GetFromIDSupplyOrderHistory(idSupplyOrderHistory uint32) (result SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history = ?", idSupplyOrderHistory).Find(&result).Error

	return
}

// GetBatchFromIDSupplyOrderHistory 批量唯一主键查找
func (obj *_SupplyOrderHistoryMgr) GetBatchFromIDSupplyOrderHistory(idSupplyOrderHistorys []uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history IN (?)", idSupplyOrderHistorys).Find(&results).Error

	return
}

// GetFromIDSupplyOrder 通过id_supply_order获取内容
func (obj *_SupplyOrderHistoryMgr) GetFromIDSupplyOrder(idSupplyOrder uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error

	return
}

// GetBatchFromIDSupplyOrder 批量唯一主键查找
func (obj *_SupplyOrderHistoryMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_SupplyOrderHistoryMgr) GetFromIDEmployee(idEmployee uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_SupplyOrderHistoryMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromEmployeeLastname 通过employee_lastname获取内容
func (obj *_SupplyOrderHistoryMgr) GetFromEmployeeLastname(employeeLastname string) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname = ?", employeeLastname).Find(&results).Error

	return
}

// GetBatchFromEmployeeLastname 批量唯一主键查找
func (obj *_SupplyOrderHistoryMgr) GetBatchFromEmployeeLastname(employeeLastnames []string) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname IN (?)", employeeLastnames).Find(&results).Error

	return
}

// GetFromEmployeeFirstname 通过employee_firstname获取内容
func (obj *_SupplyOrderHistoryMgr) GetFromEmployeeFirstname(employeeFirstname string) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname = ?", employeeFirstname).Find(&results).Error

	return
}

// GetBatchFromEmployeeFirstname 批量唯一主键查找
func (obj *_SupplyOrderHistoryMgr) GetBatchFromEmployeeFirstname(employeeFirstnames []string) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname IN (?)", employeeFirstnames).Find(&results).Error

	return
}

// GetFromIDState 通过id_state获取内容
func (obj *_SupplyOrderHistoryMgr) GetFromIDState(idState uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

// GetBatchFromIDState 批量唯一主键查找
func (obj *_SupplyOrderHistoryMgr) GetBatchFromIDState(idStates []uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_SupplyOrderHistoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_SupplyOrderHistoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SupplyOrderHistoryMgr) FetchByPrimaryKey(idSupplyOrderHistory uint32) (result SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history = ?", idSupplyOrderHistory).Find(&result).Error

	return
}

// FetchIndexByIDSupplyOrder  获取多个内容
func (obj *_SupplyOrderHistoryMgr) FetchIndexByIDSupplyOrder(idSupplyOrder uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error

	return
}

// FetchIndexByIDEmployee  获取多个内容
func (obj *_SupplyOrderHistoryMgr) FetchIndexByIDEmployee(idEmployee uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// FetchIndexByIDState  获取多个内容
func (obj *_SupplyOrderHistoryMgr) FetchIndexByIDState(idState uint32) (results []*SupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}
