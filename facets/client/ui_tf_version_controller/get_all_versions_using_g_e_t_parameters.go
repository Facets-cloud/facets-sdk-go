// Code generated by go-swagger; DO NOT EDIT.

package ui_tf_version_controller

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

// NewGetAllVersionsUsingGETParams creates a new GetAllVersionsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllVersionsUsingGETParams() *GetAllVersionsUsingGETParams {
	return &GetAllVersionsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllVersionsUsingGETParamsWithTimeout creates a new GetAllVersionsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllVersionsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllVersionsUsingGETParams {
	return &GetAllVersionsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllVersionsUsingGETParamsWithContext creates a new GetAllVersionsUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllVersionsUsingGETParamsWithContext(ctx context.Context) *GetAllVersionsUsingGETParams {
	return &GetAllVersionsUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllVersionsUsingGETParamsWithHTTPClient creates a new GetAllVersionsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllVersionsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllVersionsUsingGETParams {
	return &GetAllVersionsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAllVersionsUsingGETParams contains all the parameters to send to the API endpoint

	for the get all versions using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAllVersionsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all versions using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllVersionsUsingGETParams) WithDefaults() *GetAllVersionsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all versions using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllVersionsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all versions using g e t params
func (o *GetAllVersionsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllVersionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all versions using g e t params
func (o *GetAllVersionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all versions using g e t params
func (o *GetAllVersionsUsingGETParams) WithContext(ctx context.Context) *GetAllVersionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all versions using g e t params
func (o *GetAllVersionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all versions using g e t params
func (o *GetAllVersionsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllVersionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all versions using g e t params
func (o *GetAllVersionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllVersionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
