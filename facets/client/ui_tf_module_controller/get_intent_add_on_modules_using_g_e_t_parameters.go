// Code generated by go-swagger; DO NOT EDIT.

package ui_tf_module_controller

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

// NewGetIntentAddOnModulesUsingGETParams creates a new GetIntentAddOnModulesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntentAddOnModulesUsingGETParams() *GetIntentAddOnModulesUsingGETParams {
	return &GetIntentAddOnModulesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntentAddOnModulesUsingGETParamsWithTimeout creates a new GetIntentAddOnModulesUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetIntentAddOnModulesUsingGETParamsWithTimeout(timeout time.Duration) *GetIntentAddOnModulesUsingGETParams {
	return &GetIntentAddOnModulesUsingGETParams{
		timeout: timeout,
	}
}

// NewGetIntentAddOnModulesUsingGETParamsWithContext creates a new GetIntentAddOnModulesUsingGETParams object
// with the ability to set a context for a request.
func NewGetIntentAddOnModulesUsingGETParamsWithContext(ctx context.Context) *GetIntentAddOnModulesUsingGETParams {
	return &GetIntentAddOnModulesUsingGETParams{
		Context: ctx,
	}
}

// NewGetIntentAddOnModulesUsingGETParamsWithHTTPClient creates a new GetIntentAddOnModulesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntentAddOnModulesUsingGETParamsWithHTTPClient(client *http.Client) *GetIntentAddOnModulesUsingGETParams {
	return &GetIntentAddOnModulesUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetIntentAddOnModulesUsingGETParams contains all the parameters to send to the API endpoint

	for the get intent add on modules using g e t operation.

	Typically these are written to a http.Request.
*/
type GetIntentAddOnModulesUsingGETParams struct {

	/* Cloud.

	   cloud
	*/
	Cloud *string

	/* Flavor.

	   flavor
	*/
	Flavor string

	/* Intent.

	   intent
	*/
	Intent string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get intent add on modules using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntentAddOnModulesUsingGETParams) WithDefaults() *GetIntentAddOnModulesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get intent add on modules using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntentAddOnModulesUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) WithTimeout(timeout time.Duration) *GetIntentAddOnModulesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) WithContext(ctx context.Context) *GetIntentAddOnModulesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) WithHTTPClient(client *http.Client) *GetIntentAddOnModulesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloud adds the cloud to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) WithCloud(cloud *string) *GetIntentAddOnModulesUsingGETParams {
	o.SetCloud(cloud)
	return o
}

// SetCloud adds the cloud to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) SetCloud(cloud *string) {
	o.Cloud = cloud
}

// WithFlavor adds the flavor to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) WithFlavor(flavor string) *GetIntentAddOnModulesUsingGETParams {
	o.SetFlavor(flavor)
	return o
}

// SetFlavor adds the flavor to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) SetFlavor(flavor string) {
	o.Flavor = flavor
}

// WithIntent adds the intent to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) WithIntent(intent string) *GetIntentAddOnModulesUsingGETParams {
	o.SetIntent(intent)
	return o
}

// SetIntent adds the intent to the get intent add on modules using g e t params
func (o *GetIntentAddOnModulesUsingGETParams) SetIntent(intent string) {
	o.Intent = intent
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntentAddOnModulesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cloud != nil {

		// query param cloud
		var qrCloud string

		if o.Cloud != nil {
			qrCloud = *o.Cloud
		}
		qCloud := qrCloud
		if qCloud != "" {

			if err := r.SetQueryParam("cloud", qCloud); err != nil {
				return err
			}
		}
	}

	// path param flavor
	if err := r.SetPathParam("flavor", o.Flavor); err != nil {
		return err
	}

	// path param intent
	if err := r.SetPathParam("intent", o.Intent); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
