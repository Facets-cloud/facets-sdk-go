// Code generated by go-swagger; DO NOT EDIT.

package ui_domain_mapping_controller

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

// NewDeleteDomainMappingUsingDELETEParams creates a new DeleteDomainMappingUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDomainMappingUsingDELETEParams() *DeleteDomainMappingUsingDELETEParams {
	return &DeleteDomainMappingUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDomainMappingUsingDELETEParamsWithTimeout creates a new DeleteDomainMappingUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteDomainMappingUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteDomainMappingUsingDELETEParams {
	return &DeleteDomainMappingUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteDomainMappingUsingDELETEParamsWithContext creates a new DeleteDomainMappingUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteDomainMappingUsingDELETEParamsWithContext(ctx context.Context) *DeleteDomainMappingUsingDELETEParams {
	return &DeleteDomainMappingUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteDomainMappingUsingDELETEParamsWithHTTPClient creates a new DeleteDomainMappingUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDomainMappingUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteDomainMappingUsingDELETEParams {
	return &DeleteDomainMappingUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteDomainMappingUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete domain mapping using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteDomainMappingUsingDELETEParams struct {

	/* Alias.

	   alias
	*/
	Alias string

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

// WithDefaults hydrates default values in the delete domain mapping using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDomainMappingUsingDELETEParams) WithDefaults() *DeleteDomainMappingUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete domain mapping using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDomainMappingUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteDomainMappingUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) WithContext(ctx context.Context) *DeleteDomainMappingUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteDomainMappingUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlias adds the alias to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) WithAlias(alias string) *DeleteDomainMappingUsingDELETEParams {
	o.SetAlias(alias)
	return o
}

// SetAlias adds the alias to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) SetAlias(alias string) {
	o.Alias = alias
}

// WithClusterID adds the clusterID to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) WithClusterID(clusterID string) *DeleteDomainMappingUsingDELETEParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithResourceName adds the resourceName to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) WithResourceName(resourceName string) *DeleteDomainMappingUsingDELETEParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WithResourceType adds the resourceType to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) WithResourceType(resourceType string) *DeleteDomainMappingUsingDELETEParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the delete domain mapping using d e l e t e params
func (o *DeleteDomainMappingUsingDELETEParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDomainMappingUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param alias
	qrAlias := o.Alias
	qAlias := qrAlias
	if qAlias != "" {

		if err := r.SetQueryParam("alias", qAlias); err != nil {
			return err
		}
	}

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
