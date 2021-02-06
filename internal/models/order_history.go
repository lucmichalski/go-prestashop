package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgOrderHistoryMgr struct {
	*_BaseMgr
}

// EgOrderHistoryMgr open func
func EgOrderHistoryMgr(db *gorm.DB) *_EgOrderHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_history"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderHistoryMgr) GetTableName() string {
	return "eg_order_history"
}

// Get 获取
func (obj *_EgOrderHistoryMgr) Get() (result EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderHistoryMgr) Gets() (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderHistory id_order_history获取 
func (obj *_EgOrderHistoryMgr) WithIDOrderHistory(idOrderHistory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_history"] = idOrderHistory })
}

// WithIDEmployee id_employee获取 
func (obj *_EgOrderHistoryMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDOrder id_order获取 
func (obj *_EgOrderHistoryMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithIDOrderState id_order_state获取 
func (obj *_EgOrderHistoryMgr) WithIDOrderState(idOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_state"] = idOrderState })
}

// WithDateAdd date_add获取 
func (obj *_EgOrderHistoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderHistoryMgr) GetByOption(opts ...Option) (result EgOrderHistory, err error) {
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
func (obj *_EgOrderHistoryMgr) GetByOptions(opts ...Option) (results []*EgOrderHistory, err error) {
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
func (obj *_EgOrderHistoryMgr)  GetFromIDOrderHistory(idOrderHistory uint32) (result EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history = ?", idOrderHistory).Find(&result).Error
	
	return
}

// GetBatchFromIDOrderHistory 批量唯一主键查找 
func (obj *_EgOrderHistoryMgr) GetBatchFromIDOrderHistory(idOrderHistorys []uint32) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history IN (?)", idOrderHistorys).Find(&results).Error
	
	return
}
 
// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgOrderHistoryMgr) GetFromIDEmployee(idEmployee uint32) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgOrderHistoryMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromIDOrder 通过id_order获取内容  
func (obj *_EgOrderHistoryMgr) GetFromIDOrder(idOrder uint32) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDOrder 批量唯一主键查找 
func (obj *_EgOrderHistoryMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error
	
	return
}
 
// GetFromIDOrderState 通过id_order_state获取内容  
func (obj *_EgOrderHistoryMgr) GetFromIDOrderState(idOrderState uint32) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderState 批量唯一主键查找 
func (obj *_EgOrderHistoryMgr) GetBatchFromIDOrderState(idOrderStates []uint32) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state IN (?)", idOrderStates).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgOrderHistoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgOrderHistoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderHistoryMgr) FetchByPrimaryKey(idOrderHistory uint32 ) (result EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_history = ?", idOrderHistory).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDEmployee  获取多个内容
 func (obj *_EgOrderHistoryMgr) FetchIndexByIDEmployee(idEmployee uint32 ) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}
 
 // FetchIndexByOrderHistoryOrder  获取多个内容
 func (obj *_EgOrderHistoryMgr) FetchIndexByOrderHistoryOrder(idOrder uint32 ) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDOrderState  获取多个内容
 func (obj *_EgOrderHistoryMgr) FetchIndexByIDOrderState(idOrderState uint32 ) (results []*EgOrderHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&results).Error
	
	return
}
 

	
