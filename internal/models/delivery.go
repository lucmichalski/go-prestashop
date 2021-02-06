package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _DeliveryMgr struct {
	*_BaseMgr
}

func DeliveryMgr(db *gorm.DB) *_DeliveryMgr {
	if db == nil {
		panic(fmt.Errorf("DeliveryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DeliveryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_delivery"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_DeliveryMgr) GetTableName() string {
	return "ps_delivery"
}

func (obj *_DeliveryMgr) Get() (result Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_DeliveryMgr) Gets() (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_DeliveryMgr) WithIDDelivery(idDelivery uint32) Option {
	return optionFunc(func(o *options) { o.query["id_delivery"] = idDelivery })
}

func (obj *_DeliveryMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_DeliveryMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_DeliveryMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

func (obj *_DeliveryMgr) WithIDRangePrice(idRangePrice uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_price"] = idRangePrice })
}

func (obj *_DeliveryMgr) WithIDRangeWeight(idRangeWeight uint32) Option {
	return optionFunc(func(o *options) { o.query["id_range_weight"] = idRangeWeight })
}

func (obj *_DeliveryMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

func (obj *_DeliveryMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

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


func (obj *_DeliveryMgr) GetFromIDDelivery(idDelivery uint32) (result Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery = ?", idDelivery).Find(&result).Error

	return
}

func (obj *_DeliveryMgr) GetBatchFromIDDelivery(idDeliverys []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery IN (?)", idDeliverys).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetFromIDShop(idShop uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetBatchFromIDShop(idShops []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetFromIDCarrier(idCarrier uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetFromIDRangePrice(idRangePrice uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetBatchFromIDRangePrice(idRangePrices []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price IN (?)", idRangePrices).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetFromIDRangeWeight(idRangeWeight uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetBatchFromIDRangeWeight(idRangeWeights []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight IN (?)", idRangeWeights).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetFromIDZone(idZone uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetBatchFromIDZone(idZones []uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetFromPrice(price float64) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) GetBatchFromPrice(prices []float64) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}


func (obj *_DeliveryMgr) FetchByPrimaryKey(idDelivery uint32) (result Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_delivery = ?", idDelivery).Find(&result).Error

	return
}

func (obj *_DeliveryMgr) FetchIndexByIDCarrier(idCarrier uint32, idZone uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_zone = ?", idCarrier, idZone).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) FetchIndexByIDRangePrice(idRangePrice uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_price = ?", idRangePrice).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) FetchIndexByIDRangeWeight(idRangeWeight uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_range_weight = ?", idRangeWeight).Find(&results).Error

	return
}

func (obj *_DeliveryMgr) FetchIndexByIDZone(idZone uint32) (results []*Delivery, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}
