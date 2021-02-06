package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgSpecificPriceRuleConditionGroupMgr struct {
	*_BaseMgr
}

// EgSpecificPriceRuleConditionGroupMgr open func
func EgSpecificPriceRuleConditionGroupMgr(db *gorm.DB) *_EgSpecificPriceRuleConditionGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgSpecificPriceRuleConditionGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSpecificPriceRuleConditionGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_specific_price_rule_condition_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSpecificPriceRuleConditionGroupMgr) GetTableName() string {
	return "eg_specific_price_rule_condition_group"
}

// Get 获取
func (obj *_EgSpecificPriceRuleConditionGroupMgr) Get() (result EgSpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSpecificPriceRuleConditionGroupMgr) Gets() (results []*EgSpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSpecificPriceRuleConditionGroup id_specific_price_rule_condition_group获取 
func (obj *_EgSpecificPriceRuleConditionGroupMgr) WithIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule_condition_group"] = idSpecificPriceRuleConditionGroup })
}

// WithIDSpecificPriceRule id_specific_price_rule获取 
func (obj *_EgSpecificPriceRuleConditionGroupMgr) WithIDSpecificPriceRule(idSpecificPriceRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_rule"] = idSpecificPriceRule })
}


// GetByOption 功能选项模式获取
func (obj *_EgSpecificPriceRuleConditionGroupMgr) GetByOption(opts ...Option) (result EgSpecificPriceRuleConditionGroup, err error) {
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
func (obj *_EgSpecificPriceRuleConditionGroupMgr) GetByOptions(opts ...Option) (results []*EgSpecificPriceRuleConditionGroup, err error) {
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


// GetFromIDSpecificPriceRuleConditionGroup 通过id_specific_price_rule_condition_group获取内容  
func (obj *_EgSpecificPriceRuleConditionGroupMgr) GetFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroup uint32) (results []*EgSpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ?", idSpecificPriceRuleConditionGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDSpecificPriceRuleConditionGroup 批量唯一主键查找 
func (obj *_EgSpecificPriceRuleConditionGroupMgr) GetBatchFromIDSpecificPriceRuleConditionGroup(idSpecificPriceRuleConditionGroups []uint32) (results []*EgSpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group IN (?)", idSpecificPriceRuleConditionGroups).Find(&results).Error
	
	return
}
 
// GetFromIDSpecificPriceRule 通过id_specific_price_rule获取内容  
func (obj *_EgSpecificPriceRuleConditionGroupMgr) GetFromIDSpecificPriceRule(idSpecificPriceRule uint32) (results []*EgSpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule = ?", idSpecificPriceRule).Find(&results).Error
	
	return
}

// GetBatchFromIDSpecificPriceRule 批量唯一主键查找 
func (obj *_EgSpecificPriceRuleConditionGroupMgr) GetBatchFromIDSpecificPriceRule(idSpecificPriceRules []uint32) (results []*EgSpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule IN (?)", idSpecificPriceRules).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSpecificPriceRuleConditionGroupMgr) FetchByPrimaryKey(idSpecificPriceRuleConditionGroup uint32 ,idSpecificPriceRule uint32 ) (result EgSpecificPriceRuleConditionGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_rule_condition_group = ? AND id_specific_price_rule = ?", idSpecificPriceRuleConditionGroup , idSpecificPriceRule).Find(&result).Error
	
	return
}
 

 

	

