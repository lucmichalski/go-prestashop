package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgBadgeLangMgr struct {
	*_BaseMgr
}

// EgBadgeLangMgr open func
func EgBadgeLangMgr(db *gorm.DB) *_EgBadgeLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgBadgeLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgBadgeLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_badge_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgBadgeLangMgr) GetTableName() string {
	return "eg_badge_lang"
}

// Get 获取
func (obj *_EgBadgeLangMgr) Get() (result EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgBadgeLangMgr) Gets() (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDBadge id_badge获取 
func (obj *_EgBadgeLangMgr) WithIDBadge(idBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_badge"] = idBadge })
}

// WithIDLang id_lang获取 
func (obj *_EgBadgeLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgBadgeLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取 
func (obj *_EgBadgeLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithGroupName group_name获取 
func (obj *_EgBadgeLangMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}


// GetByOption 功能选项模式获取
func (obj *_EgBadgeLangMgr) GetByOption(opts ...Option) (result EgBadgeLang, err error) {
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
func (obj *_EgBadgeLangMgr) GetByOptions(opts ...Option) (results []*EgBadgeLang, err error) {
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


// GetFromIDBadge 通过id_badge获取内容  
func (obj *_EgBadgeLangMgr) GetFromIDBadge(idBadge int) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&results).Error
	
	return
}

// GetBatchFromIDBadge 批量唯一主键查找 
func (obj *_EgBadgeLangMgr) GetBatchFromIDBadge(idBadges []int) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge IN (?)", idBadges).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgBadgeLangMgr) GetFromIDLang(idLang int) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgBadgeLangMgr) GetBatchFromIDLang(idLangs []int) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgBadgeLangMgr) GetFromName(name string) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgBadgeLangMgr) GetBatchFromName(names []string) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgBadgeLangMgr) GetFromDescription(description string) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgBadgeLangMgr) GetBatchFromDescription(descriptions []string) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromGroupName 通过group_name获取内容  
func (obj *_EgBadgeLangMgr) GetFromGroupName(groupName string) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_name = ?", groupName).Find(&results).Error
	
	return
}

// GetBatchFromGroupName 批量唯一主键查找 
func (obj *_EgBadgeLangMgr) GetBatchFromGroupName(groupNames []string) (results []*EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_name IN (?)", groupNames).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgBadgeLangMgr) FetchByPrimaryKey(idBadge int ,idLang int ) (result EgBadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ? AND id_lang = ?", idBadge , idLang).Find(&result).Error
	
	return
}
 

 

	

