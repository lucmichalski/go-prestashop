package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgOrderDetailMgr struct {
	*_BaseMgr
}

// EgOrderDetailMgr open func
func EgOrderDetailMgr(db *gorm.DB) *_EgOrderDetailMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_detail"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderDetailMgr) GetTableName() string {
	return "eg_order_detail"
}

// Get 获取
func (obj *_EgOrderDetailMgr) Get() (result EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderDetailMgr) Gets() (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderDetail id_order_detail获取 
func (obj *_EgOrderDetailMgr) WithIDOrderDetail(idOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

// WithIDOrder id_order获取 
func (obj *_EgOrderDetailMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithIDOrderInvoice id_order_invoice获取 
func (obj *_EgOrderDetailMgr) WithIDOrderInvoice(idOrderInvoice int) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

// WithIDWarehouse id_warehouse获取 
func (obj *_EgOrderDetailMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

// WithIDShop id_shop获取 
func (obj *_EgOrderDetailMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithProductID product_id获取 
func (obj *_EgOrderDetailMgr) WithProductID(productID uint32) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

// WithProductAttributeID product_attribute_id获取 
func (obj *_EgOrderDetailMgr) WithProductAttributeID(productAttributeID uint32) Option {
	return optionFunc(func(o *options) { o.query["product_attribute_id"] = productAttributeID })
}

// WithIDCustomization id_customization获取 
func (obj *_EgOrderDetailMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

// WithProductName product_name获取 
func (obj *_EgOrderDetailMgr) WithProductName(productName string) Option {
	return optionFunc(func(o *options) { o.query["product_name"] = productName })
}

// WithProductQuantity product_quantity获取 
func (obj *_EgOrderDetailMgr) WithProductQuantity(productQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity"] = productQuantity })
}

// WithProductQuantityInStock product_quantity_in_stock获取 
func (obj *_EgOrderDetailMgr) WithProductQuantityInStock(productQuantityInStock int) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_in_stock"] = productQuantityInStock })
}

// WithProductQuantityRefunded product_quantity_refunded获取 
func (obj *_EgOrderDetailMgr) WithProductQuantityRefunded(productQuantityRefunded uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_refunded"] = productQuantityRefunded })
}

// WithProductQuantityReturn product_quantity_return获取 
func (obj *_EgOrderDetailMgr) WithProductQuantityReturn(productQuantityReturn uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_return"] = productQuantityReturn })
}

// WithProductQuantityReinjected product_quantity_reinjected获取 
func (obj *_EgOrderDetailMgr) WithProductQuantityReinjected(productQuantityReinjected uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_reinjected"] = productQuantityReinjected })
}

// WithProductPrice product_price获取 
func (obj *_EgOrderDetailMgr) WithProductPrice(productPrice float64) Option {
	return optionFunc(func(o *options) { o.query["product_price"] = productPrice })
}

// WithReductionPercent reduction_percent获取 
func (obj *_EgOrderDetailMgr) WithReductionPercent(reductionPercent float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_percent"] = reductionPercent })
}

// WithReductionAmount reduction_amount获取 
func (obj *_EgOrderDetailMgr) WithReductionAmount(reductionAmount float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_amount"] = reductionAmount })
}

// WithReductionAmountTaxIncl reduction_amount_tax_incl获取 
func (obj *_EgOrderDetailMgr) WithReductionAmountTaxIncl(reductionAmountTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_amount_tax_incl"] = reductionAmountTaxIncl })
}

// WithReductionAmountTaxExcl reduction_amount_tax_excl获取 
func (obj *_EgOrderDetailMgr) WithReductionAmountTaxExcl(reductionAmountTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_amount_tax_excl"] = reductionAmountTaxExcl })
}

// WithGroupReduction group_reduction获取 
func (obj *_EgOrderDetailMgr) WithGroupReduction(groupReduction float64) Option {
	return optionFunc(func(o *options) { o.query["group_reduction"] = groupReduction })
}

// WithProductQuantityDiscount product_quantity_discount获取 
func (obj *_EgOrderDetailMgr) WithProductQuantityDiscount(productQuantityDiscount float64) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_discount"] = productQuantityDiscount })
}

// WithProductEan13 product_ean13获取 
func (obj *_EgOrderDetailMgr) WithProductEan13(productEan13 string) Option {
	return optionFunc(func(o *options) { o.query["product_ean13"] = productEan13 })
}

// WithProductIsbn product_isbn获取 
func (obj *_EgOrderDetailMgr) WithProductIsbn(productIsbn string) Option {
	return optionFunc(func(o *options) { o.query["product_isbn"] = productIsbn })
}

// WithProductUpc product_upc获取 
func (obj *_EgOrderDetailMgr) WithProductUpc(productUpc string) Option {
	return optionFunc(func(o *options) { o.query["product_upc"] = productUpc })
}

// WithProductMpn product_mpn获取 
func (obj *_EgOrderDetailMgr) WithProductMpn(productMpn string) Option {
	return optionFunc(func(o *options) { o.query["product_mpn"] = productMpn })
}

// WithProductReference product_reference获取 
func (obj *_EgOrderDetailMgr) WithProductReference(productReference string) Option {
	return optionFunc(func(o *options) { o.query["product_reference"] = productReference })
}

// WithProductSupplierReference product_supplier_reference获取 
func (obj *_EgOrderDetailMgr) WithProductSupplierReference(productSupplierReference string) Option {
	return optionFunc(func(o *options) { o.query["product_supplier_reference"] = productSupplierReference })
}

// WithProductWeight product_weight获取 
func (obj *_EgOrderDetailMgr) WithProductWeight(productWeight float64) Option {
	return optionFunc(func(o *options) { o.query["product_weight"] = productWeight })
}

// WithIDTaxRulesGroup id_tax_rules_group获取 
func (obj *_EgOrderDetailMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithTaxComputationMethod tax_computation_method获取 
func (obj *_EgOrderDetailMgr) WithTaxComputationMethod(taxComputationMethod bool) Option {
	return optionFunc(func(o *options) { o.query["tax_computation_method"] = taxComputationMethod })
}

// WithTaxName tax_name获取 
func (obj *_EgOrderDetailMgr) WithTaxName(taxName string) Option {
	return optionFunc(func(o *options) { o.query["tax_name"] = taxName })
}

// WithTaxRate tax_rate获取 
func (obj *_EgOrderDetailMgr) WithTaxRate(taxRate float64) Option {
	return optionFunc(func(o *options) { o.query["tax_rate"] = taxRate })
}

// WithEcotax ecotax获取 
func (obj *_EgOrderDetailMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

// WithEcotaxTaxRate ecotax_tax_rate获取 
func (obj *_EgOrderDetailMgr) WithEcotaxTaxRate(ecotaxTaxRate float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax_tax_rate"] = ecotaxTaxRate })
}

// WithDiscountQuantityApplied discount_quantity_applied获取 
func (obj *_EgOrderDetailMgr) WithDiscountQuantityApplied(discountQuantityApplied bool) Option {
	return optionFunc(func(o *options) { o.query["discount_quantity_applied"] = discountQuantityApplied })
}

// WithDownloadHash download_hash获取 
func (obj *_EgOrderDetailMgr) WithDownloadHash(downloadHash string) Option {
	return optionFunc(func(o *options) { o.query["download_hash"] = downloadHash })
}

// WithDownloadNb download_nb获取 
func (obj *_EgOrderDetailMgr) WithDownloadNb(downloadNb uint32) Option {
	return optionFunc(func(o *options) { o.query["download_nb"] = downloadNb })
}

// WithDownloadDeadline download_deadline获取 
func (obj *_EgOrderDetailMgr) WithDownloadDeadline(downloadDeadline time.Time) Option {
	return optionFunc(func(o *options) { o.query["download_deadline"] = downloadDeadline })
}

// WithTotalPriceTaxIncl total_price_tax_incl获取 
func (obj *_EgOrderDetailMgr) WithTotalPriceTaxIncl(totalPriceTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_price_tax_incl"] = totalPriceTaxIncl })
}

// WithTotalPriceTaxExcl total_price_tax_excl获取 
func (obj *_EgOrderDetailMgr) WithTotalPriceTaxExcl(totalPriceTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_price_tax_excl"] = totalPriceTaxExcl })
}

// WithUnitPriceTaxIncl unit_price_tax_incl获取 
func (obj *_EgOrderDetailMgr) WithUnitPriceTaxIncl(unitPriceTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_tax_incl"] = unitPriceTaxIncl })
}

// WithUnitPriceTaxExcl unit_price_tax_excl获取 
func (obj *_EgOrderDetailMgr) WithUnitPriceTaxExcl(unitPriceTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_tax_excl"] = unitPriceTaxExcl })
}

// WithTotalShippingPriceTaxIncl total_shipping_price_tax_incl获取 
func (obj *_EgOrderDetailMgr) WithTotalShippingPriceTaxIncl(totalShippingPriceTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_price_tax_incl"] = totalShippingPriceTaxIncl })
}

// WithTotalShippingPriceTaxExcl total_shipping_price_tax_excl获取 
func (obj *_EgOrderDetailMgr) WithTotalShippingPriceTaxExcl(totalShippingPriceTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_price_tax_excl"] = totalShippingPriceTaxExcl })
}

// WithPurchaseSupplierPrice purchase_supplier_price获取 
func (obj *_EgOrderDetailMgr) WithPurchaseSupplierPrice(purchaseSupplierPrice float64) Option {
	return optionFunc(func(o *options) { o.query["purchase_supplier_price"] = purchaseSupplierPrice })
}

// WithOriginalProductPrice original_product_price获取 
func (obj *_EgOrderDetailMgr) WithOriginalProductPrice(originalProductPrice float64) Option {
	return optionFunc(func(o *options) { o.query["original_product_price"] = originalProductPrice })
}

// WithOriginalWholesalePrice original_wholesale_price获取 
func (obj *_EgOrderDetailMgr) WithOriginalWholesalePrice(originalWholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["original_wholesale_price"] = originalWholesalePrice })
}

// WithTotalRefundedTaxExcl total_refunded_tax_excl获取 
func (obj *_EgOrderDetailMgr) WithTotalRefundedTaxExcl(totalRefundedTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_refunded_tax_excl"] = totalRefundedTaxExcl })
}

// WithTotalRefundedTaxIncl total_refunded_tax_incl获取 
func (obj *_EgOrderDetailMgr) WithTotalRefundedTaxIncl(totalRefundedTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_refunded_tax_incl"] = totalRefundedTaxIncl })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderDetailMgr) GetByOption(opts ...Option) (result EgOrderDetail, err error) {
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
func (obj *_EgOrderDetailMgr) GetByOptions(opts ...Option) (results []*EgOrderDetail, err error) {
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
func (obj *_EgOrderDetailMgr)  GetFromIDOrderDetail(idOrderDetail uint32) (result EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&result).Error
	
	return
}

// GetBatchFromIDOrderDetail 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromIDOrderDetail(idOrderDetails []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error
	
	return
}
 
// GetFromIDOrder 通过id_order获取内容  
func (obj *_EgOrderDetailMgr) GetFromIDOrder(idOrder uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDOrder 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error
	
	return
}
 
// GetFromIDOrderInvoice 通过id_order_invoice获取内容  
func (obj *_EgOrderDetailMgr) GetFromIDOrderInvoice(idOrderInvoice int) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderInvoice 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []int) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error
	
	return
}
 
// GetFromIDWarehouse 通过id_warehouse获取内容  
func (obj *_EgOrderDetailMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgOrderDetailMgr) GetFromIDShop(idShop uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromProductID 通过product_id获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductID(productID uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id = ?", productID).Find(&results).Error
	
	return
}

// GetBatchFromProductID 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductID(productIDs []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id IN (?)", productIDs).Find(&results).Error
	
	return
}
 
// GetFromProductAttributeID 通过product_attribute_id获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductAttributeID(productAttributeID uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_attribute_id = ?", productAttributeID).Find(&results).Error
	
	return
}

// GetBatchFromProductAttributeID 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductAttributeID(productAttributeIDs []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_attribute_id IN (?)", productAttributeIDs).Find(&results).Error
	
	return
}
 
