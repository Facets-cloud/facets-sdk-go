// Code generated by go-swagger; DO NOT EDIT.

package ui_o_auth_controller

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

// NewGetAllIntegrationsUsingGETParams creates a new GetAllIntegrationsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllIntegrationsUsingGETParams() *GetAllIntegrationsUsingGETParams {
	return &GetAllIntegrationsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllIntegrationsUsingGETParamsWithTimeout creates a new GetAllIntegrationsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllIntegrationsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllIntegrationsUsingGETParams {
	return &GetAllIntegrationsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllIntegrationsUsingGETParamsWithContext creates a new GetAllIntegrationsUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllIntegrationsUsingGETParamsWithContext(ctx context.Context) *GetAllIntegrationsUsingGETParams {
	return &GetAllIntegrationsUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllIntegrationsUsingGETParamsWithHTTPClient creates a new GetAllIntegrationsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllIntegrationsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllIntegrationsUsingGETParams {
	return &GetAllIntegrationsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAllIntegrationsUsingGETParams contains all the parameters to send to the API endpoint

	for the get all integrations using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAllIntegrationsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all integrations using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllIntegrationsUsingGETParams) WithDefaults() *GetAllIntegrationsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all integrations using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllIntegrationsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all integrations using g e t params
func (o *GetAllIntegrationsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllIntegrationsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all integrations using g e t params
func (o *GetAllIntegrationsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all integrations using g e t params
func (o *GetAllIntegrationsUsingGETParams) WithContext(ctx context.Context) *GetAllIntegrationsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all integrations using g e t params
func (o *GetAllIntegrationsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all integrations using g e t params
func (o *GetAllIntegrationsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllIntegrationsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all integrations using g e t params
func (o *GetAllIntegrationsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllIntegrationsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
