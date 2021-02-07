package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TagCountMgr struct {
	*_BaseMgr
}

func TagCountMgr(db *gorm.DB) *_TagCountMgr {
	if db == nil {
		panic(fmt.Errorf("TagCountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TagCountMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tag_count"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_TagCountMgr) GetTableName() string {
	return "ps_tag_count"
}

func (obj *_TagCountMgr) Get() (result TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_TagCountMgr) Gets() (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_TagCountMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

func (obj *_TagCountMgr) WithIDTag(idTag uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tag"] = idTag })
}

func (obj *_TagCountMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_TagCountMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_TagCountMgr) WithCounter(counter uint32) Option {
	return optionFunc(func(o *options) { o.query["counter"] = counter })
}

func (obj *_TagCountMgr) GetByOption(opts ...Option) (result TagCount, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_TagCountMgr) GetByOptions(opts ...Option) (results []*TagCount, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetFromIDGroup(idGroup uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetFromIDTag(idTag uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetBatchFromIDTag(idTags []uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag IN (?)", idTags).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetFromIDLang(idLang uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetBatchFromIDLang(idLangs []uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetFromIDShop(idShop uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetBatchFromIDShop(idShops []uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetFromCounter(counter uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter = ?", counter).Find(&results).Error

	return
}

func (obj *_TagCountMgr) GetBatchFromCounter(counters []uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("counter IN (?)", counters).Find(&results).Error

	return
}

func (obj *_TagCountMgr) FetchByPrimaryKey(idGroup uint32, idTag uint32) (result TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_tag = ?", idGroup, idTag).Find(&result).Error

	return
}

func (obj *_TagCountMgr) FetchIndexByIDGroup(idGroup uint32, idLang uint32, idShop uint32, counter uint32) (results []*TagCount, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ? AND id_lang = ? AND id_shop = ? AND counter = ?", idGroup, idLang, idShop, counter).Find(&results).Error

	return
}