// GetFromIDCustomization 通过id_customization获取内容  
func (obj *_EgOrderDetailMgr) GetFromIDCustomization(idCustomization uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomization 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error
	
	return
}
 
// GetFromProductName 通过product_name获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductName(productName string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_name = ?", productName).Find(&results).Error
	
	return
}

// GetBatchFromProductName 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductName(productNames []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_name IN (?)", productNames).Find(&results).Error
	
	return
}
 
// GetFromProductQuantity 通过product_quantity获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductQuantity(productQuantity uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity = ?", productQuantity).Find(&results).Error
	
	return
}

// GetBatchFromProductQuantity 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductQuantity(productQuantitys []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity IN (?)", productQuantitys).Find(&results).Error
	
	return
}
 
// GetFromProductQuantityInStock 通过product_quantity_in_stock获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductQuantityInStock(productQuantityInStock int) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_in_stock = ?", productQuantityInStock).Find(&results).Error
	
	return
}

// GetBatchFromProductQuantityInStock 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductQuantityInStock(productQuantityInStocks []int) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_in_stock IN (?)", productQuantityInStocks).Find(&results).Error
	
	return
}
 
// GetFromProductQuantityRefunded 通过product_quantity_refunded获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductQuantityRefunded(productQuantityRefunded uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_refunded = ?", productQuantityRefunded).Find(&results).Error
	
	return
}

