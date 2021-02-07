package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _RiskLangMgr struct {
	*_BaseMgr
}

func RiskLangMgr(db *gorm.DB) *_RiskLangMgr {
	if db == nil {
		panic(fmt.Errorf("RiskLangMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RiskLangMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_risk_lang"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_RiskLangMgr) GetTableName() string {
	return "ps_risk_lang"
}

func (obj *_RiskLangMgr) Get() (result RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_RiskLangMgr) Gets() (results []*RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

func (obj *_RiskLangMgr) WithIDRisk(idRisk uint32) Option {
	return optionFunc(func(o *options) { o.query["id_risk"] = idRisk })
}

func (obj *_RiskLangMgr) WithIDLang(idLang uint32) Option {
	return optionFunc(func(o *options) { o.query["id_lang"] = idLang })
}

func (obj *_RiskLangMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_RiskLangMgr) GetByOption(opts ...Option) (result RiskLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_RiskLangMgr) GetByOptions(opts ...Option) (results []*RiskLang, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

func (obj *_RiskLangMgr) GetFromIDRisk(idRisk uint32) (results []*RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&results).Error

	return
}

func (obj *_RiskLangMgr) GetBatchFromIDRisk(idRisks []uint32) (results []*RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk IN (?)", idRisks).Find(&results).Error

	return
}

func (obj *_RiskLangMgr) GetFromIDLang(idLang uint32) (results []*RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang = ?", idLang).Find(&results).Error

	return
}

func (obj *_RiskLangMgr) GetBatchFromIDLang(idLangs []uint32) (results []*RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_lang IN (?)", idLangs).Find(&results).Error

	return
}

func (obj *_RiskLangMgr) GetFromName(name string) (results []*RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_RiskLangMgr) GetBatchFromName(names []string) (results []*RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_RiskLangMgr) FetchByPrimaryKey(idRisk uint32, idLang uint32) (result RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ? AND id_lang = ?", idRisk, idLang).Find(&result).Error

	return
}

func (obj *_RiskLangMgr) FetchIndexByIDRisk(idRisk uint32) (results []*RiskLang, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_risk = ?", idRisk).Find(&results).Error

	return
}
