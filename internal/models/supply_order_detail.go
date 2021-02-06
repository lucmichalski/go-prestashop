package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgSupplyOrderDetailMgr struct {
	*_BaseMgr
}

// EgSupplyOrderDetailMgr open func
func EgSupplyOrderDetailMgr(db *gorm.DB) *_EgSupplyOrderDetailMgr {
	if db == nil {
		panic(fmt.Errorf("EgSupplyOrderDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSupplyOrderDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_supply_order_detail"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSupplyOrderDetailMgr) GetTableName() string {
	return "eg_supply_order_detail"
}

// Get 获取
func (obj *_EgSupplyOrderDetailMgr) Get() (result EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSupplyOrderDetailMgr) Gets() (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplyOrderDetail id_supply_order_detail获取 
func (obj *_EgSupplyOrderDetailMgr) WithIDSupplyOrderDetail(idSupplyOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_detail"] = idSupplyOrderDetail })
}

// WithIDSupplyOrder id_supply_order获取 
func (obj *_EgSupplyOrderDetailMgr) WithIDSupplyOrder(idSupplyOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

// WithIDCurrency id_currency获取 
func (obj *_EgSupplyOrderDetailMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDProduct id_product获取 
func (obj *_EgSupplyOrderDetailMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDProductAttribute id_product_attribute获取 
func (obj *_EgSupplyOrderDetailMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithReference reference获取 
func (obj *_EgSupplyOrderDetailMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

// WithSupplierReference supplier_reference获取 
func (obj *_EgSupplyOrderDetailMgr) WithSupplierReference(supplierReference string) Option {
	return optionFunc(func(o *options) { o.query["supplier_reference"] = supplierReference })
}

// WithName name获取 
func (obj *_EgSupplyOrderDetailMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithEan13 ean13获取 
func (obj *_EgSupplyOrderDetailMgr) WithEan13(ean13 string) Option {
	return optionFunc(func(o *options) { o.query["ean13"] = ean13 })
}

// WithIsbn isbn获取 
func (obj *_EgSupplyOrderDetailMgr) WithIsbn(isbn string) Option {
	return optionFunc(func(o *options) { o.query["isbn"] = isbn })
}

// WithUpc upc获取 
func (obj *_EgSupplyOrderDetailMgr) WithUpc(upc string) Option {
	return optionFunc(func(o *options) { o.query["upc"] = upc })
}

// WithMpn mpn获取 
func (obj *_EgSupplyOrderDetailMgr) WithMpn(mpn string) Option {
	return optionFunc(func(o *options) { o.query["mpn"] = mpn })
}

// WithExchangeRate exchange_rate获取 
func (obj *_EgSupplyOrderDetailMgr) WithExchangeRate(exchangeRate float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_rate"] = exchangeRate })
}

// WithUnitPriceTe unit_price_te获取 
func (obj *_EgSupplyOrderDetailMgr) WithUnitPriceTe(unitPriceTe float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_te"] = unitPriceTe })
}

// WithQuantityExpected quantity_expected获取 
func (obj *_EgSupplyOrderDetailMgr) WithQuantityExpected(quantityExpected uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity_expected"] = quantityExpected })
}

// WithQuantityReceived quantity_received获取 
func (obj *_EgSupplyOrderDetailMgr) WithQuantityReceived(quantityReceived uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity_received"] = quantityReceived })
}

// WithPriceTe price_te获取 
func (obj *_EgSupplyOrderDetailMgr) WithPriceTe(priceTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_te"] = priceTe })
}

// WithDiscountRate discount_rate获取 
func (obj *_EgSupplyOrderDetailMgr) WithDiscountRate(discountRate float64) Option {
	return optionFunc(func(o *options) { o.query["discount_rate"] = discountRate })
}

// WithDiscountValueTe discount_value_te获取 
func (obj *_EgSupplyOrderDetailMgr) WithDiscountValueTe(discountValueTe float64) Option {
	return optionFunc(func(o *options) { o.query["discount_value_te"] = discountValueTe })
}

// WithPriceWithDiscountTe price_with_discount_te获取 
func (obj *_EgSupplyOrderDetailMgr) WithPriceWithDiscountTe(priceWithDiscountTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_with_discount_te"] = priceWithDiscountTe })
}

// WithTaxRate tax_rate获取 
func (obj *_EgSupplyOrderDetailMgr) WithTaxRate(taxRate float64) Option {
	return optionFunc(func(o *options) { o.query["tax_rate"] = taxRate })
}

// WithTaxValue tax_value获取 
func (obj *_EgSupplyOrderDetailMgr) WithTaxValue(taxValue float64) Option {
	return optionFunc(func(o *options) { o.query["tax_value"] = taxValue })
}

// WithPriceTi price_ti获取 
func (obj *_EgSupplyOrderDetailMgr) WithPriceTi(priceTi float64) Option {
	return optionFunc(func(o *options) { o.query["price_ti"] = priceTi })
}

// WithTaxValueWithOrderDiscount tax_value_with_order_discount获取 
func (obj *_EgSupplyOrderDetailMgr) WithTaxValueWithOrderDiscount(taxValueWithOrderDiscount float64) Option {
	return optionFunc(func(o *options) { o.query["tax_value_with_order_discount"] = taxValueWithOrderDiscount })
}

// WithPriceWithOrderDiscountTe price_with_order_discount_te获取 
func (obj *_EgSupplyOrderDetailMgr) WithPriceWithOrderDiscountTe(priceWithOrderDiscountTe float64) Option {
	return optionFunc(func(o *options) { o.query["price_with_order_discount_te"] = priceWithOrderDiscountTe })
}


// GetByOption 功能选项模式获取
func (obj *_EgSupplyOrderDetailMgr) GetByOption(opts ...Option) (result EgSupplyOrderDetail, err error) {
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
func (obj *_EgSupplyOrderDetailMgr) GetByOptions(opts ...Option) (results []*EgSupplyOrderDetail, err error) {
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


// GetFromIDSupplyOrderDetail 通过id_supply_order_detail获取内容  
func (obj *_EgSupplyOrderDetailMgr)  GetFromIDSupplyOrderDetail(idSupplyOrderDetail uint32) (result EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail = ?", idSupplyOrderDetail).Find(&result).Error
	
	return
}

// GetBatchFromIDSupplyOrderDetail 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromIDSupplyOrderDetail(idSupplyOrderDetails []uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail IN (?)", idSupplyOrderDetails).Find(&results).Error
	
	return
}
 
// GetFromIDSupplyOrder 通过id_supply_order获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromIDSupplyOrder(idSupplyOrder uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplyOrder 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error
	
	return
}
 
// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromIDProduct(idProduct uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDProductAttribute 通过id_product_attribute获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error
	
	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error
	
	return
}
 
// GetFromReference 通过reference获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromReference(reference string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}

// GetBatchFromReference 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromReference(references []string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error
	
	return
}
 
// GetFromSupplierReference 通过supplier_reference获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromSupplierReference(supplierReference string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference = ?", supplierReference).Find(&results).Error
	
	return
}

// GetBatchFromSupplierReference 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromSupplierReference(supplierReferences []string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_reference IN (?)", supplierReferences).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromName(name string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromName(names []string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromEan13 通过ean13获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromEan13(ean13 string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 = ?", ean13).Find(&results).Error
	
	return
}

// GetBatchFromEan13 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromEan13(ean13s []string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ean13 IN (?)", ean13s).Find(&results).Error
	
	return
}
 
// GetFromIsbn 通过isbn获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromIsbn(isbn string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn = ?", isbn).Find(&results).Error
	
	return
}

// GetBatchFromIsbn 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromIsbn(isbns []string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("isbn IN (?)", isbns).Find(&results).Error
	
	return
}
 
// GetFromUpc 通过upc获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromUpc(upc string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc = ?", upc).Find(&results).Error
	
	return
}

// GetBatchFromUpc 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromUpc(upcs []string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("upc IN (?)", upcs).Find(&results).Error
	
	return
}
 
// GetFromMpn 通过mpn获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromMpn(mpn string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn = ?", mpn).Find(&results).Error
	
	return
}

