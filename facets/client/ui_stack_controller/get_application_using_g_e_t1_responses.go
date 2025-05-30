// Code generated by go-swagger; DO NOT EDIT.

package ui_stack_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetApplicationUsingGET1Reader is a Reader for the GetApplicationUsingGET1 structure.
type GetApplicationUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}] getApplicationUsingGET_1", response, response.Code())
	}
}

// NewGetApplicationUsingGET1OK creates a GetApplicationUsingGET1OK with default headers values
func NewGetApplicationUsingGET1OK() *GetApplicationUsingGET1OK {
	return &GetApplicationUsingGET1OK{}
}

/*
GetApplicationUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetApplicationUsingGET1OK struct {
	Payload interface{}
}

// IsSuccess returns true when this get application using g e t1 o k response has a 2xx status code
func (o *GetApplicationUsingGET1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application using g e t1 o k response has a 3xx status code
func (o *GetApplicationUsingGET1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application using g e t1 o k response has a 4xx status code
func (o *GetApplicationUsingGET1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application using g e t1 o k response has a 5xx status code
func (o *GetApplicationUsingGET1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application using g e t1 o k response a status code equal to that given
func (o *GetApplicationUsingGET1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application using g e t1 o k response
func (o *GetApplicationUsingGET1OK) Code() int {
	return 200
}

func (o *GetApplicationUsingGET1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}][%d] getApplicationUsingGET1OK %s", 200, payload)
}

func (o *GetApplicationUsingGET1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}][%d] getApplicationUsingGET1OK %s", 200, payload)
}

func (o *GetApplicationUsingGET1OK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetApplicationUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationUsingGET1Unauthorized creates a GetApplicationUsingGET1Unauthorized with default headers values
func NewGetApplicationUsingGET1Unauthorized() *GetApplicationUsingGET1Unauthorized {
	return &GetApplicationUsingGET1Unauthorized{}
}

/*
GetApplicationUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetApplicationUsingGET1Unauthorized struct {
}

// IsSuccess returns true when this get application using g e t1 unauthorized response has a 2xx status code
func (o *GetApplicationUsingGET1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application using g e t1 unauthorized response has a 3xx status code
func (o *GetApplicationUsingGET1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application using g e t1 unauthorized response has a 4xx status code
func (o *GetApplicationUsingGET1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application using g e t1 unauthorized response has a 5xx status code
func (o *GetApplicationUsingGET1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get application using g e t1 unauthorized response a status code equal to that given
func (o *GetApplicationUsingGET1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get application using g e t1 unauthorized response
func (o *GetApplicationUsingGET1Unauthorized) Code() int {
	return 401
}

func (o *GetApplicationUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}][%d] getApplicationUsingGET1Unauthorized", 401)
}

func (o *GetApplicationUsingGET1Unauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}][%d] getApplicationUsingGET1Unauthorized", 401)
}

func (o *GetApplicationUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationUsingGET1Forbidden creates a GetApplicationUsingGET1Forbidden with default headers values
func NewGetApplicationUsingGET1Forbidden() *GetApplicationUsingGET1Forbidden {
	return &GetApplicationUsingGET1Forbidden{}
}

/*
GetApplicationUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetApplicationUsingGET1Forbidden struct {
}

// IsSuccess returns true when this get application using g e t1 forbidden response has a 2xx status code
func (o *GetApplicationUsingGET1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application using g e t1 forbidden response has a 3xx status code
func (o *GetApplicationUsingGET1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application using g e t1 forbidden response has a 4xx status code
func (o *GetApplicationUsingGET1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application using g e t1 forbidden response has a 5xx status code
func (o *GetApplicationUsingGET1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get application using g e t1 forbidden response a status code equal to that given
func (o *GetApplicationUsingGET1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get application using g e t1 forbidden response
func (o *GetApplicationUsingGET1Forbidden) Code() int {
	return 403
}

func (o *GetApplicationUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}][%d] getApplicationUsingGET1Forbidden", 403)
}

func (o *GetApplicationUsingGET1Forbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}][%d] getApplicationUsingGET1Forbidden", 403)
}

func (o *GetApplicationUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationUsingGET1NotFound creates a GetApplicationUsingGET1NotFound with default headers values
func NewGetApplicationUsingGET1NotFound() *GetApplicationUsingGET1NotFound {
	return &GetApplicationUsingGET1NotFound{}
}

/*
GetApplicationUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetApplicationUsingGET1NotFound struct {
}

// IsSuccess returns true when this get application using g e t1 not found response has a 2xx status code
func (o *GetApplicationUsingGET1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application using g e t1 not found response has a 3xx status code
func (o *GetApplicationUsingGET1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application using g e t1 not found response has a 4xx status code
func (o *GetApplicationUsingGET1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application using g e t1 not found response has a 5xx status code
func (o *GetApplicationUsingGET1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application using g e t1 not found response a status code equal to that given
func (o *GetApplicationUsingGET1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application using g e t1 not found response
func (o *GetApplicationUsingGET1NotFound) Code() int {
	return 404
}

func (o *GetApplicationUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}][%d] getApplicationUsingGET1NotFound", 404)
}

func (o *GetApplicationUsingGET1NotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/{stackName}/{resourceType}/{appName}][%d] getApplicationUsingGET1NotFound", 404)
}

func (o *GetApplicationUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
