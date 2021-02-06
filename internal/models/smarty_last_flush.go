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

// SmartyLastFlushMgr open func
func SmartyLastFlushMgr(db *gorm.DB) *_SmartyLastFlushMgr {
	if db == nil {
		panic(fmt.Errorf("SmartyLastFlushMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SmartyLastFlushMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_smarty_last_flush"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SmartyLastFlushMgr) GetTableName() string {
	return "eg_smarty_last_flush"
}

// Get 获取
func (obj *_SmartyLastFlushMgr) Get() (result SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SmartyLastFlushMgr) Gets() (results []*SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithType type获取
func (obj *_SmartyLastFlushMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithLastFlush last_flush获取
func (obj *_SmartyLastFlushMgr) WithLastFlush(lastFlush time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_flush"] = lastFlush })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromType 通过type获取内容
func (obj *_SmartyLastFlushMgr) GetFromType(_type string) (result SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&result).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_SmartyLastFlushMgr) GetBatchFromType(_types []string) (results []*SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromLastFlush 通过last_flush获取内容
func (obj *_SmartyLastFlushMgr) GetFromLastFlush(lastFlush time.Time) (results []*SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_flush = ?", lastFlush).Find(&results).Error

	return
}

// GetBatchFromLastFlush 批量唯一主键查找
func (obj *_SmartyLastFlushMgr) GetBatchFromLastFlush(lastFlushs []time.Time) (results []*SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("last_flush IN (?)", lastFlushs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SmartyLastFlushMgr) FetchByPrimaryKey(_type string) (result SmartyLastFlush, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&result).Error

	return
}
