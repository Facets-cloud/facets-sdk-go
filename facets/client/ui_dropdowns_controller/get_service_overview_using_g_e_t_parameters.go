// Code generated by go-swagger; DO NOT EDIT.

package ui_dropdowns_controller

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

// NewGetServiceOverviewUsingGETParams creates a new GetServiceOverviewUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetServiceOverviewUsingGETParams() *GetServiceOverviewUsingGETParams {
	return &GetServiceOverviewUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceOverviewUsingGETParamsWithTimeout creates a new GetServiceOverviewUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetServiceOverviewUsingGETParamsWithTimeout(timeout time.Duration) *GetServiceOverviewUsingGETParams {
	return &GetServiceOverviewUsingGETParams{
		timeout: timeout,
	}
}

// NewGetServiceOverviewUsingGETParamsWithContext creates a new GetServiceOverviewUsingGETParams object
// with the ability to set a context for a request.
func NewGetServiceOverviewUsingGETParamsWithContext(ctx context.Context) *GetServiceOverviewUsingGETParams {
	return &GetServiceOverviewUsingGETParams{
		Context: ctx,
	}
}

// NewGetServiceOverviewUsingGETParamsWithHTTPClient creates a new GetServiceOverviewUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetServiceOverviewUsingGETParamsWithHTTPClient(client *http.Client) *GetServiceOverviewUsingGETParams {
	return &GetServiceOverviewUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetServiceOverviewUsingGETParams contains all the parameters to send to the API endpoint

	for the get service overview using g e t operation.

	Typically these are written to a http.Request.
*/
type GetServiceOverviewUsingGETParams struct {

	/* ServiceName.

	   serviceName
	*/
	ServiceName string

	/* StackName.

	   stackName
	*/
	StackName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get service overview using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceOverviewUsingGETParams) WithDefaults() *GetServiceOverviewUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get service overview using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetServiceOverviewUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) WithTimeout(timeout time.Duration) *GetServiceOverviewUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) WithContext(ctx context.Context) *GetServiceOverviewUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) WithHTTPClient(client *http.Client) *GetServiceOverviewUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) WithServiceName(serviceName string) *GetServiceOverviewUsingGETParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithStackName adds the stackName to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) WithStackName(stackName string) *GetServiceOverviewUsingGETParams {
	o.SetStackName(stackName)
	return o
}

// SetStackName adds the stackName to the get service overview using g e t params
func (o *GetServiceOverviewUsingGETParams) SetStackName(stackName string) {
	o.StackName = stackName
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceOverviewUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	// path param stackName
	if err := r.SetPathParam("stackName", o.StackName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
