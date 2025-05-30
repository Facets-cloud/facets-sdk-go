// Code generated by go-swagger; DO NOT EDIT.

package ui_local_cluster_controller

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

// NewCreateClusterUsingPOST3Params creates a new CreateClusterUsingPOST3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterUsingPOST3Params() *CreateClusterUsingPOST3Params {
	return &CreateClusterUsingPOST3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterUsingPOST3ParamsWithTimeout creates a new CreateClusterUsingPOST3Params object
// with the ability to set a timeout on a request.
func NewCreateClusterUsingPOST3ParamsWithTimeout(timeout time.Duration) *CreateClusterUsingPOST3Params {
	return &CreateClusterUsingPOST3Params{
		timeout: timeout,
	}
}

// NewCreateClusterUsingPOST3ParamsWithContext creates a new CreateClusterUsingPOST3Params object
// with the ability to set a context for a request.
func NewCreateClusterUsingPOST3ParamsWithContext(ctx context.Context) *CreateClusterUsingPOST3Params {
	return &CreateClusterUsingPOST3Params{
		Context: ctx,
	}
}

// NewCreateClusterUsingPOST3ParamsWithHTTPClient creates a new CreateClusterUsingPOST3Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterUsingPOST3ParamsWithHTTPClient(client *http.Client) *CreateClusterUsingPOST3Params {
	return &CreateClusterUsingPOST3Params{
		HTTPClient: client,
	}
}

/*
CreateClusterUsingPOST3Params contains all the parameters to send to the API endpoint

	for the create cluster using p o s t3 operation.

	Typically these are written to a http.Request.
*/
type CreateClusterUsingPOST3Params struct {

	/* Request.

	   request
	*/
	Request *models.LocalClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster using p o s t3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterUsingPOST3Params) WithDefaults() *CreateClusterUsingPOST3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster using p o s t3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterUsingPOST3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster using p o s t3 params
func (o *CreateClusterUsingPOST3Params) WithTimeout(timeout time.Duration) *CreateClusterUsingPOST3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster using p o s t3 params
func (o *CreateClusterUsingPOST3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster using p o s t3 params
func (o *CreateClusterUsingPOST3Params) WithContext(ctx context.Context) *CreateClusterUsingPOST3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster using p o s t3 params
func (o *CreateClusterUsingPOST3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster using p o s t3 params
func (o *CreateClusterUsingPOST3Params) WithHTTPClient(client *http.Client) *CreateClusterUsingPOST3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster using p o s t3 params
func (o *CreateClusterUsingPOST3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create cluster using p o s t3 params
func (o *CreateClusterUsingPOST3Params) WithRequest(request *models.LocalClusterRequest) *CreateClusterUsingPOST3Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create cluster using p o s t3 params
func (o *CreateClusterUsingPOST3Params) SetRequest(request *models.LocalClusterRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterUsingPOST3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
