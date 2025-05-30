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

// NewValidateGitlabAccountUsingPOSTParams creates a new ValidateGitlabAccountUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateGitlabAccountUsingPOSTParams() *ValidateGitlabAccountUsingPOSTParams {
	return &ValidateGitlabAccountUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateGitlabAccountUsingPOSTParamsWithTimeout creates a new ValidateGitlabAccountUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewValidateGitlabAccountUsingPOSTParamsWithTimeout(timeout time.Duration) *ValidateGitlabAccountUsingPOSTParams {
	return &ValidateGitlabAccountUsingPOSTParams{
		timeout: timeout,
	}
}

// NewValidateGitlabAccountUsingPOSTParamsWithContext creates a new ValidateGitlabAccountUsingPOSTParams object
// with the ability to set a context for a request.
func NewValidateGitlabAccountUsingPOSTParamsWithContext(ctx context.Context) *ValidateGitlabAccountUsingPOSTParams {
	return &ValidateGitlabAccountUsingPOSTParams{
		Context: ctx,
	}
}

// NewValidateGitlabAccountUsingPOSTParamsWithHTTPClient creates a new ValidateGitlabAccountUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateGitlabAccountUsingPOSTParamsWithHTTPClient(client *http.Client) *ValidateGitlabAccountUsingPOSTParams {
	return &ValidateGitlabAccountUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
ValidateGitlabAccountUsingPOSTParams contains all the parameters to send to the API endpoint

	for the validate gitlab account using p o s t operation.

	Typically these are written to a http.Request.
*/
type ValidateGitlabAccountUsingPOSTParams struct {

	/* GitLabAccount.

	   gitLabAccount
	*/
	GitLabAccount *models.GitLabAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate gitlab account using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateGitlabAccountUsingPOSTParams) WithDefaults() *ValidateGitlabAccountUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate gitlab account using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateGitlabAccountUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate gitlab account using p o s t params
func (o *ValidateGitlabAccountUsingPOSTParams) WithTimeout(timeout time.Duration) *ValidateGitlabAccountUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate gitlab account using p o s t params
func (o *ValidateGitlabAccountUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate gitlab account using p o s t params
func (o *ValidateGitlabAccountUsingPOSTParams) WithContext(ctx context.Context) *ValidateGitlabAccountUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate gitlab account using p o s t params
func (o *ValidateGitlabAccountUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate gitlab account using p o s t params
func (o *ValidateGitlabAccountUsingPOSTParams) WithHTTPClient(client *http.Client) *ValidateGitlabAccountUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate gitlab account using p o s t params
func (o *ValidateGitlabAccountUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGitLabAccount adds the gitLabAccount to the validate gitlab account using p o s t params
func (o *ValidateGitlabAccountUsingPOSTParams) WithGitLabAccount(gitLabAccount *models.GitLabAccount) *ValidateGitlabAccountUsingPOSTParams {
	o.SetGitLabAccount(gitLabAccount)
	return o
}

// SetGitLabAccount adds the gitLabAccount to the validate gitlab account using p o s t params
func (o *ValidateGitlabAccountUsingPOSTParams) SetGitLabAccount(gitLabAccount *models.GitLabAccount) {
	o.GitLabAccount = gitLabAccount
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateGitlabAccountUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GitLabAccount != nil {
		if err := r.SetBodyParam(o.GitLabAccount); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
