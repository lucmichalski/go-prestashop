package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TaxRuleMgr struct {
	*_BaseMgr
}

func TaxRuleMgr(db *gorm.DB) *_TaxRuleMgr {
	if db == nil {
		panic(fmt.Errorf("TaxRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TaxRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tax_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_TaxRuleMgr) GetTableName() string {
	return "ps_tax_rule"
}

func (obj *_TaxRuleMgr) Get() (result TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_TaxRuleMgr) Gets() (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) WithIDTaxRule(idTaxRule int) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rule"] = idTaxRule })
}

func (obj *_TaxRuleMgr) WithIDTaxRulesGroup(idTaxRulesGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

func (obj *_TaxRuleMgr) WithIDCountry(idCountry int) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_TaxRuleMgr) WithIDState(idState int) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

func (obj *_TaxRuleMgr) WithZipcodeFrom(zipcodeFrom string) Option {
	return optionFunc(func(o *options) { o.query["zipcode_from"] = zipcodeFrom })
}

func (obj *_TaxRuleMgr) WithZipcodeTo(zipcodeTo string) Option {
	return optionFunc(func(o *options) { o.query["zipcode_to"] = zipcodeTo })
}

func (obj *_TaxRuleMgr) WithIDTax(idTax int) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

func (obj *_TaxRuleMgr) WithBehavior(behavior int) Option {
	return optionFunc(func(o *options) { o.query["behavior"] = behavior })
}

func (obj *_TaxRuleMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_TaxRuleMgr) GetByOption(opts ...Option) (result TaxRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_TaxRuleMgr) GetByOptions(opts ...Option) (results []*TaxRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromIDTaxRule(idTaxRule int) (result TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rule = ?", idTaxRule).Find(&result).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromIDTaxRule(idTaxRules []int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rule IN (?)", idTaxRules).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromIDCountry(idCountry int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromIDCountry(idCountrys []int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromIDState(idState int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromIDState(idStates []int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromZipcodeFrom(zipcodeFrom string) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zipcode_from = ?", zipcodeFrom).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromZipcodeFrom(zipcodeFroms []string) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zipcode_from IN (?)", zipcodeFroms).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromZipcodeTo(zipcodeTo string) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zipcode_to = ?", zipcodeTo).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromZipcodeTo(zipcodeTos []string) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zipcode_to IN (?)", zipcodeTos).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromIDTax(idTax int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromIDTax(idTaxs []int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromBehavior(behavior int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("behavior = ?", behavior).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromBehavior(behaviors []int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("behavior IN (?)", behaviors).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetFromDescription(description string) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) GetBatchFromDescription(descriptions []string) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) FetchByPrimaryKey(idTaxRule int) (result TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rule = ?", idTaxRule).Find(&result).Error

	return
}

func (obj *_TaxRuleMgr) FetchIndexByIDTaxRulesGroup(idTaxRulesGroup int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) FetchIndexByCategoryGetproducts(idTaxRulesGroup int, idCountry int, idState int, zipcodeFrom string) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ? AND id_country = ? AND id_state = ? AND zipcode_from = ?", idTaxRulesGroup, idCountry, idState, zipcodeFrom).Find(&results).Error

	return
}

func (obj *_TaxRuleMgr) FetchIndexByIDTax(idTax int) (results []*TaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error

	return
}
