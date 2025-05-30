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

	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

// NewCreateDraftClusterUsingPOSTParams creates a new CreateDraftClusterUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDraftClusterUsingPOSTParams() *CreateDraftClusterUsingPOSTParams {
	return &CreateDraftClusterUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDraftClusterUsingPOSTParamsWithTimeout creates a new CreateDraftClusterUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateDraftClusterUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateDraftClusterUsingPOSTParams {
	return &CreateDraftClusterUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateDraftClusterUsingPOSTParamsWithContext creates a new CreateDraftClusterUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateDraftClusterUsingPOSTParamsWithContext(ctx context.Context) *CreateDraftClusterUsingPOSTParams {
	return &CreateDraftClusterUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateDraftClusterUsingPOSTParamsWithHTTPClient creates a new CreateDraftClusterUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDraftClusterUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateDraftClusterUsingPOSTParams {
	return &CreateDraftClusterUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CreateDraftClusterUsingPOSTParams contains all the parameters to send to the API endpoint

	for the create draft cluster using p o s t operation.

	Typically these are written to a http.Request.
*/
type CreateDraftClusterUsingPOSTParams struct {

	/* ClusterRequest.

	   clusterRequest
	*/
	ClusterRequest *models.DraftClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create draft cluster using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDraftClusterUsingPOSTParams) WithDefaults() *CreateDraftClusterUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create draft cluster using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDraftClusterUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create draft cluster using p o s t params
func (o *CreateDraftClusterUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateDraftClusterUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create draft cluster using p o s t params
func (o *CreateDraftClusterUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create draft cluster using p o s t params
func (o *CreateDraftClusterUsingPOSTParams) WithContext(ctx context.Context) *CreateDraftClusterUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create draft cluster using p o s t params
func (o *CreateDraftClusterUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create draft cluster using p o s t params
func (o *CreateDraftClusterUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateDraftClusterUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create draft cluster using p o s t params
func (o *CreateDraftClusterUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterRequest adds the clusterRequest to the create draft cluster using p o s t params
func (o *CreateDraftClusterUsingPOSTParams) WithClusterRequest(clusterRequest *models.DraftClusterRequest) *CreateDraftClusterUsingPOSTParams {
	o.SetClusterRequest(clusterRequest)
	return o
}

// SetClusterRequest adds the clusterRequest to the create draft cluster using p o s t params
func (o *CreateDraftClusterUsingPOSTParams) SetClusterRequest(clusterRequest *models.DraftClusterRequest) {
	o.ClusterRequest = clusterRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDraftClusterUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ClusterRequest != nil {
		if err := r.SetBodyParam(o.ClusterRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
