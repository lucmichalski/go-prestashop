package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GenderMgr struct {
	*_BaseMgr
}

// GenderMgr open func
func GenderMgr(db *gorm.DB) *_GenderMgr {
	if db == nil {
		panic(fmt.Errorf("GenderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GenderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_gender"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GenderMgr) GetTableName() string {
	return "eg_gender"
}

// Get 获取
func (obj *_GenderMgr) Get() (result Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GenderMgr) Gets() (results []*Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGender id_gender获取
func (obj *_GenderMgr) WithIDGender(idGender int) Option {
	return optionFunc(func(o *options) { o.query["id_gender"] = idGender })
}

// WithType type获取
func (obj *_GenderMgr) WithType(_type bool) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_GenderMgr) GetByOption(opts ...Option) (result Gender, err error) {
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
func (obj *_GenderMgr) GetByOptions(opts ...Option) (results []*Gender, err error) {
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

// GetFromIDGender 通过id_gender获取内容
func (obj *_GenderMgr) GetFromIDGender(idGender int) (result Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&result).Error

	return
}

// GetBatchFromIDGender 批量唯一主键查找
func (obj *_GenderMgr) GetBatchFromIDGender(idGenders []int) (results []*Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender IN (?)", idGenders).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_GenderMgr) GetFromType(_type bool) (results []*Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_GenderMgr) GetBatchFromType(_types []bool) (results []*Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_GenderMgr) FetchByPrimaryKey(idGender int) (result Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&result).Error

	return
}
