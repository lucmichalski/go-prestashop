package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _CarrierMgr struct {
	*_BaseMgr
}

// CarrierMgr open func
func CarrierMgr(db *gorm.DB) *_CarrierMgr {
	if db == nil {
		panic(fmt.Errorf("CarrierMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CarrierMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_carrier"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CarrierMgr) GetTableName() string {
	return "ps_carrier"
}

// Get 获取
func (obj *_CarrierMgr) Get() (result Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CarrierMgr) Gets() (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDCarrier id_carrier获取
func (obj *_CarrierMgr) WithIDCarrier(idCarrier uint32) Option {
	return optionFunc(func(o *options) { o.query["id_carrier"] = idCarrier })
}

// WithIDReference id_reference获取
func (obj *_CarrierMgr) WithIDReference(idReference uint32) Option {
	return optionFunc(func(o *options) { o.query["id_reference"] = idReference })
}

// WithIDTaxRulesGroup id_tax_rules_group获取
func (obj *_CarrierMgr) WithIDTaxRulesGroup(idTaxRulesGroup uint32) Option {
	return optionFunc(func(o *options) { o.query["id_tax_rules_group"] = idTaxRulesGroup })
}

// WithName name获取
func (obj *_CarrierMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithURL url获取
func (obj *_CarrierMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithActive active获取
func (obj *_CarrierMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithDeleted deleted获取
func (obj *_CarrierMgr) WithDeleted(deleted bool) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithShippingHandling shipping_handling获取
func (obj *_CarrierMgr) WithShippingHandling(shippingHandling bool) Option {
	return optionFunc(func(o *options) { o.query["shipping_handling"] = shippingHandling })
}

// WithRangeBehavior range_behavior获取
func (obj *_CarrierMgr) WithRangeBehavior(rangeBehavior bool) Option {
	return optionFunc(func(o *options) { o.query["range_behavior"] = rangeBehavior })
}

// WithIsModule is_module获取
func (obj *_CarrierMgr) WithIsModule(isModule bool) Option {
	return optionFunc(func(o *options) { o.query["is_module"] = isModule })
}

// WithIsFree is_free获取
func (obj *_CarrierMgr) WithIsFree(isFree bool) Option {
	return optionFunc(func(o *options) { o.query["is_free"] = isFree })
}

// WithShippingExternal shipping_external获取
func (obj *_CarrierMgr) WithShippingExternal(shippingExternal bool) Option {
	return optionFunc(func(o *options) { o.query["shipping_external"] = shippingExternal })
}

// WithNeedRange need_range获取
func (obj *_CarrierMgr) WithNeedRange(needRange bool) Option {
	return optionFunc(func(o *options) { o.query["need_range"] = needRange })
}

// WithExternalModuleName external_module_name获取
func (obj *_CarrierMgr) WithExternalModuleName(externalModuleName string) Option {
	return optionFunc(func(o *options) { o.query["external_module_name"] = externalModuleName })
}

// WithShippingMethod shipping_method获取
func (obj *_CarrierMgr) WithShippingMethod(shippingMethod int) Option {
	return optionFunc(func(o *options) { o.query["shipping_method"] = shippingMethod })
}

// WithPosition position获取
func (obj *_CarrierMgr) WithPosition(position uint32) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithMaxWidth max_width获取
func (obj *_CarrierMgr) WithMaxWidth(maxWidth int) Option {
	return optionFunc(func(o *options) { o.query["max_width"] = maxWidth })
}

// WithMaxHeight max_height获取
func (obj *_CarrierMgr) WithMaxHeight(maxHeight int) Option {
	return optionFunc(func(o *options) { o.query["max_height"] = maxHeight })
}

// WithMaxDepth max_depth获取
func (obj *_CarrierMgr) WithMaxDepth(maxDepth int) Option {
	return optionFunc(func(o *options) { o.query["max_depth"] = maxDepth })
}

// WithMaxWeight max_weight获取
func (obj *_CarrierMgr) WithMaxWeight(maxWeight float64) Option {
	return optionFunc(func(o *options) { o.query["max_weight"] = maxWeight })
}

// WithGrade grade获取
func (obj *_CarrierMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

// GetByOption 功能选项模式获取
func (obj *_CarrierMgr) GetByOption(opts ...Option) (result Carrier, err error) {
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
func (obj *_CarrierMgr) GetByOptions(opts ...Option) (results []*Carrier, err error) {
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

// GetFromIDCarrier 通过id_carrier获取内容
func (obj *_CarrierMgr) GetFromIDCarrier(idCarrier uint32) (result Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&result).Error

	return
}

// GetBatchFromIDCarrier 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromIDCarrier(idCarriers []uint32) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier IN (?)", idCarriers).Find(&results).Error

	return
}

// GetFromIDReference 通过id_reference获取内容
func (obj *_CarrierMgr) GetFromIDReference(idReference uint32) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference = ?", idReference).Find(&results).Error

	return
}

// GetBatchFromIDReference 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromIDReference(idReferences []uint32) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference IN (?)", idReferences).Find(&results).Error

	return
}

// GetFromIDTaxRulesGroup 通过id_tax_rules_group获取内容
func (obj *_CarrierMgr) GetFromIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}

// GetBatchFromIDTaxRulesGroup 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromIDTaxRulesGroup(idTaxRulesGroups []uint32) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group IN (?)", idTaxRulesGroups).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_CarrierMgr) GetFromName(name string) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromName(names []string) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容
func (obj *_CarrierMgr) GetFromURL(url string) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromURL(urls []string) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_CarrierMgr) GetFromActive(active bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromActive(actives []bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容
func (obj *_CarrierMgr) GetFromDeleted(deleted bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromDeleted(deleteds []bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("deleted IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromShippingHandling 通过shipping_handling获取内容
func (obj *_CarrierMgr) GetFromShippingHandling(shippingHandling bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_handling = ?", shippingHandling).Find(&results).Error

	return
}

// GetBatchFromShippingHandling 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromShippingHandling(shippingHandlings []bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_handling IN (?)", shippingHandlings).Find(&results).Error

	return
}

// GetFromRangeBehavior 通过range_behavior获取内容
func (obj *_CarrierMgr) GetFromRangeBehavior(rangeBehavior bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("range_behavior = ?", rangeBehavior).Find(&results).Error

	return
}

// GetBatchFromRangeBehavior 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromRangeBehavior(rangeBehaviors []bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("range_behavior IN (?)", rangeBehaviors).Find(&results).Error

	return
}

// GetFromIsModule 通过is_module获取内容
func (obj *_CarrierMgr) GetFromIsModule(isModule bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module = ?", isModule).Find(&results).Error

	return
}

// GetBatchFromIsModule 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromIsModule(isModules []bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_module IN (?)", isModules).Find(&results).Error

	return
}

// GetFromIsFree 通过is_free获取内容
func (obj *_CarrierMgr) GetFromIsFree(isFree bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_free = ?", isFree).Find(&results).Error

	return
}

// GetBatchFromIsFree 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromIsFree(isFrees []bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_free IN (?)", isFrees).Find(&results).Error

	return
}

// GetFromShippingExternal 通过shipping_external获取内容
func (obj *_CarrierMgr) GetFromShippingExternal(shippingExternal bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_external = ?", shippingExternal).Find(&results).Error

	return
}

// GetBatchFromShippingExternal 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromShippingExternal(shippingExternals []bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_external IN (?)", shippingExternals).Find(&results).Error

	return
}

// GetFromNeedRange 通过need_range获取内容
func (obj *_CarrierMgr) GetFromNeedRange(needRange bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("need_range = ?", needRange).Find(&results).Error

	return
}

// GetBatchFromNeedRange 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromNeedRange(needRanges []bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("need_range IN (?)", needRanges).Find(&results).Error

	return
}

// GetFromExternalModuleName 通过external_module_name获取内容
func (obj *_CarrierMgr) GetFromExternalModuleName(externalModuleName string) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("external_module_name = ?", externalModuleName).Find(&results).Error

	return
}

// GetBatchFromExternalModuleName 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromExternalModuleName(externalModuleNames []string) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("external_module_name IN (?)", externalModuleNames).Find(&results).Error

	return
}

// GetFromShippingMethod 通过shipping_method获取内容
func (obj *_CarrierMgr) GetFromShippingMethod(shippingMethod int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_method = ?", shippingMethod).Find(&results).Error

	return
}

// GetBatchFromShippingMethod 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromShippingMethod(shippingMethods []int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("shipping_method IN (?)", shippingMethods).Find(&results).Error

	return
}

