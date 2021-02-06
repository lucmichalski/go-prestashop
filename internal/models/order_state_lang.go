package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderStateLangMgr struct {
	*_BaseMgr
}

func OrderStateLangMgr(db *gorm.DB) *_OrderStateLangMgr {
	if db == nil {
		panic(fmt.Errorf("OrderStateLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderStateLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_state_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_OrderStateLangMgr) GetTableName() string {
	return "ps_order_state_lang"
}

func (obj *_OrderStateLangMgr) Get() (result OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_OrderStateLangMgr) Gets() (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_OrderStateLangMgr) WithIDOrderState(idOrderState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_state"] = idOrderState })
}

func (obj *_OrderStateLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_OrderStateLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_OrderStateLangMgr) WithTemplate(template string) Option {
	return optionFunc(func(o *options) { o.query["template"] = template })
}

func (obj *_OrderStateLangMgr) GetByOption(opts ...Option) (result OrderStateLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_OrderStateLangMgr) GetByOptions(opts ...Option) (results []*OrderStateLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_OrderStateLangMgr) GetFromIDOrderState(idOrderState uint32) (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ?", idOrderState).Find(&results).Error

	return
}

func (obj *_OrderStateLangMgr) GetBatchFromIDOrderState(idOrderStates []uint32) (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state IN (?)", idOrderStates).Find(&results).Error

	return
}

func (obj *_OrderStateLangMgr) GetFromIDLang(idLang uint32) (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_OrderStateLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_OrderStateLangMgr) GetFromName(name string) (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_OrderStateLangMgr) GetBatchFromName(names []string) (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_OrderStateLangMgr) GetFromTemplate(template string) (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template = ?", template).Find(&results).Error

	return
}

func (obj *_OrderStateLangMgr) GetBatchFromTemplate(templates []string) (results []*OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("template IN (?)", templates).Find(&results).Error

	return
}


func (obj *_OrderStateLangMgr) FetchByPrimaryKey(idOrderState uint32, idLang uint32) (result OrderStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_state = ? AND id_lang = ?", idOrderState, idLang).Find(&result).Error

	return
}
