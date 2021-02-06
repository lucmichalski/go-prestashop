package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgCountryShopMgr struct {
	*_BaseMgr
}

// EgCountryShopMgr open func
func EgCountryShopMgr(db *gorm.DB) *_EgCountryShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgCountryShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCountryShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_country_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCountryShopMgr) GetTableName() string {
	return "eg_country_shop"
}

// Get 获取
func (obj *_EgCountryShopMgr) Get() (result EgCountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCountryShopMgr) Gets() (results []*EgCountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCountry id_country获取 
func (obj *_EgCountryShopMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithIDShop id_shop获取 
func (obj *_EgCountryShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgCountryShopMgr) GetByOption(opts ...Option) (result EgCountryShop, err error) {
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
func (obj *_EgCountryShopMgr) GetByOptions(opts ...Option) (results []*EgCountryShop, err error) {
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


// GetFromIDCountry 通过id_country获取内容  
func (obj *_EgCountryShopMgr) GetFromIDCountry(idCountry uint32) (results []*EgCountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error
	
	return
}

// GetBatchFromIDCountry 批量唯一主键查找 
func (obj *_EgCountryShopMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*EgCountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCountryShopMgr) GetFromIDShop(idShop uint32) (results []*EgCountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCountryShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCountryShopMgr) FetchByPrimaryKey(idCountry uint32 ,idShop uint32 ) (result EgCountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ? AND id_shop = ?", idCountry , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgCountryShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgCountryShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

