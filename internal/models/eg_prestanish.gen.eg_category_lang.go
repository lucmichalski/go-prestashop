package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCategoryLangMgr struct {
	*_BaseMgr
}

// EgCategoryLangMgr open func
func EgCategoryLangMgr(db *gorm.DB) *_EgCategoryLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgCategoryLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCategoryLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_category_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCategoryLangMgr) GetTableName() string {
	return "eg_category_lang"
}

// Get 获取
func (obj *_EgCategoryLangMgr) Get() (result EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCategoryLangMgr) Gets() (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCategory id_category获取 
func (obj *_EgCategoryLangMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithIDShop id_shop获取 
func (obj *_EgCategoryLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取 
func (obj *_EgCategoryLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgCategoryLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取 
func (obj *_EgCategoryLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithLinkRewrite link_rewrite获取 
func (obj *_EgCategoryLangMgr) WithLinkRewrite(linkRewrite string) Option {
	return optionFunc(func(o *options) { o.query["link_rewrite"] = linkRewrite })
}

// WithMetaTitle meta_title获取 
func (obj *_EgCategoryLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// WithMetaKeywords meta_keywords获取 
func (obj *_EgCategoryLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithMetaDescription meta_description获取 
func (obj *_EgCategoryLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}


// GetByOption 功能选项模式获取
func (obj *_EgCategoryLangMgr) GetByOption(opts ...Option) (result EgCategoryLang, err error) {
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
func (obj *_EgCategoryLangMgr) GetByOptions(opts ...Option) (results []*EgCategoryLang, err error) {
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
func (obj *_EgCategoryLangMgr) GetFromIDCategory(idCategory uint32) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCategory 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCategoryLangMgr) GetFromIDShop(idShop uint32) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCategoryLangMgr) GetFromIDLang(idLang uint32) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCategoryLangMgr) GetFromName(name string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromName(names []string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgCategoryLangMgr) GetFromDescription(description string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromDescription(descriptions []string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromLinkRewrite 通过link_rewrite获取内容  
func (obj *_EgCategoryLangMgr) GetFromLinkRewrite(linkRewrite string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite = ?", linkRewrite).Find(&results).Error
	
	return
}

// GetBatchFromLinkRewrite 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromLinkRewrite(linkRewrites []string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite IN (?)", linkRewrites).Find(&results).Error
	
	return
}
 
// GetFromMetaTitle 通过meta_title获取内容  
func (obj *_EgCategoryLangMgr) GetFromMetaTitle(metaTitle string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error
	
	return
}

// GetBatchFromMetaTitle 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error
	
	return
}
 
// GetFromMetaKeywords 通过meta_keywords获取内容  
func (obj *_EgCategoryLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error
	
	return
}

// GetBatchFromMetaKeywords 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error
	
	return
}
 
// GetFromMetaDescription 通过meta_description获取内容  
func (obj *_EgCategoryLangMgr) GetFromMetaDescription(metaDescription string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error
	
	return
}

// GetBatchFromMetaDescription 批量唯一主键查找 
func (obj *_EgCategoryLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCategoryLangMgr) FetchByPrimaryKey(idCategory uint32 ,idShop uint32 ,idLang uint32 ) (result EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND id_shop = ? AND id_lang = ?", idCategory , idShop , idLang).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByCategoryName  获取多个内容
 func (obj *_EgCategoryLangMgr) FetchIndexByCategoryName(name string ) (results []*EgCategoryLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}
 

	

