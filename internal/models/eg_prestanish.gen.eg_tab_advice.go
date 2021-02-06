package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgTabAdviceMgr struct {
	*_BaseMgr
}

// EgTabAdviceMgr open func
func EgTabAdviceMgr(db *gorm.DB) *_EgTabAdviceMgr {
	if db == nil {
		panic(fmt.Errorf("EgTabAdviceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTabAdviceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tab_advice"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTabAdviceMgr) GetTableName() string {
	return "eg_tab_advice"
}

// Get 获取
func (obj *_EgTabAdviceMgr) Get() (result EgTabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTabAdviceMgr) Gets() (results []*EgTabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTab id_tab获取 
func (obj *_EgTabAdviceMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

// WithIDAdvice id_advice获取 
func (obj *_EgTabAdviceMgr) WithIDAdvice(idAdvice int) Option {
	return optionFunc(func(o *options) { o.query["id_advice"] = idAdvice })
}


// GetByOption 功能选项模式获取
func (obj *_EgTabAdviceMgr) GetByOption(opts ...Option) (result EgTabAdvice, err error) {
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
func (obj *_EgTabAdviceMgr) GetByOptions(opts ...Option) (results []*EgTabAdvice, err error) {
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


// GetFromIDTab 通过id_tab获取内容  
func (obj *_EgTabAdviceMgr) GetFromIDTab(idTab int) (results []*EgTabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error
	
	return
}

// GetBatchFromIDTab 批量唯一主键查找 
func (obj *_EgTabAdviceMgr) GetBatchFromIDTab(idTabs []int) (results []*EgTabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error
	
	return
}
 
// GetFromIDAdvice 通过id_advice获取内容  
func (obj *_EgTabAdviceMgr) GetFromIDAdvice(idAdvice int) (results []*EgTabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice = ?", idAdvice).Find(&results).Error
	
	return
}

// GetBatchFromIDAdvice 批量唯一主键查找 
func (obj *_EgTabAdviceMgr) GetBatchFromIDAdvice(idAdvices []int) (results []*EgTabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_advice IN (?)", idAdvices).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTabAdviceMgr) FetchByPrimaryKey(idTab int ,idAdvice int ) (result EgTabAdvice, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ? AND id_advice = ?", idTab , idAdvice).Find(&result).Error
	
	return
}
 

 

	

