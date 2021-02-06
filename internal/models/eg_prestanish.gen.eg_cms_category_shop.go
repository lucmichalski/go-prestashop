package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgCmsCategoryShopMgr struct {
	*_BaseMgr
}

// EgCmsCategoryShopMgr open func
func EgCmsCategoryShopMgr(db *gorm.DB) *_EgCmsCategoryShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgCmsCategoryShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCmsCategoryShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms_category_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCmsCategoryShopMgr) GetTableName() string {
	return "eg_cms_category_shop"
}

// Get 获取
func (obj *_EgCmsCategoryShopMgr) Get() (result EgCmsCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCmsCategoryShopMgr) Gets() (results []*EgCmsCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCmsCategory id_cms_category获取 
func (obj *_EgCmsCategoryShopMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

// WithIDShop id_shop获取 
func (obj *_EgCmsCategoryShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgCmsCategoryShopMgr) GetByOption(opts ...Option) (result EgCmsCategoryShop, err error) {
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
func (obj *_EgCmsCategoryShopMgr) GetByOptions(opts ...Option) (results []*EgCmsCategoryShop, err error) {
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


// GetFromIDCmsCategory 通过id_cms_category获取内容  
func (obj *_EgCmsCategoryShopMgr) GetFromIDCmsCategory(idCmsCategory uint32) (results []*EgCmsCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCmsCategory 批量唯一主键查找 
func (obj *_EgCmsCategoryShopMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*EgCmsCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCmsCategoryShopMgr) GetFromIDShop(idShop uint32) (results []*EgCmsCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCmsCategoryShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCmsCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCmsCategoryShopMgr) FetchByPrimaryKey(idCmsCategory uint32 ,idShop uint32 ) (result EgCmsCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ? AND id_shop = ?", idCmsCategory , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgCmsCategoryShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgCmsCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

