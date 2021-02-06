package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgCustomerThreadMgr struct {
	*_BaseMgr
}

// EgCustomerThreadMgr open func
func EgCustomerThreadMgr(db *gorm.DB) *_EgCustomerThreadMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomerThreadMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomerThreadMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer_thread"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomerThreadMgr) GetTableName() string {
	return "eg_customer_thread"
}

// Get 获取
func (obj *_EgCustomerThreadMgr) Get() (result EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomerThreadMgr) Gets() (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomerThread id_customer_thread获取 
func (obj *_EgCustomerThreadMgr) WithIDCustomerThread(idCustomerThread uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer_thread"] = idCustomerThread })
}

// WithIDShop id_shop获取 
func (obj *_EgCustomerThreadMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取 
func (obj *_EgCustomerThreadMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDContact id_contact获取 
func (obj *_EgCustomerThreadMgr) WithIDContact(idContact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_contact"] = idContact })
}

// WithIDCustomer id_customer获取 
func (obj *_EgCustomerThreadMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDOrder id_order获取 
func (obj *_EgCustomerThreadMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithIDProduct id_product获取 
func (obj *_EgCustomerThreadMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithStatus status获取 
func (obj *_EgCustomerThreadMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithEmail email获取 
func (obj *_EgCustomerThreadMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithToken token获取 
func (obj *_EgCustomerThreadMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

// WithDateAdd date_add获取 
func (obj *_EgCustomerThreadMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgCustomerThreadMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomerThreadMgr) GetByOption(opts ...Option) (result EgCustomerThread, err error) {
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
func (obj *_EgCustomerThreadMgr) GetByOptions(opts ...Option) (results []*EgCustomerThread, err error) {
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


// GetFromIDCustomerThread 通过id_customer_thread获取内容  
func (obj *_EgCustomerThreadMgr)  GetFromIDCustomerThread(idCustomerThread uint32) (result EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread = ?", idCustomerThread).Find(&result).Error
	
	return
}

// GetBatchFromIDCustomerThread 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromIDCustomerThread(idCustomerThreads []uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread IN (?)", idCustomerThreads).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCustomerThreadMgr) GetFromIDShop(idShop uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCustomerThreadMgr) GetFromIDLang(idLang uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDContact 通过id_contact获取内容  
func (obj *_EgCustomerThreadMgr) GetFromIDContact(idContact uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&results).Error
	
	return
}

// GetBatchFromIDContact 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromIDContact(idContacts []uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact IN (?)", idContacts).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgCustomerThreadMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDOrder 通过id_order获取内容  
func (obj *_EgCustomerThreadMgr) GetFromIDOrder(idOrder uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDOrder 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgCustomerThreadMgr) GetFromIDProduct(idProduct uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromStatus 通过status获取内容  
func (obj *_EgCustomerThreadMgr) GetFromStatus(status string) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error
	
	return
}

// GetBatchFromStatus 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromStatus(statuss []string) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error
	
	return
}
 
// GetFromEmail 通过email获取内容  
func (obj *_EgCustomerThreadMgr) GetFromEmail(email string) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error
	
	return
}

// GetBatchFromEmail 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromEmail(emails []string) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error
	
	return
}
 
// GetFromToken 通过token获取内容  
func (obj *_EgCustomerThreadMgr) GetFromToken(token string) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error
	
	return
}

// GetBatchFromToken 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromToken(tokens []string) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgCustomerThreadMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgCustomerThreadMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgCustomerThreadMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomerThreadMgr) FetchByPrimaryKey(idCustomerThread uint32 ) (result EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread = ?", idCustomerThread).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgCustomerThreadMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDLang  获取多个内容
 func (obj *_EgCustomerThreadMgr) FetchIndexByIDLang(idLang uint32 ) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDContact  获取多个内容
 func (obj *_EgCustomerThreadMgr) FetchIndexByIDContact(idContact uint32 ) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCustomer  获取多个内容
 func (obj *_EgCustomerThreadMgr) FetchIndexByIDCustomer(idCustomer uint32 ) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDOrder  获取多个内容
 func (obj *_EgCustomerThreadMgr) FetchIndexByIDOrder(idOrder uint32 ) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDProduct  获取多个内容
 func (obj *_EgCustomerThreadMgr) FetchIndexByIDProduct(idProduct uint32 ) (results []*EgCustomerThread, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}
 

	

