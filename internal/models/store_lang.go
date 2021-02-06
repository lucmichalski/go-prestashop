package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StoreLangMgr struct {
	*_BaseMgr
}

// StoreLangMgr open func
func StoreLangMgr(db *gorm.DB) *_StoreLangMgr {
	if db == nil {
		panic(fmt.Errorf("StoreLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StoreLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_store_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StoreLangMgr) GetTableName() string {
	return "eg_store_lang"
}

// Get 获取
func (obj *_StoreLangMgr) Get() (result StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StoreLangMgr) Gets() (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDStore id_store获取
func (obj *_StoreLangMgr) WithIDStore(idStore uint32) Option {
	return optionFunc(func(o *options) { o.query["id_store"] = idStore })
}

// WithIDLang id_lang获取
func (obj *_StoreLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_StoreLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAddress1 address1获取
func (obj *_StoreLangMgr) WithAddress1(address1 string) Option {
	return optionFunc(func(o *options) { o.query["address1"] = address1 })
}

// WithAddress2 address2获取
func (obj *_StoreLangMgr) WithAddress2(address2 string) Option {
	return optionFunc(func(o *options) { o.query["address2"] = address2 })
}

// WithHours hours获取
func (obj *_StoreLangMgr) WithHours(hours string) Option {
	return optionFunc(func(o *options) { o.query["hours"] = hours })
}

// WithNote note获取
func (obj *_StoreLangMgr) WithNote(note string) Option {
	return optionFunc(func(o *options) { o.query["note"] = note })
}

// GetByOption 功能选项模式获取
func (obj *_StoreLangMgr) GetByOption(opts ...Option) (result StoreLang, err error) {
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
func (obj *_StoreLangMgr) GetByOptions(opts ...Option) (results []*StoreLang, err error) {
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

// GetFromIDStore 通过id_store获取内容
func (obj *_StoreLangMgr) GetFromIDStore(idStore uint32) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ?", idStore).Find(&results).Error

	return
}

// GetBatchFromIDStore 批量唯一主键查找
func (obj *_StoreLangMgr) GetBatchFromIDStore(idStores []uint32) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store IN (?)", idStores).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_StoreLangMgr) GetFromIDLang(idLang uint32) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_StoreLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_StoreLangMgr) GetFromName(name string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_StoreLangMgr) GetBatchFromName(names []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromAddress1 通过address1获取内容
func (obj *_StoreLangMgr) GetFromAddress1(address1 string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 = ?", address1).Find(&results).Error

	return
}

// GetBatchFromAddress1 批量唯一主键查找
func (obj *_StoreLangMgr) GetBatchFromAddress1(address1s []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address1 IN (?)", address1s).Find(&results).Error

	return
}

// GetFromAddress2 通过address2获取内容
func (obj *_StoreLangMgr) GetFromAddress2(address2 string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 = ?", address2).Find(&results).Error

	return
}

// GetBatchFromAddress2 批量唯一主键查找
func (obj *_StoreLangMgr) GetBatchFromAddress2(address2s []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("address2 IN (?)", address2s).Find(&results).Error

	return
}

// GetFromHours 通过hours获取内容
func (obj *_StoreLangMgr) GetFromHours(hours string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hours = ?", hours).Find(&results).Error

	return
}

// GetBatchFromHours 批量唯一主键查找
func (obj *_StoreLangMgr) GetBatchFromHours(hourss []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hours IN (?)", hourss).Find(&results).Error

	return
}

// GetFromNote 通过note获取内容
func (obj *_StoreLangMgr) GetFromNote(note string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note = ?", note).Find(&results).Error

	return
}

// GetBatchFromNote 批量唯一主键查找
func (obj *_StoreLangMgr) GetBatchFromNote(notes []string) (results []*StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("note IN (?)", notes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_StoreLangMgr) FetchByPrimaryKey(idStore uint32, idLang uint32) (result StoreLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_store = ? AND id_lang = ?", idStore, idLang).Find(&result).Error

	return
}
