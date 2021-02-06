package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentReportMgr struct {
	*_BaseMgr
}

// ProductCommentReportMgr open func
func ProductCommentReportMgr(db *gorm.DB) *_ProductCommentReportMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentReportMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentReportMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment_report"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductCommentReportMgr) GetTableName() string {
	return "ps_product_comment_report"
}

// Get 获取
func (obj *_ProductCommentReportMgr) Get() (result ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductCommentReportMgr) Gets() (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomer id_customer获取
func (obj *_ProductCommentReportMgr) WithIDCustomer(idCustomer int) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDProductComment id_product_comment获取
func (obj *_ProductCommentReportMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_ProductCommentReportMgr) GetFromIDCustomer(idCustomer int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_ProductCommentReportMgr) GetBatchFromIDCustomer(idCustomers []int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDProductComment 通过id_product_comment获取内容
func (obj *_ProductCommentReportMgr) GetFromIDProductComment(idProductComment int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}

// GetBatchFromIDProductComment 批量唯一主键查找
func (obj *_ProductCommentReportMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductCommentReportMgr) FetchByPrimaryKey(idCustomer int, idProductComment int) (result ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_product_comment = ?", idCustomer, idProductComment).Find(&result).Error

	return
}

// FetchIndexByIDXD22C9649ACF38A54  获取多个内容
func (obj *_ProductCommentReportMgr) FetchIndexByIDXD22C9649ACF38A54(idProductComment int) (results []*ProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}
