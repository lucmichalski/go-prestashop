package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderReturnStateLangMgr struct {
	*_BaseMgr
}

func OrderReturnStateLangMgr(db *gorm.DB) *_OrderReturnStateLangMgr {
	if db == nil {
		panic(fmt.Errorf("OrderReturnStateLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderReturnStateLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_return_state_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderReturnStateLangMgr) GetTableName() string {
	return "ps_order_return_state_lang"
}

func (obj *_OrderReturnStateLangMgr) Get() (result OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderReturnStateLangMgr) Gets() (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_OrderReturnStateLangMgr) WithIDOrderReturnState(idOrderReturnState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return_state"] = idOrderReturnState })
}

func (obj *_OrderReturnStateLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_OrderReturnStateLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_OrderReturnStateLangMgr) GetByOption(opts ...Option) (result OrderReturnStateLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderReturnStateLangMgr) GetByOptions(opts ...Option) (results []*OrderReturnStateLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_OrderReturnStateLangMgr) GetFromIDOrderReturnState(idOrderReturnState uint32) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ?", idOrderReturnState).Find(&results).Error

	return
}

func (obj *_OrderReturnStateLangMgr) GetBatchFromIDOrderReturnState(idOrderReturnStates []uint32) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state IN (?)", idOrderReturnStates).Find(&results).Error

	return
}

func (obj *_OrderReturnStateLangMgr) GetFromIDLang(idLang uint32) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_OrderReturnStateLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_OrderReturnStateLangMgr) GetFromName(name string) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_OrderReturnStateLangMgr) GetBatchFromName(names []string) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_OrderReturnStateLangMgr) FetchByPrimaryKey(idOrderReturnState uint32, idLang uint32) (result OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ? AND id_lang = ?", idOrderReturnState, idLang).Find(&result).Error

	return
}
