package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _EgGenderLangMgr struct {
	*_BaseMgr
}

// EgGenderLangMgr open func
func EgGenderLangMgr(db *gorm.DB) *_EgGenderLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgGenderLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgGenderLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_gender_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgGenderLangMgr) GetTableName() string {
	return "eg_gender_lang"
}

// Get 获取
func (obj *_EgGenderLangMgr) Get() (result EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgGenderLangMgr) Gets() (results []*EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDGender id_gender获取 
func (obj *_EgGenderLangMgr) WithIDGender(idGender uint32) Option {
	return optionFunc(func(o *options) { o.query["id_gender"] = idGender })
}

// WithIDLang id_lang获取 
func (obj *_EgGenderLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgGenderLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgGenderLangMgr) GetByOption(opts ...Option) (result EgGenderLang, err error) {
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
func (obj *_EgGenderLangMgr) GetByOptions(opts ...Option) (results []*EgGenderLang, err error) {
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


// GetFromIDGender 通过id_gender获取内容  
func (obj *_EgGenderLangMgr) GetFromIDGender(idGender uint32) (results []*EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error
	
	return
}

// GetBatchFromIDGender 批量唯一主键查找 
func (obj *_EgGenderLangMgr) GetBatchFromIDGender(idGenders []uint32) (results []*EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender IN (?)", idGenders).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgGenderLangMgr) GetFromIDLang(idLang uint32) (results []*EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgGenderLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgGenderLangMgr) GetFromName(name string) (results []*EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgGenderLangMgr) GetBatchFromName(names []string) (results []*EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgGenderLangMgr) FetchByPrimaryKey(idGender uint32 ,idLang uint32 ) (result EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ? AND id_lang = ?", idGender , idLang).Find(&result).Error
	
	return
}
 

 
 // FetchIndexByIDGender  获取多个内容
 func (obj *_EgGenderLangMgr) FetchIndexByIDGender(idGender uint32 ) (results []*EgGenderLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_gender = ?", idGender).Find(&results).Error
	
	return
}
 

	

