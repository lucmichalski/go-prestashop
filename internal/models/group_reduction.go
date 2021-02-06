package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgGroupReductionMgr struct {
	*_BaseMgr
}

// EgGroupReductionMgr open func
func EgGroupReductionMgr(db *gorm.DB) *_EgGroupReductionMgr {
	if db == nil {
		panic(fmt.Errorf("EgGroupReductionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgGroupReductionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_group_reduction"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgGroupReductionMgr) GetTableName() string {
	return "eg_group_reduction"
}

// Get 获取
func (obj *_EgGroupReductionMgr) Get() (result EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgGroupReductionMgr) Gets() (results []*EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGroupReduction id_group_reduction获取 
func (obj *_EgGroupReductionMgr) WithIDGroupReduction(idGroupReduction string) Option {
	return optionFunc(func(o *options) { o.query["id_group_reduction"] = idGroupReduction })
}

// WithIDGroup id_group获取 
func (obj *_EgGroupReductionMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithIDCategory id_category获取 
func (obj *_EgGroupReductionMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithReduction reduction获取 
func (obj *_EgGroupReductionMgr) WithReduction(reduction float64) Option {
	return optionFunc(func(o *options) { o.query["reduction"] = reduction })
}


// GetByOption 功能选项模式获取
func (obj *_EgGroupReductionMgr) GetByOption(opts ...Option) (result EgGroupReduction, err error) {
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
func (obj *_EgGroupReductionMgr) GetByOptions(opts ...Option) (results []*EgGroupReduction, err error) {
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


// GetFromIDGroupReduction 通过id_group_reduction获取内容  
func (obj *_EgGroupReductionMgr)  GetFromIDGroupReduction(idGroupReduction string) (result EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group_reduction = ?", idGroupReduction).Find(&result).Error
	
	return
}

// GetBatchFromIDGroupReduction 批量唯一主键查找 
func (obj *_EgGroupReductionMgr) GetBatchFromIDGroupReduction(idGroupReductions []string) (results []*EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group_reduction IN (?)", idGroupReductions).Find(&results).Error
	
	return
}
 
// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgGroupReductionMgr) GetFromIDGroup(idGroup uint32) (results []*EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgGroupReductionMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
// GetFromIDCategory 通过id_category获取内容  
func (obj *_EgGroupReductionMgr) GetFromIDCategory(idCategory uint32) (results []*EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCategory 批量唯一主键查找 
func (obj *_EgGroupReductionMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error
	
	return
}
 
// GetFromReduction 通过reduction获取内容  
func (obj *_EgGroupReductionMgr) GetFromReduction(reduction float64) (results []*EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction = ?", reduction).Find(&results).Error
	
	return
}

// GetBatchFromReduction 批量唯一主键查找 
func (obj *_EgGroupReductionMgr) GetBatchFromReduction(reductions []float64) (results []*EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reduction IN (?)", reductions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgGroupReductionMgr) FetchByPrimaryKey(idGroupReduction string ) (result EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group_reduction = ?", idGroupReduction).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDGroup primay or index 获取唯一内容
 func (obj *_EgGroupReductionMgr) FetchUniqueIndexByIDGroup(idGroup uint32 ,idCategory uint32 ) (result EgGroupReduction, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_category = ?", idGroup , idCategory).Find(&result).Error
	
	return
}
 

 

	

