package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TagMgr struct {
	*_BaseMgr
}

// TagMgr open func
func TagMgr(db *gorm.DB) *_TagMgr {
	if db == nil {
		panic(fmt.Errorf("TagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TagMgr) GetTableName() string {
	return "ps_tag"
}

// Get 获取
func (obj *_TagMgr) Get() (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TagMgr) Gets() (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTag id_tag获取
func (obj *_TagMgr) WithIDTag(idTag uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tag"] = idTag })
}

// WithIDLang id_lang获取
func (obj *_TagMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_TagMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_TagMgr) GetByOption(opts ...Option) (result Tag, err error) {
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
func (obj *_TagMgr) GetByOptions(opts ...Option) (results []*Tag, err error) {
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

// GetFromIDTag 通过id_tag获取内容
func (obj *_TagMgr) GetFromIDTag(idTag uint32) (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&result).Error

	return
}

// GetBatchFromIDTag 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromIDTag(idTags []uint32) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag IN (?)", idTags).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_TagMgr) GetFromIDLang(idLang uint32) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromIDLang(idLangs []uint32) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TagMgr) GetFromName(name string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_TagMgr) GetBatchFromName(names []string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TagMgr) FetchByPrimaryKey(idTag uint32) (result Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tag = ?", idTag).Find(&result).Error

	return
}

// FetchIndexByIDLang  获取多个内容
func (obj *_TagMgr) FetchIndexByIDLang(idLang uint32) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// FetchIndexByTagName  获取多个内容
func (obj *_TagMgr) FetchIndexByTagName(name string) (results []*Tag, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
