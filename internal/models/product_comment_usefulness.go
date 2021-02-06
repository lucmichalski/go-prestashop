package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProductCommentUsefulnessMgr struct {
	*_BaseMgr
}

// ProductCommentUsefulnessMgr open func
func ProductCommentUsefulnessMgr(db *gorm.DB) *_ProductCommentUsefulnessMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCommentUsefulnessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCommentUsefulnessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_comment_usefulness"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductCommentUsefulnessMgr) GetTableName() string {
	return "eg_product_comment_usefulness"
}

// Get 获取
func (obj *_ProductCommentUsefulnessMgr) Get() (result ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductCommentUsefulnessMgr) Gets() (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomer id_customer获取
func (obj *_ProductCommentUsefulnessMgr) WithIDCustomer(idCustomer int) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDProductComment id_product_comment获取
func (obj *_ProductCommentUsefulnessMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

// WithUsefulness usefulness获取
func (obj *_ProductCommentUsefulnessMgr) WithUsefulness(usefulness bool) Option {
	return optionFunc(func(o *options) { o.query["usefulness"] = usefulness })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCustomer 通过id_customer获取内容
func (obj *_ProductCommentUsefulnessMgr) GetFromIDCustomer(idCustomer int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error

	return
}

// GetBatchFromIDCustomer 批量唯一主键查找
func (obj *_ProductCommentUsefulnessMgr) GetBatchFromIDCustomer(idCustomers []int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error

	return
}

// GetFromIDProductComment 通过id_product_comment获取内容
func (obj *_ProductCommentUsefulnessMgr) GetFromIDProductComment(idProductComment int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}

// GetBatchFromIDProductComment 批量唯一主键查找
func (obj *_ProductCommentUsefulnessMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error

	return
}

// GetFromUsefulness 通过usefulness获取内容
func (obj *_ProductCommentUsefulnessMgr) GetFromUsefulness(usefulness bool) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usefulness = ?", usefulness).Find(&results).Error

	return
}

// GetBatchFromUsefulness 批量唯一主键查找
func (obj *_ProductCommentUsefulnessMgr) GetBatchFromUsefulness(usefulnesss []bool) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usefulness IN (?)", usefulnesss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductCommentUsefulnessMgr) FetchByPrimaryKey(idCustomer int, idProductComment int) (result ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_product_comment = ?", idCustomer, idProductComment).Find(&result).Error

	return
}

// FetchIndexByIDXE681E112ACF38A54  获取多个内容
func (obj *_ProductCommentUsefulnessMgr) FetchIndexByIDXE681E112ACF38A54(idProductComment int) (results []*ProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error

	return
}
