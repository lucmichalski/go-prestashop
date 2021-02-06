package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgProductCommentMgr struct {
	*_BaseMgr
}

// EgProductCommentMgr open func
func EgProductCommentMgr(db *gorm.DB) *_EgProductCommentMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductCommentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductCommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_comment"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductCommentMgr) GetTableName() string {
	return "eg_product_comment"
}

// Get 获取
func (obj *_EgProductCommentMgr) Get() (result EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductCommentMgr) Gets() (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductComment id_product_comment获取 
func (obj *_EgProductCommentMgr) WithIDProductComment(idProductComment int) Option {
	return optionFunc(func(o *options) { o.query["id_product_comment"] = idProductComment })
}

// WithIDProduct id_product获取 
func (obj *_EgProductCommentMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDCustomer id_customer获取 
func (obj *_EgProductCommentMgr) WithIDCustomer(idCustomer int) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDGuest id_guest获取 
func (obj *_EgProductCommentMgr) WithIDGuest(idGuest int) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

// WithCustomerName customer_name获取 
func (obj *_EgProductCommentMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithTitle title获取 
func (obj *_EgProductCommentMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 
func (obj *_EgProductCommentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithGrade grade获取 
func (obj *_EgProductCommentMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

// WithValidate validate获取 
func (obj *_EgProductCommentMgr) WithValidate(validate bool) Option {
	return optionFunc(func(o *options) { o.query["validate"] = validate })
}

// WithDeleted deleted获取 
func (obj *_EgProductCommentMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithDateAdd date_add获取 
func (obj *_EgProductCommentMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductCommentMgr) GetByOption(opts ...Option) (result EgProductComment, err error) {
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
func (obj *_EgProductCommentMgr) GetByOptions(opts ...Option) (results []*EgProductComment, err error) {
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


// GetFromIDProductComment 通过id_product_comment获取内容  
func (obj *_EgProductCommentMgr)  GetFromIDProductComment(idProductComment int) (result EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&result).Error
	
	return
}

// GetBatchFromIDProductComment 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromIDProductComment(idProductComments []int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment IN (?)", idProductComments).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgProductCommentMgr) GetFromIDProduct(idProduct int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromIDProduct(idProducts []int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgProductCommentMgr) GetFromIDCustomer(idCustomer int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromIDCustomer(idCustomers []int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDGuest 通过id_guest获取内容  
func (obj *_EgProductCommentMgr) GetFromIDGuest(idGuest int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error
	
	return
}

// GetBatchFromIDGuest 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromIDGuest(idGuests []int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error
	
	return
}
 
// GetFromCustomerName 通过customer_name获取内容  
func (obj *_EgProductCommentMgr) GetFromCustomerName(customerName string) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_name = ?", customerName).Find(&results).Error
	
	return
}

// GetBatchFromCustomerName 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromCustomerName(customerNames []string) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_name IN (?)", customerNames).Find(&results).Error
	
	return
}
 
// GetFromTitle 通过title获取内容  
func (obj *_EgProductCommentMgr) GetFromTitle(title string) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error
	
	return
}

// GetBatchFromTitle 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromTitle(titles []string) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error
	
	return
}
 
// GetFromContent 通过content获取内容  
func (obj *_EgProductCommentMgr) GetFromContent(content string) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error
	
	return
}

// GetBatchFromContent 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromContent(contents []string) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error
	
	return
}
 
// GetFromGrade 通过grade获取内容  
func (obj *_EgProductCommentMgr) GetFromGrade(grade int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade = ?", grade).Find(&results).Error
	
	return
}

// GetBatchFromGrade 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromGrade(grades []int) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade IN (?)", grades).Find(&results).Error
	
	return
}
 
// GetFromValidate 通过validate获取内容  
func (obj *_EgProductCommentMgr) GetFromValidate(validate bool) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validate = ?", validate).Find(&results).Error
	
	return
}

// GetBatchFromValidate 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromValidate(validates []bool) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validate IN (?)", validates).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgProductCommentMgr) GetFromDeleted(deleted bool) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgProductCommentMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgProductCommentMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductCommentMgr) FetchByPrimaryKey(idProductComment int ) (result EgProductComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_comment = ?", idProductComment).Find(&result).Error
	
	return
}
 

 

	

