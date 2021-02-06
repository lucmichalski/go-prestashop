package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BadgeLangMgr struct {
	*_BaseMgr
}

// BadgeLangMgr open func
func BadgeLangMgr(db *gorm.DB) *_BadgeLangMgr {
	if db == nil {
		panic(fmt.Errorf("BadgeLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BadgeLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_badge_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BadgeLangMgr) GetTableName() string {
	return "ps_badge_lang"
}

// Get 获取
func (obj *_BadgeLangMgr) Get() (result BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BadgeLangMgr) Gets() (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDBadge id_badge获取
func (obj *_BadgeLangMgr) WithIDBadge(idBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_badge"] = idBadge })
}

// WithIDLang id_lang获取
func (obj *_BadgeLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_BadgeLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取
func (obj *_BadgeLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithGroupName group_name获取
func (obj *_BadgeLangMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}

// GetByOption 功能选项模式获取
func (obj *_BadgeLangMgr) GetByOption(opts ...Option) (result BadgeLang, err error) {
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
func (obj *_BadgeLangMgr) GetByOptions(opts ...Option) (results []*BadgeLang, err error) {
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

// GetFromIDBadge 通过id_badge获取内容
func (obj *_BadgeLangMgr) GetFromIDBadge(idBadge int) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&results).Error

	return
}

// GetBatchFromIDBadge 批量唯一主键查找
func (obj *_BadgeLangMgr) GetBatchFromIDBadge(idBadges []int) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge IN (?)", idBadges).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_BadgeLangMgr) GetFromIDLang(idLang int) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_BadgeLangMgr) GetBatchFromIDLang(idLangs []int) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_BadgeLangMgr) GetFromName(name string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_BadgeLangMgr) GetBatchFromName(names []string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容
func (obj *_BadgeLangMgr) GetFromDescription(description string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量唯一主键查找
func (obj *_BadgeLangMgr) GetBatchFromDescription(descriptions []string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromGroupName 通过group_name获取内容
func (obj *_BadgeLangMgr) GetFromGroupName(groupName string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_name = ?", groupName).Find(&results).Error

	return
}

// GetBatchFromGroupName 批量唯一主键查找
func (obj *_BadgeLangMgr) GetBatchFromGroupName(groupNames []string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_name IN (?)", groupNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BadgeLangMgr) FetchByPrimaryKey(idBadge int, idLang int) (result BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ? AND id_lang = ?", idBadge, idLang).Find(&result).Error

	return
}
