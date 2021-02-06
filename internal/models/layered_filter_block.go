package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LayeredFilterBlockMgr struct {
	*_BaseMgr
}

// LayeredFilterBlockMgr open func
func LayeredFilterBlockMgr(db *gorm.DB) *_LayeredFilterBlockMgr {
	if db == nil {
		panic(fmt.Errorf("LayeredFilterBlockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LayeredFilterBlockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_layered_filter_block"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LayeredFilterBlockMgr) GetTableName() string {
	return "ps_layered_filter_block"
}

// Get 获取
func (obj *_LayeredFilterBlockMgr) Get() (result LayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LayeredFilterBlockMgr) Gets() (results []*LayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithHash hash获取
func (obj *_LayeredFilterBlockMgr) WithHash(hash string) Option {
	return optionFunc(func(o *options) { o.query["hash"] = hash })
}

// WithData data获取
func (obj *_LayeredFilterBlockMgr) WithData(data string) Option {
	return optionFunc(func(o *options) { o.query["data"] = data })
}

// GetByOption 功能选项模式获取
func (obj *_LayeredFilterBlockMgr) GetByOption(opts ...Option) (result LayeredFilterBlock, err error) {
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
func (obj *_LayeredFilterBlockMgr) GetByOptions(opts ...Option) (results []*LayeredFilterBlock, err error) {
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

// GetFromHash 通过hash获取内容
func (obj *_LayeredFilterBlockMgr) GetFromHash(hash string) (result LayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hash = ?", hash).Find(&result).Error

	return
}

// GetBatchFromHash 批量唯一主键查找
func (obj *_LayeredFilterBlockMgr) GetBatchFromHash(hashs []string) (results []*LayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hash IN (?)", hashs).Find(&results).Error

	return
}

// GetFromData 通过data获取内容
func (obj *_LayeredFilterBlockMgr) GetFromData(data string) (results []*LayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("data = ?", data).Find(&results).Error

	return
}

// GetBatchFromData 批量唯一主键查找
func (obj *_LayeredFilterBlockMgr) GetBatchFromData(datas []string) (results []*LayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("data IN (?)", datas).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LayeredFilterBlockMgr) FetchByPrimaryKey(hash string) (result LayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hash = ?", hash).Find(&result).Error

	return
}
