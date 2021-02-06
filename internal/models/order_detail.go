package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderDetailMgr struct {
	*_BaseMgr
}

func OrderDetailMgr(db *gorm.DB) *_OrderDetailMgr {
	if db == nil {
		panic(fmt.Errorf("OrderDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_detail"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderDetailMgr) GetTableName() string {
	return "ps_order_detail"
}

func (obj *_OrderDetailMgr) Get() (result OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderDetailMgr) Gets() (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderDetailMgr) WithIDOrderDetail(idOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

func (obj *_OrderDetailMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

func (obj *_OrderDetailMgr) WithIDOrderInvoice(idOrderInvoice int) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

func (obj *_OrderDetailMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

func (obj *_OrderDetailMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_OrderDetailMgr) WithProductID(productID uint32) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

func (obj *_OrderDetailMgr) WithProductAttributeID(productAttributeID uint32) Option {
	return optionFunc(func(o *options) { o.query["product_attribute_id"] = productAttributeID })
}

func (obj *_OrderDetailMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

func (obj *_OrderDetailMgr) WithProductName(productName string) Option {
	return optionFunc(func(o *options) { o.query["product_name"] = productName })
}

func (obj *_OrderDetailMgr) WithProductQuantity(productQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity"] = productQuantity })
}

func (obj *_OrderDetailMgr) WithProductQuantityInStock(productQuantityInStock int) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_in_stock"] = productQuantityInStock })
}

func (obj *_OrderDetailMgr) WithProductQuantityRefunded(productQuantityRefunded uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_refunded"] = productQuantityRefunded })
}

func (obj *_OrderDetailMgr) WithProductQuantityReturn(productQuantityReturn uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_return"] = productQuantityReturn })
}

func (obj *_OrderDetailMgr) WithProductQuantityReinjected(productQuantityReinjected uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_reinjected"] = productQuantityReinjected })
}

func (obj *_OrderDetailMgr) WithProductPrice(productPrice float64) Option {
	return optionFunc(func(o *options) { o.query["product_price"] = productPrice })
}

func (obj *_OrderDetailMgr) WithReductionPercent(reductionPercent float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_percent"] = reductionPercent })
}

func (obj *_OrderDetailMgr) WithReductionAmount(reductionAmount float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_amount"] = reductionAmount })
}

func (obj *_OrderDetailMgr) WithReductionAmountTaxIncl(reductionAmountTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_amount_tax_incl"] = reductionAmountTaxIncl })
}

func (obj *_OrderDetailMgr) WithReductionAmountTaxExcl(reductionAmountTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_amount_tax_excl"] = reductionAmountTaxExcl })
}

func (obj *_OrderDetailMgr) WithGroupReduction(groupReduction float64) Option {
	return optionFunc(func(o *options) { o.query["group_reduction"] = groupReduction })
}

func (obj *_OrderDetailMgr) WithProductQuantityDiscount(productQuantityDiscount float64) Option {
	return optionFunc(func(o *options) { o.query["product_quantity_discount"] = productQuantityDiscount })
}

func (obj *_OrderDetailMgr) WithProductEan13(productEan13 string) Option {
	return optionFunc(func(o *options) { o.query["product_ean13"] = productEan13 })
}

func (obj *_OrderDetailMgr) WithProductIsbn(productIsbn string) Option {
	return optionFunc(func(o *options) { o.query["product_isbn"] = productIsbn })
}

func (obj *_OrderDetailMgr) WithProductUpc(productUpc string) Option {
	return optionFunc(func(o *options) { o.query["product_upc"] = productUpc })
}

func (obj *_OrderDetailMgr) WithProductMpn(productMpn string) Option {
	return optionFunc(func(o *options) { o.query["product_mpn"] = productMpn })
}

func (obj *_OrderDetailMgr) WithProductReference(productReference string) Option {
	return optionFunc(func(o *options) { o.query["product_reference"] = productReference })
}

func (obj *_OrderDetailMgr) WithProductSupplierReference(productSupplierReference string) Option {
	return optionFunc(func(o *options) { o.query["product_supplier_reference"] = productSupplierReference })
}

func (obj *_OrderDetailMgr) WithProductWeight(productWeight float64) Option {
	return optionFunc(func(o *options) { o.query["product_weight"] = productWeight })
}

func (obj *_OrderDetailMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

func (obj *_OrderDetailMgr) WithTaxComputationMethod(taxComputationMethod bool) Option {
	return optionFunc(func(o *options) { o.query["tax_computation_method"] = taxComputationMethod })
}

func (obj *_OrderDetailMgr) WithTaxName(taxName string) Option {
	return optionFunc(func(o *options) { o.query["tax_name"] = taxName })
}

func (obj *_OrderDetailMgr) WithTaxRate(taxRate float64) Option {
	return optionFunc(func(o *options) { o.query["tax_rate"] = taxRate })
}

func (obj *_OrderDetailMgr) WithEcotax(ecotax float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax"] = ecotax })
}

func (obj *_OrderDetailMgr) WithEcotaxTaxRate(ecotaxTaxRate float64) Option {
	return optionFunc(func(o *options) { o.query["ecotax_tax_rate"] = ecotaxTaxRate })
}

func (obj *_OrderDetailMgr) WithDiscountQuantityApplied(discountQuantityApplied bool) Option {
	return optionFunc(func(o *options) { o.query["discount_quantity_applied"] = discountQuantityApplied })
}

func (obj *_OrderDetailMgr) WithDownloadHash(downloadHash string) Option {
	return optionFunc(func(o *options) { o.query["download_hash"] = downloadHash })
}

func (obj *_OrderDetailMgr) WithDownloadNb(downloadNb uint32) Option {
	return optionFunc(func(o *options) { o.query["download_nb"] = downloadNb })
}

func (obj *_OrderDetailMgr) WithDownloadDeadline(downloadDeadline time.Time) Option {
	return optionFunc(func(o *options) { o.query["download_deadline"] = downloadDeadline })
}

func (obj *_OrderDetailMgr) WithTotalPriceTaxIncl(totalPriceTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_price_tax_incl"] = totalPriceTaxIncl })
}

func (obj *_OrderDetailMgr) WithTotalPriceTaxExcl(totalPriceTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_price_tax_excl"] = totalPriceTaxExcl })
}

func (obj *_OrderDetailMgr) WithUnitPriceTaxIncl(unitPriceTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_tax_incl"] = unitPriceTaxIncl })
}

func (obj *_OrderDetailMgr) WithUnitPriceTaxExcl(unitPriceTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price_tax_excl"] = unitPriceTaxExcl })
}

func (obj *_OrderDetailMgr) WithTotalShippingPriceTaxIncl(totalShippingPriceTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_price_tax_incl"] = totalShippingPriceTaxIncl })
}

func (obj *_OrderDetailMgr) WithTotalShippingPriceTaxExcl(totalShippingPriceTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_price_tax_excl"] = totalShippingPriceTaxExcl })
}

func (obj *_OrderDetailMgr) WithPurchaseSupplierPrice(purchaseSupplierPrice float64) Option {
	return optionFunc(func(o *options) { o.query["purchase_supplier_price"] = purchaseSupplierPrice })
}

func (obj *_OrderDetailMgr) WithOriginalProductPrice(originalProductPrice float64) Option {
	return optionFunc(func(o *options) { o.query["original_product_price"] = originalProductPrice })
}

func (obj *_OrderDetailMgr) WithOriginalWholesalePrice(originalWholesalePrice float64) Option {
	return optionFunc(func(o *options) { o.query["original_wholesale_price"] = originalWholesalePrice })
}

func (obj *_OrderDetailMgr) WithTotalRefundedTaxExcl(totalRefundedTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_refunded_tax_excl"] = totalRefundedTaxExcl })
}

func (obj *_OrderDetailMgr) WithTotalRefundedTaxIncl(totalRefundedTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_refunded_tax_incl"] = totalRefundedTaxIncl })
}

func (obj *_OrderDetailMgr) GetByOption(opts ...Option) (result OrderDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderDetailMgr) GetByOptions(opts ...Option) (results []*OrderDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_OrderDetailMgr) GetFromIDOrderDetail(idOrderDetail uint32) (result OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&result).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromIDOrderDetail(idOrderDetails []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromIDOrder(idOrder uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromIDOrderInvoice(idOrderInvoice int) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []int) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromIDShop(idShop uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromIDShop(idShops []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductID(productID uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id = ?", productID).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductID(productIDs []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id IN (?)", productIDs).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductAttributeID(productAttributeID uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_attribute_id = ?", productAttributeID).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductAttributeID(productAttributeIDs []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_attribute_id IN (?)", productAttributeIDs).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromIDCustomization(idCustomization uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductName(productName string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_name = ?", productName).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductName(productNames []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_name IN (?)", productNames).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductQuantity(productQuantity uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity = ?", productQuantity).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductQuantity(productQuantitys []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity IN (?)", productQuantitys).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductQuantityInStock(productQuantityInStock int) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_in_stock = ?", productQuantityInStock).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductQuantityInStock(productQuantityInStocks []int) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_in_stock IN (?)", productQuantityInStocks).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductQuantityRefunded(productQuantityRefunded uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_refunded = ?", productQuantityRefunded).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductQuantityRefunded(productQuantityRefundeds []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_refunded IN (?)", productQuantityRefundeds).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductQuantityReturn(productQuantityReturn uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_return = ?", productQuantityReturn).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductQuantityReturn(productQuantityReturns []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_return IN (?)", productQuantityReturns).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductQuantityReinjected(productQuantityReinjected uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_reinjected = ?", productQuantityReinjected).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductQuantityReinjected(productQuantityReinjecteds []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_reinjected IN (?)", productQuantityReinjecteds).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductPrice(productPrice float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_price = ?", productPrice).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductPrice(productPrices []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_price IN (?)", productPrices).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromReductionPercent(reductionPercent float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_percent = ?", reductionPercent).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromReductionPercent(reductionPercents []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_percent IN (?)", reductionPercents).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromReductionAmount(reductionAmount float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount = ?", reductionAmount).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromReductionAmount(reductionAmounts []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount IN (?)", reductionAmounts).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromReductionAmountTaxIncl(reductionAmountTaxIncl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount_tax_incl = ?", reductionAmountTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromReductionAmountTaxIncl(reductionAmountTaxIncls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount_tax_incl IN (?)", reductionAmountTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromReductionAmountTaxExcl(reductionAmountTaxExcl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount_tax_excl = ?", reductionAmountTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromReductionAmountTaxExcl(reductionAmountTaxExcls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount_tax_excl IN (?)", reductionAmountTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromGroupReduction(groupReduction float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_reduction = ?", groupReduction).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromGroupReduction(groupReductions []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_reduction IN (?)", groupReductions).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductQuantityDiscount(productQuantityDiscount float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_discount = ?", productQuantityDiscount).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductQuantityDiscount(productQuantityDiscounts []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity_discount IN (?)", productQuantityDiscounts).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductEan13(productEan13 string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_ean13 = ?", productEan13).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductEan13(productEan13s []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_ean13 IN (?)", productEan13s).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductIsbn(productIsbn string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_isbn = ?", productIsbn).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductIsbn(productIsbns []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_isbn IN (?)", productIsbns).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductUpc(productUpc string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_upc = ?", productUpc).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductUpc(productUpcs []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_upc IN (?)", productUpcs).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductMpn(productMpn string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_mpn = ?", productMpn).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductMpn(productMpns []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_mpn IN (?)", productMpns).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductReference(productReference string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_reference = ?", productReference).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductReference(productReferences []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_reference IN (?)", productReferences).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductSupplierReference(productSupplierReference string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_reference = ?", productSupplierReference).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductSupplierReference(productSupplierReferences []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_reference IN (?)", productSupplierReferences).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromProductWeight(productWeight float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_weight = ?", productWeight).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromProductWeight(productWeights []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_weight IN (?)", productWeights).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTaxComputationMethod(taxComputationMethod bool) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_computation_method = ?", taxComputationMethod).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTaxComputationMethod(taxComputationMethods []bool) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_computation_method IN (?)", taxComputationMethods).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTaxName(taxName string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_name = ?", taxName).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTaxName(taxNames []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_name IN (?)", taxNames).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTaxRate(taxRate float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_rate = ?", taxRate).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTaxRate(taxRates []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_rate IN (?)", taxRates).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromEcotax(ecotax float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax = ?", ecotax).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromEcotax(ecotaxs []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax IN (?)", ecotaxs).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromEcotaxTaxRate(ecotaxTaxRate float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax_tax_rate = ?", ecotaxTaxRate).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromEcotaxTaxRate(ecotaxTaxRates []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ecotax_tax_rate IN (?)", ecotaxTaxRates).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromDiscountQuantityApplied(discountQuantityApplied bool) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_quantity_applied = ?", discountQuantityApplied).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromDiscountQuantityApplied(discountQuantityApplieds []bool) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_quantity_applied IN (?)", discountQuantityApplieds).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromDownloadHash(downloadHash string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_hash = ?", downloadHash).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromDownloadHash(downloadHashs []string) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_hash IN (?)", downloadHashs).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromDownloadNb(downloadNb uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_nb = ?", downloadNb).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromDownloadNb(downloadNbs []uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_nb IN (?)", downloadNbs).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromDownloadDeadline(downloadDeadline time.Time) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_deadline = ?", downloadDeadline).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromDownloadDeadline(downloadDeadlines []time.Time) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("download_deadline IN (?)", downloadDeadlines).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTotalPriceTaxIncl(totalPriceTaxIncl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_incl = ?", totalPriceTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTotalPriceTaxIncl(totalPriceTaxIncls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_incl IN (?)", totalPriceTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTotalPriceTaxExcl(totalPriceTaxExcl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_excl = ?", totalPriceTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTotalPriceTaxExcl(totalPriceTaxExcls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_price_tax_excl IN (?)", totalPriceTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromUnitPriceTaxIncl(unitPriceTaxIncl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_incl = ?", unitPriceTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromUnitPriceTaxIncl(unitPriceTaxIncls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_incl IN (?)", unitPriceTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromUnitPriceTaxExcl(unitPriceTaxExcl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_excl = ?", unitPriceTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromUnitPriceTaxExcl(unitPriceTaxExcls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("unit_price_tax_excl IN (?)", unitPriceTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTotalShippingPriceTaxIncl(totalShippingPriceTaxIncl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_price_tax_incl = ?", totalShippingPriceTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTotalShippingPriceTaxIncl(totalShippingPriceTaxIncls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_price_tax_incl IN (?)", totalShippingPriceTaxIncls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTotalShippingPriceTaxExcl(totalShippingPriceTaxExcl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_price_tax_excl = ?", totalShippingPriceTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTotalShippingPriceTaxExcl(totalShippingPriceTaxExcls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_price_tax_excl IN (?)", totalShippingPriceTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromPurchaseSupplierPrice(purchaseSupplierPrice float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("purchase_supplier_price = ?", purchaseSupplierPrice).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromPurchaseSupplierPrice(purchaseSupplierPrices []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("purchase_supplier_price IN (?)", purchaseSupplierPrices).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromOriginalProductPrice(originalProductPrice float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_product_price = ?", originalProductPrice).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromOriginalProductPrice(originalProductPrices []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_product_price IN (?)", originalProductPrices).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromOriginalWholesalePrice(originalWholesalePrice float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_wholesale_price = ?", originalWholesalePrice).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromOriginalWholesalePrice(originalWholesalePrices []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("original_wholesale_price IN (?)", originalWholesalePrices).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTotalRefundedTaxExcl(totalRefundedTaxExcl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_refunded_tax_excl = ?", totalRefundedTaxExcl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTotalRefundedTaxExcl(totalRefundedTaxExcls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_refunded_tax_excl IN (?)", totalRefundedTaxExcls).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetFromTotalRefundedTaxIncl(totalRefundedTaxIncl float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_refunded_tax_incl = ?", totalRefundedTaxIncl).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) GetBatchFromTotalRefundedTaxIncl(totalRefundedTaxIncls []float64) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_refunded_tax_incl IN (?)", totalRefundedTaxIncls).Find(&results).Error

	return
}


func (obj *_OrderDetailMgr) FetchByPrimaryKey(idOrderDetail uint32) (result OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&result).Error

	return
}

func (obj *_OrderDetailMgr) FetchIndexByIDOrderIDOrderDetail(idOrderDetail uint32, idOrder uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ? AND id_order = ?", idOrderDetail, idOrder).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) FetchIndexByOrderDetailOrder(idOrder uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) FetchIndexByProductID(productID uint32, productAttributeID uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_id = ? AND product_attribute_id = ?", productID, productAttributeID).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) FetchIndexByProductAttributeID(productAttributeID uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_attribute_id = ?", productAttributeID).Find(&results).Error

	return
}

func (obj *_OrderDetailMgr) FetchIndexByIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}
