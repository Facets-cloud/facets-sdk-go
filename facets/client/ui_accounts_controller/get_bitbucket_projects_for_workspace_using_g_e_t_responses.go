// Code generated by go-swagger; DO NOT EDIT.

package ui_accounts_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetBitbucketProjectsForWorkspaceUsingGETReader is a Reader for the GetBitbucketProjectsForWorkspaceUsingGET structure.
type GetBitbucketProjectsForWorkspaceUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBitbucketProjectsForWorkspaceUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBitbucketProjectsForWorkspaceUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBitbucketProjectsForWorkspaceUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBitbucketProjectsForWorkspaceUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBitbucketProjectsForWorkspaceUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects] getBitbucketProjectsForWorkspaceUsingGET", response, response.Code())
	}
}

// NewGetBitbucketProjectsForWorkspaceUsingGETOK creates a GetBitbucketProjectsForWorkspaceUsingGETOK with default headers values
func NewGetBitbucketProjectsForWorkspaceUsingGETOK() *GetBitbucketProjectsForWorkspaceUsingGETOK {
	return &GetBitbucketProjectsForWorkspaceUsingGETOK{}
}

/*
GetBitbucketProjectsForWorkspaceUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetBitbucketProjectsForWorkspaceUsingGETOK struct {
	Payload []string
}

// IsSuccess returns true when this get bitbucket projects for workspace using g e t o k response has a 2xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bitbucket projects for workspace using g e t o k response has a 3xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bitbucket projects for workspace using g e t o k response has a 4xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bitbucket projects for workspace using g e t o k response has a 5xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bitbucket projects for workspace using g e t o k response a status code equal to that given
func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bitbucket projects for workspace using g e t o k response
func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) Code() int {
	return 200
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects][%d] getBitbucketProjectsForWorkspaceUsingGETOK %s", 200, payload)
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects][%d] getBitbucketProjectsForWorkspaceUsingGETOK %s", 200, payload)
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) GetPayload() []string {
	return o.Payload
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBitbucketProjectsForWorkspaceUsingGETUnauthorized creates a GetBitbucketProjectsForWorkspaceUsingGETUnauthorized with default headers values
func NewGetBitbucketProjectsForWorkspaceUsingGETUnauthorized() *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized {
	return &GetBitbucketProjectsForWorkspaceUsingGETUnauthorized{}
}

/*
GetBitbucketProjectsForWorkspaceUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetBitbucketProjectsForWorkspaceUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get bitbucket projects for workspace using g e t unauthorized response has a 2xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bitbucket projects for workspace using g e t unauthorized response has a 3xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bitbucket projects for workspace using g e t unauthorized response has a 4xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bitbucket projects for workspace using g e t unauthorized response has a 5xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get bitbucket projects for workspace using g e t unauthorized response a status code equal to that given
func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get bitbucket projects for workspace using g e t unauthorized response
func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects][%d] getBitbucketProjectsForWorkspaceUsingGETUnauthorized", 401)
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects][%d] getBitbucketProjectsForWorkspaceUsingGETUnauthorized", 401)
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBitbucketProjectsForWorkspaceUsingGETForbidden creates a GetBitbucketProjectsForWorkspaceUsingGETForbidden with default headers values
func NewGetBitbucketProjectsForWorkspaceUsingGETForbidden() *GetBitbucketProjectsForWorkspaceUsingGETForbidden {
	return &GetBitbucketProjectsForWorkspaceUsingGETForbidden{}
}

/*
GetBitbucketProjectsForWorkspaceUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetBitbucketProjectsForWorkspaceUsingGETForbidden struct {
}

// IsSuccess returns true when this get bitbucket projects for workspace using g e t forbidden response has a 2xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bitbucket projects for workspace using g e t forbidden response has a 3xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bitbucket projects for workspace using g e t forbidden response has a 4xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bitbucket projects for workspace using g e t forbidden response has a 5xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get bitbucket projects for workspace using g e t forbidden response a status code equal to that given
func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get bitbucket projects for workspace using g e t forbidden response
func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) Code() int {
	return 403
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects][%d] getBitbucketProjectsForWorkspaceUsingGETForbidden", 403)
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects][%d] getBitbucketProjectsForWorkspaceUsingGETForbidden", 403)
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBitbucketProjectsForWorkspaceUsingGETNotFound creates a GetBitbucketProjectsForWorkspaceUsingGETNotFound with default headers values
func NewGetBitbucketProjectsForWorkspaceUsingGETNotFound() *GetBitbucketProjectsForWorkspaceUsingGETNotFound {
	return &GetBitbucketProjectsForWorkspaceUsingGETNotFound{}
}

/*
GetBitbucketProjectsForWorkspaceUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetBitbucketProjectsForWorkspaceUsingGETNotFound struct {
}

// IsSuccess returns true when this get bitbucket projects for workspace using g e t not found response has a 2xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bitbucket projects for workspace using g e t not found response has a 3xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bitbucket projects for workspace using g e t not found response has a 4xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bitbucket projects for workspace using g e t not found response has a 5xx status code
func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get bitbucket projects for workspace using g e t not found response a status code equal to that given
func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get bitbucket projects for workspace using g e t not found response
func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) Code() int {
	return 404
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects][%d] getBitbucketProjectsForWorkspaceUsingGETNotFound", 404)
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/accounts/{accountId}/workspaces/{workspace}/projects][%d] getBitbucketProjectsForWorkspaceUsingGETNotFound", 404)
}

func (o *GetBitbucketProjectsForWorkspaceUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
