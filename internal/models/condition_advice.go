package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ConditionAdviceMgr struct {
	*_BaseMgr
}

func ConditionAdviceMgr(db *gorm.DB) *_ConditionAdviceMgr {
	if db == nil {
		panic(fmt.Errorf("ConditionAdviceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConditionAdviceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_condition_advice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ConditionAdviceMgr) GetTableName() string {
	return "ps_condition_advice"
}

func (obj *_ConditionAdviceMgr) Get() (result ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ConditionAdviceMgr) Gets() (results []*ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ConditionAdviceMgr) WithIDCondition(idCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_condition"] = idCondition })
}

func (obj *_ConditionAdviceMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}

func (obj *_ConditionAdviceMgr) WithDisplay(display bool) Option {
	return optionFunc(func(o *options) { o.query["display"] = display })
}

func (obj *_ConditionAdviceMgr) GetByOption(opts ...Option) (result ConditionAdvice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ConditionAdviceMgr) GetByOptions(opts ...Option) (results []*ConditionAdvice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ConditionAdviceMgr) GetFromIDCondition(idCondition int) (results []*ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ?", idCondition).Find(&results).Error

	return
}

func (obj *_ConditionAdviceMgr) GetBatchFromIDCondition(idConditions []int) (results []*ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition IN (?)", idConditions).Find(&results).Error

	return
}

func (obj *_ConditionAdviceMgr) GetFromIDAdvice(idAdvice int) (results []*ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&results).Error

	return
}

func (obj *_ConditionAdviceMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error

	return
}

func (obj *_ConditionAdviceMgr) GetFromDisplay(display bool) (results []*ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display = ?", display).Find(&results).Error

	return
}

func (obj *_ConditionAdviceMgr) GetBatchFromDisplay(displays []bool) (results []*ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display IN (?)", displays).Find(&results).Error

	return
}

func (obj *_ConditionAdviceMgr) FetchByPrimaryKey(idCondition int, idAdvice int) (result ConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ? AND id_advice = ?", idCondition, idAdvice).Find(&result).Error

	return
}
