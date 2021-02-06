package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LinksmenutopMgr struct {
	*_BaseMgr
}

// LinksmenutopMgr open func
func LinksmenutopMgr(db *gorm.DB) *_LinksmenutopMgr {
	if db == nil {
		panic(fmt.Errorf("LinksmenutopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LinksmenutopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_linksmenutop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LinksmenutopMgr) GetTableName() string {
	return "ps_linksmenutop"
}

// Get 获取
func (obj *_LinksmenutopMgr) Get() (result Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LinksmenutopMgr) Gets() (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLinksmenutop id_linksmenutop获取
func (obj *_LinksmenutopMgr) WithIDLinksmenutop(idLinksmenutop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_linksmenutop"] = idLinksmenutop })
}

// WithIDShop id_shop获取
func (obj *_LinksmenutopMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithNewWindow new_window获取
func (obj *_LinksmenutopMgr) WithNewWindow(newWindow bool) Option {
	return optionFunc(func(o *options) { o.query["new_window"] = newWindow })
}

// GetByOption 功能选项模式获取
func (obj *_LinksmenutopMgr) GetByOption(opts ...Option) (result Linksmenutop, err error) {
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
func (obj *_LinksmenutopMgr) GetByOptions(opts ...Option) (results []*Linksmenutop, err error) {
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

// GetFromIDLinksmenutop 通过id_linksmenutop获取内容
func (obj *_LinksmenutopMgr) GetFromIDLinksmenutop(idLinksmenutop uint32) (result Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ?", idLinksmenutop).Find(&result).Error

	return
}

// GetBatchFromIDLinksmenutop 批量唯一主键查找
func (obj *_LinksmenutopMgr) GetBatchFromIDLinksmenutop(idLinksmenutops []uint32) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop IN (?)", idLinksmenutops).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_LinksmenutopMgr) GetFromIDShop(idShop uint32) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_LinksmenutopMgr) GetBatchFromIDShop(idShops []uint32) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromNewWindow 通过new_window获取内容
func (obj *_LinksmenutopMgr) GetFromNewWindow(newWindow bool) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window = ?", newWindow).Find(&results).Error

	return
}

// GetBatchFromNewWindow 批量唯一主键查找
func (obj *_LinksmenutopMgr) GetBatchFromNewWindow(newWindows []bool) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("new_window IN (?)", newWindows).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_LinksmenutopMgr) FetchByPrimaryKey(idLinksmenutop uint32) (result Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_linksmenutop = ?", idLinksmenutop).Find(&result).Error

	return
}

// FetchIndexByIDShop  获取多个内容
func (obj *_LinksmenutopMgr) FetchIndexByIDShop(idShop uint32) (results []*Linksmenutop, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}
