package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ManufacturerMgr struct {
	*_BaseMgr
}

func ManufacturerMgr(db *gorm.DB) *_ManufacturerMgr {
	if db == nil {
		panic(fmt.Errorf("ManufacturerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ManufacturerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_manufacturer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ManufacturerMgr) GetTableName() string {
	return "ps_manufacturer"
}

func (obj *_ManufacturerMgr) Get() (result Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ManufacturerMgr) Gets() (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

func (obj *_ManufacturerMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ManufacturerMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ManufacturerMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_ManufacturerMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ManufacturerMgr) GetByOption(opts ...Option) (result Manufacturer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ManufacturerMgr) GetByOptions(opts ...Option) (results []*Manufacturer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetFromIDManufacturer(idManufacturer uint32) (result Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&result).Error

	return
}

func (obj *_ManufacturerMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetFromName(name string) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetBatchFromName(names []string) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetFromDateAdd(dateAdd time.Time) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetFromDateUpd(dateUpd time.Time) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetFromActive(active bool) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) GetBatchFromActive(actives []bool) (results []*Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_ManufacturerMgr) FetchByPrimaryKey(idManufacturer uint32) (result Manufacturer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&result).Error

	return
}
