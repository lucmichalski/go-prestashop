package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCartRuleCarrierMgr struct {
	*_BaseMgr
}

// EgCartRuleCarrierMgr open func
func EgCartRuleCarrierMgr(db *gorm.DB) *_EgCartRuleCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("EgCartRuleCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCartRuleCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cart_rule_carrier"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCartRuleCarrierMgr) GetTableName() string {
	return "eg_cart_rule_carrier"
}

// Get 获取
func (obj *_EgCartRuleCarrierMgr) Get() (result EgCartRuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCartRuleCarrierMgr) Gets() (results []*EgCartRuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCartRule id_cart_rule获取 
func (obj *_EgCartRuleCarrierMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithIDCarrier id_carrier获取 
func (obj *_EgCartRuleCarrierMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}


// GetByOption 功能选项模式获取
func (obj *_EgCartRuleCarrierMgr) GetByOption(opts ...Option) (result EgCartRuleCarrier, err error) {
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
func (obj *_EgCartRuleCarrierMgr) GetByOptions(opts ...Option) (results []*EgCartRuleCarrier, err error) {
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
func (obj *_EgCartRuleCarrierMgr) GetFromIDCartRule(idCartRule uint32) (results []*EgCartRuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error
	
	return
}

// GetBatchFromIDCartRule 批量唯一主键查找 
func (obj *_EgCartRuleCarrierMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*EgCartRuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error
	
	return
}
 
// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgCartRuleCarrierMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgCartRuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgCartRuleCarrierMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgCartRuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCartRuleCarrierMgr) FetchByPrimaryKey(idCartRule uint32 ,idCarrier uint32 ) (result EgCartRuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ? AND id_carrier = ?", idCartRule , idCarrier).Find(&result).Error
	
	return
}
 

 

	

