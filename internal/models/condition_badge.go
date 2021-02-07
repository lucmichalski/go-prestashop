package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ConditionBadgeMgr struct {
	*_BaseMgr
}

func ConditionBadgeMgr(db *gorm.DB) *_ConditionBadgeMgr {
	if db == nil {
		panic(fmt.Errorf("ConditionBadgeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConditionBadgeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_condition_badge"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ConditionBadgeMgr) GetTableName() string {
	return "ps_condition_badge"
}

func (obj *_ConditionBadgeMgr) Get() (result ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ConditionBadgeMgr) Gets() (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ConditionBadgeMgr) WithIDCondition(idCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_condition"] = idCondition })
}

func (obj *_ConditionBadgeMgr) WithIDBadge(idBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_badge"] = idBadge })
}

func (obj *_ConditionBadgeMgr) GetByOption(opts ...Option) (result ConditionBadge, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ConditionBadgeMgr) GetByOptions(opts ...Option) (results []*ConditionBadge, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ConditionBadgeMgr) GetFromIDCondition(idCondition int) (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ?", idCondition).Find(&results).Error

	return
}

func (obj *_ConditionBadgeMgr) GetBatchFromIDCondition(idConditions []int) (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition IN (?)", idConditions).Find(&results).Error

	return
}

func (obj *_ConditionBadgeMgr) GetFromIDBadge(idBadge int) (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&results).Error

	return
}

func (obj *_ConditionBadgeMgr) GetBatchFromIDBadge(idBadges []int) (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge IN (?)", idBadges).Find(&results).Error

	return
}

func (obj *_ConditionBadgeMgr) FetchByPrimaryKey(idCondition int, idBadge int) (result ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ? AND id_badge = ?", idCondition, idBadge).Find(&result).Error

	return
}
