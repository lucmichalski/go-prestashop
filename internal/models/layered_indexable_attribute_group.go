package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgLayeredIndexableAttributeGroupMgr struct {
	*_BaseMgr
}

// EgLayeredIndexableAttributeGroupMgr open func
func EgLayeredIndexableAttributeGroupMgr(db *gorm.DB) *_EgLayeredIndexableAttributeGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgLayeredIndexableAttributeGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLayeredIndexableAttributeGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_indexable_attribute_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLayeredIndexableAttributeGroupMgr) GetTableName() string {
	return "eg_layered_indexable_attribute_group"
}

// Get 获取
func (obj *_EgLayeredIndexableAttributeGroupMgr) Get() (result EgLayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLayeredIndexableAttributeGroupMgr) Gets() (results []*EgLayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeGroup id_attribute_group获取 
func (obj *_EgLayeredIndexableAttributeGroupMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIndexable indexable获取 
func (obj *_EgLayeredIndexableAttributeGroupMgr) WithIndexable(indexable bool) Option {
	return optionFunc(func(o *options) { o.query["indexable"] = indexable })
}


// GetByOption 功能选项模式获取
func (obj *_EgLayeredIndexableAttributeGroupMgr) GetByOption(opts ...Option) (result EgLayeredIndexableAttributeGroup, err error) {
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
func (obj *_EgLayeredIndexableAttributeGroupMgr) GetByOptions(opts ...Option) (results []*EgLayeredIndexableAttributeGroup, err error) {
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


// GetFromIDAttributeGroup 通过id_attribute_group获取内容  
func (obj *_EgLayeredIndexableAttributeGroupMgr)  GetFromIDAttributeGroup(idAttributeGroup int) (result EgLayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error
	
	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找 
func (obj *_EgLayeredIndexableAttributeGroupMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*EgLayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error
	
	return
}
 
// GetFromIndexable 通过indexable获取内容  
func (obj *_EgLayeredIndexableAttributeGroupMgr) GetFromIndexable(indexable bool) (results []*EgLayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable = ?", indexable).Find(&results).Error
	
	return
}

// GetBatchFromIndexable 批量唯一主键查找 
func (obj *_EgLayeredIndexableAttributeGroupMgr) GetBatchFromIndexable(indexables []bool) (results []*EgLayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexable IN (?)", indexables).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLayeredIndexableAttributeGroupMgr) FetchByPrimaryKey(idAttributeGroup int ) (result EgLayeredIndexableAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error
	
	return
}
 

 

	

