package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LangShopMgr struct {
	*_BaseMgr
}

// LangShopMgr open func
func LangShopMgr(db *gorm.DB) *_LangShopMgr {
	if db == nil {
		panic(fmt.Errorf("LangShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LangShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_lang_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LangShopMgr) GetTableName() string {
	return "ps_lang_shop"
}

// Get 获取
func (obj *_LangShopMgr) Get() (result LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LangShopMgr) Gets() (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLang id_lang获取
func (obj *_LangShopMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取
func (obj *_LangShopMgr) WithIDShop(idShop int) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_LangShopMgr) GetByOption(opts ...Option) (result LangShop, err error) {
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
func (obj *_LangShopMgr) GetByOptions(opts ...Option) (results []*LangShop, err error) {
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

// GetFromIDLang 通过id_lang获取内容
func (obj *_LangShopMgr) GetFromIDLang(idLang int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_LangShopMgr) GetBatchFromIDLang(idLangs []int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_LangShopMgr) GetFromIDShop(idShop int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_LangShopMgr) GetBatchFromIDShop(idShops []int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LangShopMgr) FetchByPrimaryKey(idLang int, idShop int) (result LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ? AND id_shop = ?", idLang, idShop).Find(&result).Error

	return
}

// FetchIndexByIDXA5D79262BA299860  获取多个内容
func (obj *_LangShopMgr) FetchIndexByIDXA5D79262BA299860(idLang int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// FetchIndexByIDXA5D79262274A50A0  获取多个内容
func (obj *_LangShopMgr) FetchIndexByIDXA5D79262274A50A0(idShop int) (results []*LangShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
