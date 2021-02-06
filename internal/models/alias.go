package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgAliasMgr struct {
	*_BaseMgr
}

// EgAliasMgr open func
func EgAliasMgr(db *gorm.DB) *_EgAliasMgr {
	if db == nil {
		panic(fmt.Errorf("EgAliasMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgAliasMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_alias"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgAliasMgr) GetTableName() string {
	return "eg_alias"
}

// Get 获取
func (obj *_EgAliasMgr) Get() (result EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgAliasMgr) Gets() (results []*EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAlias id_alias获取 
func (obj *_EgAliasMgr) WithIDAlias(idAlias uint32) Option {
	return optionFunc(func(o *options) { o.query["id_alias"] = idAlias })
}

// WithAlias alias获取 
func (obj *_EgAliasMgr) WithAlias(alias string) Option {
	return optionFunc(func(o *options) { o.query["alias"] = alias })
}

// WithSearch search获取 
func (obj *_EgAliasMgr) WithSearch(search string) Option {
	return optionFunc(func(o *options) { o.query["search"] = search })
}

// WithActive active获取 
func (obj *_EgAliasMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}


// GetByOption 功能选项模式获取
func (obj *_EgAliasMgr) GetByOption(opts ...Option) (result EgAlias, err error) {
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
func (obj *_EgAliasMgr) GetByOptions(opts ...Option) (results []*EgAlias, err error) {
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


// GetFromIDAlias 通过id_alias获取内容  
func (obj *_EgAliasMgr)  GetFromIDAlias(idAlias uint32) (result EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_alias = ?", idAlias).Find(&result).Error
	
	return
}

// GetBatchFromIDAlias 批量唯一主键查找 
func (obj *_EgAliasMgr) GetBatchFromIDAlias(idAliass []uint32) (results []*EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_alias IN (?)", idAliass).Find(&results).Error
	
	return
}
 
// GetFromAlias 通过alias获取内容  
func (obj *_EgAliasMgr)  GetFromAlias(alias string) (result EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error
	
	return
}

// GetBatchFromAlias 批量唯一主键查找 
func (obj *_EgAliasMgr) GetBatchFromAlias(aliass []string) (results []*EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias IN (?)", aliass).Find(&results).Error
	
	return
}
 
// GetFromSearch 通过search获取内容  
func (obj *_EgAliasMgr) GetFromSearch(search string) (results []*EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("search = ?", search).Find(&results).Error
	
	return
}

// GetBatchFromSearch 批量唯一主键查找 
func (obj *_EgAliasMgr) GetBatchFromSearch(searchs []string) (results []*EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("search IN (?)", searchs).Find(&results).Error
	
	return
}
 
// GetFromActive 通过active获取内容  
func (obj *_EgAliasMgr) GetFromActive(active bool) (results []*EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error
	
	return
}

// GetBatchFromActive 批量唯一主键查找 
func (obj *_EgAliasMgr) GetBatchFromActive(actives []bool) (results []*EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgAliasMgr) FetchByPrimaryKey(idAlias uint32 ) (result EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_alias = ?", idAlias).Find(&result).Error
	
	return
}
 
 // FetchUniqueByAlias primay or index 获取唯一内容
 func (obj *_EgAliasMgr) FetchUniqueByAlias(alias string ) (result EgAlias, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("alias = ?", alias).Find(&result).Error
	
	return
}
 

 

	

