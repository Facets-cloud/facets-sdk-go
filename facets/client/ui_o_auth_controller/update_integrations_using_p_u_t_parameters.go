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

	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

// NewUpdateIntegrationsUsingPUTParams creates a new UpdateIntegrationsUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIntegrationsUsingPUTParams() *UpdateIntegrationsUsingPUTParams {
	return &UpdateIntegrationsUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIntegrationsUsingPUTParamsWithTimeout creates a new UpdateIntegrationsUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateIntegrationsUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateIntegrationsUsingPUTParams {
	return &UpdateIntegrationsUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateIntegrationsUsingPUTParamsWithContext creates a new UpdateIntegrationsUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateIntegrationsUsingPUTParamsWithContext(ctx context.Context) *UpdateIntegrationsUsingPUTParams {
	return &UpdateIntegrationsUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateIntegrationsUsingPUTParamsWithHTTPClient creates a new UpdateIntegrationsUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIntegrationsUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateIntegrationsUsingPUTParams {
	return &UpdateIntegrationsUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateIntegrationsUsingPUTParams contains all the parameters to send to the API endpoint

	for the update integrations using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateIntegrationsUsingPUTParams struct {

	/* Client.

	   client
	*/
	Client *models.CustomOAuth2ClientRegistration

	/* RegistrationID.

	   registrationId
	*/
	RegistrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update integrations using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIntegrationsUsingPUTParams) WithDefaults() *UpdateIntegrationsUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update integrations using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIntegrationsUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateIntegrationsUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) WithContext(ctx context.Context) *UpdateIntegrationsUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateIntegrationsUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClient adds the client to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) WithClient(client *models.CustomOAuth2ClientRegistration) *UpdateIntegrationsUsingPUTParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) SetClient(client *models.CustomOAuth2ClientRegistration) {
	o.Client = client
}

// WithRegistrationID adds the registrationID to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) WithRegistrationID(registrationID string) *UpdateIntegrationsUsingPUTParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the update integrations using p u t params
func (o *UpdateIntegrationsUsingPUTParams) SetRegistrationID(registrationID string) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIntegrationsUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Client != nil {
		if err := r.SetBodyParam(o.Client); err != nil {
			return err
		}
	}

	// path param registrationId
	if err := r.SetPathParam("registrationId", o.RegistrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
