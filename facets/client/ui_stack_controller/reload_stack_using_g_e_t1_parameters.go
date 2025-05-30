// Code generated by go-swagger; DO NOT EDIT.

package ui_stack_controller

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

// NewReloadStackUsingGET1Params creates a new ReloadStackUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReloadStackUsingGET1Params() *ReloadStackUsingGET1Params {
	return &ReloadStackUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewReloadStackUsingGET1ParamsWithTimeout creates a new ReloadStackUsingGET1Params object
// with the ability to set a timeout on a request.
func NewReloadStackUsingGET1ParamsWithTimeout(timeout time.Duration) *ReloadStackUsingGET1Params {
	return &ReloadStackUsingGET1Params{
		timeout: timeout,
	}
}

// NewReloadStackUsingGET1ParamsWithContext creates a new ReloadStackUsingGET1Params object
// with the ability to set a context for a request.
func NewReloadStackUsingGET1ParamsWithContext(ctx context.Context) *ReloadStackUsingGET1Params {
	return &ReloadStackUsingGET1Params{
		Context: ctx,
	}
}

// NewReloadStackUsingGET1ParamsWithHTTPClient creates a new ReloadStackUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewReloadStackUsingGET1ParamsWithHTTPClient(client *http.Client) *ReloadStackUsingGET1Params {
	return &ReloadStackUsingGET1Params{
		HTTPClient: client,
	}
}

/*
ReloadStackUsingGET1Params contains all the parameters to send to the API endpoint

	for the reload stack using g e t 1 operation.

	Typically these are written to a http.Request.
*/
type ReloadStackUsingGET1Params struct {

	/* StackName.

	   stackName
	*/
	StackName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reload stack using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReloadStackUsingGET1Params) WithDefaults() *ReloadStackUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reload stack using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReloadStackUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reload stack using g e t 1 params
func (o *ReloadStackUsingGET1Params) WithTimeout(timeout time.Duration) *ReloadStackUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reload stack using g e t 1 params
func (o *ReloadStackUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reload stack using g e t 1 params
func (o *ReloadStackUsingGET1Params) WithContext(ctx context.Context) *ReloadStackUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reload stack using g e t 1 params
func (o *ReloadStackUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reload stack using g e t 1 params
func (o *ReloadStackUsingGET1Params) WithHTTPClient(client *http.Client) *ReloadStackUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reload stack using g e t 1 params
func (o *ReloadStackUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackName adds the stackName to the reload stack using g e t 1 params
func (o *ReloadStackUsingGET1Params) WithStackName(stackName string) *ReloadStackUsingGET1Params {
	o.SetStackName(stackName)
	return o
}

// SetStackName adds the stackName to the reload stack using g e t 1 params
func (o *ReloadStackUsingGET1Params) SetStackName(stackName string) {
	o.StackName = stackName
}

// WriteToRequest writes these params to a swagger request
func (o *ReloadStackUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stackName
	if err := r.SetPathParam("stackName", o.StackName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
