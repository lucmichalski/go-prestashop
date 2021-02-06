package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderReturnStateLangMgr struct {
	*_BaseMgr
}

// OrderReturnStateLangMgr open func
func OrderReturnStateLangMgr(db *gorm.DB) *_OrderReturnStateLangMgr {
	if db == nil {
		panic(fmt.Errorf("OrderReturnStateLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderReturnStateLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_order_return_state_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderReturnStateLangMgr) GetTableName() string {
	return "eg_order_return_state_lang"
}

// Get 获取
func (obj *_OrderReturnStateLangMgr) Get() (result OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderReturnStateLangMgr) Gets() (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderReturnState id_order_return_state获取
func (obj *_OrderReturnStateLangMgr) WithIDOrderReturnState(idOrderReturnState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_return_state"] = idOrderReturnState })
}

// WithIDLang id_lang获取
func (obj *_OrderReturnStateLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_OrderReturnStateLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_OrderReturnStateLangMgr) GetByOption(opts ...Option) (result OrderReturnStateLang, err error) {
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
func (obj *_OrderReturnStateLangMgr) GetByOptions(opts ...Option) (results []*OrderReturnStateLang, err error) {
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

// GetFromIDOrderReturnState 通过id_order_return_state获取内容
func (obj *_OrderReturnStateLangMgr) GetFromIDOrderReturnState(idOrderReturnState uint32) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ?", idOrderReturnState).Find(&results).Error

	return
}

// GetBatchFromIDOrderReturnState 批量唯一主键查找
func (obj *_OrderReturnStateLangMgr) GetBatchFromIDOrderReturnState(idOrderReturnStates []uint32) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state IN (?)", idOrderReturnStates).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_OrderReturnStateLangMgr) GetFromIDLang(idLang uint32) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_OrderReturnStateLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_OrderReturnStateLangMgr) GetFromName(name string) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_OrderReturnStateLangMgr) GetBatchFromName(names []string) (results []*OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderReturnStateLangMgr) FetchByPrimaryKey(idOrderReturnState uint32, idLang uint32) (result OrderReturnStateLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_return_state = ? AND id_lang = ?", idOrderReturnState, idLang).Find(&result).Error

	return
}
