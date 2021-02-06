package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOrderSlipDetailTaxMgr struct {
	*_BaseMgr
}

// EgOrderSlipDetailTaxMgr open func
func EgOrderSlipDetailTaxMgr(db *gorm.DB) *_EgOrderSlipDetailTaxMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderSlipDetailTaxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderSlipDetailTaxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_slip_detail_tax"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderSlipDetailTaxMgr) GetTableName() string {
	return "eg_order_slip_detail_tax"
}

// Get 获取
func (obj *_EgOrderSlipDetailTaxMgr) Get() (result EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderSlipDetailTaxMgr) Gets() (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderSlipDetail id_order_slip_detail获取 
func (obj *_EgOrderSlipDetailTaxMgr) WithIDOrderSlipDetail(idOrderSlipDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_slip_detail"] = idOrderSlipDetail })
}

// WithIDTax id_tax获取 
func (obj *_EgOrderSlipDetailTaxMgr) WithIDTax(idTax uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithUnitAmount unit_amount获取 
func (obj *_EgOrderSlipDetailTaxMgr) WithUnitAmount(unitAmount float64) Option {
	return optionFunc(func(o *options) { o.query["unit_amount"] = unitAmount })
}

// WithTotalAmount total_amount获取 
func (obj *_EgOrderSlipDetailTaxMgr) WithTotalAmount(totalAmount float64) Option {
	return optionFunc(func(o *options) { o.query["total_amount"] = totalAmount })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderSlipDetailTaxMgr) GetByOption(opts ...Option) (result EgOrderSlipDetailTax, err error) {
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
func (obj *_EgOrderSlipDetailTaxMgr) GetByOptions(opts ...Option) (results []*EgOrderSlipDetailTax, err error) {
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


// GetFromIDOrderSlipDetail 通过id_order_slip_detail获取内容  
func (obj *_EgOrderSlipDetailTaxMgr) GetFromIDOrderSlipDetail(idOrderSlipDetail uint32) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip_detail = ?", idOrderSlipDetail).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderSlipDetail 批量唯一主键查找 
func (obj *_EgOrderSlipDetailTaxMgr) GetBatchFromIDOrderSlipDetail(idOrderSlipDetails []uint32) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip_detail IN (?)", idOrderSlipDetails).Find(&results).Error
	
	return
}
 
// GetFromIDTax 通过id_tax获取内容  
func (obj *_EgOrderSlipDetailTaxMgr) GetFromIDTax(idTax uint32) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}

// GetBatchFromIDTax 批量唯一主键查找 
func (obj *_EgOrderSlipDetailTaxMgr) GetBatchFromIDTax(idTaxs []uint32) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error
	
	return
}
 
// GetFromUnitAmount 通过unit_amount获取内容  
func (obj *_EgOrderSlipDetailTaxMgr) GetFromUnitAmount(unitAmount float64) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount = ?", unitAmount).Find(&results).Error
	
	return
}

// GetBatchFromUnitAmount 批量唯一主键查找 
func (obj *_EgOrderSlipDetailTaxMgr) GetBatchFromUnitAmount(unitAmounts []float64) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_amount IN (?)", unitAmounts).Find(&results).Error
	
	return
}
 
// GetFromTotalAmount 通过total_amount获取内容  
func (obj *_EgOrderSlipDetailTaxMgr) GetFromTotalAmount(totalAmount float64) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount = ?", totalAmount).Find(&results).Error
	
	return
}

// GetBatchFromTotalAmount 批量唯一主键查找 
func (obj *_EgOrderSlipDetailTaxMgr) GetBatchFromTotalAmount(totalAmounts []float64) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_amount IN (?)", totalAmounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 
 // FetchIndexByIDOrderSlipDetail  获取多个内容
 func (obj *_EgOrderSlipDetailTaxMgr) FetchIndexByIDOrderSlipDetail(idOrderSlipDetail uint32 ) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip_detail = ?", idOrderSlipDetail).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDTax  获取多个内容
 func (obj *_EgOrderSlipDetailTaxMgr) FetchIndexByIDTax(idTax uint32 ) (results []*EgOrderSlipDetailTax, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}
 

	

