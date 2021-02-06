package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SearchEngineMgr struct {
	*_BaseMgr
}

func SearchEngineMgr(db *gorm.DB) *_SearchEngineMgr {
	if db == nil {
		panic(fmt.Errorf("SearchEngineMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SearchEngineMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_search_engine"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SearchEngineMgr) GetTableName() string {
	return "ps_search_engine"
}

func (obj *_SearchEngineMgr) Get() (result SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SearchEngineMgr) Gets() (results []*SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SearchEngineMgr) WithIDSearchEngine(idSearchEngine uint32) Option {
	return optionFunc(func(o *options) { o.query["id_search_engine"] = idSearchEngine })
}

func (obj *_SearchEngineMgr) WithServer(server string) Option {
	return optionFunc(func(o *options) { o.query["server"] = server })
}

func (obj *_SearchEngineMgr) WithGetvar(getvar string) Option {
	return optionFunc(func(o *options) { o.query["getvar"] = getvar })
}

func (obj *_SearchEngineMgr) GetByOption(opts ...Option) (result SearchEngine, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SearchEngineMgr) GetByOptions(opts ...Option) (results []*SearchEngine, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SearchEngineMgr) GetFromIDSearchEngine(idSearchEngine uint32) (result SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_search_engine = ?", idSearchEngine).Find(&result).Error

	return
}

func (obj *_SearchEngineMgr) GetBatchFromIDSearchEngine(idSearchEngines []uint32) (results []*SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_search_engine IN (?)", idSearchEngines).Find(&results).Error

	return
}

func (obj *_SearchEngineMgr) GetFromServer(server string) (results []*SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("server = ?", server).Find(&results).Error

	return
}

func (obj *_SearchEngineMgr) GetBatchFromServer(servers []string) (results []*SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("server IN (?)", servers).Find(&results).Error

	return
}

func (obj *_SearchEngineMgr) GetFromGetvar(getvar string) (results []*SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("getvar = ?", getvar).Find(&results).Error

	return
}

func (obj *_SearchEngineMgr) GetBatchFromGetvar(getvars []string) (results []*SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("getvar IN (?)", getvars).Find(&results).Error

	return
}


func (obj *_SearchEngineMgr) FetchByPrimaryKey(idSearchEngine uint32) (result SearchEngine, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_search_engine = ?", idSearchEngine).Find(&result).Error

	return
}
