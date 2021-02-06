package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgSupplierShopMgr struct {
	*_BaseMgr
}

// EgSupplierShopMgr open func
func EgSupplierShopMgr(db *gorm.DB) *_EgSupplierShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgSupplierShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSupplierShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_supplier_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSupplierShopMgr) GetTableName() string {
	return "eg_supplier_shop"
}

// Get 获取
func (obj *_EgSupplierShopMgr) Get() (result EgSupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSupplierShopMgr) Gets() (results []*EgSupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSupplier id_supplier获取 
func (obj *_EgSupplierShopMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithIDShop id_shop获取 
func (obj *_EgSupplierShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgSupplierShopMgr) GetByOption(opts ...Option) (result EgSupplierShop, err error) {
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
func (obj *_EgSupplierShopMgr) GetByOptions(opts ...Option) (results []*EgSupplierShop, err error) {
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


// GetFromIDSupplier 通过id_supplier获取内容  
func (obj *_EgSupplierShopMgr) GetFromIDSupplier(idSupplier uint32) (results []*EgSupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplier 批量唯一主键查找 
func (obj *_EgSupplierShopMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*EgSupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgSupplierShopMgr) GetFromIDShop(idShop uint32) (results []*EgSupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgSupplierShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgSupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSupplierShopMgr) FetchByPrimaryKey(idSupplier uint32 ,idShop uint32 ) (result EgSupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ? AND id_shop = ?", idSupplier , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgSupplierShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgSupplierShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

