package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StateMgr struct {
	*_BaseMgr
}

func StateMgr(db *gorm.DB) *_StateMgr {
	if db == nil {
		panic(fmt.Errorf("StateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_state"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_StateMgr) GetTableName() string {
	return "ps_state"
}

func (obj *_StateMgr) Get() (result State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_StateMgr) Gets() (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_StateMgr) WithIDState(idState uint32) Option {
	return optionFunc(func(o *options) { o.query["id_state"] = idState })
}

func (obj *_StateMgr) WithIDCountry(idCountry uint32) Option {
	return optionFunc(func(o *options) { o.query["id_country"] = idCountry })
}

func (obj *_StateMgr) WithIDZone(idZone uint32) Option {
	return optionFunc(func(o *options) { o.query["id_zone"] = idZone })
}

func (obj *_StateMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_StateMgr) WithIsoCode(isoCode string) Option {
	return optionFunc(func(o *options) { o.query["iso_code"] = isoCode })
}

func (obj *_StateMgr) WithTaxBehavior(taxBehavior int16) Option {
	return optionFunc(func(o *options) { o.query["tax_behavior"] = taxBehavior })
}

func (obj *_StateMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_StateMgr) GetByOption(opts ...Option) (result State, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_StateMgr) GetByOptions(opts ...Option) (results []*State, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_StateMgr) GetFromIDState(idState uint32) (result State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&result).Error

	return
}

func (obj *_StateMgr) GetBatchFromIDState(idStates []uint32) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state IN (?)", idStates).Find(&results).Error

	return
}

func (obj *_StateMgr) GetFromIDCountry(idCountry uint32) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_StateMgr) GetBatchFromIDCountry(idCountrys []uint32) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country IN (?)", idCountrys).Find(&results).Error

	return
}

func (obj *_StateMgr) GetFromIDZone(idZone uint32) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}

func (obj *_StateMgr) GetBatchFromIDZone(idZones []uint32) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone IN (?)", idZones).Find(&results).Error

	return
}

func (obj *_StateMgr) GetFromName(name string) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_StateMgr) GetBatchFromName(names []string) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_StateMgr) GetFromIsoCode(isoCode string) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code = ?", isoCode).Find(&results).Error

	return
}

func (obj *_StateMgr) GetBatchFromIsoCode(isoCodes []string) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("iso_code IN (?)", isoCodes).Find(&results).Error

	return
}

func (obj *_StateMgr) GetFromTaxBehavior(taxBehavior int16) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_behavior = ?", taxBehavior).Find(&results).Error

	return
}

func (obj *_StateMgr) GetBatchFromTaxBehavior(taxBehaviors []int16) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tax_behavior IN (?)", taxBehaviors).Find(&results).Error

	return
}

func (obj *_StateMgr) GetFromActive(active bool) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_StateMgr) GetBatchFromActive(actives []bool) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}


func (obj *_StateMgr) FetchByPrimaryKey(idState uint32) (result State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_state = ?", idState).Find(&result).Error

	return
}

func (obj *_StateMgr) FetchIndexByIDCountry(idCountry uint32) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_country = ?", idCountry).Find(&results).Error

	return
}

func (obj *_StateMgr) FetchIndexByIDZone(idZone uint32) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_zone = ?", idZone).Find(&results).Error

	return
}

func (obj *_StateMgr) FetchIndexByName(name string) (results []*State, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
