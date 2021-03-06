package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _FeatureLangMgr struct {
	*_BaseMgr
}

func FeatureLangMgr(db *gorm.DB) *_FeatureLangMgr {
	if db == nil {
		panic(fmt.Errorf("FeatureLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FeatureLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_feature_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_FeatureLangMgr) GetTableName() string {
	return "ps_feature_lang"
}

func (obj *_FeatureLangMgr) Get() (result FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_FeatureLangMgr) Gets() (results []*FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_FeatureLangMgr) WithIDFeature(idFeature uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

func (obj *_FeatureLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_FeatureLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_FeatureLangMgr) GetByOption(opts ...Option) (result FeatureLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_FeatureLangMgr) GetByOptions(opts ...Option) (results []*FeatureLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_FeatureLangMgr) GetFromIDFeature(idFeature uint32) (results []*FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error

	return
}

func (obj *_FeatureLangMgr) GetBatchFromIDFeature(idFeatures []uint32) (results []*FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error

	return
}

func (obj *_FeatureLangMgr) GetFromIDLang(idLang uint32) (results []*FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_FeatureLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_FeatureLangMgr) GetFromName(name string) (results []*FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_FeatureLangMgr) GetBatchFromName(names []string) (results []*FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_FeatureLangMgr) FetchByPrimaryKey(idFeature uint32, idLang uint32) (result FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ? AND id_lang = ?", idFeature, idLang).Find(&result).Error

	return
}

func (obj *_FeatureLangMgr) FetchIndexByIDLang(idLang uint32, name string) (results []*FeatureLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ? AND name = ?", idLang, name).Find(&results).Error

	return
}
