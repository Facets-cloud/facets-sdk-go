// Code generated by go-swagger; DO NOT EDIT.

package common_cluster_controller

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

// NewNotifyAlertsUsingPOSTParams creates a new NotifyAlertsUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotifyAlertsUsingPOSTParams() *NotifyAlertsUsingPOSTParams {
	return &NotifyAlertsUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotifyAlertsUsingPOSTParamsWithTimeout creates a new NotifyAlertsUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewNotifyAlertsUsingPOSTParamsWithTimeout(timeout time.Duration) *NotifyAlertsUsingPOSTParams {
	return &NotifyAlertsUsingPOSTParams{
		timeout: timeout,
	}
}

// NewNotifyAlertsUsingPOSTParamsWithContext creates a new NotifyAlertsUsingPOSTParams object
// with the ability to set a context for a request.
func NewNotifyAlertsUsingPOSTParamsWithContext(ctx context.Context) *NotifyAlertsUsingPOSTParams {
	return &NotifyAlertsUsingPOSTParams{
		Context: ctx,
	}
}

// NewNotifyAlertsUsingPOSTParamsWithHTTPClient creates a new NotifyAlertsUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotifyAlertsUsingPOSTParamsWithHTTPClient(client *http.Client) *NotifyAlertsUsingPOSTParams {
	return &NotifyAlertsUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
NotifyAlertsUsingPOSTParams contains all the parameters to send to the API endpoint

	for the notify alerts using p o s t operation.

	Typically these are written to a http.Request.
*/
type NotifyAlertsUsingPOSTParams struct {

	/* Alerts.

	   alerts
	*/
	Alerts *models.Response

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notify alerts using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifyAlertsUsingPOSTParams) WithDefaults() *NotifyAlertsUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notify alerts using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifyAlertsUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) WithTimeout(timeout time.Duration) *NotifyAlertsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) WithContext(ctx context.Context) *NotifyAlertsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) WithHTTPClient(client *http.Client) *NotifyAlertsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlerts adds the alerts to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) WithAlerts(alerts *models.Response) *NotifyAlertsUsingPOSTParams {
	o.SetAlerts(alerts)
	return o
}

// SetAlerts adds the alerts to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) SetAlerts(alerts *models.Response) {
	o.Alerts = alerts
}

// WithClusterID adds the clusterID to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) WithClusterID(clusterID string) *NotifyAlertsUsingPOSTParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the notify alerts using p o s t params
func (o *NotifyAlertsUsingPOSTParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *NotifyAlertsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Alerts != nil {
		if err := r.SetBodyParam(o.Alerts); err != nil {
			return err
		}
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
