package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _InfoLangMgr struct {
	*_BaseMgr
}

func InfoLangMgr(db *gorm.DB) *_InfoLangMgr {
	if db == nil {
		panic(fmt.Errorf("InfoLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfoLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_info_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_InfoLangMgr) GetTableName() string {
	return "ps_info_lang"
}

func (obj *_InfoLangMgr) Get() (result InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_InfoLangMgr) Gets() (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_InfoLangMgr) WithIDInfo(idInfo uint32) Option {
	return optionFunc(func(o *options) { o.query["id_info"] = idInfo })
}

func (obj *_InfoLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_InfoLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_InfoLangMgr) WithText(text string) Option {
	return optionFunc(func(o *options) { o.query["text"] = text })
}

func (obj *_InfoLangMgr) GetByOption(opts ...Option) (result InfoLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_InfoLangMgr) GetByOptions(opts ...Option) (results []*InfoLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_InfoLangMgr) GetFromIDInfo(idInfo uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&results).Error

	return
}

func (obj *_InfoLangMgr) GetBatchFromIDInfo(idInfos []uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info IN (?)", idInfos).Find(&results).Error

	return
}

func (obj *_InfoLangMgr) GetFromIDShop(idShop uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_InfoLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_InfoLangMgr) GetFromIDLang(idLang uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_InfoLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_InfoLangMgr) GetFromText(text string) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text = ?", text).Find(&results).Error

	return
}

func (obj *_InfoLangMgr) GetBatchFromText(texts []string) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text IN (?)", texts).Find(&results).Error

	return
}


func (obj *_InfoLangMgr) FetchByPrimaryKey(idInfo uint32, idShop uint32, idLang uint32) (result InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ? AND id_shop = ? AND id_lang = ?", idInfo, idShop, idLang).Find(&result).Error

	return
}
