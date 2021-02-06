package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CartCartRuleMgr struct {
	*_BaseMgr
}

// CartCartRuleMgr open func
func CartCartRuleMgr(db *gorm.DB) *_CartCartRuleMgr {
	if db == nil {
		panic(fmt.Errorf("CartCartRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CartCartRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cart_cart_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CartCartRuleMgr) GetTableName() string {
	return "eg_cart_cart_rule"
}

// Get 获取
func (obj *_CartCartRuleMgr) Get() (result CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CartCartRuleMgr) Gets() (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCart id_cart获取
func (obj *_CartCartRuleMgr) WithIDCart(idCart uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart"] = idCart })
}

// WithIDCartRule id_cart_rule获取
func (obj *_CartCartRuleMgr) WithIDCartRule(idCartRule uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cart_rule"] = idCartRule })
}

// GetByOption 功能选项模式获取
func (obj *_CartCartRuleMgr) GetByOption(opts ...Option) (result CartCartRule, err error) {
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
func (obj *_CartCartRuleMgr) GetByOptions(opts ...Option) (results []*CartCartRule, err error) {
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

// GetFromIDCart 通过id_cart获取内容
func (obj *_CartCartRuleMgr) GetFromIDCart(idCart uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ?", idCart).Find(&results).Error

	return
}

// GetBatchFromIDCart 批量唯一主键查找
func (obj *_CartCartRuleMgr) GetBatchFromIDCart(idCarts []uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart IN (?)", idCarts).Find(&results).Error

	return
}

// GetFromIDCartRule 通过id_cart_rule获取内容
func (obj *_CartCartRuleMgr) GetFromIDCartRule(idCartRule uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}

// GetBatchFromIDCartRule 批量唯一主键查找
func (obj *_CartCartRuleMgr) GetBatchFromIDCartRule(idCartRules []uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule IN (?)", idCartRules).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CartCartRuleMgr) FetchByPrimaryKey(idCart uint32, idCartRule uint32) (result CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart = ? AND id_cart_rule = ?", idCart, idCartRule).Find(&result).Error

	return
}

// FetchIndexByIDCartRule  获取多个内容
func (obj *_CartCartRuleMgr) FetchIndexByIDCartRule(idCartRule uint32) (results []*CartCartRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cart_rule = ?", idCartRule).Find(&results).Error

	return
}
