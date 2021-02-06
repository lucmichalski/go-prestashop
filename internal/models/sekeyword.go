package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _SekeywordMgr struct {
	*_BaseMgr
}

// SekeywordMgr open func
func SekeywordMgr(db *gorm.DB) *_SekeywordMgr {
	if db == nil {
		panic(fmt.Errorf("SekeywordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SekeywordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_sekeyword"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SekeywordMgr) GetTableName() string {
	return "ps_sekeyword"
}

// Get 获取
func (obj *_SekeywordMgr) Get() (result Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SekeywordMgr) Gets() (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDSekeyword id_sekeyword获取
func (obj *_SekeywordMgr) WithIDSekeyword(idSekeyword uint32) Option {
	return optionFunc(func(o *options) { o.query["id_sekeyword"] = idSekeyword })
}

// WithIDShop id_shop获取
func (obj *_SekeywordMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDShopGroup id_shop_group获取
func (obj *_SekeywordMgr) WithIDShopGroup(idShopGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_group"] = idShopGroup })
}

// WithKeyword keyword获取
func (obj *_SekeywordMgr) WithKeyword(keyword string) Option {
	return optionFunc(func(o *options) { o.query["keyword"] = keyword })
}

// WithDateAdd date_add获取
func (obj *_SekeywordMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
func (obj *_SekeywordMgr) GetByOption(opts ...Option) (result Sekeyword, err error) {
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
func (obj *_SekeywordMgr) GetByOptions(opts ...Option) (results []*Sekeyword, err error) {
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

// GetFromIDSekeyword 通过id_sekeyword获取内容
func (obj *_SekeywordMgr) GetFromIDSekeyword(idSekeyword uint32) (result Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword = ?", idSekeyword).Find(&result).Error

	return
}

// GetBatchFromIDSekeyword 批量唯一主键查找
func (obj *_SekeywordMgr) GetBatchFromIDSekeyword(idSekeywords []uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword IN (?)", idSekeywords).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_SekeywordMgr) GetFromIDShop(idShop uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_SekeywordMgr) GetBatchFromIDShop(idShops []uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromIDShopGroup 通过id_shop_group获取内容
func (obj *_SekeywordMgr) GetFromIDShopGroup(idShopGroup uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group = ?", idShopGroup).Find(&results).Error

	return
}

// GetBatchFromIDShopGroup 批量唯一主键查找
func (obj *_SekeywordMgr) GetBatchFromIDShopGroup(idShopGroups []uint32) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_group IN (?)", idShopGroups).Find(&results).Error

	return
}

// GetFromKeyword 通过keyword获取内容
func (obj *_SekeywordMgr) GetFromKeyword(keyword string) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keyword = ?", keyword).Find(&results).Error

	return
}

// GetBatchFromKeyword 批量唯一主键查找
func (obj *_SekeywordMgr) GetBatchFromKeyword(keywords []string) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("keyword IN (?)", keywords).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_SekeywordMgr) GetFromDateAdd(dateAdd time.Time) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_SekeywordMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SekeywordMgr) FetchByPrimaryKey(idSekeyword uint32) (result Sekeyword, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_sekeyword = ?", idSekeyword).Find(&result).Error

	return
}
