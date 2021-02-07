package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AliasMgr struct {
	*_BaseMgr
}

func AliasMgr(db *gorm.DB) *_AliasMgr {
	if db == nil {
		panic(fmt.Errorf("AliasMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AliasMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_alias"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AliasMgr) GetTableName() string {
	return "ps_alias"
}

func (obj *_AliasMgr) Get() (result Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AliasMgr) Gets() (results []*Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_AliasMgr) WithIDAlias(idAlias uint32) Option {
	return optionFunc(func(o *options) { o.query["id_alias"] = idAlias })
}

func (obj *_AliasMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

func (obj *_AliasMgr) WithSearch(search string) Option {
	return optionFunc(func(o *options) { o.query["search"] = search })
}

func (obj *_AliasMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_AliasMgr) GetByOption(opts ...Option) (result Alias, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AliasMgr) GetByOptions(opts ...Option) (results []*Alias, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_AliasMgr) GetFromIDAlias(idAlias uint32) (result Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_alias = ?", idAlias).Find(&result).Error

	return
}

func (obj *_AliasMgr) GetBatchFromIDAlias(idAliass []uint32) (results []*Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_alias IN (?)", idAliass).Find(&results).Error

	return
}

func (obj *_AliasMgr) GetFromAlias(alias string) (result Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error

	return
}

func (obj *_AliasMgr) GetBatchFromAlias(aliass []string) (results []*Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error

	return
}

func (obj *_AliasMgr) GetFromSearch(search string) (results []*Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("search = ?", search).Find(&results).Error

	return
}

func (obj *_AliasMgr) GetBatchFromSearch(searchs []string) (results []*Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("search IN (?)", searchs).Find(&results).Error

	return
}

func (obj *_AliasMgr) GetFromActive(active bool) (results []*Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_AliasMgr) GetBatchFromActive(actives []bool) (results []*Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_AliasMgr) FetchByPrimaryKey(idAlias uint32) (result Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_alias = ?", idAlias).Find(&result).Error

	return
}

func (obj *_AliasMgr) FetchUniqueByAlias(alias string) (result Alias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error

	return
}
