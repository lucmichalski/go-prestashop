package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgCustomerMessageMgr struct {
	*_BaseMgr
}

// EgCustomerMessageMgr open func
func EgCustomerMessageMgr(db *gorm.DB) *_EgCustomerMessageMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomerMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomerMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer_message"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomerMessageMgr) GetTableName() string {
	return "eg_customer_message"
}

// Get 获取
func (obj *_EgCustomerMessageMgr) Get() (result EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomerMessageMgr) Gets() (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomerMessage id_customer_message获取 
func (obj *_EgCustomerMessageMgr) WithIDCustomerMessage(idCustomerMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer_message"] = idCustomerMessage })
}

// WithIDCustomerThread id_customer_thread获取 
func (obj *_EgCustomerMessageMgr) WithIDCustomerThread(idCustomerThread int) Option {
	return optionFunc(func(o *options) { o.query["id_customer_thread"] = idCustomerThread })
}

// WithIDEmployee id_employee获取 
func (obj *_EgCustomerMessageMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithMessage message获取 
func (obj *_EgCustomerMessageMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithFileName file_name获取 
func (obj *_EgCustomerMessageMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

// WithIPAddress ip_address获取 
func (obj *_EgCustomerMessageMgr) WithIPAddress(ipAddress string) Option {
	return optionFunc(func(o *options) { o.query["ip_address"] = ipAddress })
}

// WithUserAgent user_agent获取 
func (obj *_EgCustomerMessageMgr) WithUserAgent(userAgent string) Option {
	return optionFunc(func(o *options) { o.query["user_agent"] = userAgent })
}

// WithDateAdd date_add获取 
func (obj *_EgCustomerMessageMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgCustomerMessageMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithPrivate private获取 
func (obj *_EgCustomerMessageMgr) WithPrivate(private int8) Option {
	return optionFunc(func(o *options) { o.query["private"] = private })
}

// WithRead read获取 
func (obj *_EgCustomerMessageMgr) WithRead(read bool) Option {
	return optionFunc(func(o *options) { o.query["read"] = read })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomerMessageMgr) GetByOption(opts ...Option) (result EgCustomerMessage, err error) {
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
func (obj *_EgCustomerMessageMgr) GetByOptions(opts ...Option) (results []*EgCustomerMessage, err error) {
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


// GetFromIDCustomerMessage 通过id_customer_message获取内容  
func (obj *_EgCustomerMessageMgr)  GetFromIDCustomerMessage(idCustomerMessage uint32) (result EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message = ?", idCustomerMessage).Find(&result).Error
	
	return
}

// GetBatchFromIDCustomerMessage 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromIDCustomerMessage(idCustomerMessages []uint32) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message IN (?)", idCustomerMessages).Find(&results).Error
	
	return
}
 
// GetFromIDCustomerThread 通过id_customer_thread获取内容  
func (obj *_EgCustomerMessageMgr) GetFromIDCustomerThread(idCustomerThread int) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread = ?", idCustomerThread).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomerThread 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromIDCustomerThread(idCustomerThreads []int) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread IN (?)", idCustomerThreads).Find(&results).Error
	
	return
}
 
// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgCustomerMessageMgr) GetFromIDEmployee(idEmployee uint32) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromMessage 通过message获取内容  
func (obj *_EgCustomerMessageMgr) GetFromMessage(message string) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error
	
	return
}

// GetBatchFromMessage 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromMessage(messages []string) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error
	
	return
}
 
// GetFromFileName 通过file_name获取内容  
func (obj *_EgCustomerMessageMgr) GetFromFileName(fileName string) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name = ?", fileName).Find(&results).Error
	
	return
}

// GetBatchFromFileName 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromFileName(fileNames []string) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("file_name IN (?)", fileNames).Find(&results).Error
	
	return
}
 
// GetFromIPAddress 通过ip_address获取内容  
func (obj *_EgCustomerMessageMgr) GetFromIPAddress(ipAddress string) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address = ?", ipAddress).Find(&results).Error
	
	return
}

// GetBatchFromIPAddress 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromIPAddress(ipAddresss []string) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip_address IN (?)", ipAddresss).Find(&results).Error
	
	return
}
 
// GetFromUserAgent 通过user_agent获取内容  
func (obj *_EgCustomerMessageMgr) GetFromUserAgent(userAgent string) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_agent = ?", userAgent).Find(&results).Error
	
	return
}

// GetBatchFromUserAgent 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromUserAgent(userAgents []string) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("user_agent IN (?)", userAgents).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgCustomerMessageMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgCustomerMessageMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
// GetFromPrivate 通过private获取内容  
func (obj *_EgCustomerMessageMgr) GetFromPrivate(private int8) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private = ?", private).Find(&results).Error
	
	return
}

// GetBatchFromPrivate 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromPrivate(privates []int8) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("private IN (?)", privates).Find(&results).Error
	
	return
}
 
// GetFromRead 通过read获取内容  
func (obj *_EgCustomerMessageMgr) GetFromRead(read bool) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("read = ?", read).Find(&results).Error
	
	return
}

// GetBatchFromRead 批量唯一主键查找 
func (obj *_EgCustomerMessageMgr) GetBatchFromRead(reads []bool) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("read IN (?)", reads).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomerMessageMgr) FetchByPrimaryKey(idCustomerMessage uint32 ) (result EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_message = ?", idCustomerMessage).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCustomerThread  获取多个内容
 func (obj *_EgCustomerMessageMgr) FetchIndexByIDCustomerThread(idCustomerThread int ) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer_thread = ?", idCustomerThread).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDEmployee  获取多个内容
 func (obj *_EgCustomerMessageMgr) FetchIndexByIDEmployee(idEmployee uint32 ) (results []*EgCustomerMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}
 

	

