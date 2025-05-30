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

// NewGetClusterInfoUsingGETParams creates a new GetClusterInfoUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterInfoUsingGETParams() *GetClusterInfoUsingGETParams {
	return &GetClusterInfoUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterInfoUsingGETParamsWithTimeout creates a new GetClusterInfoUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetClusterInfoUsingGETParamsWithTimeout(timeout time.Duration) *GetClusterInfoUsingGETParams {
	return &GetClusterInfoUsingGETParams{
		timeout: timeout,
	}
}

// NewGetClusterInfoUsingGETParamsWithContext creates a new GetClusterInfoUsingGETParams object
// with the ability to set a context for a request.
func NewGetClusterInfoUsingGETParamsWithContext(ctx context.Context) *GetClusterInfoUsingGETParams {
	return &GetClusterInfoUsingGETParams{
		Context: ctx,
	}
}

// NewGetClusterInfoUsingGETParamsWithHTTPClient creates a new GetClusterInfoUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterInfoUsingGETParamsWithHTTPClient(client *http.Client) *GetClusterInfoUsingGETParams {
	return &GetClusterInfoUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetClusterInfoUsingGETParams contains all the parameters to send to the API endpoint

	for the get cluster info using g e t operation.

	Typically these are written to a http.Request.
*/
type GetClusterInfoUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster info using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterInfoUsingGETParams) WithDefaults() *GetClusterInfoUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster info using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterInfoUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster info using g e t params
func (o *GetClusterInfoUsingGETParams) WithTimeout(timeout time.Duration) *GetClusterInfoUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster info using g e t params
func (o *GetClusterInfoUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster info using g e t params
func (o *GetClusterInfoUsingGETParams) WithContext(ctx context.Context) *GetClusterInfoUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster info using g e t params
func (o *GetClusterInfoUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster info using g e t params
func (o *GetClusterInfoUsingGETParams) WithHTTPClient(client *http.Client) *GetClusterInfoUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster info using g e t params
func (o *GetClusterInfoUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster info using g e t params
func (o *GetClusterInfoUsingGETParams) WithClusterID(clusterID string) *GetClusterInfoUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster info using g e t params
func (o *GetClusterInfoUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterInfoUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
