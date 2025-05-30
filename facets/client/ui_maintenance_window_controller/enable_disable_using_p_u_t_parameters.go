// Code generated by go-swagger; DO NOT EDIT.

package ui_maintenance_window_controller

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
	"github.com/go-openapi/swag"
)

// NewEnableDisableUsingPUTParams creates a new EnableDisableUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableDisableUsingPUTParams() *EnableDisableUsingPUTParams {
	return &EnableDisableUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableDisableUsingPUTParamsWithTimeout creates a new EnableDisableUsingPUTParams object
// with the ability to set a timeout on a request.
func NewEnableDisableUsingPUTParamsWithTimeout(timeout time.Duration) *EnableDisableUsingPUTParams {
	return &EnableDisableUsingPUTParams{
		timeout: timeout,
	}
}

// NewEnableDisableUsingPUTParamsWithContext creates a new EnableDisableUsingPUTParams object
// with the ability to set a context for a request.
func NewEnableDisableUsingPUTParamsWithContext(ctx context.Context) *EnableDisableUsingPUTParams {
	return &EnableDisableUsingPUTParams{
		Context: ctx,
	}
}

// NewEnableDisableUsingPUTParamsWithHTTPClient creates a new EnableDisableUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableDisableUsingPUTParamsWithHTTPClient(client *http.Client) *EnableDisableUsingPUTParams {
	return &EnableDisableUsingPUTParams{
		HTTPClient: client,
	}
}

/*
EnableDisableUsingPUTParams contains all the parameters to send to the API endpoint

	for the enable disable using p u t operation.

	Typically these are written to a http.Request.
*/
type EnableDisableUsingPUTParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* Disabled.

	   disabled
	*/
	Disabled bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable disable using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableDisableUsingPUTParams) WithDefaults() *EnableDisableUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable disable using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableDisableUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) WithTimeout(timeout time.Duration) *EnableDisableUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) WithContext(ctx context.Context) *EnableDisableUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) WithHTTPClient(client *http.Client) *EnableDisableUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) WithClusterID(clusterID string) *EnableDisableUsingPUTParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDisabled adds the disabled to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) WithDisabled(disabled bool) *EnableDisableUsingPUTParams {
	o.SetDisabled(disabled)
	return o
}

// SetDisabled adds the disabled to the enable disable using p u t params
func (o *EnableDisableUsingPUTParams) SetDisabled(disabled bool) {
	o.Disabled = disabled
}

// WriteToRequest writes these params to a swagger request
func (o *EnableDisableUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	// query param disabled
	qrDisabled := o.Disabled
	qDisabled := swag.FormatBool(qrDisabled)
	if qDisabled != "" {

		if err := r.SetQueryParam("disabled", qDisabled); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
