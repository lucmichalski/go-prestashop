package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AddressFormatMgr struct {
	*_BaseMgr
}

// AddressFormatMgr open func
func AddressFormatMgr(db *gorm.DB) *_AddressFormatMgr {
	if db == nil {
		panic(fmt.Errorf("AddressFormatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AddressFormatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_address_format"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AddressFormatMgr) GetTableName() string {
	return "eg_address_format"
}

// Get 获取
func (obj *_AddressFormatMgr) Get() (result AddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AddressFormatMgr) Gets() (results []*AddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCountry id_country获取
func (obj *_AddressFormatMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

// WithFormat format获取
func (obj *_AddressFormatMgr) WithFormat(format string) Option {
	return optionFunc(func(o *options) { o.query["format"] = format })
}

// GetByOption 功能选项模式获取
func (obj *_AddressFormatMgr) GetByOption(opts ...Option) (result AddressFormat, err error) {
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
func (obj *_AddressFormatMgr) GetByOptions(opts ...Option) (results []*AddressFormat, err error) {
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

// GetFromIDCountry 通过id_country获取内容
func (obj *_AddressFormatMgr) GetFromIDCountry(idCountry uint32) (result AddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&result).Error

	return
}

// GetBatchFromIDCountry 批量唯一主键查找
func (obj *_AddressFormatMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*AddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

// GetFromFormat 通过format获取内容
func (obj *_AddressFormatMgr) GetFromFormat(format string) (results []*AddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("format = ?", format).Find(&results).Error

	return
}

// GetBatchFromFormat 批量唯一主键查找
func (obj *_AddressFormatMgr) GetBatchFromFormat(formats []string) (results []*AddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("format IN (?)", formats).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AddressFormatMgr) FetchByPrimaryKey(idCountry uint32) (result AddressFormat, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&result).Error

	return
}
