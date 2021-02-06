package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ZoneMgr struct {
	*_BaseMgr
}

// ZoneMgr open func
func ZoneMgr(db *gorm.DB) *_ZoneMgr {
	if db == nil {
		panic(fmt.Errorf("ZoneMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZoneMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_zone"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZoneMgr) GetTableName() string {
	return "eg_zone"
}

// Get 获取
func (obj *_ZoneMgr) Get() (result Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZoneMgr) Gets() (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDZone id_zone获取
func (obj *_ZoneMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

// WithName name获取
func (obj *_ZoneMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithActive active获取
func (obj *_ZoneMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// GetByOption 功能选项模式获取
func (obj *_ZoneMgr) GetByOption(opts ...Option) (result Zone, err error) {
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
func (obj *_ZoneMgr) GetByOptions(opts ...Option) (results []*Zone, err error) {
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

// GetFromIDZone 通过id_zone获取内容
func (obj *_ZoneMgr) GetFromIDZone(idZone uint32) (result Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&result).Error

	return
}

// GetBatchFromIDZone 批量唯一主键查找
func (obj *_ZoneMgr) GetBatchFromIDZone(idZones []uint32) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ZoneMgr) GetFromName(name string) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ZoneMgr) GetBatchFromName(names []string) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_ZoneMgr) GetFromActive(active bool) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_ZoneMgr) GetBatchFromActive(actives []bool) (results []*Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ZoneMgr) FetchByPrimaryKey(idZone uint32) (result Zone, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&result).Error

	return
}
