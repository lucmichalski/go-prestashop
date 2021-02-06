package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImportMatchMgr struct {
	*_BaseMgr
}

func ImportMatchMgr(db *gorm.DB) *_ImportMatchMgr {
	if db == nil {
		panic(fmt.Errorf("ImportMatchMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImportMatchMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_import_match"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ImportMatchMgr) GetTableName() string {
	return "ps_import_match"
}

func (obj *_ImportMatchMgr) Get() (result ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ImportMatchMgr) Gets() (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ImportMatchMgr) WithIDImportMatch(idImportMatch int) Option {
	return optionFunc(func(o *options) { o.query["id_import_match"] = idImportMatch })
}

func (obj *_ImportMatchMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ImportMatchMgr) WithMatch(match string) Option {
	return optionFunc(func(o *options) { o.query["match"] = match })
}

func (obj *_ImportMatchMgr) WithSkip(skip int) Option {
	return optionFunc(func(o *options) { o.query["skip"] = skip })
}

func (obj *_ImportMatchMgr) GetByOption(opts ...Option) (result ImportMatch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ImportMatchMgr) GetByOptions(opts ...Option) (results []*ImportMatch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ImportMatchMgr) GetFromIDImportMatch(idImportMatch int) (result ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_import_match = ?", idImportMatch).Find(&result).Error

	return
}

func (obj *_ImportMatchMgr) GetBatchFromIDImportMatch(idImportMatchs []int) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_import_match IN (?)", idImportMatchs).Find(&results).Error

	return
}

func (obj *_ImportMatchMgr) GetFromName(name string) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ImportMatchMgr) GetBatchFromName(names []string) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ImportMatchMgr) GetFromMatch(match string) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("match = ?", match).Find(&results).Error

	return
}

func (obj *_ImportMatchMgr) GetBatchFromMatch(matchs []string) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("match IN (?)", matchs).Find(&results).Error

	return
}

func (obj *_ImportMatchMgr) GetFromSkip(skip int) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("skip = ?", skip).Find(&results).Error

	return
}

func (obj *_ImportMatchMgr) GetBatchFromSkip(skips []int) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("skip IN (?)", skips).Find(&results).Error

	return
}


func (obj *_ImportMatchMgr) FetchByPrimaryKey(idImportMatch int) (result ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_import_match = ?", idImportMatch).Find(&result).Error

	return
}
