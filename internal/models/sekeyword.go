package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SekeywordMgr struct {
	*_BaseMgr
}

func SekeywordMgr(db *gorm.DB) *_SekeywordMgr {
	if db == nil {
		panic(fmt.Errorf("SekeywordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SekeywordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_sekeyword"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SekeywordMgr) GetTableName() string {
	return "ps_sekeyword"
}

func (obj *_SekeywordMgr) Get() (result Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SekeywordMgr) Gets() (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) WithIDSekeyword(idSekeyword uint32) Option {
	return optionFunc(func(o *options) { o.query["id_sekeyword"] = idSekeyword })
}

func (obj *_SekeywordMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_SekeywordMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

func (obj *_SekeywordMgr) WithKeyword(keyword string) Option {
	return optionFunc(func(o *options) { o.query["keyword"] = keyword })
}

func (obj *_SekeywordMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_SekeywordMgr) GetByOption(opts ...Option) (result Sekeyword, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SekeywordMgr) GetByOptions(opts ...Option) (results []*Sekeyword, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetFromIDSekeyword(idSekeyword uint32) (result Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword = ?", idSekeyword).Find(&result).Error

	return
}

func (obj *_SekeywordMgr) GetBatchFromIDSekeyword(idSekeywords []uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword IN (?)", idSekeywords).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetFromIDShop(idShop uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetBatchFromIDShop(idShops []uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetFromKeyword(keyword string) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keyword = ?", keyword).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetBatchFromKeyword(keywords []string) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keyword IN (?)", keywords).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetFromDateAdd(dateAdd time.Time) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_SekeywordMgr) FetchByPrimaryKey(idSekeyword uint32) (result Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword = ?", idSekeyword).Find(&result).Error

	return
}
