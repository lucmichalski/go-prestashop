package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CartRuleMgr struct {
	*_BaseMgr
}

func CartRuleMgr(db *gorm.DB) *_CartRuleMgr {
	if db == nil {
		panic(fmt.Errorf("CartRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cart_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CartRuleMgr) GetTableName() string {
	return "ps_cart_rule"
}

func (obj *_CartRuleMgr) Get() (result CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CartRuleMgr) Gets() (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CartRuleMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

func (obj *_CartRuleMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_CartRuleMgr) WithDateFrom(dateFrom time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_from"] = dateFrom })
}

func (obj *_CartRuleMgr) WithDateTo(dateTo time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_to"] = dateTo })
}

func (obj *_CartRuleMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_CartRuleMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

func (obj *_CartRuleMgr) WithQuantityPerUser(quantityPerUser uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity_per_user"] = quantityPerUser })
}

func (obj *_CartRuleMgr) WithPriority(priority uint32) Option {
	return optionFunc(func(o *options) { o.query["priority"] = priority })
}

func (obj *_CartRuleMgr) WithPartialUse(partialUse bool) Option {
	return optionFunc(func(o *options) { o.query["partial_use"] = partialUse })
}

func (obj *_CartRuleMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

func (obj *_CartRuleMgr) WithMinimumAmount(minimumAmount float64) Option {
	return optionFunc(func(o *options) { o.query["minimum_amount"] = minimumAmount })
}

func (obj *_CartRuleMgr) WithMinimumAmountTax(minimumAmountTax bool) Option {
	return optionFunc(func(o *options) { o.query["minimum_amount_tax"] = minimumAmountTax })
}

func (obj *_CartRuleMgr) WithMinimumAmountCurrency(minimumAmountCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["minimum_amount_currency"] = minimumAmountCurrency })
}

func (obj *_CartRuleMgr) WithMinimumAmountShipping(minimumAmountShipping bool) Option {
	return optionFunc(func(o *options) { o.query["minimum_amount_shipping"] = minimumAmountShipping })
}

func (obj *_CartRuleMgr) WithCountryRestriction(countryRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["country_restriction"] = countryRestriction })
}

func (obj *_CartRuleMgr) WithCarrierRestriction(carrierRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["carrier_restriction"] = carrierRestriction })
}

func (obj *_CartRuleMgr) WithGroupRestriction(groupRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["group_restriction"] = groupRestriction })
}

func (obj *_CartRuleMgr) WithCartRuleRestriction(cartRuleRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["cart_rule_restriction"] = cartRuleRestriction })
}

func (obj *_CartRuleMgr) WithProductRestriction(productRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["product_restriction"] = productRestriction })
}

func (obj *_CartRuleMgr) WithShopRestriction(shopRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["shop_restriction"] = shopRestriction })
}

func (obj *_CartRuleMgr) WithFreeShipping(freeShipping bool) Option {
	return optionFunc(func(o *options) { o.query["free_shipping"] = freeShipping })
}

func (obj *_CartRuleMgr) WithReductionPercent(reductionPercent float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_percent"] = reductionPercent })
}

func (obj *_CartRuleMgr) WithReductionAmount(reductionAmount float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_amount"] = reductionAmount })
}

func (obj *_CartRuleMgr) WithReductionTax(reductionTax bool) Option {
	return optionFunc(func(o *options) { o.query["reduction_tax"] = reductionTax })
}

func (obj *_CartRuleMgr) WithReductionCurrency(reductionCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["reduction_currency"] = reductionCurrency })
}

func (obj *_CartRuleMgr) WithReductionProduct(reductionProduct int) Option {
	return optionFunc(func(o *options) { o.query["reduction_product"] = reductionProduct })
}

func (obj *_CartRuleMgr) WithReductionExcludeSpecial(reductionExcludeSpecial bool) Option {
	return optionFunc(func(o *options) { o.query["reduction_exclude_special"] = reductionExcludeSpecial })
}

func (obj *_CartRuleMgr) WithGiftProduct(giftProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["gift_product"] = giftProduct })
}

func (obj *_CartRuleMgr) WithGiftProductAttribute(giftProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["gift_product_attribute"] = giftProductAttribute })
}

func (obj *_CartRuleMgr) WithHighlight(highlight bool) Option {
	return optionFunc(func(o *options) { o.query["highlight"] = highlight })
}

