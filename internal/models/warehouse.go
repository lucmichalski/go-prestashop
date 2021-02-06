package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgWarehouseMgr struct {
	*_BaseMgr
}

// EgWarehouseMgr open func
func EgWarehouseMgr(db *gorm.DB) *_EgWarehouseMgr {
	if db == nil {
		panic(fmt.Errorf("EgWarehouseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgWarehouseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_warehouse"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgWarehouseMgr) GetTableName() string {
	return "eg_warehouse"
}

// Get 获取
func (obj *_EgWarehouseMgr) Get() (result EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgWarehouseMgr) Gets() (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWarehouse id_warehouse获取 
func (obj *_EgWarehouseMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}

// WithIDCurrency id_currency获取 
func (obj *_EgWarehouseMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}

// WithIDAddress id_address获取 
func (obj *_EgWarehouseMgr) WithIDAddress(idAddress uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address"] = idAddress })
}

// WithIDEmployee id_employee获取 
func (obj *_EgWarehouseMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithReference reference获取 
func (obj *_EgWarehouseMgr) WithReference(reference string) Option {
	return optionFunc(func(o *options) { o.query["reference"] = reference })
}

// WithName name获取 
func (obj *_EgWarehouseMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithManagementType management_type获取 
func (obj *_EgWarehouseMgr) WithManagementType(managementType string) Option {
	return optionFunc(func(o *options) { o.query["management_type"] = managementType })
}

// WithDeleted deleted获取 
func (obj *_EgWarehouseMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}


// GetByOption 功能选项模式获取
func (obj *_EgWarehouseMgr) GetByOption(opts ...Option) (result EgWarehouse, err error) {
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
func (obj *_EgWarehouseMgr) GetByOptions(opts ...Option) (results []*EgWarehouse, err error) {
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


// GetFromIDWarehouse 通过id_warehouse获取内容  
func (obj *_EgWarehouseMgr)  GetFromIDWarehouse(idWarehouse uint32) (result EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&result).Error
	
	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找 
func (obj *_EgWarehouseMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error
	
	return
}
 
// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgWarehouseMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgWarehouseMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
// GetFromIDAddress 通过id_address获取内容  
func (obj *_EgWarehouseMgr) GetFromIDAddress(idAddress uint32) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address = ?", idAddress).Find(&results).Error
	
	return
}

// GetBatchFromIDAddress 批量唯一主键查找 
func (obj *_EgWarehouseMgr) GetBatchFromIDAddress(idAddresss []uint32) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address IN (?)", idAddresss).Find(&results).Error
	
	return
}
 
// GetFromIDEmployee 通过id_employee获取内容  
func (obj *_EgWarehouseMgr) GetFromIDEmployee(idEmployee uint32) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error
	
	return
}

// GetBatchFromIDEmployee 批量唯一主键查找 
func (obj *_EgWarehouseMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error
	
	return
}
 
// GetFromReference 通过reference获取内容  
func (obj *_EgWarehouseMgr) GetFromReference(reference string) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference = ?", reference).Find(&results).Error
	
	return
}

// GetBatchFromReference 批量唯一主键查找 
func (obj *_EgWarehouseMgr) GetBatchFromReference(references []string) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("reference IN (?)", references).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgWarehouseMgr) GetFromName(name string) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgWarehouseMgr) GetBatchFromName(names []string) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromManagementType 通过management_type获取内容  
func (obj *_EgWarehouseMgr) GetFromManagementType(managementType string) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("management_type = ?", managementType).Find(&results).Error
	
	return
}

// GetBatchFromManagementType 批量唯一主键查找 
func (obj *_EgWarehouseMgr) GetBatchFromManagementType(managementTypes []string) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("management_type IN (?)", managementTypes).Find(&results).Error
	
	return
}
 
// GetFromDeleted 通过deleted获取内容  
func (obj *_EgWarehouseMgr) GetFromDeleted(deleted bool) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error
	
	return
}

// GetBatchFromDeleted 批量唯一主键查找 
func (obj *_EgWarehouseMgr) GetBatchFromDeleted(deleteds []bool) (results []*EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgWarehouseMgr) FetchByPrimaryKey(idWarehouse uint32 ) (result EgWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&result).Error
	
	return
}
 

 

	

