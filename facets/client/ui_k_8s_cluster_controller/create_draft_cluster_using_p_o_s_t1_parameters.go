// Code generated by go-swagger; DO NOT EDIT.

package ui_k_8s_cluster_controller

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

// NewCreateDraftClusterUsingPOST1Params creates a new CreateDraftClusterUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDraftClusterUsingPOST1Params() *CreateDraftClusterUsingPOST1Params {
	return &CreateDraftClusterUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDraftClusterUsingPOST1ParamsWithTimeout creates a new CreateDraftClusterUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewCreateDraftClusterUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateDraftClusterUsingPOST1Params {
	return &CreateDraftClusterUsingPOST1Params{
		timeout: timeout,
	}
}

// NewCreateDraftClusterUsingPOST1ParamsWithContext creates a new CreateDraftClusterUsingPOST1Params object
// with the ability to set a context for a request.
func NewCreateDraftClusterUsingPOST1ParamsWithContext(ctx context.Context) *CreateDraftClusterUsingPOST1Params {
	return &CreateDraftClusterUsingPOST1Params{
		Context: ctx,
	}
}

// NewCreateDraftClusterUsingPOST1ParamsWithHTTPClient creates a new CreateDraftClusterUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDraftClusterUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateDraftClusterUsingPOST1Params {
	return &CreateDraftClusterUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
CreateDraftClusterUsingPOST1Params contains all the parameters to send to the API endpoint

	for the create draft cluster using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type CreateDraftClusterUsingPOST1Params struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* Request.

	   request
	*/
	Request *models.KubernetesClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create draft cluster using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDraftClusterUsingPOST1Params) WithDefaults() *CreateDraftClusterUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create draft cluster using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDraftClusterUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateDraftClusterUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) WithContext(ctx context.Context) *CreateDraftClusterUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateDraftClusterUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) WithClusterID(clusterID string) *CreateDraftClusterUsingPOST1Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithRequest adds the request to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) WithRequest(request *models.KubernetesClusterRequest) *CreateDraftClusterUsingPOST1Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create draft cluster using p o s t 1 params
func (o *CreateDraftClusterUsingPOST1Params) SetRequest(request *models.KubernetesClusterRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDraftClusterUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
