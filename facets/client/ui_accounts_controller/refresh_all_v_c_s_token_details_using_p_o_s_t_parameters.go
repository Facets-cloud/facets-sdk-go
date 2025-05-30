// Code generated by go-swagger; DO NOT EDIT.

package ui_accounts_controller

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
)

// NewRefreshAllVCSTokenDetailsUsingPOSTParams creates a new RefreshAllVCSTokenDetailsUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRefreshAllVCSTokenDetailsUsingPOSTParams() *RefreshAllVCSTokenDetailsUsingPOSTParams {
	return &RefreshAllVCSTokenDetailsUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshAllVCSTokenDetailsUsingPOSTParamsWithTimeout creates a new RefreshAllVCSTokenDetailsUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewRefreshAllVCSTokenDetailsUsingPOSTParamsWithTimeout(timeout time.Duration) *RefreshAllVCSTokenDetailsUsingPOSTParams {
	return &RefreshAllVCSTokenDetailsUsingPOSTParams{
		timeout: timeout,
	}
}

// NewRefreshAllVCSTokenDetailsUsingPOSTParamsWithContext creates a new RefreshAllVCSTokenDetailsUsingPOSTParams object
// with the ability to set a context for a request.
func NewRefreshAllVCSTokenDetailsUsingPOSTParamsWithContext(ctx context.Context) *RefreshAllVCSTokenDetailsUsingPOSTParams {
	return &RefreshAllVCSTokenDetailsUsingPOSTParams{
		Context: ctx,
	}
}

// NewRefreshAllVCSTokenDetailsUsingPOSTParamsWithHTTPClient creates a new RefreshAllVCSTokenDetailsUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewRefreshAllVCSTokenDetailsUsingPOSTParamsWithHTTPClient(client *http.Client) *RefreshAllVCSTokenDetailsUsingPOSTParams {
	return &RefreshAllVCSTokenDetailsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
RefreshAllVCSTokenDetailsUsingPOSTParams contains all the parameters to send to the API endpoint

	for the refresh all v c s token details using p o s t operation.

	Typically these are written to a http.Request.
*/
type RefreshAllVCSTokenDetailsUsingPOSTParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the refresh all v c s token details using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) WithDefaults() *RefreshAllVCSTokenDetailsUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the refresh all v c s token details using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the refresh all v c s token details using p o s t params
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) WithTimeout(timeout time.Duration) *RefreshAllVCSTokenDetailsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh all v c s token details using p o s t params
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh all v c s token details using p o s t params
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) WithContext(ctx context.Context) *RefreshAllVCSTokenDetailsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh all v c s token details using p o s t params
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh all v c s token details using p o s t params
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) WithHTTPClient(client *http.Client) *RefreshAllVCSTokenDetailsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh all v c s token details using p o s t params
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshAllVCSTokenDetailsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
