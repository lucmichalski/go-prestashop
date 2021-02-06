package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCustomizationMgr struct {
	*_BaseMgr
}

// EgCustomizationMgr open func
func EgCustomizationMgr(db *gorm.DB) *_EgCustomizationMgr {
	if db == nil {
		panic(fmt.Errorf("EgCustomizationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCustomizationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customization"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCustomizationMgr) GetTableName() string {
	return "eg_customization"
}

// Get 获取
func (obj *_EgCustomizationMgr) Get() (result EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCustomizationMgr) Gets() (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCustomization id_customization获取 
func (obj *_EgCustomizationMgr) WithIDCustomization(idCustomization uint32) Option {
	return optionFunc(func(o *options) { o.query["id_customization"] = idCustomization })
}

// WithIDProductAttribute id_product_attribute获取 
func (obj *_EgCustomizationMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDAddressDelivery id_address_delivery获取 
func (obj *_EgCustomizationMgr) WithIDAddressDelivery(idAddressDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_address_delivery"] = idAddressDelivery })
}

// WithIDCart id_cart获取 
func (obj *_EgCustomizationMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDProduct id_product获取 
func (obj *_EgCustomizationMgr) WithIDProduct(idProduct int) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithQuantity quantity获取 
func (obj *_EgCustomizationMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithQuantityRefunded quantity_refunded获取 
func (obj *_EgCustomizationMgr) WithQuantityRefunded(quantityRefunded int) Option {
	return optionFunc(func(o *options) { o.query["quantity_refunded"] = quantityRefunded })
}

// WithQuantityReturned quantity_returned获取 
func (obj *_EgCustomizationMgr) WithQuantityReturned(quantityReturned int) Option {
	return optionFunc(func(o *options) { o.query["quantity_returned"] = quantityReturned })
}

// WithInCart in_cart获取 
func (obj *_EgCustomizationMgr) WithInCart(inCart bool) Option {
	return optionFunc(func(o *options) { o.query["in_cart"] = inCart })
}


// GetByOption 功能选项模式获取
func (obj *_EgCustomizationMgr) GetByOption(opts ...Option) (result EgCustomization, err error) {
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
func (obj *_EgCustomizationMgr) GetByOptions(opts ...Option) (results []*EgCustomization, err error) {
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


// GetFromIDCustomization 通过id_customization获取内容  
func (obj *_EgCustomizationMgr) GetFromIDCustomization(idCustomization uint32) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ?", idCustomization).Find(&results).Error
	
	return
}

// GetBatchFromIDCustomization 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromIDCustomization(idCustomizations []uint32) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization IN (?)", idCustomizations).Find(&results).Error
	
	return
}
 
// GetFromIDProductAttribute 通过id_product_attribute获取内容  
func (obj *_EgCustomizationMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error
	
	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error
	
	return
}
 
// GetFromIDAddressDelivery 通过id_address_delivery获取内容  
func (obj *_EgCustomizationMgr) GetFromIDAddressDelivery(idAddressDelivery uint32) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery = ?", idAddressDelivery).Find(&results).Error
	
	return
}

// GetBatchFromIDAddressDelivery 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromIDAddressDelivery(idAddressDeliverys []uint32) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_address_delivery IN (?)", idAddressDeliverys).Find(&results).Error
	
	return
}
 
// GetFromIDCart 通过id_cart获取内容  
func (obj *_EgCustomizationMgr) GetFromIDCart(idCart uint32) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error
	
	return
}

// GetBatchFromIDCart 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromIDCart(idCarts []uint32) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgCustomizationMgr) GetFromIDProduct(idProduct int) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromIDProduct(idProducts []int) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromQuantity 通过quantity获取内容  
func (obj *_EgCustomizationMgr) GetFromQuantity(quantity int) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error
	
	return
}

// GetBatchFromQuantity 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromQuantity(quantitys []int) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error
	
	return
}
 
// GetFromQuantityRefunded 通过quantity_refunded获取内容  
func (obj *_EgCustomizationMgr) GetFromQuantityRefunded(quantityRefunded int) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_refunded = ?", quantityRefunded).Find(&results).Error
	
	return
}

// GetBatchFromQuantityRefunded 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromQuantityRefunded(quantityRefundeds []int) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_refunded IN (?)", quantityRefundeds).Find(&results).Error
	
	return
}
 
// GetFromQuantityReturned 通过quantity_returned获取内容  
func (obj *_EgCustomizationMgr) GetFromQuantityReturned(quantityReturned int) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_returned = ?", quantityReturned).Find(&results).Error
	
	return
}

// GetBatchFromQuantityReturned 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromQuantityReturned(quantityReturneds []int) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity_returned IN (?)", quantityReturneds).Find(&results).Error
	
	return
}
 
// GetFromInCart 通过in_cart获取内容  
func (obj *_EgCustomizationMgr) GetFromInCart(inCart bool) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("in_cart = ?", inCart).Find(&results).Error
	
	return
}

// GetBatchFromInCart 批量唯一主键查找 
func (obj *_EgCustomizationMgr) GetBatchFromInCart(inCarts []bool) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("in_cart IN (?)", inCarts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCustomizationMgr) FetchByPrimaryKey(idCustomization uint32 ,idAddressDelivery uint32 ,idCart uint32 ,idProduct int ) (result EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_customization = ? AND id_address_delivery = ? AND id_cart = ? AND id_product = ?", idCustomization , idAddressDelivery , idCart , idProduct).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDProductAttribute  获取多个内容
 func (obj *_EgCustomizationMgr) FetchIndexByIDProductAttribute(idProductAttribute uint32 ) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDCartProduct  获取多个内容
 func (obj *_EgCustomizationMgr) FetchIndexByIDCartProduct(idProductAttribute uint32 ,idCart uint32 ,idProduct int ) (results []*EgCustomization, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ? AND id_cart = ? AND id_product = ?", idProductAttribute , idCart , idProduct).Find(&results).Error
	
	return
}
 

	

