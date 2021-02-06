package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _OrderSlipMgr struct {
	*_BaseMgr
}

// OrderSlipMgr open func
func OrderSlipMgr(db *gorm.DB) *_OrderSlipMgr {
	if db == nil {
		panic(fmt.Errorf("OrderSlipMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderSlipMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_slip"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderSlipMgr) GetTableName() string {
	return "ps_order_slip"
}

// Get 获取
func (obj *_OrderSlipMgr) Get() (result OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderSlipMgr) Gets() (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderSlip id_order_slip获取
func (obj *_OrderSlipMgr) WithIDOrderSlip(idOrderSlip uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_slip"] = idOrderSlip })
}

// WithConversionRate conversion_rate获取
func (obj *_OrderSlipMgr) WithConversionRate(conversionRate float64) Option {
	return optionFunc(func(o *options) { o.query["conversion_rate"] = conversionRate })
}

// WithIDCustomer id_customer获取
func (obj *_OrderSlipMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDOrder id_order获取
func (obj *_OrderSlipMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithTotalProductsTaxExcl total_products_tax_excl获取
func (obj *_OrderSlipMgr) WithTotalProductsTaxExcl(totalProductsTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_products_tax_excl"] = totalProductsTaxExcl })
}

// WithTotalProductsTaxIncl total_products_tax_incl获取
func (obj *_OrderSlipMgr) WithTotalProductsTaxIncl(totalProductsTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_products_tax_incl"] = totalProductsTaxIncl })
}

// WithTotalShippingTaxExcl total_shipping_tax_excl获取
func (obj *_OrderSlipMgr) WithTotalShippingTaxExcl(totalShippingTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_excl"] = totalShippingTaxExcl })
}

// WithTotalShippingTaxIncl total_shipping_tax_incl获取
func (obj *_OrderSlipMgr) WithTotalShippingTaxIncl(totalShippingTaxIncl float64) Option {
	return optionFunc(func(o *options) { o.query["total_shipping_tax_incl"] = totalShippingTaxIncl })
}

// WithShippingCost shipping_cost获取
func (obj *_OrderSlipMgr) WithShippingCost(shippingCost uint8) Option {
	return optionFunc(func(o *options) { o.query["shipping_cost"] = shippingCost })
}

// WithAmount amount获取
func (obj *_OrderSlipMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithShippingCostAmount shipping_cost_amount获取
func (obj *_OrderSlipMgr) WithShippingCostAmount(shippingCostAmount float64) Option {
	return optionFunc(func(o *options) { o.query["shipping_cost_amount"] = shippingCostAmount })
}

// WithPartial partial获取
func (obj *_OrderSlipMgr) WithPartial(partial bool) Option {
	return optionFunc(func(o *options) { o.query["partial"] = partial })
}

// WithOrderSlipType order_slip_type获取
func (obj *_OrderSlipMgr) WithOrderSlipType(orderSlipType bool) Option {
	return optionFunc(func(o *options) { o.query["order_slip_type"] = orderSlipType })
}

// WithDateAdd date_add获取
func (obj *_OrderSlipMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_OrderSlipMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_OrderSlipMgr) GetByOption(opts ...Option) (result OrderSlip, err error) {
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
func (obj *_OrderSlipMgr) GetByOptions(opts ...Option) (results []*OrderSlip, err error) {
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

// GetFromIDOrderSlip 通过id_order_slip获取内容
func (obj *_OrderSlipMgr) GetFromIDOrderSlip(idOrderSlip uint32) (result OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip = ?", idOrderSlip).Find(&result).Error

	return
}

// GetBatchFromIDOrderSlip 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromIDOrderSlip(idOrderSlips []uint32) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip IN (?)", idOrderSlips).Find(&results).Error

	return
}

// GetFromConversionRate 通过conversion_rate获取内容
func (obj *_OrderSlipMgr) GetFromConversionRate(conversionRate float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate = ?", conversionRate).Find(&results).Error

	return
}

// GetBatchFromConversionRate 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromConversionRate(conversionRates []float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("conversion_rate IN (?)", conversionRates).Find(&results).Error

	return
}

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_OrderSlipMgr) GetFromIDCustomer(idCustomer uint32) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_OrderSlipMgr) GetFromIDOrder(idOrder uint32) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromTotalProductsTaxExcl 通过total_products_tax_excl获取内容
func (obj *_OrderSlipMgr) GetFromTotalProductsTaxExcl(totalProductsTaxExcl float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_tax_excl = ?", totalProductsTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalProductsTaxExcl 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromTotalProductsTaxExcl(totalProductsTaxExcls []float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_tax_excl IN (?)", totalProductsTaxExcls).Find(&results).Error

	return
}

// GetFromTotalProductsTaxIncl 通过total_products_tax_incl获取内容
func (obj *_OrderSlipMgr) GetFromTotalProductsTaxIncl(totalProductsTaxIncl float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_tax_incl = ?", totalProductsTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalProductsTaxIncl 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromTotalProductsTaxIncl(totalProductsTaxIncls []float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_products_tax_incl IN (?)", totalProductsTaxIncls).Find(&results).Error

	return
}

// GetFromTotalShippingTaxExcl 通过total_shipping_tax_excl获取内容
func (obj *_OrderSlipMgr) GetFromTotalShippingTaxExcl(totalShippingTaxExcl float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl = ?", totalShippingTaxExcl).Find(&results).Error

	return
}

// GetBatchFromTotalShippingTaxExcl 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromTotalShippingTaxExcl(totalShippingTaxExcls []float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_excl IN (?)", totalShippingTaxExcls).Find(&results).Error

	return
}

// GetFromTotalShippingTaxIncl 通过total_shipping_tax_incl获取内容
func (obj *_OrderSlipMgr) GetFromTotalShippingTaxIncl(totalShippingTaxIncl float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl = ?", totalShippingTaxIncl).Find(&results).Error

	return
}

// GetBatchFromTotalShippingTaxIncl 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromTotalShippingTaxIncl(totalShippingTaxIncls []float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_shipping_tax_incl IN (?)", totalShippingTaxIncls).Find(&results).Error

	return
}

// GetFromShippingCost 通过shipping_cost获取内容
func (obj *_OrderSlipMgr) GetFromShippingCost(shippingCost uint8) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_cost = ?", shippingCost).Find(&results).Error

	return
}

