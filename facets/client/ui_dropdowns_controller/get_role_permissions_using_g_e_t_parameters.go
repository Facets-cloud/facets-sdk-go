// Code generated by go-swagger; DO NOT EDIT.

package ui_dropdowns_controller

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

// NewGetRolePermissionsUsingGETParams creates a new GetRolePermissionsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRolePermissionsUsingGETParams() *GetRolePermissionsUsingGETParams {
	return &GetRolePermissionsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRolePermissionsUsingGETParamsWithTimeout creates a new GetRolePermissionsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetRolePermissionsUsingGETParamsWithTimeout(timeout time.Duration) *GetRolePermissionsUsingGETParams {
	return &GetRolePermissionsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetRolePermissionsUsingGETParamsWithContext creates a new GetRolePermissionsUsingGETParams object
// with the ability to set a context for a request.
func NewGetRolePermissionsUsingGETParamsWithContext(ctx context.Context) *GetRolePermissionsUsingGETParams {
	return &GetRolePermissionsUsingGETParams{
		Context: ctx,
	}
}

// NewGetRolePermissionsUsingGETParamsWithHTTPClient creates a new GetRolePermissionsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRolePermissionsUsingGETParamsWithHTTPClient(client *http.Client) *GetRolePermissionsUsingGETParams {
	return &GetRolePermissionsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetRolePermissionsUsingGETParams contains all the parameters to send to the API endpoint

	for the get role permissions using g e t operation.

	Typically these are written to a http.Request.
*/
type GetRolePermissionsUsingGETParams struct {

	/* Role.

	   role
	*/
	Role string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get role permissions using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolePermissionsUsingGETParams) WithDefaults() *GetRolePermissionsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get role permissions using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRolePermissionsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get role permissions using g e t params
func (o *GetRolePermissionsUsingGETParams) WithTimeout(timeout time.Duration) *GetRolePermissionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role permissions using g e t params
func (o *GetRolePermissionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role permissions using g e t params
func (o *GetRolePermissionsUsingGETParams) WithContext(ctx context.Context) *GetRolePermissionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role permissions using g e t params
func (o *GetRolePermissionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role permissions using g e t params
func (o *GetRolePermissionsUsingGETParams) WithHTTPClient(client *http.Client) *GetRolePermissionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role permissions using g e t params
func (o *GetRolePermissionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRole adds the role to the get role permissions using g e t params
func (o *GetRolePermissionsUsingGETParams) WithRole(role string) *GetRolePermissionsUsingGETParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the get role permissions using g e t params
func (o *GetRolePermissionsUsingGETParams) SetRole(role string) {
	o.Role = role
}

// WriteToRequest writes these params to a swagger request
func (o *GetRolePermissionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param role
	if err := r.SetPathParam("role", o.Role); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
