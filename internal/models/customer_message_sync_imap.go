package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomerMessageSyncImapMgr struct {
	*_BaseMgr
}

// CustomerMessageSyncImapMgr open func
func CustomerMessageSyncImapMgr(db *gorm.DB) *_CustomerMessageSyncImapMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerMessageSyncImapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerMessageSyncImapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_customer_message_sync_imap"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CustomerMessageSyncImapMgr) GetTableName() string {
	return "eg_customer_message_sync_imap"
}

// Get 获取
func (obj *_CustomerMessageSyncImapMgr) Get() (result CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CustomerMessageSyncImapMgr) Gets() (results []*CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithMd5Header md5_header获取
func (obj *_CustomerMessageSyncImapMgr) WithMd5Header(md5Header []byte) Option {
	return optionFunc(func(o *options) { o.query["md5_header"] = md5Header })
}

// GetByOption 功能选项模式获取
func (obj *_CustomerMessageSyncImapMgr) GetByOption(opts ...Option) (result CustomerMessageSyncImap, err error) {
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
func (obj *_CustomerMessageSyncImapMgr) GetByOptions(opts ...Option) (results []*CustomerMessageSyncImap, err error) {
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

// GetFromMd5Header 通过md5_header获取内容
func (obj *_CustomerMessageSyncImapMgr) GetFromMd5Header(md5Header []byte) (results []*CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("md5_header = ?", md5Header).Find(&results).Error

	return
}

// GetBatchFromMd5Header 批量唯一主键查找
func (obj *_CustomerMessageSyncImapMgr) GetBatchFromMd5Header(md5Headers [][]byte) (results []*CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("md5_header IN (?)", md5Headers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByMd5HeaderIndex  获取多个内容
func (obj *_CustomerMessageSyncImapMgr) FetchIndexByMd5HeaderIndex(md5Header []byte) (results []*CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("md5_header = ?", md5Header).Find(&results).Error

	return
}
