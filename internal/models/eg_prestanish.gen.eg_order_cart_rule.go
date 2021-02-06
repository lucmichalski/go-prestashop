package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOrderCartRuleMgr struct {
	*_BaseMgr
}

// EgOrderCartRuleMgr open func
func EgOrderCartRuleMgr(db *gorm.DB) *_EgOrderCartRuleMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderCartRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderCartRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_cart_rule"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderCartRuleMgr) GetTableName() string {
	return "eg_order_cart_rule"
}

// Get 获取
func (obj *_EgOrderCartRuleMgr) Get() (result EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderCartRuleMgr) Gets() (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderCartRule id_order_cart_rule获取 
func (obj *_EgOrderCartRuleMgr) WithIDOrderCartRule(idOrderCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_cart_rule"] = idOrderCartRule })
}

// WithIDOrder id_order获取 
func (obj *_EgOrderCartRuleMgr) WithIDOrder(idOrder uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order"] = idOrder })
}

// WithIDCartRule id_cart_rule获取 
func (obj *_EgOrderCartRuleMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// WithIDOrderInvoice id_order_invoice获取 
func (obj *_EgOrderCartRuleMgr) WithIDOrderInvoice(idOrderInvoice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_invoice"] = idOrderInvoice })
}

// WithName name获取 
func (obj *_EgOrderCartRuleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取 
func (obj *_EgOrderCartRuleMgr) WithValue(value float64) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithValueTaxExcl value_tax_excl获取 
func (obj *_EgOrderCartRuleMgr) WithValueTaxExcl(valueTaxExcl float64) Option {
	return optionFunc(func(o *options) { o.query["value_tax_excl"] = valueTaxExcl })
}

// WithFreeShipping free_shipping获取 
func (obj *_EgOrderCartRuleMgr) WithFreeShipping(freeShipping bool) Option {
	return optionFunc(func(o *options) { o.query["free_shipping"] = freeShipping })
}

// WithDeleted deleted获取 
func (obj *_EgOrderCartRuleMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderCartRuleMgr) GetByOption(opts ...Option) (result EgOrderCartRule, err error) {
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
func (obj *_EgOrderCartRuleMgr) GetByOptions(opts ...Option) (results []*EgOrderCartRule, err error) {
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


// GetFromIDOrderCartRule 通过id_order_cart_rule获取内容  
func (obj *_EgOrderCartRuleMgr)  GetFromIDOrderCartRule(idOrderCartRule uint32) (result EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule = ?", idOrderCartRule).Find(&result).Error
	
	return
}

// GetBatchFromIDOrderCartRule 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromIDOrderCartRule(idOrderCartRules []uint32) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule IN (?)", idOrderCartRules).Find(&results).Error
	
	return
}
 
// GetFromIDOrder 通过id_order获取内容  
func (obj *_EgOrderCartRuleMgr) GetFromIDOrder(idOrder uint32) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}

// GetBatchFromIDOrder 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromIDOrder(idOrders []uint32) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order IN (?)", idOrders).Find(&results).Error
	
	return
}
 
// GetFromIDCartRule 通过id_cart_rule获取内容  
func (obj *_EgOrderCartRuleMgr) GetFromIDCartRule(idCartRule uint32) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error
	
	return
}

// GetBatchFromIDCartRule 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error
	
	return
}
 
// GetFromIDOrderInvoice 通过id_order_invoice获取内容  
func (obj *_EgOrderCartRuleMgr) GetFromIDOrderInvoice(idOrderInvoice uint32) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice = ?", idOrderInvoice).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderInvoice 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromIDOrderInvoice(idOrderInvoices []uint32) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_invoice IN (?)", idOrderInvoices).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgOrderCartRuleMgr) GetFromName(name string) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromName(names []string) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_EgOrderCartRuleMgr) GetFromValue(value float64) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromValue(values []float64) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromValueTaxExcl 通过value_tax_excl获取内容  
func (obj *_EgOrderCartRuleMgr) GetFromValueTaxExcl(valueTaxExcl float64) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value_tax_excl = ?", valueTaxExcl).Find(&results).Error
	
	return
}

// GetBatchFromValueTaxExcl 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromValueTaxExcl(valueTaxExcls []float64) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value_tax_excl IN (?)", valueTaxExcls).Find(&results).Error
	
	return
}
 
// GetFromFreeShipping 通过free_shipping获取内容  
func (obj *_EgOrderCartRuleMgr) GetFromFreeShipping(freeShipping bool) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping = ?", freeShipping).Find(&results).Error
	
	return
}

// GetBatchFromFreeShipping 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromFreeShipping(freeShippings []bool) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("free_shipping IN (?)", freeShippings).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgOrderCartRuleMgr) GetFromDeleted(deleted bool) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgOrderCartRuleMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderCartRuleMgr) FetchByPrimaryKey(idOrderCartRule uint32 ) (result EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_cart_rule = ?", idOrderCartRule).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDOrder  获取多个内容
 func (obj *_EgOrderCartRuleMgr) FetchIndexByIDOrder(idOrder uint32 ) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order = ?", idOrder).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCartRule  获取多个内容
 func (obj *_EgOrderCartRuleMgr) FetchIndexByIDCartRule(idCartRule uint32 ) (results []*EgOrderCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error
	
	return
}
 

	

