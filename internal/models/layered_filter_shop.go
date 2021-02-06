package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredFilterShopMgr struct {
	*_BaseMgr
}

// LayeredFilterShopMgr open func
func LayeredFilterShopMgr(db *gorm.DB) *_LayeredFilterShopMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredFilterShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredFilterShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_filter_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredFilterShopMgr) GetTableName() string {
	return "ps_layered_filter_shop"
}

// Get 获取
func (obj *_LayeredFilterShopMgr) Get() (result LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredFilterShopMgr) Gets() (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLayeredFilter id_layered_filter获取
func (obj *_LayeredFilterShopMgr) WithIDLayeredFilter(idLayeredFilter uint32) Option {
	return optionFunc(func(o *options) { o.query["id_layered_filter"] = idLayeredFilter })
}

// WithIDShop id_shop获取
func (obj *_LayeredFilterShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_LayeredFilterShopMgr) GetByOption(opts ...Option) (result LayeredFilterShop, err error) {
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
func (obj *_LayeredFilterShopMgr) GetByOptions(opts ...Option) (results []*LayeredFilterShop, err error) {
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

// GetFromIDLayeredFilter 通过id_layered_filter获取内容
func (obj *_LayeredFilterShopMgr) GetFromIDLayeredFilter(idLayeredFilter uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ?", idLayeredFilter).Find(&results).Error

	return
}

// GetBatchFromIDLayeredFilter 批量唯一主键查找
func (obj *_LayeredFilterShopMgr) GetBatchFromIDLayeredFilter(idLayeredFilters []uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter IN (?)", idLayeredFilters).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_LayeredFilterShopMgr) GetFromIDShop(idShop uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_LayeredFilterShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredFilterShopMgr) FetchByPrimaryKey(idLayeredFilter uint32, idShop uint32) (result LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_layered_filter = ? AND id_shop = ?", idLayeredFilter, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_LayeredFilterShopMgr) FetchIndexByIDShop(idShop uint32) (results []*LayeredFilterShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
