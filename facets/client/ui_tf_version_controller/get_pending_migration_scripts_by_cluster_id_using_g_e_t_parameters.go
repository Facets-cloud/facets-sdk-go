// Code generated by go-swagger; DO NOT EDIT.

package ui_tf_version_controller

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

// NewGetPendingMigrationScriptsByClusterIDUsingGETParams creates a new GetPendingMigrationScriptsByClusterIDUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPendingMigrationScriptsByClusterIDUsingGETParams() *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	return &GetPendingMigrationScriptsByClusterIDUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPendingMigrationScriptsByClusterIDUsingGETParamsWithTimeout creates a new GetPendingMigrationScriptsByClusterIDUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetPendingMigrationScriptsByClusterIDUsingGETParamsWithTimeout(timeout time.Duration) *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	return &GetPendingMigrationScriptsByClusterIDUsingGETParams{
		timeout: timeout,
	}
}

// NewGetPendingMigrationScriptsByClusterIDUsingGETParamsWithContext creates a new GetPendingMigrationScriptsByClusterIDUsingGETParams object
// with the ability to set a context for a request.
func NewGetPendingMigrationScriptsByClusterIDUsingGETParamsWithContext(ctx context.Context) *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	return &GetPendingMigrationScriptsByClusterIDUsingGETParams{
		Context: ctx,
	}
}

// NewGetPendingMigrationScriptsByClusterIDUsingGETParamsWithHTTPClient creates a new GetPendingMigrationScriptsByClusterIDUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPendingMigrationScriptsByClusterIDUsingGETParamsWithHTTPClient(client *http.Client) *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	return &GetPendingMigrationScriptsByClusterIDUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetPendingMigrationScriptsByClusterIDUsingGETParams contains all the parameters to send to the API endpoint

	for the get pending migration scripts by cluster Id using g e t operation.

	Typically these are written to a http.Request.
*/
type GetPendingMigrationScriptsByClusterIDUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pending migration scripts by cluster Id using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) WithDefaults() *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pending migration scripts by cluster Id using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pending migration scripts by cluster Id using g e t params
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) WithTimeout(timeout time.Duration) *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pending migration scripts by cluster Id using g e t params
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pending migration scripts by cluster Id using g e t params
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) WithContext(ctx context.Context) *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pending migration scripts by cluster Id using g e t params
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pending migration scripts by cluster Id using g e t params
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) WithHTTPClient(client *http.Client) *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pending migration scripts by cluster Id using g e t params
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get pending migration scripts by cluster Id using g e t params
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) WithClusterID(clusterID string) *GetPendingMigrationScriptsByClusterIDUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get pending migration scripts by cluster Id using g e t params
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPendingMigrationScriptsByClusterIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
