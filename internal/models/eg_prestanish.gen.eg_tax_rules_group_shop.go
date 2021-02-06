package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgTaxRulesGroupShopMgr struct {
	*_BaseMgr
}

// EgTaxRulesGroupShopMgr open func
func EgTaxRulesGroupShopMgr(db *gorm.DB) *_EgTaxRulesGroupShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgTaxRulesGroupShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTaxRulesGroupShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tax_rules_group_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTaxRulesGroupShopMgr) GetTableName() string {
	return "eg_tax_rules_group_shop"
}

// Get 获取
func (obj *_EgTaxRulesGroupShopMgr) Get() (result EgTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTaxRulesGroupShopMgr) Gets() (results []*EgTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTaxRulesGroup id_tax_rules_group获取 
func (obj *_EgTaxRulesGroupShopMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithIDShop id_shop获取 
func (obj *_EgTaxRulesGroupShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgTaxRulesGroupShopMgr) GetByOption(opts ...Option) (result EgTaxRulesGroupShop, err error) {
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
func (obj *_EgTaxRulesGroupShopMgr) GetByOptions(opts ...Option) (results []*EgTaxRulesGroupShop, err error) {
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


// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容  
func (obj *_EgTaxRulesGroupShopMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*EgTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找 
func (obj *_EgTaxRulesGroupShopMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*EgTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgTaxRulesGroupShopMgr) GetFromIDShop(idShop uint32) (results []*EgTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgTaxRulesGroupShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTaxRulesGroupShopMgr) FetchByPrimaryKey(idTaxRulesGroup uint32 ,idShop uint32 ) (result EgTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ? AND id_shop = ?", idTaxRulesGroup , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgTaxRulesGroupShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgTaxRulesGroupShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

