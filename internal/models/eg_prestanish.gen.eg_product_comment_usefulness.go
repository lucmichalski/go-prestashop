package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProductCommentUsefulnessMgr struct {
	*_BaseMgr
}

// EgProductCommentUsefulnessMgr open func
func EgProductCommentUsefulnessMgr(db *gorm.DB) *_EgProductCommentUsefulnessMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductCommentUsefulnessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductCommentUsefulnessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_comment_usefulness"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductCommentUsefulnessMgr) GetTableName() string {
	return "eg_product_comment_usefulness"
}

// Get 获取
func (obj *_EgProductCommentUsefulnessMgr) Get() (result EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductCommentUsefulnessMgr) Gets() (results []*EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomer id_customer获取 
func (obj *_EgProductCommentUsefulnessMgr) WithIDCustomer(idCustomer int) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDProductComment id_product_comment获取 
func (obj *_EgProductCommentUsefulnessMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

// WithUsefulness usefulness获取 
func (obj *_EgProductCommentUsefulnessMgr) WithUsefulness(usefulness bool) Option {
	return optionFunc(func(o *options) { o.query["usefulness"] = usefulness })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductCommentUsefulnessMgr) GetByOption(opts ...Option) (result EgProductCommentUsefulness, err error) {
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
func (obj *_EgProductCommentUsefulnessMgr) GetByOptions(opts ...Option) (results []*EgProductCommentUsefulness, err error) {
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
func (obj *_EgProductCommentUsefulnessMgr) GetFromIDCustomer(idCustomer int) (results []*EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgProductCommentUsefulnessMgr) GetBatchFromIDCustomer(idCustomers []int) (results []*EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDProductComment 通过id_product_comment获取内容  
func (obj *_EgProductCommentUsefulnessMgr) GetFromIDProductComment(idProductComment int) (results []*EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error
	
	return
}

// GetBatchFromIDProductComment 批量唯一主键查找 
func (obj *_EgProductCommentUsefulnessMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error
	
	return
}
 
// GetFromUsefulness 通过usefulness获取内容  
func (obj *_EgProductCommentUsefulnessMgr) GetFromUsefulness(usefulness bool) (results []*EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usefulness = ?", usefulness).Find(&results).Error
	
	return
}

// GetBatchFromUsefulness 批量唯一主键查找 
func (obj *_EgProductCommentUsefulnessMgr) GetBatchFromUsefulness(usefulnesss []bool) (results []*EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("usefulness IN (?)", usefulnesss).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductCommentUsefulnessMgr) FetchByPrimaryKey(idCustomer int ,idProductComment int ) (result EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ? AND id_product_comment = ?", idCustomer , idProductComment).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDXE681E112ACF38A54  获取多个内容
 func (obj *_EgProductCommentUsefulnessMgr) FetchIndexByIDXE681E112ACF38A54(idProductComment int ) (results []*EgProductCommentUsefulness, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&results).Error
	
	return
}
 

	

