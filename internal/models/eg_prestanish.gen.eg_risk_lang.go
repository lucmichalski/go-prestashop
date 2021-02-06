package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgRiskLangMgr struct {
	*_BaseMgr
}

// EgRiskLangMgr open func
func EgRiskLangMgr(db *gorm.DB) *_EgRiskLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgRiskLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgRiskLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_risk_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgRiskLangMgr) GetTableName() string {
	return "eg_risk_lang"
}

// Get 获取
func (obj *_EgRiskLangMgr) Get() (result EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgRiskLangMgr) Gets() (results []*EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDRisk id_risk获取 
func (obj *_EgRiskLangMgr) WithIDRisk(idRisk uint32) Option {
	return optionFunc(func(o *options) { o.query["id_risk"] = idRisk })
}

// WithIDLang id_lang获取 
func (obj *_EgRiskLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgRiskLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgRiskLangMgr) GetByOption(opts ...Option) (result EgRiskLang, err error) {
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
func (obj *_EgRiskLangMgr) GetByOptions(opts ...Option) (results []*EgRiskLang, err error) {
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


// GetFromIDRisk 通过id_risk获取内容  
func (obj *_EgRiskLangMgr) GetFromIDRisk(idRisk uint32) (results []*EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&results).Error
	
	return
}

// GetBatchFromIDRisk 批量唯一主键查找 
func (obj *_EgRiskLangMgr) GetBatchFromIDRisk(idRisks []uint32) (results []*EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk IN (?)", idRisks).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgRiskLangMgr) GetFromIDLang(idLang uint32) (results []*EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgRiskLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgRiskLangMgr) GetFromName(name string) (results []*EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgRiskLangMgr) GetBatchFromName(names []string) (results []*EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgRiskLangMgr) FetchByPrimaryKey(idRisk uint32 ,idLang uint32 ) (result EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ? AND id_lang = ?", idRisk , idLang).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDRisk  获取多个内容
 func (obj *_EgRiskLangMgr) FetchIndexByIDRisk(idRisk uint32 ) (results []*EgRiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&results).Error
	
	return
}
 

	

