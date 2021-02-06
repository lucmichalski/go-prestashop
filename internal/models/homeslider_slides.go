package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HomesliderSlidesMgr struct {
	*_BaseMgr
}

// HomesliderSlidesMgr open func
func HomesliderSlidesMgr(db *gorm.DB) *_HomesliderSlidesMgr {
	if db == nil {
		panic(fmt.Errorf("HomesliderSlidesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HomesliderSlidesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_homeslider_slides"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HomesliderSlidesMgr) GetTableName() string {
	return "ps_homeslider_slides"
}

// Get 获取
func (obj *_HomesliderSlidesMgr) Get() (result HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_HomesliderSlidesMgr) Gets() (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDHomesliderSlides id_homeslider_slides获取
func (obj *_HomesliderSlidesMgr) WithIDHomesliderSlides(idHomesliderSlides uint32) Option {
	return optionFunc(func(o *options) { o.query["id_homeslider_slides"] = idHomesliderSlides })
}

// WithPosition position获取
func (obj *_HomesliderSlidesMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithActive active获取
func (obj *_HomesliderSlidesMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// GetByOption 功能选项模式获取
func (obj *_HomesliderSlidesMgr) GetByOption(opts ...Option) (result HomesliderSlides, err error) {
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
func (obj *_HomesliderSlidesMgr) GetByOptions(opts ...Option) (results []*HomesliderSlides, err error) {
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

// GetFromIDHomesliderSlides 通过id_homeslider_slides获取内容
func (obj *_HomesliderSlidesMgr) GetFromIDHomesliderSlides(idHomesliderSlides uint32) (result HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&result).Error

	return
}

// GetBatchFromIDHomesliderSlides 批量唯一主键查找
func (obj *_HomesliderSlidesMgr) GetBatchFromIDHomesliderSlides(idHomesliderSlidess []uint32) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides IN (?)", idHomesliderSlidess).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_HomesliderSlidesMgr) GetFromPosition(position uint32) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_HomesliderSlidesMgr) GetBatchFromPosition(positions []uint32) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_HomesliderSlidesMgr) GetFromActive(active bool) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_HomesliderSlidesMgr) GetBatchFromActive(actives []bool) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_HomesliderSlidesMgr) FetchByPrimaryKey(idHomesliderSlides uint32) (result HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&result).Error

	return
}
