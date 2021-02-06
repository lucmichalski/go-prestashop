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

// PsgdprConsentMgr open func
func PsgdprConsentMgr(db *gorm.DB) *_PsgdprConsentMgr {
	if db == nil {
		panic(fmt.Errorf("PsgdprConsentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PsgdprConsentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_psgdpr_consent"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PsgdprConsentMgr) GetTableName() string {
	return "eg_psgdpr_consent"
}

// Get 获取
func (obj *_PsgdprConsentMgr) Get() (result PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PsgdprConsentMgr) Gets() (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGdprConsent id_gdpr_consent获取
func (obj *_PsgdprConsentMgr) WithIDGdprConsent(idGdprConsent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gdpr_consent"] = idGdprConsent })
}

// WithIDModule id_module获取
func (obj *_PsgdprConsentMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithActive active获取
func (obj *_PsgdprConsentMgr) WithActive(active int) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithError error获取
func (obj *_PsgdprConsentMgr) WithError(error int) Option {
	return optionFunc(func(o *options) { o.query["error"] = error })
}

// WithErrorMessage error_message获取
func (obj *_PsgdprConsentMgr) WithErrorMessage(errorMessage string) Option {
	return optionFunc(func(o *options) { o.query["error_message"] = errorMessage })
}

// WithDateAdd date_add获取
func (obj *_PsgdprConsentMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_PsgdprConsentMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDGdprConsent 通过id_gdpr_consent获取内容
func (obj *_PsgdprConsentMgr) GetFromIDGdprConsent(idGdprConsent uint32) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ?", idGdprConsent).Find(&results).Error

	return
}

// GetBatchFromIDGdprConsent 批量唯一主键查找
func (obj *_PsgdprConsentMgr) GetBatchFromIDGdprConsent(idGdprConsents []uint32) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent IN (?)", idGdprConsents).Find(&results).Error

	return
}

// GetFromIDModule 通过id_module获取内容
func (obj *_PsgdprConsentMgr) GetFromIDModule(idModule uint32) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_PsgdprConsentMgr) GetBatchFromIDModule(idModules []uint32) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_PsgdprConsentMgr) GetFromActive(active int) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_PsgdprConsentMgr) GetBatchFromActive(actives []int) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromError 通过error获取内容
func (obj *_PsgdprConsentMgr) GetFromError(error int) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error = ?", error).Find(&results).Error

	return
}

// GetBatchFromError 批量唯一主键查找
func (obj *_PsgdprConsentMgr) GetBatchFromError(errors []int) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error IN (?)", errors).Find(&results).Error

	return
}

// GetFromErrorMessage 通过error_message获取内容
func (obj *_PsgdprConsentMgr) GetFromErrorMessage(errorMessage string) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_message = ?", errorMessage).Find(&results).Error

	return
}

// GetBatchFromErrorMessage 批量唯一主键查找
func (obj *_PsgdprConsentMgr) GetBatchFromErrorMessage(errorMessages []string) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_message IN (?)", errorMessages).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_PsgdprConsentMgr) GetFromDateAdd(dateAdd time.Time) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_PsgdprConsentMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_PsgdprConsentMgr) GetFromDateUpd(dateUpd time.Time) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_PsgdprConsentMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PsgdprConsentMgr) FetchByPrimaryKey(idGdprConsent uint32, idModule uint32) (result PsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ? AND id_module = ?", idGdprConsent, idModule).Find(&result).Error

	return
}
