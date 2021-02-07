package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TabMgr struct {
	*_BaseMgr
}

func TabMgr(db *gorm.DB) *_TabMgr {
	if db == nil {
		panic(fmt.Errorf("TabMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TabMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_tab"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_TabMgr) GetTableName() string {
	return "ps_tab"
}

func (obj *_TabMgr) Get() (result Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_TabMgr) Gets() (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_TabMgr) WithIDTab(idTab int) Option {
	return optionFunc(func(o *options) { o.query["id_tab"] = idTab })
}

func (obj *_TabMgr) WithIDParent(idParent int) Option {
	return optionFunc(func(o *options) { o.query["id_parent"] = idParent })
}

func (obj *_TabMgr) WithPosition(position int) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_TabMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

func (obj *_TabMgr) WithClassName(className string) Option {
	return optionFunc(func(o *options) { o.query["class_name"] = className })
}

func (obj *_TabMgr) WithRouteName(routeName string) Option {
	return optionFunc(func(o *options) { o.query["route_name"] = routeName })
}

func (obj *_TabMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_TabMgr) WithEnabled(enabled bool) Option {
	return optionFunc(func(o *options) { o.query["enabled"] = enabled })
}

func (obj *_TabMgr) WithHideHostMode(hideHostMode bool) Option {
	return optionFunc(func(o *options) { o.query["hide_host_mode"] = hideHostMode })
}

func (obj *_TabMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

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

func (obj *_TabMgr) GetFromIDTab(idTab int) (result Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&result).Error

	return
}

func (obj *_TabMgr) GetBatchFromIDTab(idTabs []int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab IN (?)", idTabs).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromIDParent(idParent int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromIDParent(idParents []int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent IN (?)", idParents).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromPosition(position int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromPosition(positions []int) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromModule(module string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromModule(modules []string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromClassName(className string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name = ?", className).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromClassName(classNames []string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("class_name IN (?)", classNames).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromRouteName(routeName string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("route_name = ?", routeName).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromRouteName(routeNames []string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("route_name IN (?)", routeNames).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromActive(active bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromActive(actives []bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromEnabled(enabled bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enabled = ?", enabled).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromEnabled(enableds []bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("enabled IN (?)", enableds).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromHideHostMode(hideHostMode bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide_host_mode = ?", hideHostMode).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromHideHostMode(hideHostModes []bool) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hide_host_mode IN (?)", hideHostModes).Find(&results).Error

	return
}

func (obj *_TabMgr) GetFromIcon(icon string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

func (obj *_TabMgr) GetBatchFromIcon(icons []string) (results []*Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

func (obj *_TabMgr) FetchByPrimaryKey(idTab int) (result Tab, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tab = ?", idTab).Find(&result).Error

	return
}
