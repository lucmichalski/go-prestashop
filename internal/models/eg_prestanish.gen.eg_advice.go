package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgAdviceMgr struct {
	*_BaseMgr
}

// EgAdviceMgr open func
func EgAdviceMgr(db *gorm.DB) *_EgAdviceMgr {
	if db == nil {
		panic(fmt.Errorf("EgAdviceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAdviceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_advice"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAdviceMgr) GetTableName() string {
	return "eg_advice"
}

// Get 获取
func (obj *_EgAdviceMgr) Get() (result EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAdviceMgr) Gets() (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAdvice id_advice获取 
func (obj *_EgAdviceMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}

// WithIDPsAdvice id_ps_advice获取 
func (obj *_EgAdviceMgr) WithIDPsAdvice(idPsAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_ps_advice"] = idPsAdvice })
}

// WithIDTab id_tab获取 
func (obj *_EgAdviceMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

// WithIDsTab ids_tab获取 
func (obj *_EgAdviceMgr) WithIDsTab(idsTab string) Option {
	return optionFunc(func(o *options) { o.query["ids_tab"] = idsTab })
}

// WithValidated validated获取 
func (obj *_EgAdviceMgr) WithValidated(validated bool) Option {
	return optionFunc(func(o *options) { o.query["validated"] = validated })
}

// WithHide hide获取 
func (obj *_EgAdviceMgr) WithHide(hide bool) Option {
	return optionFunc(func(o *options) { o.query["hide"] = hide })
}

// WithLocation location获取 
func (obj *_EgAdviceMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithSelector selector获取 
func (obj *_EgAdviceMgr) WithSelector(selector string) Option {
	return optionFunc(func(o *options) { o.query["selector"] = selector })
}

// WithStartDay start_day获取 
func (obj *_EgAdviceMgr) WithStartDay(startDay int) Option {
	return optionFunc(func(o *options) { o.query["start_day"] = startDay })
}

// WithStopDay stop_day获取 
func (obj *_EgAdviceMgr) WithStopDay(stopDay int) Option {
	return optionFunc(func(o *options) { o.query["stop_day"] = stopDay })
}

// WithWeight weight获取 
func (obj *_EgAdviceMgr) WithWeight(weight int) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}


// GetByOption 功能选项模式获取
func (obj *_EgAdviceMgr) GetByOption(opts ...Option) (result EgAdvice, err error) {
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
func (obj *_EgAdviceMgr) GetByOptions(opts ...Option) (results []*EgAdvice, err error) {
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


// GetFromIDAdvice 通过id_advice获取内容  
func (obj *_EgAdviceMgr)  GetFromIDAdvice(idAdvice int) (result EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&result).Error
	
	return
}

// GetBatchFromIDAdvice 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error
	
	return
}
 
// GetFromIDPsAdvice 通过id_ps_advice获取内容  
func (obj *_EgAdviceMgr) GetFromIDPsAdvice(idPsAdvice int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_advice = ?", idPsAdvice).Find(&results).Error
	
	return
}

// GetBatchFromIDPsAdvice 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromIDPsAdvice(idPsAdvices []int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_ps_advice IN (?)", idPsAdvices).Find(&results).Error
	
	return
}
 
// GetFromIDTab 通过id_tab获取内容  
func (obj *_EgAdviceMgr) GetFromIDTab(idTab int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error
	
	return
}

// GetBatchFromIDTab 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromIDTab(idTabs []int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error
	
	return
}
 
// GetFromIDsTab 通过ids_tab获取内容  
func (obj *_EgAdviceMgr) GetFromIDsTab(idsTab string) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ids_tab = ?", idsTab).Find(&results).Error
	
	return
}

// GetBatchFromIDsTab 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromIDsTab(idsTabs []string) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ids_tab IN (?)", idsTabs).Find(&results).Error
	
	return
}
 
// GetFromValidated 通过validated获取内容  
func (obj *_EgAdviceMgr) GetFromValidated(validated bool) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated = ?", validated).Find(&results).Error
	
	return
}

// GetBatchFromValidated 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromValidated(validateds []bool) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("validated IN (?)", validateds).Find(&results).Error
	
	return
}
 
// GetFromHide 通过hide获取内容  
func (obj *_EgAdviceMgr) GetFromHide(hide bool) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide = ?", hide).Find(&results).Error
	
	return
}

// GetBatchFromHide 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromHide(hides []bool) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide IN (?)", hides).Find(&results).Error
	
	return
}
 
// GetFromLocation 通过location获取内容  
func (obj *_EgAdviceMgr) GetFromLocation(location string) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error
	
	return
}

// GetBatchFromLocation 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromLocation(locations []string) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error
	
	return
}
 
// GetFromSelector 通过selector获取内容  
func (obj *_EgAdviceMgr) GetFromSelector(selector string) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("selector = ?", selector).Find(&results).Error
	
	return
}

// GetBatchFromSelector 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromSelector(selectors []string) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("selector IN (?)", selectors).Find(&results).Error
	
	return
}
 
// GetFromStartDay 通过start_day获取内容  
func (obj *_EgAdviceMgr) GetFromStartDay(startDay int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_day = ?", startDay).Find(&results).Error
	
	return
}

// GetBatchFromStartDay 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromStartDay(startDays []int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_day IN (?)", startDays).Find(&results).Error
	
	return
}
 
// GetFromStopDay 通过stop_day获取内容  
func (obj *_EgAdviceMgr) GetFromStopDay(stopDay int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stop_day = ?", stopDay).Find(&results).Error
	
	return
}

// GetBatchFromStopDay 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromStopDay(stopDays []int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stop_day IN (?)", stopDays).Find(&results).Error
	
	return
}
 
// GetFromWeight 通过weight获取内容  
func (obj *_EgAdviceMgr) GetFromWeight(weight int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error
	
	return
}

// GetBatchFromWeight 批量唯一主键查找 
func (obj *_EgAdviceMgr) GetBatchFromWeight(weights []int) (results []*EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAdviceMgr) FetchByPrimaryKey(idAdvice int ) (result EgAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&result).Error
	
	return
}
 

 

	

