package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgSpecificPricePriorityMgr struct {
	*_BaseMgr
}

// EgSpecificPricePriorityMgr open func
func EgSpecificPricePriorityMgr(db *gorm.DB) *_EgSpecificPricePriorityMgr {
	if db == nil {
		panic(fmt.Errorf("EgSpecificPricePriorityMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSpecificPricePriorityMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_specific_price_priority"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSpecificPricePriorityMgr) GetTableName() string {
	return "eg_specific_price_priority"
}

// Get 获取
func (obj *_EgSpecificPricePriorityMgr) Get() (result EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSpecificPricePriorityMgr) Gets() (results []*EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSpecificPricePriority id_specific_price_priority获取 
func (obj *_EgSpecificPricePriorityMgr) WithIDSpecificPricePriority(idSpecificPricePriority int) Option {
	return optionFunc(func(o *options) { o.query["id_specific_price_priority"] = idSpecificPricePriority })
}

// WithIDProduct id_product获取 
func (obj *_EgSpecificPricePriorityMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithPriority priority获取 
func (obj *_EgSpecificPricePriorityMgr) WithPriority(priority string) Option {
	return optionFunc(func(o *options) { o.query["priority"] = priority })
}


// GetByOption 功能选项模式获取
func (obj *_EgSpecificPricePriorityMgr) GetByOption(opts ...Option) (result EgSpecificPricePriority, err error) {
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
func (obj *_EgSpecificPricePriorityMgr) GetByOptions(opts ...Option) (results []*EgSpecificPricePriority, err error) {
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


// GetFromIDSpecificPricePriority 通过id_specific_price_priority获取内容  
func (obj *_EgSpecificPricePriorityMgr) GetFromIDSpecificPricePriority(idSpecificPricePriority int) (results []*EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_priority = ?", idSpecificPricePriority).Find(&results).Error
	
	return
}

// GetBatchFromIDSpecificPricePriority 批量唯一主键查找 
func (obj *_EgSpecificPricePriorityMgr) GetBatchFromIDSpecificPricePriority(idSpecificPricePrioritys []int) (results []*EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_priority IN (?)", idSpecificPricePrioritys).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgSpecificPricePriorityMgr) GetFromIDProduct(idProduct int) (results []*EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgSpecificPricePriorityMgr) GetBatchFromIDProduct(idProducts []int) (results []*EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromPriority 通过priority获取内容  
func (obj *_EgSpecificPricePriorityMgr) GetFromPriority(priority string) (results []*EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("priority = ?", priority).Find(&results).Error
	
	return
}

// GetBatchFromPriority 批量唯一主键查找 
func (obj *_EgSpecificPricePriorityMgr) GetBatchFromPriority(prioritys []string) (results []*EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("priority IN (?)", prioritys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSpecificPricePriorityMgr) FetchByPrimaryKey(idSpecificPricePriority int ,idProduct int ) (result EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_specific_price_priority = ? AND id_product = ?", idSpecificPricePriority , idProduct).Find(&result).Error
	
	return
}
 
 // FetchUniqueByIDProduct primay or index 获取唯一内容
 func (obj *_EgSpecificPricePriorityMgr) FetchUniqueByIDProduct(idProduct int ) (result EgSpecificPricePriority, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&result).Error
	
	return
}
 

 

	

