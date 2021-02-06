package	model	
import (	
"time"	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCartRuleMgr struct {
	*_BaseMgr
}

// EgCartRuleMgr open func
func EgCartRuleMgr(db *gorm.DB) *_EgCartRuleMgr {
	if db == nil {
		panic(fmt.Errorf("EgCartRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCartRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cart_rule"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCartRuleMgr) GetTableName() string {
	return "eg_cart_rule"
}

// Get 获取
func (obj *_EgCartRuleMgr) Get() (result EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCartRuleMgr) Gets() (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCartRule id_cart_rule获取 
func (obj *_EgCartRuleMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithIDCustomer id_customer获取 
func (obj *_EgCartRuleMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithDateFrom date_from获取 
func (obj *_EgCartRuleMgr) WithDateFrom(dateFrom time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_from"] = dateFrom })
}

// WithDateTo date_to获取 
func (obj *_EgCartRuleMgr) WithDateTo(dateTo time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_to"] = dateTo })
}

// WithDescription description获取 
func (obj *_EgCartRuleMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithQuantity quantity获取 
func (obj *_EgCartRuleMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithQuantityPerUser quantity_per_user获取 
func (obj *_EgCartRuleMgr) WithQuantityPerUser(quantityPerUser uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity_per_user"] = quantityPerUser })
}

// WithPriority priority获取 
func (obj *_EgCartRuleMgr) WithPriority(priority uint32) Option {
	return optionFunc(func(o *options) { o.query["priority"] = priority })
}

// WithPartialUse partial_use获取 
func (obj *_EgCartRuleMgr) WithPartialUse(partialUse bool) Option {
	return optionFunc(func(o *options) { o.query["partial_use"] = partialUse })
}

// WithCode code获取 
func (obj *_EgCartRuleMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithMinimumAmount minimum_amount获取 
func (obj *_EgCartRuleMgr) WithMinimumAmount(minimumAmount float64) Option {
	return optionFunc(func(o *options) { o.query["minimum_amount"] = minimumAmount })
}

// WithMinimumAmountTax minimum_amount_tax获取 
func (obj *_EgCartRuleMgr) WithMinimumAmountTax(minimumAmountTax bool) Option {
	return optionFunc(func(o *options) { o.query["minimum_amount_tax"] = minimumAmountTax })
}

// WithMinimumAmountCurrency minimum_amount_currency获取 
func (obj *_EgCartRuleMgr) WithMinimumAmountCurrency(minimumAmountCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["minimum_amount_currency"] = minimumAmountCurrency })
}

// WithMinimumAmountShipping minimum_amount_shipping获取 
func (obj *_EgCartRuleMgr) WithMinimumAmountShipping(minimumAmountShipping bool) Option {
	return optionFunc(func(o *options) { o.query["minimum_amount_shipping"] = minimumAmountShipping })
}

// WithCountryRestriction country_restriction获取 
func (obj *_EgCartRuleMgr) WithCountryRestriction(countryRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["country_restriction"] = countryRestriction })
}

// WithCarrierRestriction carrier_restriction获取 
func (obj *_EgCartRuleMgr) WithCarrierRestriction(carrierRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["carrier_restriction"] = carrierRestriction })
}

// WithGroupRestriction group_restriction获取 
func (obj *_EgCartRuleMgr) WithGroupRestriction(groupRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["group_restriction"] = groupRestriction })
}

// WithCartRuleRestriction cart_rule_restriction获取 
func (obj *_EgCartRuleMgr) WithCartRuleRestriction(cartRuleRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["cart_rule_restriction"] = cartRuleRestriction })
}

// WithProductRestriction product_restriction获取 
func (obj *_EgCartRuleMgr) WithProductRestriction(productRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["product_restriction"] = productRestriction })
}

// WithShopRestriction shop_restriction获取 
func (obj *_EgCartRuleMgr) WithShopRestriction(shopRestriction bool) Option {
	return optionFunc(func(o *options) { o.query["shop_restriction"] = shopRestriction })
}

