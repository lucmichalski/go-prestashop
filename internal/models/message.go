package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MessageMgr struct {
	*_BaseMgr
}

// MessageMgr open func
func MessageMgr(db *gorm.DB) *_MessageMgr {
	if db == nil {
		panic(fmt.Errorf("MessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MessageMgr) GetTableName() string {
	return "ps_message"
}

// Get 获取
func (obj *_MessageMgr) Get() (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MessageMgr) Gets() (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMessage id_message获取
func (obj *_MessageMgr) WithIDMessage(idMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_message"] = idMessage })
}

// WithIDCart id_cart获取
func (obj *_MessageMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDCustomer id_customer获取
func (obj *_MessageMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDEmployee id_employee获取
func (obj *_MessageMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDOrder id_order获取
func (obj *_MessageMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithMessage message获取
func (obj *_MessageMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithPrivate private获取
func (obj *_MessageMgr) WithPrivate(private bool) Option {
	return optionFunc(func(o *options) { o.query["private"] = private })
}

// WithDateAdd date_add获取
func (obj *_MessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
func (obj *_MessageMgr) GetByOption(opts ...Option) (result Message, err error) {
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
func (obj *_MessageMgr) GetByOptions(opts ...Option) (results []*Message, err error) {
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

// GetFromIDMessage 通过id_message获取内容
func (obj *_MessageMgr) GetFromIDMessage(idMessage uint32) (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ?", idMessage).Find(&result).Error

	return
}

// GetBatchFromIDMessage 批量唯一主键查找
func (obj *_MessageMgr) GetBatchFromIDMessage(idMessages []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message IN (?)", idMessages).Find(&results).Error

	return
}

// GetFromIDCart 通过id_cart获取内容
func (obj *_MessageMgr) GetFromIDCart(idCart uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

// GetBatchFromIDCart 批量唯一主键查找
func (obj *_MessageMgr) GetBatchFromIDCart(idCarts []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_MessageMgr) GetFromIDCustomer(idCustomer uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_MessageMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_MessageMgr) GetFromIDEmployee(idEmployee uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_MessageMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromIDOrder 通过id_order获取内容
func (obj *_MessageMgr) GetFromIDOrder(idOrder uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

// GetBatchFromIDOrder 批量唯一主键查找
func (obj *_MessageMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容
func (obj *_MessageMgr) GetFromMessage(message string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量唯一主键查找
func (obj *_MessageMgr) GetBatchFromMessage(messages []string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

// GetFromPrivate 通过private获取内容
func (obj *_MessageMgr) GetFromPrivate(private bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private = ?", private).Find(&results).Error

	return
}

// GetBatchFromPrivate 批量唯一主键查找
func (obj *_MessageMgr) GetBatchFromPrivate(privates []bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private IN (?)", privates).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_MessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_MessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MessageMgr) FetchByPrimaryKey(idMessage uint32) (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ?", idMessage).Find(&result).Error

	return
}

// FetchIndexByIDCart  获取多个内容
func (obj *_MessageMgr) FetchIndexByIDCart(idCart uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

// FetchIndexByIDCustomer  获取多个内容
func (obj *_MessageMgr) FetchIndexByIDCustomer(idCustomer uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// FetchIndexByIDEmployee  获取多个内容
func (obj *_MessageMgr) FetchIndexByIDEmployee(idEmployee uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// FetchIndexByMessageOrder  获取多个内容
func (obj *_MessageMgr) FetchIndexByMessageOrder(idOrder uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}
