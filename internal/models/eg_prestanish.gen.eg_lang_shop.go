package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLangShopMgr struct {
	*_BaseMgr
}

// EgLangShopMgr open func
func EgLangShopMgr(db *gorm.DB) *_EgLangShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgLangShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLangShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_lang_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLangShopMgr) GetTableName() string {
	return "eg_lang_shop"
}

// Get 获取
func (obj *_EgLangShopMgr) Get() (result EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLangShopMgr) Gets() (results []*EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLang id_lang获取 
func (obj *_EgLangShopMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取 
func (obj *_EgLangShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgLangShopMgr) GetByOption(opts ...Option) (result EgLangShop, err error) {
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
func (obj *_EgLangShopMgr) GetByOptions(opts ...Option) (results []*EgLangShop, err error) {
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


// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgLangShopMgr) GetFromIDLang(idLang int) (results []*EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgLangShopMgr) GetBatchFromIDLang(idLangs []int) (results []*EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgLangShopMgr) GetFromIDShop(idShop int) (results []*EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgLangShopMgr) GetBatchFromIDShop(idShops []int) (results []*EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLangShopMgr) FetchByPrimaryKey(idLang int ,idShop int ) (result EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ? AND id_shop = ?", idLang , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDXA5D79262BA299860  获取多个内容
 func (obj *_EgLangShopMgr) FetchIndexByIDXA5D79262BA299860(idLang int ) (results []*EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDXA5D79262274A50A0  获取多个内容
 func (obj *_EgLangShopMgr) FetchIndexByIDXA5D79262274A50A0(idShop int ) (results []*EgLangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

