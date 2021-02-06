package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _QuickAccessMgr struct {
	*_BaseMgr
}

// QuickAccessMgr open func
func QuickAccessMgr(db *gorm.DB) *_QuickAccessMgr {
	if db == nil {
		panic(fmt.Errorf("QuickAccessMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_QuickAccessMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_quick_access"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_QuickAccessMgr) GetTableName() string {
	return "eg_quick_access"
}

// Get 获取
func (obj *_QuickAccessMgr) Get() (result QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_QuickAccessMgr) Gets() (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDQuickAccess id_quick_access获取
func (obj *_QuickAccessMgr) WithIDQuickAccess(idQuickAccess uint32) Option {
	return optionFunc(func(o *options) { o.query["id_quick_access"] = idQuickAccess })
}

// WithNewWindow new_window获取
func (obj *_QuickAccessMgr) WithNewWindow(newWindow bool) Option {
	return optionFunc(func(o *options) { o.query["new_window"] = newWindow })
}

// WithLink link获取
func (obj *_QuickAccessMgr) WithLink(link string) Option {
	return optionFunc(func(o *options) { o.query["link"] = link })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDQuickAccess 通过id_quick_access获取内容
func (obj *_QuickAccessMgr) GetFromIDQuickAccess(idQuickAccess uint32) (result QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ?", idQuickAccess).Find(&result).Error

	return
}

// GetBatchFromIDQuickAccess 批量唯一主键查找
func (obj *_QuickAccessMgr) GetBatchFromIDQuickAccess(idQuickAccesss []uint32) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access IN (?)", idQuickAccesss).Find(&results).Error

	return
}

// GetFromNewWindow 通过new_window获取内容
func (obj *_QuickAccessMgr) GetFromNewWindow(newWindow bool) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window = ?", newWindow).Find(&results).Error

	return
}

// GetBatchFromNewWindow 批量唯一主键查找
func (obj *_QuickAccessMgr) GetBatchFromNewWindow(newWindows []bool) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window IN (?)", newWindows).Find(&results).Error

	return
}

// GetFromLink 通过link获取内容
func (obj *_QuickAccessMgr) GetFromLink(link string) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link = ?", link).Find(&results).Error

	return
}

// GetBatchFromLink 批量唯一主键查找
func (obj *_QuickAccessMgr) GetBatchFromLink(links []string) (results []*QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("link IN (?)", links).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_QuickAccessMgr) FetchByPrimaryKey(idQuickAccess uint32) (result QuickAccess, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ?", idQuickAccess).Find(&result).Error

	return
}
