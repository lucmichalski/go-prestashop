package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _PsgdprLogMgr struct {
	*_BaseMgr
}

func PsgdprLogMgr(db *gorm.DB) *_PsgdprLogMgr {
	if db == nil {
		panic(fmt.Errorf("PsgdprLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PsgdprLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_psgdpr_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_PsgdprLogMgr) GetTableName() string {
	return "ps_psgdpr_log"
}

func (obj *_PsgdprLogMgr) Get() (result PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_PsgdprLogMgr) Gets() (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_PsgdprLogMgr) WithIDGdprLog(idGdprLog uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gdpr_log"] = idGdprLog })
}

func (obj *_PsgdprLogMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_PsgdprLogMgr) WithIDGuest(idGuest uint32) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

func (obj *_PsgdprLogMgr) WithClientName(clientName string) Option {
	return optionFunc(func(o *options) { o.query["client_name"] = clientName })
}

func (obj *_PsgdprLogMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

func (obj *_PsgdprLogMgr) WithRequestType(requestType int) Option {
	return optionFunc(func(o *options) { o.query["request_type"] = requestType })
}

func (obj *_PsgdprLogMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_PsgdprLogMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_PsgdprLogMgr) GetByOption(opts ...Option) (result PsgdprLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_PsgdprLogMgr) GetByOptions(opts ...Option) (results []*PsgdprLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_PsgdprLogMgr) GetFromIDGdprLog(idGdprLog uint32) (result PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log = ?", idGdprLog).Find(&result).Error

	return
}

func (obj *_PsgdprLogMgr) GetBatchFromIDGdprLog(idGdprLogs []uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log IN (?)", idGdprLogs).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetFromIDCustomer(idCustomer uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetFromIDGuest(idGuest uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetBatchFromIDGuest(idGuests []uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetFromClientName(clientName string) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("client_name = ?", clientName).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetBatchFromClientName(clientNames []string) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("client_name IN (?)", clientNames).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetFromIDModule(idModule uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetBatchFromIDModule(idModules []uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetFromRequestType(requestType int) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_type = ?", requestType).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetBatchFromRequestType(requestTypes []int) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_type IN (?)", requestTypes).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetFromDateAdd(dateAdd time.Time) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetFromDateUpd(dateUpd time.Time) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_PsgdprLogMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}


func (obj *_PsgdprLogMgr) FetchByPrimaryKey(idGdprLog uint32) (result PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log = ?", idGdprLog).Find(&result).Error

	return
}
