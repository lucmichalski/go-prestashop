package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LinkBlockShopMgr struct {
	*_BaseMgr
}

func LinkBlockShopMgr(db *gorm.DB) *_LinkBlockShopMgr {
	if db == nil {
		panic(fmt.Errorf("LinkBlockShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LinkBlockShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_link_block_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LinkBlockShopMgr) GetTableName() string {
	return "ps_link_block_shop"
}

func (obj *_LinkBlockShopMgr) Get() (result LinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LinkBlockShopMgr) Gets() (results []*LinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_LinkBlockShopMgr) WithIDLinkBlock(idLinkBlock uint32) Option {
	return optionFunc(func(o *options) { o.query["id_link_block"] = idLinkBlock })
}

func (obj *_LinkBlockShopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_LinkBlockShopMgr) GetByOption(opts ...Option) (result LinkBlockShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LinkBlockShopMgr) GetByOptions(opts ...Option) (results []*LinkBlockShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_LinkBlockShopMgr) GetFromIDLinkBlock(idLinkBlock uint32) (results []*LinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ?", idLinkBlock).Find(&results).Error

	return
}

func (obj *_LinkBlockShopMgr) GetBatchFromIDLinkBlock(idLinkBlocks []uint32) (results []*LinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block IN (?)", idLinkBlocks).Find(&results).Error

	return
}

func (obj *_LinkBlockShopMgr) GetFromIDShop(idShop uint32) (results []*LinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_LinkBlockShopMgr) GetBatchFromIDShop(idShops []uint32) (results []*LinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_LinkBlockShopMgr) FetchByPrimaryKey(idLinkBlock uint32, idShop uint32) (result LinkBlockShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ? AND id_shop = ?", idLinkBlock, idShop).Find(&result).Error

	return
}
