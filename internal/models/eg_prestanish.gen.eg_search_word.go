package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgSearchWordMgr struct {
	*_BaseMgr
}

// EgSearchWordMgr open func
func EgSearchWordMgr(db *gorm.DB) *_EgSearchWordMgr {
	if db == nil {
		panic(fmt.Errorf("EgSearchWordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgSearchWordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_search_word"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgSearchWordMgr) GetTableName() string {
	return "eg_search_word"
}

// Get 获取
func (obj *_EgSearchWordMgr) Get() (result EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgSearchWordMgr) Gets() (results []*EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDWord id_word获取 
func (obj *_EgSearchWordMgr) WithIDWord(idWord uint32) Option {
	return optionFunc(func(o *options) { o.query["id_word"] = idWord })
}

// WithIDShop id_shop获取 
func (obj *_EgSearchWordMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithIDLang id_lang获取 
func (obj *_EgSearchWordMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithWord word获取 
func (obj *_EgSearchWordMgr) WithWord(word string) Option {
	return optionFunc(func(o *options) { o.query["word"] = word })
}


// GetByOption 功能选项模式获取
func (obj *_EgSearchWordMgr) GetByOption(opts ...Option) (result EgSearchWord, err error) {
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
func (obj *_EgSearchWordMgr) GetByOptions(opts ...Option) (results []*EgSearchWord, err error) {
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


// GetFromIDWord 通过id_word获取内容  
func (obj *_EgSearchWordMgr)  GetFromIDWord(idWord uint32) (result EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word = ?", idWord).Find(&result).Error
	
	return
}

// GetBatchFromIDWord 批量唯一主键查找 
func (obj *_EgSearchWordMgr) GetBatchFromIDWord(idWords []uint32) (results []*EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word IN (?)", idWords).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgSearchWordMgr) GetFromIDShop(idShop uint32) (results []*EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgSearchWordMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgSearchWordMgr) GetFromIDLang(idLang uint32) (results []*EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgSearchWordMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromWord 通过word获取内容  
func (obj *_EgSearchWordMgr) GetFromWord(word string) (results []*EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("word = ?", word).Find(&results).Error
	
	return
}

// GetBatchFromWord 批量唯一主键查找 
func (obj *_EgSearchWordMgr) GetBatchFromWord(words []string) (results []*EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("word IN (?)", words).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgSearchWordMgr) FetchByPrimaryKey(idWord uint32 ) (result EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_word = ?", idWord).Find(&result).Error
	
	return
}
 
 // FetchUniqueIndexByIDLang primay or index 获取唯一内容
 func (obj *_EgSearchWordMgr) FetchUniqueIndexByIDLang(idShop uint32 ,idLang uint32 ,word string ) (result EgSearchWord, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ? AND id_lang = ? AND word = ?", idShop , idLang , word).Find(&result).Error
	
	return
}
 

 

	

