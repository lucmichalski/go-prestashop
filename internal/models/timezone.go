package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TimezoneMgr struct {
	*_BaseMgr
}

func TimezoneMgr(db *gorm.DB) *_TimezoneMgr {
	if db == nil {
		panic(fmt.Errorf("TimezoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TimezoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_timezone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_TimezoneMgr) GetTableName() string {
	return "ps_timezone"
}

func (obj *_TimezoneMgr) Get() (result Timezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_TimezoneMgr) Gets() (results []*Timezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_TimezoneMgr) WithIDTimezone(idTimezone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_timezone"] = idTimezone })
}

func (obj *_TimezoneMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_TimezoneMgr) GetByOption(opts ...Option) (result Timezone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_TimezoneMgr) GetByOptions(opts ...Option) (results []*Timezone, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_TimezoneMgr) GetFromIDTimezone(idTimezone uint32) (result Timezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_timezone = ?", idTimezone).Find(&result).Error

	return
}

func (obj *_TimezoneMgr) GetBatchFromIDTimezone(idTimezones []uint32) (results []*Timezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_timezone IN (?)", idTimezones).Find(&results).Error

	return
}

func (obj *_TimezoneMgr) GetFromName(name string) (results []*Timezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_TimezoneMgr) GetBatchFromName(names []string) (results []*Timezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_TimezoneMgr) FetchByPrimaryKey(idTimezone uint32) (result Timezone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_timezone = ?", idTimezone).Find(&result).Error

	return
}
