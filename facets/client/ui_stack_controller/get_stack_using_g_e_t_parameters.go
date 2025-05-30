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

// NewGetStackUsingGETParams creates a new GetStackUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStackUsingGETParams() *GetStackUsingGETParams {
	return &GetStackUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStackUsingGETParamsWithTimeout creates a new GetStackUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetStackUsingGETParamsWithTimeout(timeout time.Duration) *GetStackUsingGETParams {
	return &GetStackUsingGETParams{
		timeout: timeout,
	}
}

// NewGetStackUsingGETParamsWithContext creates a new GetStackUsingGETParams object
// with the ability to set a context for a request.
func NewGetStackUsingGETParamsWithContext(ctx context.Context) *GetStackUsingGETParams {
	return &GetStackUsingGETParams{
		Context: ctx,
	}
}

// NewGetStackUsingGETParamsWithHTTPClient creates a new GetStackUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStackUsingGETParamsWithHTTPClient(client *http.Client) *GetStackUsingGETParams {
	return &GetStackUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetStackUsingGETParams contains all the parameters to send to the API endpoint

	for the get stack using g e t operation.

	Typically these are written to a http.Request.
*/
type GetStackUsingGETParams struct {

	/* StackName.

	   stackName
	*/
	StackName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get stack using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStackUsingGETParams) WithDefaults() *GetStackUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get stack using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStackUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get stack using g e t params
func (o *GetStackUsingGETParams) WithTimeout(timeout time.Duration) *GetStackUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stack using g e t params
func (o *GetStackUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stack using g e t params
func (o *GetStackUsingGETParams) WithContext(ctx context.Context) *GetStackUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stack using g e t params
func (o *GetStackUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stack using g e t params
func (o *GetStackUsingGETParams) WithHTTPClient(client *http.Client) *GetStackUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stack using g e t params
func (o *GetStackUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackName adds the stackName to the get stack using g e t params
func (o *GetStackUsingGETParams) WithStackName(stackName string) *GetStackUsingGETParams {
	o.SetStackName(stackName)
	return o
}

// SetStackName adds the stackName to the get stack using g e t params
func (o *GetStackUsingGETParams) SetStackName(stackName string) {
	o.StackName = stackName
}

// WriteToRequest writes these params to a swagger request
func (o *GetStackUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
