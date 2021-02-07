package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PsgdprConsentLangMgr struct {
	*_BaseMgr
}

func PsgdprConsentLangMgr(db *gorm.DB) *_PsgdprConsentLangMgr {
	if db == nil {
		panic(fmt.Errorf("PsgdprConsentLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PsgdprConsentLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_psgdpr_consent_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PsgdprConsentLangMgr) GetTableName() string {
	return "ps_psgdpr_consent_lang"
}

func (obj *_PsgdprConsentLangMgr) Get() (result PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PsgdprConsentLangMgr) Gets() (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) WithIDGdprConsent(idGdprConsent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gdpr_consent"] = idGdprConsent })
}

func (obj *_PsgdprConsentLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_PsgdprConsentLangMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

func (obj *_PsgdprConsentLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_PsgdprConsentLangMgr) GetByOption(opts ...Option) (result PsgdprConsentLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetByOptions(opts ...Option) (results []*PsgdprConsentLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetFromIDGdprConsent(idGdprConsent uint32) (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ?", idGdprConsent).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetBatchFromIDGdprConsent(idGdprConsents []uint32) (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent IN (?)", idGdprConsents).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetFromIDLang(idLang uint32) (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetFromMessage(message string) (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetBatchFromMessage(messages []string) (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetFromIDShop(idShop uint32) (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_PsgdprConsentLangMgr) FetchByPrimaryKey(idGdprConsent uint32, idLang uint32, idShop uint32) (result PsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ? AND id_lang = ? AND id_shop = ?", idGdprConsent, idLang, idShop).Find(&result).Error

	return
}
