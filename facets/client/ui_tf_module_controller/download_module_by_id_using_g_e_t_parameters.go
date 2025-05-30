// Code generated by go-swagger; DO NOT EDIT.

package ui_tf_module_controller

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

// NewDownloadModuleByIDUsingGETParams creates a new DownloadModuleByIDUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadModuleByIDUsingGETParams() *DownloadModuleByIDUsingGETParams {
	return &DownloadModuleByIDUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadModuleByIDUsingGETParamsWithTimeout creates a new DownloadModuleByIDUsingGETParams object
// with the ability to set a timeout on a request.
func NewDownloadModuleByIDUsingGETParamsWithTimeout(timeout time.Duration) *DownloadModuleByIDUsingGETParams {
	return &DownloadModuleByIDUsingGETParams{
		timeout: timeout,
	}
}

// NewDownloadModuleByIDUsingGETParamsWithContext creates a new DownloadModuleByIDUsingGETParams object
// with the ability to set a context for a request.
func NewDownloadModuleByIDUsingGETParamsWithContext(ctx context.Context) *DownloadModuleByIDUsingGETParams {
	return &DownloadModuleByIDUsingGETParams{
		Context: ctx,
	}
}

// NewDownloadModuleByIDUsingGETParamsWithHTTPClient creates a new DownloadModuleByIDUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadModuleByIDUsingGETParamsWithHTTPClient(client *http.Client) *DownloadModuleByIDUsingGETParams {
	return &DownloadModuleByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*
DownloadModuleByIDUsingGETParams contains all the parameters to send to the API endpoint

	for the download module by Id using g e t operation.

	Typically these are written to a http.Request.
*/
type DownloadModuleByIDUsingGETParams struct {

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the download module by Id using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadModuleByIDUsingGETParams) WithDefaults() *DownloadModuleByIDUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download module by Id using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadModuleByIDUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download module by Id using g e t params
func (o *DownloadModuleByIDUsingGETParams) WithTimeout(timeout time.Duration) *DownloadModuleByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download module by Id using g e t params
func (o *DownloadModuleByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download module by Id using g e t params
func (o *DownloadModuleByIDUsingGETParams) WithContext(ctx context.Context) *DownloadModuleByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download module by Id using g e t params
func (o *DownloadModuleByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download module by Id using g e t params
func (o *DownloadModuleByIDUsingGETParams) WithHTTPClient(client *http.Client) *DownloadModuleByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download module by Id using g e t params
func (o *DownloadModuleByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the download module by Id using g e t params
func (o *DownloadModuleByIDUsingGETParams) WithID(id string) *DownloadModuleByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the download module by Id using g e t params
func (o *DownloadModuleByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadModuleByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
