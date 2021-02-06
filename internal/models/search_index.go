package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgSearchIndexMgr struct {
	*_BaseMgr
}

// EgSearchIndexMgr open func
func EgSearchIndexMgr(db *gorm.DB) *_EgSearchIndexMgr {
	if db == nil {
		panic(fmt.Errorf("EgSearchIndexMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSearchIndexMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_search_index"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSearchIndexMgr) GetTableName() string {
	return "eg_search_index"
}

// Get 获取
func (obj *_EgSearchIndexMgr) Get() (result EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSearchIndexMgr) Gets() (results []*EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProduct id_product获取 
func (obj *_EgSearchIndexMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithIDWord id_word获取 
func (obj *_EgSearchIndexMgr) WithIDWord(idWord uint32) Option {
	return optionFunc(func(o *options) { o.query["id_word"] = idWord })
}

// WithWeight weight获取 
func (obj *_EgSearchIndexMgr) WithWeight(weight uint16) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}


// GetByOption 功能选项模式获取
func (obj *_EgSearchIndexMgr) GetByOption(opts ...Option) (result EgSearchIndex, err error) {
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
func (obj *_EgSearchIndexMgr) GetByOptions(opts ...Option) (results []*EgSearchIndex, err error) {
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


// GetFromIDProduct 通过id_product获取内容  
func (obj *_EgSearchIndexMgr) GetFromIDProduct(idProduct uint32) (results []*EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error
	
	return
}

// GetBatchFromIDProduct 批量唯一主键查找 
func (obj *_EgSearchIndexMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error
	
	return
}
 
// GetFromIDWord 通过id_word获取内容  
func (obj *_EgSearchIndexMgr) GetFromIDWord(idWord uint32) (results []*EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word = ?", idWord).Find(&results).Error
	
	return
}

// GetBatchFromIDWord 批量唯一主键查找 
func (obj *_EgSearchIndexMgr) GetBatchFromIDWord(idWords []uint32) (results []*EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word IN (?)", idWords).Find(&results).Error
	
	return
}
 
// GetFromWeight 通过weight获取内容  
func (obj *_EgSearchIndexMgr) GetFromWeight(weight uint16) (results []*EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight = ?", weight).Find(&results).Error
	
	return
}

// GetBatchFromWeight 批量唯一主键查找 
func (obj *_EgSearchIndexMgr) GetBatchFromWeight(weights []uint16) (results []*EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("weight IN (?)", weights).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSearchIndexMgr) FetchByPrimaryKey(idProduct uint32 ,idWord uint32 ) (result EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND id_word = ?", idProduct , idWord).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDProduct  获取多个内容
 func (obj *_EgSearchIndexMgr) FetchIndexByIDProduct(idProduct uint32 ,weight uint16 ) (results []*EgSearchIndex, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ? AND weight = ?", idProduct , weight).Find(&results).Error
	
	return
}
 

	

