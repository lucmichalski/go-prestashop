package	model	
import (	
"context"	
"gorm.io/gorm"	
"time"	
"fmt"	
)	

type _EgSupplyOrderHistoryMgr struct {
	*_BaseMgr
}

// EgSupplyOrderHistoryMgr open func
func EgSupplyOrderHistoryMgr(db *gorm.DB) *_EgSupplyOrderHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("EgSupplyOrderHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSupplyOrderHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_supply_order_history"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSupplyOrderHistoryMgr) GetTableName() string {
	return "eg_supply_order_history"
}

// Get 获取
func (obj *_EgSupplyOrderHistoryMgr) Get() (result EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSupplyOrderHistoryMgr) Gets() (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplyOrderHistory id_supply_order_history获取 
func (obj *_EgSupplyOrderHistoryMgr) WithIDSupplyOrderHistory(idSupplyOrderHistory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_history"] = idSupplyOrderHistory })
}

// WithIDSupplyOrder id_supply_order获取 
func (obj *_EgSupplyOrderHistoryMgr) WithIDSupplyOrder(idSupplyOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

// WithIDEmployee id_employee获取 
func (obj *_EgSupplyOrderHistoryMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithEmployeeLastname employee_lastname获取 
func (obj *_EgSupplyOrderHistoryMgr) WithEmployeeLastname(employeeLastname string) Option {
	return optionFunc(func(o *options) { o.query["employee_lastname"] = employeeLastname })
}

// WithEmployeeFirstname employee_firstname获取 
func (obj *_EgSupplyOrderHistoryMgr) WithEmployeeFirstname(employeeFirstname string) Option {
	return optionFunc(func(o *options) { o.query["employee_firstname"] = employeeFirstname })
}

// WithIDState id_state获取 
func (obj *_EgSupplyOrderHistoryMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

// WithDateAdd date_add获取 
func (obj *_EgSupplyOrderHistoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgSupplyOrderHistoryMgr) GetByOption(opts ...Option) (result EgSupplyOrderHistory, err error) {
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
func (obj *_EgSupplyOrderHistoryMgr) GetByOptions(opts ...Option) (results []*EgSupplyOrderHistory, err error) {
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
func (obj *_EgSupplyOrderHistoryMgr)  GetFromIDSupplyOrderHistory(idSupplyOrderHistory uint32) (result EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history = ?", idSupplyOrderHistory).Find(&result).Error
	
	return
}

// GetBatchFromIDSupplyOrderHistory 批量唯一主键查找 
func (obj *_EgSupplyOrderHistoryMgr) GetBatchFromIDSupplyOrderHistory(idSupplyOrderHistorys []uint32) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history IN (?)", idSupplyOrderHistorys).Find(&results).Error
	
	return
}
 
// GetFromIDSupplyOrder 通过id_supply_order获取内容  
func (obj *_EgSupplyOrderHistoryMgr) GetFromIDSupplyOrder(idSupplyOrder uint32) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplyOrder 批量唯一主键查找 
func (obj *_EgSupplyOrderHistoryMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []uint32) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error
	
	return
}
 
// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgSupplyOrderHistoryMgr) GetFromIDEmployee(idEmployee uint32) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgSupplyOrderHistoryMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromEmployeeLastname 通过employee_lastname获取内容  
func (obj *_EgSupplyOrderHistoryMgr) GetFromEmployeeLastname(employeeLastname string) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname = ?", employeeLastname).Find(&results).Error
	
	return
}

// GetBatchFromEmployeeLastname 批量唯一主键查找 
func (obj *_EgSupplyOrderHistoryMgr) GetBatchFromEmployeeLastname(employeeLastnames []string) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_lastname IN (?)", employeeLastnames).Find(&results).Error
	
	return
}
 
// GetFromEmployeeFirstname 通过employee_firstname获取内容  
func (obj *_EgSupplyOrderHistoryMgr) GetFromEmployeeFirstname(employeeFirstname string) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname = ?", employeeFirstname).Find(&results).Error
	
	return
}

// GetBatchFromEmployeeFirstname 批量唯一主键查找 
func (obj *_EgSupplyOrderHistoryMgr) GetBatchFromEmployeeFirstname(employeeFirstnames []string) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee_firstname IN (?)", employeeFirstnames).Find(&results).Error
	
	return
}
 
// GetFromIDState 通过id_state获取内容  
func (obj *_EgSupplyOrderHistoryMgr) GetFromIDState(idState uint32) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error
	
	return
}

// GetBatchFromIDState 批量唯一主键查找 
func (obj *_EgSupplyOrderHistoryMgr) GetBatchFromIDState(idStates []uint32) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgSupplyOrderHistoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgSupplyOrderHistoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSupplyOrderHistoryMgr) FetchByPrimaryKey(idSupplyOrderHistory uint32 ) (result EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_history = ?", idSupplyOrderHistory).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDSupplyOrder  获取多个内容
 func (obj *_EgSupplyOrderHistoryMgr) FetchIndexByIDSupplyOrder(idSupplyOrder uint32 ) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDEmployee  获取多个内容
 func (obj *_EgSupplyOrderHistoryMgr) FetchIndexByIDEmployee(idEmployee uint32 ) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDState  获取多个内容
 func (obj *_EgSupplyOrderHistoryMgr) FetchIndexByIDState(idState uint32 ) (results []*EgSupplyOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error
	
	return
}
 

	

