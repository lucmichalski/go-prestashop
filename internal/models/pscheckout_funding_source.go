package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PscheckoutFundingSourceMgr struct {
	*_BaseMgr
}

// PscheckoutFundingSourceMgr open func
func PscheckoutFundingSourceMgr(db *gorm.DB) *_PscheckoutFundingSourceMgr {
	if db == nil {
		panic(fmt.Errorf("PscheckoutFundingSourceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PscheckoutFundingSourceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_pscheckout_funding_source"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PscheckoutFundingSourceMgr) GetTableName() string {
	return "ps_pscheckout_funding_source"
}

// Get 获取
func (obj *_PscheckoutFundingSourceMgr) Get() (result PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PscheckoutFundingSourceMgr) Gets() (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithName name获取
func (obj *_PscheckoutFundingSourceMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithActive active获取
func (obj *_PscheckoutFundingSourceMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithPosition position获取
func (obj *_PscheckoutFundingSourceMgr) WithPosition(position uint8) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithIDShop id_shop获取
func (obj *_PscheckoutFundingSourceMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// GetByOption 功能选项模式获取
func (obj *_PscheckoutFundingSourceMgr) GetByOption(opts ...Option) (result PscheckoutFundingSource, err error) {
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
func (obj *_PscheckoutFundingSourceMgr) GetByOptions(opts ...Option) (results []*PscheckoutFundingSource, err error) {
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

// GetFromName 通过name获取内容
func (obj *_PscheckoutFundingSourceMgr) GetFromName(name string) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_PscheckoutFundingSourceMgr) GetBatchFromName(names []string) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_PscheckoutFundingSourceMgr) GetFromActive(active bool) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_PscheckoutFundingSourceMgr) GetBatchFromActive(actives []bool) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_PscheckoutFundingSourceMgr) GetFromPosition(position uint8) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_PscheckoutFundingSourceMgr) GetBatchFromPosition(positions []uint8) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_PscheckoutFundingSourceMgr) GetFromIDShop(idShop uint32) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_PscheckoutFundingSourceMgr) GetBatchFromIDShop(idShops []uint32) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PscheckoutFundingSourceMgr) FetchByPrimaryKey(name string, idShop uint32) (result PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ? AND id_shop = ?", name, idShop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_PscheckoutFundingSourceMgr) FetchIndexByIDShop(idShop uint32) (results []*PscheckoutFundingSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
