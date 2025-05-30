// Code generated by go-swagger; DO NOT EDIT.

package public_ap_is

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

// NewLinkGitlabAccountUsingPOSTParams creates a new LinkGitlabAccountUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLinkGitlabAccountUsingPOSTParams() *LinkGitlabAccountUsingPOSTParams {
	return &LinkGitlabAccountUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLinkGitlabAccountUsingPOSTParamsWithTimeout creates a new LinkGitlabAccountUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewLinkGitlabAccountUsingPOSTParamsWithTimeout(timeout time.Duration) *LinkGitlabAccountUsingPOSTParams {
	return &LinkGitlabAccountUsingPOSTParams{
		timeout: timeout,
	}
}

// NewLinkGitlabAccountUsingPOSTParamsWithContext creates a new LinkGitlabAccountUsingPOSTParams object
// with the ability to set a context for a request.
func NewLinkGitlabAccountUsingPOSTParamsWithContext(ctx context.Context) *LinkGitlabAccountUsingPOSTParams {
	return &LinkGitlabAccountUsingPOSTParams{
		Context: ctx,
	}
}

// NewLinkGitlabAccountUsingPOSTParamsWithHTTPClient creates a new LinkGitlabAccountUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewLinkGitlabAccountUsingPOSTParamsWithHTTPClient(client *http.Client) *LinkGitlabAccountUsingPOSTParams {
	return &LinkGitlabAccountUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
LinkGitlabAccountUsingPOSTParams contains all the parameters to send to the API endpoint

	for the link gitlab account using p o s t operation.

	Typically these are written to a http.Request.
*/
type LinkGitlabAccountUsingPOSTParams struct {

	/* Payload.

	   payload
	*/
	Payload *models.OneTimePayloadGitlabOauthAppPayload

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the link gitlab account using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LinkGitlabAccountUsingPOSTParams) WithDefaults() *LinkGitlabAccountUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the link gitlab account using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LinkGitlabAccountUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the link gitlab account using p o s t params
func (o *LinkGitlabAccountUsingPOSTParams) WithTimeout(timeout time.Duration) *LinkGitlabAccountUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the link gitlab account using p o s t params
func (o *LinkGitlabAccountUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the link gitlab account using p o s t params
func (o *LinkGitlabAccountUsingPOSTParams) WithContext(ctx context.Context) *LinkGitlabAccountUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the link gitlab account using p o s t params
func (o *LinkGitlabAccountUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the link gitlab account using p o s t params
func (o *LinkGitlabAccountUsingPOSTParams) WithHTTPClient(client *http.Client) *LinkGitlabAccountUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the link gitlab account using p o s t params
func (o *LinkGitlabAccountUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the link gitlab account using p o s t params
func (o *LinkGitlabAccountUsingPOSTParams) WithPayload(payload *models.OneTimePayloadGitlabOauthAppPayload) *LinkGitlabAccountUsingPOSTParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the link gitlab account using p o s t params
func (o *LinkGitlabAccountUsingPOSTParams) SetPayload(payload *models.OneTimePayloadGitlabOauthAppPayload) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *LinkGitlabAccountUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