// GetBatchFromShippingCost 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromShippingCost(shippingCosts []uint8) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_cost IN (?)", shippingCosts).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容
func (obj *_OrderSlipMgr) GetFromAmount(amount float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromAmount(amounts []float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromShippingCostAmount 通过shipping_cost_amount获取内容
func (obj *_OrderSlipMgr) GetFromShippingCostAmount(shippingCostAmount float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_cost_amount = ?", shippingCostAmount).Find(&results).Error

	return
}

// GetBatchFromShippingCostAmount 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromShippingCostAmount(shippingCostAmounts []float64) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_cost_amount IN (?)", shippingCostAmounts).Find(&results).Error

	return
}

// GetFromPartial 通过partial获取内容
func (obj *_OrderSlipMgr) GetFromPartial(partial bool) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("partial = ?", partial).Find(&results).Error

	return
}

// GetBatchFromPartial 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromPartial(partials []bool) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("partial IN (?)", partials).Find(&results).Error

	return
}

// GetFromOrderSlipType 通过order_slip_type获取内容
func (obj *_OrderSlipMgr) GetFromOrderSlipType(orderSlipType bool) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_slip_type = ?", orderSlipType).Find(&results).Error

	return
}

// GetBatchFromOrderSlipType 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromOrderSlipType(orderSlipTypes []bool) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_slip_type IN (?)", orderSlipTypes).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_OrderSlipMgr) GetFromDateAdd(dateAdd time.Time) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_OrderSlipMgr) GetFromDateUpd(dateUpd time.Time) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_OrderSlipMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderSlipMgr) FetchByPrimaryKey(idOrderSlip uint32) (result OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_slip = ?", idOrderSlip).Find(&result).Error

	return
}

// FetchIndexByOrderSlipCustomer  获取多个内容
func (obj *_OrderSlipMgr) FetchIndexByOrderSlipCustomer(idCustomer uint32) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// FetchIndexByIDOrder  获取多个内容
func (obj *_OrderSlipMgr) FetchIndexByIDOrder(idOrder uint32) (results []*OrderSlip, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}