// GetBatchFromProductQuantityRefunded 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductQuantityRefunded(productQuantityRefundeds []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_refunded IN (?)", productQuantityRefundeds).Find(&results).Error
	
	return
}
 
// GetFromProductQuantityReturn 通过product_quantity_return获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductQuantityReturn(productQuantityReturn uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_return = ?", productQuantityReturn).Find(&results).Error
	
	return
}

// GetBatchFromProductQuantityReturn 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductQuantityReturn(productQuantityReturns []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_return IN (?)", productQuantityReturns).Find(&results).Error
	
	return
}
 
// GetFromProductQuantityReinjected 通过product_quantity_reinjected获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductQuantityReinjected(productQuantityReinjected uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_reinjected = ?", productQuantityReinjected).Find(&results).Error
	
	return
}

// GetBatchFromProductQuantityReinjected 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductQuantityReinjected(productQuantityReinjecteds []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_reinjected IN (?)", productQuantityReinjecteds).Find(&results).Error
	
	return
}
 
// GetFromProductPrice 通过product_price获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductPrice(productPrice float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_price = ?", productPrice).Find(&results).Error
	
	return
}

// GetBatchFromProductPrice 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductPrice(productPrices []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_price IN (?)", productPrices).Find(&results).Error
	
	return
}
 
