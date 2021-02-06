package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ProductCommentMgr struct {
	*_BaseMgr
}

func ProductCommentMgr(db *gorm.DB) *_ProductCommentMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductCommentMgr) GetTableName() string {
	return "ps_product_comment"
}

func (obj *_ProductCommentMgr) Get() (result ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductCommentMgr) Gets() (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductCommentMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

func (obj *_ProductCommentMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductCommentMgr) WithIDCustomer(idCustomer int) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_ProductCommentMgr) WithIDGuest(idGuest int) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

func (obj *_ProductCommentMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

func (obj *_ProductCommentMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

func (obj *_ProductCommentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

func (obj *_ProductCommentMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

func (obj *_ProductCommentMgr) WithValidate(validate bool) Option {
	return optionFunc(func(o *options) { o.query["validate"] = validate })
}

func (obj *_ProductCommentMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_ProductCommentMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ProductCommentMgr) GetByOption(opts ...Option) (result ProductComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductCommentMgr) GetByOptions(opts ...Option) (results []*ProductComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProductCommentMgr) GetFromIDProductComment(idProductComment int) (result ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&result).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromIDProduct(idProduct int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromIDProduct(idProducts []int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromIDCustomer(idCustomer int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromIDCustomer(idCustomers []int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromIDGuest(idGuest int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromIDGuest(idGuests []int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromCustomerName(customerName string) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_name = ?", customerName).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromCustomerName(customerNames []string) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_name IN (?)", customerNames).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromTitle(title string) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromTitle(titles []string) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromContent(content string) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromContent(contents []string) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromGrade(grade int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade = ?", grade).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromGrade(grades []int) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade IN (?)", grades).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromValidate(validate bool) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validate = ?", validate).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromValidate(validates []bool) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validate IN (?)", validates).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromDeleted(deleted bool) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromDeleted(deleteds []bool) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetFromDateAdd(dateAdd time.Time) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ProductCommentMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_ProductCommentMgr) FetchByPrimaryKey(idProductComment int) (result ProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&result).Error

	return
}
