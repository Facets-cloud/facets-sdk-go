// Code generated by go-swagger; DO NOT EDIT.

package ui_settings_controller

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

// NewGetAllSettingsForEntityUsingGETParams creates a new GetAllSettingsForEntityUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllSettingsForEntityUsingGETParams() *GetAllSettingsForEntityUsingGETParams {
	return &GetAllSettingsForEntityUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllSettingsForEntityUsingGETParamsWithTimeout creates a new GetAllSettingsForEntityUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllSettingsForEntityUsingGETParamsWithTimeout(timeout time.Duration) *GetAllSettingsForEntityUsingGETParams {
	return &GetAllSettingsForEntityUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllSettingsForEntityUsingGETParamsWithContext creates a new GetAllSettingsForEntityUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllSettingsForEntityUsingGETParamsWithContext(ctx context.Context) *GetAllSettingsForEntityUsingGETParams {
	return &GetAllSettingsForEntityUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllSettingsForEntityUsingGETParamsWithHTTPClient creates a new GetAllSettingsForEntityUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllSettingsForEntityUsingGETParamsWithHTTPClient(client *http.Client) *GetAllSettingsForEntityUsingGETParams {
	return &GetAllSettingsForEntityUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetAllSettingsForEntityUsingGETParams contains all the parameters to send to the API endpoint

	for the get all settings for entity using g e t operation.

	Typically these are written to a http.Request.
*/
type GetAllSettingsForEntityUsingGETParams struct {

	/* Entity.

	   entity
	*/
	Entity string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all settings for entity using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSettingsForEntityUsingGETParams) WithDefaults() *GetAllSettingsForEntityUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all settings for entity using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllSettingsForEntityUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all settings for entity using g e t params
func (o *GetAllSettingsForEntityUsingGETParams) WithTimeout(timeout time.Duration) *GetAllSettingsForEntityUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all settings for entity using g e t params
func (o *GetAllSettingsForEntityUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all settings for entity using g e t params
func (o *GetAllSettingsForEntityUsingGETParams) WithContext(ctx context.Context) *GetAllSettingsForEntityUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all settings for entity using g e t params
func (o *GetAllSettingsForEntityUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all settings for entity using g e t params
func (o *GetAllSettingsForEntityUsingGETParams) WithHTTPClient(client *http.Client) *GetAllSettingsForEntityUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all settings for entity using g e t params
func (o *GetAllSettingsForEntityUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the get all settings for entity using g e t params
func (o *GetAllSettingsForEntityUsingGETParams) WithEntity(entity string) *GetAllSettingsForEntityUsingGETParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the get all settings for entity using g e t params
func (o *GetAllSettingsForEntityUsingGETParams) SetEntity(entity string) {
	o.Entity = entity
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllSettingsForEntityUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
