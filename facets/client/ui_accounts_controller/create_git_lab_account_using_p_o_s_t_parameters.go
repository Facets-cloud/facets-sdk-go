// Code generated by go-swagger; DO NOT EDIT.

package ui_accounts_controller

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

// NewCreateGitLabAccountUsingPOSTParams creates a new CreateGitLabAccountUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGitLabAccountUsingPOSTParams() *CreateGitLabAccountUsingPOSTParams {
	return &CreateGitLabAccountUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGitLabAccountUsingPOSTParamsWithTimeout creates a new CreateGitLabAccountUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateGitLabAccountUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateGitLabAccountUsingPOSTParams {
	return &CreateGitLabAccountUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateGitLabAccountUsingPOSTParamsWithContext creates a new CreateGitLabAccountUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateGitLabAccountUsingPOSTParamsWithContext(ctx context.Context) *CreateGitLabAccountUsingPOSTParams {
	return &CreateGitLabAccountUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateGitLabAccountUsingPOSTParamsWithHTTPClient creates a new CreateGitLabAccountUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGitLabAccountUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateGitLabAccountUsingPOSTParams {
	return &CreateGitLabAccountUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CreateGitLabAccountUsingPOSTParams contains all the parameters to send to the API endpoint

	for the create git lab account using p o s t operation.

	Typically these are written to a http.Request.
*/
type CreateGitLabAccountUsingPOSTParams struct {

	/* GitlabAccountRequest.

	   gitlabAccountRequest
	*/
	GitlabAccountRequest *models.GitLabAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create git lab account using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGitLabAccountUsingPOSTParams) WithDefaults() *CreateGitLabAccountUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create git lab account using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGitLabAccountUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create git lab account using p o s t params
func (o *CreateGitLabAccountUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateGitLabAccountUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create git lab account using p o s t params
func (o *CreateGitLabAccountUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create git lab account using p o s t params
func (o *CreateGitLabAccountUsingPOSTParams) WithContext(ctx context.Context) *CreateGitLabAccountUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create git lab account using p o s t params
func (o *CreateGitLabAccountUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create git lab account using p o s t params
func (o *CreateGitLabAccountUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateGitLabAccountUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create git lab account using p o s t params
func (o *CreateGitLabAccountUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGitlabAccountRequest adds the gitlabAccountRequest to the create git lab account using p o s t params
func (o *CreateGitLabAccountUsingPOSTParams) WithGitlabAccountRequest(gitlabAccountRequest *models.GitLabAccount) *CreateGitLabAccountUsingPOSTParams {
	o.SetGitlabAccountRequest(gitlabAccountRequest)
	return o
}

// SetGitlabAccountRequest adds the gitlabAccountRequest to the create git lab account using p o s t params
func (o *CreateGitLabAccountUsingPOSTParams) SetGitlabAccountRequest(gitlabAccountRequest *models.GitLabAccount) {
	o.GitlabAccountRequest = gitlabAccountRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGitLabAccountUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GitlabAccountRequest != nil {
		if err := r.SetBodyParam(o.GitlabAccountRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
