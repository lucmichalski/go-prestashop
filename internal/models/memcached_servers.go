package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgMemcachedServersMgr struct {
	*_BaseMgr
}

// EgMemcachedServersMgr open func
func EgMemcachedServersMgr(db *gorm.DB) *_EgMemcachedServersMgr {
	if db == nil {
		panic(fmt.Errorf("EgMemcachedServersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgMemcachedServersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_memcached_servers"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgMemcachedServersMgr) GetTableName() string {
	return "eg_memcached_servers"
}

// Get 获取
func (obj *_EgMemcachedServersMgr) Get() (result EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgMemcachedServersMgr) Gets() (results []*EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDMemcachedServer id_memcached_server获取 
func (obj *_EgMemcachedServersMgr) WithIDMemcachedServer(idMemcachedServer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_memcached_server"] = idMemcachedServer })
}

// WithIP ip获取 
func (obj *_EgMemcachedServersMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithPort port获取 
func (obj *_EgMemcachedServersMgr) WithPort(port uint32) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithWeight weight获取 
func (obj *_EgMemcachedServersMgr) WithWeight(weight uint32) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}


// GetByOption 功能选项模式获取
func (obj *_EgMemcachedServersMgr) GetByOption(opts ...Option) (result EgMemcachedServers, err error) {
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
func (obj *_EgMemcachedServersMgr) GetByOptions(opts ...Option) (results []*EgMemcachedServers, err error) {
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
func (obj *_EgMemcachedServersMgr)  GetFromIDMemcachedServer(idMemcachedServer uint32) (result EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server = ?", idMemcachedServer).Find(&result).Error
	
	return
}

// GetBatchFromIDMemcachedServer 批量唯一主键查找 
func (obj *_EgMemcachedServersMgr) GetBatchFromIDMemcachedServer(idMemcachedServers []uint32) (results []*EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server IN (?)", idMemcachedServers).Find(&results).Error
	
	return
}
 
// GetFromIP 通过ip获取内容  
func (obj *_EgMemcachedServersMgr) GetFromIP(ip string) (results []*EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量唯一主键查找 
func (obj *_EgMemcachedServersMgr) GetBatchFromIP(ips []string) (results []*EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("ip IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过port获取内容  
func (obj *_EgMemcachedServersMgr) GetFromPort(port uint32) (results []*EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("port = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量唯一主键查找 
func (obj *_EgMemcachedServersMgr) GetBatchFromPort(ports []uint32) (results []*EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("port IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromWeight 通过weight获取内容  
func (obj *_EgMemcachedServersMgr) GetFromWeight(weight uint32) (results []*EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error
	
	return
}

// GetBatchFromWeight 批量唯一主键查找 
func (obj *_EgMemcachedServersMgr) GetBatchFromWeight(weights []uint32) (results []*EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgMemcachedServersMgr) FetchByPrimaryKey(idMemcachedServer uint32 ) (result EgMemcachedServers, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_memcached_server = ?", idMemcachedServer).Find(&result).Error
	
	return
}
 

 

	

