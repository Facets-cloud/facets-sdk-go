// Code generated by go-swagger; DO NOT EDIT.

package ui_no_auth_user_controller

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

// NewValidateTokenUsingPOSTParams creates a new ValidateTokenUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateTokenUsingPOSTParams() *ValidateTokenUsingPOSTParams {
	return &ValidateTokenUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateTokenUsingPOSTParamsWithTimeout creates a new ValidateTokenUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewValidateTokenUsingPOSTParamsWithTimeout(timeout time.Duration) *ValidateTokenUsingPOSTParams {
	return &ValidateTokenUsingPOSTParams{
		timeout: timeout,
	}
}

// NewValidateTokenUsingPOSTParamsWithContext creates a new ValidateTokenUsingPOSTParams object
// with the ability to set a context for a request.
func NewValidateTokenUsingPOSTParamsWithContext(ctx context.Context) *ValidateTokenUsingPOSTParams {
	return &ValidateTokenUsingPOSTParams{
		Context: ctx,
	}
}

// NewValidateTokenUsingPOSTParamsWithHTTPClient creates a new ValidateTokenUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateTokenUsingPOSTParamsWithHTTPClient(client *http.Client) *ValidateTokenUsingPOSTParams {
	return &ValidateTokenUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
ValidateTokenUsingPOSTParams contains all the parameters to send to the API endpoint

	for the validate token using p o s t operation.

	Typically these are written to a http.Request.
*/
type ValidateTokenUsingPOSTParams struct {

	/* ValidationRequest.

	   validationRequest
	*/
	ValidationRequest *models.TokenValidationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate token using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateTokenUsingPOSTParams) WithDefaults() *ValidateTokenUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate token using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateTokenUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate token using p o s t params
func (o *ValidateTokenUsingPOSTParams) WithTimeout(timeout time.Duration) *ValidateTokenUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate token using p o s t params
func (o *ValidateTokenUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate token using p o s t params
func (o *ValidateTokenUsingPOSTParams) WithContext(ctx context.Context) *ValidateTokenUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate token using p o s t params
func (o *ValidateTokenUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate token using p o s t params
func (o *ValidateTokenUsingPOSTParams) WithHTTPClient(client *http.Client) *ValidateTokenUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate token using p o s t params
func (o *ValidateTokenUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidationRequest adds the validationRequest to the validate token using p o s t params
func (o *ValidateTokenUsingPOSTParams) WithValidationRequest(validationRequest *models.TokenValidationRequest) *ValidateTokenUsingPOSTParams {
	o.SetValidationRequest(validationRequest)
	return o
}

// SetValidationRequest adds the validationRequest to the validate token using p o s t params
func (o *ValidateTokenUsingPOSTParams) SetValidationRequest(validationRequest *models.TokenValidationRequest) {
	o.ValidationRequest = validationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateTokenUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ValidationRequest != nil {
		if err := r.SetBodyParam(o.ValidationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
