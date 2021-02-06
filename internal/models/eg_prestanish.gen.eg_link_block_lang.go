package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _EgLinkBlockLangMgr struct {
	*_BaseMgr
}

// EgLinkBlockLangMgr open func
func EgLinkBlockLangMgr(db *gorm.DB) *_EgLinkBlockLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgLinkBlockLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgLinkBlockLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_link_block_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgLinkBlockLangMgr) GetTableName() string {
	return "eg_link_block_lang"
}

// Get 获取
func (obj *_EgLinkBlockLangMgr) Get() (result EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgLinkBlockLangMgr) Gets() (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDLinkBlock id_link_block获取 
func (obj *_EgLinkBlockLangMgr) WithIDLinkBlock(idLinkBlock uint32) Option {
	return optionFunc(func(o *options) { o.query["id_link_block"] = idLinkBlock })
}

// WithIDLang id_lang获取 
func (obj *_EgLinkBlockLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取 
func (obj *_EgLinkBlockLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCustomContent custom_content获取 
func (obj *_EgLinkBlockLangMgr) WithCustomContent(customContent string) Option {
	return optionFunc(func(o *options) { o.query["custom_content"] = customContent })
}


// GetByOption 功能选项模式获取
func (obj *_EgLinkBlockLangMgr) GetByOption(opts ...Option) (result EgLinkBlockLang, err error) {
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
func (obj *_EgLinkBlockLangMgr) GetByOptions(opts ...Option) (results []*EgLinkBlockLang, err error) {
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


// GetFromIDLinkBlock 通过id_link_block获取内容  
func (obj *_EgLinkBlockLangMgr) GetFromIDLinkBlock(idLinkBlock uint32) (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ?", idLinkBlock).Find(&results).Error
	
	return
}

// GetBatchFromIDLinkBlock 批量唯一主键查找 
func (obj *_EgLinkBlockLangMgr) GetBatchFromIDLinkBlock(idLinkBlocks []uint32) (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block IN (?)", idLinkBlocks).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgLinkBlockLangMgr) GetFromIDLang(idLang uint32) (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgLinkBlockLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgLinkBlockLangMgr) GetFromName(name string) (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgLinkBlockLangMgr) GetBatchFromName(names []string) (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromCustomContent 通过custom_content获取内容  
func (obj *_EgLinkBlockLangMgr) GetFromCustomContent(customContent string) (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom_content = ?", customContent).Find(&results).Error
	
	return
}

// GetBatchFromCustomContent 批量唯一主键查找 
func (obj *_EgLinkBlockLangMgr) GetBatchFromCustomContent(customContents []string) (results []*EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom_content IN (?)", customContents).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgLinkBlockLangMgr) FetchByPrimaryKey(idLinkBlock uint32 ,idLang uint32 ) (result EgLinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ? AND id_lang = ?", idLinkBlock , idLang).Find(&result).Error
	
	return
}
 

 

	

