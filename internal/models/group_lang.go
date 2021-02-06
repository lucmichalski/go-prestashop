package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GroupLangMgr struct {
	*_BaseMgr
}

func GroupLangMgr(db *gorm.DB) *_GroupLangMgr {
	if db == nil {
		panic(fmt.Errorf("GroupLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GroupLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_group_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_GroupLangMgr) GetTableName() string {
	return "ps_group_lang"
}

func (obj *_GroupLangMgr) Get() (result GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_GroupLangMgr) Gets() (results []*GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_GroupLangMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_GroupLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_GroupLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_GroupLangMgr) GetByOption(opts ...Option) (result GroupLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_GroupLangMgr) GetByOptions(opts ...Option) (results []*GroupLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_GroupLangMgr) GetFromIDGroup(idGroup uint32) (results []*GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_GroupLangMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_GroupLangMgr) GetFromIDLang(idLang uint32) (results []*GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_GroupLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_GroupLangMgr) GetFromName(name string) (results []*GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_GroupLangMgr) GetBatchFromName(names []string) (results []*GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_GroupLangMgr) FetchByPrimaryKey(idGroup uint32, idLang uint32) (result GroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_lang = ?", idGroup, idLang).Find(&result).Error

	return
}
