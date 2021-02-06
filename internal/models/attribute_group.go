package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgAttributeGroupMgr struct {
	*_BaseMgr
}

// EgAttributeGroupMgr open func
func EgAttributeGroupMgr(db *gorm.DB) *_EgAttributeGroupMgr {
	if db == nil {
		panic(fmt.Errorf("EgAttributeGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAttributeGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attribute_group"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAttributeGroupMgr) GetTableName() string {
	return "eg_attribute_group"
}

// Get 获取
func (obj *_EgAttributeGroupMgr) Get() (result EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAttributeGroupMgr) Gets() (results []*EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeGroup id_attribute_group获取 
func (obj *_EgAttributeGroupMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIsColorGroup is_color_group获取 
func (obj *_EgAttributeGroupMgr) WithIsColorGroup(isColorGroup bool) Option {
	return optionFunc(func(o *options) { o.query["is_color_group"] = isColorGroup })
}

// WithGroupType group_type获取 
func (obj *_EgAttributeGroupMgr) WithGroupType(groupType string) Option {
	return optionFunc(func(o *options) { o.query["group_type"] = groupType })
}

// WithPosition position获取 
func (obj *_EgAttributeGroupMgr) WithPosition(position int) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}


// GetByOption 功能选项模式获取
func (obj *_EgAttributeGroupMgr) GetByOption(opts ...Option) (result EgAttributeGroup, err error) {
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
func (obj *_EgAttributeGroupMgr) GetByOptions(opts ...Option) (results []*EgAttributeGroup, err error) {
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
func (obj *_EgAttributeGroupMgr)  GetFromIDAttributeGroup(idAttributeGroup int) (result EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error
	
	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找 
func (obj *_EgAttributeGroupMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error
	
	return
}
 
// GetFromIsColorGroup 通过is_color_group获取内容  
func (obj *_EgAttributeGroupMgr) GetFromIsColorGroup(isColorGroup bool) (results []*EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_color_group = ?", isColorGroup).Find(&results).Error
	
	return
}

// GetBatchFromIsColorGroup 批量唯一主键查找 
func (obj *_EgAttributeGroupMgr) GetBatchFromIsColorGroup(isColorGroups []bool) (results []*EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_color_group IN (?)", isColorGroups).Find(&results).Error
	
	return
}
 
// GetFromGroupType 通过group_type获取内容  
func (obj *_EgAttributeGroupMgr) GetFromGroupType(groupType string) (results []*EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_type = ?", groupType).Find(&results).Error
	
	return
}

// GetBatchFromGroupType 批量唯一主键查找 
func (obj *_EgAttributeGroupMgr) GetBatchFromGroupType(groupTypes []string) (results []*EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_type IN (?)", groupTypes).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgAttributeGroupMgr) GetFromPosition(position int) (results []*EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgAttributeGroupMgr) GetBatchFromPosition(positions []int) (results []*EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAttributeGroupMgr) FetchByPrimaryKey(idAttributeGroup int ) (result EgAttributeGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&result).Error
	
	return
}
 

 

	

