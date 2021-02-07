package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OperatingSystemMgr struct {
	*_BaseMgr
}

func OperatingSystemMgr(db *gorm.DB) *_OperatingSystemMgr {
	if db == nil {
		panic(fmt.Errorf("OperatingSystemMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OperatingSystemMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_operating_system"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OperatingSystemMgr) GetTableName() string {
	return "ps_operating_system"
}

func (obj *_OperatingSystemMgr) Get() (result OperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OperatingSystemMgr) Gets() (results []*OperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_OperatingSystemMgr) WithIDOperatingSystem(idOperatingSystem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_operating_system"] = idOperatingSystem })
}

func (obj *_OperatingSystemMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_OperatingSystemMgr) GetByOption(opts ...Option) (result OperatingSystem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OperatingSystemMgr) GetByOptions(opts ...Option) (results []*OperatingSystem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_OperatingSystemMgr) GetFromIDOperatingSystem(idOperatingSystem uint32) (result OperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system = ?", idOperatingSystem).Find(&result).Error

	return
}

func (obj *_OperatingSystemMgr) GetBatchFromIDOperatingSystem(idOperatingSystems []uint32) (results []*OperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system IN (?)", idOperatingSystems).Find(&results).Error

	return
}

func (obj *_OperatingSystemMgr) GetFromName(name string) (results []*OperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_OperatingSystemMgr) GetBatchFromName(names []string) (results []*OperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_OperatingSystemMgr) FetchByPrimaryKey(idOperatingSystem uint32) (result OperatingSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_operating_system = ?", idOperatingSystem).Find(&result).Error

	return
}
