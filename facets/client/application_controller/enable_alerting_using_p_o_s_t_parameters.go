// Code generated by go-swagger; DO NOT EDIT.

package application_controller

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

// NewEnableAlertingUsingPOSTParams creates a new EnableAlertingUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableAlertingUsingPOSTParams() *EnableAlertingUsingPOSTParams {
	return &EnableAlertingUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableAlertingUsingPOSTParamsWithTimeout creates a new EnableAlertingUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewEnableAlertingUsingPOSTParamsWithTimeout(timeout time.Duration) *EnableAlertingUsingPOSTParams {
	return &EnableAlertingUsingPOSTParams{
		timeout: timeout,
	}
}

// NewEnableAlertingUsingPOSTParamsWithContext creates a new EnableAlertingUsingPOSTParams object
// with the ability to set a context for a request.
func NewEnableAlertingUsingPOSTParamsWithContext(ctx context.Context) *EnableAlertingUsingPOSTParams {
	return &EnableAlertingUsingPOSTParams{
		Context: ctx,
	}
}

// NewEnableAlertingUsingPOSTParamsWithHTTPClient creates a new EnableAlertingUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableAlertingUsingPOSTParamsWithHTTPClient(client *http.Client) *EnableAlertingUsingPOSTParams {
	return &EnableAlertingUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
EnableAlertingUsingPOSTParams contains all the parameters to send to the API endpoint

	for the enable alerting using p o s t operation.

	Typically these are written to a http.Request.
*/
type EnableAlertingUsingPOSTParams struct {

	/* ApplicationFamily.

	   applicationFamily
	*/
	ApplicationFamily string

	/* ApplicationID.

	   applicationId
	*/
	ApplicationID string

	/* Environment.

	   environment
	*/
	Environment string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable alerting using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableAlertingUsingPOSTParams) WithDefaults() *EnableAlertingUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable alerting using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableAlertingUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) WithTimeout(timeout time.Duration) *EnableAlertingUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) WithContext(ctx context.Context) *EnableAlertingUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) WithHTTPClient(client *http.Client) *EnableAlertingUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationFamily adds the applicationFamily to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) WithApplicationFamily(applicationFamily string) *EnableAlertingUsingPOSTParams {
	o.SetApplicationFamily(applicationFamily)
	return o
}

// SetApplicationFamily adds the applicationFamily to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) SetApplicationFamily(applicationFamily string) {
	o.ApplicationFamily = applicationFamily
}

// WithApplicationID adds the applicationID to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) WithApplicationID(applicationID string) *EnableAlertingUsingPOSTParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) SetApplicationID(applicationID string) {
	o.ApplicationID = applicationID
}

// WithEnvironment adds the environment to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) WithEnvironment(environment string) *EnableAlertingUsingPOSTParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the enable alerting using p o s t params
func (o *EnableAlertingUsingPOSTParams) SetEnvironment(environment string) {
	o.Environment = environment
}

// WriteToRequest writes these params to a swagger request
func (o *EnableAlertingUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param applicationFamily
	if err := r.SetPathParam("applicationFamily", o.ApplicationFamily); err != nil {
		return err
	}

	// path param applicationId
	if err := r.SetPathParam("applicationId", o.ApplicationID); err != nil {
		return err
	}

	// path param environment
	if err := r.SetPathParam("environment", o.Environment); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
