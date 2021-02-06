package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CustomerMessageMgr struct {
	*_BaseMgr
}

// CustomerMessageMgr open func
func CustomerMessageMgr(db *gorm.DB) *_CustomerMessageMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CustomerMessageMgr) GetTableName() string {
	return "eg_customer_message"
}

// Get 获取
func (obj *_CustomerMessageMgr) Get() (result CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CustomerMessageMgr) Gets() (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomerMessage id_customer_message获取
func (obj *_CustomerMessageMgr) WithIDCustomerMessage(idCustomerMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer_message"] = idCustomerMessage })
}

// WithIDCustomerThread id_customer_thread获取
func (obj *_CustomerMessageMgr) WithIDCustomerThread(idCustomerThread int) Option {
	return optionFunc(func(o *options) { o.query["id_customer_thread"] = idCustomerThread })
}

// WithIDEmployee id_employee获取
func (obj *_CustomerMessageMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithMessage message获取
func (obj *_CustomerMessageMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithFileName file_name获取
func (obj *_CustomerMessageMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

// WithIPAddress ip_address获取
func (obj *_CustomerMessageMgr) WithIPAddress(ipAddress string) Option {
	return optionFunc(func(o *options) { o.query["ip_address"] = ipAddress })
}

// WithUserAgent user_agent获取
func (obj *_CustomerMessageMgr) WithUserAgent(userAgent string) Option {
	return optionFunc(func(o *options) { o.query["user_agent"] = userAgent })
}

// WithDateAdd date_add获取
func (obj *_CustomerMessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_CustomerMessageMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithPrivate private获取
func (obj *_CustomerMessageMgr) WithPrivate(private int8) Option {
	return optionFunc(func(o *options) { o.query["private"] = private })
}

// WithRead read获取
func (obj *_CustomerMessageMgr) WithRead(read bool) Option {
	return optionFunc(func(o *options) { o.query["read"] = read })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCustomerMessage 通过id_customer_message获取内容
func (obj *_CustomerMessageMgr) GetFromIDCustomerMessage(idCustomerMessage uint32) (result CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message = ?", idCustomerMessage).Find(&result).Error

	return
}

// GetBatchFromIDCustomerMessage 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromIDCustomerMessage(idCustomerMessages []uint32) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message IN (?)", idCustomerMessages).Find(&results).Error

	return
}

// GetFromIDCustomerThread 通过id_customer_thread获取内容
func (obj *_CustomerMessageMgr) GetFromIDCustomerThread(idCustomerThread int) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread = ?", idCustomerThread).Find(&results).Error

	return
}

// GetBatchFromIDCustomerThread 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromIDCustomerThread(idCustomerThreads []int) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread IN (?)", idCustomerThreads).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_CustomerMessageMgr) GetFromIDEmployee(idEmployee uint32) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容
func (obj *_CustomerMessageMgr) GetFromMessage(message string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromMessage(messages []string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

// GetFromFileName 通过file_name获取内容
func (obj *_CustomerMessageMgr) GetFromFileName(fileName string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error

	return
}

// GetBatchFromFileName 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromFileName(fileNames []string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error

	return
}

// GetFromIPAddress 通过ip_address获取内容
func (obj *_CustomerMessageMgr) GetFromIPAddress(ipAddress string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address = ?", ipAddress).Find(&results).Error

	return
}

// GetBatchFromIPAddress 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromIPAddress(ipAddresss []string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address IN (?)", ipAddresss).Find(&results).Error

	return
}

// GetFromUserAgent 通过user_agent获取内容
func (obj *_CustomerMessageMgr) GetFromUserAgent(userAgent string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_agent = ?", userAgent).Find(&results).Error

	return
}

// GetBatchFromUserAgent 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromUserAgent(userAgents []string) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_agent IN (?)", userAgents).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_CustomerMessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_CustomerMessageMgr) GetFromDateUpd(dateUpd time.Time) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

// GetFromPrivate 通过private获取内容
func (obj *_CustomerMessageMgr) GetFromPrivate(private int8) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private = ?", private).Find(&results).Error

	return
}

// GetBatchFromPrivate 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromPrivate(privates []int8) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private IN (?)", privates).Find(&results).Error

	return
}

// GetFromRead 通过read获取内容
func (obj *_CustomerMessageMgr) GetFromRead(read bool) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("read = ?", read).Find(&results).Error

	return
}

// GetBatchFromRead 批量唯一主键查找
func (obj *_CustomerMessageMgr) GetBatchFromRead(reads []bool) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("read IN (?)", reads).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CustomerMessageMgr) FetchByPrimaryKey(idCustomerMessage uint32) (result CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message = ?", idCustomerMessage).Find(&result).Error

	return
}

// FetchIndexByIDCustomerThread  获取多个内容
func (obj *_CustomerMessageMgr) FetchIndexByIDCustomerThread(idCustomerThread int) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread = ?", idCustomerThread).Find(&results).Error

	return
}

// FetchIndexByIDEmployee  获取多个内容
func (obj *_CustomerMessageMgr) FetchIndexByIDEmployee(idEmployee uint32) (results []*CustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}
