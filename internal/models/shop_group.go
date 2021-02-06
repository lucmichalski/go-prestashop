package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ShopGroupMgr struct {
	*_BaseMgr
}

func ShopGroupMgr(db *gorm.DB) *_ShopGroupMgr {
	if db == nil {
		panic(fmt.Errorf("ShopGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ShopGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_shop_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ShopGroupMgr) GetTableName() string {
	return "ps_shop_group"
}

func (obj *_ShopGroupMgr) Get() (result ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ShopGroupMgr) Gets() (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ShopGroupMgr) WithIDShopGroup(idShopGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_ShopGroupMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ShopGroupMgr) WithShareCustomer(shareCustomer bool) Option {
	return optionFunc(func(o *options) { o.query["share_customer"] = shareCustomer })
}

func (obj *_ShopGroupMgr) WithShareOrder(shareOrder bool) Option {
	return optionFunc(func(o *options) { o.query["share_order"] = shareOrder })
}

func (obj *_ShopGroupMgr) WithShareStock(shareStock bool) Option {
	return optionFunc(func(o *options) { o.query["share_stock"] = shareStock })
}

func (obj *_ShopGroupMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ShopGroupMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_ShopGroupMgr) GetByOption(opts ...Option) (result ShopGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ShopGroupMgr) GetByOptions(opts ...Option) (results []*ShopGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ShopGroupMgr) GetFromIDShopGroup(idShopGroup int) (result ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&result).Error

	return
}

func (obj *_ShopGroupMgr) GetBatchFromIDShopGroup(idShopGroups []int) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetFromName(name string) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetBatchFromName(names []string) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetFromShareCustomer(shareCustomer bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_customer = ?", shareCustomer).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetBatchFromShareCustomer(shareCustomers []bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_customer IN (?)", shareCustomers).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetFromShareOrder(shareOrder bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_order = ?", shareOrder).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetBatchFromShareOrder(shareOrders []bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_order IN (?)", shareOrders).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetFromShareStock(shareStock bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_stock = ?", shareStock).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetBatchFromShareStock(shareStocks []bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("share_stock IN (?)", shareStocks).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetFromActive(active bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetBatchFromActive(actives []bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetFromDeleted(deleted bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_ShopGroupMgr) GetBatchFromDeleted(deleteds []bool) (results []*ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}


func (obj *_ShopGroupMgr) FetchByPrimaryKey(idShopGroup int) (result ShopGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&result).Error

	return
}
