package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOrderInvoiceTaxMgr struct {
	*_BaseMgr
}

// EgOrderInvoiceTaxMgr open func
func EgOrderInvoiceTaxMgr(db *gorm.DB) *_EgOrderInvoiceTaxMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderInvoiceTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderInvoiceTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_invoice_tax"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderInvoiceTaxMgr) GetTableName() string {
	return "eg_order_invoice_tax"
}

// Get 获取
func (obj *_EgOrderInvoiceTaxMgr) Get() (result EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderInvoiceTaxMgr) Gets() (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderInvoice id_order_invoice获取 
func (obj *_EgOrderInvoiceTaxMgr) WithIDOrderInvoice(idOrderInvoice int) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

// WithType type获取 
func (obj *_EgOrderInvoiceTaxMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithIDTax id_tax获取 
func (obj *_EgOrderInvoiceTaxMgr) WithIDTax(idTax int) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithAmount amount获取 
func (obj *_EgOrderInvoiceTaxMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderInvoiceTaxMgr) GetByOption(opts ...Option) (result EgOrderInvoiceTax, err error) {
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
func (obj *_EgOrderInvoiceTaxMgr) GetByOptions(opts ...Option) (results []*EgOrderInvoiceTax, err error) {
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
func (obj *_EgOrderInvoiceTaxMgr) GetFromIDOrderInvoice(idOrderInvoice int) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderInvoice 批量唯一主键查找 
func (obj *_EgOrderInvoiceTaxMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []int) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error
	
	return
}
 
// GetFromType 通过type获取内容  
func (obj *_EgOrderInvoiceTaxMgr) GetFromType(_type string) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量唯一主键查找 
func (obj *_EgOrderInvoiceTaxMgr) GetBatchFromType(_types []string) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromIDTax 通过id_tax获取内容  
func (obj *_EgOrderInvoiceTaxMgr) GetFromIDTax(idTax int) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}

// GetBatchFromIDTax 批量唯一主键查找 
func (obj *_EgOrderInvoiceTaxMgr) GetBatchFromIDTax(idTaxs []int) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error
	
	return
}
 
// GetFromAmount 通过amount获取内容  
func (obj *_EgOrderInvoiceTaxMgr) GetFromAmount(amount float64) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error
	
	return
}

// GetBatchFromAmount 批量唯一主键查找 
func (obj *_EgOrderInvoiceTaxMgr) GetBatchFromAmount(amounts []float64) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 
 // FetchIndexByIDTax  获取多个内容
 func (obj *_EgOrderInvoiceTaxMgr) FetchIndexByIDTax(idTax int ) (results []*EgOrderInvoiceTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}
 

	

