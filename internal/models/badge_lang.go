package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BadgeLangMgr struct {
	*_BaseMgr
}

func BadgeLangMgr(db *gorm.DB) *_BadgeLangMgr {
	if db == nil {
		panic(fmt.Errorf("BadgeLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BadgeLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_badge_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_BadgeLangMgr) GetTableName() string {
	return "ps_badge_lang"
}

func (obj *_BadgeLangMgr) Get() (result BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_BadgeLangMgr) Gets() (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) WithIDBadge(idBadge int) Option {
	return optionFunc(func(o *options) { o.query["id_badge"] = idBadge })
}

func (obj *_BadgeLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_BadgeLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_BadgeLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_BadgeLangMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}

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

func (obj *_BadgeLangMgr) GetFromIDBadge(idBadge int) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ?", idBadge).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetBatchFromIDBadge(idBadges []int) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge IN (?)", idBadges).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetFromIDLang(idLang int) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetBatchFromIDLang(idLangs []int) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetFromName(name string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetBatchFromName(names []string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetFromDescription(description string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetBatchFromDescription(descriptions []string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetFromGroupName(groupName string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_name = ?", groupName).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) GetBatchFromGroupName(groupNames []string) (results []*BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("group_name IN (?)", groupNames).Find(&results).Error

	return
}

func (obj *_BadgeLangMgr) FetchByPrimaryKey(idBadge int, idLang int) (result BadgeLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_badge = ? AND id_lang = ?", idBadge, idLang).Find(&result).Error

	return
}
