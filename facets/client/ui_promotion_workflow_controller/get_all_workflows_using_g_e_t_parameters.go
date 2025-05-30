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

// NewGetAllWorkflowsUsingGETParams creates a new GetAllWorkflowsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllWorkflowsUsingGETParams() *GetAllWorkflowsUsingGETParams {
	return &GetAllWorkflowsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllWorkflowsUsingGETParamsWithTimeout creates a new GetAllWorkflowsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllWorkflowsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllWorkflowsUsingGETParams {
	return &GetAllWorkflowsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllWorkflowsUsingGETParamsWithContext creates a new GetAllWorkflowsUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllWorkflowsUsingGETParamsWithContext(ctx context.Context) *GetAllWorkflowsUsingGETParams {
	return &GetAllWorkflowsUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllWorkflowsUsingGETParamsWithHTTPClient creates a new GetAllWorkflowsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllWorkflowsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllWorkflowsUsingGETParams {
	return &GetAllWorkflowsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAllWorkflowsUsingGETParams contains all the parameters to send to the API endpoint

	for the get all workflows using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAllWorkflowsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all workflows using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllWorkflowsUsingGETParams) WithDefaults() *GetAllWorkflowsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all workflows using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllWorkflowsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all workflows using g e t params
func (o *GetAllWorkflowsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllWorkflowsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all workflows using g e t params
func (o *GetAllWorkflowsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all workflows using g e t params
func (o *GetAllWorkflowsUsingGETParams) WithContext(ctx context.Context) *GetAllWorkflowsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all workflows using g e t params
func (o *GetAllWorkflowsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all workflows using g e t params
func (o *GetAllWorkflowsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllWorkflowsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all workflows using g e t params
func (o *GetAllWorkflowsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllWorkflowsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
