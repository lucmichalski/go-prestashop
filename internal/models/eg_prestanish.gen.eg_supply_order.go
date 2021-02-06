package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgSupplyOrderMgr struct {
	*_BaseMgr
}

// EgSupplyOrderMgr open func
func EgSupplyOrderMgr(db *gorm.DB) *_EgSupplyOrderMgr {
	if db == nil {
		panic(fmt.Errorf("EgSupplyOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSupplyOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_supply_order"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSupplyOrderMgr) GetTableName() string {
	return "eg_supply_order"
}

// Get 获取
func (obj *_EgSupplyOrderMgr) Get() (result EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSupplyOrderMgr) Gets() (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplyOrder id_supply_order获取 
func (obj *_EgSupplyOrderMgr) WithIDSupplyOrder(idSupplyOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

// WithIDSupplier id_supplier获取 
func (obj *_EgSupplyOrderMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithSupplierName supplier_name获取 
func (obj *_EgSupplyOrderMgr) WithSupplierName(supplierName string) Option {
	return optionFunc(func(o *options) { o.query["supplier_name"] = supplierName })
}

// WithIDLang id_lang获取 
func (obj *_EgSupplyOrderMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDWarehouse id_warehouse获取 
func (obj *_EgSupplyOrderMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

// WithIDSupplyOrderState id_supply_order_state获取 
func (obj *_EgSupplyOrderMgr) WithIDSupplyOrderState(idSupplyOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_state"] = idSupplyOrderState })
}

// WithIDCurrency id_currency获取 
func (obj *_EgSupplyOrderMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDRefCurrency id_ref_currency获取 
func (obj *_EgSupplyOrderMgr) WithIDRefCurrency(idRefCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_ref_currency"] = idRefCurrency })
}

// WithReference reference获取 
func (obj *_EgSupplyOrderMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

// WithDateAdd date_add获取 
func (obj *_EgSupplyOrderMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgSupplyOrderMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithDateDeliveryExpected date_delivery_expected获取 
func (obj *_EgSupplyOrderMgr) WithDateDeliveryExpected(dateDeliveryExpected time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_delivery_expected"] = dateDeliveryExpected })
}

// WithTotalTe total_te获取 
func (obj *_EgSupplyOrderMgr) WithTotalTe(totalTe float64) Option {
	return optionFunc(func(o *options) { o.query["total_te"] = totalTe })
}

// WithTotalWithDiscountTe total_with_discount_te获取 
func (obj *_EgSupplyOrderMgr) WithTotalWithDiscountTe(totalWithDiscountTe float64) Option {
	return optionFunc(func(o *options) { o.query["total_with_discount_te"] = totalWithDiscountTe })
}

// WithTotalTax total_tax获取 
func (obj *_EgSupplyOrderMgr) WithTotalTax(totalTax float64) Option {
	return optionFunc(func(o *options) { o.query["total_tax"] = totalTax })
}

// WithTotalTi total_ti获取 
func (obj *_EgSupplyOrderMgr) WithTotalTi(totalTi float64) Option {
	return optionFunc(func(o *options) { o.query["total_ti"] = totalTi })
}

// WithDiscountRate discount_rate获取 
func (obj *_EgSupplyOrderMgr) WithDiscountRate(discountRate float64) Option {
	return optionFunc(func(o *options) { o.query["discount_rate"] = discountRate })
}

// WithDiscountValueTe discount_value_te获取 
func (obj *_EgSupplyOrderMgr) WithDiscountValueTe(discountValueTe float64) Option {
	return optionFunc(func(o *options) { o.query["discount_value_te"] = discountValueTe })
}

// WithIsTemplate is_template获取 
func (obj *_EgSupplyOrderMgr) WithIsTemplate(isTemplate bool) Option {
	return optionFunc(func(o *options) { o.query["is_template"] = isTemplate })
}


// GetByOption 功能选项模式获取
func (obj *_EgSupplyOrderMgr) GetByOption(opts ...Option) (result EgSupplyOrder, err error) {
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
func (obj *_EgSupplyOrderMgr) GetByOptions(opts ...Option) (results []*EgSupplyOrder, err error) {
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


// GetFromIDSupplyOrder 通过id_supply_order获取内容  
func (obj *_EgSupplyOrderMgr)  GetFromIDSupplyOrder(idSupplyOrder uint32) (result EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&result).Error
	
	return
}

// GetBatchFromIDSupplyOrder 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error
	
	return
}
 
// GetFromIDSupplier 通过id_supplier获取内容  
func (obj *_EgSupplyOrderMgr) GetFromIDSupplier(idSupplier uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplier 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error
	
	return
}
 
// GetFromSupplierName 通过supplier_name获取内容  
func (obj *_EgSupplyOrderMgr) GetFromSupplierName(supplierName string) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_name = ?", supplierName).Find(&results).Error
	
	return
}

// GetBatchFromSupplierName 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromSupplierName(supplierNames []string) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_name IN (?)", supplierNames).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgSupplyOrderMgr) GetFromIDLang(idLang uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDWarehouse 通过id_warehouse获取内容  
func (obj *_EgSupplyOrderMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error
	
	return
}
 
// GetFromIDSupplyOrderState 通过id_supply_order_state获取内容  
func (obj *_EgSupplyOrderMgr) GetFromIDSupplyOrderState(idSupplyOrderState uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplyOrderState 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromIDSupplyOrderState(idSupplyOrderStates []uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state IN (?)", idSupplyOrderStates).Find(&results).Error
	
	return
}
 
// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgSupplyOrderMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromIDRefCurrency 通过id_ref_currency获取内容  
func (obj *_EgSupplyOrderMgr) GetFromIDRefCurrency(idRefCurrency uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ref_currency = ?", idRefCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDRefCurrency 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromIDRefCurrency(idRefCurrencys []uint32) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ref_currency IN (?)", idRefCurrencys).Find(&results).Error
	
	return
}
 
// GetFromReference 通过reference获取内容  
func (obj *_EgSupplyOrderMgr) GetFromReference(reference string) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}

// GetBatchFromReference 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromReference(references []string) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgSupplyOrderMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgSupplyOrderMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromDateDeliveryExpected 通过date_delivery_expected获取内容  
func (obj *_EgSupplyOrderMgr) GetFromDateDeliveryExpected(dateDeliveryExpected time.Time) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_delivery_expected = ?", dateDeliveryExpected).Find(&results).Error
	
	return
}

// GetBatchFromDateDeliveryExpected 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromDateDeliveryExpected(dateDeliveryExpecteds []time.Time) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_delivery_expected IN (?)", dateDeliveryExpecteds).Find(&results).Error
	
	return
}
 
// GetFromTotalTe 通过total_te获取内容  
func (obj *_EgSupplyOrderMgr) GetFromTotalTe(totalTe float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_te = ?", totalTe).Find(&results).Error
	
	return
}

// GetBatchFromTotalTe 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromTotalTe(totalTes []float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_te IN (?)", totalTes).Find(&results).Error
	
	return
}
 
// GetFromTotalWithDiscountTe 通过total_with_discount_te获取内容  
func (obj *_EgSupplyOrderMgr) GetFromTotalWithDiscountTe(totalWithDiscountTe float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_with_discount_te = ?", totalWithDiscountTe).Find(&results).Error
	
	return
}

// GetBatchFromTotalWithDiscountTe 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromTotalWithDiscountTe(totalWithDiscountTes []float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_with_discount_te IN (?)", totalWithDiscountTes).Find(&results).Error
	
	return
}
 
// GetFromTotalTax 通过total_tax获取内容  
func (obj *_EgSupplyOrderMgr) GetFromTotalTax(totalTax float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_tax = ?", totalTax).Find(&results).Error
	
	return
}

// GetBatchFromTotalTax 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromTotalTax(totalTaxs []float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_tax IN (?)", totalTaxs).Find(&results).Error
	
	return
}
 
// GetFromTotalTi 通过total_ti获取内容  
func (obj *_EgSupplyOrderMgr) GetFromTotalTi(totalTi float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_ti = ?", totalTi).Find(&results).Error
	
	return
}

// GetBatchFromTotalTi 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromTotalTi(totalTis []float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_ti IN (?)", totalTis).Find(&results).Error
	
	return
}
 
// GetFromDiscountRate 通过discount_rate获取内容  
func (obj *_EgSupplyOrderMgr) GetFromDiscountRate(discountRate float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_rate = ?", discountRate).Find(&results).Error
	
	return
}

// GetBatchFromDiscountRate 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromDiscountRate(discountRates []float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_rate IN (?)", discountRates).Find(&results).Error
	
	return
}
 
// GetFromDiscountValueTe 通过discount_value_te获取内容  
func (obj *_EgSupplyOrderMgr) GetFromDiscountValueTe(discountValueTe float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_value_te = ?", discountValueTe).Find(&results).Error
	
	return
}

// GetBatchFromDiscountValueTe 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromDiscountValueTe(discountValueTes []float64) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_value_te IN (?)", discountValueTes).Find(&results).Error
	
	return
}
 
// GetFromIsTemplate 通过is_template获取内容  
func (obj *_EgSupplyOrderMgr) GetFromIsTemplate(isTemplate bool) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_template = ?", isTemplate).Find(&results).Error
	
	return
}

// GetBatchFromIsTemplate 批量唯一主键查找 
func (obj *_EgSupplyOrderMgr) GetBatchFromIsTemplate(isTemplates []bool) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_template IN (?)", isTemplates).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSupplyOrderMgr) FetchByPrimaryKey(idSupplyOrder uint32 ) (result EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDSupplier  获取多个内容
 func (obj *_EgSupplyOrderMgr) FetchIndexByIDSupplier(idSupplier uint32 ) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDWarehouse  获取多个内容
 func (obj *_EgSupplyOrderMgr) FetchIndexByIDWarehouse(idWarehouse uint32 ) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}
 
 // FetchIndexByReference  获取多个内容
 func (obj *_EgSupplyOrderMgr) FetchIndexByReference(reference string ) (results []*EgSupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}
 

	

