// Code generated by go-swagger; DO NOT EDIT.

package ui_opa_controller

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

// EditPolicyUsingPUTReader is a Reader for the EditPolicyUsingPUT structure.
type EditPolicyUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditPolicyUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditPolicyUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewEditPolicyUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEditPolicyUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEditPolicyUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEditPolicyUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /cc-ui/v1/opa/{policyName}] editPolicyUsingPUT", response, response.Code())
	}
}

// NewEditPolicyUsingPUTOK creates a EditPolicyUsingPUTOK with default headers values
func NewEditPolicyUsingPUTOK() *EditPolicyUsingPUTOK {
	return &EditPolicyUsingPUTOK{}
}

/*
EditPolicyUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type EditPolicyUsingPUTOK struct {
	Payload *models.OpaPolicy
}

// IsSuccess returns true when this edit policy using p u t o k response has a 2xx status code
func (o *EditPolicyUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edit policy using p u t o k response has a 3xx status code
func (o *EditPolicyUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit policy using p u t o k response has a 4xx status code
func (o *EditPolicyUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edit policy using p u t o k response has a 5xx status code
func (o *EditPolicyUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edit policy using p u t o k response a status code equal to that given
func (o *EditPolicyUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edit policy using p u t o k response
func (o *EditPolicyUsingPUTOK) Code() int {
	return 200
}

func (o *EditPolicyUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTOK %s", 200, payload)
}

func (o *EditPolicyUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTOK %s", 200, payload)
}

func (o *EditPolicyUsingPUTOK) GetPayload() *models.OpaPolicy {
	return o.Payload
}

func (o *EditPolicyUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpaPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditPolicyUsingPUTCreated creates a EditPolicyUsingPUTCreated with default headers values
func NewEditPolicyUsingPUTCreated() *EditPolicyUsingPUTCreated {
	return &EditPolicyUsingPUTCreated{}
}

/*
EditPolicyUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type EditPolicyUsingPUTCreated struct {
}

// IsSuccess returns true when this edit policy using p u t created response has a 2xx status code
func (o *EditPolicyUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edit policy using p u t created response has a 3xx status code
func (o *EditPolicyUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit policy using p u t created response has a 4xx status code
func (o *EditPolicyUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this edit policy using p u t created response has a 5xx status code
func (o *EditPolicyUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this edit policy using p u t created response a status code equal to that given
func (o *EditPolicyUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the edit policy using p u t created response
func (o *EditPolicyUsingPUTCreated) Code() int {
	return 201
}

func (o *EditPolicyUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTCreated", 201)
}

func (o *EditPolicyUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTCreated", 201)
}

func (o *EditPolicyUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditPolicyUsingPUTUnauthorized creates a EditPolicyUsingPUTUnauthorized with default headers values
func NewEditPolicyUsingPUTUnauthorized() *EditPolicyUsingPUTUnauthorized {
	return &EditPolicyUsingPUTUnauthorized{}
}

/*
EditPolicyUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type EditPolicyUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this edit policy using p u t unauthorized response has a 2xx status code
func (o *EditPolicyUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edit policy using p u t unauthorized response has a 3xx status code
func (o *EditPolicyUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit policy using p u t unauthorized response has a 4xx status code
func (o *EditPolicyUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edit policy using p u t unauthorized response has a 5xx status code
func (o *EditPolicyUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edit policy using p u t unauthorized response a status code equal to that given
func (o *EditPolicyUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edit policy using p u t unauthorized response
func (o *EditPolicyUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *EditPolicyUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTUnauthorized", 401)
}

func (o *EditPolicyUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTUnauthorized", 401)
}

func (o *EditPolicyUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditPolicyUsingPUTForbidden creates a EditPolicyUsingPUTForbidden with default headers values
func NewEditPolicyUsingPUTForbidden() *EditPolicyUsingPUTForbidden {
	return &EditPolicyUsingPUTForbidden{}
}

/*
EditPolicyUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EditPolicyUsingPUTForbidden struct {
}

// IsSuccess returns true when this edit policy using p u t forbidden response has a 2xx status code
func (o *EditPolicyUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edit policy using p u t forbidden response has a 3xx status code
func (o *EditPolicyUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit policy using p u t forbidden response has a 4xx status code
func (o *EditPolicyUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edit policy using p u t forbidden response has a 5xx status code
func (o *EditPolicyUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edit policy using p u t forbidden response a status code equal to that given
func (o *EditPolicyUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edit policy using p u t forbidden response
func (o *EditPolicyUsingPUTForbidden) Code() int {
	return 403
}

func (o *EditPolicyUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTForbidden", 403)
}

func (o *EditPolicyUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTForbidden", 403)
}

func (o *EditPolicyUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditPolicyUsingPUTNotFound creates a EditPolicyUsingPUTNotFound with default headers values
func NewEditPolicyUsingPUTNotFound() *EditPolicyUsingPUTNotFound {
	return &EditPolicyUsingPUTNotFound{}
}

/*
EditPolicyUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type EditPolicyUsingPUTNotFound struct {
}

// IsSuccess returns true when this edit policy using p u t not found response has a 2xx status code
func (o *EditPolicyUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edit policy using p u t not found response has a 3xx status code
func (o *EditPolicyUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edit policy using p u t not found response has a 4xx status code
func (o *EditPolicyUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edit policy using p u t not found response has a 5xx status code
func (o *EditPolicyUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edit policy using p u t not found response a status code equal to that given
func (o *EditPolicyUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edit policy using p u t not found response
func (o *EditPolicyUsingPUTNotFound) Code() int {
	return 404
}

func (o *EditPolicyUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTNotFound", 404)
}

func (o *EditPolicyUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/opa/{policyName}][%d] editPolicyUsingPUTNotFound", 404)
}

func (o *EditPolicyUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
