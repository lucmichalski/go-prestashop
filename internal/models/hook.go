package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HookMgr struct {
	*_BaseMgr
}

func HookMgr(db *gorm.DB) *_HookMgr {
	if db == nil {
		panic(fmt.Errorf("HookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_hook"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_HookMgr) GetTableName() string {
	return "ps_hook"
}

func (obj *_HookMgr) Get() (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_HookMgr) Gets() (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_HookMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

func (obj *_HookMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_HookMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

func (obj *_HookMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_HookMgr) WithPosition(position bool) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_HookMgr) GetByOption(opts ...Option) (result Hook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_HookMgr) GetByOptions(opts ...Option) (results []*Hook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_HookMgr) GetFromIDHook(idHook uint32) (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&result).Error

	return
}

func (obj *_HookMgr) GetBatchFromIDHook(idHooks []uint32) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error

	return
}

func (obj *_HookMgr) GetFromName(name string) (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}

func (obj *_HookMgr) GetBatchFromName(names []string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_HookMgr) GetFromTitle(title string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

func (obj *_HookMgr) GetBatchFromTitle(titles []string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

func (obj *_HookMgr) GetFromDescription(description string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_HookMgr) GetBatchFromDescription(descriptions []string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_HookMgr) GetFromPosition(position bool) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_HookMgr) GetBatchFromPosition(positions []bool) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}


func (obj *_HookMgr) FetchByPrimaryKey(idHook uint32) (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&result).Error

	return
}

func (obj *_HookMgr) FetchUniqueByHookName(name string) (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}
