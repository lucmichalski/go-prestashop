package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgTaxLangMgr struct {
	*_BaseMgr
}

// EgTaxLangMgr open func
func EgTaxLangMgr(db *gorm.DB) *_EgTaxLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgTaxLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTaxLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tax_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTaxLangMgr) GetTableName() string {
	return "eg_tax_lang"
}

// Get 获取
func (obj *_EgTaxLangMgr) Get() (result EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTaxLangMgr) Gets() (results []*EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTax id_tax获取 
func (obj *_EgTaxLangMgr) WithIDTax(idTax uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithIDLang id_lang获取 
func (obj *_EgTaxLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgTaxLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgTaxLangMgr) GetByOption(opts ...Option) (result EgTaxLang, err error) {
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
func (obj *_EgTaxLangMgr) GetByOptions(opts ...Option) (results []*EgTaxLang, err error) {
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


// GetFromIDTax 通过id_tax获取内容  
func (obj *_EgTaxLangMgr) GetFromIDTax(idTax uint32) (results []*EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}

// GetBatchFromIDTax 批量唯一主键查找 
func (obj *_EgTaxLangMgr) GetBatchFromIDTax(idTaxs []uint32) (results []*EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgTaxLangMgr) GetFromIDLang(idLang uint32) (results []*EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgTaxLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgTaxLangMgr) GetFromName(name string) (results []*EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgTaxLangMgr) GetBatchFromName(names []string) (results []*EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTaxLangMgr) FetchByPrimaryKey(idTax uint32 ,idLang uint32 ) (result EgTaxLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ? AND id_lang = ?", idTax , idLang).Find(&result).Error
	
	return
}
 

 

	

