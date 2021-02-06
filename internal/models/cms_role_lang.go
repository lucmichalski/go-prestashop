package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsRoleLangMgr struct {
	*_BaseMgr
}

func CmsRoleLangMgr(db *gorm.DB) *_CmsRoleLangMgr {
	if db == nil {
		panic(fmt.Errorf("CmsRoleLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsRoleLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms_role_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CmsRoleLangMgr) GetTableName() string {
	return "ps_cms_role_lang"
}

func (obj *_CmsRoleLangMgr) Get() (result CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CmsRoleLangMgr) Gets() (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CmsRoleLangMgr) WithIDCmsRole(idCmsRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_role"] = idCmsRole })
}

func (obj *_CmsRoleLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_CmsRoleLangMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

func (obj *_CmsRoleLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

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


func (obj *_CmsRoleLangMgr) GetFromIDCmsRole(idCmsRole uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ?", idCmsRole).Find(&results).Error

	return
}

func (obj *_CmsRoleLangMgr) GetBatchFromIDCmsRole(idCmsRoles []uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role IN (?)", idCmsRoles).Find(&results).Error

	return
}

func (obj *_CmsRoleLangMgr) GetFromIDLang(idLang uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_CmsRoleLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_CmsRoleLangMgr) GetFromIDShop(idShop uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

func (obj *_CmsRoleLangMgr) GetBatchFromIDShop(idShops []uint32) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

func (obj *_CmsRoleLangMgr) GetFromName(name string) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_CmsRoleLangMgr) GetBatchFromName(names []string) (results []*CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}


func (obj *_CmsRoleLangMgr) FetchByPrimaryKey(idCmsRole uint32, idLang uint32, idShop uint32) (result CmsRoleLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ? AND id_lang = ? AND id_shop = ?", idCmsRole, idLang, idShop).Find(&result).Error

	return
}
