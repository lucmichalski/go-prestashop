package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SearchIndexMgr struct {
	*_BaseMgr
}

// SearchIndexMgr open func
func SearchIndexMgr(db *gorm.DB) *_SearchIndexMgr {
	if db == nil {
		panic(fmt.Errorf("SearchIndexMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SearchIndexMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_search_index"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SearchIndexMgr) GetTableName() string {
	return "eg_search_index"
}

// Get 获取
func (obj *_SearchIndexMgr) Get() (result SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SearchIndexMgr) Gets() (results []*SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取
func (obj *_SearchIndexMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDWord id_word获取
func (obj *_SearchIndexMgr) WithIDWord(idWord uint32) Option {
	return optionFunc(func(o *options) { o.query["id_word"] = idWord })
}

// WithWeight weight获取
func (obj *_SearchIndexMgr) WithWeight(weight uint16) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// GetByOption 功能选项模式获取
func (obj *_SearchIndexMgr) GetByOption(opts ...Option) (result SearchIndex, err error) {
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
func (obj *_SearchIndexMgr) GetByOptions(opts ...Option) (results []*SearchIndex, err error) {
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

// GetFromIDProduct 通过id_product获取内容
func (obj *_SearchIndexMgr) GetFromIDProduct(idProduct uint32) (results []*SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_SearchIndexMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDWord 通过id_word获取内容
func (obj *_SearchIndexMgr) GetFromIDWord(idWord uint32) (results []*SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word = ?", idWord).Find(&results).Error

	return
}

// GetBatchFromIDWord 批量唯一主键查找
func (obj *_SearchIndexMgr) GetBatchFromIDWord(idWords []uint32) (results []*SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word IN (?)", idWords).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容
func (obj *_SearchIndexMgr) GetFromWeight(weight uint16) (results []*SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量唯一主键查找
func (obj *_SearchIndexMgr) GetBatchFromWeight(weights []uint16) (results []*SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SearchIndexMgr) FetchByPrimaryKey(idProduct uint32, idWord uint32) (result SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_word = ?", idProduct, idWord).Find(&result).Error

	return
}

// FetchIndexByIDProduct  获取多个内容
func (obj *_SearchIndexMgr) FetchIndexByIDProduct(idProduct uint32, weight uint16) (results []*SearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND weight = ?", idProduct, weight).Find(&results).Error

	return
}
