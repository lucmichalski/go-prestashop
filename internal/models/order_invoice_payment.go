package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOrderInvoicePaymentMgr struct {
	*_BaseMgr
}

// EgOrderInvoicePaymentMgr open func
func EgOrderInvoicePaymentMgr(db *gorm.DB) *_EgOrderInvoicePaymentMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderInvoicePaymentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderInvoicePaymentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_invoice_payment"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderInvoicePaymentMgr) GetTableName() string {
	return "eg_order_invoice_payment"
}

// Get 获取
func (obj *_EgOrderInvoicePaymentMgr) Get() (result EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderInvoicePaymentMgr) Gets() (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderInvoice id_order_invoice获取 
func (obj *_EgOrderInvoicePaymentMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

// WithIDOrderPayment id_order_payment获取 
func (obj *_EgOrderInvoicePaymentMgr) WithIDOrderPayment(idOrderPayment uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_payment"] = idOrderPayment })
}

// WithIDOrder id_order获取 
func (obj *_EgOrderInvoicePaymentMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderInvoicePaymentMgr) GetByOption(opts ...Option) (result EgOrderInvoicePayment, err error) {
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
func (obj *_EgOrderInvoicePaymentMgr) GetByOptions(opts ...Option) (results []*EgOrderInvoicePayment, err error) {
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


// GetFromIDOrderInvoice 通过id_order_invoice获取内容  
func (obj *_EgOrderInvoicePaymentMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderInvoice 批量唯一主键查找 
func (obj *_EgOrderInvoicePaymentMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error
	
	return
}
 
// GetFromIDOrderPayment 通过id_order_payment获取内容  
func (obj *_EgOrderInvoicePaymentMgr) GetFromIDOrderPayment(idOrderPayment uint32) (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment = ?", idOrderPayment).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderPayment 批量唯一主键查找 
func (obj *_EgOrderInvoicePaymentMgr) GetBatchFromIDOrderPayment(idOrderPayments []uint32) (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment IN (?)", idOrderPayments).Find(&results).Error
	
	return
}
 
// GetFromIDOrder 通过id_order获取内容  
func (obj *_EgOrderInvoicePaymentMgr) GetFromIDOrder(idOrder uint32) (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDOrder 批量唯一主键查找 
func (obj *_EgOrderInvoicePaymentMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderInvoicePaymentMgr) FetchByPrimaryKey(idOrderInvoice uint32 ,idOrderPayment uint32 ) (result EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ? AND id_order_payment = ?", idOrderInvoice , idOrderPayment).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByOrderPayment  获取多个内容
 func (obj *_EgOrderInvoicePaymentMgr) FetchIndexByOrderPayment(idOrderPayment uint32 ) (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_payment = ?", idOrderPayment).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDOrder  获取多个内容
 func (obj *_EgOrderInvoicePaymentMgr) FetchIndexByIDOrder(idOrder uint32 ) (results []*EgOrderInvoicePayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}
 

	

