package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCustomizationFieldLangMgr struct {
	*_BaseMgr
}

// EgCustomizationFieldLangMgr open func
func EgCustomizationFieldLangMgr(db *gorm.DB) *_EgCustomizationFieldLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomizationFieldLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomizationFieldLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customization_field_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomizationFieldLangMgr) GetTableName() string {
	return "eg_customization_field_lang"
}

// Get 获取
func (obj *_EgCustomizationFieldLangMgr) Get() (result EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomizationFieldLangMgr) Gets() (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomizationField id_customization_field获取 
func (obj *_EgCustomizationFieldLangMgr) WithIDCustomizationField(idCustomizationField uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization_field"] = idCustomizationField })
}

// WithIDLang id_lang获取 
func (obj *_EgCustomizationFieldLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取 
func (obj *_EgCustomizationFieldLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithName name获取 
func (obj *_EgCustomizationFieldLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomizationFieldLangMgr) GetByOption(opts ...Option) (result EgCustomizationFieldLang, err error) {
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
func (obj *_EgCustomizationFieldLangMgr) GetByOptions(opts ...Option) (results []*EgCustomizationFieldLang, err error) {
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


// GetFromIDCustomizationField 通过id_customization_field获取内容  
func (obj *_EgCustomizationFieldLangMgr) GetFromIDCustomizationField(idCustomizationField uint32) (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field = ?", idCustomizationField).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomizationField 批量唯一主键查找 
func (obj *_EgCustomizationFieldLangMgr) GetBatchFromIDCustomizationField(idCustomizationFields []uint32) (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field IN (?)", idCustomizationFields).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCustomizationFieldLangMgr) GetFromIDLang(idLang uint32) (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCustomizationFieldLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCustomizationFieldLangMgr) GetFromIDShop(idShop uint32) (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCustomizationFieldLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCustomizationFieldLangMgr) GetFromName(name string) (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCustomizationFieldLangMgr) GetBatchFromName(names []string) (results []*EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomizationFieldLangMgr) FetchByPrimaryKey(idCustomizationField uint32 ,idLang uint32 ,idShop uint32 ) (result EgCustomizationFieldLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization_field = ? AND id_lang = ? AND id_shop = ?", idCustomizationField , idLang , idShop).Find(&result).Error
	
	return
}
 

 

	

