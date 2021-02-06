package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCustomerSessionMgr struct {
	*_BaseMgr
}

// EgCustomerSessionMgr open func
func EgCustomerSessionMgr(db *gorm.DB) *_EgCustomerSessionMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomerSessionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomerSessionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer_session"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomerSessionMgr) GetTableName() string {
	return "eg_customer_session"
}

// Get 获取
func (obj *_EgCustomerSessionMgr) Get() (result EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomerSessionMgr) Gets() (results []*EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomerSession id_customer_session获取 
func (obj *_EgCustomerSessionMgr) WithIDCustomerSession(idCustomerSession uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer_session"] = idCustomerSession })
}

// WithIDCustomer id_customer获取 
func (obj *_EgCustomerSessionMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithToken token获取 
func (obj *_EgCustomerSessionMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomerSessionMgr) GetByOption(opts ...Option) (result EgCustomerSession, err error) {
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
func (obj *_EgCustomerSessionMgr) GetByOptions(opts ...Option) (results []*EgCustomerSession, err error) {
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


// GetFromIDCustomerSession 通过id_customer_session获取内容  
func (obj *_EgCustomerSessionMgr)  GetFromIDCustomerSession(idCustomerSession uint32) (result EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session = ?", idCustomerSession).Find(&result).Error
	
	return
}

// GetBatchFromIDCustomerSession 批量唯一主键查找 
func (obj *_EgCustomerSessionMgr) GetBatchFromIDCustomerSession(idCustomerSessions []uint32) (results []*EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session IN (?)", idCustomerSessions).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgCustomerSessionMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgCustomerSessionMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromToken 通过token获取内容  
func (obj *_EgCustomerSessionMgr) GetFromToken(token string) (results []*EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error
	
	return
}

// GetBatchFromToken 批量唯一主键查找 
func (obj *_EgCustomerSessionMgr) GetBatchFromToken(tokens []string) (results []*EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomerSessionMgr) FetchByPrimaryKey(idCustomerSession uint32 ) (result EgCustomerSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_session = ?", idCustomerSession).Find(&result).Error
	
	return
}
 

 

	

