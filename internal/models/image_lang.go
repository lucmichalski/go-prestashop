package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImageLangMgr struct {
	*_BaseMgr
}

func ImageLangMgr(db *gorm.DB) *_ImageLangMgr {
	if db == nil {
		panic(fmt.Errorf("ImageLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImageLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_image_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ImageLangMgr) GetTableName() string {
	return "ps_image_lang"
}

func (obj *_ImageLangMgr) Get() (result ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ImageLangMgr) Gets() (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_ImageLangMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

func (obj *_ImageLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_ImageLangMgr) WithLegend(legend string) Option {
	return optionFunc(func(o *options) { o.query["legend"] = legend })
}

func (obj *_ImageLangMgr) GetByOption(opts ...Option) (result ImageLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ImageLangMgr) GetByOptions(opts ...Option) (results []*ImageLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_ImageLangMgr) GetFromIDImage(idImage uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error

	return
}

func (obj *_ImageLangMgr) GetBatchFromIDImage(idImages []uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error

	return
}

func (obj *_ImageLangMgr) GetFromIDLang(idLang uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_ImageLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_ImageLangMgr) GetFromLegend(legend string) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend = ?", legend).Find(&results).Error

	return
}

func (obj *_ImageLangMgr) GetBatchFromLegend(legends []string) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend IN (?)", legends).Find(&results).Error

	return
}

func (obj *_ImageLangMgr) FetchByPrimaryKey(idImage uint32, idLang uint32) (result ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_lang = ?", idImage, idLang).Find(&result).Error

	return
}

func (obj *_ImageLangMgr) FetchIndexByIDImage(idImage uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error

	return
}
