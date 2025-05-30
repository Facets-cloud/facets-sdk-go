// Code generated by go-swagger; DO NOT EDIT.

package gcp_cluster_controller

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

// NewCreateClusterUsingPOST2Params creates a new CreateClusterUsingPOST2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterUsingPOST2Params() *CreateClusterUsingPOST2Params {
	return &CreateClusterUsingPOST2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterUsingPOST2ParamsWithTimeout creates a new CreateClusterUsingPOST2Params object
// with the ability to set a timeout on a request.
func NewCreateClusterUsingPOST2ParamsWithTimeout(timeout time.Duration) *CreateClusterUsingPOST2Params {
	return &CreateClusterUsingPOST2Params{
		timeout: timeout,
	}
}

// NewCreateClusterUsingPOST2ParamsWithContext creates a new CreateClusterUsingPOST2Params object
// with the ability to set a context for a request.
func NewCreateClusterUsingPOST2ParamsWithContext(ctx context.Context) *CreateClusterUsingPOST2Params {
	return &CreateClusterUsingPOST2Params{
		Context: ctx,
	}
}

// NewCreateClusterUsingPOST2ParamsWithHTTPClient creates a new CreateClusterUsingPOST2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterUsingPOST2ParamsWithHTTPClient(client *http.Client) *CreateClusterUsingPOST2Params {
	return &CreateClusterUsingPOST2Params{
		HTTPClient: client,
	}
}

/*
CreateClusterUsingPOST2Params contains all the parameters to send to the API endpoint

	for the create cluster using p o s t 2 operation.

	Typically these are written to a http.Request.
*/
type CreateClusterUsingPOST2Params struct {

	/* Request.

	   request
	*/
	Request *models.GCPClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterUsingPOST2Params) WithDefaults() *CreateClusterUsingPOST2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterUsingPOST2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster using p o s t 2 params
func (o *CreateClusterUsingPOST2Params) WithTimeout(timeout time.Duration) *CreateClusterUsingPOST2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster using p o s t 2 params
func (o *CreateClusterUsingPOST2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster using p o s t 2 params
func (o *CreateClusterUsingPOST2Params) WithContext(ctx context.Context) *CreateClusterUsingPOST2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster using p o s t 2 params
func (o *CreateClusterUsingPOST2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster using p o s t 2 params
func (o *CreateClusterUsingPOST2Params) WithHTTPClient(client *http.Client) *CreateClusterUsingPOST2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster using p o s t 2 params
func (o *CreateClusterUsingPOST2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create cluster using p o s t 2 params
func (o *CreateClusterUsingPOST2Params) WithRequest(request *models.GCPClusterRequest) *CreateClusterUsingPOST2Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create cluster using p o s t 2 params
func (o *CreateClusterUsingPOST2Params) SetRequest(request *models.GCPClusterRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterUsingPOST2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
