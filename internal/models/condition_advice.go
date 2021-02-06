package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgConditionAdviceMgr struct {
	*_BaseMgr
}

// EgConditionAdviceMgr open func
func EgConditionAdviceMgr(db *gorm.DB) *_EgConditionAdviceMgr {
	if db == nil {
		panic(fmt.Errorf("EgConditionAdviceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgConditionAdviceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_condition_advice"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgConditionAdviceMgr) GetTableName() string {
	return "eg_condition_advice"
}

// Get 获取
func (obj *_EgConditionAdviceMgr) Get() (result EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgConditionAdviceMgr) Gets() (results []*EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCondition id_condition获取 
func (obj *_EgConditionAdviceMgr) WithIDCondition(idCondition int) Option {
	return optionFunc(func(o *options) { o.query["id_condition"] = idCondition })
}

// WithIDAdvice id_advice获取 
func (obj *_EgConditionAdviceMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}

// WithDisplay display获取 
func (obj *_EgConditionAdviceMgr) WithDisplay(display bool) Option {
	return optionFunc(func(o *options) { o.query["display"] = display })
}


// GetByOption 功能选项模式获取
func (obj *_EgConditionAdviceMgr) GetByOption(opts ...Option) (result EgConditionAdvice, err error) {
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
func (obj *_EgConditionAdviceMgr) GetByOptions(opts ...Option) (results []*EgConditionAdvice, err error) {
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
func (obj *_EgConditionAdviceMgr) GetFromIDCondition(idCondition int) (results []*EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ?", idCondition).Find(&results).Error
	
	return
}

// GetBatchFromIDCondition 批量唯一主键查找 
func (obj *_EgConditionAdviceMgr) GetBatchFromIDCondition(idConditions []int) (results []*EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition IN (?)", idConditions).Find(&results).Error
	
	return
}
 
// GetFromIDAdvice 通过id_advice获取内容  
func (obj *_EgConditionAdviceMgr) GetFromIDAdvice(idAdvice int) (results []*EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&results).Error
	
	return
}

// GetBatchFromIDAdvice 批量唯一主键查找 
func (obj *_EgConditionAdviceMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error
	
	return
}
 
// GetFromDisplay 通过display获取内容  
func (obj *_EgConditionAdviceMgr) GetFromDisplay(display bool) (results []*EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display = ?", display).Find(&results).Error
	
	return
}

// GetBatchFromDisplay 批量唯一主键查找 
func (obj *_EgConditionAdviceMgr) GetBatchFromDisplay(displays []bool) (results []*EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display IN (?)", displays).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgConditionAdviceMgr) FetchByPrimaryKey(idCondition int ,idAdvice int ) (result EgConditionAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_condition = ? AND id_advice = ?", idCondition , idAdvice).Find(&result).Error
	
	return
}
 

 

	

