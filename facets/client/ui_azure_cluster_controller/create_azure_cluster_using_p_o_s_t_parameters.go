// Code generated by go-swagger; DO NOT EDIT.

package ui_azure_cluster_controller

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

// NewCreateAzureClusterUsingPOSTParams creates a new CreateAzureClusterUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAzureClusterUsingPOSTParams() *CreateAzureClusterUsingPOSTParams {
	return &CreateAzureClusterUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAzureClusterUsingPOSTParamsWithTimeout creates a new CreateAzureClusterUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateAzureClusterUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateAzureClusterUsingPOSTParams {
	return &CreateAzureClusterUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateAzureClusterUsingPOSTParamsWithContext creates a new CreateAzureClusterUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateAzureClusterUsingPOSTParamsWithContext(ctx context.Context) *CreateAzureClusterUsingPOSTParams {
	return &CreateAzureClusterUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateAzureClusterUsingPOSTParamsWithHTTPClient creates a new CreateAzureClusterUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAzureClusterUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateAzureClusterUsingPOSTParams {
	return &CreateAzureClusterUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CreateAzureClusterUsingPOSTParams contains all the parameters to send to the API endpoint

	for the create azure cluster using p o s t operation.

	Typically these are written to a http.Request.
*/
type CreateAzureClusterUsingPOSTParams struct {

	/* Request.

	   request
	*/
	Request *models.AzureClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create azure cluster using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAzureClusterUsingPOSTParams) WithDefaults() *CreateAzureClusterUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create azure cluster using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAzureClusterUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create azure cluster using p o s t params
func (o *CreateAzureClusterUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateAzureClusterUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create azure cluster using p o s t params
func (o *CreateAzureClusterUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create azure cluster using p o s t params
func (o *CreateAzureClusterUsingPOSTParams) WithContext(ctx context.Context) *CreateAzureClusterUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create azure cluster using p o s t params
func (o *CreateAzureClusterUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create azure cluster using p o s t params
func (o *CreateAzureClusterUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateAzureClusterUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create azure cluster using p o s t params
func (o *CreateAzureClusterUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create azure cluster using p o s t params
func (o *CreateAzureClusterUsingPOSTParams) WithRequest(request *models.AzureClusterRequest) *CreateAzureClusterUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create azure cluster using p o s t params
func (o *CreateAzureClusterUsingPOSTParams) SetRequest(request *models.AzureClusterRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAzureClusterUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
