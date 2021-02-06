package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImportMatchMgr struct {
	*_BaseMgr
}

// ImportMatchMgr open func
func ImportMatchMgr(db *gorm.DB) *_ImportMatchMgr {
	if db == nil {
		panic(fmt.Errorf("ImportMatchMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImportMatchMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_import_match"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImportMatchMgr) GetTableName() string {
	return "eg_import_match"
}

// Get 获取
func (obj *_ImportMatchMgr) Get() (result ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImportMatchMgr) Gets() (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDImportMatch id_import_match获取
func (obj *_ImportMatchMgr) WithIDImportMatch(idImportMatch int) Option {
	return optionFunc(func(o *options) { o.query["id_import_match"] = idImportMatch })
}

// WithName name获取
func (obj *_ImportMatchMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithMatch match获取
func (obj *_ImportMatchMgr) WithMatch(match string) Option {
	return optionFunc(func(o *options) { o.query["match"] = match })
}

// WithSkip skip获取
func (obj *_ImportMatchMgr) WithSkip(skip int) Option {
	return optionFunc(func(o *options) { o.query["skip"] = skip })
}

// GetByOption 功能选项模式获取
func (obj *_ImportMatchMgr) GetByOption(opts ...Option) (result ImportMatch, err error) {
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
func (obj *_ImportMatchMgr) GetByOptions(opts ...Option) (results []*ImportMatch, err error) {
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

// GetFromIDImportMatch 通过id_import_match获取内容
func (obj *_ImportMatchMgr) GetFromIDImportMatch(idImportMatch int) (result ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_import_match = ?", idImportMatch).Find(&result).Error

	return
}

// GetBatchFromIDImportMatch 批量唯一主键查找
func (obj *_ImportMatchMgr) GetBatchFromIDImportMatch(idImportMatchs []int) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_import_match IN (?)", idImportMatchs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ImportMatchMgr) GetFromName(name string) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ImportMatchMgr) GetBatchFromName(names []string) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromMatch 通过match获取内容
func (obj *_ImportMatchMgr) GetFromMatch(match string) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("match = ?", match).Find(&results).Error

	return
}

// GetBatchFromMatch 批量唯一主键查找
func (obj *_ImportMatchMgr) GetBatchFromMatch(matchs []string) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("match IN (?)", matchs).Find(&results).Error

	return
}

// GetFromSkip 通过skip获取内容
func (obj *_ImportMatchMgr) GetFromSkip(skip int) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("skip = ?", skip).Find(&results).Error

	return
}

// GetBatchFromSkip 批量唯一主键查找
func (obj *_ImportMatchMgr) GetBatchFromSkip(skips []int) (results []*ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("skip IN (?)", skips).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ImportMatchMgr) FetchByPrimaryKey(idImportMatch int) (result ImportMatch, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_import_match = ?", idImportMatch).Find(&result).Error

	return
}
