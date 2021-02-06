package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgProductCommentReportMgr struct {
	*_BaseMgr
}

// EgProductCommentReportMgr open func
func EgProductCommentReportMgr(db *gorm.DB) *_EgProductCommentReportMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductCommentReportMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductCommentReportMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_comment_report"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductCommentReportMgr) GetTableName() string {
	return "eg_product_comment_report"
}

// Get 获取
func (obj *_EgProductCommentReportMgr) Get() (result EgProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductCommentReportMgr) Gets() (results []*EgProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomer id_customer获取 
func (obj *_EgProductCommentReportMgr) WithIDCustomer(idCustomer int) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDProductComment id_product_comment获取 
func (obj *_EgProductCommentReportMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductCommentReportMgr) GetByOption(opts ...Option) (result EgProductCommentReport, err error) {
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
func (obj *_EgProductCommentReportMgr) GetByOptions(opts ...Option) (results []*EgProductCommentReport, err error) {
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
func (obj *_EgProductCommentReportMgr) GetFromIDCustomer(idCustomer int) (results []*EgProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgProductCommentReportMgr) GetBatchFromIDCustomer(idCustomers []int) (results []*EgProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDProductComment 通过id_product_comment获取内容  
func (obj *_EgProductCommentReportMgr) GetFromIDProductComment(idProductComment int) (results []*EgProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error
	
	return
}

// GetBatchFromIDProductComment 批量唯一主键查找 
func (obj *_EgProductCommentReportMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*EgProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductCommentReportMgr) FetchByPrimaryKey(idCustomer int ,idProductComment int ) (result EgProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_product_comment = ?", idCustomer , idProductComment).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDXD22C9649ACF38A54  获取多个内容
 func (obj *_EgProductCommentReportMgr) FetchIndexByIDXD22C9649ACF38A54(idProductComment int ) (results []*EgProductCommentReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error
	
	return
}
 

	

