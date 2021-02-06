package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SpecificPriceRuleMgr struct {
	*_BaseMgr
}

func SpecificPriceRuleMgr(db *gorm.DB) *_SpecificPriceRuleMgr {
	if db == nil {
		panic(fmt.Errorf("SpecificPriceRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SpecificPriceRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_specific_price_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SpecificPriceRuleMgr) GetTableName() string {
	return "ps_specific_price_rule"
}

func (obj *_SpecificPriceRuleMgr) Get() (result SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleMgr) Gets() (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SpecificPriceRuleMgr) WithIDSpecificPriceRule(idSpecificPriceRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule"] = idSpecificPriceRule })
}

func (obj *_SpecificPriceRuleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_SpecificPriceRuleMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_SpecificPriceRuleMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

func (obj *_SpecificPriceRuleMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_SpecificPriceRuleMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_SpecificPriceRuleMgr) WithFromQuantity(fromQuantity string) Option {
	return optionFunc(func(o *options) { o.query["from_quantity"] = fromQuantity })
}

func (obj *_SpecificPriceRuleMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

func (obj *_SpecificPriceRuleMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}

func (obj *_SpecificPriceRuleMgr) WithReductionTax(reductionTax bool) Option {
	return optionFunc(func(o *options) { o.query["reduction_tax"] = reductionTax })
}

func (obj *_SpecificPriceRuleMgr) WithReductionType(reductionType string) Option {
	return optionFunc(func(o *options) { o.query["reduction_type"] = reductionType })
}

func (obj *_SpecificPriceRuleMgr) WithFrom(from time.Time) Option {
	return optionFunc(func(o *options) { o.query["from"] = from })
}

func (obj *_SpecificPriceRuleMgr) WithTo(to time.Time) Option {
	return optionFunc(func(o *options) { o.query["to"] = to })
}

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


func (obj *_SpecificPriceRuleMgr) GetFromIDSpecificPriceRule(idSpecificPriceRule uint32) (result SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromIDSpecificPriceRule(idSpecificPriceRules []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule IN (?)", idSpecificPriceRules).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromName(name string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromName(names []string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromIDShop(idShop uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromIDShop(idShops []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromIDCurrency(idCurrency uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromIDCountry(idCountry uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromIDGroup(idGroup uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromFromQuantity(fromQuantity string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_quantity = ?", fromQuantity).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromFromQuantity(fromQuantitys []string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_quantity IN (?)", fromQuantitys).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromPrice(price float64) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromPrice(prices []float64) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromReduction(reduction float64) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromReduction(reductions []float64) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromReductionTax(reductionTax bool) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax = ?", reductionTax).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromReductionTax(reductionTaxs []bool) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_tax IN (?)", reductionTaxs).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromReductionType(reductionType string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_type = ?", reductionType).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromReductionType(reductionTypes []string) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction_type IN (?)", reductionTypes).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromFrom(from time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from = ?", from).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromFrom(froms []time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from IN (?)", froms).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetFromTo(to time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to = ?", to).Find(&results).Error

	return
}

func (obj *_SpecificPriceRuleMgr) GetBatchFromTo(tos []time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to IN (?)", tos).Find(&results).Error

	return
}


func (obj *_SpecificPriceRuleMgr) FetchByPrimaryKey(idSpecificPriceRule uint32) (result SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&result).Error

	return
}

func (obj *_SpecificPriceRuleMgr) FetchIndexByIDProduct(idShop uint32, idCurrency uint32, idCountry uint32, idGroup uint32, fromQuantity string, from time.Time, to time.Time) (results []*SpecificPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_currency = ? AND id_country = ? AND id_group = ? AND from_quantity = ? AND from = ? AND to = ?", idShop, idCurrency, idCountry, idGroup, fromQuantity, from, to).Find(&results).Error

	return
}