// GetFromReductionPercent 通过reduction_percent获取内容  
func (obj *_EgOrderDetailMgr) GetFromReductionPercent(reductionPercent float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_percent = ?", reductionPercent).Find(&results).Error
	
	return
}

// GetBatchFromReductionPercent 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromReductionPercent(reductionPercents []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_percent IN (?)", reductionPercents).Find(&results).Error
	
	return
}
 
// GetFromReductionAmount 通过reduction_amount获取内容  
func (obj *_EgOrderDetailMgr) GetFromReductionAmount(reductionAmount float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount = ?", reductionAmount).Find(&results).Error
	
	return
}

// GetBatchFromReductionAmount 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromReductionAmount(reductionAmounts []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount IN (?)", reductionAmounts).Find(&results).Error
	
	return
}
 
// GetFromReductionAmountTaxIncl 通过reduction_amount_tax_incl获取内容  
func (obj *_EgOrderDetailMgr) GetFromReductionAmountTaxIncl(reductionAmountTaxIncl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount_tax_incl = ?", reductionAmountTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromReductionAmountTaxIncl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromReductionAmountTaxIncl(reductionAmountTaxIncls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount_tax_incl IN (?)", reductionAmountTaxIncls).Find(&results).Error
	
	return
}
 
// GetFromReductionAmountTaxExcl 通过reduction_amount_tax_excl获取内容  
func (obj *_EgOrderDetailMgr) GetFromReductionAmountTaxExcl(reductionAmountTaxExcl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount_tax_excl = ?", reductionAmountTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromReductionAmountTaxExcl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromReductionAmountTaxExcl(reductionAmountTaxExcls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount_tax_excl IN (?)", reductionAmountTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromGroupReduction 通过group_reduction获取内容  
func (obj *_EgOrderDetailMgr) GetFromGroupReduction(groupReduction float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_reduction = ?", groupReduction).Find(&results).Error
	
	return
}

// GetBatchFromGroupReduction 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromGroupReduction(groupReductions []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_reduction IN (?)", groupReductions).Find(&results).Error
	
	return
}
 
