package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BadgeMgr struct {
	*_BaseMgr
}

func BadgeMgr(db *gorm.DB) *_BadgeMgr {
	if db == nil {
		panic(fmt.Errorf("BadgeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BadgeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_badge"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_BadgeMgr) GetTableName() string {
	return "ps_badge"
}

func (obj *_BadgeMgr) Get() (result Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_BadgeMgr) Gets() (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_BadgeMgr) WithIDBadge(idBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_badge"] = idBadge })
}

func (obj *_BadgeMgr) WithIDPsBadge(idPsBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_ps_badge"] = idPsBadge })
}

func (obj *_BadgeMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

func (obj *_BadgeMgr) WithIDGroup(idGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_BadgeMgr) WithGroupPosition(groupPosition int) Option {
	return optionFunc(func(o *options) { o.query["group_position"] = groupPosition })
}

func (obj *_BadgeMgr) WithScoring(scoring int) Option {
	return optionFunc(func(o *options) { o.query["scoring"] = scoring })
}

func (obj *_BadgeMgr) WithAwb(awb int) Option {
	return optionFunc(func(o *options) { o.query["awb"] = awb })
}

func (obj *_BadgeMgr) WithValidated(validated bool) Option {
	return optionFunc(func(o *options) { o.query["validated"] = validated })
}

func (obj *_BadgeMgr) GetByOption(opts ...Option) (result Badge, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_BadgeMgr) GetByOptions(opts ...Option) (results []*Badge, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetFromIDBadge(idBadge int) (result Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&result).Error

	return
}

func (obj *_BadgeMgr) GetBatchFromIDBadge(idBadges []int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge IN (?)", idBadges).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetFromIDPsBadge(idPsBadge int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_badge = ?", idPsBadge).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetBatchFromIDPsBadge(idPsBadges []int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_badge IN (?)", idPsBadges).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetFromType(_type string) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetBatchFromType(_types []string) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetFromIDGroup(idGroup int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetBatchFromIDGroup(idGroups []int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetFromGroupPosition(groupPosition int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_position = ?", groupPosition).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetBatchFromGroupPosition(groupPositions []int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_position IN (?)", groupPositions).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetFromScoring(scoring int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("scoring = ?", scoring).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetBatchFromScoring(scorings []int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("scoring IN (?)", scorings).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetFromAwb(awb int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("awb = ?", awb).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetBatchFromAwb(awbs []int) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("awb IN (?)", awbs).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetFromValidated(validated bool) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated = ?", validated).Find(&results).Error

	return
}

func (obj *_BadgeMgr) GetBatchFromValidated(validateds []bool) (results []*Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated IN (?)", validateds).Find(&results).Error

	return
}

func (obj *_BadgeMgr) FetchByPrimaryKey(idBadge int) (result Badge, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&result).Error

	return
}
