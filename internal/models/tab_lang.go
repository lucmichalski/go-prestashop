package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TabLangMgr struct {
	*_BaseMgr
}

func TabLangMgr(db *gorm.DB) *_TabLangMgr {
	if db == nil {
		panic(fmt.Errorf("TabLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TabLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tab_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_TabLangMgr) GetTableName() string {
	return "ps_tab_lang"
}

func (obj *_TabLangMgr) Get() (result TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_TabLangMgr) Gets() (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_TabLangMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

func (obj *_TabLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_TabLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_TabLangMgr) GetByOption(opts ...Option) (result TabLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_TabLangMgr) GetByOptions(opts ...Option) (results []*TabLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_TabLangMgr) GetFromIDTab(idTab int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

func (obj *_TabLangMgr) GetBatchFromIDTab(idTabs []int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

func (obj *_TabLangMgr) GetFromIDLang(idLang int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_TabLangMgr) GetBatchFromIDLang(idLangs []int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_TabLangMgr) GetFromName(name string) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_TabLangMgr) GetBatchFromName(names []string) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_TabLangMgr) FetchByPrimaryKey(idTab int, idLang int) (result TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ? AND id_lang = ?", idTab, idLang).Find(&result).Error

	return
}

func (obj *_TabLangMgr) FetchIndexByIDX3E3D6F36ED47AB56(idTab int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

func (obj *_TabLangMgr) FetchIndexByIDX3E3D6F36BA299860(idLang int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}
