package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ShopMgr struct {
	*_BaseMgr
}

func ShopMgr(db *gorm.DB) *_ShopMgr {
	if db == nil {
		panic(fmt.Errorf("ShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ShopMgr) GetTableName() string {
	return "ps_shop"
}

func (obj *_ShopMgr) Get() (result Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ShopMgr) Gets() (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_ShopMgr) WithIDShopGroup(idShopGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_ShopMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ShopMgr) WithIDCategory(idCategory int) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

func (obj *_ShopMgr) WithThemeName(themeName string) Option {
	return optionFunc(func(o *options) { o.query["theme_name"] = themeName })
}

func (obj *_ShopMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ShopMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

func (obj *_ShopMgr) GetByOption(opts ...Option) (result Shop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ShopMgr) GetByOptions(opts ...Option) (results []*Shop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ShopMgr) GetFromIDShop(idShop int) (result Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&result).Error

	return
}

func (obj *_ShopMgr) GetBatchFromIDShop(idShops []int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetFromIDShopGroup(idShopGroup int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetBatchFromIDShopGroup(idShopGroups []int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetFromName(name string) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetBatchFromName(names []string) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetFromIDCategory(idCategory int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetBatchFromIDCategory(idCategorys []int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetFromThemeName(themeName string) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("theme_name = ?", themeName).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetBatchFromThemeName(themeNames []string) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("theme_name IN (?)", themeNames).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetFromActive(active bool) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetBatchFromActive(actives []bool) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetFromDeleted(deleted bool) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

func (obj *_ShopMgr) GetBatchFromDeleted(deleteds []bool) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}


func (obj *_ShopMgr) FetchByPrimaryKey(idShop int) (result Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&result).Error

	return
}

func (obj *_ShopMgr) FetchIndexByIDX667E487AF5C9E40(idShopGroup int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}
