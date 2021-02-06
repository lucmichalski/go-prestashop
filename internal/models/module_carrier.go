package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgModuleCarrierMgr struct {
	*_BaseMgr
}

// EgModuleCarrierMgr open func
func EgModuleCarrierMgr(db *gorm.DB) *_EgModuleCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("EgModuleCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgModuleCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_module_carrier"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgModuleCarrierMgr) GetTableName() string {
	return "eg_module_carrier"
}

// Get 获取
func (obj *_EgModuleCarrierMgr) Get() (result EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgModuleCarrierMgr) Gets() (results []*EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取 
func (obj *_EgModuleCarrierMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithIDShop id_shop获取 
func (obj *_EgModuleCarrierMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDReference id_reference获取 
func (obj *_EgModuleCarrierMgr) WithIDReference(idReference int) Option {
	return optionFunc(func(o *options) { o.query["id_reference"] = idReference })
}


// GetByOption 功能选项模式获取
func (obj *_EgModuleCarrierMgr) GetByOption(opts ...Option) (result EgModuleCarrier, err error) {
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
func (obj *_EgModuleCarrierMgr) GetByOptions(opts ...Option) (results []*EgModuleCarrier, err error) {
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
func (obj *_EgModuleCarrierMgr) GetFromIDModule(idModule uint32) (results []*EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&results).Error
	
	return
}

// GetBatchFromIDModule 批量唯一主键查找 
func (obj *_EgModuleCarrierMgr) GetBatchFromIDModule(idModules []uint32) (results []*EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgModuleCarrierMgr) GetFromIDShop(idShop uint32) (results []*EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgModuleCarrierMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDReference 通过id_reference获取内容  
func (obj *_EgModuleCarrierMgr) GetFromIDReference(idReference int) (results []*EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference = ?", idReference).Find(&results).Error
	
	return
}

// GetBatchFromIDReference 批量唯一主键查找 
func (obj *_EgModuleCarrierMgr) GetBatchFromIDReference(idReferences []int) (results []*EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference IN (?)", idReferences).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgModuleCarrierMgr) FetchByPrimaryKey(idModule uint32 ,idShop uint32 ,idReference int ) (result EgModuleCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ? AND id_shop = ? AND id_reference = ?", idModule , idShop , idReference).Find(&result).Error
	
	return
}
 

 

	

