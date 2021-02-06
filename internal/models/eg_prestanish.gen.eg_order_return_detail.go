package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgOrderReturnDetailMgr struct {
	*_BaseMgr
}

// EgOrderReturnDetailMgr open func
func EgOrderReturnDetailMgr(db *gorm.DB) *_EgOrderReturnDetailMgr {
	if db == nil {
		panic(fmt.Errorf("EgOrderReturnDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgOrderReturnDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_return_detail"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgOrderReturnDetailMgr) GetTableName() string {
	return "eg_order_return_detail"
}

// Get 获取
func (obj *_EgOrderReturnDetailMgr) Get() (result EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgOrderReturnDetailMgr) Gets() (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderReturn id_order_return获取 
func (obj *_EgOrderReturnDetailMgr) WithIDOrderReturn(idOrderReturn uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return"] = idOrderReturn })
}

// WithIDOrderDetail id_order_detail获取 
func (obj *_EgOrderReturnDetailMgr) WithIDOrderDetail(idOrderDetail uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_detail"] = idOrderDetail })
}

// WithIDCustomization id_customization获取 
func (obj *_EgOrderReturnDetailMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

// WithProductQuantity product_quantity获取 
func (obj *_EgOrderReturnDetailMgr) WithProductQuantity(productQuantity uint32) Option {
	return optionFunc(func(o *options) { o.query["product_quantity"] = productQuantity })
}


// GetByOption 功能选项模式获取
func (obj *_EgOrderReturnDetailMgr) GetByOption(opts ...Option) (result EgOrderReturnDetail, err error) {
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
func (obj *_EgOrderReturnDetailMgr) GetByOptions(opts ...Option) (results []*EgOrderReturnDetail, err error) {
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


// GetFromIDOrderReturn 通过id_order_return获取内容  
func (obj *_EgOrderReturnDetailMgr) GetFromIDOrderReturn(idOrderReturn uint32) (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return = ?", idOrderReturn).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderReturn 批量唯一主键查找 
func (obj *_EgOrderReturnDetailMgr) GetBatchFromIDOrderReturn(idOrderReturns []uint32) (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return IN (?)", idOrderReturns).Find(&results).Error
	
	return
}
 
// GetFromIDOrderDetail 通过id_order_detail获取内容  
func (obj *_EgOrderReturnDetailMgr) GetFromIDOrderDetail(idOrderDetail uint32) (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail = ?", idOrderDetail).Find(&results).Error
	
	return
}

// GetBatchFromIDOrderDetail 批量唯一主键查找 
func (obj *_EgOrderReturnDetailMgr) GetBatchFromIDOrderDetail(idOrderDetails []uint32) (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_detail IN (?)", idOrderDetails).Find(&results).Error
	
	return
}
 
// GetFromIDCustomization 通过id_customization获取内容  
func (obj *_EgOrderReturnDetailMgr) GetFromIDCustomization(idCustomization uint32) (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomization 批量唯一主键查找 
func (obj *_EgOrderReturnDetailMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error
	
	return
}
 
// GetFromProductQuantity 通过product_quantity获取内容  
func (obj *_EgOrderReturnDetailMgr) GetFromProductQuantity(productQuantity uint32) (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity = ?", productQuantity).Find(&results).Error
	
	return
}

// GetBatchFromProductQuantity 批量唯一主键查找 
func (obj *_EgOrderReturnDetailMgr) GetBatchFromProductQuantity(productQuantitys []uint32) (results []*EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_quantity IN (?)", productQuantitys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgOrderReturnDetailMgr) FetchByPrimaryKey(idOrderReturn uint32 ,idOrderDetail uint32 ,idCustomization uint32 ) (result EgOrderReturnDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return = ? AND id_order_detail = ? AND id_customization = ?", idOrderReturn , idOrderDetail , idCustomization).Find(&result).Error
	
	return
}
 

 

	

