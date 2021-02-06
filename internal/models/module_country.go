package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleCountryMgr struct {
	*_BaseMgr
}

func ModuleCountryMgr(db *gorm.DB) *_ModuleCountryMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ModuleCountryMgr) GetTableName() string {
	return "ps_module_country"
}

func (obj *_ModuleCountryMgr) Get() (result ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ModuleCountryMgr) Gets() (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ModuleCountryMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_ModuleCountryMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ModuleCountryMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_ModuleCountryMgr) GetByOption(opts ...Option) (result ModuleCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ModuleCountryMgr) GetByOptions(opts ...Option) (results []*ModuleCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ModuleCountryMgr) GetFromIDModule(idModule uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_ModuleCountryMgr) GetBatchFromIDModule(idModules []uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_ModuleCountryMgr) GetFromIDShop(idShop uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ModuleCountryMgr) GetBatchFromIDShop(idShops []uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ModuleCountryMgr) GetFromIDCountry(idCountry uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_ModuleCountryMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}


func (obj *_ModuleCountryMgr) FetchByPrimaryKey(idModule uint32, idShop uint32, idCountry uint32) (result ModuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_country = ?", idModule, idShop, idCountry).Find(&result).Error

	return
}
