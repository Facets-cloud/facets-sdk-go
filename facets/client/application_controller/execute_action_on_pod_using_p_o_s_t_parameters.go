// Code generated by go-swagger; DO NOT EDIT.

package application_controller

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

// NewExecuteActionOnPodUsingPOSTParams creates a new ExecuteActionOnPodUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExecuteActionOnPodUsingPOSTParams() *ExecuteActionOnPodUsingPOSTParams {
	return &ExecuteActionOnPodUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteActionOnPodUsingPOSTParamsWithTimeout creates a new ExecuteActionOnPodUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewExecuteActionOnPodUsingPOSTParamsWithTimeout(timeout time.Duration) *ExecuteActionOnPodUsingPOSTParams {
	return &ExecuteActionOnPodUsingPOSTParams{
		timeout: timeout,
	}
}

// NewExecuteActionOnPodUsingPOSTParamsWithContext creates a new ExecuteActionOnPodUsingPOSTParams object
// with the ability to set a context for a request.
func NewExecuteActionOnPodUsingPOSTParamsWithContext(ctx context.Context) *ExecuteActionOnPodUsingPOSTParams {
	return &ExecuteActionOnPodUsingPOSTParams{
		Context: ctx,
	}
}

// NewExecuteActionOnPodUsingPOSTParamsWithHTTPClient creates a new ExecuteActionOnPodUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewExecuteActionOnPodUsingPOSTParamsWithHTTPClient(client *http.Client) *ExecuteActionOnPodUsingPOSTParams {
	return &ExecuteActionOnPodUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
ExecuteActionOnPodUsingPOSTParams contains all the parameters to send to the API endpoint

	for the execute action on pod using p o s t operation.

	Typically these are written to a http.Request.
*/
type ExecuteActionOnPodUsingPOSTParams struct {

	/* ApplicationAction.

	   applicationAction
	*/
	ApplicationAction *models.ApplicationAction

	/* ApplicationFamily.

	   applicationFamily
	*/
	ApplicationFamily string

	/* ApplicationID.

	   applicationId
	*/
	ApplicationID string

	/* Environment.

	   environment
	*/
	Environment string

	/* PodName.

	   podName
	*/
	PodName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the execute action on pod using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteActionOnPodUsingPOSTParams) WithDefaults() *ExecuteActionOnPodUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the execute action on pod using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExecuteActionOnPodUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) WithTimeout(timeout time.Duration) *ExecuteActionOnPodUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) WithContext(ctx context.Context) *ExecuteActionOnPodUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) WithHTTPClient(client *http.Client) *ExecuteActionOnPodUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationAction adds the applicationAction to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) WithApplicationAction(applicationAction *models.ApplicationAction) *ExecuteActionOnPodUsingPOSTParams {
	o.SetApplicationAction(applicationAction)
	return o
}

// SetApplicationAction adds the applicationAction to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) SetApplicationAction(applicationAction *models.ApplicationAction) {
	o.ApplicationAction = applicationAction
}

// WithApplicationFamily adds the applicationFamily to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) WithApplicationFamily(applicationFamily string) *ExecuteActionOnPodUsingPOSTParams {
	o.SetApplicationFamily(applicationFamily)
	return o
}

// SetApplicationFamily adds the applicationFamily to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) SetApplicationFamily(applicationFamily string) {
	o.ApplicationFamily = applicationFamily
}

// WithApplicationID adds the applicationID to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) WithApplicationID(applicationID string) *ExecuteActionOnPodUsingPOSTParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) SetApplicationID(applicationID string) {
	o.ApplicationID = applicationID
}

// WithEnvironment adds the environment to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) WithEnvironment(environment string) *ExecuteActionOnPodUsingPOSTParams {
	o.SetEnvironment(environment)
	return o
}

// SetEnvironment adds the environment to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) SetEnvironment(environment string) {
	o.Environment = environment
}

// WithPodName adds the podName to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) WithPodName(podName string) *ExecuteActionOnPodUsingPOSTParams {
	o.SetPodName(podName)
	return o
}

// SetPodName adds the podName to the execute action on pod using p o s t params
func (o *ExecuteActionOnPodUsingPOSTParams) SetPodName(podName string) {
	o.PodName = podName
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteActionOnPodUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ApplicationAction != nil {
		if err := r.SetBodyParam(o.ApplicationAction); err != nil {
			return err
		}
	}

	// path param applicationFamily
	if err := r.SetPathParam("applicationFamily", o.ApplicationFamily); err != nil {
		return err
	}

	// path param applicationId
	if err := r.SetPathParam("applicationId", o.ApplicationID); err != nil {
		return err
	}

	// path param environment
	if err := r.SetPathParam("environment", o.Environment); err != nil {
		return err
	}

	// path param podName
	if err := r.SetPathParam("podName", o.PodName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
