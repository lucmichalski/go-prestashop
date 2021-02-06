package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProductGroupReductionCacheMgr struct {
	*_BaseMgr
}

// EgProductGroupReductionCacheMgr open func
func EgProductGroupReductionCacheMgr(db *gorm.DB) *_EgProductGroupReductionCacheMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductGroupReductionCacheMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductGroupReductionCacheMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_group_reduction_cache"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductGroupReductionCacheMgr) GetTableName() string {
	return "eg_product_group_reduction_cache"
}

// Get 获取
func (obj *_EgProductGroupReductionCacheMgr) Get() (result EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductGroupReductionCacheMgr) Gets() (results []*EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgProductGroupReductionCacheMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDGroup id_group获取 
func (obj *_EgProductGroupReductionCacheMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithReduction reduction获取 
func (obj *_EgProductGroupReductionCacheMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductGroupReductionCacheMgr) GetByOption(opts ...Option) (result EgProductGroupReductionCache, err error) {
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
func (obj *_EgProductGroupReductionCacheMgr) GetByOptions(opts ...Option) (results []*EgProductGroupReductionCache, err error) {
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


// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgProductGroupReductionCacheMgr) GetFromIDProduct(idProduct uint32) (results []*EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductGroupReductionCacheMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgProductGroupReductionCacheMgr) GetFromIDGroup(idGroup uint32) (results []*EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgProductGroupReductionCacheMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
// GetFromReduction 通过reduction获取内容  
func (obj *_EgProductGroupReductionCacheMgr) GetFromReduction(reduction float64) (results []*EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error
	
	return
}

// GetBatchFromReduction 批量唯一主键查找 
func (obj *_EgProductGroupReductionCacheMgr) GetBatchFromReduction(reductions []float64) (results []*EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductGroupReductionCacheMgr) FetchByPrimaryKey(idProduct uint32 ,idGroup uint32 ) (result EgProductGroupReductionCache, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_group = ?", idProduct , idGroup).Find(&result).Error
	
	return
}
 

 

	

