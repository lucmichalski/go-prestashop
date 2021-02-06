package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCmsShopMgr struct {
	*_BaseMgr
}

// EgCmsShopMgr open func
func EgCmsShopMgr(db *gorm.DB) *_EgCmsShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgCmsShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCmsShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCmsShopMgr) GetTableName() string {
	return "eg_cms_shop"
}

// Get 获取
func (obj *_EgCmsShopMgr) Get() (result EgCmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCmsShopMgr) Gets() (results []*EgCmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCms id_cms获取 
func (obj *_EgCmsShopMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

// WithIDShop id_shop获取 
func (obj *_EgCmsShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgCmsShopMgr) GetByOption(opts ...Option) (result EgCmsShop, err error) {
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
func (obj *_EgCmsShopMgr) GetByOptions(opts ...Option) (results []*EgCmsShop, err error) {
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
func (obj *_EgCmsShopMgr) GetFromIDCms(idCms uint32) (results []*EgCmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error
	
	return
}

// GetBatchFromIDCms 批量唯一主键查找 
func (obj *_EgCmsShopMgr) GetBatchFromIDCms(idCmss []uint32) (results []*EgCmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCmsShopMgr) GetFromIDShop(idShop uint32) (results []*EgCmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCmsShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCmsShopMgr) FetchByPrimaryKey(idCms uint32 ,idShop uint32 ) (result EgCmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ? AND id_shop = ?", idCms , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgCmsShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgCmsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

