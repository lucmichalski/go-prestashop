package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgStoreShopMgr struct {
	*_BaseMgr
}

// EgStoreShopMgr open func
func EgStoreShopMgr(db *gorm.DB) *_EgStoreShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgStoreShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgStoreShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_store_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgStoreShopMgr) GetTableName() string {
	return "eg_store_shop"
}

// Get 获取
func (obj *_EgStoreShopMgr) Get() (result EgStoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgStoreShopMgr) Gets() (results []*EgStoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStore id_store获取 
func (obj *_EgStoreShopMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

// WithIDShop id_shop获取 
func (obj *_EgStoreShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgStoreShopMgr) GetByOption(opts ...Option) (result EgStoreShop, err error) {
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
func (obj *_EgStoreShopMgr) GetByOptions(opts ...Option) (results []*EgStoreShop, err error) {
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


// GetFromIDStore 通过id_store获取内容  
func (obj *_EgStoreShopMgr) GetFromIDStore(idStore uint32) (results []*EgStoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&results).Error
	
	return
}

// GetBatchFromIDStore 批量唯一主键查找 
func (obj *_EgStoreShopMgr) GetBatchFromIDStore(idStores []uint32) (results []*EgStoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgStoreShopMgr) GetFromIDShop(idShop uint32) (results []*EgStoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgStoreShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgStoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgStoreShopMgr) FetchByPrimaryKey(idStore uint32 ,idShop uint32 ) (result EgStoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ? AND id_shop = ?", idStore , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgStoreShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgStoreShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

