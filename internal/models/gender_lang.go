package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GenderLangMgr struct {
	*_BaseMgr
}

// GenderLangMgr open func
func GenderLangMgr(db *gorm.DB) *_GenderLangMgr {
	if db == nil {
		panic(fmt.Errorf("GenderLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GenderLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_gender_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GenderLangMgr) GetTableName() string {
	return "eg_gender_lang"
}

// Get 获取
func (obj *_GenderLangMgr) Get() (result GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GenderLangMgr) Gets() (results []*GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGender id_gender获取
func (obj *_GenderLangMgr) WithIDGender(idGender uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gender"] = idGender })
}

// WithIDLang id_lang获取
func (obj *_GenderLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_GenderLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_GenderLangMgr) GetByOption(opts ...Option) (result GenderLang, err error) {
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
func (obj *_GenderLangMgr) GetByOptions(opts ...Option) (results []*GenderLang, err error) {
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

// GetFromIDGender 通过id_gender获取内容
func (obj *_GenderLangMgr) GetFromIDGender(idGender uint32) (results []*GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error

	return
}

// GetBatchFromIDGender 批量唯一主键查找
func (obj *_GenderLangMgr) GetBatchFromIDGender(idGenders []uint32) (results []*GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender IN (?)", idGenders).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_GenderLangMgr) GetFromIDLang(idLang uint32) (results []*GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_GenderLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_GenderLangMgr) GetFromName(name string) (results []*GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_GenderLangMgr) GetBatchFromName(names []string) (results []*GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GenderLangMgr) FetchByPrimaryKey(idGender uint32, idLang uint32) (result GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ? AND id_lang = ?", idGender, idLang).Find(&result).Error

	return
}

// FetchIndexByIDGender  获取多个内容
func (obj *_GenderLangMgr) FetchIndexByIDGender(idGender uint32) (results []*GenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error

	return
}
