package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProfileLangMgr struct {
	*_BaseMgr
}

func ProfileLangMgr(db *gorm.DB) *_ProfileLangMgr {
	if db == nil {
		panic(fmt.Errorf("ProfileLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProfileLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_profile_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProfileLangMgr) GetTableName() string {
	return "ps_profile_lang"
}

func (obj *_ProfileLangMgr) Get() (result ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProfileLangMgr) Gets() (results []*ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProfileLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_ProfileLangMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}

func (obj *_ProfileLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ProfileLangMgr) GetByOption(opts ...Option) (result ProfileLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProfileLangMgr) GetByOptions(opts ...Option) (results []*ProfileLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProfileLangMgr) GetFromIDLang(idLang uint32) (results []*ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_ProfileLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_ProfileLangMgr) GetFromIDProfile(idProfile uint32) (results []*ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&results).Error

	return
}

func (obj *_ProfileLangMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error

	return
}

func (obj *_ProfileLangMgr) GetFromName(name string) (results []*ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ProfileLangMgr) GetBatchFromName(names []string) (results []*ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_ProfileLangMgr) FetchByPrimaryKey(idLang uint32, idProfile uint32) (result ProfileLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ? AND id_profile = ?", idLang, idProfile).Find(&result).Error

	return
}
