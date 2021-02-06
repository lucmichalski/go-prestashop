package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCountryLangMgr struct {
	*_BaseMgr
}

// EgCountryLangMgr open func
func EgCountryLangMgr(db *gorm.DB) *_EgCountryLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgCountryLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCountryLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_country_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCountryLangMgr) GetTableName() string {
	return "eg_country_lang"
}

// Get 获取
func (obj *_EgCountryLangMgr) Get() (result EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCountryLangMgr) Gets() (results []*EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCountry id_country获取 
func (obj *_EgCountryLangMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDLang id_lang获取 
func (obj *_EgCountryLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgCountryLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgCountryLangMgr) GetByOption(opts ...Option) (result EgCountryLang, err error) {
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
func (obj *_EgCountryLangMgr) GetByOptions(opts ...Option) (results []*EgCountryLang, err error) {
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


// GetFromIDCountry 通过id_country获取内容  
func (obj *_EgCountryLangMgr) GetFromIDCountry(idCountry uint32) (results []*EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}

// GetBatchFromIDCountry 批量唯一主键查找 
func (obj *_EgCountryLangMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCountryLangMgr) GetFromIDLang(idLang uint32) (results []*EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCountryLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCountryLangMgr) GetFromName(name string) (results []*EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCountryLangMgr) GetBatchFromName(names []string) (results []*EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCountryLangMgr) FetchByPrimaryKey(idCountry uint32 ,idLang uint32 ) (result EgCountryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ? AND id_lang = ?", idCountry , idLang).Find(&result).Error
	
	return
}
 

 

	

