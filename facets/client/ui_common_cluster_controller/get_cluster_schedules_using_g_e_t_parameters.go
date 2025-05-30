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

// NewGetClusterSchedulesUsingGETParams creates a new GetClusterSchedulesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterSchedulesUsingGETParams() *GetClusterSchedulesUsingGETParams {
	return &GetClusterSchedulesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterSchedulesUsingGETParamsWithTimeout creates a new GetClusterSchedulesUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetClusterSchedulesUsingGETParamsWithTimeout(timeout time.Duration) *GetClusterSchedulesUsingGETParams {
	return &GetClusterSchedulesUsingGETParams{
		timeout: timeout,
	}
}

// NewGetClusterSchedulesUsingGETParamsWithContext creates a new GetClusterSchedulesUsingGETParams object
// with the ability to set a context for a request.
func NewGetClusterSchedulesUsingGETParamsWithContext(ctx context.Context) *GetClusterSchedulesUsingGETParams {
	return &GetClusterSchedulesUsingGETParams{
		Context: ctx,
	}
}

// NewGetClusterSchedulesUsingGETParamsWithHTTPClient creates a new GetClusterSchedulesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterSchedulesUsingGETParamsWithHTTPClient(client *http.Client) *GetClusterSchedulesUsingGETParams {
	return &GetClusterSchedulesUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetClusterSchedulesUsingGETParams contains all the parameters to send to the API endpoint

	for the get cluster schedules using g e t operation.

	Typically these are written to a http.Request.
*/
type GetClusterSchedulesUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster schedules using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterSchedulesUsingGETParams) WithDefaults() *GetClusterSchedulesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster schedules using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterSchedulesUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster schedules using g e t params
func (o *GetClusterSchedulesUsingGETParams) WithTimeout(timeout time.Duration) *GetClusterSchedulesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster schedules using g e t params
func (o *GetClusterSchedulesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster schedules using g e t params
func (o *GetClusterSchedulesUsingGETParams) WithContext(ctx context.Context) *GetClusterSchedulesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster schedules using g e t params
func (o *GetClusterSchedulesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster schedules using g e t params
func (o *GetClusterSchedulesUsingGETParams) WithHTTPClient(client *http.Client) *GetClusterSchedulesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster schedules using g e t params
func (o *GetClusterSchedulesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster schedules using g e t params
func (o *GetClusterSchedulesUsingGETParams) WithClusterID(clusterID string) *GetClusterSchedulesUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster schedules using g e t params
func (o *GetClusterSchedulesUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterSchedulesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
