package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgLayeredFilterBlockMgr struct {
	*_BaseMgr
}

// EgLayeredFilterBlockMgr open func
func EgLayeredFilterBlockMgr(db *gorm.DB) *_EgLayeredFilterBlockMgr {
	if db == nil {
		panic(fmt.Errorf("EgLayeredFilterBlockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLayeredFilterBlockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_layered_filter_block"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLayeredFilterBlockMgr) GetTableName() string {
	return "eg_layered_filter_block"
}

// Get 获取
func (obj *_EgLayeredFilterBlockMgr) Get() (result EgLayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLayeredFilterBlockMgr) Gets() (results []*EgLayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithHash hash获取 
func (obj *_EgLayeredFilterBlockMgr) WithHash(hash string) Option {
	return optionFunc(func(o *options) { o.query["hash"] = hash })
}

// WithData data获取 
func (obj *_EgLayeredFilterBlockMgr) WithData(data string) Option {
	return optionFunc(func(o *options) { o.query["data"] = data })
}


// GetByOption 功能选项模式获取
func (obj *_EgLayeredFilterBlockMgr) GetByOption(opts ...Option) (result EgLayeredFilterBlock, err error) {
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
func (obj *_EgLayeredFilterBlockMgr) GetByOptions(opts ...Option) (results []*EgLayeredFilterBlock, err error) {
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


// GetFromHash 通过hash获取内容  
func (obj *_EgLayeredFilterBlockMgr)  GetFromHash(hash string) (result EgLayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hash = ?", hash).Find(&result).Error
	
	return
}

// GetBatchFromHash 批量唯一主键查找 
func (obj *_EgLayeredFilterBlockMgr) GetBatchFromHash(hashs []string) (results []*EgLayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hash IN (?)", hashs).Find(&results).Error
	
	return
}
 
// GetFromData 通过data获取内容  
func (obj *_EgLayeredFilterBlockMgr) GetFromData(data string) (results []*EgLayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("data = ?", data).Find(&results).Error
	
	return
}

// GetBatchFromData 批量唯一主键查找 
func (obj *_EgLayeredFilterBlockMgr) GetBatchFromData(datas []string) (results []*EgLayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("data IN (?)", datas).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLayeredFilterBlockMgr) FetchByPrimaryKey(hash string ) (result EgLayeredFilterBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hash = ?", hash).Find(&result).Error
	
	return
}
 

 

	

