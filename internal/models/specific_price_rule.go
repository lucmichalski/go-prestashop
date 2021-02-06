package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SpecificPriceRuleMgr struct {
	*_BaseMgr
}

// SpecificPriceRuleMgr open func
func SpecificPriceRuleMgr(db *gorm.DB) *_SpecificPriceRuleMgr {
	if db == nil {
		panic(fmt.Errorf("SpecificPriceRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpecificPriceRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_specific_price_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SpecificPriceRuleMgr) GetTableName() string {
	return "ps_specific_price_rule"
}

// Get 获取
func (obj *_SpecificPriceRuleMgr) Get() (result SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SpecificPriceRuleMgr) Gets() (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSpecificPriceRule id_specific_price_rule获取
func (obj *_SpecificPriceRuleMgr) WithIDSpecificPriceRule(idSpecificPriceRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule"] = idSpecificPriceRule })
}

// WithName name获取
func (obj *_SpecificPriceRuleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIDShop id_shop获取
func (obj *_SpecificPriceRuleMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDCurrency id_currency获取
func (obj *_SpecificPriceRuleMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDCountry id_country获取
func (obj *_SpecificPriceRuleMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDGroup id_group获取
func (obj *_SpecificPriceRuleMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithFromQuantity from_quantity获取
func (obj *_SpecificPriceRuleMgr) WithFromQuantity(fromQuantity string) Option {
	return optionFunc(func(o *options) { o.query["from_quantity"] = fromQuantity })
}

// WithPrice price获取
func (obj *_SpecificPriceRuleMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithReduction reduction获取
func (obj *_SpecificPriceRuleMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}

// WithReductionTax reduction_tax获取
func (obj *_SpecificPriceRuleMgr) WithReductionTax(reductionTax bool) Option {
	return optionFunc(func(o *options) { o.query["reduction_tax"] = reductionTax })
}

// WithReductionType reduction_type获取
func (obj *_SpecificPriceRuleMgr) WithReductionType(reductionType string) Option {
	return optionFunc(func(o *options) { o.query["reduction_type"] = reductionType })
}

// WithFrom from获取
func (obj *_SpecificPriceRuleMgr) WithFrom(from time.Time) Option {
	return optionFunc(func(o *options) { o.query["from"] = from })
}

// WithTo to获取
func (obj *_SpecificPriceRuleMgr) WithTo(to time.Time) Option {
	return optionFunc(func(o *options) { o.query["to"] = to })
}

// GetByOption 功能选项模式获取
func (obj *_SpecificPriceRuleMgr) GetByOption(opts ...Option) (result SpecificPriceRule, err error) {
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
func (obj *_SpecificPriceRuleMgr) GetByOptions(opts ...Option) (results []*SpecificPriceRule, err error) {
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

// GetFromIDSpecificPriceRule 通过id_specific_price_rule获取内容
func (obj *_SpecificPriceRuleMgr) GetFromIDSpecificPriceRule(idSpecificPriceRule uint32) (result SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&result).Error

	return
}

// GetBatchFromIDSpecificPriceRule 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromIDSpecificPriceRule(idSpecificPriceRules []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule IN (?)", idSpecificPriceRules).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_SpecificPriceRuleMgr) GetFromName(name string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromName(names []string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_SpecificPriceRuleMgr) GetFromIDShop(idShop uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromIDShop(idShops []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDCurrency 通过id_currency获取内容
func (obj *_SpecificPriceRuleMgr) GetFromIDCurrency(idCurrency uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

// GetBatchFromIDCurrency 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

// GetFromIDCountry 通过id_country获取内容
func (obj *_SpecificPriceRuleMgr) GetFromIDCountry(idCountry uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

// GetBatchFromIDCountry 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

// GetFromIDGroup 通过id_group获取内容
func (obj *_SpecificPriceRuleMgr) GetFromIDGroup(idGroup uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

// GetFromFromQuantity 通过from_quantity获取内容
func (obj *_SpecificPriceRuleMgr) GetFromFromQuantity(fromQuantity string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_quantity = ?", fromQuantity).Find(&results).Error

	return
}

// GetBatchFromFromQuantity 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromFromQuantity(fromQuantitys []string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_quantity IN (?)", fromQuantitys).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容
func (obj *_SpecificPriceRuleMgr) GetFromPrice(price float64) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromPrice(prices []float64) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromReduction 通过reduction获取内容
func (obj *_SpecificPriceRuleMgr) GetFromReduction(reduction float64) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error

	return
}

// GetBatchFromReduction 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromReduction(reductions []float64) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error

	return
}

// GetFromReductionTax 通过reduction_tax获取内容
func (obj *_SpecificPriceRuleMgr) GetFromReductionTax(reductionTax bool) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax = ?", reductionTax).Find(&results).Error

	return
}

// GetBatchFromReductionTax 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromReductionTax(reductionTaxs []bool) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax IN (?)", reductionTaxs).Find(&results).Error

	return
}

// GetFromReductionType 通过reduction_type获取内容
func (obj *_SpecificPriceRuleMgr) GetFromReductionType(reductionType string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_type = ?", reductionType).Find(&results).Error

	return
}

// GetBatchFromReductionType 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromReductionType(reductionTypes []string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_type IN (?)", reductionTypes).Find(&results).Error

	return
}

// GetFromFrom 通过from获取内容
func (obj *_SpecificPriceRuleMgr) GetFromFrom(from time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from = ?", from).Find(&results).Error

	return
}

// GetBatchFromFrom 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromFrom(froms []time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from IN (?)", froms).Find(&results).Error

	return
}

// GetFromTo 通过to获取内容
func (obj *_SpecificPriceRuleMgr) GetFromTo(to time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to = ?", to).Find(&results).Error

	return
}

// GetBatchFromTo 批量唯一主键查找
func (obj *_SpecificPriceRuleMgr) GetBatchFromTo(tos []time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to IN (?)", tos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SpecificPriceRuleMgr) FetchByPrimaryKey(idSpecificPriceRule uint32) (result SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&result).Error

	return
}

// FetchIndexByIDProduct  获取多个内容
func (obj *_SpecificPriceRuleMgr) FetchIndexByIDProduct(idShop uint32, idCurrency uint32, idCountry uint32, idGroup uint32, fromQuantity string, from time.Time, to time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_currency = ? AND id_country = ? AND id_group = ? AND from_quantity = ? AND from = ? AND to = ?", idShop, idCurrency, idCountry, idGroup, fromQuantity, from, to).Find(&results).Error

	return
}
