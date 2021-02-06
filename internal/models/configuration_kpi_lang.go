package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgConfigurationKpiLangMgr struct {
	*_BaseMgr
}

// EgConfigurationKpiLangMgr open func
func EgConfigurationKpiLangMgr(db *gorm.DB) *_EgConfigurationKpiLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgConfigurationKpiLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgConfigurationKpiLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_configuration_kpi_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgConfigurationKpiLangMgr) GetTableName() string {
	return "eg_configuration_kpi_lang"
}

// Get 获取
func (obj *_EgConfigurationKpiLangMgr) Get() (result EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgConfigurationKpiLangMgr) Gets() (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDConfigurationKpi id_configuration_kpi获取 
func (obj *_EgConfigurationKpiLangMgr) WithIDConfigurationKpi(idConfigurationKpi uint32) Option {
	return optionFunc(func(o *options) { o.query["id_configuration_kpi"] = idConfigurationKpi })
}

// WithIDLang id_lang获取 
func (obj *_EgConfigurationKpiLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithValue value获取 
func (obj *_EgConfigurationKpiLangMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithDateUpd date_upd获取 
func (obj *_EgConfigurationKpiLangMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgConfigurationKpiLangMgr) GetByOption(opts ...Option) (result EgConfigurationKpiLang, err error) {
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
func (obj *_EgConfigurationKpiLangMgr) GetByOptions(opts ...Option) (results []*EgConfigurationKpiLang, err error) {
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


// GetFromIDConfigurationKpi 通过id_configuration_kpi获取内容  
func (obj *_EgConfigurationKpiLangMgr) GetFromIDConfigurationKpi(idConfigurationKpi uint32) (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi = ?", idConfigurationKpi).Find(&results).Error
	
	return
}

// GetBatchFromIDConfigurationKpi 批量唯一主键查找 
func (obj *_EgConfigurationKpiLangMgr) GetBatchFromIDConfigurationKpi(idConfigurationKpis []uint32) (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi IN (?)", idConfigurationKpis).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgConfigurationKpiLangMgr) GetFromIDLang(idLang uint32) (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgConfigurationKpiLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_EgConfigurationKpiLangMgr) GetFromValue(value string) (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量唯一主键查找 
func (obj *_EgConfigurationKpiLangMgr) GetBatchFromValue(values []string) (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgConfigurationKpiLangMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgConfigurationKpiLangMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgConfigurationKpiLangMgr) FetchByPrimaryKey(idConfigurationKpi uint32 ,idLang uint32 ) (result EgConfigurationKpiLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_configuration_kpi = ? AND id_lang = ?", idConfigurationKpi , idLang).Find(&result).Error
	
	return
}
 

 

	

