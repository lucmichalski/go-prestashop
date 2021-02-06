package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgMessageMgr struct {
	*_BaseMgr
}

// EgMessageMgr open func
func EgMessageMgr(db *gorm.DB) *_EgMessageMgr {
	if db == nil {
		panic(fmt.Errorf("EgMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_message"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgMessageMgr) GetTableName() string {
	return "eg_message"
}

// Get 获取
func (obj *_EgMessageMgr) Get() (result EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgMessageMgr) Gets() (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMessage id_message获取 
func (obj *_EgMessageMgr) WithIDMessage(idMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_message"] = idMessage })
}

// WithIDCart id_cart获取 
func (obj *_EgMessageMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDCustomer id_customer获取 
func (obj *_EgMessageMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDEmployee id_employee获取 
func (obj *_EgMessageMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithIDOrder id_order获取 
func (obj *_EgMessageMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithMessage message获取 
func (obj *_EgMessageMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithPrivate private获取 
func (obj *_EgMessageMgr) WithPrivate(private bool) Option {
	return optionFunc(func(o *options) { o.query["private"] = private })
}

// WithDateAdd date_add获取 
func (obj *_EgMessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}


// GetByOption 功能选项模式获取
func (obj *_EgMessageMgr) GetByOption(opts ...Option) (result EgMessage, err error) {
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
func (obj *_EgMessageMgr) GetByOptions(opts ...Option) (results []*EgMessage, err error) {
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
func (obj *_EgMessageMgr)  GetFromIDMessage(idMessage uint32) (result EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ?", idMessage).Find(&result).Error
	
	return
}

// GetBatchFromIDMessage 批量唯一主键查找 
func (obj *_EgMessageMgr) GetBatchFromIDMessage(idMessages []uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message IN (?)", idMessages).Find(&results).Error
	
	return
}
 
// GetFromIDCart 通过id_cart获取内容  
func (obj *_EgMessageMgr) GetFromIDCart(idCart uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error
	
	return
}

// GetBatchFromIDCart 批量唯一主键查找 
func (obj *_EgMessageMgr) GetBatchFromIDCart(idCarts []uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgMessageMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgMessageMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgMessageMgr) GetFromIDEmployee(idEmployee uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgMessageMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromIDOrder 通过id_order获取内容  
func (obj *_EgMessageMgr) GetFromIDOrder(idOrder uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDOrder 批量唯一主键查找 
func (obj *_EgMessageMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error
	
	return
}
 
// GetFromMessage 通过message获取内容  
func (obj *_EgMessageMgr) GetFromMessage(message string) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error
	
	return
}

// GetBatchFromMessage 批量唯一主键查找 
func (obj *_EgMessageMgr) GetBatchFromMessage(messages []string) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error
	
	return
}
 
// GetFromPrivate 通过private获取内容  
func (obj *_EgMessageMgr) GetFromPrivate(private bool) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private = ?", private).Find(&results).Error
	
	return
}

// GetBatchFromPrivate 批量唯一主键查找 
func (obj *_EgMessageMgr) GetBatchFromPrivate(privates []bool) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private IN (?)", privates).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgMessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgMessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgMessageMgr) FetchByPrimaryKey(idMessage uint32 ) (result EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ?", idMessage).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCart  获取多个内容
 func (obj *_EgMessageMgr) FetchIndexByIDCart(idCart uint32 ) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCustomer  获取多个内容
 func (obj *_EgMessageMgr) FetchIndexByIDCustomer(idCustomer uint32 ) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDEmployee  获取多个内容
 func (obj *_EgMessageMgr) FetchIndexByIDEmployee(idEmployee uint32 ) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}
 
 // FetchIndexByMessageOrder  获取多个内容
 func (obj *_EgMessageMgr) FetchIndexByMessageOrder(idOrder uint32 ) (results []*EgMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}
 

	

