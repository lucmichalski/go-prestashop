package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _PsgdprConsentMgr struct {
	*_BaseMgr
}

func PsgdprConsentMgr(db *gorm.DB) *_PsgdprConsentMgr {
	if db == nil {
		panic(fmt.Errorf("PsgdprConsentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PsgdprConsentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_psgdpr_consent"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PsgdprConsentMgr) GetTableName() string {
	return "ps_psgdpr_consent"
}

func (obj *_PsgdprConsentMgr) Get() (result PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PsgdprConsentMgr) Gets() (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PsgdprConsentMgr) WithIDGdprConsent(idGdprConsent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gdpr_consent"] = idGdprConsent })
}

func (obj *_PsgdprConsentMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_PsgdprConsentMgr) WithActive(active int) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_PsgdprConsentMgr) WithError(error int) Option {
	return optionFunc(func(o *options) { o.query["error"] = error })
}

func (obj *_PsgdprConsentMgr) WithErrorMessage(errorMessage string) Option {
	return optionFunc(func(o *options) { o.query["error_message"] = errorMessage })
}

func (obj *_PsgdprConsentMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_PsgdprConsentMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_PsgdprConsentMgr) GetByOption(opts ...Option) (result PsgdprConsent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_PsgdprConsentMgr) GetByOptions(opts ...Option) (results []*PsgdprConsent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_PsgdprConsentMgr) GetFromIDGdprConsent(idGdprConsent uint32) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ?", idGdprConsent).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetBatchFromIDGdprConsent(idGdprConsents []uint32) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent IN (?)", idGdprConsents).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetFromIDModule(idModule uint32) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetBatchFromIDModule(idModules []uint32) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetFromActive(active int) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetBatchFromActive(actives []int) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetFromError(error int) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error = ?", error).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetBatchFromError(errors []int) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error IN (?)", errors).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetFromErrorMessage(errorMessage string) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_message = ?", errorMessage).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetBatchFromErrorMessage(errorMessages []string) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_message IN (?)", errorMessages).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetFromDateAdd(dateAdd time.Time) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetFromDateUpd(dateUpd time.Time) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_PsgdprConsentMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_PsgdprConsentMgr) FetchByPrimaryKey(idGdprConsent uint32, idModule uint32) (result PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ? AND id_module = ?", idGdprConsent, idModule).Find(&result).Error

	return
}
