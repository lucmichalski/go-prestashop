package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ConditionMgr struct {
	*_BaseMgr
}

func ConditionMgr(db *gorm.DB) *_ConditionMgr {
	if db == nil {
		panic(fmt.Errorf("ConditionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConditionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_condition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ConditionMgr) GetTableName() string {
	return "ps_condition"
}

func (obj *_ConditionMgr) Get() (result Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ConditionMgr) Gets() (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ConditionMgr) WithIDCondition(idCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_condition"] = idCondition })
}

func (obj *_ConditionMgr) WithIDPsCondition(idPsCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_ps_condition"] = idPsCondition })
}

func (obj *_ConditionMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

func (obj *_ConditionMgr) WithRequest(request string) Option {
	return optionFunc(func(o *options) { o.query["request"] = request })
}

func (obj *_ConditionMgr) WithOperator(operator string) Option {
	return optionFunc(func(o *options) { o.query["operator"] = operator })
}

func (obj *_ConditionMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

func (obj *_ConditionMgr) WithResult(result string) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

func (obj *_ConditionMgr) WithCalculationType(calculationType string) Option {
	return optionFunc(func(o *options) { o.query["calculation_type"] = calculationType })
}

func (obj *_ConditionMgr) WithCalculationDetail(calculationDetail string) Option {
	return optionFunc(func(o *options) { o.query["calculation_detail"] = calculationDetail })
}

func (obj *_ConditionMgr) WithValidated(validated bool) Option {
	return optionFunc(func(o *options) { o.query["validated"] = validated })
}

func (obj *_ConditionMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ConditionMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_ConditionMgr) GetByOption(opts ...Option) (result Condition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ConditionMgr) GetByOptions(opts ...Option) (results []*Condition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ConditionMgr) GetFromIDCondition(idCondition int) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ?", idCondition).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromIDCondition(idConditions []int) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition IN (?)", idConditions).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromIDPsCondition(idPsCondition int) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_condition = ?", idPsCondition).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromIDPsCondition(idPsConditions []int) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_condition IN (?)", idPsConditions).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromType(_type string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromType(_types []string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromRequest(request string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request = ?", request).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromRequest(requests []string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request IN (?)", requests).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromOperator(operator string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("operator = ?", operator).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromOperator(operators []string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("operator IN (?)", operators).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromValue(value string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromValue(values []string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromResult(result string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("result = ?", result).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromResult(results []string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("result IN (?)", results).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromCalculationType(calculationType string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calculation_type = ?", calculationType).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromCalculationType(calculationTypes []string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calculation_type IN (?)", calculationTypes).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromCalculationDetail(calculationDetail string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calculation_detail = ?", calculationDetail).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromCalculationDetail(calculationDetails []string) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calculation_detail IN (?)", calculationDetails).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromValidated(validated bool) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated = ?", validated).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromValidated(validateds []bool) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated IN (?)", validateds).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromDateAdd(dateAdd time.Time) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetFromDateUpd(dateUpd time.Time) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_ConditionMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_ConditionMgr) FetchByPrimaryKey(idCondition int, idPsCondition int) (result Condition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ? AND id_ps_condition = ?", idCondition, idPsCondition).Find(&result).Error

	return
}
