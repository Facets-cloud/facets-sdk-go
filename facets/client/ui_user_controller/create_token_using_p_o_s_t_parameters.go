// Code generated by go-swagger; DO NOT EDIT.

package ui_user_controller

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

// NewCreateTokenUsingPOSTParams creates a new CreateTokenUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTokenUsingPOSTParams() *CreateTokenUsingPOSTParams {
	return &CreateTokenUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTokenUsingPOSTParamsWithTimeout creates a new CreateTokenUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateTokenUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateTokenUsingPOSTParams {
	return &CreateTokenUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateTokenUsingPOSTParamsWithContext creates a new CreateTokenUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateTokenUsingPOSTParamsWithContext(ctx context.Context) *CreateTokenUsingPOSTParams {
	return &CreateTokenUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateTokenUsingPOSTParamsWithHTTPClient creates a new CreateTokenUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTokenUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateTokenUsingPOSTParams {
	return &CreateTokenUsingPOSTParams{
		HTTPClient: client,
	}
}

/*
CreateTokenUsingPOSTParams contains all the parameters to send to the API endpoint

	for the create token using p o s t operation.

	Typically these are written to a http.Request.
*/
type CreateTokenUsingPOSTParams struct {

	// CreatedOn.
	//
	// Format: date-time
	CreatedOn *strfmt.DateTime

	// Description.
	Description *string

	// Name.
	Name *string

	// Token.
	Token *string

	// TokenID.
	TokenID *string

	// UserName.
	UserName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create token using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTokenUsingPOSTParams) WithDefaults() *CreateTokenUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create token using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTokenUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateTokenUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithContext(ctx context.Context) *CreateTokenUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateTokenUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedOn adds the createdOn to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithCreatedOn(createdOn *strfmt.DateTime) *CreateTokenUsingPOSTParams {
	o.SetCreatedOn(createdOn)
	return o
}

// SetCreatedOn adds the createdOn to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetCreatedOn(createdOn *strfmt.DateTime) {
	o.CreatedOn = createdOn
}

// WithDescription adds the description to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithDescription(description *string) *CreateTokenUsingPOSTParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetDescription(description *string) {
	o.Description = description
}

// WithName adds the name to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithName(name *string) *CreateTokenUsingPOSTParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetName(name *string) {
	o.Name = name
}

// WithToken adds the token to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithToken(token *string) *CreateTokenUsingPOSTParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetToken(token *string) {
	o.Token = token
}

// WithTokenID adds the tokenID to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithTokenID(tokenID *string) *CreateTokenUsingPOSTParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetTokenID(tokenID *string) {
	o.TokenID = tokenID
}

// WithUserName adds the userName to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) WithUserName(userName *string) *CreateTokenUsingPOSTParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the create token using p o s t params
func (o *CreateTokenUsingPOSTParams) SetUserName(userName *string) {
	o.UserName = userName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTokenUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatedOn != nil {

		// query param createdOn
		var qrCreatedOn strfmt.DateTime

		if o.CreatedOn != nil {
			qrCreatedOn = *o.CreatedOn
		}
		qCreatedOn := qrCreatedOn.String()
		if qCreatedOn != "" {

			if err := r.SetQueryParam("createdOn", qCreatedOn); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Token != nil {

		// query param token
		var qrToken string

		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {

			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}
	}

	if o.TokenID != nil {

		// query param tokenId
		var qrTokenID string

		if o.TokenID != nil {
			qrTokenID = *o.TokenID
		}
		qTokenID := qrTokenID
		if qTokenID != "" {

			if err := r.SetQueryParam("tokenId", qTokenID); err != nil {
				return err
			}
		}
	}

	if o.UserName != nil {

		// query param userName
		var qrUserName string

		if o.UserName != nil {
			qrUserName = *o.UserName
		}
		qUserName := qrUserName
		if qUserName != "" {

			if err := r.SetQueryParam("userName", qUserName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
