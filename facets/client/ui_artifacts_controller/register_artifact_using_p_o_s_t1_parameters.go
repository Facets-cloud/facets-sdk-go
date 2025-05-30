// Code generated by go-swagger; DO NOT EDIT.

package ui_artifacts_controller

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

// NewRegisterArtifactUsingPOST1Params creates a new RegisterArtifactUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterArtifactUsingPOST1Params() *RegisterArtifactUsingPOST1Params {
	return &RegisterArtifactUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterArtifactUsingPOST1ParamsWithTimeout creates a new RegisterArtifactUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewRegisterArtifactUsingPOST1ParamsWithTimeout(timeout time.Duration) *RegisterArtifactUsingPOST1Params {
	return &RegisterArtifactUsingPOST1Params{
		timeout: timeout,
	}
}

// NewRegisterArtifactUsingPOST1ParamsWithContext creates a new RegisterArtifactUsingPOST1Params object
// with the ability to set a context for a request.
func NewRegisterArtifactUsingPOST1ParamsWithContext(ctx context.Context) *RegisterArtifactUsingPOST1Params {
	return &RegisterArtifactUsingPOST1Params{
		Context: ctx,
	}
}

// NewRegisterArtifactUsingPOST1ParamsWithHTTPClient creates a new RegisterArtifactUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterArtifactUsingPOST1ParamsWithHTTPClient(client *http.Client) *RegisterArtifactUsingPOST1Params {
	return &RegisterArtifactUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
RegisterArtifactUsingPOST1Params contains all the parameters to send to the API endpoint

	for the register artifact using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type RegisterArtifactUsingPOST1Params struct {

	/* Artifact.

	   artifact
	*/
	Artifact *models.Artifact

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register artifact using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterArtifactUsingPOST1Params) WithDefaults() *RegisterArtifactUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register artifact using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterArtifactUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register artifact using p o s t 1 params
func (o *RegisterArtifactUsingPOST1Params) WithTimeout(timeout time.Duration) *RegisterArtifactUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register artifact using p o s t 1 params
func (o *RegisterArtifactUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register artifact using p o s t 1 params
func (o *RegisterArtifactUsingPOST1Params) WithContext(ctx context.Context) *RegisterArtifactUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register artifact using p o s t 1 params
func (o *RegisterArtifactUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register artifact using p o s t 1 params
func (o *RegisterArtifactUsingPOST1Params) WithHTTPClient(client *http.Client) *RegisterArtifactUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register artifact using p o s t 1 params
func (o *RegisterArtifactUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifact adds the artifact to the register artifact using p o s t 1 params
func (o *RegisterArtifactUsingPOST1Params) WithArtifact(artifact *models.Artifact) *RegisterArtifactUsingPOST1Params {
	o.SetArtifact(artifact)
	return o
}

// SetArtifact adds the artifact to the register artifact using p o s t 1 params
func (o *RegisterArtifactUsingPOST1Params) SetArtifact(artifact *models.Artifact) {
	o.Artifact = artifact
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterArtifactUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Artifact != nil {
		if err := r.SetBodyParam(o.Artifact); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
