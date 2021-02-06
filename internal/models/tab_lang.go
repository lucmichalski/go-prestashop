package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TabLangMgr struct {
	*_BaseMgr
}

// TabLangMgr open func
func TabLangMgr(db *gorm.DB) *_TabLangMgr {
	if db == nil {
		panic(fmt.Errorf("TabLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TabLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tab_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TabLangMgr) GetTableName() string {
	return "ps_tab_lang"
}

// Get 获取
func (obj *_TabLangMgr) Get() (result TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TabLangMgr) Gets() (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTab id_tab获取
func (obj *_TabLangMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

// WithIDLang id_lang获取
func (obj *_TabLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_TabLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_TabLangMgr) GetByOption(opts ...Option) (result TabLang, err error) {
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
func (obj *_TabLangMgr) GetByOptions(opts ...Option) (results []*TabLang, err error) {
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

// GetFromIDTab 通过id_tab获取内容
func (obj *_TabLangMgr) GetFromIDTab(idTab int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

// GetBatchFromIDTab 批量唯一主键查找
func (obj *_TabLangMgr) GetBatchFromIDTab(idTabs []int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_TabLangMgr) GetFromIDLang(idLang int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_TabLangMgr) GetBatchFromIDLang(idLangs []int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TabLangMgr) GetFromName(name string) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_TabLangMgr) GetBatchFromName(names []string) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TabLangMgr) FetchByPrimaryKey(idTab int, idLang int) (result TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ? AND id_lang = ?", idTab, idLang).Find(&result).Error

	return
}

// FetchIndexByIDX3E3D6F36ED47AB56  获取多个内容
func (obj *_TabLangMgr) FetchIndexByIDX3E3D6F36ED47AB56(idTab int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&results).Error

	return
}

// FetchIndexByIDX3E3D6F36BA299860  获取多个内容
func (obj *_TabLangMgr) FetchIndexByIDX3E3D6F36BA299860(idLang int) (results []*TabLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}
