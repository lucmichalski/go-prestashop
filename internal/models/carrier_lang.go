package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CarrierLangMgr struct {
	*_BaseMgr
}

// CarrierLangMgr open func
func CarrierLangMgr(db *gorm.DB) *_CarrierLangMgr {
	if db == nil {
		panic(fmt.Errorf("CarrierLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CarrierLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_carrier_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarrierLangMgr) GetTableName() string {
	return "ps_carrier_lang"
}

// Get 获取
func (obj *_CarrierLangMgr) Get() (result CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarrierLangMgr) Gets() (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取
func (obj *_CarrierLangMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDShop id_shop获取
func (obj *_CarrierLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取
func (obj *_CarrierLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithDelay delay获取
func (obj *_CarrierLangMgr) WithDelay(delay string) Option {
	return optionFunc(func(o *options) { o.query["delay"] = delay })
}

// GetByOption 功能选项模式获取
func (obj *_CarrierLangMgr) GetByOption(opts ...Option) (result CarrierLang, err error) {
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
func (obj *_CarrierLangMgr) GetByOptions(opts ...Option) (results []*CarrierLang, err error) {
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

// GetFromIDCarrier 通过id_carrier获取内容
func (obj *_CarrierLangMgr) GetFromIDCarrier(idCarrier uint32) (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_CarrierLangMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_CarrierLangMgr) GetFromIDShop(idShop uint32) (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_CarrierLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_CarrierLangMgr) GetFromIDLang(idLang uint32) (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_CarrierLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromDelay 通过delay获取内容
func (obj *_CarrierLangMgr) GetFromDelay(delay string) (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delay = ?", delay).Find(&results).Error

	return
}

// GetBatchFromDelay 批量唯一主键查找
func (obj *_CarrierLangMgr) GetBatchFromDelay(delays []string) (results []*CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("delay IN (?)", delays).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarrierLangMgr) FetchByPrimaryKey(idCarrier uint32, idShop uint32, idLang uint32) (result CarrierLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_shop = ? AND id_lang = ?", idCarrier, idShop, idLang).Find(&result).Error

	return
}
