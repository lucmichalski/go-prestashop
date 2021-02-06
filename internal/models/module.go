package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ModuleMgr struct {
	*_BaseMgr
}

// ModuleMgr open func
func ModuleMgr(db *gorm.DB) *_ModuleMgr {
	if db == nil {
		panic(fmt.Errorf("ModuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ModuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_module"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ModuleMgr) GetTableName() string {
	return "eg_module"
}

// Get 获取
func (obj *_ModuleMgr) Get() (result Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ModuleMgr) Gets() (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDModule id_module获取
func (obj *_ModuleMgr) WithIDModule(idModule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_module"] = idModule })
}

// WithName name获取
func (obj *_ModuleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithActive active获取
func (obj *_ModuleMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithVersion version获取
func (obj *_ModuleMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_ModuleMgr) GetByOption(opts ...Option) (result Module, err error) {
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
func (obj *_ModuleMgr) GetByOptions(opts ...Option) (results []*Module, err error) {
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

// GetFromIDModule 通过id_module获取内容
func (obj *_ModuleMgr) GetFromIDModule(idModule uint32) (result Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&result).Error

	return
}

// GetBatchFromIDModule 批量唯一主键查找
func (obj *_ModuleMgr) GetBatchFromIDModule(idModules []uint32) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module IN (?)", idModules).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ModuleMgr) GetFromName(name string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ModuleMgr) GetBatchFromName(names []string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_ModuleMgr) GetFromActive(active bool) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_ModuleMgr) GetBatchFromActive(actives []bool) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_ModuleMgr) GetFromVersion(version string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("version = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量唯一主键查找
func (obj *_ModuleMgr) GetBatchFromVersion(versions []string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("version IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ModuleMgr) FetchByPrimaryKey(idModule uint32) (result Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_module = ?", idModule).Find(&result).Error

	return
}

// FetchUniqueByNameUNIQUE primay or index 获取唯一内容
func (obj *_ModuleMgr) FetchUniqueByNameUNIQUE(name string) (result Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}

// FetchIndexByName  获取多个内容
func (obj *_ModuleMgr) FetchIndexByName(name string) (results []*Module, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
