// Code generated by go-swagger; DO NOT EDIT.

package ui_ci_cd_controller

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

// NewRegisterArtifactSaasUsingPOST1Params creates a new RegisterArtifactSaasUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterArtifactSaasUsingPOST1Params() *RegisterArtifactSaasUsingPOST1Params {
	return &RegisterArtifactSaasUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterArtifactSaasUsingPOST1ParamsWithTimeout creates a new RegisterArtifactSaasUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewRegisterArtifactSaasUsingPOST1ParamsWithTimeout(timeout time.Duration) *RegisterArtifactSaasUsingPOST1Params {
	return &RegisterArtifactSaasUsingPOST1Params{
		timeout: timeout,
	}
}

// NewRegisterArtifactSaasUsingPOST1ParamsWithContext creates a new RegisterArtifactSaasUsingPOST1Params object
// with the ability to set a context for a request.
func NewRegisterArtifactSaasUsingPOST1ParamsWithContext(ctx context.Context) *RegisterArtifactSaasUsingPOST1Params {
	return &RegisterArtifactSaasUsingPOST1Params{
		Context: ctx,
	}
}

// NewRegisterArtifactSaasUsingPOST1ParamsWithHTTPClient creates a new RegisterArtifactSaasUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterArtifactSaasUsingPOST1ParamsWithHTTPClient(client *http.Client) *RegisterArtifactSaasUsingPOST1Params {
	return &RegisterArtifactSaasUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
RegisterArtifactSaasUsingPOST1Params contains all the parameters to send to the API endpoint

	for the register artifact saas using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type RegisterArtifactSaasUsingPOST1Params struct {

	/* ArtifactRequest.

	   artifactRequest
	*/
	ArtifactRequest *models.SaasArtifactRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register artifact saas using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterArtifactSaasUsingPOST1Params) WithDefaults() *RegisterArtifactSaasUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register artifact saas using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterArtifactSaasUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register artifact saas using p o s t 1 params
func (o *RegisterArtifactSaasUsingPOST1Params) WithTimeout(timeout time.Duration) *RegisterArtifactSaasUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register artifact saas using p o s t 1 params
func (o *RegisterArtifactSaasUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register artifact saas using p o s t 1 params
func (o *RegisterArtifactSaasUsingPOST1Params) WithContext(ctx context.Context) *RegisterArtifactSaasUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register artifact saas using p o s t 1 params
func (o *RegisterArtifactSaasUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register artifact saas using p o s t 1 params
func (o *RegisterArtifactSaasUsingPOST1Params) WithHTTPClient(client *http.Client) *RegisterArtifactSaasUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register artifact saas using p o s t 1 params
func (o *RegisterArtifactSaasUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactRequest adds the artifactRequest to the register artifact saas using p o s t 1 params
func (o *RegisterArtifactSaasUsingPOST1Params) WithArtifactRequest(artifactRequest *models.SaasArtifactRequest) *RegisterArtifactSaasUsingPOST1Params {
	o.SetArtifactRequest(artifactRequest)
	return o
}

// SetArtifactRequest adds the artifactRequest to the register artifact saas using p o s t 1 params
func (o *RegisterArtifactSaasUsingPOST1Params) SetArtifactRequest(artifactRequest *models.SaasArtifactRequest) {
	o.ArtifactRequest = artifactRequest
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterArtifactSaasUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ArtifactRequest != nil {
		if err := r.SetBodyParam(o.ArtifactRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
