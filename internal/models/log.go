package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _LogMgr struct {
	*_BaseMgr
}

func LogMgr(db *gorm.DB) *_LogMgr {
	if db == nil {
		panic(fmt.Errorf("LogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LogMgr) GetTableName() string {
	return "ps_log"
}

func (obj *_LogMgr) Get() (result Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LogMgr) Gets() (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LogMgr) WithIDLog(idLog uint32) Option {
	return optionFunc(func(o *options) { o.query["id_log"] = idLog })
}

func (obj *_LogMgr) WithSeverity(severity bool) Option {
	return optionFunc(func(o *options) { o.query["severity"] = severity })
}

func (obj *_LogMgr) WithErrorCode(errorCode int) Option {
	return optionFunc(func(o *options) { o.query["error_code"] = errorCode })
}

func (obj *_LogMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

func (obj *_LogMgr) WithObjectType(objectType string) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

func (obj *_LogMgr) WithObjectID(objectID uint32) Option {
	return optionFunc(func(o *options) { o.query["object_id"] = objectID })
}

func (obj *_LogMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_LogMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_LogMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_LogMgr) GetByOption(opts ...Option) (result Log, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LogMgr) GetByOptions(opts ...Option) (results []*Log, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LogMgr) GetFromIDLog(idLog uint32) (result Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_log = ?", idLog).Find(&result).Error

	return
}

func (obj *_LogMgr) GetBatchFromIDLog(idLogs []uint32) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_log IN (?)", idLogs).Find(&results).Error

	return
}

func (obj *_LogMgr) GetFromSeverity(severity bool) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("severity = ?", severity).Find(&results).Error

	return
}

func (obj *_LogMgr) GetBatchFromSeverity(severitys []bool) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("severity IN (?)", severitys).Find(&results).Error

	return
}

func (obj *_LogMgr) GetFromErrorCode(errorCode int) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_code = ?", errorCode).Find(&results).Error

	return
}

func (obj *_LogMgr) GetBatchFromErrorCode(errorCodes []int) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_code IN (?)", errorCodes).Find(&results).Error

	return
}

func (obj *_LogMgr) GetFromMessage(message string) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

func (obj *_LogMgr) GetBatchFromMessage(messages []string) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

func (obj *_LogMgr) GetFromObjectType(objectType string) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_type = ?", objectType).Find(&results).Error

	return
}

func (obj *_LogMgr) GetBatchFromObjectType(objectTypes []string) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_type IN (?)", objectTypes).Find(&results).Error

	return
}

func (obj *_LogMgr) GetFromObjectID(objectID uint32) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_id = ?", objectID).Find(&results).Error

	return
}

func (obj *_LogMgr) GetBatchFromObjectID(objectIDs []uint32) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_id IN (?)", objectIDs).Find(&results).Error

	return
}

func (obj *_LogMgr) GetFromIDEmployee(idEmployee uint32) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_LogMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_LogMgr) GetFromDateAdd(dateAdd time.Time) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_LogMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_LogMgr) GetFromDateUpd(dateUpd time.Time) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_LogMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_LogMgr) FetchByPrimaryKey(idLog uint32) (result Log, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_log = ?", idLog).Find(&result).Error

	return
}
