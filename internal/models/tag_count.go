package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgTagCountMgr struct {
	*_BaseMgr
}

// EgTagCountMgr open func
func EgTagCountMgr(db *gorm.DB) *_EgTagCountMgr {
	if db == nil {
		panic(fmt.Errorf("EgTagCountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgTagCountMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tag_count"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgTagCountMgr) GetTableName() string {
	return "eg_tag_count"
}

// Get 获取
func (obj *_EgTagCountMgr) Get() (result EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgTagCountMgr) Gets() (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGroup id_group获取 
func (obj *_EgTagCountMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// WithIDTag id_tag获取 
func (obj *_EgTagCountMgr) WithIDTag(idTag uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tag"] = idTag })
}

// WithIDLang id_lang获取 
func (obj *_EgTagCountMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取 
func (obj *_EgTagCountMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithCounter counter获取 
func (obj *_EgTagCountMgr) WithCounter(counter uint32) Option {
	return optionFunc(func(o *options) { o.query["counter"] = counter })
}


// GetByOption 功能选项模式获取
func (obj *_EgTagCountMgr) GetByOption(opts ...Option) (result EgTagCount, err error) {
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
func (obj *_EgTagCountMgr) GetByOptions(opts ...Option) (results []*EgTagCount, err error) {
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


// GetFromIDGroup 通过id_group获取内容  
func (obj *_EgTagCountMgr) GetFromIDGroup(idGroup uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDGroup 批量唯一主键查找 
func (obj *_EgTagCountMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error
	
	return
}
 
// GetFromIDTag 通过id_tag获取内容  
func (obj *_EgTagCountMgr) GetFromIDTag(idTag uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&results).Error
	
	return
}

// GetBatchFromIDTag 批量唯一主键查找 
func (obj *_EgTagCountMgr) GetBatchFromIDTag(idTags []uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag IN (?)", idTags).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgTagCountMgr) GetFromIDLang(idLang uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgTagCountMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgTagCountMgr) GetFromIDShop(idShop uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgTagCountMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromCounter 通过counter获取内容  
func (obj *_EgTagCountMgr) GetFromCounter(counter uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter = ?", counter).Find(&results).Error
	
	return
}

// GetBatchFromCounter 批量唯一主键查找 
func (obj *_EgTagCountMgr) GetBatchFromCounter(counters []uint32) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter IN (?)", counters).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgTagCountMgr) FetchByPrimaryKey(idGroup uint32 ,idTag uint32 ) (result EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_tag = ?", idGroup , idTag).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDGroup  获取多个内容
 func (obj *_EgTagCountMgr) FetchIndexByIDGroup(idGroup uint32 ,idLang uint32 ,idShop uint32 ,counter uint32 ) (results []*EgTagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_lang = ? AND id_shop = ? AND counter = ?", idGroup , idLang , idShop , counter).Find(&results).Error
	
	return
}
 

	
