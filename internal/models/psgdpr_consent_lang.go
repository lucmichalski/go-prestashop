package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgPsgdprConsentLangMgr struct {
	*_BaseMgr
}

// EgPsgdprConsentLangMgr open func
func EgPsgdprConsentLangMgr(db *gorm.DB) *_EgPsgdprConsentLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgPsgdprConsentLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPsgdprConsentLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_psgdpr_consent_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPsgdprConsentLangMgr) GetTableName() string {
	return "eg_psgdpr_consent_lang"
}

// Get 获取
func (obj *_EgPsgdprConsentLangMgr) Get() (result EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPsgdprConsentLangMgr) Gets() (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGdprConsent id_gdpr_consent获取 
func (obj *_EgPsgdprConsentLangMgr) WithIDGdprConsent(idGdprConsent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gdpr_consent"] = idGdprConsent })
}

// WithIDLang id_lang获取 
func (obj *_EgPsgdprConsentLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithMessage message获取 
func (obj *_EgPsgdprConsentLangMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// WithIDShop id_shop获取 
func (obj *_EgPsgdprConsentLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgPsgdprConsentLangMgr) GetByOption(opts ...Option) (result EgPsgdprConsentLang, err error) {
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
func (obj *_EgPsgdprConsentLangMgr) GetByOptions(opts ...Option) (results []*EgPsgdprConsentLang, err error) {
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
func (obj *_EgPsgdprConsentLangMgr) GetFromIDGdprConsent(idGdprConsent uint32) (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ?", idGdprConsent).Find(&results).Error
	
	return
}

// GetBatchFromIDGdprConsent 批量唯一主键查找 
func (obj *_EgPsgdprConsentLangMgr) GetBatchFromIDGdprConsent(idGdprConsents []uint32) (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent IN (?)", idGdprConsents).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgPsgdprConsentLangMgr) GetFromIDLang(idLang uint32) (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgPsgdprConsentLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromMessage 通过message获取内容  
func (obj *_EgPsgdprConsentLangMgr) GetFromMessage(message string) (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error
	
	return
}

// GetBatchFromMessage 批量唯一主键查找 
func (obj *_EgPsgdprConsentLangMgr) GetBatchFromMessage(messages []string) (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgPsgdprConsentLangMgr) GetFromIDShop(idShop uint32) (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgPsgdprConsentLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPsgdprConsentLangMgr) FetchByPrimaryKey(idGdprConsent uint32 ,idLang uint32 ,idShop uint32 ) (result EgPsgdprConsentLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gdpr_consent = ? AND id_lang = ? AND id_shop = ?", idGdprConsent , idLang , idShop).Find(&result).Error
	
	return
}
 

 

	

