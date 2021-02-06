package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCategoryShopMgr struct {
	*_BaseMgr
}

// EgCategoryShopMgr open func
func EgCategoryShopMgr(db *gorm.DB) *_EgCategoryShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgCategoryShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCategoryShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_category_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCategoryShopMgr) GetTableName() string {
	return "eg_category_shop"
}

// Get 获取
func (obj *_EgCategoryShopMgr) Get() (result EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCategoryShopMgr) Gets() (results []*EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCategory id_category获取 
func (obj *_EgCategoryShopMgr) WithIDCategory(idCategory int) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithIDShop id_shop获取 
func (obj *_EgCategoryShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithPosition position获取 
func (obj *_EgCategoryShopMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}


// GetByOption 功能选项模式获取
func (obj *_EgCategoryShopMgr) GetByOption(opts ...Option) (result EgCategoryShop, err error) {
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
func (obj *_EgCategoryShopMgr) GetByOptions(opts ...Option) (results []*EgCategoryShop, err error) {
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


// GetFromIDCategory 通过id_category获取内容  
func (obj *_EgCategoryShopMgr) GetFromIDCategory(idCategory int) (results []*EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error
	
	return
}

// GetBatchFromIDCategory 批量唯一主键查找 
func (obj *_EgCategoryShopMgr) GetBatchFromIDCategory(idCategorys []int) (results []*EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCategoryShopMgr) GetFromIDShop(idShop int) (results []*EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCategoryShopMgr) GetBatchFromIDShop(idShops []int) (results []*EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromPosition 通过position获取内容  
func (obj *_EgCategoryShopMgr) GetFromPosition(position uint32) (results []*EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error
	
	return
}

// GetBatchFromPosition 批量唯一主键查找 
func (obj *_EgCategoryShopMgr) GetBatchFromPosition(positions []uint32) (results []*EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCategoryShopMgr) FetchByPrimaryKey(idCategory int ,idShop int ) (result EgCategoryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ? AND id_shop = ?", idCategory , idShop).Find(&result).Error
	
	return
}
 

 

	

