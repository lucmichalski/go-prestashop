package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ImageTypeMgr struct {
	*_BaseMgr
}

func ImageTypeMgr(db *gorm.DB) *_ImageTypeMgr {
	if db == nil {
		panic(fmt.Errorf("ImageTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ImageTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_image_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ImageTypeMgr) GetTableName() string {
	return "ps_image_type"
}

func (obj *_ImageTypeMgr) Get() (result ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ImageTypeMgr) Gets() (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ImageTypeMgr) WithIDImageType(idImageType uint32) Option {
	return optionFunc(func(o *options) { o.query["id_image_type"] = idImageType })
}

func (obj *_ImageTypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ImageTypeMgr) WithWidth(width uint32) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

func (obj *_ImageTypeMgr) WithHeight(height uint32) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

func (obj *_ImageTypeMgr) WithProducts(products bool) Option {
	return optionFunc(func(o *options) { o.query["products"] = products })
}

func (obj *_ImageTypeMgr) WithCategories(categories bool) Option {
	return optionFunc(func(o *options) { o.query["categories"] = categories })
}

func (obj *_ImageTypeMgr) WithManufacturers(manufacturers bool) Option {
	return optionFunc(func(o *options) { o.query["manufacturers"] = manufacturers })
}

func (obj *_ImageTypeMgr) WithSuppliers(suppliers bool) Option {
	return optionFunc(func(o *options) { o.query["suppliers"] = suppliers })
}

func (obj *_ImageTypeMgr) WithStores(stores bool) Option {
	return optionFunc(func(o *options) { o.query["stores"] = stores })
}

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


func (obj *_ImageTypeMgr) GetFromIDImageType(idImageType uint32) (result ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image_type = ?", idImageType).Find(&result).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromIDImageType(idImageTypes []uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image_type IN (?)", idImageTypes).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetFromName(name string) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromName(names []string) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetFromWidth(width uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("width = ?", width).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromWidth(widths []uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("width IN (?)", widths).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetFromHeight(height uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromHeight(heights []uint32) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetFromProducts(products bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("products = ?", products).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromProducts(productss []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("products IN (?)", productss).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetFromCategories(categories bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("categories = ?", categories).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromCategories(categoriess []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("categories IN (?)", categoriess).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetFromManufacturers(manufacturers bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("manufacturers = ?", manufacturers).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromManufacturers(manufacturerss []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("manufacturers IN (?)", manufacturerss).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetFromSuppliers(suppliers bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("suppliers = ?", suppliers).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromSuppliers(supplierss []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("suppliers IN (?)", supplierss).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetFromStores(stores bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stores = ?", stores).Find(&results).Error

	return
}

func (obj *_ImageTypeMgr) GetBatchFromStores(storess []bool) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("stores IN (?)", storess).Find(&results).Error

	return
}


func (obj *_ImageTypeMgr) FetchByPrimaryKey(idImageType uint32) (result ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_image_type = ?", idImageType).Find(&result).Error

	return
}

func (obj *_ImageTypeMgr) FetchIndexByImageTypeName(name string) (results []*ImageType, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}
