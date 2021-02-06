package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _EgConditionMgr struct {
	*_BaseMgr
}

// EgConditionMgr open func
func EgConditionMgr(db *gorm.DB) *_EgConditionMgr {
	if db == nil {
		panic(fmt.Errorf("EgConditionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgConditionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_condition"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgConditionMgr) GetTableName() string {
	return "eg_condition"
}

// Get 获取
func (obj *_EgConditionMgr) Get() (result EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgConditionMgr) Gets() (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCondition id_condition获取 
func (obj *_EgConditionMgr) WithIDCondition(idCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_condition"] = idCondition })
}

// WithIDPsCondition id_ps_condition获取 
func (obj *_EgConditionMgr) WithIDPsCondition(idPsCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_ps_condition"] = idPsCondition })
}

// WithType type获取 
func (obj *_EgConditionMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithRequest request获取 
func (obj *_EgConditionMgr) WithRequest(request string) Option {
	return optionFunc(func(o *options) { o.query["request"] = request })
}

// WithOperator operator获取 
func (obj *_EgConditionMgr) WithOperator(operator string) Option {
	return optionFunc(func(o *options) { o.query["operator"] = operator })
}

// WithValue value获取 
func (obj *_EgConditionMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithResult result获取 
func (obj *_EgConditionMgr) WithResult(result string) Option {
	return optionFunc(func(o *options) { o.query["result"] = result })
}

// WithCalculationType calculation_type获取 
func (obj *_EgConditionMgr) WithCalculationType(calculationType string) Option {
	return optionFunc(func(o *options) { o.query["calculation_type"] = calculationType })
}

// WithCalculationDetail calculation_detail获取 
func (obj *_EgConditionMgr) WithCalculationDetail(calculationDetail string) Option {
	return optionFunc(func(o *options) { o.query["calculation_detail"] = calculationDetail })
}

// WithValidated validated获取 
func (obj *_EgConditionMgr) WithValidated(validated bool) Option {
	return optionFunc(func(o *options) { o.query["validated"] = validated })
}

// WithDateAdd date_add获取 
func (obj *_EgConditionMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取 
func (obj *_EgConditionMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}


// GetByOption 功能选项模式获取
func (obj *_EgConditionMgr) GetByOption(opts ...Option) (result EgCondition, err error) {
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
func (obj *_EgConditionMgr) GetByOptions(opts ...Option) (results []*EgCondition, err error) {
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


// GetFromIDCondition 通过id_condition获取内容  
func (obj *_EgConditionMgr) GetFromIDCondition(idCondition int) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ?", idCondition).Find(&results).Error
	
	return
}

// GetBatchFromIDCondition 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromIDCondition(idConditions []int) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition IN (?)", idConditions).Find(&results).Error
	
	return
}
 
// GetFromIDPsCondition 通过id_ps_condition获取内容  
func (obj *_EgConditionMgr) GetFromIDPsCondition(idPsCondition int) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_condition = ?", idPsCondition).Find(&results).Error
	
	return
}

// GetBatchFromIDPsCondition 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromIDPsCondition(idPsConditions []int) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_condition IN (?)", idPsConditions).Find(&results).Error
	
	return
}
 
// GetFromType 通过type获取内容  
func (obj *_EgConditionMgr) GetFromType(_type string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromType(_types []string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromRequest 通过request获取内容  
func (obj *_EgConditionMgr) GetFromRequest(request string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request = ?", request).Find(&results).Error
	
	return
}

// GetBatchFromRequest 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromRequest(requests []string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request IN (?)", requests).Find(&results).Error
	
	return
}
 
// GetFromOperator 通过operator获取内容  
func (obj *_EgConditionMgr) GetFromOperator(operator string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("operator = ?", operator).Find(&results).Error
	
	return
}

// GetBatchFromOperator 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromOperator(operators []string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("operator IN (?)", operators).Find(&results).Error
	
	return
}
 
// GetFromValue 通过value获取内容  
func (obj *_EgConditionMgr) GetFromValue(value string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error
	
	return
}

// GetBatchFromValue 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromValue(values []string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error
	
	return
}
 
// GetFromResult 通过result获取内容  
func (obj *_EgConditionMgr) GetFromResult(result string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("result = ?", result).Find(&results).Error
	
	return
}

// GetBatchFromResult 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromResult(results []string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("result IN (?)", results).Find(&results).Error
	
	return
}
 
// GetFromCalculationType 通过calculation_type获取内容  
func (obj *_EgConditionMgr) GetFromCalculationType(calculationType string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calculation_type = ?", calculationType).Find(&results).Error
	
	return
}

// GetBatchFromCalculationType 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromCalculationType(calculationTypes []string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calculation_type IN (?)", calculationTypes).Find(&results).Error
	
	return
}
 
// GetFromCalculationDetail 通过calculation_detail获取内容  
func (obj *_EgConditionMgr) GetFromCalculationDetail(calculationDetail string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calculation_detail = ?", calculationDetail).Find(&results).Error
	
	return
}

// GetBatchFromCalculationDetail 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromCalculationDetail(calculationDetails []string) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("calculation_detail IN (?)", calculationDetails).Find(&results).Error
	
	return
}
 
// GetFromValidated 通过validated获取内容  
func (obj *_EgConditionMgr) GetFromValidated(validated bool) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated = ?", validated).Find(&results).Error
	
	return
}

// GetBatchFromValidated 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromValidated(validateds []bool) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated IN (?)", validateds).Find(&results).Error
	
	return
}
 
// GetFromDateAdd 通过date_add获取内容  
func (obj *_EgConditionMgr) GetFromDateAdd(dateAdd time.Time) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error
	
	return
}

// GetBatchFromDateAdd 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error
	
	return
}
 
// GetFromDateUpd 通过date_upd获取内容  
func (obj *_EgConditionMgr) GetFromDateUpd(dateUpd time.Time) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error
	
	return
}

// GetBatchFromDateUpd 批量唯一主键查找 
func (obj *_EgConditionMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgConditionMgr) FetchByPrimaryKey(idCondition int ,idPsCondition int ) (result EgCondition, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ? AND id_ps_condition = ?", idCondition , idPsCondition).Find(&result).Error
	
	return
}
 

 

	

