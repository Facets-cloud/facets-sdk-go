// Code generated by go-swagger; DO NOT EDIT.

package ui_coder_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

// CreateWorkspaceWithExistingBranchUsingPOSTReader is a Reader for the CreateWorkspaceWithExistingBranchUsingPOST structure.
type CreateWorkspaceWithExistingBranchUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWorkspaceWithExistingBranchUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateWorkspaceWithExistingBranchUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateWorkspaceWithExistingBranchUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateWorkspaceWithExistingBranchUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateWorkspaceWithExistingBranchUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateWorkspaceWithExistingBranchUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch] createWorkspaceWithExistingBranchUsingPOST", response, response.Code())
	}
}

// NewCreateWorkspaceWithExistingBranchUsingPOSTOK creates a CreateWorkspaceWithExistingBranchUsingPOSTOK with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOSTOK() *CreateWorkspaceWithExistingBranchUsingPOSTOK {
	return &CreateWorkspaceWithExistingBranchUsingPOSTOK{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type CreateWorkspaceWithExistingBranchUsingPOSTOK struct {
	Payload *models.CoderWorkspaceResponse
}

// IsSuccess returns true when this create workspace with existing branch using p o s t o k response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create workspace with existing branch using p o s t o k response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t o k response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create workspace with existing branch using p o s t o k response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t o k response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create workspace with existing branch using p o s t o k response
func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) Code() int {
	return 200
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTOK %s", 200, payload)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTOK %s", 200, payload)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) GetPayload() *models.CoderWorkspaceResponse {
	return o.Payload
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoderWorkspaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkspaceWithExistingBranchUsingPOSTCreated creates a CreateWorkspaceWithExistingBranchUsingPOSTCreated with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOSTCreated() *CreateWorkspaceWithExistingBranchUsingPOSTCreated {
	return &CreateWorkspaceWithExistingBranchUsingPOSTCreated{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type CreateWorkspaceWithExistingBranchUsingPOSTCreated struct {
}

// IsSuccess returns true when this create workspace with existing branch using p o s t created response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create workspace with existing branch using p o s t created response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t created response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create workspace with existing branch using p o s t created response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t created response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create workspace with existing branch using p o s t created response
func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) Code() int {
	return 201
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTCreated", 201)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTCreated", 201)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkspaceWithExistingBranchUsingPOSTUnauthorized creates a CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOSTUnauthorized() *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized {
	return &CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this create workspace with existing branch using p o s t unauthorized response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create workspace with existing branch using p o s t unauthorized response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t unauthorized response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create workspace with existing branch using p o s t unauthorized response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t unauthorized response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create workspace with existing branch using p o s t unauthorized response
func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTUnauthorized", 401)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTUnauthorized", 401)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkspaceWithExistingBranchUsingPOSTForbidden creates a CreateWorkspaceWithExistingBranchUsingPOSTForbidden with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOSTForbidden() *CreateWorkspaceWithExistingBranchUsingPOSTForbidden {
	return &CreateWorkspaceWithExistingBranchUsingPOSTForbidden{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateWorkspaceWithExistingBranchUsingPOSTForbidden struct {
}

// IsSuccess returns true when this create workspace with existing branch using p o s t forbidden response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create workspace with existing branch using p o s t forbidden response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t forbidden response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create workspace with existing branch using p o s t forbidden response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t forbidden response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create workspace with existing branch using p o s t forbidden response
func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) Code() int {
	return 403
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTForbidden", 403)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTForbidden", 403)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkspaceWithExistingBranchUsingPOSTNotFound creates a CreateWorkspaceWithExistingBranchUsingPOSTNotFound with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOSTNotFound() *CreateWorkspaceWithExistingBranchUsingPOSTNotFound {
	return &CreateWorkspaceWithExistingBranchUsingPOSTNotFound{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateWorkspaceWithExistingBranchUsingPOSTNotFound struct {
}

// IsSuccess returns true when this create workspace with existing branch using p o s t not found response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create workspace with existing branch using p o s t not found response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t not found response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create workspace with existing branch using p o s t not found response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t not found response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create workspace with existing branch using p o s t not found response
func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) Code() int {
	return 404
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTNotFound", 404)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOSTNotFound", 404)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
