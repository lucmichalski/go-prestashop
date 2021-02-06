package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _OrderMessageLangMgr struct {
	*_BaseMgr
}

// OrderMessageLangMgr open func
func OrderMessageLangMgr(db *gorm.DB) *_OrderMessageLangMgr {
	if db == nil {
		panic(fmt.Errorf("OrderMessageLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderMessageLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_order_message_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderMessageLangMgr) GetTableName() string {
	return "ps_order_message_lang"
}

// Get 获取
func (obj *_OrderMessageLangMgr) Get() (result OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderMessageLangMgr) Gets() (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDOrderMessage id_order_message获取
func (obj *_OrderMessageLangMgr) WithIDOrderMessage(idOrderMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_order_message"] = idOrderMessage })
}

// WithIDLang id_lang获取
func (obj *_OrderMessageLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_OrderMessageLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithMessage message获取
func (obj *_OrderMessageLangMgr) WithMessage(message string) Option {
	return optionFunc(func(o *options) { o.query["message"] = message })
}

// GetByOption 功能选项模式获取
func (obj *_OrderMessageLangMgr) GetByOption(opts ...Option) (result OrderMessageLang, err error) {
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
func (obj *_OrderMessageLangMgr) GetByOptions(opts ...Option) (results []*OrderMessageLang, err error) {
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

// GetFromIDOrderMessage 通过id_order_message获取内容
func (obj *_OrderMessageLangMgr) GetFromIDOrderMessage(idOrderMessage uint32) (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message = ?", idOrderMessage).Find(&results).Error

	return
}

// GetBatchFromIDOrderMessage 批量唯一主键查找
func (obj *_OrderMessageLangMgr) GetBatchFromIDOrderMessage(idOrderMessages []uint32) (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message IN (?)", idOrderMessages).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_OrderMessageLangMgr) GetFromIDLang(idLang uint32) (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_OrderMessageLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_OrderMessageLangMgr) GetFromName(name string) (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_OrderMessageLangMgr) GetBatchFromName(names []string) (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromMessage 通过message获取内容
func (obj *_OrderMessageLangMgr) GetFromMessage(message string) (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message = ?", message).Find(&results).Error

	return
}

// GetBatchFromMessage 批量唯一主键查找
func (obj *_OrderMessageLangMgr) GetBatchFromMessage(messages []string) (results []*OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("message IN (?)", messages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_OrderMessageLangMgr) FetchByPrimaryKey(idOrderMessage uint32, idLang uint32) (result OrderMessageLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_order_message = ? AND id_lang = ?", idOrderMessage, idLang).Find(&result).Error

	return
}
