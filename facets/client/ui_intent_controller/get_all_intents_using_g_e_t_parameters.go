// Code generated by go-swagger; DO NOT EDIT.

package ui_intent_controller

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

// NewGetAllIntentsUsingGETParams creates a new GetAllIntentsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllIntentsUsingGETParams() *GetAllIntentsUsingGETParams {
	return &GetAllIntentsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllIntentsUsingGETParamsWithTimeout creates a new GetAllIntentsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllIntentsUsingGETParamsWithTimeout(timeout time.Duration) *GetAllIntentsUsingGETParams {
	return &GetAllIntentsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllIntentsUsingGETParamsWithContext creates a new GetAllIntentsUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllIntentsUsingGETParamsWithContext(ctx context.Context) *GetAllIntentsUsingGETParams {
	return &GetAllIntentsUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllIntentsUsingGETParamsWithHTTPClient creates a new GetAllIntentsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllIntentsUsingGETParamsWithHTTPClient(client *http.Client) *GetAllIntentsUsingGETParams {
	return &GetAllIntentsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAllIntentsUsingGETParams contains all the parameters to send to the API endpoint

	for the get all intents using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAllIntentsUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all intents using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllIntentsUsingGETParams) WithDefaults() *GetAllIntentsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all intents using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllIntentsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all intents using g e t params
func (o *GetAllIntentsUsingGETParams) WithTimeout(timeout time.Duration) *GetAllIntentsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all intents using g e t params
func (o *GetAllIntentsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all intents using g e t params
func (o *GetAllIntentsUsingGETParams) WithContext(ctx context.Context) *GetAllIntentsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all intents using g e t params
func (o *GetAllIntentsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all intents using g e t params
func (o *GetAllIntentsUsingGETParams) WithHTTPClient(client *http.Client) *GetAllIntentsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all intents using g e t params
func (o *GetAllIntentsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllIntentsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
