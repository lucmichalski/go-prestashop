package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CarrierZoneMgr struct {
	*_BaseMgr
}

func CarrierZoneMgr(db *gorm.DB) *_CarrierZoneMgr {
	if db == nil {
		panic(fmt.Errorf("CarrierZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CarrierZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_carrier_zone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CarrierZoneMgr) GetTableName() string {
	return "ps_carrier_zone"
}

func (obj *_CarrierZoneMgr) Get() (result CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CarrierZoneMgr) Gets() (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CarrierZoneMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

func (obj *_CarrierZoneMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

func (obj *_CarrierZoneMgr) GetByOption(opts ...Option) (result CarrierZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CarrierZoneMgr) GetByOptions(opts ...Option) (results []*CarrierZone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CarrierZoneMgr) GetFromIDCarrier(idCarrier uint32) (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

func (obj *_CarrierZoneMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

func (obj *_CarrierZoneMgr) GetFromIDZone(idZone uint32) (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}

func (obj *_CarrierZoneMgr) GetBatchFromIDZone(idZones []uint32) (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error

	return
}

func (obj *_CarrierZoneMgr) FetchByPrimaryKey(idCarrier uint32, idZone uint32) (result CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_zone = ?", idCarrier, idZone).Find(&result).Error

	return
}
