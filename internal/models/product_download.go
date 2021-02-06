package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ProductDownloadMgr struct {
	*_BaseMgr
}

// ProductDownloadMgr open func
func ProductDownloadMgr(db *gorm.DB) *_ProductDownloadMgr {
	if db == nil {
		panic(fmt.Errorf("ProductDownloadMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductDownloadMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_download"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductDownloadMgr) GetTableName() string {
	return "ps_product_download"
}

// Get 获取
func (obj *_ProductDownloadMgr) Get() (result ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductDownloadMgr) Gets() (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithIDProductDownload id_product_download获取
func (obj *_ProductDownloadMgr) WithIDProductDownload(idProductDownload uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_download"] = idProductDownload })
}

// WithIDProduct id_product获取
func (obj *_ProductDownloadMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

// WithDisplayFilename display_filename获取
func (obj *_ProductDownloadMgr) WithDisplayFilename(displayFilename string) Option {
	return optionFunc(func(o *options) { o.query["display_filename"] = displayFilename })
}

// WithFilename filename获取
func (obj *_ProductDownloadMgr) WithFilename(filename string) Option {
	return optionFunc(func(o *options) { o.query["filename"] = filename })
}

// WithDateAdd date_add获取
func (obj *_ProductDownloadMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

// WithDateExpiration date_expiration获取
func (obj *_ProductDownloadMgr) WithDateExpiration(dateExpiration time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_expiration"] = dateExpiration })
}

// WithNbDaysAccessible nb_days_accessible获取
func (obj *_ProductDownloadMgr) WithNbDaysAccessible(nbDaysAccessible uint32) Option {
	return optionFunc(func(o *options) { o.query["nb_days_accessible"] = nbDaysAccessible })
}

// WithNbDownloadable nb_downloadable获取
func (obj *_ProductDownloadMgr) WithNbDownloadable(nbDownloadable uint32) Option {
	return optionFunc(func(o *options) { o.query["nb_downloadable"] = nbDownloadable })
}

// WithActive active获取
func (obj *_ProductDownloadMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

// WithIsShareable is_shareable获取
func (obj *_ProductDownloadMgr) WithIsShareable(isShareable bool) Option {
	return optionFunc(func(o *options) { o.query["is_shareable"] = isShareable })
}

// GetByOption 功能选项模式获取
func (obj *_ProductDownloadMgr) GetByOption(opts ...Option) (result ProductDownload, err error) {
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
func (obj *_ProductDownloadMgr) GetByOptions(opts ...Option) (results []*ProductDownload, err error) {
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

// GetFromIDProductDownload 通过id_product_download获取内容
func (obj *_ProductDownloadMgr) GetFromIDProductDownload(idProductDownload uint32) (result ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download = ?", idProductDownload).Find(&result).Error

	return
}

// GetBatchFromIDProductDownload 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromIDProductDownload(idProductDownloads []uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download IN (?)", idProductDownloads).Find(&results).Error

	return
}

// GetFromIDProduct 通过id_product获取内容
func (obj *_ProductDownloadMgr) GetFromIDProduct(idProduct uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

// GetBatchFromIDProduct 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

// GetFromDisplayFilename 通过display_filename获取内容
func (obj *_ProductDownloadMgr) GetFromDisplayFilename(displayFilename string) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display_filename = ?", displayFilename).Find(&results).Error

	return
}

// GetBatchFromDisplayFilename 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromDisplayFilename(displayFilenames []string) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display_filename IN (?)", displayFilenames).Find(&results).Error

	return
}

// GetFromFilename 通过filename获取内容
func (obj *_ProductDownloadMgr) GetFromFilename(filename string) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filename = ?", filename).Find(&results).Error

	return
}

// GetBatchFromFilename 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromFilename(filenames []string) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filename IN (?)", filenames).Find(&results).Error

	return
}

// GetFromDateAdd 通过date_add获取内容
func (obj *_ProductDownloadMgr) GetFromDateAdd(dateAdd time.Time) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

// GetBatchFromDateAdd 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

// GetFromDateExpiration 通过date_expiration获取内容
func (obj *_ProductDownloadMgr) GetFromDateExpiration(dateExpiration time.Time) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_expiration = ?", dateExpiration).Find(&results).Error

	return
}

// GetBatchFromDateExpiration 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromDateExpiration(dateExpirations []time.Time) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_expiration IN (?)", dateExpirations).Find(&results).Error

	return
}

// GetFromNbDaysAccessible 通过nb_days_accessible获取内容
func (obj *_ProductDownloadMgr) GetFromNbDaysAccessible(nbDaysAccessible uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_days_accessible = ?", nbDaysAccessible).Find(&results).Error

	return
}

// GetBatchFromNbDaysAccessible 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromNbDaysAccessible(nbDaysAccessibles []uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_days_accessible IN (?)", nbDaysAccessibles).Find(&results).Error

	return
}

// GetFromNbDownloadable 通过nb_downloadable获取内容
func (obj *_ProductDownloadMgr) GetFromNbDownloadable(nbDownloadable uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_downloadable = ?", nbDownloadable).Find(&results).Error

	return
}

// GetBatchFromNbDownloadable 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromNbDownloadable(nbDownloadables []uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_downloadable IN (?)", nbDownloadables).Find(&results).Error

	return
}

// GetFromActive 通过active获取内容
func (obj *_ProductDownloadMgr) GetFromActive(active bool) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

// GetBatchFromActive 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromActive(actives []bool) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

// GetFromIsShareable 通过is_shareable获取内容
func (obj *_ProductDownloadMgr) GetFromIsShareable(isShareable bool) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_shareable = ?", isShareable).Find(&results).Error

	return
}

// GetBatchFromIsShareable 批量唯一主键查找
func (obj *_ProductDownloadMgr) GetBatchFromIsShareable(isShareables []bool) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_shareable IN (?)", isShareables).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ProductDownloadMgr) FetchByPrimaryKey(idProductDownload uint32) (result ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download = ?", idProductDownload).Find(&result).Error

	return
}
