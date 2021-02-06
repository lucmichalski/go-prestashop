package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgPsgdprConsentMgr struct {
	*_BaseMgr
}

// EgPsgdprConsentMgr open func
func EgPsgdprConsentMgr(db *gorm.DB) *_EgPsgdprConsentMgr {
	if db == nil {
		panic(fmt.Errorf("EgPsgdprConsentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPsgdprConsentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_psgdpr_consent"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPsgdprConsentMgr) GetTableName() string {
	return "eg_psgdpr_consent"
}

// Get 获取
func (obj *_EgPsgdprConsentMgr) Get() (result EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPsgdprConsentMgr) Gets() (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGdprConsent id_gdpr_consent获取 
func (obj *_EgPsgdprConsentMgr) WithIDGdprConsent(idGdprConsent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gdpr_consent"] = idGdprConsent })
}

// WithIDModule id_module获取 
func (obj *_EgPsgdprConsentMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithActive active获取 
func (obj *_EgPsgdprConsentMgr) WithActive(active int) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithError error获取 
func (obj *_EgPsgdprConsentMgr) WithError(error int) Option {
	return optionFunc(func(o *options) { o.query["error"] = error })
}

// WithErrorMessage error_message获取 
func (obj *_EgPsgdprConsentMgr) WithErrorMessage(errorMessage string) Option {
	return optionFunc(func(o *options) { o.query["error_message"] = errorMessage })
}

// WithDateAdd date_add获取 
func (obj *_EgPsgdprConsentMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgPsgdprConsentMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgPsgdprConsentMgr) GetByOption(opts ...Option) (result EgPsgdprConsent, err error) {
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
func (obj *_EgPsgdprConsentMgr) GetByOptions(opts ...Option) (results []*EgPsgdprConsent, err error) {
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


// GetFromIDGdprConsent 通过id_gdpr_consent获取内容  
func (obj *_EgPsgdprConsentMgr) GetFromIDGdprConsent(idGdprConsent uint32) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ?", idGdprConsent).Find(&results).Error
	
	return
}

// GetBatchFromIDGdprConsent 批量唯一主键查找 
func (obj *_EgPsgdprConsentMgr) GetBatchFromIDGdprConsent(idGdprConsents []uint32) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent IN (?)", idGdprConsents).Find(&results).Error
	
	return
}
 
// GetFromIDModule 通过id_module获取内容  
func (obj *_EgPsgdprConsentMgr) GetFromIDModule(idModule uint32) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}

// GetBatchFromIDModule 批量唯一主键查找 
func (obj *_EgPsgdprConsentMgr) GetBatchFromIDModule(idModules []uint32) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgPsgdprConsentMgr) GetFromActive(active int) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgPsgdprConsentMgr) GetBatchFromActive(actives []int) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
// GetFromError 通过error获取内容  
func (obj *_EgPsgdprConsentMgr) GetFromError(error int) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error = ?", error).Find(&results).Error
	
	return
}

// GetBatchFromError 批量唯一主键查找 
func (obj *_EgPsgdprConsentMgr) GetBatchFromError(errors []int) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error IN (?)", errors).Find(&results).Error
	
	return
}
 
// GetFromErrorMessage 通过error_message获取内容  
func (obj *_EgPsgdprConsentMgr) GetFromErrorMessage(errorMessage string) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_message = ?", errorMessage).Find(&results).Error
	
	return
}

// GetBatchFromErrorMessage 批量唯一主键查找 
func (obj *_EgPsgdprConsentMgr) GetBatchFromErrorMessage(errorMessages []string) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("error_message IN (?)", errorMessages).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgPsgdprConsentMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgPsgdprConsentMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgPsgdprConsentMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgPsgdprConsentMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPsgdprConsentMgr) FetchByPrimaryKey(idGdprConsent uint32 ,idModule uint32 ) (result EgPsgdprConsent, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ? AND id_module = ?", idGdprConsent , idModule).Find(&result).Error
	
	return
}
 

 

	

