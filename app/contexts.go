// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "demo": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/cratermoon/goaelog/design
// --out=$(GOPATH)/src/github.com/cratermoon/goaelog
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// TimeDatetimeContext provides the datetime time action context.
type TimeDatetimeContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Long bool
	Tz   string
}

// NewTimeDatetimeContext parses the incoming request URL and body, performs validations and creates the
// context used by the datetime controller time action.
func NewTimeDatetimeContext(ctx context.Context, r *http.Request, service *goa.Service) (*TimeDatetimeContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := TimeDatetimeContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLong := req.Params["long"]
	if len(paramLong) > 0 {
		rawLong := paramLong[0]
		if long, err2 := strconv.ParseBool(rawLong); err2 == nil {
			rctx.Long = long
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("long", rawLong, "boolean"))
		}
	}
	paramTz := req.Params["tz"]
	if len(paramTz) > 0 {
		rawTz := paramTz[0]
		rctx.Tz = rawTz
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *TimeDatetimeContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// GenerateUUIDContext provides the uuid generate action context.
type GenerateUUIDContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGenerateUUIDContext parses the incoming request URL and body, performs validations and creates the
// context used by the uuid controller generate action.
func NewGenerateUUIDContext(ctx context.Context, r *http.Request, service *goa.Service) (*GenerateUUIDContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GenerateUUIDContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GenerateUUIDContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}