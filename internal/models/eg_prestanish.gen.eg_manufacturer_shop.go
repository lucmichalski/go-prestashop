package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgManufacturerShopMgr struct {
	*_BaseMgr
}

// EgManufacturerShopMgr open func
func EgManufacturerShopMgr(db *gorm.DB) *_EgManufacturerShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgManufacturerShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgManufacturerShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_manufacturer_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgManufacturerShopMgr) GetTableName() string {
	return "eg_manufacturer_shop"
}

// Get 获取
func (obj *_EgManufacturerShopMgr) Get() (result EgManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgManufacturerShopMgr) Gets() (results []*EgManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDManufacturer id_manufacturer获取 
func (obj *_EgManufacturerShopMgr) WithIDManufacturer(idManufacturer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_manufacturer"] = idManufacturer })
}

// WithIDShop id_shop获取 
func (obj *_EgManufacturerShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgManufacturerShopMgr) GetByOption(opts ...Option) (result EgManufacturerShop, err error) {
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
func (obj *_EgManufacturerShopMgr) GetByOptions(opts ...Option) (results []*EgManufacturerShop, err error) {
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


// GetFromIDManufacturer 通过id_manufacturer获取内容  
func (obj *_EgManufacturerShopMgr) GetFromIDManufacturer(idManufacturer uint32) (results []*EgManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ?", idManufacturer).Find(&results).Error
	
	return
}

// GetBatchFromIDManufacturer 批量唯一主键查找 
func (obj *_EgManufacturerShopMgr) GetBatchFromIDManufacturer(idManufacturers []uint32) (results []*EgManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer IN (?)", idManufacturers).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgManufacturerShopMgr) GetFromIDShop(idShop uint32) (results []*EgManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgManufacturerShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgManufacturerShopMgr) FetchByPrimaryKey(idManufacturer uint32 ,idShop uint32 ) (result EgManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_manufacturer = ? AND id_shop = ?", idManufacturer , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgManufacturerShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgManufacturerShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

