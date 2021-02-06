package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HomesliderSlidesMgr struct {
	*_BaseMgr
}

func HomesliderSlidesMgr(db *gorm.DB) *_HomesliderSlidesMgr {
	if db == nil {
		panic(fmt.Errorf("HomesliderSlidesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HomesliderSlidesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_homeslider_slides"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_HomesliderSlidesMgr) GetTableName() string {
	return "ps_homeslider_slides"
}

func (obj *_HomesliderSlidesMgr) Get() (result HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_HomesliderSlidesMgr) Gets() (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_HomesliderSlidesMgr) WithIDHomesliderSlides(idHomesliderSlides uint32) Option {
	return optionFunc(func(o *options) { o.query["id_homeslider_slides"] = idHomesliderSlides })
}

func (obj *_HomesliderSlidesMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_HomesliderSlidesMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

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


func (obj *_HomesliderSlidesMgr) GetFromIDHomesliderSlides(idHomesliderSlides uint32) (result HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&result).Error

	return
}

func (obj *_HomesliderSlidesMgr) GetBatchFromIDHomesliderSlides(idHomesliderSlidess []uint32) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides IN (?)", idHomesliderSlidess).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesMgr) GetFromPosition(position uint32) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesMgr) GetBatchFromPosition(positions []uint32) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesMgr) GetFromActive(active bool) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesMgr) GetBatchFromActive(actives []bool) (results []*HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}


func (obj *_HomesliderSlidesMgr) FetchByPrimaryKey(idHomesliderSlides uint32) (result HomesliderSlides, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&result).Error

	return
}
