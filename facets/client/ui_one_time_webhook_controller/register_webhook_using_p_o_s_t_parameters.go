// Code generated by go-swagger; DO NOT EDIT.

package ui_one_time_webhook_controller

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

// NewRegisterWebhookUsingPOSTParams creates a new RegisterWebhookUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegisterWebhookUsingPOSTParams() *RegisterWebhookUsingPOSTParams {
	return &RegisterWebhookUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterWebhookUsingPOSTParamsWithTimeout creates a new RegisterWebhookUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewRegisterWebhookUsingPOSTParamsWithTimeout(timeout time.Duration) *RegisterWebhookUsingPOSTParams {
	return &RegisterWebhookUsingPOSTParams{
		timeout: timeout,
	}
}

// NewRegisterWebhookUsingPOSTParamsWithContext creates a new RegisterWebhookUsingPOSTParams object
// with the ability to set a context for a request.
func NewRegisterWebhookUsingPOSTParamsWithContext(ctx context.Context) *RegisterWebhookUsingPOSTParams {
	return &RegisterWebhookUsingPOSTParams{
		Context: ctx,
	}
}

// NewRegisterWebhookUsingPOSTParamsWithHTTPClient creates a new RegisterWebhookUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegisterWebhookUsingPOSTParamsWithHTTPClient(client *http.Client) *RegisterWebhookUsingPOSTParams {
	return &RegisterWebhookUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
RegisterWebhookUsingPOSTParams contains all the parameters to send to the API endpoint

	for the register webhook using p o s t operation.

	Typically these are written to a http.Request.
*/
type RegisterWebhookUsingPOSTParams struct {

	/* Webhook.

	   webhook
	*/
	Webhook *models.OneTimeWebhook

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the register webhook using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterWebhookUsingPOSTParams) WithDefaults() *RegisterWebhookUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the register webhook using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegisterWebhookUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the register webhook using p o s t params
func (o *RegisterWebhookUsingPOSTParams) WithTimeout(timeout time.Duration) *RegisterWebhookUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register webhook using p o s t params
func (o *RegisterWebhookUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register webhook using p o s t params
func (o *RegisterWebhookUsingPOSTParams) WithContext(ctx context.Context) *RegisterWebhookUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register webhook using p o s t params
func (o *RegisterWebhookUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register webhook using p o s t params
func (o *RegisterWebhookUsingPOSTParams) WithHTTPClient(client *http.Client) *RegisterWebhookUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register webhook using p o s t params
func (o *RegisterWebhookUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWebhook adds the webhook to the register webhook using p o s t params
func (o *RegisterWebhookUsingPOSTParams) WithWebhook(webhook *models.OneTimeWebhook) *RegisterWebhookUsingPOSTParams {
	o.SetWebhook(webhook)
	return o
}

// SetWebhook adds the webhook to the register webhook using p o s t params
func (o *RegisterWebhookUsingPOSTParams) SetWebhook(webhook *models.OneTimeWebhook) {
	o.Webhook = webhook
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterWebhookUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Webhook != nil {
		if err := r.SetBodyParam(o.Webhook); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
