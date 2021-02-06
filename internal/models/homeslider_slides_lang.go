package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _HomesliderSlidesLangMgr struct {
	*_BaseMgr
}

func HomesliderSlidesLangMgr(db *gorm.DB) *_HomesliderSlidesLangMgr {
	if db == nil {
		panic(fmt.Errorf("HomesliderSlidesLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HomesliderSlidesLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_homeslider_slides_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_HomesliderSlidesLangMgr) GetTableName() string {
	return "ps_homeslider_slides_lang"
}

func (obj *_HomesliderSlidesLangMgr) Get() (result HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) Gets() (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_HomesliderSlidesLangMgr) WithIDHomesliderSlides(idHomesliderSlides uint32) Option {
	return optionFunc(func(o *options) { o.query["id_homeslider_slides"] = idHomesliderSlides })
}

func (obj *_HomesliderSlidesLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_HomesliderSlidesLangMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

func (obj *_HomesliderSlidesLangMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

func (obj *_HomesliderSlidesLangMgr) WithLegend(legend string) Option {
	return optionFunc(func(o *options) { o.query["legend"] = legend })
}

func (obj *_HomesliderSlidesLangMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

func (obj *_HomesliderSlidesLangMgr) WithImage(image string) Option {
	return optionFunc(func(o *options) { o.query["image"] = image })
}

func (obj *_HomesliderSlidesLangMgr) GetByOption(opts ...Option) (result HomesliderSlidesLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetByOptions(opts ...Option) (results []*HomesliderSlidesLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_HomesliderSlidesLangMgr) GetFromIDHomesliderSlides(idHomesliderSlides uint32) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ?", idHomesliderSlides).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetBatchFromIDHomesliderSlides(idHomesliderSlidess []uint32) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides IN (?)", idHomesliderSlidess).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetFromIDLang(idLang uint32) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetFromTitle(title string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetBatchFromTitle(titles []string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetFromDescription(description string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetBatchFromDescription(descriptions []string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetFromLegend(legend string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend = ?", legend).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetBatchFromLegend(legends []string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend IN (?)", legends).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetFromURL(url string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetBatchFromURL(urls []string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetFromImage(image string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("image = ?", image).Find(&results).Error

	return
}

func (obj *_HomesliderSlidesLangMgr) GetBatchFromImage(images []string) (results []*HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("image IN (?)", images).Find(&results).Error

	return
}


func (obj *_HomesliderSlidesLangMgr) FetchByPrimaryKey(idHomesliderSlides uint32, idLang uint32) (result HomesliderSlidesLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_homeslider_slides = ? AND id_lang = ?", idHomesliderSlides, idLang).Find(&result).Error

	return
}