// GetBatchFromMpn 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromMpn(mpns []string) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mpn IN (?)", mpns).Find(&results).Error
	
	return
}
 
// GetFromExchangeRate 通过exchange_rate获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromExchangeRate(exchangeRate float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("exchange_rate = ?", exchangeRate).Find(&results).Error
	
	return
}

// GetBatchFromExchangeRate 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromExchangeRate(exchangeRates []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("exchange_rate IN (?)", exchangeRates).Find(&results).Error
	
	return
}
 
// GetFromUnitPriceTe 通过unit_price_te获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromUnitPriceTe(unitPriceTe float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_te = ?", unitPriceTe).Find(&results).Error
	
	return
}

// GetBatchFromUnitPriceTe 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromUnitPriceTe(unitPriceTes []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_te IN (?)", unitPriceTes).Find(&results).Error
	
	return
}
 
// GetFromQuantityExpected 通过quantity_expected获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromQuantityExpected(quantityExpected uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_expected = ?", quantityExpected).Find(&results).Error
	
	return
}

// GetBatchFromQuantityExpected 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromQuantityExpected(quantityExpecteds []uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_expected IN (?)", quantityExpecteds).Find(&results).Error
	
	return
}
 
// GetFromQuantityReceived 通过quantity_received获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromQuantityReceived(quantityReceived uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_received = ?", quantityReceived).Find(&results).Error
	
	return
}

