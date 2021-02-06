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

func CmsCategoryMgr(db *gorm.DB) *_CmsCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("CmsCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CmsCategoryMgr) GetTableName() string {
	return "ps_cms_category"
}

func (obj *_CmsCategoryMgr) Get() (result CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CmsCategoryMgr) Gets() (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CmsCategoryMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

func (obj *_CmsCategoryMgr) WithIDParent(idParent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_parent"] = idParent })
}

func (obj *_CmsCategoryMgr) WithLevelDepth(levelDepth uint8) Option {
	return optionFunc(func(o *options) { o.query["level_depth"] = levelDepth })
}

func (obj *_CmsCategoryMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_CmsCategoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_CmsCategoryMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_CmsCategoryMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

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


func (obj *_CmsCategoryMgr) GetFromIDCmsCategory(idCmsCategory uint32) (result CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&result).Error

	return
}

func (obj *_CmsCategoryMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetFromIDParent(idParent uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetBatchFromIDParent(idParents []uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent IN (?)", idParents).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetFromLevelDepth(levelDepth uint8) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth = ?", levelDepth).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetBatchFromLevelDepth(levelDepths []uint8) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth IN (?)", levelDepths).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetFromActive(active bool) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetBatchFromActive(actives []bool) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetFromDateUpd(dateUpd time.Time) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetFromPosition(position uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_CmsCategoryMgr) GetBatchFromPosition(positions []uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}


func (obj *_CmsCategoryMgr) FetchByPrimaryKey(idCmsCategory uint32) (result CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&result).Error

	return
}

func (obj *_CmsCategoryMgr) FetchIndexByCategoryParent(idParent uint32) (results []*CmsCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error

	return
}
