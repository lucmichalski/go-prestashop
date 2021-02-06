package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CarrierZoneMgr struct {
	*_BaseMgr
}

// CarrierZoneMgr open func
func CarrierZoneMgr(db *gorm.DB) *_CarrierZoneMgr {
	if db == nil {
		panic(fmt.Errorf("CarrierZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CarrierZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_carrier_zone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarrierZoneMgr) GetTableName() string {
	return "ps_carrier_zone"
}

// Get 获取
func (obj *_CarrierZoneMgr) Get() (result CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarrierZoneMgr) Gets() (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取
func (obj *_CarrierZoneMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDZone id_zone获取
func (obj *_CarrierZoneMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCarrier 通过id_carrier获取内容
func (obj *_CarrierZoneMgr) GetFromIDCarrier(idCarrier uint32) (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_CarrierZoneMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromIDZone 通过id_zone获取内容
func (obj *_CarrierZoneMgr) GetFromIDZone(idZone uint32) (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}

// GetBatchFromIDZone 批量唯一主键查找
func (obj *_CarrierZoneMgr) GetBatchFromIDZone(idZones []uint32) (results []*CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarrierZoneMgr) FetchByPrimaryKey(idCarrier uint32, idZone uint32) (result CarrierZone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_zone = ?", idCarrier, idZone).Find(&result).Error

	return
}
