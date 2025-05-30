// Code generated by go-swagger; DO NOT EDIT.

package ui_resource_status_controller

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
	"github.com/go-openapi/swag"
)

// NewGetResourceStatusUsingGETParams creates a new GetResourceStatusUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceStatusUsingGETParams() *GetResourceStatusUsingGETParams {
	return &GetResourceStatusUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceStatusUsingGETParamsWithTimeout creates a new GetResourceStatusUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetResourceStatusUsingGETParamsWithTimeout(timeout time.Duration) *GetResourceStatusUsingGETParams {
	return &GetResourceStatusUsingGETParams{
		timeout: timeout,
	}
}

// NewGetResourceStatusUsingGETParamsWithContext creates a new GetResourceStatusUsingGETParams object
// with the ability to set a context for a request.
func NewGetResourceStatusUsingGETParamsWithContext(ctx context.Context) *GetResourceStatusUsingGETParams {
	return &GetResourceStatusUsingGETParams{
		Context: ctx,
	}
}

// NewGetResourceStatusUsingGETParamsWithHTTPClient creates a new GetResourceStatusUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceStatusUsingGETParamsWithHTTPClient(client *http.Client) *GetResourceStatusUsingGETParams {
	return &GetResourceStatusUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetResourceStatusUsingGETParams contains all the parameters to send to the API endpoint

	for the get resource status using g e t operation.

	Typically these are written to a http.Request.
*/
type GetResourceStatusUsingGETParams struct {

	/* EnvironmentName.

	   environmentName
	*/
	EnvironmentName string

	/* ProjectName.

	   projectName
	*/
	ProjectName string

	/* Resources.

	   resources
	*/
	Resources []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource status using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceStatusUsingGETParams) WithDefaults() *GetResourceStatusUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource status using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceStatusUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) WithTimeout(timeout time.Duration) *GetResourceStatusUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) WithContext(ctx context.Context) *GetResourceStatusUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) WithHTTPClient(client *http.Client) *GetResourceStatusUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentName adds the environmentName to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) WithEnvironmentName(environmentName string) *GetResourceStatusUsingGETParams {
	o.SetEnvironmentName(environmentName)
	return o
}

// SetEnvironmentName adds the environmentName to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) SetEnvironmentName(environmentName string) {
	o.EnvironmentName = environmentName
}

// WithProjectName adds the projectName to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) WithProjectName(projectName string) *GetResourceStatusUsingGETParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithResources adds the resources to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) WithResources(resources []string) *GetResourceStatusUsingGETParams {
	o.SetResources(resources)
	return o
}

// SetResources adds the resources to the get resource status using g e t params
func (o *GetResourceStatusUsingGETParams) SetResources(resources []string) {
	o.Resources = resources
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceStatusUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environmentName
	if err := r.SetPathParam("environmentName", o.EnvironmentName); err != nil {
		return err
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	if o.Resources != nil {

		// binding items for resources
		joinedResources := o.bindParamResources(reg)

		// query array param resources
		if err := r.SetQueryParam("resources", joinedResources...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetResourceStatusUsingGET binds the parameter resources
func (o *GetResourceStatusUsingGETParams) bindParamResources(formats strfmt.Registry) []string {
	resourcesIR := o.Resources

	var resourcesIC []string
	for _, resourcesIIR := range resourcesIR { // explode []string

		resourcesIIV := resourcesIIR // string as string
		resourcesIC = append(resourcesIC, resourcesIIV)
	}

	// items.CollectionFormat: ""
	resourcesIS := swag.JoinByFormat(resourcesIC, "")

	return resourcesIS
}
