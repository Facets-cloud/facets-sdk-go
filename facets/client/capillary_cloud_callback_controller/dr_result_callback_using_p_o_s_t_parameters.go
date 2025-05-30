// Code generated by go-swagger; DO NOT EDIT.

package capillary_cloud_callback_controller

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

// NewDrResultCallbackUsingPOSTParams creates a new DrResultCallbackUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDrResultCallbackUsingPOSTParams() *DrResultCallbackUsingPOSTParams {
	return &DrResultCallbackUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDrResultCallbackUsingPOSTParamsWithTimeout creates a new DrResultCallbackUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewDrResultCallbackUsingPOSTParamsWithTimeout(timeout time.Duration) *DrResultCallbackUsingPOSTParams {
	return &DrResultCallbackUsingPOSTParams{
		timeout: timeout,
	}
}

// NewDrResultCallbackUsingPOSTParamsWithContext creates a new DrResultCallbackUsingPOSTParams object
// with the ability to set a context for a request.
func NewDrResultCallbackUsingPOSTParamsWithContext(ctx context.Context) *DrResultCallbackUsingPOSTParams {
	return &DrResultCallbackUsingPOSTParams{
		Context: ctx,
	}
}

// NewDrResultCallbackUsingPOSTParamsWithHTTPClient creates a new DrResultCallbackUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewDrResultCallbackUsingPOSTParamsWithHTTPClient(client *http.Client) *DrResultCallbackUsingPOSTParams {
	return &DrResultCallbackUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
DrResultCallbackUsingPOSTParams contains all the parameters to send to the API endpoint

	for the dr result callback using p o s t operation.

	Typically these are written to a http.Request.
*/
type DrResultCallbackUsingPOSTParams struct {

	/* Callback.

	   callback
	*/
	Callback *models.DRResult

	/* Cluster.

	   cluster
	*/
	Cluster string

	/* InstanceName.

	   instanceName
	*/
	InstanceName string

	/* ModuleType.

	   moduleType
	*/
	ModuleType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dr result callback using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DrResultCallbackUsingPOSTParams) WithDefaults() *DrResultCallbackUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dr result callback using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DrResultCallbackUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) WithTimeout(timeout time.Duration) *DrResultCallbackUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) WithContext(ctx context.Context) *DrResultCallbackUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) WithHTTPClient(client *http.Client) *DrResultCallbackUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallback adds the callback to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) WithCallback(callback *models.DRResult) *DrResultCallbackUsingPOSTParams {
	o.SetCallback(callback)
	return o
}

// SetCallback adds the callback to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) SetCallback(callback *models.DRResult) {
	o.Callback = callback
}

// WithCluster adds the cluster to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) WithCluster(cluster string) *DrResultCallbackUsingPOSTParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) SetCluster(cluster string) {
	o.Cluster = cluster
}

// WithInstanceName adds the instanceName to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) WithInstanceName(instanceName string) *DrResultCallbackUsingPOSTParams {
	o.SetInstanceName(instanceName)
	return o
}

// SetInstanceName adds the instanceName to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) SetInstanceName(instanceName string) {
	o.InstanceName = instanceName
}

// WithModuleType adds the moduleType to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) WithModuleType(moduleType string) *DrResultCallbackUsingPOSTParams {
	o.SetModuleType(moduleType)
	return o
}

// SetModuleType adds the moduleType to the dr result callback using p o s t params
func (o *DrResultCallbackUsingPOSTParams) SetModuleType(moduleType string) {
	o.ModuleType = moduleType
}

// WriteToRequest writes these params to a swagger request
func (o *DrResultCallbackUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Callback != nil {
		if err := r.SetBodyParam(o.Callback); err != nil {
			return err
		}
	}

	// path param cluster
	if err := r.SetPathParam("cluster", o.Cluster); err != nil {
		return err
	}

	// path param instanceName
	if err := r.SetPathParam("instanceName", o.InstanceName); err != nil {
		return err
	}

	// path param moduleType
	if err := r.SetPathParam("moduleType", o.ModuleType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
