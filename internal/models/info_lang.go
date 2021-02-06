package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _InfoLangMgr struct {
	*_BaseMgr
}

// InfoLangMgr open func
func InfoLangMgr(db *gorm.DB) *_InfoLangMgr {
	if db == nil {
		panic(fmt.Errorf("InfoLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfoLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_info_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_InfoLangMgr) GetTableName() string {
	return "ps_info_lang"
}

// Get 获取
func (obj *_InfoLangMgr) Get() (result InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_InfoLangMgr) Gets() (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDInfo id_info获取
func (obj *_InfoLangMgr) WithIDInfo(idInfo uint32) Option {
	return optionFunc(func(o *options) { o.query["id_info"] = idInfo })
}

// WithIDShop id_shop获取
func (obj *_InfoLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取
func (obj *_InfoLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithText text获取
func (obj *_InfoLangMgr) WithText(text string) Option {
	return optionFunc(func(o *options) { o.query["text"] = text })
}

// GetByOption 功能选项模式获取
func (obj *_InfoLangMgr) GetByOption(opts ...Option) (result InfoLang, err error) {
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
func (obj *_InfoLangMgr) GetByOptions(opts ...Option) (results []*InfoLang, err error) {
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

// GetFromIDInfo 通过id_info获取内容
func (obj *_InfoLangMgr) GetFromIDInfo(idInfo uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&results).Error

	return
}

// GetBatchFromIDInfo 批量唯一主键查找
func (obj *_InfoLangMgr) GetBatchFromIDInfo(idInfos []uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info IN (?)", idInfos).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_InfoLangMgr) GetFromIDShop(idShop uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_InfoLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_InfoLangMgr) GetFromIDLang(idLang uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_InfoLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromText 通过text获取内容
func (obj *_InfoLangMgr) GetFromText(text string) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text = ?", text).Find(&results).Error

	return
}

// GetBatchFromText 批量唯一主键查找
func (obj *_InfoLangMgr) GetBatchFromText(texts []string) (results []*InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("text IN (?)", texts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_InfoLangMgr) FetchByPrimaryKey(idInfo uint32, idShop uint32, idLang uint32) (result InfoLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ? AND id_shop = ? AND id_lang = ?", idInfo, idShop, idLang).Find(&result).Error

	return
}
