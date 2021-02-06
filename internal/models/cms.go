package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsMgr struct {
	*_BaseMgr
}

func CmsMgr(db *gorm.DB) *_CmsMgr {
	if db == nil {
		panic(fmt.Errorf("CmsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CmsMgr) GetTableName() string {
	return "ps_cms"
}

func (obj *_CmsMgr) Get() (result Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CmsMgr) Gets() (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CmsMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

func (obj *_CmsMgr) WithIDCmsCategory(idCmsCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_category"] = idCmsCategory })
}

func (obj *_CmsMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_CmsMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_CmsMgr) WithIndexation(indexation bool) Option {
	return optionFunc(func(o *options) { o.query["indexation"] = indexation })
}

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


func (obj *_CmsMgr) GetFromIDCms(idCms uint32) (result Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&result).Error

	return
}

func (obj *_CmsMgr) GetBatchFromIDCms(idCmss []uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error

	return
}

func (obj *_CmsMgr) GetFromIDCmsCategory(idCmsCategory uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category = ?", idCmsCategory).Find(&results).Error

	return
}

func (obj *_CmsMgr) GetBatchFromIDCmsCategory(idCmsCategorys []uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_category IN (?)", idCmsCategorys).Find(&results).Error

	return
}

func (obj *_CmsMgr) GetFromPosition(position uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_CmsMgr) GetBatchFromPosition(positions []uint32) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_CmsMgr) GetFromActive(active bool) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_CmsMgr) GetBatchFromActive(actives []bool) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_CmsMgr) GetFromIndexation(indexation bool) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexation = ?", indexation).Find(&results).Error

	return
}

func (obj *_CmsMgr) GetBatchFromIndexation(indexations []bool) (results []*Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("indexation IN (?)", indexations).Find(&results).Error

	return
}


func (obj *_CmsMgr) FetchByPrimaryKey(idCms uint32) (result Cms, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&result).Error

	return
}
