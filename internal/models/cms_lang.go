package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCmsLangMgr struct {
	*_BaseMgr
}

// EgCmsLangMgr open func
func EgCmsLangMgr(db *gorm.DB) *_EgCmsLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgCmsLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCmsLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCmsLangMgr) GetTableName() string {
	return "eg_cms_lang"
}

// Get 获取
func (obj *_EgCmsLangMgr) Get() (result EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCmsLangMgr) Gets() (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCms id_cms获取 
func (obj *_EgCmsLangMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

// WithIDLang id_lang获取 
func (obj *_EgCmsLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取 
func (obj *_EgCmsLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithMetaTitle meta_title获取 
func (obj *_EgCmsLangMgr) WithMetaTitle(metaTitle string) Option {
	return optionFunc(func(o *options) { o.query["meta_title"] = metaTitle })
}

// WithHeadSeoTitle head_seo_title获取 
func (obj *_EgCmsLangMgr) WithHeadSeoTitle(headSeoTitle string) Option {
	return optionFunc(func(o *options) { o.query["head_seo_title"] = headSeoTitle })
}

// WithMetaDescription meta_description获取 
func (obj *_EgCmsLangMgr) WithMetaDescription(metaDescription string) Option {
	return optionFunc(func(o *options) { o.query["meta_description"] = metaDescription })
}

// WithMetaKeywords meta_keywords获取 
func (obj *_EgCmsLangMgr) WithMetaKeywords(metaKeywords string) Option {
	return optionFunc(func(o *options) { o.query["meta_keywords"] = metaKeywords })
}

// WithContent content获取 
func (obj *_EgCmsLangMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithLinkRewrite link_rewrite获取 
func (obj *_EgCmsLangMgr) WithLinkRewrite(linkRewrite string) Option {
	return optionFunc(func(o *options) { o.query["link_rewrite"] = linkRewrite })
}


// GetByOption 功能选项模式获取
func (obj *_EgCmsLangMgr) GetByOption(opts ...Option) (result EgCmsLang, err error) {
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
func (obj *_EgCmsLangMgr) GetByOptions(opts ...Option) (results []*EgCmsLang, err error) {
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


// GetFromIDCms 通过id_cms获取内容  
func (obj *_EgCmsLangMgr) GetFromIDCms(idCms uint32) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error
	
	return
}

// GetBatchFromIDCms 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromIDCms(idCmss []uint32) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCmsLangMgr) GetFromIDLang(idLang uint32) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCmsLangMgr) GetFromIDShop(idShop uint32) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromMetaTitle 通过meta_title获取内容  
func (obj *_EgCmsLangMgr) GetFromMetaTitle(metaTitle string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title = ?", metaTitle).Find(&results).Error
	
	return
}

// GetBatchFromMetaTitle 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromMetaTitle(metaTitles []string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_title IN (?)", metaTitles).Find(&results).Error
	
	return
}
 
// GetFromHeadSeoTitle 通过head_seo_title获取内容  
func (obj *_EgCmsLangMgr) GetFromHeadSeoTitle(headSeoTitle string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("head_seo_title = ?", headSeoTitle).Find(&results).Error
	
	return
}

// GetBatchFromHeadSeoTitle 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromHeadSeoTitle(headSeoTitles []string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("head_seo_title IN (?)", headSeoTitles).Find(&results).Error
	
	return
}
 
// GetFromMetaDescription 通过meta_description获取内容  
func (obj *_EgCmsLangMgr) GetFromMetaDescription(metaDescription string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description = ?", metaDescription).Find(&results).Error
	
	return
}

// GetBatchFromMetaDescription 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromMetaDescription(metaDescriptions []string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_description IN (?)", metaDescriptions).Find(&results).Error
	
	return
}
 
// GetFromMetaKeywords 通过meta_keywords获取内容  
func (obj *_EgCmsLangMgr) GetFromMetaKeywords(metaKeywords string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords = ?", metaKeywords).Find(&results).Error
	
	return
}

// GetBatchFromMetaKeywords 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromMetaKeywords(metaKeywordss []string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("meta_keywords IN (?)", metaKeywordss).Find(&results).Error
	
	return
}
 
// GetFromContent 通过content获取内容  
func (obj *_EgCmsLangMgr) GetFromContent(content string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error
	
	return
}

// GetBatchFromContent 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromContent(contents []string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error
	
	return
}
 
// GetFromLinkRewrite 通过link_rewrite获取内容  
func (obj *_EgCmsLangMgr) GetFromLinkRewrite(linkRewrite string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite = ?", linkRewrite).Find(&results).Error
	
	return
}

// GetBatchFromLinkRewrite 批量唯一主键查找 
func (obj *_EgCmsLangMgr) GetBatchFromLinkRewrite(linkRewrites []string) (results []*EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link_rewrite IN (?)", linkRewrites).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCmsLangMgr) FetchByPrimaryKey(idCms uint32 ,idLang uint32 ,idShop uint32 ) (result EgCmsLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ? AND id_lang = ? AND id_shop = ?", idCms , idLang , idShop).Find(&result).Error
	
	return
}
 

 

	

