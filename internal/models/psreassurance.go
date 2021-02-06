package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _PsreassuranceMgr struct {
	*_BaseMgr
}

// PsreassuranceMgr open func
func PsreassuranceMgr(db *gorm.DB) *_PsreassuranceMgr {
	if db == nil {
		panic(fmt.Errorf("PsreassuranceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PsreassuranceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_psreassurance"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PsreassuranceMgr) GetTableName() string {
	return "ps_psreassurance"
}

// Get 获取
func (obj *_PsreassuranceMgr) Get() (result Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PsreassuranceMgr) Gets() (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDPsreassurance id_psreassurance获取
func (obj *_PsreassuranceMgr) WithIDPsreassurance(idPsreassurance uint32) Option {
	return optionFunc(func(o *options) { o.query["id_psreassurance"] = idPsreassurance })
}

// WithIcon icon获取
func (obj *_PsreassuranceMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithCustomIcon custom_icon获取
func (obj *_PsreassuranceMgr) WithCustomIcon(customIcon string) Option {
	return optionFunc(func(o *options) { o.query["custom_icon"] = customIcon })
}

// WithStatus status获取
func (obj *_PsreassuranceMgr) WithStatus(status uint32) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPosition position获取
func (obj *_PsreassuranceMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithIDShop id_shop获取
func (obj *_PsreassuranceMgr) WithIDShop(idShop uint32) Option {
	return optionFunc(func(o *options) { o.query["id_shop"] = idShop })
}

// WithTypeLink type_link获取
func (obj *_PsreassuranceMgr) WithTypeLink(typeLink uint32) Option {
	return optionFunc(func(o *options) { o.query["type_link"] = typeLink })
}

// WithIDCms id_cms获取
func (obj *_PsreassuranceMgr) WithIDCms(idCms uint32) Option {
	return optionFunc(func(o *options) { o.query["id_cms"] = idCms })
}

// WithDateAdd date_add获取
func (obj *_PsreassuranceMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateUpd date_upd获取
func (obj *_PsreassuranceMgr) WithDateUpd(dateUpd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_upd"] = dateUpd })
}

// GetByOption 功能选项模式获取
func (obj *_PsreassuranceMgr) GetByOption(opts ...Option) (result Psreassurance, err error) {
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
func (obj *_PsreassuranceMgr) GetByOptions(opts ...Option) (results []*Psreassurance, err error) {
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

// GetFromIDPsreassurance 通过id_psreassurance获取内容
func (obj *_PsreassuranceMgr) GetFromIDPsreassurance(idPsreassurance uint32) (result Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance = ?", idPsreassurance).Find(&result).Error

	return
}

// GetBatchFromIDPsreassurance 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromIDPsreassurance(idPsreassurances []uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance IN (?)", idPsreassurances).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容
func (obj *_PsreassuranceMgr) GetFromIcon(icon string) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromIcon(icons []string) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromCustomIcon 通过custom_icon获取内容
func (obj *_PsreassuranceMgr) GetFromCustomIcon(customIcon string) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom_icon = ?", customIcon).Find(&results).Error

	return
}

// GetBatchFromCustomIcon 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromCustomIcon(customIcons []string) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("custom_icon IN (?)", customIcons).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_PsreassuranceMgr) GetFromStatus(status uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromStatus(statuss []uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_PsreassuranceMgr) GetFromPosition(position uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromPosition(positions []uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

// GetFromIDShop 通过id_shop获取内容
func (obj *_PsreassuranceMgr) GetFromIDShop(idShop uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop = ?", idShop).Find(&results).Error

	return
}

// GetBatchFromIDShop 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromIDShop(idShops []uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_shop IN (?)", idShops).Find(&results).Error

	return
}

// GetFromTypeLink 通过type_link获取内容
func (obj *_PsreassuranceMgr) GetFromTypeLink(typeLink uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_link = ?", typeLink).Find(&results).Error

	return
}

// GetBatchFromTypeLink 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromTypeLink(typeLinks []uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type_link IN (?)", typeLinks).Find(&results).Error

	return
}

// GetFromIDCms 通过id_cms获取内容
func (obj *_PsreassuranceMgr) GetFromIDCms(idCms uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms = ?", idCms).Find(&results).Error

	return
}

// GetBatchFromIDCms 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromIDCms(idCmss []uint32) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_cms IN (?)", idCmss).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_PsreassuranceMgr) GetFromDateAdd(dateAdd time.Time) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateUpd 通过date_upd获取内容
func (obj *_PsreassuranceMgr) GetFromDateUpd(dateUpd time.Time) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd = ?", dateUpd).Find(&results).Error

	return
}

// GetBatchFromDateUpd 批量唯一主键查找
func (obj *_PsreassuranceMgr) GetBatchFromDateUpd(dateUpds []time.Time) (results []*Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_upd IN (?)", dateUpds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PsreassuranceMgr) FetchByPrimaryKey(idPsreassurance uint32) (result Psreassurance, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_psreassurance = ?", idPsreassurance).Find(&result).Error

	return
}
