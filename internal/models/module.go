package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleMgr struct {
	*_BaseMgr
}

func ModuleMgr(db *gorm.DB) *_ModuleMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ModuleMgr) GetTableName() string {
	return "ps_module"
}

func (obj *_ModuleMgr) Get() (result Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ModuleMgr) Gets() (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ModuleMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_ModuleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ModuleMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ModuleMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

func (obj *_ModuleMgr) GetByOption(opts ...Option) (result Module, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ModuleMgr) GetByOptions(opts ...Option) (results []*Module, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ModuleMgr) GetFromIDModule(idModule uint32) (result Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&result).Error

	return
}

func (obj *_ModuleMgr) GetBatchFromIDModule(idModules []uint32) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_ModuleMgr) GetFromName(name string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ModuleMgr) GetBatchFromName(names []string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ModuleMgr) GetFromActive(active bool) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ModuleMgr) GetBatchFromActive(actives []bool) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_ModuleMgr) GetFromVersion(version string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("version = ?", version).Find(&results).Error

	return
}

func (obj *_ModuleMgr) GetBatchFromVersion(versions []string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("version IN (?)", versions).Find(&results).Error

	return
}


func (obj *_ModuleMgr) FetchByPrimaryKey(idModule uint32) (result Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&result).Error

	return
}

func (obj *_ModuleMgr) FetchUniqueByNameUNIQUE(name string) (result Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}

func (obj *_ModuleMgr) FetchIndexByName(name string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
