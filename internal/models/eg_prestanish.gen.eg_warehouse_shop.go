package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgWarehouseShopMgr struct {
	*_BaseMgr
}

// EgWarehouseShopMgr open func
func EgWarehouseShopMgr(db *gorm.DB) *_EgWarehouseShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgWarehouseShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgWarehouseShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_warehouse_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgWarehouseShopMgr) GetTableName() string {
	return "eg_warehouse_shop"
}

// Get 获取
func (obj *_EgWarehouseShopMgr) Get() (result EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgWarehouseShopMgr) Gets() (results []*EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDShop id_shop获取 
func (obj *_EgWarehouseShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDWarehouse id_warehouse获取 
func (obj *_EgWarehouseShopMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}


// GetByOption 功能选项模式获取
func (obj *_EgWarehouseShopMgr) GetByOption(opts ...Option) (result EgWarehouseShop, err error) {
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
func (obj *_EgWarehouseShopMgr) GetByOptions(opts ...Option) (results []*EgWarehouseShop, err error) {
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


// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgWarehouseShopMgr) GetFromIDShop(idShop uint32) (results []*EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgWarehouseShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDWarehouse 通过id_warehouse获取内容  
func (obj *_EgWarehouseShopMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找 
func (obj *_EgWarehouseShopMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgWarehouseShopMgr) FetchByPrimaryKey(idShop uint32 ,idWarehouse uint32 ) (result EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_warehouse = ?", idShop , idWarehouse).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgWarehouseShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDWarehouse  获取多个内容
 func (obj *_EgWarehouseShopMgr) FetchIndexByIDWarehouse(idWarehouse uint32 ) (results []*EgWarehouseShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}
 

	

