package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgProductSupplierMgr struct {
	*_BaseMgr
}

// EgProductSupplierMgr open func
func EgProductSupplierMgr(db *gorm.DB) *_EgProductSupplierMgr {
	if db == nil {
		panic(fmt.Errorf("EgProductSupplierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgProductSupplierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_product_supplier"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgProductSupplierMgr) GetTableName() string {
	return "eg_product_supplier"
}

// Get 获取
func (obj *_EgProductSupplierMgr) Get() (result EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgProductSupplierMgr) Gets() (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductSupplier id_product_supplier获取 
func (obj *_EgProductSupplierMgr) WithIDProductSupplier(idProductSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_supplier"] = idProductSupplier })
}

// WithIDProduct id_product获取 
func (obj *_EgProductSupplierMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDProductAttribute id_product_attribute获取 
func (obj *_EgProductSupplierMgr) WithIDProductAttribute(idProductAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute"] = idProductAttribute })
}

// WithIDSupplier id_supplier获取 
func (obj *_EgProductSupplierMgr) WithIDSupplier(idSupplier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_supplier"] = idSupplier })
}

// WithProductSupplierReference product_supplier_reference获取 
func (obj *_EgProductSupplierMgr) WithProductSupplierReference(productSupplierReference string) Option {
	return optionFunc(func(o *options) { o.query["product_supplier_reference"] = productSupplierReference })
}

// WithProductSupplierPriceTe product_supplier_price_te获取 
func (obj *_EgProductSupplierMgr) WithProductSupplierPriceTe(productSupplierPriceTe float64) Option {
	return optionFunc(func(o *options) { o.query["product_supplier_price_te"] = productSupplierPriceTe })
}

// WithIDCurrency id_currency获取 
func (obj *_EgProductSupplierMgr) WithIDCurrency(idCurrency uint32) Option {
	return optionFunc(func(o *options) { o.query["id_currency"] = idCurrency })
}


// GetByOption 功能选项模式获取
func (obj *_EgProductSupplierMgr) GetByOption(opts ...Option) (result EgProductSupplier, err error) {
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
func (obj *_EgProductSupplierMgr) GetByOptions(opts ...Option) (results []*EgProductSupplier, err error) {
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


// GetFromIDProductSupplier 通过id_product_supplier获取内容  
func (obj *_EgProductSupplierMgr)  GetFromIDProductSupplier(idProductSupplier uint32) (result EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_supplier = ?", idProductSupplier).Find(&result).Error
	
	return
}

// GetBatchFromIDProductSupplier 批量唯一主键查找 
func (obj *_EgProductSupplierMgr) GetBatchFromIDProductSupplier(idProductSuppliers []uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_supplier IN (?)", idProductSuppliers).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgProductSupplierMgr) GetFromIDProduct(idProduct uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgProductSupplierMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDProductAttribute 通过id_product_attribute获取内容  
func (obj *_EgProductSupplierMgr) GetFromIDProductAttribute(idProductAttribute uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute = ?", idProductAttribute).Find(&results).Error
	
	return
}

// GetBatchFromIDProductAttribute 批量唯一主键查找 
func (obj *_EgProductSupplierMgr) GetBatchFromIDProductAttribute(idProductAttributes []uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute IN (?)", idProductAttributes).Find(&results).Error
	
	return
}
 
// GetFromIDSupplier 通过id_supplier获取内容  
func (obj *_EgProductSupplierMgr) GetFromIDSupplier(idSupplier uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier = ?", idSupplier).Find(&results).Error
	
	return
}

// GetBatchFromIDSupplier 批量唯一主键查找 
func (obj *_EgProductSupplierMgr) GetBatchFromIDSupplier(idSuppliers []uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_supplier IN (?)", idSuppliers).Find(&results).Error
	
	return
}
 
// GetFromProductSupplierReference 通过product_supplier_reference获取内容  
func (obj *_EgProductSupplierMgr) GetFromProductSupplierReference(productSupplierReference string) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_reference = ?", productSupplierReference).Find(&results).Error
	
	return
}

// GetBatchFromProductSupplierReference 批量唯一主键查找 
func (obj *_EgProductSupplierMgr) GetBatchFromProductSupplierReference(productSupplierReferences []string) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_reference IN (?)", productSupplierReferences).Find(&results).Error
	
	return
}
 
// GetFromProductSupplierPriceTe 通过product_supplier_price_te获取内容  
func (obj *_EgProductSupplierMgr) GetFromProductSupplierPriceTe(productSupplierPriceTe float64) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_price_te = ?", productSupplierPriceTe).Find(&results).Error
	
	return
}

// GetBatchFromProductSupplierPriceTe 批量唯一主键查找 
func (obj *_EgProductSupplierMgr) GetBatchFromProductSupplierPriceTe(productSupplierPriceTes []float64) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("product_supplier_price_te IN (?)", productSupplierPriceTes).Find(&results).Error
	
	return
}
 
// GetFromIDCurrency 通过id_currency获取内容  
func (obj *_EgProductSupplierMgr) GetFromIDCurrency(idCurrency uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency = ?", idCurrency).Find(&results).Error
	
	return
}

// GetBatchFromIDCurrency 批量唯一主键查找 
func (obj *_EgProductSupplierMgr) GetBatchFromIDCurrency(idCurrencys []uint32) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_currency IN (?)", idCurrencys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgProductSupplierMgr) FetchByPrimaryKey(idProductSupplier uint32 ) (result EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_supplier = ?", idProductSupplier).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDProduct primay or index 获取唯一内容
 func (obj *_EgProductSupplierMgr) FetchUniqueIndexByIDProduct(idProduct uint32 ,idProductAttribute uint32 ,idSupplier uint32 ) (result EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_product_attribute = ? AND id_supplier = ?", idProduct , idProductAttribute , idSupplier).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDSupplier  获取多个内容
 func (obj *_EgProductSupplierMgr) FetchIndexByIDSupplier(idProduct uint32 ,idSupplier uint32 ) (results []*EgProductSupplier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_supplier = ?", idProduct , idSupplier).Find(&results).Error
	
	return
}
 

	

