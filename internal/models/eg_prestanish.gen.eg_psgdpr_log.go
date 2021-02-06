package	model	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _EgPsgdprLogMgr struct {
	*_BaseMgr
}

// EgPsgdprLogMgr open func
func EgPsgdprLogMgr(db *gorm.DB) *_EgPsgdprLogMgr {
	if db == nil {
		panic(fmt.Errorf("EgPsgdprLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPsgdprLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_psgdpr_log"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPsgdprLogMgr) GetTableName() string {
	return "eg_psgdpr_log"
}

// Get 获取
func (obj *_EgPsgdprLogMgr) Get() (result EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPsgdprLogMgr) Gets() (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGdprLog id_gdpr_log获取 
func (obj *_EgPsgdprLogMgr) WithIDGdprLog(idGdprLog uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gdpr_log"] = idGdprLog })
}

// WithIDCustomer id_customer获取 
func (obj *_EgPsgdprLogMgr) WithIDCustomer(idCustomer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customer"] = idCustomer })
}

// WithIDGuest id_guest获取 
func (obj *_EgPsgdprLogMgr) WithIDGuest(idGuest uint32) Option {
	return optionFunc(func(o *options) { o.query["id_guest"] = idGuest })
}

// WithClientName client_name获取 
func (obj *_EgPsgdprLogMgr) WithClientName(clientName string) Option {
	return optionFunc(func(o *options) { o.query["client_name"] = clientName })
}

// WithIDModule id_module获取 
func (obj *_EgPsgdprLogMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithRequestType request_type获取 
func (obj *_EgPsgdprLogMgr) WithRequestType(requestType int) Option {
	return optionFunc(func(o *options) { o.query["request_type"] = requestType })
}

// WithDateAdd date_add获取 
func (obj *_EgPsgdprLogMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgPsgdprLogMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgPsgdprLogMgr) GetByOption(opts ...Option) (result EgPsgdprLog, err error) {
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
func (obj *_EgPsgdprLogMgr) GetByOptions(opts ...Option) (results []*EgPsgdprLog, err error) {
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


// GetFromIDGdprLog 通过id_gdpr_log获取内容  
func (obj *_EgPsgdprLogMgr)  GetFromIDGdprLog(idGdprLog uint32) (result EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log = ?", idGdprLog).Find(&result).Error
	
	return
}

// GetBatchFromIDGdprLog 批量唯一主键查找 
func (obj *_EgPsgdprLogMgr) GetBatchFromIDGdprLog(idGdprLogs []uint32) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log IN (?)", idGdprLogs).Find(&results).Error
	
	return
}
 
// GetFromIDCustomer 通过id_customer获取内容  
func (obj *_EgPsgdprLogMgr) GetFromIDCustomer(idCustomer uint32) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer = ?", idCustomer).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomer 批量唯一主键查找 
func (obj *_EgPsgdprLogMgr) GetBatchFromIDCustomer(idCustomers []uint32) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customer IN (?)", idCustomers).Find(&results).Error
	
	return
}
 
// GetFromIDGuest 通过id_guest获取内容  
func (obj *_EgPsgdprLogMgr) GetFromIDGuest(idGuest uint32) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest = ?", idGuest).Find(&results).Error
	
	return
}

// GetBatchFromIDGuest 批量唯一主键查找 
func (obj *_EgPsgdprLogMgr) GetBatchFromIDGuest(idGuests []uint32) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_guest IN (?)", idGuests).Find(&results).Error
	
	return
}
 
// GetFromClientName 通过client_name获取内容  
func (obj *_EgPsgdprLogMgr) GetFromClientName(clientName string) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("client_name = ?", clientName).Find(&results).Error
	
	return
}

// GetBatchFromClientName 批量唯一主键查找 
func (obj *_EgPsgdprLogMgr) GetBatchFromClientName(clientNames []string) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("client_name IN (?)", clientNames).Find(&results).Error
	
	return
}
 
// GetFromIDModule 通过id_module获取内容  
func (obj *_EgPsgdprLogMgr) GetFromIDModule(idModule uint32) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}

// GetBatchFromIDModule 批量唯一主键查找 
func (obj *_EgPsgdprLogMgr) GetBatchFromIDModule(idModules []uint32) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error
	
	return
}
 
// GetFromRequestType 通过request_type获取内容  
func (obj *_EgPsgdprLogMgr) GetFromRequestType(requestType int) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_type = ?", requestType).Find(&results).Error
	
	return
}

// GetBatchFromRequestType 批量唯一主键查找 
func (obj *_EgPsgdprLogMgr) GetBatchFromRequestType(requestTypes []int) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_type IN (?)", requestTypes).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgPsgdprLogMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgPsgdprLogMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgPsgdprLogMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgPsgdprLogMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPsgdprLogMgr) FetchByPrimaryKey(idGdprLog uint32 ) (result EgPsgdprLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_log = ?", idGdprLog).Find(&result).Error
	
	return
}
 

 

	

