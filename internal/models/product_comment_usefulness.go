package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentUsefulnessMgr struct {
	*_BaseMgr
}

func ProductCommentUsefulnessMgr(db *gorm.DB) *_ProductCommentUsefulnessMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentUsefulnessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentUsefulnessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_comment_usefulness"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductCommentUsefulnessMgr) GetTableName() string {
	return "ps_product_comment_usefulness"
}

func (obj *_ProductCommentUsefulnessMgr) Get() (result ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) Gets() (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) WithIDCustomer(idCustomer int) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

func (obj *_ProductCommentUsefulnessMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

func (obj *_ProductCommentUsefulnessMgr) WithUsefulness(usefulness bool) Option {
	return optionFunc(func(o *options) { o.query["usefulness"] = usefulness })
}

func (obj *_ProductCommentUsefulnessMgr) GetByOption(opts ...Option) (result ProductCommentUsefulness, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) GetByOptions(opts ...Option) (results []*ProductCommentUsefulness, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) GetFromIDCustomer(idCustomer int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) GetBatchFromIDCustomer(idCustomers []int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) GetFromIDProductComment(idProductComment int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) GetFromUsefulness(usefulness bool) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usefulness = ?", usefulness).Find(&results).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) GetBatchFromUsefulness(usefulnesss []bool) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usefulness IN (?)", usefulnesss).Find(&results).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) FetchByPrimaryKey(idCustomer int, idProductComment int) (result ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_product_comment = ?", idCustomer, idProductComment).Find(&result).Error

	return
}

func (obj *_ProductCommentUsefulnessMgr) FetchIndexByIDXE681E112ACF38A54(idProductComment int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}
