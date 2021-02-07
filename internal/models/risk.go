package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _RiskMgr struct {
	*_BaseMgr
}

func RiskMgr(db *gorm.DB) *_RiskMgr {
	if db == nil {
		panic(fmt.Errorf("RiskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RiskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_risk"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_RiskMgr) GetTableName() string {
	return "ps_risk"
}

func (obj *_RiskMgr) Get() (result Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_RiskMgr) Gets() (results []*Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_RiskMgr) WithIDRisk(idRisk uint32) Option {
	return optionFunc(func(o *options) { o.query["id_risk"] = idRisk })
}

func (obj *_RiskMgr) WithPercent(percent int8) Option {
	return optionFunc(func(o *options) { o.query["percent"] = percent })
}

func (obj *_RiskMgr) WithColor(color string) Option {
	return optionFunc(func(o *options) { o.query["color"] = color })
}

func (obj *_RiskMgr) GetByOption(opts ...Option) (result Risk, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_RiskMgr) GetByOptions(opts ...Option) (results []*Risk, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_RiskMgr) GetFromIDRisk(idRisk uint32) (result Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&result).Error

	return
}

func (obj *_RiskMgr) GetBatchFromIDRisk(idRisks []uint32) (results []*Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk IN (?)", idRisks).Find(&results).Error

	return
}

func (obj *_RiskMgr) GetFromPercent(percent int8) (results []*Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("percent = ?", percent).Find(&results).Error

	return
}

func (obj *_RiskMgr) GetBatchFromPercent(percents []int8) (results []*Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("percent IN (?)", percents).Find(&results).Error

	return
}

func (obj *_RiskMgr) GetFromColor(color string) (results []*Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color = ?", color).Find(&results).Error

	return
}

func (obj *_RiskMgr) GetBatchFromColor(colors []string) (results []*Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("color IN (?)", colors).Find(&results).Error

	return
}

func (obj *_RiskMgr) FetchByPrimaryKey(idRisk uint32) (result Risk, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&result).Error

	return
}
