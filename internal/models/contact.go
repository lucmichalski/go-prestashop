package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgContactMgr struct {
	*_BaseMgr
}

// EgContactMgr open func
func EgContactMgr(db *gorm.DB) *_EgContactMgr {
	if db == nil {
		panic(fmt.Errorf("EgContactMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgContactMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_contact"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgContactMgr) GetTableName() string {
	return "eg_contact"
}

// Get 获取
func (obj *_EgContactMgr) Get() (result EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgContactMgr) Gets() (results []*EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDContact id_contact获取 
func (obj *_EgContactMgr) WithIDContact(idContact uint32) Option {
	return optionFunc(func(o *options) { o.query["id_contact"] = idContact })
}

// WithEmail email获取 
func (obj *_EgContactMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithCustomerService customer_service获取 
func (obj *_EgContactMgr) WithCustomerService(customerService bool) Option {
	return optionFunc(func(o *options) { o.query["customer_service"] = customerService })
}

// WithPosition position获取 
func (obj *_EgContactMgr) WithPosition(position uint8) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}


// GetByOption 功能选项模式获取
func (obj *_EgContactMgr) GetByOption(opts ...Option) (result EgContact, err error) {
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
func (obj *_EgContactMgr) GetByOptions(opts ...Option) (results []*EgContact, err error) {
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


// GetFromIDContact 通过id_contact获取内容  
func (obj *_EgContactMgr)  GetFromIDContact(idContact uint32) (result EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&result).Error
	
	return
}

// GetBatchFromIDContact 批量唯一主键查找 
func (obj *_EgContactMgr) GetBatchFromIDContact(idContacts []uint32) (results []*EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact IN (?)", idContacts).Find(&results).Error
	
	return
}
 
// GetFromEmail 通过email获取内容  
func (obj *_EgContactMgr) GetFromEmail(email string) (results []*EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error
	
	return
}

// GetBatchFromEmail 批量唯一主键查找 
func (obj *_EgContactMgr) GetBatchFromEmail(emails []string) (results []*EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error
	
	return
}
 
// GetFromCustomerService 通过customer_service获取内容  
func (obj *_EgContactMgr) GetFromCustomerService(customerService bool) (results []*EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_service = ?", customerService).Find(&results).Error
	
	return
}

// GetBatchFromCustomerService 批量唯一主键查找 
func (obj *_EgContactMgr) GetBatchFromCustomerService(customerServices []bool) (results []*EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("customer_service IN (?)", customerServices).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgContactMgr) GetFromPosition(position uint8) (results []*EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgContactMgr) GetBatchFromPosition(positions []uint8) (results []*EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgContactMgr) FetchByPrimaryKey(idContact uint32 ) (result EgContact, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_contact = ?", idContact).Find(&result).Error
	
	return
}
 

 

	

