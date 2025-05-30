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
)

// NewGetAllArtifactoriesUsingGET1Params creates a new GetAllArtifactoriesUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllArtifactoriesUsingGET1Params() *GetAllArtifactoriesUsingGET1Params {
	return &GetAllArtifactoriesUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllArtifactoriesUsingGET1ParamsWithTimeout creates a new GetAllArtifactoriesUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetAllArtifactoriesUsingGET1ParamsWithTimeout(timeout time.Duration) *GetAllArtifactoriesUsingGET1Params {
	return &GetAllArtifactoriesUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetAllArtifactoriesUsingGET1ParamsWithContext creates a new GetAllArtifactoriesUsingGET1Params object
// with the ability to set a context for a request.
func NewGetAllArtifactoriesUsingGET1ParamsWithContext(ctx context.Context) *GetAllArtifactoriesUsingGET1Params {
	return &GetAllArtifactoriesUsingGET1Params{
		Context: ctx,
	}
}

// NewGetAllArtifactoriesUsingGET1ParamsWithHTTPClient creates a new GetAllArtifactoriesUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllArtifactoriesUsingGET1ParamsWithHTTPClient(client *http.Client) *GetAllArtifactoriesUsingGET1Params {
	return &GetAllArtifactoriesUsingGET1Params{
		HTTPClient: client,
	}
}

/*
GetAllArtifactoriesUsingGET1Params contains all the parameters to send to the API endpoint

	for the get all artifactories using g e t 1 operation.

	Typically these are written to a http.Request.
*/
type GetAllArtifactoriesUsingGET1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all artifactories using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllArtifactoriesUsingGET1Params) WithDefaults() *GetAllArtifactoriesUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all artifactories using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllArtifactoriesUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all artifactories using g e t 1 params
func (o *GetAllArtifactoriesUsingGET1Params) WithTimeout(timeout time.Duration) *GetAllArtifactoriesUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all artifactories using g e t 1 params
func (o *GetAllArtifactoriesUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all artifactories using g e t 1 params
func (o *GetAllArtifactoriesUsingGET1Params) WithContext(ctx context.Context) *GetAllArtifactoriesUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all artifactories using g e t 1 params
func (o *GetAllArtifactoriesUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all artifactories using g e t 1 params
func (o *GetAllArtifactoriesUsingGET1Params) WithHTTPClient(client *http.Client) *GetAllArtifactoriesUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all artifactories using g e t 1 params
func (o *GetAllArtifactoriesUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllArtifactoriesUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
