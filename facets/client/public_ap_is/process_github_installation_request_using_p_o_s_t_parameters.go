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

// NewProcessGithubInstallationRequestUsingPOSTParams creates a new ProcessGithubInstallationRequestUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProcessGithubInstallationRequestUsingPOSTParams() *ProcessGithubInstallationRequestUsingPOSTParams {
	return &ProcessGithubInstallationRequestUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProcessGithubInstallationRequestUsingPOSTParamsWithTimeout creates a new ProcessGithubInstallationRequestUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewProcessGithubInstallationRequestUsingPOSTParamsWithTimeout(timeout time.Duration) *ProcessGithubInstallationRequestUsingPOSTParams {
	return &ProcessGithubInstallationRequestUsingPOSTParams{
		timeout: timeout,
	}
}

// NewProcessGithubInstallationRequestUsingPOSTParamsWithContext creates a new ProcessGithubInstallationRequestUsingPOSTParams object
// with the ability to set a context for a request.
func NewProcessGithubInstallationRequestUsingPOSTParamsWithContext(ctx context.Context) *ProcessGithubInstallationRequestUsingPOSTParams {
	return &ProcessGithubInstallationRequestUsingPOSTParams{
		Context: ctx,
	}
}

// NewProcessGithubInstallationRequestUsingPOSTParamsWithHTTPClient creates a new ProcessGithubInstallationRequestUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewProcessGithubInstallationRequestUsingPOSTParamsWithHTTPClient(client *http.Client) *ProcessGithubInstallationRequestUsingPOSTParams {
	return &ProcessGithubInstallationRequestUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
ProcessGithubInstallationRequestUsingPOSTParams contains all the parameters to send to the API endpoint

	for the process github installation request using p o s t operation.

	Typically these are written to a http.Request.
*/
type ProcessGithubInstallationRequestUsingPOSTParams struct {

	/* GithubAppInstallationRequest.

	   githubAppInstallationRequest
	*/
	GithubAppInstallationRequest *models.GithubAppInstallationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the process github installation request using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProcessGithubInstallationRequestUsingPOSTParams) WithDefaults() *ProcessGithubInstallationRequestUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the process github installation request using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProcessGithubInstallationRequestUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the process github installation request using p o s t params
func (o *ProcessGithubInstallationRequestUsingPOSTParams) WithTimeout(timeout time.Duration) *ProcessGithubInstallationRequestUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the process github installation request using p o s t params
func (o *ProcessGithubInstallationRequestUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the process github installation request using p o s t params
func (o *ProcessGithubInstallationRequestUsingPOSTParams) WithContext(ctx context.Context) *ProcessGithubInstallationRequestUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the process github installation request using p o s t params
func (o *ProcessGithubInstallationRequestUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the process github installation request using p o s t params
func (o *ProcessGithubInstallationRequestUsingPOSTParams) WithHTTPClient(client *http.Client) *ProcessGithubInstallationRequestUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the process github installation request using p o s t params
func (o *ProcessGithubInstallationRequestUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGithubAppInstallationRequest adds the githubAppInstallationRequest to the process github installation request using p o s t params
func (o *ProcessGithubInstallationRequestUsingPOSTParams) WithGithubAppInstallationRequest(githubAppInstallationRequest *models.GithubAppInstallationRequest) *ProcessGithubInstallationRequestUsingPOSTParams {
	o.SetGithubAppInstallationRequest(githubAppInstallationRequest)
	return o
}

// SetGithubAppInstallationRequest adds the githubAppInstallationRequest to the process github installation request using p o s t params
func (o *ProcessGithubInstallationRequestUsingPOSTParams) SetGithubAppInstallationRequest(githubAppInstallationRequest *models.GithubAppInstallationRequest) {
	o.GithubAppInstallationRequest = githubAppInstallationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ProcessGithubInstallationRequestUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GithubAppInstallationRequest != nil {
		if err := r.SetBodyParam(o.GithubAppInstallationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
