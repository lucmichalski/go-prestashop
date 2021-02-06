package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleCarrierMgr struct {
	*_BaseMgr
}

func ModuleCarrierMgr(db *gorm.DB) *_ModuleCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_module_carrier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ModuleCarrierMgr) GetTableName() string {
	return "ps_module_carrier"
}

func (obj *_ModuleCarrierMgr) Get() (result ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ModuleCarrierMgr) Gets() (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ModuleCarrierMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_ModuleCarrierMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ModuleCarrierMgr) WithIDReference(idReference int) Option {
	return optionFunc(func(o *options) { o.query["id_reference"] = idReference })
}

func (obj *_ModuleCarrierMgr) GetByOption(opts ...Option) (result ModuleCarrier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ModuleCarrierMgr) GetByOptions(opts ...Option) (results []*ModuleCarrier, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ModuleCarrierMgr) GetFromIDModule(idModule uint32) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_ModuleCarrierMgr) GetBatchFromIDModule(idModules []uint32) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_ModuleCarrierMgr) GetFromIDShop(idShop uint32) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_ModuleCarrierMgr) GetBatchFromIDShop(idShops []uint32) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ModuleCarrierMgr) GetFromIDReference(idReference int) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference = ?", idReference).Find(&results).Error

	return
}

func (obj *_ModuleCarrierMgr) GetBatchFromIDReference(idReferences []int) (results []*ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference IN (?)", idReferences).Find(&results).Error

	return
}


func (obj *_ModuleCarrierMgr) FetchByPrimaryKey(idModule uint32, idShop uint32, idReference int) (result ModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_reference = ?", idModule, idShop, idReference).Find(&result).Error

	return
}