// GetFromPosition 通过position获取内容
func (obj *_CarrierMgr) GetFromPosition(position uint32) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position = ?", position).Find(&results).Error

	return
}

// GetBatchFromPosition 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromPosition(positions []uint32) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("position IN (?)", positions).Find(&results).Error

	return
}

// GetFromMaxWidth 通过max_width获取内容
func (obj *_CarrierMgr) GetFromMaxWidth(maxWidth int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_width = ?", maxWidth).Find(&results).Error

	return
}

// GetBatchFromMaxWidth 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromMaxWidth(maxWidths []int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_width IN (?)", maxWidths).Find(&results).Error

	return
}

// GetFromMaxHeight 通过max_height获取内容
func (obj *_CarrierMgr) GetFromMaxHeight(maxHeight int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_height = ?", maxHeight).Find(&results).Error

	return
}

// GetBatchFromMaxHeight 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromMaxHeight(maxHeights []int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_height IN (?)", maxHeights).Find(&results).Error

	return
}

// GetFromMaxDepth 通过max_depth获取内容
func (obj *_CarrierMgr) GetFromMaxDepth(maxDepth int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_depth = ?", maxDepth).Find(&results).Error

	return
}

// GetBatchFromMaxDepth 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromMaxDepth(maxDepths []int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_depth IN (?)", maxDepths).Find(&results).Error

	return
}

