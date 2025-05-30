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

// NewListIngressesByLabelsUsingGETParams creates a new ListIngressesByLabelsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListIngressesByLabelsUsingGETParams() *ListIngressesByLabelsUsingGETParams {
	return &ListIngressesByLabelsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListIngressesByLabelsUsingGETParamsWithTimeout creates a new ListIngressesByLabelsUsingGETParams object
// with the ability to set a timeout on a request.
func NewListIngressesByLabelsUsingGETParamsWithTimeout(timeout time.Duration) *ListIngressesByLabelsUsingGETParams {
	return &ListIngressesByLabelsUsingGETParams{
		timeout: timeout,
	}
}

// NewListIngressesByLabelsUsingGETParamsWithContext creates a new ListIngressesByLabelsUsingGETParams object
// with the ability to set a context for a request.
func NewListIngressesByLabelsUsingGETParamsWithContext(ctx context.Context) *ListIngressesByLabelsUsingGETParams {
	return &ListIngressesByLabelsUsingGETParams{
		Context: ctx,
	}
}

// NewListIngressesByLabelsUsingGETParamsWithHTTPClient creates a new ListIngressesByLabelsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewListIngressesByLabelsUsingGETParamsWithHTTPClient(client *http.Client) *ListIngressesByLabelsUsingGETParams {
	return &ListIngressesByLabelsUsingGETParams{
		HTTPClient: client,
	}
}

/*
ListIngressesByLabelsUsingGETParams contains all the parameters to send to the API endpoint

	for the list ingresses by labels using g e t operation.

	Typically these are written to a http.Request.
*/
type ListIngressesByLabelsUsingGETParams struct {

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

// WithDefaults hydrates default values in the list ingresses by labels using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListIngressesByLabelsUsingGETParams) WithDefaults() *ListIngressesByLabelsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list ingresses by labels using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListIngressesByLabelsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) WithTimeout(timeout time.Duration) *ListIngressesByLabelsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) WithContext(ctx context.Context) *ListIngressesByLabelsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) WithHTTPClient(client *http.Client) *ListIngressesByLabelsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) WithClusterID(clusterID string) *ListIngressesByLabelsUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLabels adds the labels to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) WithLabels(labels []string) *ListIngressesByLabelsUsingGETParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the list ingresses by labels using g e t params
func (o *ListIngressesByLabelsUsingGETParams) SetLabels(labels []string) {
	o.Labels = labels
}

// WriteToRequest writes these params to a swagger request
func (o *ListIngressesByLabelsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamListIngressesByLabelsUsingGET binds the parameter labels
func (o *ListIngressesByLabelsUsingGETParams) bindParamLabels(formats strfmt.Registry) []string {
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
