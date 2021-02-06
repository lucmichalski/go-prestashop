package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _InfoMgr struct {
	*_BaseMgr
}

func InfoMgr(db *gorm.DB) *_InfoMgr {
	if db == nil {
		panic(fmt.Errorf("InfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_InfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_InfoMgr) GetTableName() string {
	return "ps_info"
}

func (obj *_InfoMgr) Get() (result Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_InfoMgr) Gets() (results []*Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_InfoMgr) WithIDInfo(idInfo uint32) Option {
	return optionFunc(func(o *options) { o.query["id_info"] = idInfo })
}

func (obj *_InfoMgr) GetByOption(opts ...Option) (result Info, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_InfoMgr) GetByOptions(opts ...Option) (results []*Info, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_InfoMgr) GetFromIDInfo(idInfo uint32) (result Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&result).Error

	return
}

func (obj *_InfoMgr) GetBatchFromIDInfo(idInfos []uint32) (results []*Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info IN (?)", idInfos).Find(&results).Error

	return
}


func (obj *_InfoMgr) FetchByPrimaryKey(idInfo uint32) (result Info, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_info = ?", idInfo).Find(&result).Error

	return
}
