package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StoreLangMgr struct {
	*_BaseMgr
}

func StoreLangMgr(db *gorm.DB) *_StoreLangMgr {
	if db == nil {
		panic(fmt.Errorf("StoreLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoreLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_store_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StoreLangMgr) GetTableName() string {
	return "ps_store_lang"
}

func (obj *_StoreLangMgr) Get() (result StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StoreLangMgr) Gets() (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_StoreLangMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

func (obj *_StoreLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_StoreLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_StoreLangMgr) WithAddress1(address1 string) Option {
	return optionFunc(func(o *options) { o.query["address1"] = address1 })
}

func (obj *_StoreLangMgr) WithAddress2(address2 string) Option {
	return optionFunc(func(o *options) { o.query["address2"] = address2 })
}

func (obj *_StoreLangMgr) WithHours(hours string) Option {
	return optionFunc(func(o *options) { o.query["hours"] = hours })
}

func (obj *_StoreLangMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

func (obj *_StoreLangMgr) GetByOption(opts ...Option) (result StoreLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_StoreLangMgr) GetByOptions(opts ...Option) (results []*StoreLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_StoreLangMgr) GetFromIDStore(idStore uint32) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetBatchFromIDStore(idStores []uint32) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetFromIDLang(idLang uint32) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetFromName(name string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetBatchFromName(names []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetFromAddress1(address1 string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 = ?", address1).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetBatchFromAddress1(address1s []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 IN (?)", address1s).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetFromAddress2(address2 string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 = ?", address2).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetBatchFromAddress2(address2s []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 IN (?)", address2s).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetFromHours(hours string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hours = ?", hours).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetBatchFromHours(hourss []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hours IN (?)", hourss).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetFromNote(note string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note = ?", note).Find(&results).Error

	return
}

func (obj *_StoreLangMgr) GetBatchFromNote(notes []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note IN (?)", notes).Find(&results).Error

	return
}


func (obj *_StoreLangMgr) FetchByPrimaryKey(idStore uint32, idLang uint32) (result StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ? AND id_lang = ?", idStore, idLang).Find(&result).Error

	return
}
