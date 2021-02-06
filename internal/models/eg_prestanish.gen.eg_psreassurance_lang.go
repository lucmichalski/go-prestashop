package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgPsreassuranceLangMgr struct {
	*_BaseMgr
}

// EgPsreassuranceLangMgr open func
func EgPsreassuranceLangMgr(db *gorm.DB) *_EgPsreassuranceLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgPsreassuranceLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPsreassuranceLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_psreassurance_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPsreassuranceLangMgr) GetTableName() string {
	return "eg_psreassurance_lang"
}

// Get 获取
func (obj *_EgPsreassuranceLangMgr) Get() (result EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPsreassuranceLangMgr) Gets() (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPsreassurance id_psreassurance获取 
func (obj *_EgPsreassuranceLangMgr) WithIDPsreassurance(idPsreassurance uint32) Option {
	return optionFunc(func(o *options) { o.query["id_psreassurance"] = idPsreassurance })
}

// WithIDLang id_lang获取 
func (obj *_EgPsreassuranceLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取 
func (obj *_EgPsreassuranceLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithTitle title获取 
func (obj *_EgPsreassuranceLangMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDescription description获取 
func (obj *_EgPsreassuranceLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithLink link获取 
func (obj *_EgPsreassuranceLangMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}


// GetByOption 功能选项模式获取
func (obj *_EgPsreassuranceLangMgr) GetByOption(opts ...Option) (result EgPsreassuranceLang, err error) {
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
func (obj *_EgPsreassuranceLangMgr) GetByOptions(opts ...Option) (results []*EgPsreassuranceLang, err error) {
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


// GetFromIDPsreassurance 通过id_psreassurance获取内容  
func (obj *_EgPsreassuranceLangMgr) GetFromIDPsreassurance(idPsreassurance uint32) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance = ?", idPsreassurance).Find(&results).Error
	
	return
}

// GetBatchFromIDPsreassurance 批量唯一主键查找 
func (obj *_EgPsreassuranceLangMgr) GetBatchFromIDPsreassurance(idPsreassurances []uint32) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance IN (?)", idPsreassurances).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgPsreassuranceLangMgr) GetFromIDLang(idLang uint32) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgPsreassuranceLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgPsreassuranceLangMgr) GetFromIDShop(idShop uint32) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgPsreassuranceLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromTitle 通过title获取内容  
func (obj *_EgPsreassuranceLangMgr) GetFromTitle(title string) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error
	
	return
}

// GetBatchFromTitle 批量唯一主键查找 
func (obj *_EgPsreassuranceLangMgr) GetBatchFromTitle(titles []string) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error
	
	return
}
 
// GetFromDescription 通过description获取内容  
func (obj *_EgPsreassuranceLangMgr) GetFromDescription(description string) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error
	
	return
}

// GetBatchFromDescription 批量唯一主键查找 
func (obj *_EgPsreassuranceLangMgr) GetBatchFromDescription(descriptions []string) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error
	
	return
}
 
// GetFromLink 通过link获取内容  
func (obj *_EgPsreassuranceLangMgr) GetFromLink(link string) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error
	
	return
}

// GetBatchFromLink 批量唯一主键查找 
func (obj *_EgPsreassuranceLangMgr) GetBatchFromLink(links []string) (results []*EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPsreassuranceLangMgr) FetchByPrimaryKey(idPsreassurance uint32 ,idLang uint32 ,idShop uint32 ) (result EgPsreassuranceLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance = ? AND id_lang = ? AND id_shop = ?", idPsreassurance , idLang , idShop).Find(&result).Error
	
	return
}
 

 

	

