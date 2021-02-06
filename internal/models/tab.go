package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TabMgr struct {
	*_BaseMgr
}

// TabMgr open func
func TabMgr(db *gorm.DB) *_TabMgr {
	if db == nil {
		panic(fmt.Errorf("TabMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TabMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_tab"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TabMgr) GetTableName() string {
	return "eg_tab"
}

// Get 获取
func (obj *_TabMgr) Get() (result Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TabMgr) Gets() (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDTab id_tab获取
func (obj *_TabMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

// WithIDParent id_parent获取
func (obj *_TabMgr) WithIDParent(idParent int) Option {
	return optionFunc(func(o *options) { o.query["id_parent"] = idParent })
}

// WithPosition position获取
func (obj *_TabMgr) WithPosition(position int) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithModule module获取
func (obj *_TabMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithClassName class_name获取
func (obj *_TabMgr) WithClassName(className string) Option {
	return optionFunc(func(o *options) { o.query["class_name"] = className })
}

// WithRouteName route_name获取
func (obj *_TabMgr) WithRouteName(routeName string) Option {
	return optionFunc(func(o *options) { o.query["route_name"] = routeName })
}

// WithActive active获取
func (obj *_TabMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithEnabled enabled获取
func (obj *_TabMgr) WithEnabled(enabled bool) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

// WithHideHostMode hide_host_mode获取
func (obj *_TabMgr) WithHideHostMode(hideHostMode bool) Option {
	return optionFunc(func(o *options) { o.query["hide_host_mode"] = hideHostMode })
}

// WithIcon icon获取
func (obj *_TabMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// GetByOption 功能选项模式获取
func (obj *_TabMgr) GetByOption(opts ...Option) (result Tab, err error) {
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
func (obj *_TabMgr) GetByOptions(opts ...Option) (results []*Tab, err error) {
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

// GetFromIDTab 通过id_tab获取内容
func (obj *_TabMgr) GetFromIDTab(idTab int) (result Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&result).Error

	return
}

// GetBatchFromIDTab 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromIDTab(idTabs []int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

// GetFromIDParent 通过id_parent获取内容
func (obj *_TabMgr) GetFromIDParent(idParent int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error

	return
}

// GetBatchFromIDParent 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromIDParent(idParents []int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent IN (?)", idParents).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_TabMgr) GetFromPosition(position int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromPosition(positions []int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容
func (obj *_TabMgr) GetFromModule(module string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromModule(modules []string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}

// GetFromClassName 通过class_name获取内容
func (obj *_TabMgr) GetFromClassName(className string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name = ?", className).Find(&results).Error

	return
}

// GetBatchFromClassName 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromClassName(classNames []string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name IN (?)", classNames).Find(&results).Error

	return
}

// GetFromRouteName 通过route_name获取内容
func (obj *_TabMgr) GetFromRouteName(routeName string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("route_name = ?", routeName).Find(&results).Error

	return
}

// GetBatchFromRouteName 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromRouteName(routeNames []string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("route_name IN (?)", routeNames).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_TabMgr) GetFromActive(active bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromActive(actives []bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromEnabled 通过enabled获取内容
func (obj *_TabMgr) GetFromEnabled(enabled bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enabled = ?", enabled).Find(&results).Error

	return
}

// GetBatchFromEnabled 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromEnabled(enableds []bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enabled IN (?)", enableds).Find(&results).Error

	return
}

// GetFromHideHostMode 通过hide_host_mode获取内容
func (obj *_TabMgr) GetFromHideHostMode(hideHostMode bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide_host_mode = ?", hideHostMode).Find(&results).Error

	return
}

// GetBatchFromHideHostMode 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromHideHostMode(hideHostModes []bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide_host_mode IN (?)", hideHostModes).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容
func (obj *_TabMgr) GetFromIcon(icon string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找
func (obj *_TabMgr) GetBatchFromIcon(icons []string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_TabMgr) FetchByPrimaryKey(idTab int) (result Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&result).Error

	return
}
