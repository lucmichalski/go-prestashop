package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DeliveryMgr struct {
	*_BaseMgr
}

// DeliveryMgr open func
func DeliveryMgr(db *gorm.DB) *_DeliveryMgr {
	if db == nil {
		panic(fmt.Errorf("DeliveryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DeliveryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_delivery"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DeliveryMgr) GetTableName() string {
	return "eg_delivery"
}

// Get 获取
func (obj *_DeliveryMgr) Get() (result Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DeliveryMgr) Gets() (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDDelivery id_delivery获取
func (obj *_DeliveryMgr) WithIDDelivery(idDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_delivery"] = idDelivery })
}

// WithIDShop id_shop获取
func (obj *_DeliveryMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取
func (obj *_DeliveryMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithIDCarrier id_carrier获取
func (obj *_DeliveryMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDRangePrice id_range_price获取
func (obj *_DeliveryMgr) WithIDRangePrice(idRangePrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_price"] = idRangePrice })
}

// WithIDRangeWeight id_range_weight获取
func (obj *_DeliveryMgr) WithIDRangeWeight(idRangeWeight uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_weight"] = idRangeWeight })
}

// WithIDZone id_zone获取
func (obj *_DeliveryMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

// WithPrice price获取
func (obj *_DeliveryMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// GetByOption 功能选项模式获取
func (obj *_DeliveryMgr) GetByOption(opts ...Option) (result Delivery, err error) {
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
func (obj *_DeliveryMgr) GetByOptions(opts ...Option) (results []*Delivery, err error) {
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
func (obj *_DeliveryMgr) GetFromIDDelivery(idDelivery uint32) (result Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery = ?", idDelivery).Find(&result).Error

	return
}

// GetBatchFromIDDelivery 批量唯一主键查找
func (obj *_DeliveryMgr) GetBatchFromIDDelivery(idDeliverys []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery IN (?)", idDeliverys).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_DeliveryMgr) GetFromIDShop(idShop uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_DeliveryMgr) GetBatchFromIDShop(idShops []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_DeliveryMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_DeliveryMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromIDCarrier 通过id_carrier获取内容
func (obj *_DeliveryMgr) GetFromIDCarrier(idCarrier uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_DeliveryMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromIDRangePrice 通过id_range_price获取内容
func (obj *_DeliveryMgr) GetFromIDRangePrice(idRangePrice uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&results).Error

	return
}

// GetBatchFromIDRangePrice 批量唯一主键查找
func (obj *_DeliveryMgr) GetBatchFromIDRangePrice(idRangePrices []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price IN (?)", idRangePrices).Find(&results).Error

	return
}

// GetFromIDRangeWeight 通过id_range_weight获取内容
func (obj *_DeliveryMgr) GetFromIDRangeWeight(idRangeWeight uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&results).Error

	return
}

// GetBatchFromIDRangeWeight 批量唯一主键查找
func (obj *_DeliveryMgr) GetBatchFromIDRangeWeight(idRangeWeights []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight IN (?)", idRangeWeights).Find(&results).Error

	return
}

// GetFromIDZone 通过id_zone获取内容
func (obj *_DeliveryMgr) GetFromIDZone(idZone uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}

// GetBatchFromIDZone 批量唯一主键查找
func (obj *_DeliveryMgr) GetBatchFromIDZone(idZones []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容
func (obj *_DeliveryMgr) GetFromPrice(price float64) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量唯一主键查找
func (obj *_DeliveryMgr) GetBatchFromPrice(prices []float64) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_DeliveryMgr) FetchByPrimaryKey(idDelivery uint32) (result Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery = ?", idDelivery).Find(&result).Error

	return
}

// FetchIndexByIDCarrier  获取多个内容
func (obj *_DeliveryMgr) FetchIndexByIDCarrier(idCarrier uint32, idZone uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_zone = ?", idCarrier, idZone).Find(&results).Error

	return
}

// FetchIndexByIDRangePrice  获取多个内容
func (obj *_DeliveryMgr) FetchIndexByIDRangePrice(idRangePrice uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&results).Error

	return
}

// FetchIndexByIDRangeWeight  获取多个内容
func (obj *_DeliveryMgr) FetchIndexByIDRangeWeight(idRangeWeight uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&results).Error

	return
}

// FetchIndexByIDZone  获取多个内容
func (obj *_DeliveryMgr) FetchIndexByIDZone(idZone uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}
