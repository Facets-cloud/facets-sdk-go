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

	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

// NewPauseReleaseUsingPOSTParams creates a new PauseReleaseUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPauseReleaseUsingPOSTParams() *PauseReleaseUsingPOSTParams {
	return &PauseReleaseUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPauseReleaseUsingPOSTParamsWithTimeout creates a new PauseReleaseUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewPauseReleaseUsingPOSTParamsWithTimeout(timeout time.Duration) *PauseReleaseUsingPOSTParams {
	return &PauseReleaseUsingPOSTParams{
		timeout: timeout,
	}
}

// NewPauseReleaseUsingPOSTParamsWithContext creates a new PauseReleaseUsingPOSTParams object
// with the ability to set a context for a request.
func NewPauseReleaseUsingPOSTParamsWithContext(ctx context.Context) *PauseReleaseUsingPOSTParams {
	return &PauseReleaseUsingPOSTParams{
		Context: ctx,
	}
}

// NewPauseReleaseUsingPOSTParamsWithHTTPClient creates a new PauseReleaseUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewPauseReleaseUsingPOSTParamsWithHTTPClient(client *http.Client) *PauseReleaseUsingPOSTParams {
	return &PauseReleaseUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
PauseReleaseUsingPOSTParams contains all the parameters to send to the API endpoint

	for the pause release using p o s t operation.

	Typically these are written to a http.Request.
*/
type PauseReleaseUsingPOSTParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* PauseReleaseRequest.

	   pauseReleaseRequest
	*/
	PauseReleaseRequest *models.PauseReleaseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pause release using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseReleaseUsingPOSTParams) WithDefaults() *PauseReleaseUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pause release using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseReleaseUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) WithTimeout(timeout time.Duration) *PauseReleaseUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) WithContext(ctx context.Context) *PauseReleaseUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) WithHTTPClient(client *http.Client) *PauseReleaseUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) WithClusterID(clusterID string) *PauseReleaseUsingPOSTParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithPauseReleaseRequest adds the pauseReleaseRequest to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) WithPauseReleaseRequest(pauseReleaseRequest *models.PauseReleaseRequest) *PauseReleaseUsingPOSTParams {
	o.SetPauseReleaseRequest(pauseReleaseRequest)
	return o
}

// SetPauseReleaseRequest adds the pauseReleaseRequest to the pause release using p o s t params
func (o *PauseReleaseUsingPOSTParams) SetPauseReleaseRequest(pauseReleaseRequest *models.PauseReleaseRequest) {
	o.PauseReleaseRequest = pauseReleaseRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PauseReleaseUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}
	if o.PauseReleaseRequest != nil {
		if err := r.SetBodyParam(o.PauseReleaseRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
