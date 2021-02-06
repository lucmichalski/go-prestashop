package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CmsRoleMgr struct {
	*_BaseMgr
}

// CmsRoleMgr open func
func CmsRoleMgr(db *gorm.DB) *_CmsRoleMgr {
	if db == nil {
		panic(fmt.Errorf("CmsRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CmsRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_cms_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CmsRoleMgr) GetTableName() string {
	return "eg_cms_role"
}

// Get 获取
func (obj *_CmsRoleMgr) Get() (result CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CmsRoleMgr) Gets() (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCmsRole id_cms_role获取
func (obj *_CmsRoleMgr) WithIDCmsRole(idCmsRole uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms_role"] = idCmsRole })
}

// WithName name获取
func (obj *_CmsRoleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIDCms id_cms获取
func (obj *_CmsRoleMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

// GetByOption 功能选项模式获取
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

// GetByOptions 批量功能选项模式获取
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

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIDCmsRole 通过id_cms_role获取内容
func (obj *_CmsRoleMgr) GetFromIDCmsRole(idCmsRole uint32) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ?", idCmsRole).Find(&results).Error

	return
}

// GetBatchFromIDCmsRole 批量唯一主键查找
func (obj *_CmsRoleMgr) GetBatchFromIDCmsRole(idCmsRoles []uint32) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role IN (?)", idCmsRoles).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_CmsRoleMgr) GetFromName(name string) (result CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_CmsRoleMgr) GetBatchFromName(names []string) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromIDCms 通过id_cms获取内容
func (obj *_CmsRoleMgr) GetFromIDCms(idCms uint32) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error

	return
}

// GetBatchFromIDCms 批量唯一主键查找
func (obj *_CmsRoleMgr) GetBatchFromIDCms(idCmss []uint32) (results []*CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CmsRoleMgr) FetchByPrimaryKey(idCmsRole uint32, idCms uint32) (result CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms_role = ? AND id_cms = ?", idCmsRole, idCms).Find(&result).Error

	return
}

// FetchUniqueByName primay or index 获取唯一内容
func (obj *_CmsRoleMgr) FetchUniqueByName(name string) (result CmsRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&result).Error

	return
}