// GetFromProductQuantityDiscount 通过product_quantity_discount获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductQuantityDiscount(productQuantityDiscount float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_discount = ?", productQuantityDiscount).Find(&results).Error
	
	return
}

// GetBatchFromProductQuantityDiscount 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductQuantityDiscount(productQuantityDiscounts []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_discount IN (?)", productQuantityDiscounts).Find(&results).Error
	
	return
}
 
// GetFromProductEan13 通过product_ean13获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductEan13(productEan13 string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_ean13 = ?", productEan13).Find(&results).Error
	
	return
}

// GetBatchFromProductEan13 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductEan13(productEan13s []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_ean13 IN (?)", productEan13s).Find(&results).Error
	
	return
}
 
// GetFromProductIsbn 通过product_isbn获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductIsbn(productIsbn string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_isbn = ?", productIsbn).Find(&results).Error
	
	return
}

// GetBatchFromProductIsbn 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductIsbn(productIsbns []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_isbn IN (?)", productIsbns).Find(&results).Error
	
	return
}
 
// GetFromProductUpc 通过product_upc获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductUpc(productUpc string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_upc = ?", productUpc).Find(&results).Error
	
	return
}

// GetBatchFromProductUpc 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductUpc(productUpcs []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_upc IN (?)", productUpcs).Find(&results).Error
	
	return
}
 
// GetFromProductMpn 通过product_mpn获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductMpn(productMpn string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_mpn = ?", productMpn).Find(&results).Error
	
	return
}

// GetBatchFromProductMpn 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductMpn(productMpns []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_mpn IN (?)", productMpns).Find(&results).Error
	
	return
}
 
// GetFromProductReference 通过product_reference获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductReference(productReference string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_reference = ?", productReference).Find(&results).Error
	
	return
}

// GetBatchFromProductReference 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductReference(productReferences []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_reference IN (?)", productReferences).Find(&results).Error
	
	return
}
 
// GetFromProductSupplierReference 通过product_supplier_reference获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductSupplierReference(productSupplierReference string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_reference = ?", productSupplierReference).Find(&results).Error
	
	return
}

// GetBatchFromProductSupplierReference 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductSupplierReference(productSupplierReferences []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_reference IN (?)", productSupplierReferences).Find(&results).Error
	
	return
}
 
// GetFromProductWeight 通过product_weight获取内容  
func (obj *_EgOrderDetailMgr) GetFromProductWeight(productWeight float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_weight = ?", productWeight).Find(&results).Error
	
	return
}

// GetBatchFromProductWeight 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromProductWeight(productWeights []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_weight IN (?)", productWeights).Find(&results).Error
	
	return
}
 
// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容  
func (obj *_EgOrderDetailMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error
	
	return
}
 
// GetFromTaxComputationMethod 通过tax_computation_method获取内容  
func (obj *_EgOrderDetailMgr) GetFromTaxComputationMethod(taxComputationMethod bool) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_computation_method = ?", taxComputationMethod).Find(&results).Error
	
	return
}

// GetBatchFromTaxComputationMethod 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTaxComputationMethod(taxComputationMethods []bool) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_computation_method IN (?)", taxComputationMethods).Find(&results).Error
	
	return
}
 
// GetFromTaxName 通过tax_name获取内容  
func (obj *_EgOrderDetailMgr) GetFromTaxName(taxName string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_name = ?", taxName).Find(&results).Error
	
	return
}

// GetBatchFromTaxName 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTaxName(taxNames []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_name IN (?)", taxNames).Find(&results).Error
	
	return
}
 
// GetFromTaxRate 通过tax_rate获取内容  
func (obj *_EgOrderDetailMgr) GetFromTaxRate(taxRate float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_rate = ?", taxRate).Find(&results).Error
	
	return
}

// GetBatchFromTaxRate 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTaxRate(taxRates []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_rate IN (?)", taxRates).Find(&results).Error
	
	return
}
 
// GetFromEcotax 通过ecotax获取内容  
func (obj *_EgOrderDetailMgr) GetFromEcotax(ecotax float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error
	
	return
}

// GetBatchFromEcotax 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error
	
	return
}
 
// GetFromEcotaxTaxRate 通过ecotax_tax_rate获取内容  
func (obj *_EgOrderDetailMgr) GetFromEcotaxTaxRate(ecotaxTaxRate float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax_tax_rate = ?", ecotaxTaxRate).Find(&results).Error
	
	return
}

