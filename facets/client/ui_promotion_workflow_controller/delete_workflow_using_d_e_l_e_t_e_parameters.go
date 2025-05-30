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

// NewDeleteWorkflowUsingDELETEParams creates a new DeleteWorkflowUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWorkflowUsingDELETEParams() *DeleteWorkflowUsingDELETEParams {
	return &DeleteWorkflowUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkflowUsingDELETEParamsWithTimeout creates a new DeleteWorkflowUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteWorkflowUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteWorkflowUsingDELETEParams {
	return &DeleteWorkflowUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteWorkflowUsingDELETEParamsWithContext creates a new DeleteWorkflowUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteWorkflowUsingDELETEParamsWithContext(ctx context.Context) *DeleteWorkflowUsingDELETEParams {
	return &DeleteWorkflowUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteWorkflowUsingDELETEParamsWithHTTPClient creates a new DeleteWorkflowUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWorkflowUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteWorkflowUsingDELETEParams {
	return &DeleteWorkflowUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteWorkflowUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete workflow using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteWorkflowUsingDELETEParams struct {

	/* WorkflowID.

	   workflowId
	*/
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete workflow using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkflowUsingDELETEParams) WithDefaults() *DeleteWorkflowUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete workflow using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkflowUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete workflow using d e l e t e params
func (o *DeleteWorkflowUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteWorkflowUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workflow using d e l e t e params
func (o *DeleteWorkflowUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workflow using d e l e t e params
func (o *DeleteWorkflowUsingDELETEParams) WithContext(ctx context.Context) *DeleteWorkflowUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workflow using d e l e t e params
func (o *DeleteWorkflowUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workflow using d e l e t e params
func (o *DeleteWorkflowUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteWorkflowUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workflow using d e l e t e params
func (o *DeleteWorkflowUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkflowID adds the workflowID to the delete workflow using d e l e t e params
func (o *DeleteWorkflowUsingDELETEParams) WithWorkflowID(workflowID string) *DeleteWorkflowUsingDELETEParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the delete workflow using d e l e t e params
func (o *DeleteWorkflowUsingDELETEParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkflowUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workflowId
	if err := r.SetPathParam("workflowId", o.WorkflowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
