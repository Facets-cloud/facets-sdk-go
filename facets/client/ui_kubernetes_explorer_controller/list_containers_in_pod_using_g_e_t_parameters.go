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

// NewListContainersInPodUsingGETParams creates a new ListContainersInPodUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListContainersInPodUsingGETParams() *ListContainersInPodUsingGETParams {
	return &ListContainersInPodUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListContainersInPodUsingGETParamsWithTimeout creates a new ListContainersInPodUsingGETParams object
// with the ability to set a timeout on a request.
func NewListContainersInPodUsingGETParamsWithTimeout(timeout time.Duration) *ListContainersInPodUsingGETParams {
	return &ListContainersInPodUsingGETParams{
		timeout: timeout,
	}
}

// NewListContainersInPodUsingGETParamsWithContext creates a new ListContainersInPodUsingGETParams object
// with the ability to set a context for a request.
func NewListContainersInPodUsingGETParamsWithContext(ctx context.Context) *ListContainersInPodUsingGETParams {
	return &ListContainersInPodUsingGETParams{
		Context: ctx,
	}
}

// NewListContainersInPodUsingGETParamsWithHTTPClient creates a new ListContainersInPodUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewListContainersInPodUsingGETParamsWithHTTPClient(client *http.Client) *ListContainersInPodUsingGETParams {
	return &ListContainersInPodUsingGETParams{
		HTTPClient: client,
	}
}

/*
ListContainersInPodUsingGETParams contains all the parameters to send to the API endpoint

	for the list containers in pod using g e t operation.

	Typically these are written to a http.Request.
*/
type ListContainersInPodUsingGETParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* Labels.

	   labels
	*/
	Labels []string

	/* PodName.

	   podName
	*/
	PodName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list containers in pod using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListContainersInPodUsingGETParams) WithDefaults() *ListContainersInPodUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list containers in pod using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListContainersInPodUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) WithTimeout(timeout time.Duration) *ListContainersInPodUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) WithContext(ctx context.Context) *ListContainersInPodUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) WithHTTPClient(client *http.Client) *ListContainersInPodUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) WithClusterID(clusterID string) *ListContainersInPodUsingGETParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLabels adds the labels to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) WithLabels(labels []string) *ListContainersInPodUsingGETParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) SetLabels(labels []string) {
	o.Labels = labels
}

// WithPodName adds the podName to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) WithPodName(podName string) *ListContainersInPodUsingGETParams {
	o.SetPodName(podName)
	return o
}

// SetPodName adds the podName to the list containers in pod using g e t params
func (o *ListContainersInPodUsingGETParams) SetPodName(podName string) {
	o.PodName = podName
}

// WriteToRequest writes these params to a swagger request
func (o *ListContainersInPodUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param podName
	if err := r.SetPathParam("podName", o.PodName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListContainersInPodUsingGET binds the parameter labels
func (o *ListContainersInPodUsingGETParams) bindParamLabels(formats strfmt.Registry) []string {
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
