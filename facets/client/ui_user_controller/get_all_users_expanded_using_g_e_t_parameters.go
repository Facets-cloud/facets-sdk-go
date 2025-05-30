// Code generated by go-swagger; DO NOT EDIT.

package ui_user_controller

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

// NewGetAllUsersExpandedUsingGETParams creates a new GetAllUsersExpandedUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllUsersExpandedUsingGETParams() *GetAllUsersExpandedUsingGETParams {
	return &GetAllUsersExpandedUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllUsersExpandedUsingGETParamsWithTimeout creates a new GetAllUsersExpandedUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllUsersExpandedUsingGETParamsWithTimeout(timeout time.Duration) *GetAllUsersExpandedUsingGETParams {
	return &GetAllUsersExpandedUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllUsersExpandedUsingGETParamsWithContext creates a new GetAllUsersExpandedUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllUsersExpandedUsingGETParamsWithContext(ctx context.Context) *GetAllUsersExpandedUsingGETParams {
	return &GetAllUsersExpandedUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllUsersExpandedUsingGETParamsWithHTTPClient creates a new GetAllUsersExpandedUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllUsersExpandedUsingGETParamsWithHTTPClient(client *http.Client) *GetAllUsersExpandedUsingGETParams {
	return &GetAllUsersExpandedUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAllUsersExpandedUsingGETParams contains all the parameters to send to the API endpoint

	for the get all users expanded using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAllUsersExpandedUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all users expanded using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllUsersExpandedUsingGETParams) WithDefaults() *GetAllUsersExpandedUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all users expanded using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllUsersExpandedUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all users expanded using g e t params
func (o *GetAllUsersExpandedUsingGETParams) WithTimeout(timeout time.Duration) *GetAllUsersExpandedUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all users expanded using g e t params
func (o *GetAllUsersExpandedUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all users expanded using g e t params
func (o *GetAllUsersExpandedUsingGETParams) WithContext(ctx context.Context) *GetAllUsersExpandedUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all users expanded using g e t params
func (o *GetAllUsersExpandedUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all users expanded using g e t params
func (o *GetAllUsersExpandedUsingGETParams) WithHTTPClient(client *http.Client) *GetAllUsersExpandedUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all users expanded using g e t params
func (o *GetAllUsersExpandedUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllUsersExpandedUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
