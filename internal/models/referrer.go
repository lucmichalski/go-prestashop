package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ReferrerMgr struct {
	*_BaseMgr
}

func ReferrerMgr(db *gorm.DB) *_ReferrerMgr {
	if db == nil {
		panic(fmt.Errorf("ReferrerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ReferrerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("ps_referrer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

func (obj *_ReferrerMgr) GetTableName() string {
	return "ps_referrer"
}

func (obj *_ReferrerMgr) Get() (result Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

func (obj *_ReferrerMgr) Gets() (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}


func (obj *_ReferrerMgr) WithIDReferrer(idReferrer uint32) Option {
	return optionFunc(func(o *options) { o.query["id_referrer"] = idReferrer })
}

func (obj *_ReferrerMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

func (obj *_ReferrerMgr) WithPasswd(passwd string) Option {
	return optionFunc(func(o *options) { o.query["passwd"] = passwd })
}

func (obj *_ReferrerMgr) WithHTTPRefererRegexp(httpRefererRegexp string) Option {
	return optionFunc(func(o *options) { o.query["http_referer_regexp"] = httpRefererRegexp })
}

func (obj *_ReferrerMgr) WithHTTPRefererLike(httpRefererLike string) Option {
	return optionFunc(func(o *options) { o.query["http_referer_like"] = httpRefererLike })
}

func (obj *_ReferrerMgr) WithRequestURIRegexp(requestURIRegexp string) Option {
	return optionFunc(func(o *options) { o.query["request_uri_regexp"] = requestURIRegexp })
}

func (obj *_ReferrerMgr) WithRequestURILike(requestURILike string) Option {
	return optionFunc(func(o *options) { o.query["request_uri_like"] = requestURILike })
}

func (obj *_ReferrerMgr) WithHTTPRefererRegexpNot(httpRefererRegexpNot string) Option {
	return optionFunc(func(o *options) { o.query["http_referer_regexp_not"] = httpRefererRegexpNot })
}

func (obj *_ReferrerMgr) WithHTTPRefererLikeNot(httpRefererLikeNot string) Option {
	return optionFunc(func(o *options) { o.query["http_referer_like_not"] = httpRefererLikeNot })
}

func (obj *_ReferrerMgr) WithRequestURIRegexpNot(requestURIRegexpNot string) Option {
	return optionFunc(func(o *options) { o.query["request_uri_regexp_not"] = requestURIRegexpNot })
}

func (obj *_ReferrerMgr) WithRequestURILikeNot(requestURILikeNot string) Option {
	return optionFunc(func(o *options) { o.query["request_uri_like_not"] = requestURILikeNot })
}

func (obj *_ReferrerMgr) WithBaseFee(baseFee float64) Option {
	return optionFunc(func(o *options) { o.query["base_fee"] = baseFee })
}

func (obj *_ReferrerMgr) WithPercentFee(percentFee float64) Option {
	return optionFunc(func(o *options) { o.query["percent_fee"] = percentFee })
}

func (obj *_ReferrerMgr) WithClickFee(clickFee float64) Option {
	return optionFunc(func(o *options) { o.query["click_fee"] = clickFee })
}

func (obj *_ReferrerMgr) WithDateAdd(dateAdd time.Time) Option {
	return optionFunc(func(o *options) { o.query["date_add"] = dateAdd })
}

func (obj *_ReferrerMgr) GetByOption(opts ...Option) (result Referrer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

func (obj *_ReferrerMgr) GetByOptions(opts ...Option) (results []*Referrer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}


func (obj *_ReferrerMgr) GetFromIDReferrer(idReferrer uint32) (result Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&result).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromIDReferrer(idReferrers []uint32) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer IN (?)", idReferrers).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromName(name string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromName(names []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromPasswd(passwd string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd = ?", passwd).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromPasswd(passwds []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("passwd IN (?)", passwds).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromHTTPRefererRegexp(httpRefererRegexp string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_regexp = ?", httpRefererRegexp).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromHTTPRefererRegexp(httpRefererRegexps []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_regexp IN (?)", httpRefererRegexps).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromHTTPRefererLike(httpRefererLike string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_like = ?", httpRefererLike).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromHTTPRefererLike(httpRefererLikes []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_like IN (?)", httpRefererLikes).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromRequestURIRegexp(requestURIRegexp string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_regexp = ?", requestURIRegexp).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromRequestURIRegexp(requestURIRegexps []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_regexp IN (?)", requestURIRegexps).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromRequestURILike(requestURILike string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_like = ?", requestURILike).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromRequestURILike(requestURILikes []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_like IN (?)", requestURILikes).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromHTTPRefererRegexpNot(httpRefererRegexpNot string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_regexp_not = ?", httpRefererRegexpNot).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromHTTPRefererRegexpNot(httpRefererRegexpNots []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_regexp_not IN (?)", httpRefererRegexpNots).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromHTTPRefererLikeNot(httpRefererLikeNot string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_like_not = ?", httpRefererLikeNot).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromHTTPRefererLikeNot(httpRefererLikeNots []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("http_referer_like_not IN (?)", httpRefererLikeNots).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromRequestURIRegexpNot(requestURIRegexpNot string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_regexp_not = ?", requestURIRegexpNot).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromRequestURIRegexpNot(requestURIRegexpNots []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_regexp_not IN (?)", requestURIRegexpNots).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromRequestURILikeNot(requestURILikeNot string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_like_not = ?", requestURILikeNot).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromRequestURILikeNot(requestURILikeNots []string) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("request_uri_like_not IN (?)", requestURILikeNots).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromBaseFee(baseFee float64) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("base_fee = ?", baseFee).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromBaseFee(baseFees []float64) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("base_fee IN (?)", baseFees).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromPercentFee(percentFee float64) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("percent_fee = ?", percentFee).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromPercentFee(percentFees []float64) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("percent_fee IN (?)", percentFees).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromClickFee(clickFee float64) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("click_fee = ?", clickFee).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromClickFee(clickFees []float64) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("click_fee IN (?)", clickFees).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetFromDateAdd(dateAdd time.Time) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add = ?", dateAdd).Find(&results).Error

	return
}

func (obj *_ReferrerMgr) GetBatchFromDateAdd(dateAdds []time.Time) (results []*Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("date_add IN (?)", dateAdds).Find(&results).Error

	return
}


func (obj *_ReferrerMgr) FetchByPrimaryKey(idReferrer uint32) (result Referrer, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id_referrer = ?", idReferrer).Find(&result).Error

	return
}
