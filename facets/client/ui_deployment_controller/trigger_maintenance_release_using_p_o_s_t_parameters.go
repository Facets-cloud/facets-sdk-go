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

// NewTriggerMaintenanceReleaseUsingPOSTParams creates a new TriggerMaintenanceReleaseUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTriggerMaintenanceReleaseUsingPOSTParams() *TriggerMaintenanceReleaseUsingPOSTParams {
	return &TriggerMaintenanceReleaseUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerMaintenanceReleaseUsingPOSTParamsWithTimeout creates a new TriggerMaintenanceReleaseUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewTriggerMaintenanceReleaseUsingPOSTParamsWithTimeout(timeout time.Duration) *TriggerMaintenanceReleaseUsingPOSTParams {
	return &TriggerMaintenanceReleaseUsingPOSTParams{
		timeout: timeout,
	}
}

// NewTriggerMaintenanceReleaseUsingPOSTParamsWithContext creates a new TriggerMaintenanceReleaseUsingPOSTParams object
// with the ability to set a context for a request.
func NewTriggerMaintenanceReleaseUsingPOSTParamsWithContext(ctx context.Context) *TriggerMaintenanceReleaseUsingPOSTParams {
	return &TriggerMaintenanceReleaseUsingPOSTParams{
		Context: ctx,
	}
}

// NewTriggerMaintenanceReleaseUsingPOSTParamsWithHTTPClient creates a new TriggerMaintenanceReleaseUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewTriggerMaintenanceReleaseUsingPOSTParamsWithHTTPClient(client *http.Client) *TriggerMaintenanceReleaseUsingPOSTParams {
	return &TriggerMaintenanceReleaseUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
TriggerMaintenanceReleaseUsingPOSTParams contains all the parameters to send to the API endpoint

	for the trigger maintenance release using p o s t operation.

	Typically these are written to a http.Request.
*/
type TriggerMaintenanceReleaseUsingPOSTParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the trigger maintenance release using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerMaintenanceReleaseUsingPOSTParams) WithDefaults() *TriggerMaintenanceReleaseUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the trigger maintenance release using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TriggerMaintenanceReleaseUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the trigger maintenance release using p o s t params
func (o *TriggerMaintenanceReleaseUsingPOSTParams) WithTimeout(timeout time.Duration) *TriggerMaintenanceReleaseUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger maintenance release using p o s t params
func (o *TriggerMaintenanceReleaseUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger maintenance release using p o s t params
func (o *TriggerMaintenanceReleaseUsingPOSTParams) WithContext(ctx context.Context) *TriggerMaintenanceReleaseUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger maintenance release using p o s t params
func (o *TriggerMaintenanceReleaseUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger maintenance release using p o s t params
func (o *TriggerMaintenanceReleaseUsingPOSTParams) WithHTTPClient(client *http.Client) *TriggerMaintenanceReleaseUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger maintenance release using p o s t params
func (o *TriggerMaintenanceReleaseUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the trigger maintenance release using p o s t params
func (o *TriggerMaintenanceReleaseUsingPOSTParams) WithClusterID(clusterID string) *TriggerMaintenanceReleaseUsingPOSTParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the trigger maintenance release using p o s t params
func (o *TriggerMaintenanceReleaseUsingPOSTParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerMaintenanceReleaseUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
