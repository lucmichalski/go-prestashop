package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AttributeGroupLangMgr struct {
	*_BaseMgr
}

// AttributeGroupLangMgr open func
func AttributeGroupLangMgr(db *gorm.DB) *_AttributeGroupLangMgr {
	if db == nil {
		panic(fmt.Errorf("AttributeGroupLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AttributeGroupLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_attribute_group_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AttributeGroupLangMgr) GetTableName() string {
	return "eg_attribute_group_lang"
}

// Get 获取
func (obj *_AttributeGroupLangMgr) Get() (result AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AttributeGroupLangMgr) Gets() (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDAttributeGroup id_attribute_group获取
func (obj *_AttributeGroupLangMgr) WithIDAttributeGroup(idAttributeGroup int) Option {
	return optionFunc(func(o *options) { o.query["id_attribute_group"] = idAttributeGroup })
}

// WithIDLang id_lang获取
func (obj *_AttributeGroupLangMgr) WithIDLang(idLang int) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

// WithName name获取
func (obj *_AttributeGroupLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPublicName public_name获取
func (obj *_AttributeGroupLangMgr) WithPublicName(publicName string) Option {
	return optionFunc(func(o *options) { o.query["public_name"] = publicName })
}

// GetByOption 功能选项模式获取
func (obj *_AttributeGroupLangMgr) GetByOption(opts ...Option) (result AttributeGroupLang, err error) {
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
func (obj *_AttributeGroupLangMgr) GetByOptions(opts ...Option) (results []*AttributeGroupLang, err error) {
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

// GetFromIDAttributeGroup 通过id_attribute_group获取内容
func (obj *_AttributeGroupLangMgr) GetFromIDAttributeGroup(idAttributeGroup int) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

// GetBatchFromIDAttributeGroup 批量唯一主键查找
func (obj *_AttributeGroupLangMgr) GetBatchFromIDAttributeGroup(idAttributeGroups []int) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group IN (?)", idAttributeGroups).Find(&results).Error

	return
}

// GetFromIDLang 通过id_lang获取内容
func (obj *_AttributeGroupLangMgr) GetFromIDLang(idLang int) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

// GetBatchFromIDLang 批量唯一主键查找
func (obj *_AttributeGroupLangMgr) GetBatchFromIDLang(idLangs []int) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AttributeGroupLangMgr) GetFromName(name string) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_AttributeGroupLangMgr) GetBatchFromName(names []string) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromPublicName 通过public_name获取内容
func (obj *_AttributeGroupLangMgr) GetFromPublicName(publicName string) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_name = ?", publicName).Find(&results).Error

	return
}

// GetBatchFromPublicName 批量唯一主键查找
func (obj *_AttributeGroupLangMgr) GetBatchFromPublicName(publicNames []string) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_name IN (?)", publicNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_AttributeGroupLangMgr) FetchByPrimaryKey(idAttributeGroup int, idLang int) (result AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ? AND id_lang = ?", idAttributeGroup, idLang).Find(&result).Error

	return
}

// FetchIndexByIDX2958662667A664FB  获取多个内容
func (obj *_AttributeGroupLangMgr) FetchIndexByIDX2958662667A664FB(idAttributeGroup int) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_attribute_group = ?", idAttributeGroup).Find(&results).Error

	return
}

// FetchIndexByIDX29586626BA299860  获取多个内容
func (obj *_AttributeGroupLangMgr) FetchIndexByIDX29586626BA299860(idLang int) (results []*AttributeGroupLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}
