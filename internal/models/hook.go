package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HookMgr struct {
	*_BaseMgr
}

// HookMgr open func
func HookMgr(db *gorm.DB) *_HookMgr {
	if db == nil {
		panic(fmt.Errorf("HookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_hook"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HookMgr) GetTableName() string {
	return "eg_hook"
}

// Get 获取
func (obj *_HookMgr) Get() (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HookMgr) Gets() (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHook id_hook获取
func (obj *_HookMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

// WithName name获取
func (obj *_HookMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithTitle title获取
func (obj *_HookMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithDescription description获取
func (obj *_HookMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithPosition position获取
func (obj *_HookMgr) WithPosition(position bool) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDHook 通过id_hook获取内容
func (obj *_HookMgr) GetFromIDHook(idHook uint32) (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&result).Error

	return
}

// GetBatchFromIDHook 批量唯一主键查找
func (obj *_HookMgr) GetBatchFromIDHook(idHooks []uint32) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_HookMgr) GetFromName(name string) (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_HookMgr) GetBatchFromName(names []string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_HookMgr) GetFromTitle(title string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_HookMgr) GetBatchFromTitle(titles []string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_HookMgr) GetFromDescription(description string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_HookMgr) GetBatchFromDescription(descriptions []string) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_HookMgr) GetFromPosition(position bool) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_HookMgr) GetBatchFromPosition(positions []bool) (results []*Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_HookMgr) FetchByPrimaryKey(idHook uint32) (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&result).Error

	return
}

// FetchUniqueByHookName primay or index 获取唯一内容
func (obj *_HookMgr) FetchUniqueByHookName(name string) (result Hook, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}
