package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOrderDetailTaxMgr struct {
	*_BaseMgr
}

// EgOrderDetailTaxMgr open func
func EgOrderDetailTaxMgr(db *gorm.DB) *_EgOrderDetailTaxMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderDetailTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderDetailTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_detail_tax"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderDetailTaxMgr) GetTableName() string {
	return "eg_order_detail_tax"
}

// Get 获取
func (obj *_EgOrderDetailTaxMgr) Get() (result EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderDetailTaxMgr) Gets() (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderDetail id_order_detail获取 
func (obj *_EgOrderDetailTaxMgr) WithIDOrderDetail(idOrderDetail int) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

// WithIDTax id_tax获取 
func (obj *_EgOrderDetailTaxMgr) WithIDTax(idTax int) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithUnitAmount unit_amount获取 
func (obj *_EgOrderDetailTaxMgr) WithUnitAmount(unitAmount float64) Option {
	return optionFunc(func(o *options) { o.query["unit_amount"] = unitAmount })
}

// WithTotalAmount total_amount获取 
func (obj *_EgOrderDetailTaxMgr) WithTotalAmount(totalAmount float64) Option {
	return optionFunc(func(o *options) { o.query["total_amount"] = totalAmount })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderDetailTaxMgr) GetByOption(opts ...Option) (result EgOrderDetailTax, err error) {
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
func (obj *_EgOrderDetailTaxMgr) GetByOptions(opts ...Option) (results []*EgOrderDetailTax, err error) {
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


// GetFromIDOrderDetail 通过id_order_detail获取内容  
func (obj *_EgOrderDetailTaxMgr) GetFromIDOrderDetail(idOrderDetail int) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderDetail 批量唯一主键查找 
func (obj *_EgOrderDetailTaxMgr) GetBatchFromIDOrderDetail(idOrderDetails []int) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error
	
	return
}
 
// GetFromIDTax 通过id_tax获取内容  
func (obj *_EgOrderDetailTaxMgr) GetFromIDTax(idTax int) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}

// GetBatchFromIDTax 批量唯一主键查找 
func (obj *_EgOrderDetailTaxMgr) GetBatchFromIDTax(idTaxs []int) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error
	
	return
}
 
// GetFromUnitAmount 通过unit_amount获取内容  
func (obj *_EgOrderDetailTaxMgr) GetFromUnitAmount(unitAmount float64) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount = ?", unitAmount).Find(&results).Error
	
	return
}

// GetBatchFromUnitAmount 批量唯一主键查找 
func (obj *_EgOrderDetailTaxMgr) GetBatchFromUnitAmount(unitAmounts []float64) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount IN (?)", unitAmounts).Find(&results).Error
	
	return
}
 
// GetFromTotalAmount 通过total_amount获取内容  
func (obj *_EgOrderDetailTaxMgr) GetFromTotalAmount(totalAmount float64) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount = ?", totalAmount).Find(&results).Error
	
	return
}

// GetBatchFromTotalAmount 批量唯一主键查找 
func (obj *_EgOrderDetailTaxMgr) GetBatchFromTotalAmount(totalAmounts []float64) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount IN (?)", totalAmounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 
 // FetchIndexByIDOrderDetail  获取多个内容
 func (obj *_EgOrderDetailTaxMgr) FetchIndexByIDOrderDetail(idOrderDetail int ) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDTax  获取多个内容
 func (obj *_EgOrderDetailTaxMgr) FetchIndexByIDTax(idTax int ) (results []*EgOrderDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}
 

	