// WithFreeShipping free_shipping获取 
func (obj *_EgCartRuleMgr) WithFreeShipping(freeShipping bool) Option {
	return optionFunc(func(o *options) { o.query["free_shipping"] = freeShipping })
}

// WithReductionPercent reduction_percent获取 
func (obj *_EgCartRuleMgr) WithReductionPercent(reductionPercent float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_percent"] = reductionPercent })
}

// WithReductionAmount reduction_amount获取 
func (obj *_EgCartRuleMgr) WithReductionAmount(reductionAmount float64) Option {
	return optionFunc(func(o *options) { o.query["reduction_amount"] = reductionAmount })
}

// WithReductionTax reduction_tax获取 
func (obj *_EgCartRuleMgr) WithReductionTax(reductionTax bool) Option {
	return optionFunc(func(o *options) { o.query["reduction_tax"] = reductionTax })
}

// WithReductionCurrency reduction_currency获取 
func (obj *_EgCartRuleMgr) WithReductionCurrency(reductionCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["reduction_currency"] = reductionCurrency })
}

// WithReductionProduct reduction_product获取 
func (obj *_EgCartRuleMgr) WithReductionProduct(reductionProduct int) Option {
	return optionFunc(func(o *options) { o.query["reduction_product"] = reductionProduct })
}

// WithReductionExcludeSpecial reduction_exclude_special获取 
func (obj *_EgCartRuleMgr) WithReductionExcludeSpecial(reductionExcludeSpecial bool) Option {
	return optionFunc(func(o *options) { o.query["reduction_exclude_special"] = reductionExcludeSpecial })
}

// WithGiftProduct gift_product获取 
func (obj *_EgCartRuleMgr) WithGiftProduct(giftProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["gift_product"] = giftProduct })
}

// WithGiftProductAttribute gift_product_attribute获取 
func (obj *_EgCartRuleMgr) WithGiftProductAttribute(giftProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["gift_product_attribute"] = giftProductAttribute })
}

// WithHighlight highlight获取 
func (obj *_EgCartRuleMgr) WithHighlight(highlight bool) Option {
	return optionFunc(func(o *options) { o.query["highlight"] = highlight })
}

// WithActive active获取 
func (obj *_EgCartRuleMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDateAdd date_add获取 
func (obj *_EgCartRuleMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgCartRuleMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgCartRuleMgr) GetByOption(opts ...Option) (result EgCartRule, err error) {
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
func (obj *_EgCartRuleMgr) GetByOptions(opts ...Option) (results []*EgCartRule, err error) {
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


// GetFromIDCartRule 通过id_cart_rule获取内容  
func (obj *_EgCartRuleMgr)  GetFromIDCartRule(idCartRule uint32) (result EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&result).Error
	
	return
}

// GetBatchFromIDCartRule 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgCartRuleMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromDateFrom 通过date_from获取内容  
func (obj *_EgCartRuleMgr) GetFromDateFrom(dateFrom time.Time) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_from = ?", dateFrom).Find(&results).Error
	
	return
}

// GetBatchFromDateFrom 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromDateFrom(dateFroms []time.Time) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_from IN (?)", dateFroms).Find(&results).Error
	
	return
}
 
// GetFromDateTo 通过date_to获取内容  
func (obj *_EgCartRuleMgr) GetFromDateTo(dateTo time.Time) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to = ?", dateTo).Find(&results).Error
	
	return
}

// GetBatchFromDateTo 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromDateTo(dateTos []time.Time) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to IN (?)", dateTos).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgCartRuleMgr) GetFromDescription(description string) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromDescription(descriptions []string) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromQuantity 通过quantity获取内容  
func (obj *_EgCartRuleMgr) GetFromQuantity(quantity uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error
	
	return
}

// GetBatchFromQuantity 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromQuantity(quantitys []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error
	
	return
}
 
