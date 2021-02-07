package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentReportMgr struct {
	*_BaseMgr
}

func ProductCommentReportMgr(db *gorm.DB) *_ProductCommentReportMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentReportMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentReportMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment_report"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductCommentReportMgr) GetTableName() string {
	return "ps_product_comment_report"
}

func (obj *_ProductCommentReportMgr) Get() (result ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductCommentReportMgr) Gets() (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ProductCommentReportMgr) WithIDCustomer(idCustomer int) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_ProductCommentReportMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

func (obj *_ProductCommentReportMgr) GetByOption(opts ...Option) (result ProductCommentReport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductCommentReportMgr) GetByOptions(opts ...Option) (results []*ProductCommentReport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ProductCommentReportMgr) GetFromIDCustomer(idCustomer int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_ProductCommentReportMgr) GetBatchFromIDCustomer(idCustomers []int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_ProductCommentReportMgr) GetFromIDProductComment(idProductComment int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}

func (obj *_ProductCommentReportMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error

	return
}

func (obj *_ProductCommentReportMgr) FetchByPrimaryKey(idCustomer int, idProductComment int) (result ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_product_comment = ?", idCustomer, idProductComment).Find(&result).Error

	return
}

func (obj *_ProductCommentReportMgr) FetchIndexByIDXD22C9649ACF38A54(idProductComment int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}
