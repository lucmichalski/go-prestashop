package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HookAliasMgr struct {
	*_BaseMgr
}

// HookAliasMgr open func
func HookAliasMgr(db *gorm.DB) *_HookAliasMgr {
	if db == nil {
		panic(fmt.Errorf("HookAliasMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HookAliasMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_hook_alias"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HookAliasMgr) GetTableName() string {
	return "eg_hook_alias"
}

// Get 获取
func (obj *_HookAliasMgr) Get() (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HookAliasMgr) Gets() (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHookAlias id_hook_alias获取
func (obj *_HookAliasMgr) WithIDHookAlias(idHookAlias uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook_alias"] = idHookAlias })
}

// WithAlias alias获取
func (obj *_HookAliasMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithName name获取
func (obj *_HookAliasMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDHookAlias 通过id_hook_alias获取内容
func (obj *_HookAliasMgr) GetFromIDHookAlias(idHookAlias uint32) (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias = ?", idHookAlias).Find(&result).Error

	return
}

// GetBatchFromIDHookAlias 批量唯一主键查找
func (obj *_HookAliasMgr) GetBatchFromIDHookAlias(idHookAliass []uint32) (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias IN (?)", idHookAliass).Find(&results).Error

	return
}

// GetFromAlias 通过alias获取内容
func (obj *_HookAliasMgr) GetFromAlias(alias string) (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error

	return
}

// GetBatchFromAlias 批量唯一主键查找
func (obj *_HookAliasMgr) GetBatchFromAlias(aliass []string) (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_HookAliasMgr) GetFromName(name string) (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_HookAliasMgr) GetBatchFromName(names []string) (results []*HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_HookAliasMgr) FetchByPrimaryKey(idHookAlias uint32) (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook_alias = ?", idHookAlias).Find(&result).Error

	return
}

// FetchUniqueByAlias primay or index 获取唯一内容
func (obj *_HookAliasMgr) FetchUniqueByAlias(alias string) (result HookAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error

	return
}
