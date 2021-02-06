package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsRoleMgr struct {
	*_BaseMgr
}

func CmsRoleMgr(db *gorm.DB) *_CmsRoleMgr {
	if db == nil {
		panic(fmt.Errorf("CmsRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_cms_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CmsRoleMgr) GetTableName() string {
	return "ps_cms_role"
}

func (obj *_CmsRoleMgr) Get() (result CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CmsRoleMgr) Gets() (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CmsRoleMgr) WithIDCmsRole(idCmsRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_role"] = idCmsRole })
}

func (obj *_CmsRoleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_CmsRoleMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

func (obj *_CmsRoleMgr) GetByOption(opts ...Option) (result CmsRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CmsRoleMgr) GetByOptions(opts ...Option) (results []*CmsRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CmsRoleMgr) GetFromIDCmsRole(idCmsRole uint32) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ?", idCmsRole).Find(&results).Error

	return
}

func (obj *_CmsRoleMgr) GetBatchFromIDCmsRole(idCmsRoles []uint32) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role IN (?)", idCmsRoles).Find(&results).Error

	return
}

func (obj *_CmsRoleMgr) GetFromName(name string) (result CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}

func (obj *_CmsRoleMgr) GetBatchFromName(names []string) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_CmsRoleMgr) GetFromIDCms(idCms uint32) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error

	return
}

func (obj *_CmsRoleMgr) GetBatchFromIDCms(idCmss []uint32) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error

	return
}


func (obj *_CmsRoleMgr) FetchByPrimaryKey(idCmsRole uint32, idCms uint32) (result CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ? AND id_cms = ?", idCmsRole, idCms).Find(&result).Error

	return
}

func (obj *_CmsRoleMgr) FetchUniqueByName(name string) (result CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}
