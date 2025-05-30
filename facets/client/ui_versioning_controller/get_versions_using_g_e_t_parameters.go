// Code generated by go-swagger; DO NOT EDIT.

package ui_versioning_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetVersionsUsingGETParams creates a new GetVersionsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVersionsUsingGETParams() *GetVersionsUsingGETParams {
	return &GetVersionsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionsUsingGETParamsWithTimeout creates a new GetVersionsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetVersionsUsingGETParamsWithTimeout(timeout time.Duration) *GetVersionsUsingGETParams {
	return &GetVersionsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetVersionsUsingGETParamsWithContext creates a new GetVersionsUsingGETParams object
// with the ability to set a context for a request.
func NewGetVersionsUsingGETParamsWithContext(ctx context.Context) *GetVersionsUsingGETParams {
	return &GetVersionsUsingGETParams{
		Context: ctx,
	}
}

// NewGetVersionsUsingGETParamsWithHTTPClient creates a new GetVersionsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVersionsUsingGETParamsWithHTTPClient(client *http.Client) *GetVersionsUsingGETParams {
	return &GetVersionsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetVersionsUsingGETParams contains all the parameters to send to the API endpoint

	for the get versions using g e t operation.

	Typically these are written to a http.Request.
*/
type GetVersionsUsingGETParams struct {

	/* Page.

	   page

	   Format: int32
	*/
	Page *int32

	/* PerPage.

	   perPage

	   Format: int32
	   Default: 10
	*/
	PerPage *int32

	/* VersioningKey.

	   versioningKey
	*/
	VersioningKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get versions using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionsUsingGETParams) WithDefaults() *GetVersionsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get versions using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionsUsingGETParams) SetDefaults() {
	var (
		pageDefault = int32(0)

		perPageDefault = int32(10)
	)

	val := GetVersionsUsingGETParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get versions using g e t params
func (o *GetVersionsUsingGETParams) WithTimeout(timeout time.Duration) *GetVersionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get versions using g e t params
func (o *GetVersionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get versions using g e t params
func (o *GetVersionsUsingGETParams) WithContext(ctx context.Context) *GetVersionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get versions using g e t params
func (o *GetVersionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get versions using g e t params
func (o *GetVersionsUsingGETParams) WithHTTPClient(client *http.Client) *GetVersionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get versions using g e t params
func (o *GetVersionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get versions using g e t params
func (o *GetVersionsUsingGETParams) WithPage(page *int32) *GetVersionsUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get versions using g e t params
func (o *GetVersionsUsingGETParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get versions using g e t params
func (o *GetVersionsUsingGETParams) WithPerPage(perPage *int32) *GetVersionsUsingGETParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get versions using g e t params
func (o *GetVersionsUsingGETParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithVersioningKey adds the versioningKey to the get versions using g e t params
func (o *GetVersionsUsingGETParams) WithVersioningKey(versioningKey string) *GetVersionsUsingGETParams {
	o.SetVersioningKey(versioningKey)
	return o
}

// SetVersioningKey adds the versioningKey to the get versions using g e t params
func (o *GetVersionsUsingGETParams) SetVersioningKey(versioningKey string) {
	o.VersioningKey = versioningKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param perPage
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("perPage", qPerPage); err != nil {
				return err
			}
		}
	}

	// path param versioningKey
	if err := r.SetPathParam("versioningKey", o.VersioningKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