// GetFromMaxWeight 通过max_weight获取内容
func (obj *_CarrierMgr) GetFromMaxWeight(maxWeight float64) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_weight = ?", maxWeight).Find(&results).Error

	return
}

// GetBatchFromMaxWeight 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromMaxWeight(maxWeights []float64) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("max_weight IN (?)", maxWeights).Find(&results).Error

	return
}

// GetFromGrade 通过grade获取内容
func (obj *_CarrierMgr) GetFromGrade(grade int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade = ?", grade).Find(&results).Error

	return
}

// GetBatchFromGrade 批量唯一主键查找
func (obj *_CarrierMgr) GetBatchFromGrade(grades []int) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("grade IN (?)", grades).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_CarrierMgr) FetchByPrimaryKey(idCarrier uint32) (result Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_carrier = ?", idCarrier).Find(&result).Error

	return
}

// FetchIndexByReference  获取多个内容
func (obj *_CarrierMgr) FetchIndexByReference(idReference uint32, active bool, deleted bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_reference = ? AND active = ? AND deleted = ?", idReference, active, deleted).Find(&results).Error

	return
}

// FetchIndexByIDTaxRulesGroup  获取多个内容
func (obj *_CarrierMgr) FetchIndexByIDTaxRulesGroup(idTaxRulesGroup uint32) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_tax_rules_group = ?", idTaxRulesGroup).Find(&results).Error

	return
}

// FetchIndexByDeleted  获取多个内容
func (obj *_CarrierMgr) FetchIndexByDeleted(active bool, deleted bool) (results []*Carrier, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ? AND deleted = ?", active, deleted).Find(&results).Error

	return
}
