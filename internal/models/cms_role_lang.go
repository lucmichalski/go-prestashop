package	model	
import (	
"context"	
"gorm.io/gorm"	
"fmt"	
)	

type _EgCmsRoleLangMgr struct {
	*_BaseMgr
}

// EgCmsRoleLangMgr open func
func EgCmsRoleLangMgr(db *gorm.DB) *_EgCmsRoleLangMgr {
	if db == nil {
		panic(fmt.Errorf("EgCmsRoleLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EgCmsRoleLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms_role_lang"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EgCmsRoleLangMgr) GetTableName() string {
	return "eg_cms_role_lang"
}

// Get 获取
func (obj *_EgCmsRoleLangMgr) Get() (result EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_EgCmsRoleLangMgr) Gets() (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error
	
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCmsRole id_cms_role获取 
func (obj *_EgCmsRoleLangMgr) WithIDCmsRole(idCmsRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_role"] = idCmsRole })
}

// WithIDLang id_lang获取 
func (obj *_EgCmsRoleLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取 
func (obj *_EgCmsRoleLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithName name获取 
func (obj *_EgCmsRoleLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}


// GetByOption 功能选项模式获取
func (obj *_EgCmsRoleLangMgr) GetByOption(opts ...Option) (result EgCmsRoleLang, err error) {
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
func (obj *_EgCmsRoleLangMgr) GetByOptions(opts ...Option) (results []*EgCmsRoleLang, err error) {
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


// GetFromIDCmsRole 通过id_cms_role获取内容  
func (obj *_EgCmsRoleLangMgr) GetFromIDCmsRole(idCmsRole uint32) (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ?", idCmsRole).Find(&results).Error
	
	return
}

// GetBatchFromIDCmsRole 批量唯一主键查找 
func (obj *_EgCmsRoleLangMgr) GetBatchFromIDCmsRole(idCmsRoles []uint32) (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role IN (?)", idCmsRoles).Find(&results).Error
	
	return
}
 
// GetFromIDLang 通过id_lang获取内容  
func (obj *_EgCmsRoleLangMgr) GetFromIDLang(idLang uint32) (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error
	
	return
}

// GetBatchFromIDLang 批量唯一主键查找 
func (obj *_EgCmsRoleLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error
	
	return
}
 
// GetFromIDShop 通过id_shop获取内容  
func (obj *_EgCmsRoleLangMgr) GetFromIDShop(idShop uint32) (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error
	
	return
}

// GetBatchFromIDShop 批量唯一主键查找 
func (obj *_EgCmsRoleLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容  
func (obj *_EgCmsRoleLangMgr) GetFromName(name string) (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量唯一主键查找 
func (obj *_EgCmsRoleLangMgr) GetBatchFromName(names []string) (results []*EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primay or index 获取唯一内容
 func (obj *_EgCmsRoleLangMgr) FetchByPrimaryKey(idCmsRole uint32 ,idLang uint32 ,idShop uint32 ) (result EgCmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ? AND id_lang = ? AND id_shop = ?", idCmsRole , idLang , idShop).Find(&result).Error
	
	return
}
 

 

	

