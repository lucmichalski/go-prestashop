package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LinkBlockLangMgr struct {
	*_BaseMgr
}

func LinkBlockLangMgr(db *gorm.DB) *_LinkBlockLangMgr {
	if db == nil {
		panic(fmt.Errorf("LinkBlockLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LinkBlockLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_link_block_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LinkBlockLangMgr) GetTableName() string {
	return "ps_link_block_lang"
}

func (obj *_LinkBlockLangMgr) Get() (result LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LinkBlockLangMgr) Gets() (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) WithIDLinkBlock(idLinkBlock uint32) Option {
	return optionFunc(func(o *options) { o.query["id_link_block"] = idLinkBlock })
}

func (obj *_LinkBlockLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_LinkBlockLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_LinkBlockLangMgr) WithCustomContent(customContent string) Option {
	return optionFunc(func(o *options) { o.query["custom_content"] = customContent })
}

func (obj *_LinkBlockLangMgr) GetByOption(opts ...Option) (result LinkBlockLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LinkBlockLangMgr) GetByOptions(opts ...Option) (results []*LinkBlockLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) GetFromIDLinkBlock(idLinkBlock uint32) (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ?", idLinkBlock).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) GetBatchFromIDLinkBlock(idLinkBlocks []uint32) (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block IN (?)", idLinkBlocks).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) GetFromIDLang(idLang uint32) (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) GetFromName(name string) (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) GetBatchFromName(names []string) (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) GetFromCustomContent(customContent string) (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom_content = ?", customContent).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) GetBatchFromCustomContent(customContents []string) (results []*LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom_content IN (?)", customContents).Find(&results).Error

	return
}

func (obj *_LinkBlockLangMgr) FetchByPrimaryKey(idLinkBlock uint32, idLang uint32) (result LinkBlockLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ? AND id_lang = ?", idLinkBlock, idLang).Find(&result).Error

	return
}
