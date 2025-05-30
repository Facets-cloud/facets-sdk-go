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

// NewCreateBitBucketAccountUsingPOSTParams creates a new CreateBitBucketAccountUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBitBucketAccountUsingPOSTParams() *CreateBitBucketAccountUsingPOSTParams {
	return &CreateBitBucketAccountUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBitBucketAccountUsingPOSTParamsWithTimeout creates a new CreateBitBucketAccountUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateBitBucketAccountUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateBitBucketAccountUsingPOSTParams {
	return &CreateBitBucketAccountUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateBitBucketAccountUsingPOSTParamsWithContext creates a new CreateBitBucketAccountUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateBitBucketAccountUsingPOSTParamsWithContext(ctx context.Context) *CreateBitBucketAccountUsingPOSTParams {
	return &CreateBitBucketAccountUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateBitBucketAccountUsingPOSTParamsWithHTTPClient creates a new CreateBitBucketAccountUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBitBucketAccountUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateBitBucketAccountUsingPOSTParams {
	return &CreateBitBucketAccountUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CreateBitBucketAccountUsingPOSTParams contains all the parameters to send to the API endpoint

	for the create bit bucket account using p o s t operation.

	Typically these are written to a http.Request.
*/
type CreateBitBucketAccountUsingPOSTParams struct {

	/* BitbucketAccountRequest.

	   bitbucketAccountRequest
	*/
	BitbucketAccountRequest *models.BitBucketAccount

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create bit bucket account using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBitBucketAccountUsingPOSTParams) WithDefaults() *CreateBitBucketAccountUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create bit bucket account using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBitBucketAccountUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create bit bucket account using p o s t params
func (o *CreateBitBucketAccountUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateBitBucketAccountUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create bit bucket account using p o s t params
func (o *CreateBitBucketAccountUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create bit bucket account using p o s t params
func (o *CreateBitBucketAccountUsingPOSTParams) WithContext(ctx context.Context) *CreateBitBucketAccountUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create bit bucket account using p o s t params
func (o *CreateBitBucketAccountUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create bit bucket account using p o s t params
func (o *CreateBitBucketAccountUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateBitBucketAccountUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create bit bucket account using p o s t params
func (o *CreateBitBucketAccountUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBitbucketAccountRequest adds the bitbucketAccountRequest to the create bit bucket account using p o s t params
func (o *CreateBitBucketAccountUsingPOSTParams) WithBitbucketAccountRequest(bitbucketAccountRequest *models.BitBucketAccount) *CreateBitBucketAccountUsingPOSTParams {
	o.SetBitbucketAccountRequest(bitbucketAccountRequest)
	return o
}

// SetBitbucketAccountRequest adds the bitbucketAccountRequest to the create bit bucket account using p o s t params
func (o *CreateBitBucketAccountUsingPOSTParams) SetBitbucketAccountRequest(bitbucketAccountRequest *models.BitBucketAccount) {
	o.BitbucketAccountRequest = bitbucketAccountRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBitBucketAccountUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.BitbucketAccountRequest != nil {
		if err := r.SetBodyParam(o.BitbucketAccountRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
