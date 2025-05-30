// Code generated by go-swagger; DO NOT EDIT.

package ui_blueprint_designer_controller

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

// NewBulkEditDisabledForResourcesUsingPUT1Params creates a new BulkEditDisabledForResourcesUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkEditDisabledForResourcesUsingPUT1Params() *BulkEditDisabledForResourcesUsingPUT1Params {
	return &BulkEditDisabledForResourcesUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkEditDisabledForResourcesUsingPUT1ParamsWithTimeout creates a new BulkEditDisabledForResourcesUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewBulkEditDisabledForResourcesUsingPUT1ParamsWithTimeout(timeout time.Duration) *BulkEditDisabledForResourcesUsingPUT1Params {
	return &BulkEditDisabledForResourcesUsingPUT1Params{
		timeout: timeout,
	}
}

// NewBulkEditDisabledForResourcesUsingPUT1ParamsWithContext creates a new BulkEditDisabledForResourcesUsingPUT1Params object
// with the ability to set a context for a request.
func NewBulkEditDisabledForResourcesUsingPUT1ParamsWithContext(ctx context.Context) *BulkEditDisabledForResourcesUsingPUT1Params {
	return &BulkEditDisabledForResourcesUsingPUT1Params{
		Context: ctx,
	}
}

// NewBulkEditDisabledForResourcesUsingPUT1ParamsWithHTTPClient creates a new BulkEditDisabledForResourcesUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewBulkEditDisabledForResourcesUsingPUT1ParamsWithHTTPClient(client *http.Client) *BulkEditDisabledForResourcesUsingPUT1Params {
	return &BulkEditDisabledForResourcesUsingPUT1Params{
		HTTPClient: client,
	}
}

/*
BulkEditDisabledForResourcesUsingPUT1Params contains all the parameters to send to the API endpoint

	for the bulk edit disabled for resources using p u t 1 operation.

	Typically these are written to a http.Request.
*/
type BulkEditDisabledForResourcesUsingPUT1Params struct {

	/* ResourceEnableDisableRequestList.

	   resourceEnableDisableRequestList
	*/
	ResourceEnableDisableRequestList []*models.ResourceEnableDisableRequest

	/* StackName.

	   stackName
	*/
	StackName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk edit disabled for resources using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkEditDisabledForResourcesUsingPUT1Params) WithDefaults() *BulkEditDisabledForResourcesUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk edit disabled for resources using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkEditDisabledForResourcesUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) WithTimeout(timeout time.Duration) *BulkEditDisabledForResourcesUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) WithContext(ctx context.Context) *BulkEditDisabledForResourcesUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) WithHTTPClient(client *http.Client) *BulkEditDisabledForResourcesUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceEnableDisableRequestList adds the resourceEnableDisableRequestList to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) WithResourceEnableDisableRequestList(resourceEnableDisableRequestList []*models.ResourceEnableDisableRequest) *BulkEditDisabledForResourcesUsingPUT1Params {
	o.SetResourceEnableDisableRequestList(resourceEnableDisableRequestList)
	return o
}

// SetResourceEnableDisableRequestList adds the resourceEnableDisableRequestList to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) SetResourceEnableDisableRequestList(resourceEnableDisableRequestList []*models.ResourceEnableDisableRequest) {
	o.ResourceEnableDisableRequestList = resourceEnableDisableRequestList
}

// WithStackName adds the stackName to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) WithStackName(stackName string) *BulkEditDisabledForResourcesUsingPUT1Params {
	o.SetStackName(stackName)
	return o
}

// SetStackName adds the stackName to the bulk edit disabled for resources using p u t 1 params
func (o *BulkEditDisabledForResourcesUsingPUT1Params) SetStackName(stackName string) {
	o.StackName = stackName
}

// WriteToRequest writes these params to a swagger request
func (o *BulkEditDisabledForResourcesUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ResourceEnableDisableRequestList != nil {
		if err := r.SetBodyParam(o.ResourceEnableDisableRequestList); err != nil {
			return err
		}
	}

	// path param stackName
	if err := r.SetPathParam("stackName", o.StackName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
