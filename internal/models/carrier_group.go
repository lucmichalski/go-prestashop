package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CarrierGroupMgr struct {
	*_BaseMgr
}

// CarrierGroupMgr open func
func CarrierGroupMgr(db *gorm.DB) *_CarrierGroupMgr {
	if db == nil {
		panic(fmt.Errorf("CarrierGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CarrierGroupMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_carrier_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarrierGroupMgr) GetTableName() string {
	return "ps_carrier_group"
}

// Get 获取
func (obj *_CarrierGroupMgr) Get() (result CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarrierGroupMgr) Gets() (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取
func (obj *_CarrierGroupMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDGroup id_group获取
func (obj *_CarrierGroupMgr) WithIDGroup(idGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_group"] = idGroup })
}

// GetByOption 功能选项模式获取
func (obj *_CarrierGroupMgr) GetByOption(opts ...Option) (result CarrierGroup, err error) {
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
func (obj *_CarrierGroupMgr) GetByOptions(opts ...Option) (results []*CarrierGroup, err error) {
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
func (obj *_CarrierGroupMgr) GetFromIDCarrier(idCarrier uint32) (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&results).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_CarrierGroupMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromIDGroup 通过id_group获取内容
func (obj *_CarrierGroupMgr) GetFromIDGroup(idGroup uint32) (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group = ?", idGroup).Find(&results).Error

	return
}

// GetBatchFromIDGroup 批量唯一主键查找
func (obj *_CarrierGroupMgr) GetBatchFromIDGroup(idGroups []uint32) (results []*CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_group IN (?)", idGroups).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarrierGroupMgr) FetchByPrimaryKey(idCarrier uint32, idGroup uint32) (result CarrierGroup, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ? AND id_group = ?", idCarrier, idGroup).Find(&result).Error

	return
}
