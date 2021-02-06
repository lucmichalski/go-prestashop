package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _QuickAccessLangMgr struct {
	*_BaseMgr
}

// QuickAccessLangMgr open func
func QuickAccessLangMgr(db *gorm.DB) *_QuickAccessLangMgr {
	if db == nil {
		panic(fmt.Errorf("QuickAccessLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_QuickAccessLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_quick_access_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_QuickAccessLangMgr) GetTableName() string {
	return "eg_quick_access_lang"
}

// Get 获取
func (obj *_QuickAccessLangMgr) Get() (result QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_QuickAccessLangMgr) Gets() (results []*QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDQuickAccess id_quick_access获取
func (obj *_QuickAccessLangMgr) WithIDQuickAccess(idQuickAccess uint32) Option {
	return optionFunc(func(o *options) { o.query["id_quick_access"] = idQuickAccess })
}

// WithIDLang id_lang获取
func (obj *_QuickAccessLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_QuickAccessLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_QuickAccessLangMgr) GetByOption(opts ...Option) (result QuickAccessLang, err error) {
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
func (obj *_QuickAccessLangMgr) GetByOptions(opts ...Option) (results []*QuickAccessLang, err error) {
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
func (obj *_QuickAccessLangMgr) GetFromIDQuickAccess(idQuickAccess uint32) (results []*QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ?", idQuickAccess).Find(&results).Error

	return
}

// GetBatchFromIDQuickAccess 批量唯一主键查找
func (obj *_QuickAccessLangMgr) GetBatchFromIDQuickAccess(idQuickAccesss []uint32) (results []*QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access IN (?)", idQuickAccesss).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_QuickAccessLangMgr) GetFromIDLang(idLang uint32) (results []*QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_QuickAccessLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_QuickAccessLangMgr) GetFromName(name string) (results []*QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_QuickAccessLangMgr) GetBatchFromName(names []string) (results []*QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_QuickAccessLangMgr) FetchByPrimaryKey(idQuickAccess uint32, idLang uint32) (result QuickAccessLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_quick_access = ? AND id_lang = ?", idQuickAccess, idLang).Find(&result).Error

	return
}
