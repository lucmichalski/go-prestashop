package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HookAliasMgr struct {
	*_BaseMgr
}

func HookAliasMgr(db *gorm.DB) *_HookAliasMgr {
	if db == nil {
		panic(fmt.Errorf("HookAliasMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HookAliasMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_hook_alias"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_HookAliasMgr) GetTableName() string {
	return "ps_hook_alias"
}

func (obj *_HookAliasMgr) Get() (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_HookAliasMgr) Gets() (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_HookAliasMgr) WithIDHookAlias(idHookAlias uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook_alias"] = idHookAlias })
}

func (obj *_HookAliasMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

func (obj *_HookAliasMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_HookAliasMgr) GetByOption(opts ...Option) (result HookAlias, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_HookAliasMgr) GetByOptions(opts ...Option) (results []*HookAlias, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_HookAliasMgr) GetFromIDHookAlias(idHookAlias uint32) (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias = ?", idHookAlias).Find(&result).Error

	return
}

func (obj *_HookAliasMgr) GetBatchFromIDHookAlias(idHookAliass []uint32) (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias IN (?)", idHookAliass).Find(&results).Error

	return
}

func (obj *_HookAliasMgr) GetFromAlias(alias string) (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error

	return
}

func (obj *_HookAliasMgr) GetBatchFromAlias(aliass []string) (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error

	return
}

func (obj *_HookAliasMgr) GetFromName(name string) (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_HookAliasMgr) GetBatchFromName(names []string) (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_HookAliasMgr) FetchByPrimaryKey(idHookAlias uint32) (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias = ?", idHookAlias).Find(&result).Error

	return
}

func (obj *_HookAliasMgr) FetchUniqueByAlias(alias string) (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error

	return
}
