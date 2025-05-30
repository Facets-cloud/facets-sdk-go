// Code generated by go-swagger; DO NOT EDIT.

package ui_promotion_workflow_controller

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

// NewGetWorkflowsByStackUsingGETParams creates a new GetWorkflowsByStackUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowsByStackUsingGETParams() *GetWorkflowsByStackUsingGETParams {
	return &GetWorkflowsByStackUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowsByStackUsingGETParamsWithTimeout creates a new GetWorkflowsByStackUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowsByStackUsingGETParamsWithTimeout(timeout time.Duration) *GetWorkflowsByStackUsingGETParams {
	return &GetWorkflowsByStackUsingGETParams{
		timeout: timeout,
	}
}

// NewGetWorkflowsByStackUsingGETParamsWithContext creates a new GetWorkflowsByStackUsingGETParams object
// with the ability to set a context for a request.
func NewGetWorkflowsByStackUsingGETParamsWithContext(ctx context.Context) *GetWorkflowsByStackUsingGETParams {
	return &GetWorkflowsByStackUsingGETParams{
		Context: ctx,
	}
}

// NewGetWorkflowsByStackUsingGETParamsWithHTTPClient creates a new GetWorkflowsByStackUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowsByStackUsingGETParamsWithHTTPClient(client *http.Client) *GetWorkflowsByStackUsingGETParams {
	return &GetWorkflowsByStackUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetWorkflowsByStackUsingGETParams contains all the parameters to send to the API endpoint

	for the get workflows by stack using g e t operation.

	Typically these are written to a http.Request.
*/
type GetWorkflowsByStackUsingGETParams struct {

	/* StackName.

	   stackName
	*/
	StackName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflows by stack using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowsByStackUsingGETParams) WithDefaults() *GetWorkflowsByStackUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflows by stack using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowsByStackUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflows by stack using g e t params
func (o *GetWorkflowsByStackUsingGETParams) WithTimeout(timeout time.Duration) *GetWorkflowsByStackUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflows by stack using g e t params
func (o *GetWorkflowsByStackUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflows by stack using g e t params
func (o *GetWorkflowsByStackUsingGETParams) WithContext(ctx context.Context) *GetWorkflowsByStackUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflows by stack using g e t params
func (o *GetWorkflowsByStackUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflows by stack using g e t params
func (o *GetWorkflowsByStackUsingGETParams) WithHTTPClient(client *http.Client) *GetWorkflowsByStackUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflows by stack using g e t params
func (o *GetWorkflowsByStackUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackName adds the stackName to the get workflows by stack using g e t params
func (o *GetWorkflowsByStackUsingGETParams) WithStackName(stackName string) *GetWorkflowsByStackUsingGETParams {
	o.SetStackName(stackName)
	return o
}

// SetStackName adds the stackName to the get workflows by stack using g e t params
func (o *GetWorkflowsByStackUsingGETParams) SetStackName(stackName string) {
	o.StackName = stackName
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowsByStackUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
