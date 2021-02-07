package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _QuickAccessMgr struct {
	*_BaseMgr
}

func QuickAccessMgr(db *gorm.DB) *_QuickAccessMgr {
	if db == nil {
		panic(fmt.Errorf("QuickAccessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_QuickAccessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_quick_access"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_QuickAccessMgr) GetTableName() string {
	return "ps_quick_access"
}

func (obj *_QuickAccessMgr) Get() (result QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_QuickAccessMgr) Gets() (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_QuickAccessMgr) WithIDQuickAccess(idQuickAccess uint32) Option {
	return optionFunc(func(o *options) { o.query["id_quick_access"] = idQuickAccess })
}

func (obj *_QuickAccessMgr) WithNewWindow(newWindow bool) Option {
	return optionFunc(func(o *options) { o.query["new_window"] = newWindow })
}

func (obj *_QuickAccessMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

func (obj *_QuickAccessMgr) GetByOption(opts ...Option) (result QuickAccess, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_QuickAccessMgr) GetByOptions(opts ...Option) (results []*QuickAccess, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_QuickAccessMgr) GetFromIDQuickAccess(idQuickAccess uint32) (result QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ?", idQuickAccess).Find(&result).Error

	return
}

func (obj *_QuickAccessMgr) GetBatchFromIDQuickAccess(idQuickAccesss []uint32) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access IN (?)", idQuickAccesss).Find(&results).Error

	return
}

func (obj *_QuickAccessMgr) GetFromNewWindow(newWindow bool) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window = ?", newWindow).Find(&results).Error

	return
}

func (obj *_QuickAccessMgr) GetBatchFromNewWindow(newWindows []bool) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window IN (?)", newWindows).Find(&results).Error

	return
}

func (obj *_QuickAccessMgr) GetFromLink(link string) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error

	return
}

func (obj *_QuickAccessMgr) GetBatchFromLink(links []string) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error

	return
}

func (obj *_QuickAccessMgr) FetchByPrimaryKey(idQuickAccess uint32) (result QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ?", idQuickAccess).Find(&result).Error

	return
}
