// Code generated by go-swagger; DO NOT EDIT.

package ui_iac_repo_controller

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

// NewUpdateIacRepoUsingPUTParams creates a new UpdateIacRepoUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIacRepoUsingPUTParams() *UpdateIacRepoUsingPUTParams {
	return &UpdateIacRepoUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIacRepoUsingPUTParamsWithTimeout creates a new UpdateIacRepoUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateIacRepoUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateIacRepoUsingPUTParams {
	return &UpdateIacRepoUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateIacRepoUsingPUTParamsWithContext creates a new UpdateIacRepoUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateIacRepoUsingPUTParamsWithContext(ctx context.Context) *UpdateIacRepoUsingPUTParams {
	return &UpdateIacRepoUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateIacRepoUsingPUTParamsWithHTTPClient creates a new UpdateIacRepoUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIacRepoUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateIacRepoUsingPUTParams {
	return &UpdateIacRepoUsingPUTParams{
		HTTPClient: client,
	}
}

/*
UpdateIacRepoUsingPUTParams contains all the parameters to send to the API endpoint

	for the update iac repo using p u t operation.

	Typically these are written to a http.Request.
*/
type UpdateIacRepoUsingPUTParams struct {

	/* ID.

	   id
	*/
	ID string

	/* UpdatedIacRepo.

	   updatedIacRepo
	*/
	UpdatedIacRepo *models.IacRepo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update iac repo using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIacRepoUsingPUTParams) WithDefaults() *UpdateIacRepoUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update iac repo using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIacRepoUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateIacRepoUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) WithContext(ctx context.Context) *UpdateIacRepoUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateIacRepoUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) WithID(id string) *UpdateIacRepoUsingPUTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) SetID(id string) {
	o.ID = id
}

// WithUpdatedIacRepo adds the updatedIacRepo to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) WithUpdatedIacRepo(updatedIacRepo *models.IacRepo) *UpdateIacRepoUsingPUTParams {
	o.SetUpdatedIacRepo(updatedIacRepo)
	return o
}

// SetUpdatedIacRepo adds the updatedIacRepo to the update iac repo using p u t params
func (o *UpdateIacRepoUsingPUTParams) SetUpdatedIacRepo(updatedIacRepo *models.IacRepo) {
	o.UpdatedIacRepo = updatedIacRepo
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIacRepoUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.UpdatedIacRepo != nil {
		if err := r.SetBodyParam(o.UpdatedIacRepo); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
