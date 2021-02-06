package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsRoleLangMgr struct {
	*_BaseMgr
}

// CmsRoleLangMgr open func
func CmsRoleLangMgr(db *gorm.DB) *_CmsRoleLangMgr {
	if db == nil {
		panic(fmt.Errorf("CmsRoleLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsRoleLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms_role_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CmsRoleLangMgr) GetTableName() string {
	return "eg_cms_role_lang"
}

// Get 获取
func (obj *_CmsRoleLangMgr) Get() (result CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CmsRoleLangMgr) Gets() (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCmsRole id_cms_role获取
func (obj *_CmsRoleLangMgr) WithIDCmsRole(idCmsRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_role"] = idCmsRole })
}

// WithIDLang id_lang获取
func (obj *_CmsRoleLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithIDShop id_shop获取
func (obj *_CmsRoleLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithName name获取
func (obj *_CmsRoleLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_CmsRoleLangMgr) GetByOption(opts ...Option) (result CmsRoleLang, err error) {
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
func (obj *_CmsRoleLangMgr) GetByOptions(opts ...Option) (results []*CmsRoleLang, err error) {
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
func (obj *_CmsRoleLangMgr) GetFromIDCmsRole(idCmsRole uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ?", idCmsRole).Find(&results).Error

	return
}

// GetBatchFromIDCmsRole 批量唯一主键查找
func (obj *_CmsRoleLangMgr) GetBatchFromIDCmsRole(idCmsRoles []uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role IN (?)", idCmsRoles).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_CmsRoleLangMgr) GetFromIDLang(idLang uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_CmsRoleLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_CmsRoleLangMgr) GetFromIDShop(idShop uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_CmsRoleLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_CmsRoleLangMgr) GetFromName(name string) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_CmsRoleLangMgr) GetBatchFromName(names []string) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CmsRoleLangMgr) FetchByPrimaryKey(idCmsRole uint32, idLang uint32, idShop uint32) (result CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ? AND id_lang = ? AND id_shop = ?", idCmsRole, idLang, idShop).Find(&result).Error

	return
}
