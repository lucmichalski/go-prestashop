package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

var globalIsRelated bool = true // 全局预加载

type _BaseMgr struct {
	*gorm.DB
	ctx       context.Context
	cancel    context.CancelFunc
	timeout   time.Duration
	isRelated bool
}

func (obj *_BaseMgr) SetTimeOut(timeout time.Duration) {
	obj.ctx, obj.cancel = context.WithTimeout(context.Background(), timeout)
	obj.timeout = timeout
}

func (obj *_BaseMgr) SetCtx(c context.Context) {
	if c != nil {
		obj.ctx = c
	}
}

func (obj *_BaseMgr) GetCtx() context.Context {
	return obj.ctx
}

func (obj *_BaseMgr) Cancel(c context.Context) {
	obj.cancel()
}

func (obj *_BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

func (obj *_BaseMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

func (obj *_BaseMgr) GetIsRelated() bool {
	return obj.isRelated
}

func (obj *_BaseMgr) SetIsRelated(b bool) {
	obj.isRelated = b
}

func (obj *_BaseMgr) New() *gorm.DB {
	return obj.DB.Session(&gorm.Session{Context: obj.ctx})
}

type options struct {
	query map[string]interface{}
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func OpenRelated() {
	globalIsRelated = true
}

func CloseRelated() {
	globalIsRelated = true
}
