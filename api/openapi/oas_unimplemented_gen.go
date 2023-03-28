// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AddPage implements addPage operation.
//
// Add new page.
//
// POST /pages
func (UnimplementedHandler) AddPage(ctx context.Context, req OptAddPageReq, params AddPageParams) (r *Page, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFile implements getFile operation.
//
// Get file content.
//
// GET /pages/{id}/file/{file_id}
func (UnimplementedHandler) GetFile(ctx context.Context, params GetFileParams) (r GetFileRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPage implements getPage operation.
//
// Get page details.
//
// GET /pages/{id}
func (UnimplementedHandler) GetPage(ctx context.Context, params GetPageParams) (r GetPageRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPages implements getPages operation.
//
// Get all pages.
//
// GET /pages
func (UnimplementedHandler) GetPages(ctx context.Context) (r Pages, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
