package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AdminFilterMgr struct {
	*_BaseMgr
}

func AdminFilterMgr(db *gorm.DB) *_AdminFilterMgr {
	if db == nil {
		panic(fmt.Errorf("AdminFilterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AdminFilterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_admin_filter"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_AdminFilterMgr) GetTableName() string {
	return "ps_admin_filter"
}

func (obj *_AdminFilterMgr) Get() (result AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_AdminFilterMgr) Gets() (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_AdminFilterMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

func (obj *_AdminFilterMgr) WithEmployee(employee int) Option {
	return optionFunc(func(o *options) { o.query["employee"] = employee })
}

func (obj *_AdminFilterMgr) WithShop(shop int) Option {
	return optionFunc(func(o *options) { o.query["shop"] = shop })
}

func (obj *_AdminFilterMgr) WithController(controller string) Option {
	return optionFunc(func(o *options) { o.query["controller"] = controller })
}

func (obj *_AdminFilterMgr) WithAction(action string) Option {
	return optionFunc(func(o *options) { o.query["action"] = action })
}

func (obj *_AdminFilterMgr) WithFilter(filter string) Option {
	return optionFunc(func(o *options) { o.query["filter"] = filter })
}

func (obj *_AdminFilterMgr) WithFilterID(filterID string) Option {
	return optionFunc(func(o *options) { o.query["filter_id"] = filterID })
}

func (obj *_AdminFilterMgr) GetByOption(opts ...Option) (result AdminFilter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_AdminFilterMgr) GetByOptions(opts ...Option) (results []*AdminFilter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_AdminFilterMgr) GetFromID(id int) (result AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

func (obj *_AdminFilterMgr) GetBatchFromID(ids []int) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetFromEmployee(employee int) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee = ?", employee).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetBatchFromEmployee(employees []int) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee IN (?)", employees).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetFromShop(shop int) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop = ?", shop).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetBatchFromShop(shops []int) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shop IN (?)", shops).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetFromController(controller string) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("controller = ?", controller).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetBatchFromController(controllers []string) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("controller IN (?)", controllers).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetFromAction(action string) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("action = ?", action).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetBatchFromAction(actions []string) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("action IN (?)", actions).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetFromFilter(filter string) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter = ?", filter).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetBatchFromFilter(filters []string) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter IN (?)", filters).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetFromFilterID(filterID string) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_id = ?", filterID).Find(&results).Error

	return
}

func (obj *_AdminFilterMgr) GetBatchFromFilterID(filterIDs []string) (results []*AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filter_id IN (?)", filterIDs).Find(&results).Error

	return
}


func (obj *_AdminFilterMgr) FetchByPrimaryKey(id int) (result AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

func (obj *_AdminFilterMgr) FetchUniqueIndexByAdminFilterSearchIDIDx(employee int, shop int, controller string, action string, filterID string) (result AdminFilter, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("employee = ? AND shop = ? AND controller = ? AND action = ? AND filter_id = ?", employee, shop, controller, action, filterID).Find(&result).Error

	return
}
