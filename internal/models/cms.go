package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsMgr struct {
	*_BaseMgr
}

// CmsMgr open func
func CmsMgr(db *gorm.DB) *_CmsMgr {
	if db == nil {
		panic(fmt.Errorf("CmsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CmsMgr) GetTableName() string {
	return "ps_cms"
}

// Get 获取
func (obj *_CmsMgr) Get() (result Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CmsMgr) Gets() (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCms id_cms获取
func (obj *_CmsMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

// WithIDCmsCategory id_cms_category获取
func (obj *_CmsMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

// WithPosition position获取
func (obj *_CmsMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithActive active获取
func (obj *_CmsMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIndexation indexation获取
func (obj *_CmsMgr) WithIndexation(indexation bool) Option {
	return optionFunc(func(o *options) { o.query["indexation"] = indexation })
}

// GetByOption 功能选项模式获取
func (obj *_CmsMgr) GetByOption(opts ...Option) (result Cms, err error) {
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
func (obj *_CmsMgr) GetByOptions(opts ...Option) (results []*Cms, err error) {
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

// GetFromIDCms 通过id_cms获取内容
func (obj *_CmsMgr) GetFromIDCms(idCms uint32) (result Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&result).Error

	return
}

// GetBatchFromIDCms 批量唯一主键查找
func (obj *_CmsMgr) GetBatchFromIDCms(idCmss []uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error

	return
}

// GetFromIDCmsCategory 通过id_cms_category获取内容
func (obj *_CmsMgr) GetFromIDCmsCategory(idCmsCategory uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&results).Error

	return
}

// GetBatchFromIDCmsCategory 批量唯一主键查找
func (obj *_CmsMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_CmsMgr) GetFromPosition(position uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_CmsMgr) GetBatchFromPosition(positions []uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_CmsMgr) GetFromActive(active bool) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_CmsMgr) GetBatchFromActive(actives []bool) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromIndexation 通过indexation获取内容
func (obj *_CmsMgr) GetFromIndexation(indexation bool) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexation = ?", indexation).Find(&results).Error

	return
}

// GetBatchFromIndexation 批量唯一主键查找
func (obj *_CmsMgr) GetBatchFromIndexation(indexations []bool) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexation IN (?)", indexations).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CmsMgr) FetchByPrimaryKey(idCms uint32) (result Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&result).Error

	return
}
