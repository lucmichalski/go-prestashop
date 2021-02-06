package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgLayeredProductAttributeMgr struct {
	*_BaseMgr
}

// EgLayeredProductAttributeMgr open func
func EgLayeredProductAttributeMgr(db *gorm.DB) *_EgLayeredProductAttributeMgr {
	if db == nil {
		panic(fmt.Errorf("EgLayeredProductAttributeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLayeredProductAttributeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_product_attribute"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLayeredProductAttributeMgr) GetTableName() string {
	return "eg_layered_product_attribute"
}

// Get 获取
func (obj *_EgLayeredProductAttributeMgr) Get() (result EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLayeredProductAttributeMgr) Gets() (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttribute id_attribute获取 
func (obj *_EgLayeredProductAttributeMgr) WithIDAttribute(idAttribute uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

// WithIDProduct id_product获取 
func (obj *_EgLayeredProductAttributeMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDAttributeGroup id_attribute_group获取 
func (obj *_EgLayeredProductAttributeMgr) WithIDAttributeGroup(idAttributeGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIDShop id_shop获取 
func (obj *_EgLayeredProductAttributeMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgLayeredProductAttributeMgr) GetByOption(opts ...Option) (result EgLayeredProductAttribute, err error) {
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
func (obj *_EgLayeredProductAttributeMgr) GetByOptions(opts ...Option) (results []*EgLayeredProductAttribute, err error) {
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


// GetFromIDAttribute 通过id_attribute获取内容  
func (obj *_EgLayeredProductAttributeMgr) GetFromIDAttribute(idAttribute uint32) (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error
	
	return
}

// GetBatchFromIDAttribute 批量唯一主键查找 
func (obj *_EgLayeredProductAttributeMgr) GetBatchFromIDAttribute(idAttributes []uint32) (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error
	
	return
}
 
// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgLayeredProductAttributeMgr) GetFromIDProduct(idProduct uint32) (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgLayeredProductAttributeMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDAttributeGroup 通过id_attribute_group获取内容  
func (obj *_EgLayeredProductAttributeMgr) GetFromIDAttributeGroup(idAttributeGroup uint32) (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找 
func (obj *_EgLayeredProductAttributeMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []uint32) (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgLayeredProductAttributeMgr) GetFromIDShop(idShop uint32) (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgLayeredProductAttributeMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLayeredProductAttributeMgr) FetchByPrimaryKey(idAttribute uint32 ,idProduct uint32 ,idShop uint32 ) (result EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_product = ? AND id_shop = ?", idAttribute , idProduct , idShop).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDAttributeGroup primay or index 获取唯一内容
 func (obj *_EgLayeredProductAttributeMgr) FetchUniqueIndexByIDAttributeGroup(idAttribute uint32 ,idProduct uint32 ,idAttributeGroup uint32 ,idShop uint32 ) (result EgLayeredProductAttribute, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_product = ? AND id_attribute_group = ? AND id_shop = ?", idAttribute , idProduct , idAttributeGroup , idShop).Find(&result).Error
	
	return
}
 

 

	

