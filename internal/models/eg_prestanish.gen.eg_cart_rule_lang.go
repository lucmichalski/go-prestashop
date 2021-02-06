package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgCartRuleLangMgr struct {
	*_BaseMgr
}

// EgCartRuleLangMgr open func
func EgCartRuleLangMgr(db *gorm.DB) *_EgCartRuleLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgCartRuleLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCartRuleLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cart_rule_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCartRuleLangMgr) GetTableName() string {
	return "eg_cart_rule_lang"
}

// Get 获取
func (obj *_EgCartRuleLangMgr) Get() (result EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCartRuleLangMgr) Gets() (results []*EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCartRule id_cart_rule获取 
func (obj *_EgCartRuleLangMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithIDLang id_lang获取 
func (obj *_EgCartRuleLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgCartRuleLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgCartRuleLangMgr) GetByOption(opts ...Option) (result EgCartRuleLang, err error) {
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
func (obj *_EgCartRuleLangMgr) GetByOptions(opts ...Option) (results []*EgCartRuleLang, err error) {
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


// GetFromIDCartRule 通过id_cart_rule获取内容  
func (obj *_EgCartRuleLangMgr) GetFromIDCartRule(idCartRule uint32) (results []*EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error
	
	return
}

// GetBatchFromIDCartRule 批量唯一主键查找 
func (obj *_EgCartRuleLangMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCartRuleLangMgr) GetFromIDLang(idLang uint32) (results []*EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCartRuleLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCartRuleLangMgr) GetFromName(name string) (results []*EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCartRuleLangMgr) GetBatchFromName(names []string) (results []*EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCartRuleLangMgr) FetchByPrimaryKey(idCartRule uint32 ,idLang uint32 ) (result EgCartRuleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_lang = ?", idCartRule , idLang).Find(&result).Error
	
	return
}
 

 

	

