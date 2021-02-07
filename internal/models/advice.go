package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AdviceMgr struct {
	*_BaseMgr
}

func AdviceMgr(db *gorm.DB) *_AdviceMgr {
	if db == nil {
		panic(fmt.Errorf("AdviceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdviceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_advice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AdviceMgr) GetTableName() string {
	return "ps_advice"
}

func (obj *_AdviceMgr) Get() (result Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AdviceMgr) Gets() (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_AdviceMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}

func (obj *_AdviceMgr) WithIDPsAdvice(idPsAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_ps_advice"] = idPsAdvice })
}

func (obj *_AdviceMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

func (obj *_AdviceMgr) WithIDsTab(idsTab string) Option {
	return optionFunc(func(o *options) { o.query["ids_tab"] = idsTab })
}

func (obj *_AdviceMgr) WithValidated(validated bool) Option {
	return optionFunc(func(o *options) { o.query["validated"] = validated })
}

func (obj *_AdviceMgr) WithHide(hide bool) Option {
	return optionFunc(func(o *options) { o.query["hide"] = hide })
}

func (obj *_AdviceMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

func (obj *_AdviceMgr) WithSelector(selector string) Option {
	return optionFunc(func(o *options) { o.query["selector"] = selector })
}

func (obj *_AdviceMgr) WithStartDay(startDay int) Option {
	return optionFunc(func(o *options) { o.query["start_day"] = startDay })
}

func (obj *_AdviceMgr) WithStopDay(stopDay int) Option {
	return optionFunc(func(o *options) { o.query["stop_day"] = stopDay })
}

func (obj *_AdviceMgr) WithWeight(weight int) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

func (obj *_AdviceMgr) GetByOption(opts ...Option) (result Advice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AdviceMgr) GetByOptions(opts ...Option) (results []*Advice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromIDAdvice(idAdvice int) (result Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&result).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromIDPsAdvice(idPsAdvice int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_advice = ?", idPsAdvice).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromIDPsAdvice(idPsAdvices []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_advice IN (?)", idPsAdvices).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromIDTab(idTab int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromIDTab(idTabs []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromIDsTab(idsTab string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ids_tab = ?", idsTab).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromIDsTab(idsTabs []string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ids_tab IN (?)", idsTabs).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromValidated(validated bool) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated = ?", validated).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromValidated(validateds []bool) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated IN (?)", validateds).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromHide(hide bool) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide = ?", hide).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromHide(hides []bool) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide IN (?)", hides).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromLocation(location string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromLocation(locations []string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromSelector(selector string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("selector = ?", selector).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromSelector(selectors []string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("selector IN (?)", selectors).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromStartDay(startDay int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_day = ?", startDay).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromStartDay(startDays []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_day IN (?)", startDays).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromStopDay(stopDay int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stop_day = ?", stopDay).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromStopDay(stopDays []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stop_day IN (?)", stopDays).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetFromWeight(weight int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

func (obj *_AdviceMgr) GetBatchFromWeight(weights []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

func (obj *_AdviceMgr) FetchByPrimaryKey(idAdvice int) (result Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&result).Error

	return
}
