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
)

// NewPollWebhookUsingGETParams creates a new PollWebhookUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPollWebhookUsingGETParams() *PollWebhookUsingGETParams {
	return &PollWebhookUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPollWebhookUsingGETParamsWithTimeout creates a new PollWebhookUsingGETParams object
// with the ability to set a timeout on a request.
func NewPollWebhookUsingGETParamsWithTimeout(timeout time.Duration) *PollWebhookUsingGETParams {
	return &PollWebhookUsingGETParams{
		timeout: timeout,
	}
}

// NewPollWebhookUsingGETParamsWithContext creates a new PollWebhookUsingGETParams object
// with the ability to set a context for a request.
func NewPollWebhookUsingGETParamsWithContext(ctx context.Context) *PollWebhookUsingGETParams {
	return &PollWebhookUsingGETParams{
		Context: ctx,
	}
}

// NewPollWebhookUsingGETParamsWithHTTPClient creates a new PollWebhookUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewPollWebhookUsingGETParamsWithHTTPClient(client *http.Client) *PollWebhookUsingGETParams {
	return &PollWebhookUsingGETParams{
		HTTPClient: client,
	}
}

/*
PollWebhookUsingGETParams contains all the parameters to send to the API endpoint

	for the poll webhook using g e t operation.

	Typically these are written to a http.Request.
*/
type PollWebhookUsingGETParams struct {

	/* WebhookID.

	   webhookId
	*/
	WebhookID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the poll webhook using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PollWebhookUsingGETParams) WithDefaults() *PollWebhookUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the poll webhook using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PollWebhookUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the poll webhook using g e t params
func (o *PollWebhookUsingGETParams) WithTimeout(timeout time.Duration) *PollWebhookUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the poll webhook using g e t params
func (o *PollWebhookUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the poll webhook using g e t params
func (o *PollWebhookUsingGETParams) WithContext(ctx context.Context) *PollWebhookUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the poll webhook using g e t params
func (o *PollWebhookUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the poll webhook using g e t params
func (o *PollWebhookUsingGETParams) WithHTTPClient(client *http.Client) *PollWebhookUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the poll webhook using g e t params
func (o *PollWebhookUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWebhookID adds the webhookID to the poll webhook using g e t params
func (o *PollWebhookUsingGETParams) WithWebhookID(webhookID string) *PollWebhookUsingGETParams {
	o.SetWebhookID(webhookID)
	return o
}

// SetWebhookID adds the webhookId to the poll webhook using g e t params
func (o *PollWebhookUsingGETParams) SetWebhookID(webhookID string) {
	o.WebhookID = webhookID
}

// WriteToRequest writes these params to a swagger request
func (o *PollWebhookUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param webhookId
	if err := r.SetPathParam("webhookId", o.WebhookID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
