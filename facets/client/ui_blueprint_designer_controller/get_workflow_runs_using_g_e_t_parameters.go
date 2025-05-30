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
	"github.com/go-openapi/swag"
)

// NewGetWorkflowRunsUsingGETParams creates a new GetWorkflowRunsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowRunsUsingGETParams() *GetWorkflowRunsUsingGETParams {
	return &GetWorkflowRunsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowRunsUsingGETParamsWithTimeout creates a new GetWorkflowRunsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowRunsUsingGETParamsWithTimeout(timeout time.Duration) *GetWorkflowRunsUsingGETParams {
	return &GetWorkflowRunsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetWorkflowRunsUsingGETParamsWithContext creates a new GetWorkflowRunsUsingGETParams object
// with the ability to set a context for a request.
func NewGetWorkflowRunsUsingGETParamsWithContext(ctx context.Context) *GetWorkflowRunsUsingGETParams {
	return &GetWorkflowRunsUsingGETParams{
		Context: ctx,
	}
}

// NewGetWorkflowRunsUsingGETParamsWithHTTPClient creates a new GetWorkflowRunsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowRunsUsingGETParamsWithHTTPClient(client *http.Client) *GetWorkflowRunsUsingGETParams {
	return &GetWorkflowRunsUsingGETParams{
		HTTPClient: client,
	}
}

/*
GetWorkflowRunsUsingGETParams contains all the parameters to send to the API endpoint

	for the get workflow runs using g e t operation.

	Typically these are written to a http.Request.
*/
type GetWorkflowRunsUsingGETParams struct {

	/* Actor.

	   actor
	*/
	Actor *string

	/* Branch.

	   branch
	*/
	Branch *string

	/* Event.

	   event
	*/
	Event *string

	/* PageNumber.

	   pageNumber

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   pageSize

	   Format: int32
	   Default: 10
	*/
	PageSize *int32

	/* ResourceName.

	   resourceName
	*/
	ResourceName string

	/* ResourceType.

	   resourceType
	*/
	ResourceType string

	/* StackName.

	   stackName
	*/
	StackName string

	/* Status.

	   status
	*/
	Status *string

	/* WorkflowID.

	   workflowId
	*/
	WorkflowID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow runs using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowRunsUsingGETParams) WithDefaults() *GetWorkflowRunsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow runs using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowRunsUsingGETParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(10)
	)

	val := GetWorkflowRunsUsingGETParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithTimeout(timeout time.Duration) *GetWorkflowRunsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithContext(ctx context.Context) *GetWorkflowRunsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithHTTPClient(client *http.Client) *GetWorkflowRunsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActor adds the actor to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithActor(actor *string) *GetWorkflowRunsUsingGETParams {
	o.SetActor(actor)
	return o
}

// SetActor adds the actor to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetActor(actor *string) {
	o.Actor = actor
}

// WithBranch adds the branch to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithBranch(branch *string) *GetWorkflowRunsUsingGETParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetBranch(branch *string) {
	o.Branch = branch
}

// WithEvent adds the event to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithEvent(event *string) *GetWorkflowRunsUsingGETParams {
	o.SetEvent(event)
	return o
}

// SetEvent adds the event to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetEvent(event *string) {
	o.Event = event
}

// WithPageNumber adds the pageNumber to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithPageNumber(pageNumber *int32) *GetWorkflowRunsUsingGETParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithPageSize(pageSize *int32) *GetWorkflowRunsUsingGETParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithResourceName adds the resourceName to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithResourceName(resourceName string) *GetWorkflowRunsUsingGETParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WithResourceType adds the resourceType to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithResourceType(resourceType string) *GetWorkflowRunsUsingGETParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithStackName adds the stackName to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithStackName(stackName string) *GetWorkflowRunsUsingGETParams {
	o.SetStackName(stackName)
	return o
}

// SetStackName adds the stackName to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetStackName(stackName string) {
	o.StackName = stackName
}

// WithStatus adds the status to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithStatus(status *string) *GetWorkflowRunsUsingGETParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetStatus(status *string) {
	o.Status = status
}

// WithWorkflowID adds the workflowID to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) WithWorkflowID(workflowID *string) *GetWorkflowRunsUsingGETParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the get workflow runs using g e t params
func (o *GetWorkflowRunsUsingGETParams) SetWorkflowID(workflowID *string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowRunsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Actor != nil {

		// query param actor
		var qrActor string

		if o.Actor != nil {
			qrActor = *o.Actor
		}
		qActor := qrActor
		if qActor != "" {

			if err := r.SetQueryParam("actor", qActor); err != nil {
				return err
			}
		}
	}

	if o.Branch != nil {

		// query param branch
		var qrBranch string

		if o.Branch != nil {
			qrBranch = *o.Branch
		}
		qBranch := qrBranch
		if qBranch != "" {

			if err := r.SetQueryParam("branch", qBranch); err != nil {
				return err
			}
		}
	}

	if o.Event != nil {

		// query param event
		var qrEvent string

		if o.Event != nil {
			qrEvent = *o.Event
		}
		qEvent := qrEvent
		if qEvent != "" {

			if err := r.SetQueryParam("event", qEvent); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	// path param resourceName
	if err := r.SetPathParam("resourceName", o.ResourceName); err != nil {
		return err
	}

	// path param resourceType
	if err := r.SetPathParam("resourceType", o.ResourceType); err != nil {
		return err
	}

	// path param stackName
	if err := r.SetPathParam("stackName", o.StackName); err != nil {
		return err
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.WorkflowID != nil {

		// query param workflowId
		var qrWorkflowID string

		if o.WorkflowID != nil {
			qrWorkflowID = *o.WorkflowID
		}
		qWorkflowID := qrWorkflowID
		if qWorkflowID != "" {

			if err := r.SetQueryParam("workflowId", qWorkflowID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
