// Code generated by go-swagger; DO NOT EDIT.

package ui_resource_group_controller

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

// NewGetResourceGroupsForSessionUserUsingGETParams creates a new GetResourceGroupsForSessionUserUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceGroupsForSessionUserUsingGETParams() *GetResourceGroupsForSessionUserUsingGETParams {
	return &GetResourceGroupsForSessionUserUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceGroupsForSessionUserUsingGETParamsWithTimeout creates a new GetResourceGroupsForSessionUserUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetResourceGroupsForSessionUserUsingGETParamsWithTimeout(timeout time.Duration) *GetResourceGroupsForSessionUserUsingGETParams {
	return &GetResourceGroupsForSessionUserUsingGETParams{
		timeout: timeout,
	}
}

// NewGetResourceGroupsForSessionUserUsingGETParamsWithContext creates a new GetResourceGroupsForSessionUserUsingGETParams object
// with the ability to set a context for a request.
func NewGetResourceGroupsForSessionUserUsingGETParamsWithContext(ctx context.Context) *GetResourceGroupsForSessionUserUsingGETParams {
	return &GetResourceGroupsForSessionUserUsingGETParams{
		Context: ctx,
	}
}

// NewGetResourceGroupsForSessionUserUsingGETParamsWithHTTPClient creates a new GetResourceGroupsForSessionUserUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceGroupsForSessionUserUsingGETParamsWithHTTPClient(client *http.Client) *GetResourceGroupsForSessionUserUsingGETParams {
	return &GetResourceGroupsForSessionUserUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetResourceGroupsForSessionUserUsingGETParams contains all the parameters to send to the API endpoint

	for the get resource groups for session user using g e t operation.

	Typically these are written to a http.Request.
*/
type GetResourceGroupsForSessionUserUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource groups for session user using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceGroupsForSessionUserUsingGETParams) WithDefaults() *GetResourceGroupsForSessionUserUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource groups for session user using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceGroupsForSessionUserUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource groups for session user using g e t params
func (o *GetResourceGroupsForSessionUserUsingGETParams) WithTimeout(timeout time.Duration) *GetResourceGroupsForSessionUserUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource groups for session user using g e t params
func (o *GetResourceGroupsForSessionUserUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource groups for session user using g e t params
func (o *GetResourceGroupsForSessionUserUsingGETParams) WithContext(ctx context.Context) *GetResourceGroupsForSessionUserUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource groups for session user using g e t params
func (o *GetResourceGroupsForSessionUserUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource groups for session user using g e t params
func (o *GetResourceGroupsForSessionUserUsingGETParams) WithHTTPClient(client *http.Client) *GetResourceGroupsForSessionUserUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource groups for session user using g e t params
func (o *GetResourceGroupsForSessionUserUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceGroupsForSessionUserUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
