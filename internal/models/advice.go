package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AdviceMgr struct {
	*_BaseMgr
}

// AdviceMgr open func
func AdviceMgr(db *gorm.DB) *_AdviceMgr {
	if db == nil {
		panic(fmt.Errorf("AdviceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdviceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_advice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AdviceMgr) GetTableName() string {
	return "ps_advice"
}

// Get 获取
func (obj *_AdviceMgr) Get() (result Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AdviceMgr) Gets() (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAdvice id_advice获取
func (obj *_AdviceMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}

// WithIDPsAdvice id_ps_advice获取
func (obj *_AdviceMgr) WithIDPsAdvice(idPsAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_ps_advice"] = idPsAdvice })
}

// WithIDTab id_tab获取
func (obj *_AdviceMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

// WithIDsTab ids_tab获取
func (obj *_AdviceMgr) WithIDsTab(idsTab string) Option {
	return optionFunc(func(o *options) { o.query["ids_tab"] = idsTab })
}

// WithValidated validated获取
func (obj *_AdviceMgr) WithValidated(validated bool) Option {
	return optionFunc(func(o *options) { o.query["validated"] = validated })
}

// WithHide hide获取
func (obj *_AdviceMgr) WithHide(hide bool) Option {
	return optionFunc(func(o *options) { o.query["hide"] = hide })
}

// WithLocation location获取
func (obj *_AdviceMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithSelector selector获取
func (obj *_AdviceMgr) WithSelector(selector string) Option {
	return optionFunc(func(o *options) { o.query["selector"] = selector })
}

// WithStartDay start_day获取
func (obj *_AdviceMgr) WithStartDay(startDay int) Option {
	return optionFunc(func(o *options) { o.query["start_day"] = startDay })
}

// WithStopDay stop_day获取
func (obj *_AdviceMgr) WithStopDay(stopDay int) Option {
	return optionFunc(func(o *options) { o.query["stop_day"] = stopDay })
}

// WithWeight weight获取
func (obj *_AdviceMgr) WithWeight(weight int) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDAdvice 通过id_advice获取内容
func (obj *_AdviceMgr) GetFromIDAdvice(idAdvice int) (result Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&result).Error

	return
}

// GetBatchFromIDAdvice 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error

	return
}

// GetFromIDPsAdvice 通过id_ps_advice获取内容
func (obj *_AdviceMgr) GetFromIDPsAdvice(idPsAdvice int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_advice = ?", idPsAdvice).Find(&results).Error

	return
}

// GetBatchFromIDPsAdvice 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromIDPsAdvice(idPsAdvices []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_advice IN (?)", idPsAdvices).Find(&results).Error

	return
}

// GetFromIDTab 通过id_tab获取内容
func (obj *_AdviceMgr) GetFromIDTab(idTab int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

// GetBatchFromIDTab 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromIDTab(idTabs []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

// GetFromIDsTab 通过ids_tab获取内容
func (obj *_AdviceMgr) GetFromIDsTab(idsTab string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ids_tab = ?", idsTab).Find(&results).Error

	return
}

// GetBatchFromIDsTab 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromIDsTab(idsTabs []string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ids_tab IN (?)", idsTabs).Find(&results).Error

	return
}

// GetFromValidated 通过validated获取内容
func (obj *_AdviceMgr) GetFromValidated(validated bool) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated = ?", validated).Find(&results).Error

	return
}

// GetBatchFromValidated 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromValidated(validateds []bool) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated IN (?)", validateds).Find(&results).Error

	return
}

// GetFromHide 通过hide获取内容
func (obj *_AdviceMgr) GetFromHide(hide bool) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide = ?", hide).Find(&results).Error

	return
}

// GetBatchFromHide 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromHide(hides []bool) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide IN (?)", hides).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容
func (obj *_AdviceMgr) GetFromLocation(location string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromLocation(locations []string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

// GetFromSelector 通过selector获取内容
func (obj *_AdviceMgr) GetFromSelector(selector string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("selector = ?", selector).Find(&results).Error

	return
}

// GetBatchFromSelector 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromSelector(selectors []string) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("selector IN (?)", selectors).Find(&results).Error

	return
}

// GetFromStartDay 通过start_day获取内容
func (obj *_AdviceMgr) GetFromStartDay(startDay int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_day = ?", startDay).Find(&results).Error

	return
}

// GetBatchFromStartDay 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromStartDay(startDays []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_day IN (?)", startDays).Find(&results).Error

	return
}

// GetFromStopDay 通过stop_day获取内容
func (obj *_AdviceMgr) GetFromStopDay(stopDay int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stop_day = ?", stopDay).Find(&results).Error

	return
}

// GetBatchFromStopDay 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromStopDay(stopDays []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stop_day IN (?)", stopDays).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容
func (obj *_AdviceMgr) GetFromWeight(weight int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量唯一主键查找
func (obj *_AdviceMgr) GetBatchFromWeight(weights []int) (results []*Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AdviceMgr) FetchByPrimaryKey(idAdvice int) (result Advice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&result).Error

	return
}
