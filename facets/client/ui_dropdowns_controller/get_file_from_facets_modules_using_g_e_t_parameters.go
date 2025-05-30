// Code generated by go-swagger; DO NOT EDIT.

package ui_dropdowns_controller

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

// NewGetFileFromFacetsModulesUsingGETParams creates a new GetFileFromFacetsModulesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFileFromFacetsModulesUsingGETParams() *GetFileFromFacetsModulesUsingGETParams {
	return &GetFileFromFacetsModulesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFileFromFacetsModulesUsingGETParamsWithTimeout creates a new GetFileFromFacetsModulesUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetFileFromFacetsModulesUsingGETParamsWithTimeout(timeout time.Duration) *GetFileFromFacetsModulesUsingGETParams {
	return &GetFileFromFacetsModulesUsingGETParams{
		timeout: timeout,
	}
}

// NewGetFileFromFacetsModulesUsingGETParamsWithContext creates a new GetFileFromFacetsModulesUsingGETParams object
// with the ability to set a context for a request.
func NewGetFileFromFacetsModulesUsingGETParamsWithContext(ctx context.Context) *GetFileFromFacetsModulesUsingGETParams {
	return &GetFileFromFacetsModulesUsingGETParams{
		Context: ctx,
	}
}

// NewGetFileFromFacetsModulesUsingGETParamsWithHTTPClient creates a new GetFileFromFacetsModulesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFileFromFacetsModulesUsingGETParamsWithHTTPClient(client *http.Client) *GetFileFromFacetsModulesUsingGETParams {
	return &GetFileFromFacetsModulesUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetFileFromFacetsModulesUsingGETParams contains all the parameters to send to the API endpoint

	for the get file from facets modules using g e t operation.

	Typically these are written to a http.Request.
*/
type GetFileFromFacetsModulesUsingGETParams struct {

	/* FileName.

	   fileName
	*/
	FileName string

	/* Path.

	   path
	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get file from facets modules using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileFromFacetsModulesUsingGETParams) WithDefaults() *GetFileFromFacetsModulesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get file from facets modules using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFileFromFacetsModulesUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) WithTimeout(timeout time.Duration) *GetFileFromFacetsModulesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) WithContext(ctx context.Context) *GetFileFromFacetsModulesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) WithHTTPClient(client *http.Client) *GetFileFromFacetsModulesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileName adds the fileName to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) WithFileName(fileName string) *GetFileFromFacetsModulesUsingGETParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithPath adds the path to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) WithPath(path string) *GetFileFromFacetsModulesUsingGETParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the get file from facets modules using g e t params
func (o *GetFileFromFacetsModulesUsingGETParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *GetFileFromFacetsModulesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param fileName
	qrFileName := o.FileName
	qFileName := qrFileName
	if qFileName != "" {

		if err := r.SetQueryParam("fileName", qFileName); err != nil {
			return err
		}
	}

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
