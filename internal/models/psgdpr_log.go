package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _PsgdprLogMgr struct {
	*_BaseMgr
}

// PsgdprLogMgr open func
func PsgdprLogMgr(db *gorm.DB) *_PsgdprLogMgr {
	if db == nil {
		panic(fmt.Errorf("PsgdprLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PsgdprLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_psgdpr_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PsgdprLogMgr) GetTableName() string {
	return "ps_psgdpr_log"
}

// Get 获取
func (obj *_PsgdprLogMgr) Get() (result PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PsgdprLogMgr) Gets() (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGdprLog id_gdpr_log获取
func (obj *_PsgdprLogMgr) WithIDGdprLog(idGdprLog uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gdpr_log"] = idGdprLog })
}

// WithIDCustomer id_customer获取
func (obj *_PsgdprLogMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDGuest id_guest获取
func (obj *_PsgdprLogMgr) WithIDGuest(idGuest uint32) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

// WithClientName client_name获取
func (obj *_PsgdprLogMgr) WithClientName(clientName string) Option {
	return optionFunc(func(o *options) { o.query["client_name"] = clientName })
}

// WithIDModule id_module获取
func (obj *_PsgdprLogMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithRequestType request_type获取
func (obj *_PsgdprLogMgr) WithRequestType(requestType int) Option {
	return optionFunc(func(o *options) { o.query["request_type"] = requestType })
}

// WithDateAdd date_add获取
func (obj *_PsgdprLogMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_PsgdprLogMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDGdprLog 通过id_gdpr_log获取内容
func (obj *_PsgdprLogMgr) GetFromIDGdprLog(idGdprLog uint32) (result PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log = ?", idGdprLog).Find(&result).Error

	return
}

// GetBatchFromIDGdprLog 批量唯一主键查找
func (obj *_PsgdprLogMgr) GetBatchFromIDGdprLog(idGdprLogs []uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log IN (?)", idGdprLogs).Find(&results).Error

	return
}

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_PsgdprLogMgr) GetFromIDCustomer(idCustomer uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_PsgdprLogMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDGuest 通过id_guest获取内容
func (obj *_PsgdprLogMgr) GetFromIDGuest(idGuest uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error

	return
}

// GetBatchFromIDGuest 批量唯一主键查找
func (obj *_PsgdprLogMgr) GetBatchFromIDGuest(idGuests []uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error

	return
}

// GetFromClientName 通过client_name获取内容
func (obj *_PsgdprLogMgr) GetFromClientName(clientName string) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("client_name = ?", clientName).Find(&results).Error

	return
}

// GetBatchFromClientName 批量唯一主键查找
func (obj *_PsgdprLogMgr) GetBatchFromClientName(clientNames []string) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("client_name IN (?)", clientNames).Find(&results).Error

	return
}

// GetFromIDModule 通过id_module获取内容
func (obj *_PsgdprLogMgr) GetFromIDModule(idModule uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_PsgdprLogMgr) GetBatchFromIDModule(idModules []uint32) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromRequestType 通过request_type获取内容
func (obj *_PsgdprLogMgr) GetFromRequestType(requestType int) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_type = ?", requestType).Find(&results).Error

	return
}

// GetBatchFromRequestType 批量唯一主键查找
func (obj *_PsgdprLogMgr) GetBatchFromRequestType(requestTypes []int) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_type IN (?)", requestTypes).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_PsgdprLogMgr) GetFromDateAdd(dateAdd time.Time) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_PsgdprLogMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_PsgdprLogMgr) GetFromDateUpd(dateUpd time.Time) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_PsgdprLogMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PsgdprLogMgr) FetchByPrimaryKey(idGdprLog uint32) (result PsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log = ?", idGdprLog).Find(&result).Error

	return
}
