package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttributeLangMgr struct {
	*_BaseMgr
}

func AttributeLangMgr(db *gorm.DB) *_AttributeLangMgr {
	if db == nil {
		panic(fmt.Errorf("AttributeLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttributeLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_attribute_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AttributeLangMgr) GetTableName() string {
	return "ps_attribute_lang"
}

func (obj *_AttributeLangMgr) Get() (result AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AttributeLangMgr) Gets() (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_AttributeLangMgr) WithIDAttribute(idAttribute int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

func (obj *_AttributeLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_AttributeLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_AttributeLangMgr) GetByOption(opts ...Option) (result AttributeLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AttributeLangMgr) GetByOptions(opts ...Option) (results []*AttributeLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_AttributeLangMgr) GetFromIDAttribute(idAttribute int) (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error

	return
}

func (obj *_AttributeLangMgr) GetBatchFromIDAttribute(idAttributes []int) (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error

	return
}

func (obj *_AttributeLangMgr) GetFromIDLang(idLang int) (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_AttributeLangMgr) GetBatchFromIDLang(idLangs []int) (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_AttributeLangMgr) GetFromName(name string) (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_AttributeLangMgr) GetBatchFromName(names []string) (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_AttributeLangMgr) FetchByPrimaryKey(idAttribute int, idLang int) (result AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_lang = ?", idAttribute, idLang).Find(&result).Error

	return
}

func (obj *_AttributeLangMgr) FetchIndexByIDXEB57414F7A4F53DC(idAttribute int) (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error

	return
}

func (obj *_AttributeLangMgr) FetchIndexByIDXEB57414FBA299860(idLang int) (results []*AttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}
