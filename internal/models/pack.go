package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgPackMgr struct {
	*_BaseMgr
}

// EgPackMgr open func
func EgPackMgr(db *gorm.DB) *_EgPackMgr {
	if db == nil {
		panic(fmt.Errorf("EgPackMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgPackMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_pack"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgPackMgr) GetTableName() string {
	return "eg_pack"
}

// Get 获取
func (obj *_EgPackMgr) Get() (result EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgPackMgr) Gets() (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductPack id_product_pack获取 
func (obj *_EgPackMgr) WithIDProductPack(idProductPack uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_pack"] = idProductPack })
}

// WithIDProductItem id_product_item获取 
func (obj *_EgPackMgr) WithIDProductItem(idProductItem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_item"] = idProductItem })
}

// WithIDProductAttributeItem id_product_attribute_item获取 
func (obj *_EgPackMgr) WithIDProductAttributeItem(idProductAttributeItem uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_attribute_item"] = idProductAttributeItem })
}

// WithQuantity quantity获取 
func (obj *_EgPackMgr) WithQuantity(quantity uint32) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}


// GetByOption 功能选项模式获取
func (obj *_EgPackMgr) GetByOption(opts ...Option) (result EgPack, err error) {
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
func (obj *_EgPackMgr) GetByOptions(opts ...Option) (results []*EgPack, err error) {
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


// GetFromIDProductPack 通过id_product_pack获取内容  
func (obj *_EgPackMgr) GetFromIDProductPack(idProductPack uint32) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack = ?", idProductPack).Find(&results).Error
	
	return
}

// GetBatchFromIDProductPack 批量唯一主键查找 
func (obj *_EgPackMgr) GetBatchFromIDProductPack(idProductPacks []uint32) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack IN (?)", idProductPacks).Find(&results).Error
	
	return
}
 
// GetFromIDProductItem 通过id_product_item获取内容  
func (obj *_EgPackMgr) GetFromIDProductItem(idProductItem uint32) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item = ?", idProductItem).Find(&results).Error
	
	return
}

// GetBatchFromIDProductItem 批量唯一主键查找 
func (obj *_EgPackMgr) GetBatchFromIDProductItem(idProductItems []uint32) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item IN (?)", idProductItems).Find(&results).Error
	
	return
}
 
// GetFromIDProductAttributeItem 通过id_product_attribute_item获取内容  
func (obj *_EgPackMgr) GetFromIDProductAttributeItem(idProductAttributeItem uint32) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute_item = ?", idProductAttributeItem).Find(&results).Error
	
	return
}

// GetBatchFromIDProductAttributeItem 批量唯一主键查找 
func (obj *_EgPackMgr) GetBatchFromIDProductAttributeItem(idProductAttributeItems []uint32) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_attribute_item IN (?)", idProductAttributeItems).Find(&results).Error
	
	return
}
 
// GetFromQuantity 通过quantity获取内容  
func (obj *_EgPackMgr) GetFromQuantity(quantity uint32) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity = ?", quantity).Find(&results).Error
	
	return
}

// GetBatchFromQuantity 批量唯一主键查找 
func (obj *_EgPackMgr) GetBatchFromQuantity(quantitys []uint32) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("quantity IN (?)", quantitys).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgPackMgr) FetchByPrimaryKey(idProductPack uint32 ,idProductItem uint32 ,idProductAttributeItem uint32 ) (result EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_pack = ? AND id_product_item = ? AND id_product_attribute_item = ?", idProductPack , idProductItem , idProductAttributeItem).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByProductItem  获取多个内容
 func (obj *_EgPackMgr) FetchIndexByProductItem(idProductItem uint32 ,idProductAttributeItem uint32 ) (results []*EgPack, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_item = ? AND id_product_attribute_item = ?", idProductItem , idProductAttributeItem).Find(&results).Error
	
	return
}
 

	

