package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgConfigurationLangMgr struct {
	*_BaseMgr
}

// EgConfigurationLangMgr open func
func EgConfigurationLangMgr(db *gorm.DB) *_EgConfigurationLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgConfigurationLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgConfigurationLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_configuration_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgConfigurationLangMgr) GetTableName() string {
	return "eg_configuration_lang"
}

// Get 获取
func (obj *_EgConfigurationLangMgr) Get() (result EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgConfigurationLangMgr) Gets() (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConfiguration id_configuration获取 
func (obj *_EgConfigurationLangMgr) WithIDConfiguration(idConfiguration uint32) Option {
	return optionFunc(func(o *options) { o.query["id_configuration"] = idConfiguration })
}

// WithIDLang id_lang获取 
func (obj *_EgConfigurationLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithValue value获取 
func (obj *_EgConfigurationLangMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithDateUpd date_upd获取 
func (obj *_EgConfigurationLangMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgConfigurationLangMgr) GetByOption(opts ...Option) (result EgConfigurationLang, err error) {
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
func (obj *_EgConfigurationLangMgr) GetByOptions(opts ...Option) (results []*EgConfigurationLang, err error) {
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


// GetFromIDConfiguration 通过id_configuration获取内容  
func (obj *_EgConfigurationLangMgr) GetFromIDConfiguration(idConfiguration uint32) (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration = ?", idConfiguration).Find(&results).Error
	
	return
}

// GetBatchFromIDConfiguration 批量唯一主键查找 
func (obj *_EgConfigurationLangMgr) GetBatchFromIDConfiguration(idConfigurations []uint32) (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration IN (?)", idConfigurations).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgConfigurationLangMgr) GetFromIDLang(idLang uint32) (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgConfigurationLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_EgConfigurationLangMgr) GetFromValue(value string) (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量唯一主键查找 
func (obj *_EgConfigurationLangMgr) GetBatchFromValue(values []string) (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgConfigurationLangMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgConfigurationLangMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgConfigurationLangMgr) FetchByPrimaryKey(idConfiguration uint32 ,idLang uint32 ) (result EgConfigurationLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration = ? AND id_lang = ?", idConfiguration , idLang).Find(&result).Error
	
	return
}
 

 

	

