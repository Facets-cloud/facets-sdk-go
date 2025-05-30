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

// NewGetAllClusterResourcesByStackUsingGETParams creates a new GetAllClusterResourcesByStackUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllClusterResourcesByStackUsingGETParams() *GetAllClusterResourcesByStackUsingGETParams {
	return &GetAllClusterResourcesByStackUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllClusterResourcesByStackUsingGETParamsWithTimeout creates a new GetAllClusterResourcesByStackUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllClusterResourcesByStackUsingGETParamsWithTimeout(timeout time.Duration) *GetAllClusterResourcesByStackUsingGETParams {
	return &GetAllClusterResourcesByStackUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllClusterResourcesByStackUsingGETParamsWithContext creates a new GetAllClusterResourcesByStackUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllClusterResourcesByStackUsingGETParamsWithContext(ctx context.Context) *GetAllClusterResourcesByStackUsingGETParams {
	return &GetAllClusterResourcesByStackUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllClusterResourcesByStackUsingGETParamsWithHTTPClient creates a new GetAllClusterResourcesByStackUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllClusterResourcesByStackUsingGETParamsWithHTTPClient(client *http.Client) *GetAllClusterResourcesByStackUsingGETParams {
	return &GetAllClusterResourcesByStackUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAllClusterResourcesByStackUsingGETParams contains all the parameters to send to the API endpoint

	for the get all cluster resources by stack using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAllClusterResourcesByStackUsingGETParams struct {

	/* ResourceName.

	   resourceName
	*/
	ResourceName string

	/* ResourceType.

	   resourceType
	*/
	ResourceType string

	/* StackName.

	   stackName
	*/
	StackName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all cluster resources by stack using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllClusterResourcesByStackUsingGETParams) WithDefaults() *GetAllClusterResourcesByStackUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all cluster resources by stack using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllClusterResourcesByStackUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) WithTimeout(timeout time.Duration) *GetAllClusterResourcesByStackUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) WithContext(ctx context.Context) *GetAllClusterResourcesByStackUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) WithHTTPClient(client *http.Client) *GetAllClusterResourcesByStackUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName adds the resourceName to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) WithResourceName(resourceName string) *GetAllClusterResourcesByStackUsingGETParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WithResourceType adds the resourceType to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) WithResourceType(resourceType string) *GetAllClusterResourcesByStackUsingGETParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithStackName adds the stackName to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) WithStackName(stackName string) *GetAllClusterResourcesByStackUsingGETParams {
	o.SetStackName(stackName)
	return o
}

// SetStackName adds the stackName to the get all cluster resources by stack using g e t params
func (o *GetAllClusterResourcesByStackUsingGETParams) SetStackName(stackName string) {
	o.StackName = stackName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllClusterResourcesByStackUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceName
	if err := r.SetPathParam("resourceName", o.ResourceName); err != nil {
		return err
	}

	// path param resourceType
	if err := r.SetPathParam("resourceType", o.ResourceType); err != nil {
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
