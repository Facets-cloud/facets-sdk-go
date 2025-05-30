// Code generated by go-swagger; DO NOT EDIT.

package ui_common_cluster_controller

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

// NewDeleteClusterForceUsingDELETEParams creates a new DeleteClusterForceUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteClusterForceUsingDELETEParams() *DeleteClusterForceUsingDELETEParams {
	return &DeleteClusterForceUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterForceUsingDELETEParamsWithTimeout creates a new DeleteClusterForceUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteClusterForceUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteClusterForceUsingDELETEParams {
	return &DeleteClusterForceUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteClusterForceUsingDELETEParamsWithContext creates a new DeleteClusterForceUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteClusterForceUsingDELETEParamsWithContext(ctx context.Context) *DeleteClusterForceUsingDELETEParams {
	return &DeleteClusterForceUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteClusterForceUsingDELETEParamsWithHTTPClient creates a new DeleteClusterForceUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteClusterForceUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteClusterForceUsingDELETEParams {
	return &DeleteClusterForceUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteClusterForceUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete cluster force using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteClusterForceUsingDELETEParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cluster force using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClusterForceUsingDELETEParams) WithDefaults() *DeleteClusterForceUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cluster force using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClusterForceUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete cluster force using d e l e t e params
func (o *DeleteClusterForceUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteClusterForceUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster force using d e l e t e params
func (o *DeleteClusterForceUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster force using d e l e t e params
func (o *DeleteClusterForceUsingDELETEParams) WithContext(ctx context.Context) *DeleteClusterForceUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster force using d e l e t e params
func (o *DeleteClusterForceUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster force using d e l e t e params
func (o *DeleteClusterForceUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteClusterForceUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster force using d e l e t e params
func (o *DeleteClusterForceUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete cluster force using d e l e t e params
func (o *DeleteClusterForceUsingDELETEParams) WithClusterID(clusterID string) *DeleteClusterForceUsingDELETEParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster force using d e l e t e params
func (o *DeleteClusterForceUsingDELETEParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterForceUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
