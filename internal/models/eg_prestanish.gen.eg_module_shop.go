package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgModuleShopMgr struct {
	*_BaseMgr
}

// EgModuleShopMgr open func
func EgModuleShopMgr(db *gorm.DB) *_EgModuleShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgModuleShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgModuleShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_module_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgModuleShopMgr) GetTableName() string {
	return "eg_module_shop"
}

// Get 获取
func (obj *_EgModuleShopMgr) Get() (result EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgModuleShopMgr) Gets() (results []*EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取 
func (obj *_EgModuleShopMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDShop id_shop获取 
func (obj *_EgModuleShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithEnableDevice enable_device获取 
func (obj *_EgModuleShopMgr) WithEnableDevice(enableDevice bool) Option {
	return optionFunc(func(o *options) { o.query["enable_device"] = enableDevice })
}


// GetByOption 功能选项模式获取
func (obj *_EgModuleShopMgr) GetByOption(opts ...Option) (result EgModuleShop, err error) {
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
func (obj *_EgModuleShopMgr) GetByOptions(opts ...Option) (results []*EgModuleShop, err error) {
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


// GetFromIDModule 通过id_module获取内容  
func (obj *_EgModuleShopMgr) GetFromIDModule(idModule uint32) (results []*EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}

// GetBatchFromIDModule 批量唯一主键查找 
func (obj *_EgModuleShopMgr) GetBatchFromIDModule(idModules []uint32) (results []*EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgModuleShopMgr) GetFromIDShop(idShop uint32) (results []*EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgModuleShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromEnableDevice 通过enable_device获取内容  
func (obj *_EgModuleShopMgr) GetFromEnableDevice(enableDevice bool) (results []*EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enable_device = ?", enableDevice).Find(&results).Error
	
	return
}

// GetBatchFromEnableDevice 批量唯一主键查找 
func (obj *_EgModuleShopMgr) GetBatchFromEnableDevice(enableDevices []bool) (results []*EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enable_device IN (?)", enableDevices).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgModuleShopMgr) FetchByPrimaryKey(idModule uint32 ,idShop uint32 ) (result EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ?", idModule , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgModuleShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgModuleShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

