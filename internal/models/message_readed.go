package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _MessageReadedMgr struct {
	*_BaseMgr
}

func MessageReadedMgr(db *gorm.DB) *_MessageReadedMgr {
	if db == nil {
		panic(fmt.Errorf("MessageReadedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MessageReadedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_message_readed"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_MessageReadedMgr) GetTableName() string {
	return "ps_message_readed"
}

func (obj *_MessageReadedMgr) Get() (result MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_MessageReadedMgr) Gets() (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_MessageReadedMgr) WithIDMessage(idMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_message"] = idMessage })
}

func (obj *_MessageReadedMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

func (obj *_MessageReadedMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_MessageReadedMgr) GetByOption(opts ...Option) (result MessageReaded, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_MessageReadedMgr) GetByOptions(opts ...Option) (results []*MessageReaded, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_MessageReadedMgr) GetFromIDMessage(idMessage uint32) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ?", idMessage).Find(&results).Error

	return
}

func (obj *_MessageReadedMgr) GetBatchFromIDMessage(idMessages []uint32) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message IN (?)", idMessages).Find(&results).Error

	return
}

func (obj *_MessageReadedMgr) GetFromIDEmployee(idEmployee uint32) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

func (obj *_MessageReadedMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

func (obj *_MessageReadedMgr) GetFromDateAdd(dateAdd time.Time) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_MessageReadedMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_MessageReadedMgr) FetchByPrimaryKey(idMessage uint32, idEmployee uint32) (result MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ? AND id_employee = ?", idMessage, idEmployee).Find(&result).Error

	return
}