// GetFromQuantityPerUser 通过quantity_per_user获取内容  
func (obj *_EgCartRuleMgr) GetFromQuantityPerUser(quantityPerUser uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_per_user = ?", quantityPerUser).Find(&results).Error
	
	return
}

// GetBatchFromQuantityPerUser 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromQuantityPerUser(quantityPerUsers []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_per_user IN (?)", quantityPerUsers).Find(&results).Error
	
	return
}
 
// GetFromPriority 通过priority获取内容  
func (obj *_EgCartRuleMgr) GetFromPriority(priority uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("priority = ?", priority).Find(&results).Error
	
	return
}

// GetBatchFromPriority 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromPriority(prioritys []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("priority IN (?)", prioritys).Find(&results).Error
	
	return
}
 
// GetFromPartialUse 通过partial_use获取内容  
func (obj *_EgCartRuleMgr) GetFromPartialUse(partialUse bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("partial_use = ?", partialUse).Find(&results).Error
	
	return
}

// GetBatchFromPartialUse 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromPartialUse(partialUses []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("partial_use IN (?)", partialUses).Find(&results).Error
	
	return
}
 
// GetFromCode 通过code获取内容  
func (obj *_EgCartRuleMgr) GetFromCode(code string) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code = ?", code).Find(&results).Error
	
	return
}

// GetBatchFromCode 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromCode(codes []string) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("code IN (?)", codes).Find(&results).Error
	
	return
}
 
// GetFromMinimumAmount 通过minimum_amount获取内容  
func (obj *_EgCartRuleMgr) GetFromMinimumAmount(minimumAmount float64) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount = ?", minimumAmount).Find(&results).Error
	
	return
}

// GetBatchFromMinimumAmount 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromMinimumAmount(minimumAmounts []float64) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount IN (?)", minimumAmounts).Find(&results).Error
	
	return
}
 
// GetFromMinimumAmountTax 通过minimum_amount_tax获取内容  
func (obj *_EgCartRuleMgr) GetFromMinimumAmountTax(minimumAmountTax bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_tax = ?", minimumAmountTax).Find(&results).Error
	
	return
}

// GetBatchFromMinimumAmountTax 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromMinimumAmountTax(minimumAmountTaxs []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_tax IN (?)", minimumAmountTaxs).Find(&results).Error
	
	return
}
 
// GetFromMinimumAmountCurrency 通过minimum_amount_currency获取内容  
func (obj *_EgCartRuleMgr) GetFromMinimumAmountCurrency(minimumAmountCurrency uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_currency = ?", minimumAmountCurrency).Find(&results).Error
	
	return
}

// GetBatchFromMinimumAmountCurrency 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromMinimumAmountCurrency(minimumAmountCurrencys []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_currency IN (?)", minimumAmountCurrencys).Find(&results).Error
	
	return
}
 
// GetFromMinimumAmountShipping 通过minimum_amount_shipping获取内容  
func (obj *_EgCartRuleMgr) GetFromMinimumAmountShipping(minimumAmountShipping bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_shipping = ?", minimumAmountShipping).Find(&results).Error
	
	return
}

// GetBatchFromMinimumAmountShipping 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromMinimumAmountShipping(minimumAmountShippings []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minimum_amount_shipping IN (?)", minimumAmountShippings).Find(&results).Error
	
	return
}
 
// GetFromCountryRestriction 通过country_restriction获取内容  
func (obj *_EgCartRuleMgr) GetFromCountryRestriction(countryRestriction bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_restriction = ?", countryRestriction).Find(&results).Error
	
	return
}

// GetBatchFromCountryRestriction 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromCountryRestriction(countryRestrictions []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("country_restriction IN (?)", countryRestrictions).Find(&results).Error
	
	return
}
 
// GetFromCarrierRestriction 通过carrier_restriction获取内容  
func (obj *_EgCartRuleMgr) GetFromCarrierRestriction(carrierRestriction bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carrier_restriction = ?", carrierRestriction).Find(&results).Error
	
	return
}

