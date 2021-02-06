package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgAttributeShopMgr struct {
	*_BaseMgr
}

// EgAttributeShopMgr open func
func EgAttributeShopMgr(db *gorm.DB) *_EgAttributeShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgAttributeShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAttributeShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attribute_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAttributeShopMgr) GetTableName() string {
	return "eg_attribute_shop"
}

// Get 获取
func (obj *_EgAttributeShopMgr) Get() (result EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAttributeShopMgr) Gets() (results []*EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttribute id_attribute获取 
func (obj *_EgAttributeShopMgr) WithIDAttribute(idAttribute int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute"] = idAttribute })
}

// WithIDShop id_shop获取 
func (obj *_EgAttributeShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgAttributeShopMgr) GetByOption(opts ...Option) (result EgAttributeShop, err error) {
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
func (obj *_EgAttributeShopMgr) GetByOptions(opts ...Option) (results []*EgAttributeShop, err error) {
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
func (obj *_EgAttributeShopMgr) GetFromIDAttribute(idAttribute int) (results []*EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error
	
	return
}

// GetBatchFromIDAttribute 批量唯一主键查找 
func (obj *_EgAttributeShopMgr) GetBatchFromIDAttribute(idAttributes []int) (results []*EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute IN (?)", idAttributes).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgAttributeShopMgr) GetFromIDShop(idShop int) (results []*EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgAttributeShopMgr) GetBatchFromIDShop(idShops []int) (results []*EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAttributeShopMgr) FetchByPrimaryKey(idAttribute int ,idShop int ) (result EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ? AND id_shop = ?", idAttribute , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDX7634898F7A4F53DC  获取多个内容
 func (obj *_EgAttributeShopMgr) FetchIndexByIDX7634898F7A4F53DC(idAttribute int ) (results []*EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute = ?", idAttribute).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDX7634898F274A50A0  获取多个内容
 func (obj *_EgAttributeShopMgr) FetchIndexByIDX7634898F274A50A0(idShop int ) (results []*EgAttributeShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

