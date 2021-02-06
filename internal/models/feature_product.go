package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _FeatureProductMgr struct {
	*_BaseMgr
}

// FeatureProductMgr open func
func FeatureProductMgr(db *gorm.DB) *_FeatureProductMgr {
	if db == nil {
		panic(fmt.Errorf("FeatureProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FeatureProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_feature_product"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FeatureProductMgr) GetTableName() string {
	return "ps_feature_product"
}

// Get 获取
func (obj *_FeatureProductMgr) Get() (result FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FeatureProductMgr) Gets() (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDFeature id_feature获取
func (obj *_FeatureProductMgr) WithIDFeature(idFeature uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature"] = idFeature })
}

// WithIDProduct id_product获取
func (obj *_FeatureProductMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDFeatureValue id_feature_value获取
func (obj *_FeatureProductMgr) WithIDFeatureValue(idFeatureValue uint32) Option {
	return optionFunc(func(o *options) { o.query["id_feature_value"] = idFeatureValue })
}

// GetByOption 功能选项模式获取
func (obj *_FeatureProductMgr) GetByOption(opts ...Option) (result FeatureProduct, err error) {
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
func (obj *_FeatureProductMgr) GetByOptions(opts ...Option) (results []*FeatureProduct, err error) {
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

// GetFromIDFeature 通过id_feature获取内容
func (obj *_FeatureProductMgr) GetFromIDFeature(idFeature uint32) (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ?", idFeature).Find(&results).Error

	return
}

// GetBatchFromIDFeature 批量唯一主键查找
func (obj *_FeatureProductMgr) GetBatchFromIDFeature(idFeatures []uint32) (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature IN (?)", idFeatures).Find(&results).Error

	return
}

// GetFromIDProduct 通过id_product获取内容
func (obj *_FeatureProductMgr) GetFromIDProduct(idProduct uint32) (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_FeatureProductMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromIDFeatureValue 通过id_feature_value获取内容
func (obj *_FeatureProductMgr) GetFromIDFeatureValue(idFeatureValue uint32) (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&results).Error

	return
}

// GetBatchFromIDFeatureValue 批量唯一主键查找
func (obj *_FeatureProductMgr) GetBatchFromIDFeatureValue(idFeatureValues []uint32) (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value IN (?)", idFeatureValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_FeatureProductMgr) FetchByPrimaryKey(idFeature uint32, idProduct uint32, idFeatureValue uint32) (result FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature = ? AND id_product = ? AND id_feature_value = ?", idFeature, idProduct, idFeatureValue).Find(&result).Error

	return
}

// FetchIndexByIDProduct  获取多个内容
func (obj *_FeatureProductMgr) FetchIndexByIDProduct(idProduct uint32) (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// FetchIndexByIDFeatureValue  获取多个内容
func (obj *_FeatureProductMgr) FetchIndexByIDFeatureValue(idFeatureValue uint32) (results []*FeatureProduct, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_feature_value = ?", idFeatureValue).Find(&results).Error

	return
}
