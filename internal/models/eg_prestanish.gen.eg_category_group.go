package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgCategoryGroupMgr struct {
	*_BaseMgr
}

// EgCategoryGroupMgr open func
func EgCategoryGroupMgr(db *gorm.DB) *_EgCategoryGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgCategoryGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCategoryGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_category_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCategoryGroupMgr) GetTableName() string {
	return "eg_category_group"
}

// Get 获取
func (obj *_EgCategoryGroupMgr) Get() (result EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCategoryGroupMgr) Gets() (results []*EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCategory id_category获取 
func (obj *_EgCategoryGroupMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithIDGroup id_group获取 
func (obj *_EgCategoryGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}


// GetByOption 功能选项模式获取
func (obj *_EgCategoryGroupMgr) GetByOption(opts ...Option) (result EgCategoryGroup, err error) {
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
func (obj *_EgCategoryGroupMgr) GetByOptions(opts ...Option) (results []*EgCategoryGroup, err error) {
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


// GetFromIDCategory 通过id_category获取内容  
func (obj *_EgCategoryGroupMgr) GetFromIDCategory(idCategory uint32) (results []*EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCategory 批量唯一主键查找 
func (obj *_EgCategoryGroupMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error
	
	return
}
 
// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgCategoryGroupMgr) GetFromIDGroup(idGroup uint32) (results []*EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgCategoryGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCategoryGroupMgr) FetchByPrimaryKey(idCategory uint32 ,idGroup uint32 ) (result EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND id_group = ?", idCategory , idGroup).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCategory  获取多个内容
 func (obj *_EgCategoryGroupMgr) FetchIndexByIDCategory(idCategory uint32 ) (results []*EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDGroup  获取多个内容
 func (obj *_EgCategoryGroupMgr) FetchIndexByIDGroup(idGroup uint32 ) (results []*EgCategoryGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}
 

	

