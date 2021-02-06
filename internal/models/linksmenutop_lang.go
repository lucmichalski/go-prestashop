package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLinksmenutopLangMgr struct {
	*_BaseMgr
}

// EgLinksmenutopLangMgr open func
func EgLinksmenutopLangMgr(db *gorm.DB) *_EgLinksmenutopLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgLinksmenutopLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLinksmenutopLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_linksmenutop_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLinksmenutopLangMgr) GetTableName() string {
	return "eg_linksmenutop_lang"
}

// Get 获取
func (obj *_EgLinksmenutopLangMgr) Get() (result EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLinksmenutopLangMgr) Gets() (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLinksmenutop id_linksmenutop获取 
func (obj *_EgLinksmenutopLangMgr) WithIDLinksmenutop(idLinksmenutop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_linksmenutop"] = idLinksmenutop })
}

// WithIDLang id_lang获取 
func (obj *_EgLinksmenutopLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取 
func (obj *_EgLinksmenutopLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithLabel label获取 
func (obj *_EgLinksmenutopLangMgr) WithLabel(label string) Option {
	return optionFunc(func(o *options) { o.query["label"] = label })
}

// WithLink link获取 
func (obj *_EgLinksmenutopLangMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}


// GetByOption 功能选项模式获取
func (obj *_EgLinksmenutopLangMgr) GetByOption(opts ...Option) (result EgLinksmenutopLang, err error) {
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
func (obj *_EgLinksmenutopLangMgr) GetByOptions(opts ...Option) (results []*EgLinksmenutopLang, err error) {
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


// GetFromIDLinksmenutop 通过id_linksmenutop获取内容  
func (obj *_EgLinksmenutopLangMgr) GetFromIDLinksmenutop(idLinksmenutop uint32) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ?", idLinksmenutop).Find(&results).Error
	
	return
}

// GetBatchFromIDLinksmenutop 批量唯一主键查找 
func (obj *_EgLinksmenutopLangMgr) GetBatchFromIDLinksmenutop(idLinksmenutops []uint32) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop IN (?)", idLinksmenutops).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgLinksmenutopLangMgr) GetFromIDLang(idLang uint32) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgLinksmenutopLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgLinksmenutopLangMgr) GetFromIDShop(idShop uint32) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgLinksmenutopLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromLabel 通过label获取内容  
func (obj *_EgLinksmenutopLangMgr) GetFromLabel(label string) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("label = ?", label).Find(&results).Error
	
	return
}

// GetBatchFromLabel 批量唯一主键查找 
func (obj *_EgLinksmenutopLangMgr) GetBatchFromLabel(labels []string) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("label IN (?)", labels).Find(&results).Error
	
	return
}
 
// GetFromLink 通过link获取内容  
func (obj *_EgLinksmenutopLangMgr) GetFromLink(link string) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error
	
	return
}

// GetBatchFromLink 批量唯一主键查找 
func (obj *_EgLinksmenutopLangMgr) GetBatchFromLink(links []string) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 
 // FetchIndexByIDLinksmenutop  获取多个内容
 func (obj *_EgLinksmenutopLangMgr) FetchIndexByIDLinksmenutop(idLinksmenutop uint32 ,idLang uint32 ,idShop uint32 ) (results []*EgLinksmenutopLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ? AND id_lang = ? AND id_shop = ?", idLinksmenutop , idLang , idShop).Find(&results).Error
	
	return
}
 

	

