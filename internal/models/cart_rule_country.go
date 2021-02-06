package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCartRuleCountryMgr struct {
	*_BaseMgr
}

// EgCartRuleCountryMgr open func
func EgCartRuleCountryMgr(db *gorm.DB) *_EgCartRuleCountryMgr {
	if db == nil {
		panic(fmt.Errorf("EgCartRuleCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCartRuleCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cart_rule_country"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCartRuleCountryMgr) GetTableName() string {
	return "eg_cart_rule_country"
}

// Get 获取
func (obj *_EgCartRuleCountryMgr) Get() (result EgCartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCartRuleCountryMgr) Gets() (results []*EgCartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCartRule id_cart_rule获取 
func (obj *_EgCartRuleCountryMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithIDCountry id_country获取 
func (obj *_EgCartRuleCountryMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}


// GetByOption 功能选项模式获取
func (obj *_EgCartRuleCountryMgr) GetByOption(opts ...Option) (result EgCartRuleCountry, err error) {
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
func (obj *_EgCartRuleCountryMgr) GetByOptions(opts ...Option) (results []*EgCartRuleCountry, err error) {
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
func (obj *_EgCartRuleCountryMgr) GetFromIDCartRule(idCartRule uint32) (results []*EgCartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error
	
	return
}

// GetBatchFromIDCartRule 批量唯一主键查找 
func (obj *_EgCartRuleCountryMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*EgCartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error
	
	return
}
 
// GetFromIDCountry 通过id_country获取内容  
func (obj *_EgCartRuleCountryMgr) GetFromIDCountry(idCountry uint32) (results []*EgCartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}

// GetBatchFromIDCountry 批量唯一主键查找 
func (obj *_EgCartRuleCountryMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*EgCartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCartRuleCountryMgr) FetchByPrimaryKey(idCartRule uint32 ,idCountry uint32 ) (result EgCartRuleCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_country = ?", idCartRule , idCountry).Find(&result).Error
	
	return
}
 

 

	

