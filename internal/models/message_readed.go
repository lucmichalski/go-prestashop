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

// MessageReadedMgr open func
func MessageReadedMgr(db *gorm.DB) *_MessageReadedMgr {
	if db == nil {
		panic(fmt.Errorf("MessageReadedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MessageReadedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_message_readed"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MessageReadedMgr) GetTableName() string {
	return "eg_message_readed"
}

// Get 获取
func (obj *_MessageReadedMgr) Get() (result MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MessageReadedMgr) Gets() (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMessage id_message获取
func (obj *_MessageReadedMgr) WithIDMessage(idMessage uint32) Option {
	return optionFunc(func(o *options) { o.query["id_message"] = idMessage })
}

// WithIDEmployee id_employee获取
func (obj *_MessageReadedMgr) WithIDEmployee(idEmployee uint32) Option {
	return optionFunc(func(o *options) { o.query["id_employee"] = idEmployee })
}

// WithDateAdd date_add获取
func (obj *_MessageReadedMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDMessage 通过id_message获取内容
func (obj *_MessageReadedMgr) GetFromIDMessage(idMessage uint32) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ?", idMessage).Find(&results).Error

	return
}

// GetBatchFromIDMessage 批量唯一主键查找
func (obj *_MessageReadedMgr) GetBatchFromIDMessage(idMessages []uint32) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message IN (?)", idMessages).Find(&results).Error

	return
}

// GetFromIDEmployee 通过id_employee获取内容
func (obj *_MessageReadedMgr) GetFromIDEmployee(idEmployee uint32) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee = ?", idEmployee).Find(&results).Error

	return
}

// GetBatchFromIDEmployee 批量唯一主键查找
func (obj *_MessageReadedMgr) GetBatchFromIDEmployee(idEmployees []uint32) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_employee IN (?)", idEmployees).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_MessageReadedMgr) GetFromDateAdd(dateAdd time.Time) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_MessageReadedMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MessageReadedMgr) FetchByPrimaryKey(idMessage uint32, idEmployee uint32) (result MessageReaded, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_message = ? AND id_employee = ?", idMessage, idEmployee).Find(&result).Error

	return
}
