package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CategoryMgr struct {
	*_BaseMgr
}

func CategoryMgr(db *gorm.DB) *_CategoryMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_CategoryMgr) GetTableName() string {
	return "ps_category"
}

func (obj *_CategoryMgr) Get() (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_CategoryMgr) Gets() (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_CategoryMgr) WithIDCategory(idCategory uint32) Option {
	return optionFunc(func(o *options) { o.query["id_category"] = idCategory })
}

func (obj *_CategoryMgr) WithIDParent(idParent uint32) Option {
	return optionFunc(func(o *options) { o.query["id_parent"] = idParent })
}

func (obj *_CategoryMgr) WithIDShopDefault(idShopDefault uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop_default"] = idShopDefault })
}

func (obj *_CategoryMgr) WithLevelDepth(levelDepth uint8) Option {
	return optionFunc(func(o *options) { o.query["level_depth"] = levelDepth })
}

func (obj *_CategoryMgr) WithNleft(nleft uint32) Option {
	return optionFunc(func(o *options) { o.query["nleft"] = nleft })
}

func (obj *_CategoryMgr) WithNright(nright uint32) Option {
	return optionFunc(func(o *options) { o.query["nright"] = nright })
}

func (obj *_CategoryMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_CategoryMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_CategoryMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

func (obj *_CategoryMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

func (obj *_CategoryMgr) WithIsRootCategory(isRootCategory bool) Option {
	return optionFunc(func(o *options) { o.query["is_root_category"] = isRootCategory })
}

func (obj *_CategoryMgr) GetByOption(opts ...Option) (result Category, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_CategoryMgr) GetByOptions(opts ...Option) (results []*Category, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_CategoryMgr) GetFromIDCategory(idCategory uint32) (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&result).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromIDCategory(idCategorys []uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category IN (?)", idCategorys).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromIDParent(idParent uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromIDParent(idParents []uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent IN (?)", idParents).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromIDShopDefault(idShopDefault uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_default = ?", idShopDefault).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromIDShopDefault(idShopDefaults []uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop_default IN (?)", idShopDefaults).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromLevelDepth(levelDepth uint8) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth = ?", levelDepth).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromLevelDepth(levelDepths []uint8) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth IN (?)", levelDepths).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromNleft(nleft uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nleft = ?", nleft).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromNleft(nlefts []uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nleft IN (?)", nlefts).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromNright(nright uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nright = ?", nright).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromNright(nrights []uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nright IN (?)", nrights).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromActive(active bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromActive(actives []bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromDateAdd(dateAdd time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromDateUpd(dateUpd time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromPosition(position uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromPosition(positions []uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetFromIsRootCategory(isRootCategory bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_root_category = ?", isRootCategory).Find(&results).Error

	return
}

func (obj *_CategoryMgr) GetBatchFromIsRootCategory(isRootCategorys []bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_root_category IN (?)", isRootCategorys).Find(&results).Error

	return
}


func (obj *_CategoryMgr) FetchByPrimaryKey(idCategory uint32) (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_category = ?", idCategory).Find(&result).Error

	return
}

func (obj *_CategoryMgr) FetchIndexByCategoryParent(idParent uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_parent = ?", idParent).Find(&results).Error

	return
}

func (obj *_CategoryMgr) FetchIndexByLevelDepth(levelDepth uint8) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("level_depth = ?", levelDepth).Find(&results).Error

	return
}

func (obj *_CategoryMgr) FetchIndexByNleftrightactive(nleft uint32, nright uint32, active bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nleft = ? AND nright = ? AND active = ?", nleft, nright, active).Find(&results).Error

	return
}

func (obj *_CategoryMgr) FetchIndexByActivenleft(nleft uint32, active bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nleft = ? AND active = ?", nleft, active).Find(&results).Error

	return
}

func (obj *_CategoryMgr) FetchIndexByNright(nright uint32) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nright = ?", nright).Find(&results).Error

	return
}

func (obj *_CategoryMgr) FetchIndexByActivenright(nright uint32, active bool) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nright = ? AND active = ?", nright, active).Find(&results).Error

	return
}
