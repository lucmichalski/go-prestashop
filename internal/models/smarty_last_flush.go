package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SmartyLastFlushMgr struct {
	*_BaseMgr
}

func SmartyLastFlushMgr(db *gorm.DB) *_SmartyLastFlushMgr {
	if db == nil {
		panic(fmt.Errorf("SmartyLastFlushMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SmartyLastFlushMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_smarty_last_flush"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_SmartyLastFlushMgr) GetTableName() string {
	return "ps_smarty_last_flush"
}

func (obj *_SmartyLastFlushMgr) Get() (result SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_SmartyLastFlushMgr) Gets() (results []*SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_SmartyLastFlushMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

func (obj *_SmartyLastFlushMgr) WithLastFlush(lastFlush time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_flush"] = lastFlush })
}

func (obj *_SmartyLastFlushMgr) GetByOption(opts ...Option) (result SmartyLastFlush, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_SmartyLastFlushMgr) GetByOptions(opts ...Option) (results []*SmartyLastFlush, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_SmartyLastFlushMgr) GetFromType(_type string) (result SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&result).Error

	return
}

func (obj *_SmartyLastFlushMgr) GetBatchFromType(_types []string) (results []*SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

func (obj *_SmartyLastFlushMgr) GetFromLastFlush(lastFlush time.Time) (results []*SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_flush = ?", lastFlush).Find(&results).Error

	return
}

func (obj *_SmartyLastFlushMgr) GetBatchFromLastFlush(lastFlushs []time.Time) (results []*SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_flush IN (?)", lastFlushs).Find(&results).Error

	return
}


func (obj *_SmartyLastFlushMgr) FetchByPrimaryKey(_type string) (result SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&result).Error

	return
}
