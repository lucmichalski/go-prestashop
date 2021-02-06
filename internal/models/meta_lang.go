package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgMetaLangMgr struct {
	*_BaseMgr
}

// EgMetaLangMgr open func
func EgMetaLangMgr(db *gorm.DB) *_EgMetaLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgMetaLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgMetaLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_meta_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgMetaLangMgr) GetTableName() string {
	return "eg_meta_lang"
}

// Get 获取
func (obj *_EgMetaLangMgr) Get() (result EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgMetaLangMgr) Gets() (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMeta id_meta获取 
func (obj *_EgMetaLangMgr) WithIDMeta(idMeta uint32) Option {
	return optionFunc(func(o *options) { o.query["id_meta"] = idMeta })
}

// WithIDShop id_shop获取 
func (obj *_EgMetaLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取 
func (obj *_EgMetaLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithTitle title获取 
func (obj *_EgMetaLangMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDescription description获取 
func (obj *_EgMetaLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithKeywords keywords获取 
func (obj *_EgMetaLangMgr) WithKeywords(keywords string) Option {
	return optionFunc(func(o *options) { o.query["keywords"] = keywords })
}

// WithURLRewrite url_rewrite获取 
func (obj *_EgMetaLangMgr) WithURLRewrite(urlRewrite string) Option {
	return optionFunc(func(o *options) { o.query["url_rewrite"] = urlRewrite })
}


// GetByOption 功能选项模式获取
func (obj *_EgMetaLangMgr) GetByOption(opts ...Option) (result EgMetaLang, err error) {
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
func (obj *_EgMetaLangMgr) GetByOptions(opts ...Option) (results []*EgMetaLang, err error) {
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


// GetFromIDMeta 通过id_meta获取内容  
func (obj *_EgMetaLangMgr) GetFromIDMeta(idMeta uint32) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta = ?", idMeta).Find(&results).Error
	
	return
}

// GetBatchFromIDMeta 批量唯一主键查找 
func (obj *_EgMetaLangMgr) GetBatchFromIDMeta(idMetas []uint32) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta IN (?)", idMetas).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgMetaLangMgr) GetFromIDShop(idShop uint32) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgMetaLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgMetaLangMgr) GetFromIDLang(idLang uint32) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgMetaLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromTitle 通过title获取内容  
func (obj *_EgMetaLangMgr) GetFromTitle(title string) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error
	
	return
}

// GetBatchFromTitle 批量唯一主键查找 
func (obj *_EgMetaLangMgr) GetBatchFromTitle(titles []string) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgMetaLangMgr) GetFromDescription(description string) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgMetaLangMgr) GetBatchFromDescription(descriptions []string) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromKeywords 通过keywords获取内容  
func (obj *_EgMetaLangMgr) GetFromKeywords(keywords string) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords = ?", keywords).Find(&results).Error
	
	return
}

// GetBatchFromKeywords 批量唯一主键查找 
func (obj *_EgMetaLangMgr) GetBatchFromKeywords(keywordss []string) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keywords IN (?)", keywordss).Find(&results).Error
	
	return
}
 
// GetFromURLRewrite 通过url_rewrite获取内容  
func (obj *_EgMetaLangMgr) GetFromURLRewrite(urlRewrite string) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_rewrite = ?", urlRewrite).Find(&results).Error
	
	return
}

// GetBatchFromURLRewrite 批量唯一主键查找 
func (obj *_EgMetaLangMgr) GetBatchFromURLRewrite(urlRewrites []string) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url_rewrite IN (?)", urlRewrites).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgMetaLangMgr) FetchByPrimaryKey(idMeta uint32 ,idShop uint32 ,idLang uint32 ) (result EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_meta = ? AND id_shop = ? AND id_lang = ?", idMeta , idShop , idLang).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgMetaLangMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDLang  获取多个内容
 func (obj *_EgMetaLangMgr) FetchIndexByIDLang(idLang uint32 ) (results []*EgMetaLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}
 

	