// GetBatchFromCarrierRestriction 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromCarrierRestriction(carrierRestrictions []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("carrier_restriction IN (?)", carrierRestrictions).Find(&results).Error
	
	return
}
 
// GetFromGroupRestriction 通过group_restriction获取内容  
func (obj *_EgCartRuleMgr) GetFromGroupRestriction(groupRestriction bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_restriction = ?", groupRestriction).Find(&results).Error
	
	return
}

// GetBatchFromGroupRestriction 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromGroupRestriction(groupRestrictions []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_restriction IN (?)", groupRestrictions).Find(&results).Error
	
	return
}
 
// GetFromCartRuleRestriction 通过cart_rule_restriction获取内容  
func (obj *_EgCartRuleMgr) GetFromCartRuleRestriction(cartRuleRestriction bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cart_rule_restriction = ?", cartRuleRestriction).Find(&results).Error
	
	return
}

// GetBatchFromCartRuleRestriction 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromCartRuleRestriction(cartRuleRestrictions []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cart_rule_restriction IN (?)", cartRuleRestrictions).Find(&results).Error
	
	return
}
 
// GetFromProductRestriction 通过product_restriction获取内容  
func (obj *_EgCartRuleMgr) GetFromProductRestriction(productRestriction bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_restriction = ?", productRestriction).Find(&results).Error
	
	return
}

// GetBatchFromProductRestriction 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromProductRestriction(productRestrictions []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_restriction IN (?)", productRestrictions).Find(&results).Error
	
	return
}
 
// GetFromShopRestriction 通过shop_restriction获取内容  
func (obj *_EgCartRuleMgr) GetFromShopRestriction(shopRestriction bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop_restriction = ?", shopRestriction).Find(&results).Error
	
	return
}

// GetBatchFromShopRestriction 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromShopRestriction(shopRestrictions []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop_restriction IN (?)", shopRestrictions).Find(&results).Error
	
	return
}
 
// GetFromFreeShipping 通过free_shipping获取内容  
func (obj *_EgCartRuleMgr) GetFromFreeShipping(freeShipping bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping = ?", freeShipping).Find(&results).Error
	
	return
}

// GetBatchFromFreeShipping 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromFreeShipping(freeShippings []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping IN (?)", freeShippings).Find(&results).Error
	
	return
}
 
// GetFromReductionPercent 通过reduction_percent获取内容  
func (obj *_EgCartRuleMgr) GetFromReductionPercent(reductionPercent float64) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_percent = ?", reductionPercent).Find(&results).Error
	
	return
}

// GetBatchFromReductionPercent 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromReductionPercent(reductionPercents []float64) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_percent IN (?)", reductionPercents).Find(&results).Error
	
	return
}
 
// GetFromReductionAmount 通过reduction_amount获取内容  
func (obj *_EgCartRuleMgr) GetFromReductionAmount(reductionAmount float64) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount = ?", reductionAmount).Find(&results).Error
	
	return
}

// GetBatchFromReductionAmount 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromReductionAmount(reductionAmounts []float64) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_amount IN (?)", reductionAmounts).Find(&results).Error
	
	return
}
 
// GetFromReductionTax 通过reduction_tax获取内容  
func (obj *_EgCartRuleMgr) GetFromReductionTax(reductionTax bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax = ?", reductionTax).Find(&results).Error
	
	return
}

// GetBatchFromReductionTax 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromReductionTax(reductionTaxs []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax IN (?)", reductionTaxs).Find(&results).Error
	
	return
}
 
// GetFromReductionCurrency 通过reduction_currency获取内容  
func (obj *_EgCartRuleMgr) GetFromReductionCurrency(reductionCurrency uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_currency = ?", reductionCurrency).Find(&results).Error
	
	return
}

// GetBatchFromReductionCurrency 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromReductionCurrency(reductionCurrencys []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_currency IN (?)", reductionCurrencys).Find(&results).Error
	
	return
}
 
// GetFromReductionProduct 通过reduction_product获取内容  
func (obj *_EgCartRuleMgr) GetFromReductionProduct(reductionProduct int) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_product = ?", reductionProduct).Find(&results).Error
	
	return
}