// GetBatchFromEcotaxTaxRate 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromEcotaxTaxRate(ecotaxTaxRates []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax_tax_rate IN (?)", ecotaxTaxRates).Find(&results).Error
	
	return
}
 
// GetFromDiscountQuantityApplied 通过discount_quantity_applied获取内容  
func (obj *_EgOrderDetailMgr) GetFromDiscountQuantityApplied(discountQuantityApplied bool) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_quantity_applied = ?", discountQuantityApplied).Find(&results).Error
	
	return
}

// GetBatchFromDiscountQuantityApplied 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromDiscountQuantityApplied(discountQuantityApplieds []bool) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_quantity_applied IN (?)", discountQuantityApplieds).Find(&results).Error
	
	return
}
 
// GetFromDownloadHash 通过download_hash获取内容  
func (obj *_EgOrderDetailMgr) GetFromDownloadHash(downloadHash string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_hash = ?", downloadHash).Find(&results).Error
	
	return
}

// GetBatchFromDownloadHash 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromDownloadHash(downloadHashs []string) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_hash IN (?)", downloadHashs).Find(&results).Error
	
	return
}
 
// GetFromDownloadNb 通过download_nb获取内容  
func (obj *_EgOrderDetailMgr) GetFromDownloadNb(downloadNb uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_nb = ?", downloadNb).Find(&results).Error
	
	return
}

// GetBatchFromDownloadNb 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromDownloadNb(downloadNbs []uint32) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_nb IN (?)", downloadNbs).Find(&results).Error
	
	return
}
 
// GetFromDownloadDeadline 通过download_deadline获取内容  
func (obj *_EgOrderDetailMgr) GetFromDownloadDeadline(downloadDeadline time.Time) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_deadline = ?", downloadDeadline).Find(&results).Error
	
	return
}

// GetBatchFromDownloadDeadline 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromDownloadDeadline(downloadDeadlines []time.Time) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_deadline IN (?)", downloadDeadlines).Find(&results).Error
	
	return
}
 
// GetFromTotalPriceTaxIncl 通过total_price_tax_incl获取内容  
func (obj *_EgOrderDetailMgr) GetFromTotalPriceTaxIncl(totalPriceTaxIncl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_incl = ?", totalPriceTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromTotalPriceTaxIncl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTotalPriceTaxIncl(totalPriceTaxIncls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_incl IN (?)", totalPriceTaxIncls).Find(&results).Error
	
	return
}
 
// GetFromTotalPriceTaxExcl 通过total_price_tax_excl获取内容  
func (obj *_EgOrderDetailMgr) GetFromTotalPriceTaxExcl(totalPriceTaxExcl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_excl = ?", totalPriceTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromTotalPriceTaxExcl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTotalPriceTaxExcl(totalPriceTaxExcls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_excl IN (?)", totalPriceTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromUnitPriceTaxIncl 通过unit_price_tax_incl获取内容  
func (obj *_EgOrderDetailMgr) GetFromUnitPriceTaxIncl(unitPriceTaxIncl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_incl = ?", unitPriceTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromUnitPriceTaxIncl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromUnitPriceTaxIncl(unitPriceTaxIncls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_incl IN (?)", unitPriceTaxIncls).Find(&results).Error
	
	return
}
 
// GetFromUnitPriceTaxExcl 通过unit_price_tax_excl获取内容  
func (obj *_EgOrderDetailMgr) GetFromUnitPriceTaxExcl(unitPriceTaxExcl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_excl = ?", unitPriceTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromUnitPriceTaxExcl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromUnitPriceTaxExcl(unitPriceTaxExcls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_excl IN (?)", unitPriceTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromTotalShippingPriceTaxIncl 通过total_shipping_price_tax_incl获取内容  
func (obj *_EgOrderDetailMgr) GetFromTotalShippingPriceTaxIncl(totalShippingPriceTaxIncl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_price_tax_incl = ?", totalShippingPriceTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromTotalShippingPriceTaxIncl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTotalShippingPriceTaxIncl(totalShippingPriceTaxIncls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_price_tax_incl IN (?)", totalShippingPriceTaxIncls).Find(&results).Error
	
	return
}
 
