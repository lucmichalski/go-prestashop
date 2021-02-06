package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgDeliveryMgr struct {
	*_BaseMgr
}

// EgDeliveryMgr open func
func EgDeliveryMgr(db *gorm.DB) *_EgDeliveryMgr {
	if db == nil {
		panic(fmt.Errorf("EgDeliveryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgDeliveryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_delivery"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgDeliveryMgr) GetTableName() string {
	return "eg_delivery"
}

// Get 获取
func (obj *_EgDeliveryMgr) Get() (result EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgDeliveryMgr) Gets() (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDDelivery id_delivery获取 
func (obj *_EgDeliveryMgr) WithIDDelivery(idDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_delivery"] = idDelivery })
}

// WithIDShop id_shop获取 
func (obj *_EgDeliveryMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取 
func (obj *_EgDeliveryMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDCarrier id_carrier获取 
func (obj *_EgDeliveryMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDRangePrice id_range_price获取 
func (obj *_EgDeliveryMgr) WithIDRangePrice(idRangePrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_price"] = idRangePrice })
}

// WithIDRangeWeight id_range_weight获取 
func (obj *_EgDeliveryMgr) WithIDRangeWeight(idRangeWeight uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_weight"] = idRangeWeight })
}

// WithIDZone id_zone获取 
func (obj *_EgDeliveryMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

// WithPrice price获取 
func (obj *_EgDeliveryMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}


// GetByOption 功能选项模式获取
func (obj *_EgDeliveryMgr) GetByOption(opts ...Option) (result EgDelivery, err error) {
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
func (obj *_EgDeliveryMgr) GetByOptions(opts ...Option) (results []*EgDelivery, err error) {
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


// GetFromIDDelivery 通过id_delivery获取内容  
func (obj *_EgDeliveryMgr)  GetFromIDDelivery(idDelivery uint32) (result EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery = ?", idDelivery).Find(&result).Error
	
	return
}

// GetBatchFromIDDelivery 批量唯一主键查找 
func (obj *_EgDeliveryMgr) GetBatchFromIDDelivery(idDeliverys []uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery IN (?)", idDeliverys).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgDeliveryMgr) GetFromIDShop(idShop uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgDeliveryMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDShopGroup 通过id_shop_group获取内容  
func (obj *_EgDeliveryMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error
	
	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找 
func (obj *_EgDeliveryMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error
	
	return
}
 
// GetFromIDCarrier 通过id_carrier获取内容  
func (obj *_EgDeliveryMgr) GetFromIDCarrier(idCarrier uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error
	
	return
}

// GetBatchFromIDCarrier 批量唯一主键查找 
func (obj *_EgDeliveryMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error
	
	return
}
 
// GetFromIDRangePrice 通过id_range_price获取内容  
func (obj *_EgDeliveryMgr) GetFromIDRangePrice(idRangePrice uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&results).Error
	
	return
}

// GetBatchFromIDRangePrice 批量唯一主键查找 
func (obj *_EgDeliveryMgr) GetBatchFromIDRangePrice(idRangePrices []uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price IN (?)", idRangePrices).Find(&results).Error
	
	return
}
 
// GetFromIDRangeWeight 通过id_range_weight获取内容  
func (obj *_EgDeliveryMgr) GetFromIDRangeWeight(idRangeWeight uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&results).Error
	
	return
}

// GetBatchFromIDRangeWeight 批量唯一主键查找 
func (obj *_EgDeliveryMgr) GetBatchFromIDRangeWeight(idRangeWeights []uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight IN (?)", idRangeWeights).Find(&results).Error
	
	return
}
 
// GetFromIDZone 通过id_zone获取内容  
func (obj *_EgDeliveryMgr) GetFromIDZone(idZone uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error
	
	return
}

// GetBatchFromIDZone 批量唯一主键查找 
func (obj *_EgDeliveryMgr) GetBatchFromIDZone(idZones []uint32) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error
	
	return
}
 
// GetFromPrice 通过price获取内容  
func (obj *_EgDeliveryMgr) GetFromPrice(price float64) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error
	
	return
}

// GetBatchFromPrice 批量唯一主键查找 
func (obj *_EgDeliveryMgr) GetBatchFromPrice(prices []float64) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgDeliveryMgr) FetchByPrimaryKey(idDelivery uint32 ) (result EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery = ?", idDelivery).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDCarrier  获取多个内容
 func (obj *_EgDeliveryMgr) FetchIndexByIDCarrier(idCarrier uint32 ,idZone uint32 ) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_zone = ?", idCarrier , idZone).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDRangePrice  获取多个内容
 func (obj *_EgDeliveryMgr) FetchIndexByIDRangePrice(idRangePrice uint32 ) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDRangeWeight  获取多个内容
 func (obj *_EgDeliveryMgr) FetchIndexByIDRangeWeight(idRangeWeight uint32 ) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&results).Error
	
	return
}
 
 // FetchIndexByIDZone  获取多个内容
 func (obj *_EgDeliveryMgr) FetchIndexByIDZone(idZone uint32 ) (results []*EgDelivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error
	
	return
}
 

	

