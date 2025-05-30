// Code generated by go-swagger; DO NOT EDIT.

package application_controller

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

// CreateApplicationUsingPOSTReader is a Reader for the CreateApplicationUsingPOST structure.
type CreateApplicationUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateApplicationUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateApplicationUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateApplicationUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateApplicationUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateApplicationUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateApplicationUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /api/{applicationFamily}/applications] createApplicationUsingPOST", response, response.Code())
	}
}

// NewCreateApplicationUsingPOSTOK creates a CreateApplicationUsingPOSTOK with default headers values
func NewCreateApplicationUsingPOSTOK() *CreateApplicationUsingPOSTOK {
	return &CreateApplicationUsingPOSTOK{}
}

/*
CreateApplicationUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type CreateApplicationUsingPOSTOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this create application using p o s t o k response has a 2xx status code
func (o *CreateApplicationUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create application using p o s t o k response has a 3xx status code
func (o *CreateApplicationUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create application using p o s t o k response has a 4xx status code
func (o *CreateApplicationUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create application using p o s t o k response has a 5xx status code
func (o *CreateApplicationUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create application using p o s t o k response a status code equal to that given
func (o *CreateApplicationUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create application using p o s t o k response
func (o *CreateApplicationUsingPOSTOK) Code() int {
	return 200
}

func (o *CreateApplicationUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTOK %s", 200, payload)
}

func (o *CreateApplicationUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTOK %s", 200, payload)
}

func (o *CreateApplicationUsingPOSTOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *CreateApplicationUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateApplicationUsingPOSTCreated creates a CreateApplicationUsingPOSTCreated with default headers values
func NewCreateApplicationUsingPOSTCreated() *CreateApplicationUsingPOSTCreated {
	return &CreateApplicationUsingPOSTCreated{}
}

/*
CreateApplicationUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type CreateApplicationUsingPOSTCreated struct {
}

// IsSuccess returns true when this create application using p o s t created response has a 2xx status code
func (o *CreateApplicationUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create application using p o s t created response has a 3xx status code
func (o *CreateApplicationUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create application using p o s t created response has a 4xx status code
func (o *CreateApplicationUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create application using p o s t created response has a 5xx status code
func (o *CreateApplicationUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create application using p o s t created response a status code equal to that given
func (o *CreateApplicationUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create application using p o s t created response
func (o *CreateApplicationUsingPOSTCreated) Code() int {
	return 201
}

func (o *CreateApplicationUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTCreated", 201)
}

func (o *CreateApplicationUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTCreated", 201)
}

func (o *CreateApplicationUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateApplicationUsingPOSTUnauthorized creates a CreateApplicationUsingPOSTUnauthorized with default headers values
func NewCreateApplicationUsingPOSTUnauthorized() *CreateApplicationUsingPOSTUnauthorized {
	return &CreateApplicationUsingPOSTUnauthorized{}
}

/*
CreateApplicationUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateApplicationUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this create application using p o s t unauthorized response has a 2xx status code
func (o *CreateApplicationUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create application using p o s t unauthorized response has a 3xx status code
func (o *CreateApplicationUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create application using p o s t unauthorized response has a 4xx status code
func (o *CreateApplicationUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create application using p o s t unauthorized response has a 5xx status code
func (o *CreateApplicationUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create application using p o s t unauthorized response a status code equal to that given
func (o *CreateApplicationUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create application using p o s t unauthorized response
func (o *CreateApplicationUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *CreateApplicationUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTUnauthorized", 401)
}

func (o *CreateApplicationUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTUnauthorized", 401)
}

func (o *CreateApplicationUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateApplicationUsingPOSTForbidden creates a CreateApplicationUsingPOSTForbidden with default headers values
func NewCreateApplicationUsingPOSTForbidden() *CreateApplicationUsingPOSTForbidden {
	return &CreateApplicationUsingPOSTForbidden{}
}

/*
CreateApplicationUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateApplicationUsingPOSTForbidden struct {
}

// IsSuccess returns true when this create application using p o s t forbidden response has a 2xx status code
func (o *CreateApplicationUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create application using p o s t forbidden response has a 3xx status code
func (o *CreateApplicationUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create application using p o s t forbidden response has a 4xx status code
func (o *CreateApplicationUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create application using p o s t forbidden response has a 5xx status code
func (o *CreateApplicationUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create application using p o s t forbidden response a status code equal to that given
func (o *CreateApplicationUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create application using p o s t forbidden response
func (o *CreateApplicationUsingPOSTForbidden) Code() int {
	return 403
}

func (o *CreateApplicationUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTForbidden", 403)
}

func (o *CreateApplicationUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTForbidden", 403)
}

func (o *CreateApplicationUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateApplicationUsingPOSTNotFound creates a CreateApplicationUsingPOSTNotFound with default headers values
func NewCreateApplicationUsingPOSTNotFound() *CreateApplicationUsingPOSTNotFound {
	return &CreateApplicationUsingPOSTNotFound{}
}

/*
CreateApplicationUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateApplicationUsingPOSTNotFound struct {
}

// IsSuccess returns true when this create application using p o s t not found response has a 2xx status code
func (o *CreateApplicationUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create application using p o s t not found response has a 3xx status code
func (o *CreateApplicationUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create application using p o s t not found response has a 4xx status code
func (o *CreateApplicationUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create application using p o s t not found response has a 5xx status code
func (o *CreateApplicationUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create application using p o s t not found response a status code equal to that given
func (o *CreateApplicationUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create application using p o s t not found response
func (o *CreateApplicationUsingPOSTNotFound) Code() int {
	return 404
}

func (o *CreateApplicationUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTNotFound", 404)
}

func (o *CreateApplicationUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /api/{applicationFamily}/applications][%d] createApplicationUsingPOSTNotFound", 404)
}

func (o *CreateApplicationUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
