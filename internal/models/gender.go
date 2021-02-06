package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _GenderMgr struct {
	*_BaseMgr
}

func GenderMgr(db *gorm.DB) *_GenderMgr {
	if db == nil {
		panic(fmt.Errorf("GenderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GenderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_gender"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_GenderMgr) GetTableName() string {
	return "ps_gender"
}

func (obj *_GenderMgr) Get() (result Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_GenderMgr) Gets() (results []*Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_GenderMgr) WithIDGender(idGender int) Option {
	return optionFunc(func(o *options) { o.query["id_gender"] = idGender })
}

func (obj *_GenderMgr) WithType(_type bool) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

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


func (obj *_GenderMgr) GetFromIDGender(idGender int) (result Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&result).Error

	return
}

func (obj *_GenderMgr) GetBatchFromIDGender(idGenders []int) (results []*Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender IN (?)", idGenders).Find(&results).Error

	return
}

func (obj *_GenderMgr) GetFromType(_type bool) (results []*Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

func (obj *_GenderMgr) GetBatchFromType(_types []bool) (results []*Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}


func (obj *_GenderMgr) FetchByPrimaryKey(idGender int) (result Gender, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&result).Error

	return
}
