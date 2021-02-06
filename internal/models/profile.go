package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ProfileMgr struct {
	*_BaseMgr
}

func ProfileMgr(db *gorm.DB) *_ProfileMgr {
	if db == nil {
		panic(fmt.Errorf("ProfileMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProfileMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_profile"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProfileMgr) GetTableName() string {
	return "ps_profile"
}

func (obj *_ProfileMgr) Get() (result Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProfileMgr) Gets() (results []*Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProfileMgr) WithIDProfile(idProfile uint32) Option {
	return optionFunc(func(o *options) { o.query["id_profile"] = idProfile })
}

func (obj *_ProfileMgr) GetByOption(opts ...Option) (result Profile, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ProfileMgr) GetByOptions(opts ...Option) (results []*Profile, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ProfileMgr) GetFromIDProfile(idProfile uint32) (result Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&result).Error

	return
}

func (obj *_ProfileMgr) GetBatchFromIDProfile(idProfiles []uint32) (results []*Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile IN (?)", idProfiles).Find(&results).Error

	return
}


func (obj *_ProfileMgr) FetchByPrimaryKey(idProfile uint32) (result Profile, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_profile = ?", idProfile).Find(&result).Error

	return
}
