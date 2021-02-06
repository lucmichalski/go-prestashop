package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderReturnMgr struct {
	*_BaseMgr
}

// OrderReturnMgr open func
func OrderReturnMgr(db *gorm.DB) *_OrderReturnMgr {
	if db == nil {
		panic(fmt.Errorf("OrderReturnMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderReturnMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_return"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderReturnMgr) GetTableName() string {
	return "eg_order_return"
}

// Get 获取
func (obj *_OrderReturnMgr) Get() (result OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderReturnMgr) Gets() (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderReturn id_order_return获取
func (obj *_OrderReturnMgr) WithIDOrderReturn(idOrderReturn uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return"] = idOrderReturn })
}

// WithIDCustomer id_customer获取
func (obj *_OrderReturnMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDOrder id_order获取
func (obj *_OrderReturnMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithState state获取
func (obj *_OrderReturnMgr) WithState(state bool) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithQuestion question获取
func (obj *_OrderReturnMgr) WithQuestion(question string) Option {
	return optionFunc(func(o *options) { o.query["question"] = question })
}

// WithDateAdd date_add获取
func (obj *_OrderReturnMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_OrderReturnMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_OrderReturnMgr) GetByOption(opts ...Option) (result OrderReturn, err error) {
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
func (obj *_OrderReturnMgr) GetByOptions(opts ...Option) (results []*OrderReturn, err error) {
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

// GetFromIDOrderReturn 通过id_order_return获取内容
func (obj *_OrderReturnMgr) GetFromIDOrderReturn(idOrderReturn uint32) (result OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return = ?", idOrderReturn).Find(&result).Error

	return
}

// GetBatchFromIDOrderReturn 批量唯一主键查找
func (obj *_OrderReturnMgr) GetBatchFromIDOrderReturn(idOrderReturns []uint32) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return IN (?)", idOrderReturns).Find(&results).Error

	return
}

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_OrderReturnMgr) GetFromIDCustomer(idCustomer uint32) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_OrderReturnMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_OrderReturnMgr) GetFromIDOrder(idOrder uint32) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_OrderReturnMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromState 通过state获取内容
func (obj *_OrderReturnMgr) GetFromState(state bool) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量唯一主键查找
func (obj *_OrderReturnMgr) GetBatchFromState(states []bool) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("state IN (?)", states).Find(&results).Error

	return
}

// GetFromQuestion 通过question获取内容
func (obj *_OrderReturnMgr) GetFromQuestion(question string) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("question = ?", question).Find(&results).Error

	return
}

// GetBatchFromQuestion 批量唯一主键查找
func (obj *_OrderReturnMgr) GetBatchFromQuestion(questions []string) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("question IN (?)", questions).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_OrderReturnMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_OrderReturnMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_OrderReturnMgr) GetFromDateUpd(dateUpd time.Time) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_OrderReturnMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderReturnMgr) FetchByPrimaryKey(idOrderReturn uint32) (result OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return = ?", idOrderReturn).Find(&result).Error

	return
}

// FetchIndexByOrderReturnCustomer  获取多个内容
func (obj *_OrderReturnMgr) FetchIndexByOrderReturnCustomer(idCustomer uint32) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// FetchIndexByIDOrder  获取多个内容
func (obj *_OrderReturnMgr) FetchIndexByIDOrder(idOrder uint32) (results []*OrderReturn, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}
