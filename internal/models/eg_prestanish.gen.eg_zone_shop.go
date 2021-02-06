package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgZoneShopMgr struct {
	*_BaseMgr
}

// EgZoneShopMgr open func
func EgZoneShopMgr(db *gorm.DB) *_EgZoneShopMgr {
	if db == nil {
		panic(fmt.Errorf("EgZoneShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgZoneShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_zone_shop"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgZoneShopMgr) GetTableName() string {
	return "eg_zone_shop"
}

// Get 获取
func (obj *_EgZoneShopMgr) Get() (result EgZoneShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgZoneShopMgr) Gets() (results []*EgZoneShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDZone id_zone获取 
func (obj *_EgZoneShopMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

// WithIDShop id_shop获取 
func (obj *_EgZoneShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}


// GetByOption 功能选项模式获取
func (obj *_EgZoneShopMgr) GetByOption(opts ...Option) (result EgZoneShop, err error) {
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
func (obj *_EgZoneShopMgr) GetByOptions(opts ...Option) (results []*EgZoneShop, err error) {
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


// GetFromIDZone 通过id_zone获取内容  
func (obj *_EgZoneShopMgr) GetFromIDZone(idZone uint32) (results []*EgZoneShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error
	
	return
}

// GetBatchFromIDZone 批量唯一主键查找 
func (obj *_EgZoneShopMgr) GetBatchFromIDZone(idZones []uint32) (results []*EgZoneShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgZoneShopMgr) GetFromIDShop(idShop uint32) (results []*EgZoneShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgZoneShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgZoneShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgZoneShopMgr) FetchByPrimaryKey(idZone uint32 ,idShop uint32 ) (result EgZoneShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ? AND id_shop = ?", idZone , idShop).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDShop  获取多个内容
 func (obj *_EgZoneShopMgr) FetchIndexByIDShop(idShop uint32 ) (results []*EgZoneShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}
 

	

