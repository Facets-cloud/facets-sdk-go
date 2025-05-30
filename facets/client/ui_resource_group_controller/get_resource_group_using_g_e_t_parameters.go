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

// NewGetResourceGroupUsingGETParams creates a new GetResourceGroupUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceGroupUsingGETParams() *GetResourceGroupUsingGETParams {
	return &GetResourceGroupUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceGroupUsingGETParamsWithTimeout creates a new GetResourceGroupUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetResourceGroupUsingGETParamsWithTimeout(timeout time.Duration) *GetResourceGroupUsingGETParams {
	return &GetResourceGroupUsingGETParams{
		timeout: timeout,
	}
}

// NewGetResourceGroupUsingGETParamsWithContext creates a new GetResourceGroupUsingGETParams object
// with the ability to set a context for a request.
func NewGetResourceGroupUsingGETParamsWithContext(ctx context.Context) *GetResourceGroupUsingGETParams {
	return &GetResourceGroupUsingGETParams{
		Context: ctx,
	}
}

// NewGetResourceGroupUsingGETParamsWithHTTPClient creates a new GetResourceGroupUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceGroupUsingGETParamsWithHTTPClient(client *http.Client) *GetResourceGroupUsingGETParams {
	return &GetResourceGroupUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetResourceGroupUsingGETParams contains all the parameters to send to the API endpoint

	for the get resource group using g e t operation.

	Typically these are written to a http.Request.
*/
type GetResourceGroupUsingGETParams struct {

	/* ResourceGroupID.

	   resourceGroupId
	*/
	ResourceGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource group using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceGroupUsingGETParams) WithDefaults() *GetResourceGroupUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource group using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceGroupUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource group using g e t params
func (o *GetResourceGroupUsingGETParams) WithTimeout(timeout time.Duration) *GetResourceGroupUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource group using g e t params
func (o *GetResourceGroupUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource group using g e t params
func (o *GetResourceGroupUsingGETParams) WithContext(ctx context.Context) *GetResourceGroupUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource group using g e t params
func (o *GetResourceGroupUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource group using g e t params
func (o *GetResourceGroupUsingGETParams) WithHTTPClient(client *http.Client) *GetResourceGroupUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource group using g e t params
func (o *GetResourceGroupUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceGroupID adds the resourceGroupID to the get resource group using g e t params
func (o *GetResourceGroupUsingGETParams) WithResourceGroupID(resourceGroupID string) *GetResourceGroupUsingGETParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get resource group using g e t params
func (o *GetResourceGroupUsingGETParams) SetResourceGroupID(resourceGroupID string) {
	o.ResourceGroupID = resourceGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceGroupUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceGroupId
	if err := r.SetPathParam("resourceGroupId", o.ResourceGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
