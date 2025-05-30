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

// NewGetGitHistoryForResourceUsingGETParams creates a new GetGitHistoryForResourceUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGitHistoryForResourceUsingGETParams() *GetGitHistoryForResourceUsingGETParams {
	return &GetGitHistoryForResourceUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGitHistoryForResourceUsingGETParamsWithTimeout creates a new GetGitHistoryForResourceUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetGitHistoryForResourceUsingGETParamsWithTimeout(timeout time.Duration) *GetGitHistoryForResourceUsingGETParams {
	return &GetGitHistoryForResourceUsingGETParams{
		timeout: timeout,
	}
}

// NewGetGitHistoryForResourceUsingGETParamsWithContext creates a new GetGitHistoryForResourceUsingGETParams object
// with the ability to set a context for a request.
func NewGetGitHistoryForResourceUsingGETParamsWithContext(ctx context.Context) *GetGitHistoryForResourceUsingGETParams {
	return &GetGitHistoryForResourceUsingGETParams{
		Context: ctx,
	}
}

// NewGetGitHistoryForResourceUsingGETParamsWithHTTPClient creates a new GetGitHistoryForResourceUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGitHistoryForResourceUsingGETParamsWithHTTPClient(client *http.Client) *GetGitHistoryForResourceUsingGETParams {
	return &GetGitHistoryForResourceUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetGitHistoryForResourceUsingGETParams contains all the parameters to send to the API endpoint

	for the get git history for resource using g e t operation.

	Typically these are written to a http.Request.
*/
type GetGitHistoryForResourceUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* ResourceName.

	   resourceName
	*/
	ResourceName string

	/* ResourceType.

	   resourceType
	*/
	ResourceType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get git history for resource using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitHistoryForResourceUsingGETParams) WithDefaults() *GetGitHistoryForResourceUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get git history for resource using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitHistoryForResourceUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) WithTimeout(timeout time.Duration) *GetGitHistoryForResourceUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) WithContext(ctx context.Context) *GetGitHistoryForResourceUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) WithHTTPClient(client *http.Client) *GetGitHistoryForResourceUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) WithClusterID(clusterID string) *GetGitHistoryForResourceUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithResourceName adds the resourceName to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) WithResourceName(resourceName string) *GetGitHistoryForResourceUsingGETParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WithResourceType adds the resourceType to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) WithResourceType(resourceType string) *GetGitHistoryForResourceUsingGETParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get git history for resource using g e t params
func (o *GetGitHistoryForResourceUsingGETParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *GetGitHistoryForResourceUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	// path param resourceName
	if err := r.SetPathParam("resourceName", o.ResourceName); err != nil {
		return err
	}

	// path param resourceType
	if err := r.SetPathParam("resourceType", o.ResourceType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