// GetBatchFromQuantityReceived 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromQuantityReceived(quantityReceiveds []uint32) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_received IN (?)", quantityReceiveds).Find(&results).Error
	
	return
}
 
// GetFromPriceTe 通过price_te获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromPriceTe(priceTe float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te = ?", priceTe).Find(&results).Error
	
	return
}

// GetBatchFromPriceTe 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromPriceTe(priceTes []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_te IN (?)", priceTes).Find(&results).Error
	
	return
}
 
// GetFromDiscountRate 通过discount_rate获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromDiscountRate(discountRate float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_rate = ?", discountRate).Find(&results).Error
	
	return
}

// GetBatchFromDiscountRate 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromDiscountRate(discountRates []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_rate IN (?)", discountRates).Find(&results).Error
	
	return
}
 
// GetFromDiscountValueTe 通过discount_value_te获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromDiscountValueTe(discountValueTe float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_value_te = ?", discountValueTe).Find(&results).Error
	
	return
}

// GetBatchFromDiscountValueTe 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromDiscountValueTe(discountValueTes []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_value_te IN (?)", discountValueTes).Find(&results).Error
	
	return
}
 
// GetFromPriceWithDiscountTe 通过price_with_discount_te获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromPriceWithDiscountTe(priceWithDiscountTe float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_with_discount_te = ?", priceWithDiscountTe).Find(&results).Error
	
	return
}

// GetBatchFromPriceWithDiscountTe 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromPriceWithDiscountTe(priceWithDiscountTes []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_with_discount_te IN (?)", priceWithDiscountTes).Find(&results).Error
	
	return
}
 
// GetFromTaxRate 通过tax_rate获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromTaxRate(taxRate float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_rate = ?", taxRate).Find(&results).Error
	
	return
}

// GetBatchFromTaxRate 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromTaxRate(taxRates []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_rate IN (?)", taxRates).Find(&results).Error
	
	return
}
 
// GetFromTaxValue 通过tax_value获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromTaxValue(taxValue float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_value = ?", taxValue).Find(&results).Error
	
	return
}

// GetBatchFromTaxValue 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromTaxValue(taxValues []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_value IN (?)", taxValues).Find(&results).Error
	
	return
}
 
// GetFromPriceTi 通过price_ti获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromPriceTi(priceTi float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_ti = ?", priceTi).Find(&results).Error
	
	return
}

// GetBatchFromPriceTi 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromPriceTi(priceTis []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_ti IN (?)", priceTis).Find(&results).Error
	
	return
}
 
// GetFromTaxValueWithOrderDiscount 通过tax_value_with_order_discount获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromTaxValueWithOrderDiscount(taxValueWithOrderDiscount float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_value_with_order_discount = ?", taxValueWithOrderDiscount).Find(&results).Error
	
	return
}

// GetBatchFromTaxValueWithOrderDiscount 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromTaxValueWithOrderDiscount(taxValueWithOrderDiscounts []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_value_with_order_discount IN (?)", taxValueWithOrderDiscounts).Find(&results).Error
	
	return
}
 
// GetFromPriceWithOrderDiscountTe 通过price_with_order_discount_te获取内容  
func (obj *_EgSupplyOrderDetailMgr) GetFromPriceWithOrderDiscountTe(priceWithOrderDiscountTe float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_with_order_discount_te = ?", priceWithOrderDiscountTe).Find(&results).Error
	
	return
}

// GetBatchFromPriceWithOrderDiscountTe 批量唯一主键查找 
func (obj *_EgSupplyOrderDetailMgr) GetBatchFromPriceWithOrderDiscountTe(priceWithOrderDiscountTes []float64) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price_with_order_discount_te IN (?)", priceWithOrderDiscountTes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSupplyOrderDetailMgr) FetchByPrimaryKey(idSupplyOrderDetail uint32 ) (result EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_detail = ?", idSupplyOrderDetail).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDSupplyOrder  获取多个内容
 func (obj *_EgSupplyOrderDetailMgr) FetchIndexByIDSupplyOrder(idSupplyOrder uint32 ,idProduct uint32 ) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ? AND id_product = ?", idSupplyOrder , idProduct).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDProductProductAttribute  获取多个内容
 func (obj *_EgSupplyOrderDetailMgr) FetchIndexByIDProductProductAttribute(idProduct uint32 ,idProductAttribute uint32 ) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_product_attribute = ?", idProduct , idProductAttribute).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDProductAttribute  获取多个内容
 func (obj *_EgSupplyOrderDetailMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32 ) (results []*EgSupplyOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error
	
	return
}
 

	
