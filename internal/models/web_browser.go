package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _WebBrowserMgr struct {
	*_BaseMgr
}

func WebBrowserMgr(db *gorm.DB) *_WebBrowserMgr {
	if db == nil {
		panic(fmt.Errorf("WebBrowserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WebBrowserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_web_browser"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_WebBrowserMgr) GetTableName() string {
	return "ps_web_browser"
}

func (obj *_WebBrowserMgr) Get() (result WebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_WebBrowserMgr) Gets() (results []*WebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_WebBrowserMgr) WithIDWebBrowser(idWebBrowser uint32) Option {
	return optionFunc(func(o *options) { o.query["id_web_browser"] = idWebBrowser })
}

func (obj *_WebBrowserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_WebBrowserMgr) GetByOption(opts ...Option) (result WebBrowser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_WebBrowserMgr) GetByOptions(opts ...Option) (results []*WebBrowser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_WebBrowserMgr) GetFromIDWebBrowser(idWebBrowser uint32) (result WebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser = ?", idWebBrowser).Find(&result).Error

	return
}

func (obj *_WebBrowserMgr) GetBatchFromIDWebBrowser(idWebBrowsers []uint32) (results []*WebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser IN (?)", idWebBrowsers).Find(&results).Error

	return
}

func (obj *_WebBrowserMgr) GetFromName(name string) (results []*WebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_WebBrowserMgr) GetBatchFromName(names []string) (results []*WebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_WebBrowserMgr) FetchByPrimaryKey(idWebBrowser uint32) (result WebBrowser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_web_browser = ?", idWebBrowser).Find(&result).Error

	return
}
