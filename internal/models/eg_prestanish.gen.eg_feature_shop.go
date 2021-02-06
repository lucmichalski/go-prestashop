package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgFeatureShopMgr struct {
	*_BaseMgr
}

// EgFeatureShopMgr open func
func EgFeatureShopMgr(db *gorm.DB) *_EgFeatureShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgFeatureShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgFeatureShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_feature_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgFeatureShopMgr) GetTableName() string {
	return "eg_feature_shop"
}

// Get 获取
func (obj *_EgFeatureShopMgr) Get() (result EgFeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgFeatureShopMgr) Gets() (results []*EgFeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeature id_feature获取 
func (obj *_EgFeatureShopMgr) WithIDFeature(idFeature uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

// WithIDShop id_shop获取 
func (obj *_EgFeatureShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgFeatureShopMgr) GetByOption(opts ...Option) (result EgFeatureShop, err error) {
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
func (obj *_EgFeatureShopMgr) GetByOptions(opts ...Option) (results []*EgFeatureShop, err error) {
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


// GetFromIDFeature 通过id_feature获取内容  
func (obj *_EgFeatureShopMgr) GetFromIDFeature(idFeature uint32) (results []*EgFeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error
	
	return
}

// GetBatchFromIDFeature 批量唯一主键查找 
func (obj *_EgFeatureShopMgr) GetBatchFromIDFeature(idFeatures []uint32) (results []*EgFeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgFeatureShopMgr) GetFromIDShop(idShop uint32) (results []*EgFeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgFeatureShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgFeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgFeatureShopMgr) FetchByPrimaryKey(idFeature uint32 ,idShop uint32 ) (result EgFeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ? AND id_shop = ?", idFeature , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgFeatureShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgFeatureShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

