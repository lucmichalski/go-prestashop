package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ConditionBadgeMgr struct {
	*_BaseMgr
}

// ConditionBadgeMgr open func
func ConditionBadgeMgr(db *gorm.DB) *_ConditionBadgeMgr {
	if db == nil {
		panic(fmt.Errorf("ConditionBadgeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ConditionBadgeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_condition_badge"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ConditionBadgeMgr) GetTableName() string {
	return "ps_condition_badge"
}

// Get 获取
func (obj *_ConditionBadgeMgr) Get() (result ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ConditionBadgeMgr) Gets() (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCondition id_condition获取
func (obj *_ConditionBadgeMgr) WithIDCondition(idCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_condition"] = idCondition })
}

// WithIDBadge id_badge获取
func (obj *_ConditionBadgeMgr) WithIDBadge(idBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_badge"] = idBadge })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCondition 通过id_condition获取内容
func (obj *_ConditionBadgeMgr) GetFromIDCondition(idCondition int) (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ?", idCondition).Find(&results).Error

	return
}

// GetBatchFromIDCondition 批量唯一主键查找
func (obj *_ConditionBadgeMgr) GetBatchFromIDCondition(idConditions []int) (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition IN (?)", idConditions).Find(&results).Error

	return
}

// GetFromIDBadge 通过id_badge获取内容
func (obj *_ConditionBadgeMgr) GetFromIDBadge(idBadge int) (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&results).Error

	return
}

// GetBatchFromIDBadge 批量唯一主键查找
func (obj *_ConditionBadgeMgr) GetBatchFromIDBadge(idBadges []int) (results []*ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge IN (?)", idBadges).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ConditionBadgeMgr) FetchByPrimaryKey(idCondition int, idBadge int) (result ConditionBadge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ? AND id_badge = ?", idCondition, idBadge).Find(&result).Error

	return
}
