package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgCartRuleGroupMgr struct {
	*_BaseMgr
}

// EgCartRuleGroupMgr open func
func EgCartRuleGroupMgr(db *gorm.DB) *_EgCartRuleGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgCartRuleGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCartRuleGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cart_rule_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCartRuleGroupMgr) GetTableName() string {
	return "eg_cart_rule_group"
}

// Get 获取
func (obj *_EgCartRuleGroupMgr) Get() (result EgCartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCartRuleGroupMgr) Gets() (results []*EgCartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCartRule id_cart_rule获取 
func (obj *_EgCartRuleGroupMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithIDGroup id_group获取 
func (obj *_EgCartRuleGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}


// GetByOption 功能选项模式获取
func (obj *_EgCartRuleGroupMgr) GetByOption(opts ...Option) (result EgCartRuleGroup, err error) {
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
func (obj *_EgCartRuleGroupMgr) GetByOptions(opts ...Option) (results []*EgCartRuleGroup, err error) {
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
func (obj *_EgCartRuleGroupMgr) GetFromIDCartRule(idCartRule uint32) (results []*EgCartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error
	
	return
}

// GetBatchFromIDCartRule 批量唯一主键查找 
func (obj *_EgCartRuleGroupMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*EgCartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error
	
	return
}
 
// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgCartRuleGroupMgr) GetFromIDGroup(idGroup uint32) (results []*EgCartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgCartRuleGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgCartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCartRuleGroupMgr) FetchByPrimaryKey(idCartRule uint32 ,idGroup uint32 ) (result EgCartRuleGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_group = ?", idCartRule , idGroup).Find(&result).Error
	
	return
}
 

 

	

