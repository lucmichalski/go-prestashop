package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImageLangMgr struct {
	*_BaseMgr
}

// ImageLangMgr open func
func ImageLangMgr(db *gorm.DB) *_ImageLangMgr {
	if db == nil {
		panic(fmt.Errorf("ImageLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImageLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_image_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImageLangMgr) GetTableName() string {
	return "ps_image_lang"
}

// Get 获取
func (obj *_ImageLangMgr) Get() (result ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImageLangMgr) Gets() (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDImage id_image获取
func (obj *_ImageLangMgr) WithIDImage(idImage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image"] = idImage })
}

// WithIDLang id_lang获取
func (obj *_ImageLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithLegend legend获取
func (obj *_ImageLangMgr) WithLegend(legend string) Option {
	return optionFunc(func(o *options) { o.query["legend"] = legend })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDImage 通过id_image获取内容
func (obj *_ImageLangMgr) GetFromIDImage(idImage uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error

	return
}

// GetBatchFromIDImage 批量唯一主键查找
func (obj *_ImageLangMgr) GetBatchFromIDImage(idImages []uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image IN (?)", idImages).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_ImageLangMgr) GetFromIDLang(idLang uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_ImageLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromLegend 通过legend获取内容
func (obj *_ImageLangMgr) GetFromLegend(legend string) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend = ?", legend).Find(&results).Error

	return
}

// GetBatchFromLegend 批量唯一主键查找
func (obj *_ImageLangMgr) GetBatchFromLegend(legends []string) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("legend IN (?)", legends).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ImageLangMgr) FetchByPrimaryKey(idImage uint32, idLang uint32) (result ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ? AND id_lang = ?", idImage, idLang).Find(&result).Error

	return
}

// FetchIndexByIDImage  获取多个内容
func (obj *_ImageLangMgr) FetchIndexByIDImage(idImage uint32) (results []*ImageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image = ?", idImage).Find(&results).Error

	return
}
