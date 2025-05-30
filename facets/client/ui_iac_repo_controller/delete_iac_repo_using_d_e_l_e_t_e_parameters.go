// Code generated by go-swagger; DO NOT EDIT.

package ui_iac_repo_controller

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

// NewDeleteIacRepoUsingDELETEParams creates a new DeleteIacRepoUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIacRepoUsingDELETEParams() *DeleteIacRepoUsingDELETEParams {
	return &DeleteIacRepoUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIacRepoUsingDELETEParamsWithTimeout creates a new DeleteIacRepoUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteIacRepoUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteIacRepoUsingDELETEParams {
	return &DeleteIacRepoUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteIacRepoUsingDELETEParamsWithContext creates a new DeleteIacRepoUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteIacRepoUsingDELETEParamsWithContext(ctx context.Context) *DeleteIacRepoUsingDELETEParams {
	return &DeleteIacRepoUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteIacRepoUsingDELETEParamsWithHTTPClient creates a new DeleteIacRepoUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIacRepoUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteIacRepoUsingDELETEParams {
	return &DeleteIacRepoUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteIacRepoUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete iac repo using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteIacRepoUsingDELETEParams struct {

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete iac repo using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIacRepoUsingDELETEParams) WithDefaults() *DeleteIacRepoUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete iac repo using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIacRepoUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete iac repo using d e l e t e params
func (o *DeleteIacRepoUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteIacRepoUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete iac repo using d e l e t e params
func (o *DeleteIacRepoUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete iac repo using d e l e t e params
func (o *DeleteIacRepoUsingDELETEParams) WithContext(ctx context.Context) *DeleteIacRepoUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete iac repo using d e l e t e params
func (o *DeleteIacRepoUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete iac repo using d e l e t e params
func (o *DeleteIacRepoUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteIacRepoUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete iac repo using d e l e t e params
func (o *DeleteIacRepoUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete iac repo using d e l e t e params
func (o *DeleteIacRepoUsingDELETEParams) WithID(id string) *DeleteIacRepoUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete iac repo using d e l e t e params
func (o *DeleteIacRepoUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIacRepoUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
