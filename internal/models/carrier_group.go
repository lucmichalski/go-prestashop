package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CarrierGroupMgr struct {
	*_BaseMgr
}

func CarrierGroupMgr(db *gorm.DB) *_CarrierGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CarrierGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CarrierGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_carrier_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CarrierGroupMgr) GetTableName() string {
	return "ps_carrier_group"
}

func (obj *_CarrierGroupMgr) Get() (result CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CarrierGroupMgr) Gets() (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CarrierGroupMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

func (obj *_CarrierGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_CarrierGroupMgr) GetByOption(opts ...Option) (result CarrierGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CarrierGroupMgr) GetByOptions(opts ...Option) (results []*CarrierGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CarrierGroupMgr) GetFromIDCarrier(idCarrier uint32) (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

func (obj *_CarrierGroupMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

func (obj *_CarrierGroupMgr) GetFromIDGroup(idGroup uint32) (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_CarrierGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}


func (obj *_CarrierGroupMgr) FetchByPrimaryKey(idCarrier uint32, idGroup uint32) (result CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_group = ?", idCarrier, idGroup).Find(&result).Error

	return
}
