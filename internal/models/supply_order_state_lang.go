package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SupplyOrderStateLangMgr struct {
	*_BaseMgr
}

func SupplyOrderStateLangMgr(db *gorm.DB) *_SupplyOrderStateLangMgr {
	if db == nil {
		panic(fmt.Errorf("SupplyOrderStateLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupplyOrderStateLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_supply_order_state_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SupplyOrderStateLangMgr) GetTableName() string {
	return "ps_supply_order_state_lang"
}

func (obj *_SupplyOrderStateLangMgr) Get() (result SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) Gets() (results []*SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) WithIDSupplyOrderState(idSupplyOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supply_order_state"] = idSupplyOrderState })
}

func (obj *_SupplyOrderStateLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_SupplyOrderStateLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_SupplyOrderStateLangMgr) GetByOption(opts ...Option) (result SupplyOrderStateLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) GetByOptions(opts ...Option) (results []*SupplyOrderStateLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) GetFromIDSupplyOrderState(idSupplyOrderState uint32) (results []*SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ?", idSupplyOrderState).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) GetBatchFromIDSupplyOrderState(idSupplyOrderStates []uint32) (results []*SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state IN (?)", idSupplyOrderStates).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) GetFromIDLang(idLang uint32) (results []*SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) GetFromName(name string) (results []*SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) GetBatchFromName(names []string) (results []*SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_SupplyOrderStateLangMgr) FetchByPrimaryKey(idSupplyOrderState uint32, idLang uint32) (result SupplyOrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supply_order_state = ? AND id_lang = ?", idSupplyOrderState, idLang).Find(&result).Error

	return
}
