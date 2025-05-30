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

// NewGetClusterMetadataUsingGETParams creates a new GetClusterMetadataUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterMetadataUsingGETParams() *GetClusterMetadataUsingGETParams {
	return &GetClusterMetadataUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterMetadataUsingGETParamsWithTimeout creates a new GetClusterMetadataUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetClusterMetadataUsingGETParamsWithTimeout(timeout time.Duration) *GetClusterMetadataUsingGETParams {
	return &GetClusterMetadataUsingGETParams{
		timeout: timeout,
	}
}

// NewGetClusterMetadataUsingGETParamsWithContext creates a new GetClusterMetadataUsingGETParams object
// with the ability to set a context for a request.
func NewGetClusterMetadataUsingGETParamsWithContext(ctx context.Context) *GetClusterMetadataUsingGETParams {
	return &GetClusterMetadataUsingGETParams{
		Context: ctx,
	}
}

// NewGetClusterMetadataUsingGETParamsWithHTTPClient creates a new GetClusterMetadataUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterMetadataUsingGETParamsWithHTTPClient(client *http.Client) *GetClusterMetadataUsingGETParams {
	return &GetClusterMetadataUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetClusterMetadataUsingGETParams contains all the parameters to send to the API endpoint

	for the get cluster metadata using g e t operation.

	Typically these are written to a http.Request.
*/
type GetClusterMetadataUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster metadata using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterMetadataUsingGETParams) WithDefaults() *GetClusterMetadataUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster metadata using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterMetadataUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster metadata using g e t params
func (o *GetClusterMetadataUsingGETParams) WithTimeout(timeout time.Duration) *GetClusterMetadataUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster metadata using g e t params
func (o *GetClusterMetadataUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster metadata using g e t params
func (o *GetClusterMetadataUsingGETParams) WithContext(ctx context.Context) *GetClusterMetadataUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster metadata using g e t params
func (o *GetClusterMetadataUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster metadata using g e t params
func (o *GetClusterMetadataUsingGETParams) WithHTTPClient(client *http.Client) *GetClusterMetadataUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster metadata using g e t params
func (o *GetClusterMetadataUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster metadata using g e t params
func (o *GetClusterMetadataUsingGETParams) WithClusterID(clusterID string) *GetClusterMetadataUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster metadata using g e t params
func (o *GetClusterMetadataUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterMetadataUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
