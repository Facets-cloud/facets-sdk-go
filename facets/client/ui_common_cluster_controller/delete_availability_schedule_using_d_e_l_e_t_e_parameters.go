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

// NewDeleteAvailabilityScheduleUsingDELETEParams creates a new DeleteAvailabilityScheduleUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAvailabilityScheduleUsingDELETEParams() *DeleteAvailabilityScheduleUsingDELETEParams {
	return &DeleteAvailabilityScheduleUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAvailabilityScheduleUsingDELETEParamsWithTimeout creates a new DeleteAvailabilityScheduleUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteAvailabilityScheduleUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteAvailabilityScheduleUsingDELETEParams {
	return &DeleteAvailabilityScheduleUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteAvailabilityScheduleUsingDELETEParamsWithContext creates a new DeleteAvailabilityScheduleUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteAvailabilityScheduleUsingDELETEParamsWithContext(ctx context.Context) *DeleteAvailabilityScheduleUsingDELETEParams {
	return &DeleteAvailabilityScheduleUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteAvailabilityScheduleUsingDELETEParamsWithHTTPClient creates a new DeleteAvailabilityScheduleUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAvailabilityScheduleUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteAvailabilityScheduleUsingDELETEParams {
	return &DeleteAvailabilityScheduleUsingDELETEParams{
		HTTPClient: client,
	}
}

/*
DeleteAvailabilityScheduleUsingDELETEParams contains all the parameters to send to the API endpoint

	for the delete availability schedule using d e l e t e operation.

	Typically these are written to a http.Request.
*/
type DeleteAvailabilityScheduleUsingDELETEParams struct {

	/* AvailabilityScheduleID.

	   availabilityScheduleId
	*/
	AvailabilityScheduleID string

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete availability schedule using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAvailabilityScheduleUsingDELETEParams) WithDefaults() *DeleteAvailabilityScheduleUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete availability schedule using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAvailabilityScheduleUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteAvailabilityScheduleUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) WithContext(ctx context.Context) *DeleteAvailabilityScheduleUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteAvailabilityScheduleUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityScheduleID adds the availabilityScheduleID to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) WithAvailabilityScheduleID(availabilityScheduleID string) *DeleteAvailabilityScheduleUsingDELETEParams {
	o.SetAvailabilityScheduleID(availabilityScheduleID)
	return o
}

// SetAvailabilityScheduleID adds the availabilityScheduleId to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) SetAvailabilityScheduleID(availabilityScheduleID string) {
	o.AvailabilityScheduleID = availabilityScheduleID
}

// WithClusterID adds the clusterID to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) WithClusterID(clusterID string) *DeleteAvailabilityScheduleUsingDELETEParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete availability schedule using d e l e t e params
func (o *DeleteAvailabilityScheduleUsingDELETEParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAvailabilityScheduleUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param availabilityScheduleId
	if err := r.SetPathParam("availabilityScheduleId", o.AvailabilityScheduleID); err != nil {
		return err
	}

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
