// Code generated by go-swagger; DO NOT EDIT.

package ui_deployment_controller

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

// NewGetClusterStateUsingGETParams creates a new GetClusterStateUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterStateUsingGETParams() *GetClusterStateUsingGETParams {
	return &GetClusterStateUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterStateUsingGETParamsWithTimeout creates a new GetClusterStateUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetClusterStateUsingGETParamsWithTimeout(timeout time.Duration) *GetClusterStateUsingGETParams {
	return &GetClusterStateUsingGETParams{
		timeout: timeout,
	}
}

// NewGetClusterStateUsingGETParamsWithContext creates a new GetClusterStateUsingGETParams object
// with the ability to set a context for a request.
func NewGetClusterStateUsingGETParamsWithContext(ctx context.Context) *GetClusterStateUsingGETParams {
	return &GetClusterStateUsingGETParams{
		Context: ctx,
	}
}

// NewGetClusterStateUsingGETParamsWithHTTPClient creates a new GetClusterStateUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterStateUsingGETParamsWithHTTPClient(client *http.Client) *GetClusterStateUsingGETParams {
	return &GetClusterStateUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetClusterStateUsingGETParams contains all the parameters to send to the API endpoint

	for the get cluster state using g e t operation.

	Typically these are written to a http.Request.
*/
type GetClusterStateUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster state using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterStateUsingGETParams) WithDefaults() *GetClusterStateUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster state using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterStateUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster state using g e t params
func (o *GetClusterStateUsingGETParams) WithTimeout(timeout time.Duration) *GetClusterStateUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster state using g e t params
func (o *GetClusterStateUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster state using g e t params
func (o *GetClusterStateUsingGETParams) WithContext(ctx context.Context) *GetClusterStateUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster state using g e t params
func (o *GetClusterStateUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster state using g e t params
func (o *GetClusterStateUsingGETParams) WithHTTPClient(client *http.Client) *GetClusterStateUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster state using g e t params
func (o *GetClusterStateUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster state using g e t params
func (o *GetClusterStateUsingGETParams) WithClusterID(clusterID string) *GetClusterStateUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster state using g e t params
func (o *GetClusterStateUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterStateUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
