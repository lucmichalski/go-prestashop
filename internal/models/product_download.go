package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ProductDownloadMgr struct {
	*_BaseMgr
}

func ProductDownloadMgr(db *gorm.DB) *_ProductDownloadMgr {
	if db == nil {
		panic(fmt.Errorf("ProductDownloadMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductDownloadMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_product_download"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ProductDownloadMgr) GetTableName() string {
	return "ps_product_download"
}

func (obj *_ProductDownloadMgr) Get() (result ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ProductDownloadMgr) Gets() (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ProductDownloadMgr) WithIDProductDownload(idProductDownload uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product_download"] = idProductDownload })
}

func (obj *_ProductDownloadMgr) WithIDProduct(idProduct uint32) Option {
	return optionFunc(func(o *options) { o.query["id_product"] = idProduct })
}

func (obj *_ProductDownloadMgr) WithDisplayFilename(displayFilename string) Option {
	return optionFunc(func(o *options) { o.query["display_filename"] = displayFilename })
}

func (obj *_ProductDownloadMgr) WithFilename(filename string) Option {
	return optionFunc(func(o *options) { o.query["filename"] = filename })
}

func (obj *_ProductDownloadMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ProductDownloadMgr) WithDateExpiration(dateExpiration time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_expiration"] = dateExpiration })
}

func (obj *_ProductDownloadMgr) WithNbDaysAccessible(nbDaysAccessible uint32) Option {
	return optionFunc(func(o *options) { o.query["nb_days_accessible"] = nbDaysAccessible })
}

func (obj *_ProductDownloadMgr) WithNbDownloadable(nbDownloadable uint32) Option {
	return optionFunc(func(o *options) { o.query["nb_downloadable"] = nbDownloadable })
}

func (obj *_ProductDownloadMgr) WithActive(active bool) Option {
	return optionFunc(func(o *options) { o.query["active"] = active })
}

func (obj *_ProductDownloadMgr) WithIsShareable(isShareable bool) Option {
	return optionFunc(func(o *options) { o.query["is_shareable"] = isShareable })
}

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


func (obj *_ProductDownloadMgr) GetFromIDProductDownload(idProductDownload uint32) (result ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download = ?", idProductDownload).Find(&result).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromIDProductDownload(idProductDownloads []uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download IN (?)", idProductDownloads).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromIDProduct(idProduct uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product = ?", idProduct).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromIDProduct(idProducts []uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product IN (?)", idProducts).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromDisplayFilename(displayFilename string) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display_filename = ?", displayFilename).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromDisplayFilename(displayFilenames []string) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("display_filename IN (?)", displayFilenames).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromFilename(filename string) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filename = ?", filename).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromFilename(filenames []string) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("filename IN (?)", filenames).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromDateAdd(dateAdd time.Time) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromDateExpiration(dateExpiration time.Time) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_expiration = ?", dateExpiration).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromDateExpiration(dateExpirations []time.Time) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_expiration IN (?)", dateExpirations).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromNbDaysAccessible(nbDaysAccessible uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_days_accessible = ?", nbDaysAccessible).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromNbDaysAccessible(nbDaysAccessibles []uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_days_accessible IN (?)", nbDaysAccessibles).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromNbDownloadable(nbDownloadable uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_downloadable = ?", nbDownloadable).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromNbDownloadable(nbDownloadables []uint32) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nb_downloadable IN (?)", nbDownloadables).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromActive(active bool) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active = ?", active).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromActive(actives []bool) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("active IN (?)", actives).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetFromIsShareable(isShareable bool) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_shareable = ?", isShareable).Find(&results).Error

	return
}

func (obj *_ProductDownloadMgr) GetBatchFromIsShareable(isShareables []bool) (results []*ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_shareable IN (?)", isShareables).Find(&results).Error

	return
}


func (obj *_ProductDownloadMgr) FetchByPrimaryKey(idProductDownload uint32) (result ProductDownload, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_product_download = ?", idProductDownload).Find(&result).Error

	return
}
