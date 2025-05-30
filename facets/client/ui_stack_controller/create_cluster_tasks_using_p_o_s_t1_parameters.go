// Code generated by go-swagger; DO NOT EDIT.

package ui_stack_controller

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

// NewCreateClusterTasksUsingPOST1Params creates a new CreateClusterTasksUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateClusterTasksUsingPOST1Params() *CreateClusterTasksUsingPOST1Params {
	return &CreateClusterTasksUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterTasksUsingPOST1ParamsWithTimeout creates a new CreateClusterTasksUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewCreateClusterTasksUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateClusterTasksUsingPOST1Params {
	return &CreateClusterTasksUsingPOST1Params{
		timeout: timeout,
	}
}

// NewCreateClusterTasksUsingPOST1ParamsWithContext creates a new CreateClusterTasksUsingPOST1Params object
// with the ability to set a context for a request.
func NewCreateClusterTasksUsingPOST1ParamsWithContext(ctx context.Context) *CreateClusterTasksUsingPOST1Params {
	return &CreateClusterTasksUsingPOST1Params{
		Context: ctx,
	}
}

// NewCreateClusterTasksUsingPOST1ParamsWithHTTPClient creates a new CreateClusterTasksUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateClusterTasksUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateClusterTasksUsingPOST1Params {
	return &CreateClusterTasksUsingPOST1Params{
		HTTPClient: client,
	}
}

/*
CreateClusterTasksUsingPOST1Params contains all the parameters to send to the API endpoint

	for the create cluster tasks using p o s t 1 operation.

	Typically these are written to a http.Request.
*/
type CreateClusterTasksUsingPOST1Params struct {

	/* TaskRequest.

	   taskRequest
	*/
	TaskRequest *models.ClusterTaskRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create cluster tasks using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterTasksUsingPOST1Params) WithDefaults() *CreateClusterTasksUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create cluster tasks using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateClusterTasksUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create cluster tasks using p o s t 1 params
func (o *CreateClusterTasksUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateClusterTasksUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster tasks using p o s t 1 params
func (o *CreateClusterTasksUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster tasks using p o s t 1 params
func (o *CreateClusterTasksUsingPOST1Params) WithContext(ctx context.Context) *CreateClusterTasksUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster tasks using p o s t 1 params
func (o *CreateClusterTasksUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster tasks using p o s t 1 params
func (o *CreateClusterTasksUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateClusterTasksUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster tasks using p o s t 1 params
func (o *CreateClusterTasksUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskRequest adds the taskRequest to the create cluster tasks using p o s t 1 params
func (o *CreateClusterTasksUsingPOST1Params) WithTaskRequest(taskRequest *models.ClusterTaskRequest) *CreateClusterTasksUsingPOST1Params {
	o.SetTaskRequest(taskRequest)
	return o
}

// SetTaskRequest adds the taskRequest to the create cluster tasks using p o s t 1 params
func (o *CreateClusterTasksUsingPOST1Params) SetTaskRequest(taskRequest *models.ClusterTaskRequest) {
	o.TaskRequest = taskRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterTasksUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.TaskRequest != nil {
		if err := r.SetBodyParam(o.TaskRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
