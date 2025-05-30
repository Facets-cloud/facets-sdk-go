// Code generated by go-swagger; DO NOT EDIT.

package ui_application_controller

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

// NewListPodsUsingGETParams creates a new ListPodsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPodsUsingGETParams() *ListPodsUsingGETParams {
	return &ListPodsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPodsUsingGETParamsWithTimeout creates a new ListPodsUsingGETParams object
// with the ability to set a timeout on a request.
func NewListPodsUsingGETParamsWithTimeout(timeout time.Duration) *ListPodsUsingGETParams {
	return &ListPodsUsingGETParams{
		timeout: timeout,
	}
}

// NewListPodsUsingGETParamsWithContext creates a new ListPodsUsingGETParams object
// with the ability to set a context for a request.
func NewListPodsUsingGETParamsWithContext(ctx context.Context) *ListPodsUsingGETParams {
	return &ListPodsUsingGETParams{
		Context: ctx,
	}
}

// NewListPodsUsingGETParamsWithHTTPClient creates a new ListPodsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPodsUsingGETParamsWithHTTPClient(client *http.Client) *ListPodsUsingGETParams {
	return &ListPodsUsingGETParams{
		HTTPClient: client,
	}
}

/*
ListPodsUsingGETParams contains all the parameters to send to the API endpoint

	for the list pods using g e t operation.

	Typically these are written to a http.Request.
*/
type ListPodsUsingGETParams struct {

	/* ApplicationName.

	   applicationName
	*/
	ApplicationName string

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list pods using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPodsUsingGETParams) WithDefaults() *ListPodsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list pods using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPodsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list pods using g e t params
func (o *ListPodsUsingGETParams) WithTimeout(timeout time.Duration) *ListPodsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list pods using g e t params
func (o *ListPodsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list pods using g e t params
func (o *ListPodsUsingGETParams) WithContext(ctx context.Context) *ListPodsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list pods using g e t params
func (o *ListPodsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list pods using g e t params
func (o *ListPodsUsingGETParams) WithHTTPClient(client *http.Client) *ListPodsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list pods using g e t params
func (o *ListPodsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationName adds the applicationName to the list pods using g e t params
func (o *ListPodsUsingGETParams) WithApplicationName(applicationName string) *ListPodsUsingGETParams {
	o.SetApplicationName(applicationName)
	return o
}

// SetApplicationName adds the applicationName to the list pods using g e t params
func (o *ListPodsUsingGETParams) SetApplicationName(applicationName string) {
	o.ApplicationName = applicationName
}

// WithClusterID adds the clusterID to the list pods using g e t params
func (o *ListPodsUsingGETParams) WithClusterID(clusterID string) *ListPodsUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list pods using g e t params
func (o *ListPodsUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPodsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationName
	if err := r.SetPathParam("applicationName", o.ApplicationName); err != nil {
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
