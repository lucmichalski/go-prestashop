package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CmsCategoryMgr struct {
	*_BaseMgr
}

// CmsCategoryMgr open func
func CmsCategoryMgr(db *gorm.DB) *_CmsCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("CmsCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CmsCategoryMgr) GetTableName() string {
	return "ps_cms_category"
}

// Get 获取
func (obj *_CmsCategoryMgr) Get() (result CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CmsCategoryMgr) Gets() (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCmsCategory id_cms_category获取
func (obj *_CmsCategoryMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

// WithIDParent id_parent获取
func (obj *_CmsCategoryMgr) WithIDParent(idParent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_parent"] = idParent })
}

// WithLevelDepth level_depth获取
func (obj *_CmsCategoryMgr) WithLevelDepth(levelDepth uint8) Option {
	return optionFunc(func(o *options) { o.query["level_depth"] = levelDepth })
}

// WithActive active获取
func (obj *_CmsCategoryMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDateAdd date_add获取
func (obj *_CmsCategoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_CmsCategoryMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// WithPosition position获取
func (obj *_CmsCategoryMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// GetByOption 功能选项模式获取
func (obj *_CmsCategoryMgr) GetByOption(opts ...Option) (result CmsCategory, err error) {
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
func (obj *_CmsCategoryMgr) GetByOptions(opts ...Option) (results []*CmsCategory, err error) {
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

// GetFromIDCmsCategory 通过id_cms_category获取内容
func (obj *_CmsCategoryMgr) GetFromIDCmsCategory(idCmsCategory uint32) (result CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&result).Error

	return
}

// GetBatchFromIDCmsCategory 批量唯一主键查找
func (obj *_CmsCategoryMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error

	return
}

// GetFromIDParent 通过id_parent获取内容
func (obj *_CmsCategoryMgr) GetFromIDParent(idParent uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error

	return
}

// GetBatchFromIDParent 批量唯一主键查找
func (obj *_CmsCategoryMgr) GetBatchFromIDParent(idParents []uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent IN (?)", idParents).Find(&results).Error

	return
}

// GetFromLevelDepth 通过level_depth获取内容
func (obj *_CmsCategoryMgr) GetFromLevelDepth(levelDepth uint8) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth = ?", levelDepth).Find(&results).Error

	return
}

// GetBatchFromLevelDepth 批量唯一主键查找
func (obj *_CmsCategoryMgr) GetBatchFromLevelDepth(levelDepths []uint8) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth IN (?)", levelDepths).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_CmsCategoryMgr) GetFromActive(active bool) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_CmsCategoryMgr) GetBatchFromActive(actives []bool) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_CmsCategoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_CmsCategoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_CmsCategoryMgr) GetFromDateUpd(dateUpd time.Time) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_CmsCategoryMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_CmsCategoryMgr) GetFromPosition(position uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_CmsCategoryMgr) GetBatchFromPosition(positions []uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CmsCategoryMgr) FetchByPrimaryKey(idCmsCategory uint32) (result CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&result).Error

	return
}

// FetchIndexByCategoryParent  获取多个内容
func (obj *_CmsCategoryMgr) FetchIndexByCategoryParent(idParent uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error

	return
}
