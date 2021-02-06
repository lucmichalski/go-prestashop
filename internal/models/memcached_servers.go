package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _MemcachedServersMgr struct {
	*_BaseMgr
}

// MemcachedServersMgr open func
func MemcachedServersMgr(db *gorm.DB) *_MemcachedServersMgr {
	if db == nil {
		panic(fmt.Errorf("MemcachedServersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MemcachedServersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_memcached_servers"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MemcachedServersMgr) GetTableName() string {
	return "ps_memcached_servers"
}

// Get 获取
func (obj *_MemcachedServersMgr) Get() (result MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MemcachedServersMgr) Gets() (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMemcachedServer id_memcached_server获取
func (obj *_MemcachedServersMgr) WithIDMemcachedServer(idMemcachedServer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_memcached_server"] = idMemcachedServer })
}

// WithIP ip获取
func (obj *_MemcachedServersMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithPort port获取
func (obj *_MemcachedServersMgr) WithPort(port uint32) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithWeight weight获取
func (obj *_MemcachedServersMgr) WithWeight(weight uint32) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDMemcachedServer 通过id_memcached_server获取内容
func (obj *_MemcachedServersMgr) GetFromIDMemcachedServer(idMemcachedServer uint32) (result MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server = ?", idMemcachedServer).Find(&result).Error

	return
}

// GetBatchFromIDMemcachedServer 批量唯一主键查找
func (obj *_MemcachedServersMgr) GetBatchFromIDMemcachedServer(idMemcachedServers []uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server IN (?)", idMemcachedServers).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容
func (obj *_MemcachedServersMgr) GetFromIP(ip string) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量唯一主键查找
func (obj *_MemcachedServersMgr) GetBatchFromIP(ips []string) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip IN (?)", ips).Find(&results).Error

	return
}

// GetFromPort 通过port获取内容
func (obj *_MemcachedServersMgr) GetFromPort(port uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("port = ?", port).Find(&results).Error

	return
}

// GetBatchFromPort 批量唯一主键查找
func (obj *_MemcachedServersMgr) GetBatchFromPort(ports []uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("port IN (?)", ports).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容
func (obj *_MemcachedServersMgr) GetFromWeight(weight uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量唯一主键查找
func (obj *_MemcachedServersMgr) GetBatchFromWeight(weights []uint32) (results []*MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MemcachedServersMgr) FetchByPrimaryKey(idMemcachedServer uint32) (result MemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server = ?", idMemcachedServer).Find(&result).Error

	return
}
