// Code generated by go-swagger; DO NOT EDIT.

package artifactory_controller

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

// NewCreateECRArtifactoryUsingPOSTParams creates a new CreateECRArtifactoryUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateECRArtifactoryUsingPOSTParams() *CreateECRArtifactoryUsingPOSTParams {
	return &CreateECRArtifactoryUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateECRArtifactoryUsingPOSTParamsWithTimeout creates a new CreateECRArtifactoryUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateECRArtifactoryUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateECRArtifactoryUsingPOSTParams {
	return &CreateECRArtifactoryUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateECRArtifactoryUsingPOSTParamsWithContext creates a new CreateECRArtifactoryUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateECRArtifactoryUsingPOSTParamsWithContext(ctx context.Context) *CreateECRArtifactoryUsingPOSTParams {
	return &CreateECRArtifactoryUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateECRArtifactoryUsingPOSTParamsWithHTTPClient creates a new CreateECRArtifactoryUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateECRArtifactoryUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateECRArtifactoryUsingPOSTParams {
	return &CreateECRArtifactoryUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CreateECRArtifactoryUsingPOSTParams contains all the parameters to send to the API endpoint

	for the create e c r artifactory using p o s t operation.

	Typically these are written to a http.Request.
*/
type CreateECRArtifactoryUsingPOSTParams struct {

	/* EcrArtifactory.

	   ecrArtifactory
	*/
	EcrArtifactory *models.ECRArtifactory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create e c r artifactory using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateECRArtifactoryUsingPOSTParams) WithDefaults() *CreateECRArtifactoryUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create e c r artifactory using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateECRArtifactoryUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create e c r artifactory using p o s t params
func (o *CreateECRArtifactoryUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateECRArtifactoryUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create e c r artifactory using p o s t params
func (o *CreateECRArtifactoryUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create e c r artifactory using p o s t params
func (o *CreateECRArtifactoryUsingPOSTParams) WithContext(ctx context.Context) *CreateECRArtifactoryUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create e c r artifactory using p o s t params
func (o *CreateECRArtifactoryUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create e c r artifactory using p o s t params
func (o *CreateECRArtifactoryUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateECRArtifactoryUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create e c r artifactory using p o s t params
func (o *CreateECRArtifactoryUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEcrArtifactory adds the ecrArtifactory to the create e c r artifactory using p o s t params
func (o *CreateECRArtifactoryUsingPOSTParams) WithEcrArtifactory(ecrArtifactory *models.ECRArtifactory) *CreateECRArtifactoryUsingPOSTParams {
	o.SetEcrArtifactory(ecrArtifactory)
	return o
}

// SetEcrArtifactory adds the ecrArtifactory to the create e c r artifactory using p o s t params
func (o *CreateECRArtifactoryUsingPOSTParams) SetEcrArtifactory(ecrArtifactory *models.ECRArtifactory) {
	o.EcrArtifactory = ecrArtifactory
}

// WriteToRequest writes these params to a swagger request
func (o *CreateECRArtifactoryUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.EcrArtifactory != nil {
		if err := r.SetBodyParam(o.EcrArtifactory); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
