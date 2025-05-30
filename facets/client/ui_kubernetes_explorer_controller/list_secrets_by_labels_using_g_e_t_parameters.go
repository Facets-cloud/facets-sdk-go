// Code generated by go-swagger; DO NOT EDIT.

package ui_kubernetes_explorer_controller

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

// NewListSecretsByLabelsUsingGETParams creates a new ListSecretsByLabelsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSecretsByLabelsUsingGETParams() *ListSecretsByLabelsUsingGETParams {
	return &ListSecretsByLabelsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSecretsByLabelsUsingGETParamsWithTimeout creates a new ListSecretsByLabelsUsingGETParams object
// with the ability to set a timeout on a request.
func NewListSecretsByLabelsUsingGETParamsWithTimeout(timeout time.Duration) *ListSecretsByLabelsUsingGETParams {
	return &ListSecretsByLabelsUsingGETParams{
		timeout: timeout,
	}
}

// NewListSecretsByLabelsUsingGETParamsWithContext creates a new ListSecretsByLabelsUsingGETParams object
// with the ability to set a context for a request.
func NewListSecretsByLabelsUsingGETParamsWithContext(ctx context.Context) *ListSecretsByLabelsUsingGETParams {
	return &ListSecretsByLabelsUsingGETParams{
		Context: ctx,
	}
}

// NewListSecretsByLabelsUsingGETParamsWithHTTPClient creates a new ListSecretsByLabelsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSecretsByLabelsUsingGETParamsWithHTTPClient(client *http.Client) *ListSecretsByLabelsUsingGETParams {
	return &ListSecretsByLabelsUsingGETParams{
		HTTPClient: client,
	}
}

/*
ListSecretsByLabelsUsingGETParams contains all the parameters to send to the API endpoint

	for the list secrets by labels using g e t operation.

	Typically these are written to a http.Request.
*/
type ListSecretsByLabelsUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* Labels.

	   labels
	*/
	Labels []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list secrets by labels using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSecretsByLabelsUsingGETParams) WithDefaults() *ListSecretsByLabelsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list secrets by labels using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSecretsByLabelsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) WithTimeout(timeout time.Duration) *ListSecretsByLabelsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) WithContext(ctx context.Context) *ListSecretsByLabelsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) WithHTTPClient(client *http.Client) *ListSecretsByLabelsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) WithClusterID(clusterID string) *ListSecretsByLabelsUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLabels adds the labels to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) WithLabels(labels []string) *ListSecretsByLabelsUsingGETParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the list secrets by labels using g e t params
func (o *ListSecretsByLabelsUsingGETParams) SetLabels(labels []string) {
	o.Labels = labels
}

// WriteToRequest writes these params to a swagger request
func (o *ListSecretsByLabelsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	if o.Labels != nil {

		// binding items for labels
		joinedLabels := o.bindParamLabels(reg)

		// query array param labels
		if err := r.SetQueryParam("labels", joinedLabels...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListSecretsByLabelsUsingGET binds the parameter labels
func (o *ListSecretsByLabelsUsingGETParams) bindParamLabels(formats strfmt.Registry) []string {
	labelsIR := o.Labels

	var labelsIC []string
	for _, labelsIIR := range labelsIR { // explode []string

		labelsIIV := labelsIIR // string as string
		labelsIC = append(labelsIC, labelsIIV)
	}

	// items.CollectionFormat: ""
	labelsIS := swag.JoinByFormat(labelsIC, "")

	return labelsIS
}
