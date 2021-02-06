package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _MemcachedServersMgr struct {
	*_BaseMgr
}

func MemcachedServersMgr(db *gorm.DB) *_MemcachedServersMgr {
	if db == nil {
		panic(fmt.Errorf("MemcachedServersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MemcachedServersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_memcached_servers"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_MemcachedServersMgr) GetTableName() string {
	return "ps_memcached_servers"
}

func (obj *_MemcachedServersMgr) Get() (result MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_MemcachedServersMgr) Gets() (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_MemcachedServersMgr) WithIDMemcachedServer(idMemcachedServer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_memcached_server"] = idMemcachedServer })
}

func (obj *_MemcachedServersMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

func (obj *_MemcachedServersMgr) WithPort(port uint32) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

func (obj *_MemcachedServersMgr) WithWeight(weight uint32) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

func (obj *_MemcachedServersMgr) GetByOption(opts ...Option) (result MemcachedServers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_MemcachedServersMgr) GetByOptions(opts ...Option) (results []*MemcachedServers, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_MemcachedServersMgr) GetFromIDMemcachedServer(idMemcachedServer uint32) (result MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server = ?", idMemcachedServer).Find(&result).Error

	return
}

func (obj *_MemcachedServersMgr) GetBatchFromIDMemcachedServer(idMemcachedServers []uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server IN (?)", idMemcachedServers).Find(&results).Error

	return
}

func (obj *_MemcachedServersMgr) GetFromIP(ip string) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip = ?", ip).Find(&results).Error

	return
}

func (obj *_MemcachedServersMgr) GetBatchFromIP(ips []string) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip IN (?)", ips).Find(&results).Error

	return
}

func (obj *_MemcachedServersMgr) GetFromPort(port uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("port = ?", port).Find(&results).Error

	return
}

func (obj *_MemcachedServersMgr) GetBatchFromPort(ports []uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("port IN (?)", ports).Find(&results).Error

	return
}

func (obj *_MemcachedServersMgr) GetFromWeight(weight uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

func (obj *_MemcachedServersMgr) GetBatchFromWeight(weights []uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}


func (obj *_MemcachedServersMgr) FetchByPrimaryKey(idMemcachedServer uint32) (result MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server = ?", idMemcachedServer).Find(&result).Error

	return
}
