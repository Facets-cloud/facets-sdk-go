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

// CreateWorkspaceWithExistingBranchUsingPOST1Reader is a Reader for the CreateWorkspaceWithExistingBranchUsingPOST1 structure.
type CreateWorkspaceWithExistingBranchUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateWorkspaceWithExistingBranchUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateWorkspaceWithExistingBranchUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateWorkspaceWithExistingBranchUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateWorkspaceWithExistingBranchUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateWorkspaceWithExistingBranchUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch] createWorkspaceWithExistingBranchUsingPOST_1", response, response.Code())
	}
}

// NewCreateWorkspaceWithExistingBranchUsingPOST1OK creates a CreateWorkspaceWithExistingBranchUsingPOST1OK with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOST1OK() *CreateWorkspaceWithExistingBranchUsingPOST1OK {
	return &CreateWorkspaceWithExistingBranchUsingPOST1OK{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOST1OK describes a response with status code 200, with default header values.

OK
*/
type CreateWorkspaceWithExistingBranchUsingPOST1OK struct {
	Payload *models.CoderWorkspaceResponse
}

// IsSuccess returns true when this create workspace with existing branch using p o s t1 o k response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create workspace with existing branch using p o s t1 o k response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t1 o k response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create workspace with existing branch using p o s t1 o k response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t1 o k response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create workspace with existing branch using p o s t1 o k response
func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) Code() int {
	return 200
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1OK %s", 200, payload)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1OK %s", 200, payload)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) GetPayload() *models.CoderWorkspaceResponse {
	return o.Payload
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoderWorkspaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkspaceWithExistingBranchUsingPOST1Created creates a CreateWorkspaceWithExistingBranchUsingPOST1Created with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOST1Created() *CreateWorkspaceWithExistingBranchUsingPOST1Created {
	return &CreateWorkspaceWithExistingBranchUsingPOST1Created{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOST1Created describes a response with status code 201, with default header values.

Created
*/
type CreateWorkspaceWithExistingBranchUsingPOST1Created struct {
}

// IsSuccess returns true when this create workspace with existing branch using p o s t1 created response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create workspace with existing branch using p o s t1 created response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t1 created response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this create workspace with existing branch using p o s t1 created response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t1 created response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create workspace with existing branch using p o s t1 created response
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) Code() int {
	return 201
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1Created", 201)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1Created", 201)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkspaceWithExistingBranchUsingPOST1Unauthorized creates a CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOST1Unauthorized() *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized {
	return &CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized struct {
}

// IsSuccess returns true when this create workspace with existing branch using p o s t1 unauthorized response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create workspace with existing branch using p o s t1 unauthorized response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t1 unauthorized response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create workspace with existing branch using p o s t1 unauthorized response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t1 unauthorized response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create workspace with existing branch using p o s t1 unauthorized response
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) Code() int {
	return 401
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1Unauthorized", 401)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1Unauthorized", 401)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkspaceWithExistingBranchUsingPOST1Forbidden creates a CreateWorkspaceWithExistingBranchUsingPOST1Forbidden with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOST1Forbidden() *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden {
	return &CreateWorkspaceWithExistingBranchUsingPOST1Forbidden{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOST1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateWorkspaceWithExistingBranchUsingPOST1Forbidden struct {
}

// IsSuccess returns true when this create workspace with existing branch using p o s t1 forbidden response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create workspace with existing branch using p o s t1 forbidden response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t1 forbidden response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create workspace with existing branch using p o s t1 forbidden response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t1 forbidden response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create workspace with existing branch using p o s t1 forbidden response
func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) Code() int {
	return 403
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1Forbidden", 403)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1Forbidden", 403)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateWorkspaceWithExistingBranchUsingPOST1NotFound creates a CreateWorkspaceWithExistingBranchUsingPOST1NotFound with default headers values
func NewCreateWorkspaceWithExistingBranchUsingPOST1NotFound() *CreateWorkspaceWithExistingBranchUsingPOST1NotFound {
	return &CreateWorkspaceWithExistingBranchUsingPOST1NotFound{}
}

/*
CreateWorkspaceWithExistingBranchUsingPOST1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateWorkspaceWithExistingBranchUsingPOST1NotFound struct {
}

// IsSuccess returns true when this create workspace with existing branch using p o s t1 not found response has a 2xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create workspace with existing branch using p o s t1 not found response has a 3xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create workspace with existing branch using p o s t1 not found response has a 4xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create workspace with existing branch using p o s t1 not found response has a 5xx status code
func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create workspace with existing branch using p o s t1 not found response a status code equal to that given
func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create workspace with existing branch using p o s t1 not found response
func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) Code() int {
	return 404
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1NotFound", 404)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/coder/stack/{stackName}/resourceType/{resourceType}/resourceName/{resourceName}/existing-branch][%d] createWorkspaceWithExistingBranchUsingPOST1NotFound", 404)
}

func (o *CreateWorkspaceWithExistingBranchUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
