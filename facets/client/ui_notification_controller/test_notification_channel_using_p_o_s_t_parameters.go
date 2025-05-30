// Code generated by go-swagger; DO NOT EDIT.

package ui_notification_controller

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

// NewTestNotificationChannelUsingPOSTParams creates a new TestNotificationChannelUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestNotificationChannelUsingPOSTParams() *TestNotificationChannelUsingPOSTParams {
	return &TestNotificationChannelUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestNotificationChannelUsingPOSTParamsWithTimeout creates a new TestNotificationChannelUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewTestNotificationChannelUsingPOSTParamsWithTimeout(timeout time.Duration) *TestNotificationChannelUsingPOSTParams {
	return &TestNotificationChannelUsingPOSTParams{
		timeout: timeout,
	}
}

// NewTestNotificationChannelUsingPOSTParamsWithContext creates a new TestNotificationChannelUsingPOSTParams object
// with the ability to set a context for a request.
func NewTestNotificationChannelUsingPOSTParamsWithContext(ctx context.Context) *TestNotificationChannelUsingPOSTParams {
	return &TestNotificationChannelUsingPOSTParams{
		Context: ctx,
	}
}

// NewTestNotificationChannelUsingPOSTParamsWithHTTPClient creates a new TestNotificationChannelUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestNotificationChannelUsingPOSTParamsWithHTTPClient(client *http.Client) *TestNotificationChannelUsingPOSTParams {
	return &TestNotificationChannelUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
TestNotificationChannelUsingPOSTParams contains all the parameters to send to the API endpoint

	for the test notification channel using p o s t operation.

	Typically these are written to a http.Request.
*/
type TestNotificationChannelUsingPOSTParams struct {

	/* TestNotificationRequest.

	   testNotificationRequest
	*/
	TestNotificationRequest *models.TestNotificationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test notification channel using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestNotificationChannelUsingPOSTParams) WithDefaults() *TestNotificationChannelUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test notification channel using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestNotificationChannelUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test notification channel using p o s t params
func (o *TestNotificationChannelUsingPOSTParams) WithTimeout(timeout time.Duration) *TestNotificationChannelUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test notification channel using p o s t params
func (o *TestNotificationChannelUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test notification channel using p o s t params
func (o *TestNotificationChannelUsingPOSTParams) WithContext(ctx context.Context) *TestNotificationChannelUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test notification channel using p o s t params
func (o *TestNotificationChannelUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test notification channel using p o s t params
func (o *TestNotificationChannelUsingPOSTParams) WithHTTPClient(client *http.Client) *TestNotificationChannelUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test notification channel using p o s t params
func (o *TestNotificationChannelUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTestNotificationRequest adds the testNotificationRequest to the test notification channel using p o s t params
func (o *TestNotificationChannelUsingPOSTParams) WithTestNotificationRequest(testNotificationRequest *models.TestNotificationRequest) *TestNotificationChannelUsingPOSTParams {
	o.SetTestNotificationRequest(testNotificationRequest)
	return o
}

// SetTestNotificationRequest adds the testNotificationRequest to the test notification channel using p o s t params
func (o *TestNotificationChannelUsingPOSTParams) SetTestNotificationRequest(testNotificationRequest *models.TestNotificationRequest) {
	o.TestNotificationRequest = testNotificationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *TestNotificationChannelUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.TestNotificationRequest != nil {
		if err := r.SetBodyParam(o.TestNotificationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
