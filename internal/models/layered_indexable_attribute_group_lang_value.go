package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLayeredIndexableAttributeGroupLangValueMgr struct {
	*_BaseMgr
}

// EgLayeredIndexableAttributeGroupLangValueMgr open func
func EgLayeredIndexableAttributeGroupLangValueMgr(db *gorm.DB) *_EgLayeredIndexableAttributeGroupLangValueMgr {
	if db == nil {
		panic(fmt.Errorf("EgLayeredIndexableAttributeGroupLangValueMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLayeredIndexableAttributeGroupLangValueMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_indexable_attribute_group_lang_value"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetTableName() string {
	return "eg_layered_indexable_attribute_group_lang_value"
}

// Get 获取
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) Get() (result EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) Gets() (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeGroup id_attribute_group获取 
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIDLang id_lang获取 
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithURLName url_name获取 
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) WithURLName(urlName string) Option {
	return optionFunc(func(o *options) { o.query["url_name"] = urlName })
}

// WithMetaTitle meta_title获取 
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}


// GetByOption 功能选项模式获取
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetByOption(opts ...Option) (result EgLayeredIndexableAttributeGroupLangValue, err error) {
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
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetByOptions(opts ...Option) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
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
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetFromIDAttributeGroup(idAttributeGroup int) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找 
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetFromIDLang(idLang int) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetBatchFromIDLang(idLangs []int) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromURLName 通过url_name获取内容  
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetFromURLName(urlName string) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name = ?", urlName).Find(&results).Error
	
	return
}

// GetBatchFromURLName 批量唯一主键查找 
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetBatchFromURLName(urlNames []string) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_name IN (?)", urlNames).Find(&results).Error
	
	return
}
 
// GetFromMetaTitle 通过meta_title获取内容  
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetFromMetaTitle(metaTitle string) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error
	
	return
}

// GetBatchFromMetaTitle 批量唯一主键查找 
func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLayeredIndexableAttributeGroupLangValueMgr) FetchByPrimaryKey(idAttributeGroup int ,idLang int ) (result EgLayeredIndexableAttributeGroupLangValue, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ? AND id_lang = ?", idAttributeGroup , idLang).Find(&result).Error
	
	return
}
 

 

	

