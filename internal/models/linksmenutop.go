package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LinksmenutopMgr struct {
	*_BaseMgr
}

func LinksmenutopMgr(db *gorm.DB) *_LinksmenutopMgr {
	if db == nil {
		panic(fmt.Errorf("LinksmenutopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LinksmenutopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_linksmenutop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LinksmenutopMgr) GetTableName() string {
	return "ps_linksmenutop"
}

func (obj *_LinksmenutopMgr) Get() (result Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LinksmenutopMgr) Gets() (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LinksmenutopMgr) WithIDLinksmenutop(idLinksmenutop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_linksmenutop"] = idLinksmenutop })
}

func (obj *_LinksmenutopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_LinksmenutopMgr) WithNewWindow(newWindow bool) Option {
	return optionFunc(func(o *options) { o.query["new_window"] = newWindow })
}

func (obj *_LinksmenutopMgr) GetByOption(opts ...Option) (result Linksmenutop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LinksmenutopMgr) GetByOptions(opts ...Option) (results []*Linksmenutop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LinksmenutopMgr) GetFromIDLinksmenutop(idLinksmenutop uint32) (result Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ?", idLinksmenutop).Find(&result).Error

	return
}

func (obj *_LinksmenutopMgr) GetBatchFromIDLinksmenutop(idLinksmenutops []uint32) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop IN (?)", idLinksmenutops).Find(&results).Error

	return
}

func (obj *_LinksmenutopMgr) GetFromIDShop(idShop uint32) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_LinksmenutopMgr) GetBatchFromIDShop(idShops []uint32) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_LinksmenutopMgr) GetFromNewWindow(newWindow bool) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window = ?", newWindow).Find(&results).Error

	return
}

func (obj *_LinksmenutopMgr) GetBatchFromNewWindow(newWindows []bool) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window IN (?)", newWindows).Find(&results).Error

	return
}


func (obj *_LinksmenutopMgr) FetchByPrimaryKey(idLinksmenutop uint32) (result Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ?", idLinksmenutop).Find(&result).Error

	return
}

func (obj *_LinksmenutopMgr) FetchIndexByIDShop(idShop uint32) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
