package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgLogMgr struct {
	*_BaseMgr
}

// EgLogMgr open func
func EgLogMgr(db *gorm.DB) *_EgLogMgr {
	if db == nil {
		panic(fmt.Errorf("EgLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_log"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLogMgr) GetTableName() string {
	return "eg_log"
}

// Get 获取
func (obj *_EgLogMgr) Get() (result EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLogMgr) Gets() (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLog id_log获取 
func (obj *_EgLogMgr) WithIDLog(idLog uint32) Option {
	return optionFunc(func(o *options) { o.query["id_log"] = idLog })
}

// WithSeverity severity获取 
func (obj *_EgLogMgr) WithSeverity(severity bool) Option {
	return optionFunc(func(o *options) { o.query["severity"] = severity })
}

// WithErrorCode error_code获取 
func (obj *_EgLogMgr) WithErrorCode(errorCode int) Option {
	return optionFunc(func(o *options) { o.query["error_code"] = errorCode })
}

// WithMessage message获取 
func (obj *_EgLogMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithObjectType object_type获取 
func (obj *_EgLogMgr) WithObjectType(objectType string) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithObjectID object_id获取 
func (obj *_EgLogMgr) WithObjectID(objectID uint32) Option {
	return optionFunc(func(o *options) { o.query["object_id"] = objectID })
}

// WithIDEmployee id_employee获取 
func (obj *_EgLogMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithDateAdd date_add获取 
func (obj *_EgLogMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgLogMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgLogMgr) GetByOption(opts ...Option) (result EgLog, err error) {
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
func (obj *_EgLogMgr) GetByOptions(opts ...Option) (results []*EgLog, err error) {
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


// GetFromIDLog 通过id_log获取内容  
func (obj *_EgLogMgr)  GetFromIDLog(idLog uint32) (result EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_log = ?", idLog).Find(&result).Error
	
	return
}

// GetBatchFromIDLog 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromIDLog(idLogs []uint32) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_log IN (?)", idLogs).Find(&results).Error
	
	return
}
 
// GetFromSeverity 通过severity获取内容  
func (obj *_EgLogMgr) GetFromSeverity(severity bool) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("severity = ?", severity).Find(&results).Error
	
	return
}

// GetBatchFromSeverity 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromSeverity(severitys []bool) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("severity IN (?)", severitys).Find(&results).Error
	
	return
}
 
// GetFromErrorCode 通过error_code获取内容  
func (obj *_EgLogMgr) GetFromErrorCode(errorCode int) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_code = ?", errorCode).Find(&results).Error
	
	return
}

// GetBatchFromErrorCode 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromErrorCode(errorCodes []int) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_code IN (?)", errorCodes).Find(&results).Error
	
	return
}
 
// GetFromMessage 通过message获取内容  
func (obj *_EgLogMgr) GetFromMessage(message string) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error
	
	return
}

// GetBatchFromMessage 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromMessage(messages []string) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error
	
	return
}
 
// GetFromObjectType 通过object_type获取内容  
func (obj *_EgLogMgr) GetFromObjectType(objectType string) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_type = ?", objectType).Find(&results).Error
	
	return
}

// GetBatchFromObjectType 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromObjectType(objectTypes []string) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_type IN (?)", objectTypes).Find(&results).Error
	
	return
}
 
// GetFromObjectID 通过object_id获取内容  
func (obj *_EgLogMgr) GetFromObjectID(objectID uint32) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_id = ?", objectID).Find(&results).Error
	
	return
}

// GetBatchFromObjectID 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromObjectID(objectIDs []uint32) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("object_id IN (?)", objectIDs).Find(&results).Error
	
	return
}
 
// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgLogMgr) GetFromIDEmployee(idEmployee uint32) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgLogMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgLogMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgLogMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLogMgr) FetchByPrimaryKey(idLog uint32 ) (result EgLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_log = ?", idLog).Find(&result).Error
	
	return
}
 

 

	

