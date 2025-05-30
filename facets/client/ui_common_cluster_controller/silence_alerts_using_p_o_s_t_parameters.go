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

// NewSilenceAlertsUsingPOSTParams creates a new SilenceAlertsUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSilenceAlertsUsingPOSTParams() *SilenceAlertsUsingPOSTParams {
	return &SilenceAlertsUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSilenceAlertsUsingPOSTParamsWithTimeout creates a new SilenceAlertsUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewSilenceAlertsUsingPOSTParamsWithTimeout(timeout time.Duration) *SilenceAlertsUsingPOSTParams {
	return &SilenceAlertsUsingPOSTParams{
		timeout: timeout,
	}
}

// NewSilenceAlertsUsingPOSTParamsWithContext creates a new SilenceAlertsUsingPOSTParams object
// with the ability to set a context for a request.
func NewSilenceAlertsUsingPOSTParamsWithContext(ctx context.Context) *SilenceAlertsUsingPOSTParams {
	return &SilenceAlertsUsingPOSTParams{
		Context: ctx,
	}
}

// NewSilenceAlertsUsingPOSTParamsWithHTTPClient creates a new SilenceAlertsUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewSilenceAlertsUsingPOSTParamsWithHTTPClient(client *http.Client) *SilenceAlertsUsingPOSTParams {
	return &SilenceAlertsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
SilenceAlertsUsingPOSTParams contains all the parameters to send to the API endpoint

	for the silence alerts using p o s t operation.

	Typically these are written to a http.Request.
*/
type SilenceAlertsUsingPOSTParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* Request.

	   request
	*/
	Request *models.SilenceAlarmRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the silence alerts using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SilenceAlertsUsingPOSTParams) WithDefaults() *SilenceAlertsUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the silence alerts using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SilenceAlertsUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) WithTimeout(timeout time.Duration) *SilenceAlertsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) WithContext(ctx context.Context) *SilenceAlertsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) WithHTTPClient(client *http.Client) *SilenceAlertsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) WithClusterID(clusterID string) *SilenceAlertsUsingPOSTParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithRequest adds the request to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) WithRequest(request *models.SilenceAlarmRequest) *SilenceAlertsUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the silence alerts using p o s t params
func (o *SilenceAlertsUsingPOSTParams) SetRequest(request *models.SilenceAlarmRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *SilenceAlertsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
