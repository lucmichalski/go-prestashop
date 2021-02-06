package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SearchWordMgr struct {
	*_BaseMgr
}

// SearchWordMgr open func
func SearchWordMgr(db *gorm.DB) *_SearchWordMgr {
	if db == nil {
		panic(fmt.Errorf("SearchWordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SearchWordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_search_word"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SearchWordMgr) GetTableName() string {
	return "ps_search_word"
}

// Get 获取
func (obj *_SearchWordMgr) Get() (result SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SearchWordMgr) Gets() (results []*SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWord id_word获取
func (obj *_SearchWordMgr) WithIDWord(idWord uint32) Option {
	return optionFunc(func(o *options) { o.query["id_word"] = idWord })
}

// WithIDShop id_shop获取
func (obj *_SearchWordMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取
func (obj *_SearchWordMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithWord word获取
func (obj *_SearchWordMgr) WithWord(word string) Option {
	return optionFunc(func(o *options) { o.query["word"] = word })
}

// GetByOption 功能选项模式获取
func (obj *_SearchWordMgr) GetByOption(opts ...Option) (result SearchWord, err error) {
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
func (obj *_SearchWordMgr) GetByOptions(opts ...Option) (results []*SearchWord, err error) {
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

// GetFromIDWord 通过id_word获取内容
func (obj *_SearchWordMgr) GetFromIDWord(idWord uint32) (result SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word = ?", idWord).Find(&result).Error

	return
}

// GetBatchFromIDWord 批量唯一主键查找
func (obj *_SearchWordMgr) GetBatchFromIDWord(idWords []uint32) (results []*SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word IN (?)", idWords).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_SearchWordMgr) GetFromIDShop(idShop uint32) (results []*SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_SearchWordMgr) GetBatchFromIDShop(idShops []uint32) (results []*SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_SearchWordMgr) GetFromIDLang(idLang uint32) (results []*SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_SearchWordMgr) GetBatchFromIDLang(idLangs []uint32) (results []*SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromWord 通过word获取内容
func (obj *_SearchWordMgr) GetFromWord(word string) (results []*SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("word = ?", word).Find(&results).Error

	return
}

// GetBatchFromWord 批量唯一主键查找
func (obj *_SearchWordMgr) GetBatchFromWord(words []string) (results []*SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("word IN (?)", words).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SearchWordMgr) FetchByPrimaryKey(idWord uint32) (result SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word = ?", idWord).Find(&result).Error

	return
}

// FetchUniqueIndexByIDLang primay or index 获取唯一内容
func (obj *_SearchWordMgr) FetchUniqueIndexByIDLang(idShop uint32, idLang uint32, word string) (result SearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_lang = ? AND word = ?", idShop, idLang, word).Find(&result).Error

	return
}
