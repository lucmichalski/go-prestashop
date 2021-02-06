package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImageTypeMgr struct {
	*_BaseMgr
}

// ImageTypeMgr open func
func ImageTypeMgr(db *gorm.DB) *_ImageTypeMgr {
	if db == nil {
		panic(fmt.Errorf("ImageTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImageTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eg_image_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ImageTypeMgr) GetTableName() string {
	return "eg_image_type"
}

// Get 获取
func (obj *_ImageTypeMgr) Get() (result ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ImageTypeMgr) Gets() (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDImageType id_image_type获取
func (obj *_ImageTypeMgr) WithIDImageType(idImageType uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image_type"] = idImageType })
}

// WithName name获取
func (obj *_ImageTypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithWidth width获取
func (obj *_ImageTypeMgr) WithWidth(width uint32) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

// WithHeight height获取
func (obj *_ImageTypeMgr) WithHeight(height uint32) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithProducts products获取
func (obj *_ImageTypeMgr) WithProducts(products bool) Option {
	return optionFunc(func(o *options) { o.query["products"] = products })
}

// WithCategories categories获取
func (obj *_ImageTypeMgr) WithCategories(categories bool) Option {
	return optionFunc(func(o *options) { o.query["categories"] = categories })
}

// WithManufacturers manufacturers获取
func (obj *_ImageTypeMgr) WithManufacturers(manufacturers bool) Option {
	return optionFunc(func(o *options) { o.query["manufacturers"] = manufacturers })
}

// WithSuppliers suppliers获取
func (obj *_ImageTypeMgr) WithSuppliers(suppliers bool) Option {
	return optionFunc(func(o *options) { o.query["suppliers"] = suppliers })
}

// WithStores stores获取
func (obj *_ImageTypeMgr) WithStores(stores bool) Option {
	return optionFunc(func(o *options) { o.query["stores"] = stores })
}

// GetByOption 功能选项模式获取
func (obj *_ImageTypeMgr) GetByOption(opts ...Option) (result ImageType, err error) {
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
func (obj *_ImageTypeMgr) GetByOptions(opts ...Option) (results []*ImageType, err error) {
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

// GetFromIDImageType 通过id_image_type获取内容
func (obj *_ImageTypeMgr) GetFromIDImageType(idImageType uint32) (result ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image_type = ?", idImageType).Find(&result).Error

	return
}

// GetBatchFromIDImageType 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromIDImageType(idImageTypes []uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image_type IN (?)", idImageTypes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_ImageTypeMgr) GetFromName(name string) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromName(names []string) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromWidth 通过width获取内容
func (obj *_ImageTypeMgr) GetFromWidth(width uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("width = ?", width).Find(&results).Error

	return
}

// GetBatchFromWidth 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromWidth(widths []uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("width IN (?)", widths).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容
func (obj *_ImageTypeMgr) GetFromHeight(height uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromHeight(heights []uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error

	return
}

// GetFromProducts 通过products获取内容
func (obj *_ImageTypeMgr) GetFromProducts(products bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("products = ?", products).Find(&results).Error

	return
}

// GetBatchFromProducts 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromProducts(productss []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("products IN (?)", productss).Find(&results).Error

	return
}

// GetFromCategories 通过categories获取内容
func (obj *_ImageTypeMgr) GetFromCategories(categories bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("categories = ?", categories).Find(&results).Error

	return
}

// GetBatchFromCategories 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromCategories(categoriess []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("categories IN (?)", categoriess).Find(&results).Error

	return
}

// GetFromManufacturers 通过manufacturers获取内容
func (obj *_ImageTypeMgr) GetFromManufacturers(manufacturers bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("manufacturers = ?", manufacturers).Find(&results).Error

	return
}

// GetBatchFromManufacturers 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromManufacturers(manufacturerss []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("manufacturers IN (?)", manufacturerss).Find(&results).Error

	return
}

// GetFromSuppliers 通过suppliers获取内容
func (obj *_ImageTypeMgr) GetFromSuppliers(suppliers bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("suppliers = ?", suppliers).Find(&results).Error

	return
}

// GetBatchFromSuppliers 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromSuppliers(supplierss []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("suppliers IN (?)", supplierss).Find(&results).Error

	return
}

// GetFromStores 通过stores获取内容
func (obj *_ImageTypeMgr) GetFromStores(stores bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stores = ?", stores).Find(&results).Error

	return
}

// GetBatchFromStores 批量唯一主键查找
func (obj *_ImageTypeMgr) GetBatchFromStores(storess []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stores IN (?)", storess).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ImageTypeMgr) FetchByPrimaryKey(idImageType uint32) (result ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image_type = ?", idImageType).Find(&result).Error

	return
}

// FetchIndexByImageTypeName  获取多个内容
func (obj *_ImageTypeMgr) FetchIndexByImageTypeName(name string) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
