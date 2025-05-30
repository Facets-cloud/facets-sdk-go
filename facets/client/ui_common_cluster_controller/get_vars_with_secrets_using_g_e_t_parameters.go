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

// NewGetVarsWithSecretsUsingGETParams creates a new GetVarsWithSecretsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVarsWithSecretsUsingGETParams() *GetVarsWithSecretsUsingGETParams {
	return &GetVarsWithSecretsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVarsWithSecretsUsingGETParamsWithTimeout creates a new GetVarsWithSecretsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetVarsWithSecretsUsingGETParamsWithTimeout(timeout time.Duration) *GetVarsWithSecretsUsingGETParams {
	return &GetVarsWithSecretsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetVarsWithSecretsUsingGETParamsWithContext creates a new GetVarsWithSecretsUsingGETParams object
// with the ability to set a context for a request.
func NewGetVarsWithSecretsUsingGETParamsWithContext(ctx context.Context) *GetVarsWithSecretsUsingGETParams {
	return &GetVarsWithSecretsUsingGETParams{
		Context: ctx,
	}
}

// NewGetVarsWithSecretsUsingGETParamsWithHTTPClient creates a new GetVarsWithSecretsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVarsWithSecretsUsingGETParamsWithHTTPClient(client *http.Client) *GetVarsWithSecretsUsingGETParams {
	return &GetVarsWithSecretsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetVarsWithSecretsUsingGETParams contains all the parameters to send to the API endpoint

	for the get vars with secrets using g e t operation.

	Typically these are written to a http.Request.
*/
type GetVarsWithSecretsUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vars with secrets using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVarsWithSecretsUsingGETParams) WithDefaults() *GetVarsWithSecretsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vars with secrets using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVarsWithSecretsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vars with secrets using g e t params
func (o *GetVarsWithSecretsUsingGETParams) WithTimeout(timeout time.Duration) *GetVarsWithSecretsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vars with secrets using g e t params
func (o *GetVarsWithSecretsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vars with secrets using g e t params
func (o *GetVarsWithSecretsUsingGETParams) WithContext(ctx context.Context) *GetVarsWithSecretsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vars with secrets using g e t params
func (o *GetVarsWithSecretsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vars with secrets using g e t params
func (o *GetVarsWithSecretsUsingGETParams) WithHTTPClient(client *http.Client) *GetVarsWithSecretsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vars with secrets using g e t params
func (o *GetVarsWithSecretsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get vars with secrets using g e t params
func (o *GetVarsWithSecretsUsingGETParams) WithClusterID(clusterID string) *GetVarsWithSecretsUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get vars with secrets using g e t params
func (o *GetVarsWithSecretsUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVarsWithSecretsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