func (obj *_CartRuleMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_CartRuleMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_CartRuleMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_CartRuleMgr) GetByOption(opts ...Option) (result CartRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CartRuleMgr) GetByOptions(opts ...Option) (results []*CartRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CartRuleMgr) GetFromIDCartRule(idCartRule uint32) (result CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&result).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromIDCustomer(idCustomer uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromDateFrom(dateFrom time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_from = ?", dateFrom).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromDateFrom(dateFroms []time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_from IN (?)", dateFroms).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromDateTo(dateTo time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to = ?", dateTo).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromDateTo(dateTos []time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to IN (?)", dateTos).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromDescription(description string) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromDescription(descriptions []string) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromQuantity(quantity uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromQuantity(quantitys []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromQuantityPerUser(quantityPerUser uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_per_user = ?", quantityPerUser).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromQuantityPerUser(quantityPerUsers []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_per_user IN (?)", quantityPerUsers).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromPriority(priority uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("priority = ?", priority).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromPriority(prioritys []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("priority IN (?)", prioritys).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromPartialUse(partialUse bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("partial_use = ?", partialUse).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromPartialUse(partialUses []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("partial_use IN (?)", partialUses).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromCode(code string) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code = ?", code).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromCode(codes []string) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromMinimumAmount(minimumAmount float64) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount = ?", minimumAmount).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromMinimumAmount(minimumAmounts []float64) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount IN (?)", minimumAmounts).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromMinimumAmountTax(minimumAmountTax bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_tax = ?", minimumAmountTax).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromMinimumAmountTax(minimumAmountTaxs []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_tax IN (?)", minimumAmountTaxs).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromMinimumAmountCurrency(minimumAmountCurrency uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_currency = ?", minimumAmountCurrency).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromMinimumAmountCurrency(minimumAmountCurrencys []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_currency IN (?)", minimumAmountCurrencys).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromMinimumAmountShipping(minimumAmountShipping bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_shipping = ?", minimumAmountShipping).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromMinimumAmountShipping(minimumAmountShippings []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_shipping IN (?)", minimumAmountShippings).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromCountryRestriction(countryRestriction bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_restriction = ?", countryRestriction).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromCountryRestriction(countryRestrictions []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_restriction IN (?)", countryRestrictions).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromCarrierRestriction(carrierRestriction bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carrier_restriction = ?", carrierRestriction).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromCarrierRestriction(carrierRestrictions []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carrier_restriction IN (?)", carrierRestrictions).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromGroupRestriction(groupRestriction bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_restriction = ?", groupRestriction).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromGroupRestriction(groupRestrictions []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_restriction IN (?)", groupRestrictions).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromCartRuleRestriction(cartRuleRestriction bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cart_rule_restriction = ?", cartRuleRestriction).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromCartRuleRestriction(cartRuleRestrictions []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cart_rule_restriction IN (?)", cartRuleRestrictions).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromProductRestriction(productRestriction bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_restriction = ?", productRestriction).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromProductRestriction(productRestrictions []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_restriction IN (?)", productRestrictions).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromShopRestriction(shopRestriction bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop_restriction = ?", shopRestriction).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromShopRestriction(shopRestrictions []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop_restriction IN (?)", shopRestrictions).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromFreeShipping(freeShipping bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping = ?", freeShipping).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromFreeShipping(freeShippings []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping IN (?)", freeShippings).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromReductionPercent(reductionPercent float64) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_percent = ?", reductionPercent).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromReductionPercent(reductionPercents []float64) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_percent IN (?)", reductionPercents).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromReductionAmount(reductionAmount float64) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount = ?", reductionAmount).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromReductionAmount(reductionAmounts []float64) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount IN (?)", reductionAmounts).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromReductionTax(reductionTax bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax = ?", reductionTax).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromReductionTax(reductionTaxs []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax IN (?)", reductionTaxs).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromReductionCurrency(reductionCurrency uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_currency = ?", reductionCurrency).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromReductionCurrency(reductionCurrencys []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_currency IN (?)", reductionCurrencys).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromReductionProduct(reductionProduct int) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_product = ?", reductionProduct).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromReductionProduct(reductionProducts []int) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_product IN (?)", reductionProducts).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromReductionExcludeSpecial(reductionExcludeSpecial bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_exclude_special = ?", reductionExcludeSpecial).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromReductionExcludeSpecial(reductionExcludeSpecials []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_exclude_special IN (?)", reductionExcludeSpecials).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromGiftProduct(giftProduct uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_product = ?", giftProduct).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromGiftProduct(giftProducts []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_product IN (?)", giftProducts).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromGiftProductAttribute(giftProductAttribute uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_product_attribute = ?", giftProductAttribute).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromGiftProductAttribute(giftProductAttributes []uint32) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_product_attribute IN (?)", giftProductAttributes).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromHighlight(highlight bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("highlight = ?", highlight).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromHighlight(highlights []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("highlight IN (?)", highlights).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromActive(active bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromActive(actives []bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromDateAdd(dateAdd time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetFromDateUpd(dateUpd time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_CartRuleMgr) FetchByPrimaryKey(idCartRule uint32) (result CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&result).Error

	return
}

func (obj *_CartRuleMgr) FetchIndexByIDCustomer(idCustomer uint32, dateTo time.Time, active bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND date_to = ? AND active = ?", idCustomer, dateTo, active).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) FetchIndexByIDCustomer2(idCustomer uint32, dateTo time.Time, highlight bool, active bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND date_to = ? AND highlight = ? AND active = ?", idCustomer, dateTo, highlight, active).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) FetchIndexByDateFrom(dateFrom time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_from = ?", dateFrom).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) FetchIndexByGroupRestriction(dateTo time.Time, groupRestriction bool, active bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to = ? AND group_restriction = ? AND active = ?", dateTo, groupRestriction, active).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) FetchIndexByGroupRestriction2(dateTo time.Time, groupRestriction bool, highlight bool, active bool) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to = ? AND group_restriction = ? AND highlight = ? AND active = ?", dateTo, groupRestriction, highlight, active).Find(&results).Error

	return
}

func (obj *_CartRuleMgr) FetchIndexByDateTo(dateTo time.Time) (results []*CartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to = ?", dateTo).Find(&results).Error

	return
}
