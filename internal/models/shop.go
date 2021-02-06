package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ShopMgr struct {
	*_BaseMgr
}

// ShopMgr open func
func ShopMgr(db *gorm.DB) *_ShopMgr {
	if db == nil {
		panic(fmt.Errorf("ShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ShopMgr) GetTableName() string {
	return "eg_shop"
}

// Get 获取
func (obj *_ShopMgr) Get() (result Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ShopMgr) Gets() (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDShop id_shop获取
func (obj *_ShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取
func (obj *_ShopMgr) WithIDShopGroup(idShopGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithName name获取
func (obj *_ShopMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIDCategory id_category获取
func (obj *_ShopMgr) WithIDCategory(idCategory int) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

// WithThemeName theme_name获取
func (obj *_ShopMgr) WithThemeName(themeName string) Option {
	return optionFunc(func(o *options) { o.query["theme_name"] = themeName })
}

// WithActive active获取
func (obj *_ShopMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取
func (obj *_ShopMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDShop 通过id_shop获取内容
func (obj *_ShopMgr) GetFromIDShop(idShop int) (result Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&result).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_ShopMgr) GetBatchFromIDShop(idShops []int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_ShopMgr) GetFromIDShopGroup(idShopGroup int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_ShopMgr) GetBatchFromIDShopGroup(idShopGroups []int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ShopMgr) GetFromName(name string) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ShopMgr) GetBatchFromName(names []string) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromIDCategory 通过id_category获取内容
func (obj *_ShopMgr) GetFromIDCategory(idCategory int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&results).Error

	return
}

// GetBatchFromIDCategory 批量唯一主键查找
func (obj *_ShopMgr) GetBatchFromIDCategory(idCategorys []int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

// GetFromThemeName 通过theme_name获取内容
func (obj *_ShopMgr) GetFromThemeName(themeName string) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("theme_name = ?", themeName).Find(&results).Error

	return
}

// GetBatchFromThemeName 批量唯一主键查找
func (obj *_ShopMgr) GetBatchFromThemeName(themeNames []string) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("theme_name IN (?)", themeNames).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_ShopMgr) GetFromActive(active bool) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_ShopMgr) GetBatchFromActive(actives []bool) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_ShopMgr) GetFromDeleted(deleted bool) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_ShopMgr) GetBatchFromDeleted(deleteds []bool) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ShopMgr) FetchByPrimaryKey(idShop int) (result Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&result).Error

	return
}

// FetchIndexByIDX667E487AF5C9E40  获取多个内容
func (obj *_ShopMgr) FetchIndexByIDX667E487AF5C9E40(idShopGroup int) (results []*Shop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}
