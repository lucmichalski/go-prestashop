package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CustomerMessageSyncImapMgr struct {
	*_BaseMgr
}

func CustomerMessageSyncImapMgr(db *gorm.DB) *_CustomerMessageSyncImapMgr {
	if db == nil {
		panic(fmt.Errorf("CustomerMessageSyncImapMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CustomerMessageSyncImapMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_customer_message_sync_imap"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CustomerMessageSyncImapMgr) GetTableName() string {
	return "ps_customer_message_sync_imap"
}

func (obj *_CustomerMessageSyncImapMgr) Get() (result CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CustomerMessageSyncImapMgr) Gets() (results []*CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CustomerMessageSyncImapMgr) WithMd5Header(md5Header []byte) Option {
	return optionFunc(func(o *options) { o.query["md5_header"] = md5Header })
}

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


func (obj *_CustomerMessageSyncImapMgr) GetFromMd5Header(md5Header []byte) (results []*CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("md5_header = ?", md5Header).Find(&results).Error

	return
}

func (obj *_CustomerMessageSyncImapMgr) GetBatchFromMd5Header(md5Headers [][]byte) (results []*CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("md5_header IN (?)", md5Headers).Find(&results).Error

	return
}


func (obj *_CustomerMessageSyncImapMgr) FetchIndexByMd5HeaderIndex(md5Header []byte) (results []*CustomerMessageSyncImap, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("md5_header = ?", md5Header).Find(&results).Error

	return
}
