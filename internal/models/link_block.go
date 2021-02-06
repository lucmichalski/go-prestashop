package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _LinkBlockMgr struct {
	*_BaseMgr
}

func LinkBlockMgr(db *gorm.DB) *_LinkBlockMgr {
	if db == nil {
		panic(fmt.Errorf("LinkBlockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LinkBlockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_link_block"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_LinkBlockMgr) GetTableName() string {
	return "ps_link_block"
}

func (obj *_LinkBlockMgr) Get() (result LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_LinkBlockMgr) Gets() (results []*LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_LinkBlockMgr) WithIDLinkBlock(idLinkBlock uint32) Option {
	return optionFunc(func(o *options) { o.query["id_link_block"] = idLinkBlock })
}

func (obj *_LinkBlockMgr) WithIDHook(idHook uint32) Option {
	return optionFunc(func(o *options) { o.query["id_hook"] = idHook })
}

func (obj *_LinkBlockMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_LinkBlockMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

func (obj *_LinkBlockMgr) GetByOption(opts ...Option) (result LinkBlock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_LinkBlockMgr) GetByOptions(opts ...Option) (results []*LinkBlock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_LinkBlockMgr) GetFromIDLinkBlock(idLinkBlock uint32) (result LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ?", idLinkBlock).Find(&result).Error

	return
}

func (obj *_LinkBlockMgr) GetBatchFromIDLinkBlock(idLinkBlocks []uint32) (results []*LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block IN (?)", idLinkBlocks).Find(&results).Error

	return
}

func (obj *_LinkBlockMgr) GetFromIDHook(idHook uint32) (results []*LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook = ?", idHook).Find(&results).Error

	return
}

func (obj *_LinkBlockMgr) GetBatchFromIDHook(idHooks []uint32) (results []*LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_hook IN (?)", idHooks).Find(&results).Error

	return
}

func (obj *_LinkBlockMgr) GetFromPosition(position uint32) (results []*LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_LinkBlockMgr) GetBatchFromPosition(positions []uint32) (results []*LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_LinkBlockMgr) GetFromContent(content string) (results []*LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

func (obj *_LinkBlockMgr) GetBatchFromContent(contents []string) (results []*LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}


func (obj *_LinkBlockMgr) FetchByPrimaryKey(idLinkBlock uint32) (result LinkBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_link_block = ?", idLinkBlock).Find(&result).Error

	return
}