// GetBatchFromReductionProduct 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromReductionProduct(reductionProducts []int) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_product IN (?)", reductionProducts).Find(&results).Error
	
	return
}
 
// GetFromReductionExcludeSpecial 通过reduction_exclude_special获取内容  
func (obj *_EgCartRuleMgr) GetFromReductionExcludeSpecial(reductionExcludeSpecial bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_exclude_special = ?", reductionExcludeSpecial).Find(&results).Error
	
	return
}

// GetBatchFromReductionExcludeSpecial 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromReductionExcludeSpecial(reductionExcludeSpecials []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_exclude_special IN (?)", reductionExcludeSpecials).Find(&results).Error
	
	return
}
 
// GetFromGiftProduct 通过gift_product获取内容  
func (obj *_EgCartRuleMgr) GetFromGiftProduct(giftProduct uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_product = ?", giftProduct).Find(&results).Error
	
	return
}

// GetBatchFromGiftProduct 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromGiftProduct(giftProducts []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_product IN (?)", giftProducts).Find(&results).Error
	
	return
}
 
// GetFromGiftProductAttribute 通过gift_product_attribute获取内容  
func (obj *_EgCartRuleMgr) GetFromGiftProductAttribute(giftProductAttribute uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_product_attribute = ?", giftProductAttribute).Find(&results).Error
	
	return
}

// GetBatchFromGiftProductAttribute 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromGiftProductAttribute(giftProductAttributes []uint32) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gift_product_attribute IN (?)", giftProductAttributes).Find(&results).Error
	
	return
}
 
// GetFromHighlight 通过highlight获取内容  
func (obj *_EgCartRuleMgr) GetFromHighlight(highlight bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("highlight = ?", highlight).Find(&results).Error
	
	return
}

// GetBatchFromHighlight 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromHighlight(highlights []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("highlight IN (?)", highlights).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgCartRuleMgr) GetFromActive(active bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromActive(actives []bool) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgCartRuleMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgCartRuleMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgCartRuleMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCartRuleMgr) FetchByPrimaryKey(idCartRule uint32 ) (result EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCustomer  获取多个内容
 func (obj *_EgCartRuleMgr) FetchIndexByIDCustomer(idCustomer uint32 ,dateTo time.Time ,active bool ) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND date_to = ? AND active = ?", idCustomer , dateTo , active).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCustomer2  获取多个内容
 func (obj *_EgCartRuleMgr) FetchIndexByIDCustomer2(idCustomer uint32 ,dateTo time.Time ,highlight bool ,active bool ) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND date_to = ? AND highlight = ? AND active = ?", idCustomer , dateTo , highlight , active).Find(&results).Error
	
	return
}
 
 // FetchIndexByDateFrom  获取多个内容
 func (obj *_EgCartRuleMgr) FetchIndexByDateFrom(dateFrom time.Time ) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_from = ?", dateFrom).Find(&results).Error
	
	return
}
 
 // FetchIndexByGroupRestriction  获取多个内容
 func (obj *_EgCartRuleMgr) FetchIndexByGroupRestriction(dateTo time.Time ,groupRestriction bool ,active bool ) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to = ? AND group_restriction = ? AND active = ?", dateTo , groupRestriction , active).Find(&results).Error
	
	return
}
 
 // FetchIndexByGroupRestriction2  获取多个内容
 func (obj *_EgCartRuleMgr) FetchIndexByGroupRestriction2(dateTo time.Time ,groupRestriction bool ,highlight bool ,active bool ) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to = ? AND group_restriction = ? AND highlight = ? AND active = ?", dateTo , groupRestriction , highlight , active).Find(&results).Error
	
	return
}
 
 // FetchIndexByDateTo  获取多个内容
 func (obj *_EgCartRuleMgr) FetchIndexByDateTo(dateTo time.Time ) (results []*EgCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_to = ?", dateTo).Find(&results).Error
	
	return
}
 

	

