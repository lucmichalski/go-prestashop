package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _InfoShopMgr struct {
	*_BaseMgr
}

// InfoShopMgr open func
func InfoShopMgr(db *gorm.DB) *_InfoShopMgr {
	if db == nil {
		panic(fmt.Errorf("InfoShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfoShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_info_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_InfoShopMgr) GetTableName() string {
	return "eg_info_shop"
}

// Get 获取
func (obj *_InfoShopMgr) Get() (result InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_InfoShopMgr) Gets() (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDInfo id_info获取
func (obj *_InfoShopMgr) WithIDInfo(idInfo uint32) Option {
	return optionFunc(func(o *options) { o.query["id_info"] = idInfo })
}

// WithIDShop id_shop获取
func (obj *_InfoShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_InfoShopMgr) GetByOption(opts ...Option) (result InfoShop, err error) {
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
func (obj *_InfoShopMgr) GetByOptions(opts ...Option) (results []*InfoShop, err error) {
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

// GetFromIDInfo 通过id_info获取内容
func (obj *_InfoShopMgr) GetFromIDInfo(idInfo uint32) (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&results).Error

	return
}

// GetBatchFromIDInfo 批量唯一主键查找
func (obj *_InfoShopMgr) GetBatchFromIDInfo(idInfos []uint32) (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info IN (?)", idInfos).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_InfoShopMgr) GetFromIDShop(idShop uint32) (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_InfoShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_InfoShopMgr) FetchByPrimaryKey(idInfo uint32, idShop uint32) (result InfoShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ? AND id_shop = ?", idInfo, idShop).Find(&result).Error

	return
}
