package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SupplyOrderMgr struct {
	*_BaseMgr
}

func SupplyOrderMgr(db *gorm.DB) *_SupplyOrderMgr {
	if db == nil {
		panic(fmt.Errorf("SupplyOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplyOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supply_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SupplyOrderMgr) GetTableName() string {
	return "ps_supply_order"
}

func (obj *_SupplyOrderMgr) Get() (result SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SupplyOrderMgr) Gets() (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SupplyOrderMgr) WithIDSupplyOrder(idSupplyOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order"] = idSupplyOrder })
}

func (obj *_SupplyOrderMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

func (obj *_SupplyOrderMgr) WithSupplierName(supplierName string) Option {
	return optionFunc(func(o *options) { o.query["supplier_name"] = supplierName })
}

func (obj *_SupplyOrderMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_SupplyOrderMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

func (obj *_SupplyOrderMgr) WithIDSupplyOrderState(idSupplyOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_state"] = idSupplyOrderState })
}

func (obj *_SupplyOrderMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_SupplyOrderMgr) WithIDRefCurrency(idRefCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_ref_currency"] = idRefCurrency })
}

func (obj *_SupplyOrderMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

func (obj *_SupplyOrderMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_SupplyOrderMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_SupplyOrderMgr) WithDateDeliveryExpected(dateDeliveryExpected time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_delivery_expected"] = dateDeliveryExpected })
}

func (obj *_SupplyOrderMgr) WithTotalTe(totalTe float64) Option {
	return optionFunc(func(o *options) { o.query["total_te"] = totalTe })
}

func (obj *_SupplyOrderMgr) WithTotalWithDiscountTe(totalWithDiscountTe float64) Option {
	return optionFunc(func(o *options) { o.query["total_with_discount_te"] = totalWithDiscountTe })
}

func (obj *_SupplyOrderMgr) WithTotalTax(totalTax float64) Option {
	return optionFunc(func(o *options) { o.query["total_tax"] = totalTax })
}

func (obj *_SupplyOrderMgr) WithTotalTi(totalTi float64) Option {
	return optionFunc(func(o *options) { o.query["total_ti"] = totalTi })
}

func (obj *_SupplyOrderMgr) WithDiscountRate(discountRate float64) Option {
	return optionFunc(func(o *options) { o.query["discount_rate"] = discountRate })
}

func (obj *_SupplyOrderMgr) WithDiscountValueTe(discountValueTe float64) Option {
	return optionFunc(func(o *options) { o.query["discount_value_te"] = discountValueTe })
}

func (obj *_SupplyOrderMgr) WithIsTemplate(isTemplate bool) Option {
	return optionFunc(func(o *options) { o.query["is_template"] = isTemplate })
}

func (obj *_SupplyOrderMgr) GetByOption(opts ...Option) (result SupplyOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SupplyOrderMgr) GetByOptions(opts ...Option) (results []*SupplyOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SupplyOrderMgr) GetFromIDSupplyOrder(idSupplyOrder uint32) (result SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&result).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromIDSupplyOrder(idSupplyOrders []uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order IN (?)", idSupplyOrders).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromIDSupplier(idSupplier uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromSupplierName(supplierName string) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_name = ?", supplierName).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromSupplierName(supplierNames []string) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("supplier_name IN (?)", supplierNames).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromIDLang(idLang uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromIDLang(idLangs []uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromIDSupplyOrderState(idSupplyOrderState uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromIDSupplyOrderState(idSupplyOrderStates []uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state IN (?)", idSupplyOrderStates).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromIDCurrency(idCurrency uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromIDRefCurrency(idRefCurrency uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ref_currency = ?", idRefCurrency).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromIDRefCurrency(idRefCurrencys []uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ref_currency IN (?)", idRefCurrencys).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromReference(reference string) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromReference(references []string) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromDateAdd(dateAdd time.Time) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromDateUpd(dateUpd time.Time) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromDateDeliveryExpected(dateDeliveryExpected time.Time) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_delivery_expected = ?", dateDeliveryExpected).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromDateDeliveryExpected(dateDeliveryExpecteds []time.Time) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_delivery_expected IN (?)", dateDeliveryExpecteds).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromTotalTe(totalTe float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_te = ?", totalTe).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromTotalTe(totalTes []float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_te IN (?)", totalTes).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromTotalWithDiscountTe(totalWithDiscountTe float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_with_discount_te = ?", totalWithDiscountTe).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromTotalWithDiscountTe(totalWithDiscountTes []float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_with_discount_te IN (?)", totalWithDiscountTes).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromTotalTax(totalTax float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_tax = ?", totalTax).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromTotalTax(totalTaxs []float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_tax IN (?)", totalTaxs).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromTotalTi(totalTi float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_ti = ?", totalTi).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromTotalTi(totalTis []float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_ti IN (?)", totalTis).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromDiscountRate(discountRate float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_rate = ?", discountRate).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromDiscountRate(discountRates []float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_rate IN (?)", discountRates).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromDiscountValueTe(discountValueTe float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_value_te = ?", discountValueTe).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromDiscountValueTe(discountValueTes []float64) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("discount_value_te IN (?)", discountValueTes).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetFromIsTemplate(isTemplate bool) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_template = ?", isTemplate).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) GetBatchFromIsTemplate(isTemplates []bool) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_template IN (?)", isTemplates).Find(&results).Error

	return
}


func (obj *_SupplyOrderMgr) FetchByPrimaryKey(idSupplyOrder uint32) (result SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order = ?", idSupplyOrder).Find(&result).Error

	return
}

func (obj *_SupplyOrderMgr) FetchIndexByIDSupplier(idSupplier uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) FetchIndexByIDWarehouse(idWarehouse uint32) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error

	return
}

func (obj *_SupplyOrderMgr) FetchIndexByReference(reference string) (results []*SupplyOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error

	return
}
