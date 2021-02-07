package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MessageMgr struct {
	*_BaseMgr
}

func MessageMgr(db *gorm.DB) *_MessageMgr {
	if db == nil {
		panic(fmt.Errorf("MessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_MessageMgr) GetTableName() string {
	return "ps_message"
}

func (obj *_MessageMgr) Get() (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_MessageMgr) Gets() (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_MessageMgr) WithIDMessage(idMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_message"] = idMessage })
}

func (obj *_MessageMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

func (obj *_MessageMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_MessageMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_MessageMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

func (obj *_MessageMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

func (obj *_MessageMgr) WithPrivate(private bool) Option {
	return optionFunc(func(o *options) { o.query["private"] = private })
}

func (obj *_MessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

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

func (obj *_MessageMgr) GetFromIDMessage(idMessage uint32) (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ?", idMessage).Find(&result).Error

	return
}

func (obj *_MessageMgr) GetBatchFromIDMessage(idMessages []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message IN (?)", idMessages).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetFromIDCart(idCart uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetBatchFromIDCart(idCarts []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetFromIDCustomer(idCustomer uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetFromIDEmployee(idEmployee uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetFromIDOrder(idOrder uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetFromMessage(message string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetBatchFromMessage(messages []string) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetFromPrivate(private bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private = ?", private).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetBatchFromPrivate(privates []bool) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private IN (?)", privates).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_MessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_MessageMgr) FetchByPrimaryKey(idMessage uint32) (result Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ?", idMessage).Find(&result).Error

	return
}

func (obj *_MessageMgr) FetchIndexByIDCart(idCart uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

func (obj *_MessageMgr) FetchIndexByIDCustomer(idCustomer uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_MessageMgr) FetchIndexByIDEmployee(idEmployee uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_MessageMgr) FetchIndexByMessageOrder(idOrder uint32) (results []*Message, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error

	return
}
