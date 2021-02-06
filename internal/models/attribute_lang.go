package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgAttributeLangMgr struct {
	*_BaseMgr
}

// EgAttributeLangMgr open func
func EgAttributeLangMgr(db *gorm.DB) *_EgAttributeLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgAttributeLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAttributeLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attribute_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAttributeLangMgr) GetTableName() string {
	return "eg_attribute_lang"
}

// Get 获取
func (obj *_EgAttributeLangMgr) Get() (result EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAttributeLangMgr) Gets() (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttribute id_attribute获取 
func (obj *_EgAttributeLangMgr) WithIDAttribute(idAttribute int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

// WithIDLang id_lang获取 
func (obj *_EgAttributeLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgAttributeLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgAttributeLangMgr) GetByOption(opts ...Option) (result EgAttributeLang, err error) {
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
func (obj *_EgAttributeLangMgr) GetByOptions(opts ...Option) (results []*EgAttributeLang, err error) {
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


// GetFromIDAttribute 通过id_attribute获取内容  
func (obj *_EgAttributeLangMgr) GetFromIDAttribute(idAttribute int) (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error
	
	return
}

// GetBatchFromIDAttribute 批量唯一主键查找 
func (obj *_EgAttributeLangMgr) GetBatchFromIDAttribute(idAttributes []int) (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgAttributeLangMgr) GetFromIDLang(idLang int) (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgAttributeLangMgr) GetBatchFromIDLang(idLangs []int) (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgAttributeLangMgr) GetFromName(name string) (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgAttributeLangMgr) GetBatchFromName(names []string) (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAttributeLangMgr) FetchByPrimaryKey(idAttribute int ,idLang int ) (result EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_lang = ?", idAttribute , idLang).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDXEB57414F7A4F53DC  获取多个内容
 func (obj *_EgAttributeLangMgr) FetchIndexByIDXEB57414F7A4F53DC(idAttribute int ) (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDXEB57414FBA299860  获取多个内容
 func (obj *_EgAttributeLangMgr) FetchIndexByIDXEB57414FBA299860(idLang int ) (results []*EgAttributeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}
 

	

