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

// NewSoftDeletedEntitiesUsingGETParams creates a new SoftDeletedEntitiesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSoftDeletedEntitiesUsingGETParams() *SoftDeletedEntitiesUsingGETParams {
	return &SoftDeletedEntitiesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSoftDeletedEntitiesUsingGETParamsWithTimeout creates a new SoftDeletedEntitiesUsingGETParams object
// with the ability to set a timeout on a request.
func NewSoftDeletedEntitiesUsingGETParamsWithTimeout(timeout time.Duration) *SoftDeletedEntitiesUsingGETParams {
	return &SoftDeletedEntitiesUsingGETParams{
		timeout: timeout,
	}
}

// NewSoftDeletedEntitiesUsingGETParamsWithContext creates a new SoftDeletedEntitiesUsingGETParams object
// with the ability to set a context for a request.
func NewSoftDeletedEntitiesUsingGETParamsWithContext(ctx context.Context) *SoftDeletedEntitiesUsingGETParams {
	return &SoftDeletedEntitiesUsingGETParams{
		Context: ctx,
	}
}

// NewSoftDeletedEntitiesUsingGETParamsWithHTTPClient creates a new SoftDeletedEntitiesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewSoftDeletedEntitiesUsingGETParamsWithHTTPClient(client *http.Client) *SoftDeletedEntitiesUsingGETParams {
	return &SoftDeletedEntitiesUsingGETParams{
		HTTPClient: client,
	}
}

/*
SoftDeletedEntitiesUsingGETParams contains all the parameters to send to the API endpoint

	for the soft deleted entities using g e t operation.

	Typically these are written to a http.Request.
*/
type SoftDeletedEntitiesUsingGETParams struct {

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

	/* SortBy.

	   sortBy
	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the soft deleted entities using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftDeletedEntitiesUsingGETParams) WithDefaults() *SoftDeletedEntitiesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the soft deleted entities using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftDeletedEntitiesUsingGETParams) SetDefaults() {
	var (
		pageDefault = int32(0)

		perPageDefault = int32(10)
	)

	val := SoftDeletedEntitiesUsingGETParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) WithTimeout(timeout time.Duration) *SoftDeletedEntitiesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) WithContext(ctx context.Context) *SoftDeletedEntitiesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) WithHTTPClient(client *http.Client) *SoftDeletedEntitiesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) WithPage(page *int32) *SoftDeletedEntitiesUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) WithPerPage(perPage *int32) *SoftDeletedEntitiesUsingGETParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithSortBy adds the sortBy to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) WithSortBy(sortBy *string) *SoftDeletedEntitiesUsingGETParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the soft deleted entities using g e t params
func (o *SoftDeletedEntitiesUsingGETParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *SoftDeletedEntitiesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
