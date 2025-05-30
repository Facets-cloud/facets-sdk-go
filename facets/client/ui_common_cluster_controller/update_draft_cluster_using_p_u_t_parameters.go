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

// NewUpdateDraftClusterUsingPUTParams creates a new UpdateDraftClusterUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDraftClusterUsingPUTParams() *UpdateDraftClusterUsingPUTParams {
	return &UpdateDraftClusterUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDraftClusterUsingPUTParamsWithTimeout creates a new UpdateDraftClusterUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateDraftClusterUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateDraftClusterUsingPUTParams {
	return &UpdateDraftClusterUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateDraftClusterUsingPUTParamsWithContext creates a new UpdateDraftClusterUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateDraftClusterUsingPUTParamsWithContext(ctx context.Context) *UpdateDraftClusterUsingPUTParams {
	return &UpdateDraftClusterUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateDraftClusterUsingPUTParamsWithHTTPClient creates a new UpdateDraftClusterUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDraftClusterUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateDraftClusterUsingPUTParams {
	return &UpdateDraftClusterUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateDraftClusterUsingPUTParams contains all the parameters to send to the API endpoint

	for the update draft cluster using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateDraftClusterUsingPUTParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* ClusterRequest.

	   clusterRequest
	*/
	ClusterRequest *models.DraftClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update draft cluster using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDraftClusterUsingPUTParams) WithDefaults() *UpdateDraftClusterUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update draft cluster using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDraftClusterUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateDraftClusterUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) WithContext(ctx context.Context) *UpdateDraftClusterUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateDraftClusterUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) WithClusterID(clusterID string) *UpdateDraftClusterUsingPUTParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithClusterRequest adds the clusterRequest to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) WithClusterRequest(clusterRequest *models.DraftClusterRequest) *UpdateDraftClusterUsingPUTParams {
	o.SetClusterRequest(clusterRequest)
	return o
}

// SetClusterRequest adds the clusterRequest to the update draft cluster using p u t params
func (o *UpdateDraftClusterUsingPUTParams) SetClusterRequest(clusterRequest *models.DraftClusterRequest) {
	o.ClusterRequest = clusterRequest
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDraftClusterUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}
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
