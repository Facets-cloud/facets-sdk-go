// Code generated by go-swagger; DO NOT EDIT.

package ui_artifactory_controller

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

// NewCreateECRArtifactoryUsingPOST1Params creates a new CreateECRArtifactoryUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateECRArtifactoryUsingPOST1Params() *CreateECRArtifactoryUsingPOST1Params {
	return &CreateECRArtifactoryUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateECRArtifactoryUsingPOST1ParamsWithTimeout creates a new CreateECRArtifactoryUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewCreateECRArtifactoryUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateECRArtifactoryUsingPOST1Params {
	return &CreateECRArtifactoryUsingPOST1Params{
		timeout: timeout,
	}
}

// NewCreateECRArtifactoryUsingPOST1ParamsWithContext creates a new CreateECRArtifactoryUsingPOST1Params object
// with the ability to set a context for a request.
func NewCreateECRArtifactoryUsingPOST1ParamsWithContext(ctx context.Context) *CreateECRArtifactoryUsingPOST1Params {
	return &CreateECRArtifactoryUsingPOST1Params{
		Context: ctx,
	}
}

// NewCreateECRArtifactoryUsingPOST1ParamsWithHTTPClient creates a new CreateECRArtifactoryUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateECRArtifactoryUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateECRArtifactoryUsingPOST1Params {
	return &CreateECRArtifactoryUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
CreateECRArtifactoryUsingPOST1Params contains all the parameters to send to the API endpoint

	for the create e c r artifactory using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type CreateECRArtifactoryUsingPOST1Params struct {

	/* EcrArtifactory.

	   ecrArtifactory
	*/
	EcrArtifactory *models.ECRArtifactory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create e c r artifactory using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateECRArtifactoryUsingPOST1Params) WithDefaults() *CreateECRArtifactoryUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create e c r artifactory using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateECRArtifactoryUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create e c r artifactory using p o s t 1 params
func (o *CreateECRArtifactoryUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateECRArtifactoryUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create e c r artifactory using p o s t 1 params
func (o *CreateECRArtifactoryUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create e c r artifactory using p o s t 1 params
func (o *CreateECRArtifactoryUsingPOST1Params) WithContext(ctx context.Context) *CreateECRArtifactoryUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create e c r artifactory using p o s t 1 params
func (o *CreateECRArtifactoryUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create e c r artifactory using p o s t 1 params
func (o *CreateECRArtifactoryUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateECRArtifactoryUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create e c r artifactory using p o s t 1 params
func (o *CreateECRArtifactoryUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEcrArtifactory adds the ecrArtifactory to the create e c r artifactory using p o s t 1 params
func (o *CreateECRArtifactoryUsingPOST1Params) WithEcrArtifactory(ecrArtifactory *models.ECRArtifactory) *CreateECRArtifactoryUsingPOST1Params {
	o.SetEcrArtifactory(ecrArtifactory)
	return o
}

// SetEcrArtifactory adds the ecrArtifactory to the create e c r artifactory using p o s t 1 params
func (o *CreateECRArtifactoryUsingPOST1Params) SetEcrArtifactory(ecrArtifactory *models.ECRArtifactory) {
	o.EcrArtifactory = ecrArtifactory
}

// WriteToRequest writes these params to a swagger request
func (o *CreateECRArtifactoryUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
