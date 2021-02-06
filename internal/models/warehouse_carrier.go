package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgWarehouseCarrierMgr struct {
	*_BaseMgr
}

// EgWarehouseCarrierMgr open func
func EgWarehouseCarrierMgr(db *gorm.DB) *_EgWarehouseCarrierMgr {
	if db == nil {
		panic(fmt.Errorf("EgWarehouseCarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgWarehouseCarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_warehouse_carrier"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgWarehouseCarrierMgr) GetTableName() string {
	return "eg_warehouse_carrier"
}

// Get 获取
func (obj *_EgWarehouseCarrierMgr) Get() (result EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgWarehouseCarrierMgr) Gets() (results []*EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取 
func (obj *_EgWarehouseCarrierMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDWarehouse id_warehouse获取 
func (obj *_EgWarehouseCarrierMgr) WithIDWarehouse(idWarehouse uint32) Option {
	return optionFunc(func(o *options) { o.query["id_warehouse"] = idWarehouse })
}


// GetByOption 功能选项模式获取
func (obj *_EgWarehouseCarrierMgr) GetByOption(opts ...Option) (result EgWarehouseCarrier, err error) {
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
func (obj *_EgWarehouseCarrierMgr) GetByOptions(opts ...Option) (results []*EgWarehouseCarrier, err error) {
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


// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgWarehouseCarrierMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgWarehouseCarrierMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromIDWarehouse 通过id_warehouse获取内容  
func (obj *_EgWarehouseCarrierMgr) GetFromIDWarehouse(idWarehouse uint32) (results []*EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}

// GetBatchFromIDWarehouse 批量唯一主键查找 
func (obj *_EgWarehouseCarrierMgr) GetBatchFromIDWarehouse(idWarehouses []uint32) (results []*EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse IN (?)", idWarehouses).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgWarehouseCarrierMgr) FetchByPrimaryKey(idCarrier uint32 ,idWarehouse uint32 ) (result EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_warehouse = ?", idCarrier , idWarehouse).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCarrier  获取多个内容
 func (obj *_EgWarehouseCarrierMgr) FetchIndexByIDCarrier(idCarrier uint32 ) (results []*EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDWarehouse  获取多个内容
 func (obj *_EgWarehouseCarrierMgr) FetchIndexByIDWarehouse(idWarehouse uint32 ) (results []*EgWarehouseCarrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_warehouse = ?", idWarehouse).Find(&results).Error
	
	return
}
 

	

