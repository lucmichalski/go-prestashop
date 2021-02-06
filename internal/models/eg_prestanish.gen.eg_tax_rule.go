package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgTaxRuleMgr struct {
	*_BaseMgr
}

// EgTaxRuleMgr open func
func EgTaxRuleMgr(db *gorm.DB) *_EgTaxRuleMgr {
	if db == nil {
		panic(fmt.Errorf("EgTaxRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTaxRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tax_rule"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTaxRuleMgr) GetTableName() string {
	return "eg_tax_rule"
}

// Get 获取
func (obj *_EgTaxRuleMgr) Get() (result EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTaxRuleMgr) Gets() (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTaxRule id_tax_rule获取 
func (obj *_EgTaxRuleMgr) WithIDTaxRule(idTaxRule int) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rule"] = idTaxRule })
}

// WithIDTaxRulesGroup id_tax_rules_group获取 
func (obj *_EgTaxRuleMgr) WithIDTaxRulesGroup(idTaxRulesGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithIDCountry id_country获取 
func (obj *_EgTaxRuleMgr) WithIDCountry(idCountry int) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDState id_state获取 
func (obj *_EgTaxRuleMgr) WithIDState(idState int) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

// WithZipcodeFrom zipcode_from获取 
func (obj *_EgTaxRuleMgr) WithZipcodeFrom(zipcodeFrom string) Option {
	return optionFunc(func(o *options) { o.query["zipcode_from"] = zipcodeFrom })
}

// WithZipcodeTo zipcode_to获取 
func (obj *_EgTaxRuleMgr) WithZipcodeTo(zipcodeTo string) Option {
	return optionFunc(func(o *options) { o.query["zipcode_to"] = zipcodeTo })
}

// WithIDTax id_tax获取 
func (obj *_EgTaxRuleMgr) WithIDTax(idTax int) Option {
	return optionFunc(func(o *options) { o.query["id_tax"] = idTax })
}

// WithBehavior behavior获取 
func (obj *_EgTaxRuleMgr) WithBehavior(behavior int) Option {
	return optionFunc(func(o *options) { o.query["behavior"] = behavior })
}

// WithDescription description获取 
func (obj *_EgTaxRuleMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}


// GetByOption 功能选项模式获取
func (obj *_EgTaxRuleMgr) GetByOption(opts ...Option) (result EgTaxRule, err error) {
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
func (obj *_EgTaxRuleMgr) GetByOptions(opts ...Option) (results []*EgTaxRule, err error) {
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


// GetFromIDTaxRule 通过id_tax_rule获取内容  
func (obj *_EgTaxRuleMgr)  GetFromIDTaxRule(idTaxRule int) (result EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rule = ?", idTaxRule).Find(&result).Error
	
	return
}

// GetBatchFromIDTaxRule 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromIDTaxRule(idTaxRules []int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rule IN (?)", idTaxRules).Find(&results).Error
	
	return
}
 
// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容  
func (obj *_EgTaxRuleMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error
	
	return
}
 
// GetFromIDCountry 通过id_country获取内容  
func (obj *_EgTaxRuleMgr) GetFromIDCountry(idCountry int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}

// GetBatchFromIDCountry 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromIDCountry(idCountrys []int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error
	
	return
}
 
// GetFromIDState 通过id_state获取内容  
func (obj *_EgTaxRuleMgr) GetFromIDState(idState int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&results).Error
	
	return
}

// GetBatchFromIDState 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromIDState(idStates []int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error
	
	return
}
 
// GetFromZipcodeFrom 通过zipcode_from获取内容  
func (obj *_EgTaxRuleMgr) GetFromZipcodeFrom(zipcodeFrom string) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zipcode_from = ?", zipcodeFrom).Find(&results).Error
	
	return
}

// GetBatchFromZipcodeFrom 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromZipcodeFrom(zipcodeFroms []string) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zipcode_from IN (?)", zipcodeFroms).Find(&results).Error
	
	return
}
 
// GetFromZipcodeTo 通过zipcode_to获取内容  
func (obj *_EgTaxRuleMgr) GetFromZipcodeTo(zipcodeTo string) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zipcode_to = ?", zipcodeTo).Find(&results).Error
	
	return
}

// GetBatchFromZipcodeTo 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromZipcodeTo(zipcodeTos []string) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("zipcode_to IN (?)", zipcodeTos).Find(&results).Error
	
	return
}
 
// GetFromIDTax 通过id_tax获取内容  
func (obj *_EgTaxRuleMgr) GetFromIDTax(idTax int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}

// GetBatchFromIDTax 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromIDTax(idTaxs []int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax IN (?)", idTaxs).Find(&results).Error
	
	return
}
 
// GetFromBehavior 通过behavior获取内容  
func (obj *_EgTaxRuleMgr) GetFromBehavior(behavior int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("behavior = ?", behavior).Find(&results).Error
	
	return
}

// GetBatchFromBehavior 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromBehavior(behaviors []int) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("behavior IN (?)", behaviors).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgTaxRuleMgr) GetFromDescription(description string) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgTaxRuleMgr) GetBatchFromDescription(descriptions []string) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTaxRuleMgr) FetchByPrimaryKey(idTaxRule int ) (result EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rule = ?", idTaxRule).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDTaxRulesGroup  获取多个内容
 func (obj *_EgTaxRuleMgr) FetchIndexByIDTaxRulesGroup(idTaxRulesGroup int ) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error
	
	return
}
 
 // FetchIndexByCategoryGetproducts  获取多个内容
 func (obj *_EgTaxRuleMgr) FetchIndexByCategoryGetproducts(idTaxRulesGroup int ,idCountry int ,idState int ,zipcodeFrom string ) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ? AND id_country = ? AND id_state = ? AND zipcode_from = ?", idTaxRulesGroup , idCountry , idState , zipcodeFrom).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDTax  获取多个内容
 func (obj *_EgTaxRuleMgr) FetchIndexByIDTax(idTax int ) (results []*EgTaxRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax = ?", idTax).Find(&results).Error
	
	return
}
 

	

