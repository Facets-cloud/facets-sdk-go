// Code generated by go-swagger; DO NOT EDIT.

package ui_chat_gpt_controller

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

// NewAnalyzeKubernetesClusterUsingGETParams creates a new AnalyzeKubernetesClusterUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAnalyzeKubernetesClusterUsingGETParams() *AnalyzeKubernetesClusterUsingGETParams {
	return &AnalyzeKubernetesClusterUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAnalyzeKubernetesClusterUsingGETParamsWithTimeout creates a new AnalyzeKubernetesClusterUsingGETParams object
// with the ability to set a timeout on a request.
func NewAnalyzeKubernetesClusterUsingGETParamsWithTimeout(timeout time.Duration) *AnalyzeKubernetesClusterUsingGETParams {
	return &AnalyzeKubernetesClusterUsingGETParams{
		timeout: timeout,
	}
}

// NewAnalyzeKubernetesClusterUsingGETParamsWithContext creates a new AnalyzeKubernetesClusterUsingGETParams object
// with the ability to set a context for a request.
func NewAnalyzeKubernetesClusterUsingGETParamsWithContext(ctx context.Context) *AnalyzeKubernetesClusterUsingGETParams {
	return &AnalyzeKubernetesClusterUsingGETParams{
		Context: ctx,
	}
}

// NewAnalyzeKubernetesClusterUsingGETParamsWithHTTPClient creates a new AnalyzeKubernetesClusterUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewAnalyzeKubernetesClusterUsingGETParamsWithHTTPClient(client *http.Client) *AnalyzeKubernetesClusterUsingGETParams {
	return &AnalyzeKubernetesClusterUsingGETParams{
		HTTPClient: client,
	}
}

/*
AnalyzeKubernetesClusterUsingGETParams contains all the parameters to send to the API endpoint

	for the analyze kubernetes cluster using g e t operation.

	Typically these are written to a http.Request.
*/
type AnalyzeKubernetesClusterUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* Filters.

	   filters
	*/
	Filters []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the analyze kubernetes cluster using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AnalyzeKubernetesClusterUsingGETParams) WithDefaults() *AnalyzeKubernetesClusterUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the analyze kubernetes cluster using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AnalyzeKubernetesClusterUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) WithTimeout(timeout time.Duration) *AnalyzeKubernetesClusterUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) WithContext(ctx context.Context) *AnalyzeKubernetesClusterUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) WithHTTPClient(client *http.Client) *AnalyzeKubernetesClusterUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) WithClusterID(clusterID string) *AnalyzeKubernetesClusterUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithFilters adds the filters to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) WithFilters(filters []string) *AnalyzeKubernetesClusterUsingGETParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the analyze kubernetes cluster using g e t params
func (o *AnalyzeKubernetesClusterUsingGETParams) SetFilters(filters []string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *AnalyzeKubernetesClusterUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if o.Filters != nil {

		// binding items for filters
		joinedFilters := o.bindParamFilters(reg)

		// query array param filters
		if err := r.SetQueryParam("filters", joinedFilters...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamAnalyzeKubernetesClusterUsingGET binds the parameter filters
func (o *AnalyzeKubernetesClusterUsingGETParams) bindParamFilters(formats strfmt.Registry) []string {
	filtersIR := o.Filters

	var filtersIC []string
	for _, filtersIIR := range filtersIR { // explode []string

		filtersIIV := filtersIIR // string as string
		filtersIC = append(filtersIC, filtersIIV)
	}

	// items.CollectionFormat: "multi"
	filtersIS := swag.JoinByFormat(filtersIC, "multi")

	return filtersIS
}
