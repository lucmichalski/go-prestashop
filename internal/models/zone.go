package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ZoneMgr struct {
	*_BaseMgr
}

func ZoneMgr(db *gorm.DB) *_ZoneMgr {
	if db == nil {
		panic(fmt.Errorf("ZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_zone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ZoneMgr) GetTableName() string {
	return "ps_zone"
}

func (obj *_ZoneMgr) Get() (result Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ZoneMgr) Gets() (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ZoneMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

func (obj *_ZoneMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ZoneMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ZoneMgr) GetByOption(opts ...Option) (result Zone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ZoneMgr) GetByOptions(opts ...Option) (results []*Zone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ZoneMgr) GetFromIDZone(idZone uint32) (result Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&result).Error

	return
}

func (obj *_ZoneMgr) GetBatchFromIDZone(idZones []uint32) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error

	return
}

func (obj *_ZoneMgr) GetFromName(name string) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ZoneMgr) GetBatchFromName(names []string) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ZoneMgr) GetFromActive(active bool) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ZoneMgr) GetBatchFromActive(actives []bool) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}


func (obj *_ZoneMgr) FetchByPrimaryKey(idZone uint32) (result Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&result).Error

	return
}
