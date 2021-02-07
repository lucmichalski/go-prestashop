package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CustomerMessageMgr struct {
	*_BaseMgr
}

func CustomerMessageMgr(db *gorm.DB) *_CustomerMessageMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customer_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomerMessageMgr) GetTableName() string {
	return "ps_customer_message"
}

func (obj *_CustomerMessageMgr) Get() (result CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomerMessageMgr) Gets() (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) WithIDCustomerMessage(idCustomerMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer_message"] = idCustomerMessage })
}

func (obj *_CustomerMessageMgr) WithIDCustomerThread(idCustomerThread int) Option {
	return optionFunc(func(o *options) { o.query["id_customer_thread"] = idCustomerThread })
}

func (obj *_CustomerMessageMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_CustomerMessageMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

func (obj *_CustomerMessageMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

func (obj *_CustomerMessageMgr) WithIPAddress(ipAddress string) Option {
	return optionFunc(func(o *options) { o.query["ip_address"] = ipAddress })
}

func (obj *_CustomerMessageMgr) WithUserAgent(userAgent string) Option {
	return optionFunc(func(o *options) { o.query["user_agent"] = userAgent })
}

func (obj *_CustomerMessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_CustomerMessageMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_CustomerMessageMgr) WithPrivate(private int8) Option {
	return optionFunc(func(o *options) { o.query["private"] = private })
}

func (obj *_CustomerMessageMgr) WithRead(read bool) Option {
	return optionFunc(func(o *options) { o.query["read"] = read })
}

func (obj *_CustomerMessageMgr) GetByOption(opts ...Option) (result CustomerMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CustomerMessageMgr) GetByOptions(opts ...Option) (results []*CustomerMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromIDCustomerMessage(idCustomerMessage uint32) (result CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message = ?", idCustomerMessage).Find(&result).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromIDCustomerMessage(idCustomerMessages []uint32) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message IN (?)", idCustomerMessages).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromIDCustomerThread(idCustomerThread int) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread = ?", idCustomerThread).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromIDCustomerThread(idCustomerThreads []int) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread IN (?)", idCustomerThreads).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromIDEmployee(idEmployee uint32) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromMessage(message string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromMessage(messages []string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromFileName(fileName string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromFileName(fileNames []string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromIPAddress(ipAddress string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address = ?", ipAddress).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromIPAddress(ipAddresss []string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address IN (?)", ipAddresss).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromUserAgent(userAgent string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_agent = ?", userAgent).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromUserAgent(userAgents []string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_agent IN (?)", userAgents).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromDateUpd(dateUpd time.Time) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromPrivate(private int8) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private = ?", private).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromPrivate(privates []int8) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private IN (?)", privates).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetFromRead(read bool) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("read = ?", read).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) GetBatchFromRead(reads []bool) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("read IN (?)", reads).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) FetchByPrimaryKey(idCustomerMessage uint32) (result CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message = ?", idCustomerMessage).Find(&result).Error

	return
}

func (obj *_CustomerMessageMgr) FetchIndexByIDCustomerThread(idCustomerThread int) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread = ?", idCustomerThread).Find(&results).Error

	return
}

func (obj *_CustomerMessageMgr) FetchIndexByIDEmployee(idEmployee uint32) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}