// GetFromTotalShippingPriceTaxExcl 通过total_shipping_price_tax_excl获取内容  
func (obj *_EgOrderDetailMgr) GetFromTotalShippingPriceTaxExcl(totalShippingPriceTaxExcl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_price_tax_excl = ?", totalShippingPriceTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromTotalShippingPriceTaxExcl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTotalShippingPriceTaxExcl(totalShippingPriceTaxExcls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_price_tax_excl IN (?)", totalShippingPriceTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromPurchaseSupplierPrice 通过purchase_supplier_price获取内容  
func (obj *_EgOrderDetailMgr) GetFromPurchaseSupplierPrice(purchaseSupplierPrice float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("purchase_supplier_price = ?", purchaseSupplierPrice).Find(&results).Error
	
	return
}

// GetBatchFromPurchaseSupplierPrice 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromPurchaseSupplierPrice(purchaseSupplierPrices []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("purchase_supplier_price IN (?)", purchaseSupplierPrices).Find(&results).Error
	
	return
}
 
// GetFromOriginalProductPrice 通过original_product_price获取内容  
func (obj *_EgOrderDetailMgr) GetFromOriginalProductPrice(originalProductPrice float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_product_price = ?", originalProductPrice).Find(&results).Error
	
	return
}

// GetBatchFromOriginalProductPrice 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromOriginalProductPrice(originalProductPrices []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_product_price IN (?)", originalProductPrices).Find(&results).Error
	
	return
}
 
// GetFromOriginalWholesalePrice 通过original_wholesale_price获取内容  
func (obj *_EgOrderDetailMgr) GetFromOriginalWholesalePrice(originalWholesalePrice float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_wholesale_price = ?", originalWholesalePrice).Find(&results).Error
	
	return
}

// GetBatchFromOriginalWholesalePrice 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromOriginalWholesalePrice(originalWholesalePrices []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_wholesale_price IN (?)", originalWholesalePrices).Find(&results).Error
	
	return
}
 
// GetFromTotalRefundedTaxExcl 通过total_refunded_tax_excl获取内容  
func (obj *_EgOrderDetailMgr) GetFromTotalRefundedTaxExcl(totalRefundedTaxExcl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_refunded_tax_excl = ?", totalRefundedTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromTotalRefundedTaxExcl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTotalRefundedTaxExcl(totalRefundedTaxExcls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_refunded_tax_excl IN (?)", totalRefundedTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromTotalRefundedTaxIncl 通过total_refunded_tax_incl获取内容  
func (obj *_EgOrderDetailMgr) GetFromTotalRefundedTaxIncl(totalRefundedTaxIncl float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_refunded_tax_incl = ?", totalRefundedTaxIncl).Find(&results).Error
	
	return
}

// GetBatchFromTotalRefundedTaxIncl 批量唯一主键查找 
func (obj *_EgOrderDetailMgr) GetBatchFromTotalRefundedTaxIncl(totalRefundedTaxIncls []float64) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_refunded_tax_incl IN (?)", totalRefundedTaxIncls).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderDetailMgr) FetchByPrimaryKey(idOrderDetail uint32 ) (result EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDOrderIDOrderDetail  获取多个内容
 func (obj *_EgOrderDetailMgr) FetchIndexByIDOrderIDOrderDetail(idOrderDetail uint32 ,idOrder uint32 ) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ? AND id_order = ?", idOrderDetail , idOrder).Find(&results).Error
	
	return
}
 
 // FetchIndexByOrderDetailOrder  获取多个内容
 func (obj *_EgOrderDetailMgr) FetchIndexByOrderDetailOrder(idOrder uint32 ) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}
 
 // FetchIndexByProductID  获取多个内容
 func (obj *_EgOrderDetailMgr) FetchIndexByProductID(productID uint32 ,productAttributeID uint32 ) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id = ? AND product_attribute_id = ?", productID , productAttributeID).Find(&results).Error
	
	return
}
 
 // FetchIndexByProductAttributeID  获取多个内容
 func (obj *_EgOrderDetailMgr) FetchIndexByProductAttributeID(productAttributeID uint32 ) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_attribute_id = ?", productAttributeID).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDTaxRulesGroup  获取多个内容
 func (obj *_EgOrderDetailMgr) FetchIndexByIDTaxRulesGroup(idTaxRulesGroup uint32 ) (results []*EgOrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}
 

	

