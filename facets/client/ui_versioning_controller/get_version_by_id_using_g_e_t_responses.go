// Code generated by go-swagger; DO NOT EDIT.

package ui_versioning_controller

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

// GetVersionByIDUsingGETReader is a Reader for the GetVersionByIDUsingGET structure.
type GetVersionByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVersionByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVersionByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVersionByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/versions/id/{id}] getVersionByIdUsingGET", response, response.Code())
	}
}

// NewGetVersionByIDUsingGETOK creates a GetVersionByIDUsingGETOK with default headers values
func NewGetVersionByIDUsingGETOK() *GetVersionByIDUsingGETOK {
	return &GetVersionByIDUsingGETOK{}
}

/*
GetVersionByIDUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetVersionByIDUsingGETOK struct {
	Payload *models.Version
}

// IsSuccess returns true when this get version by Id using g e t o k response has a 2xx status code
func (o *GetVersionByIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get version by Id using g e t o k response has a 3xx status code
func (o *GetVersionByIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version by Id using g e t o k response has a 4xx status code
func (o *GetVersionByIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get version by Id using g e t o k response has a 5xx status code
func (o *GetVersionByIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get version by Id using g e t o k response a status code equal to that given
func (o *GetVersionByIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get version by Id using g e t o k response
func (o *GetVersionByIDUsingGETOK) Code() int {
	return 200
}

func (o *GetVersionByIDUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/versions/id/{id}][%d] getVersionByIdUsingGETOK %s", 200, payload)
}

func (o *GetVersionByIDUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/versions/id/{id}][%d] getVersionByIdUsingGETOK %s", 200, payload)
}

func (o *GetVersionByIDUsingGETOK) GetPayload() *models.Version {
	return o.Payload
}

func (o *GetVersionByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Version)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionByIDUsingGETUnauthorized creates a GetVersionByIDUsingGETUnauthorized with default headers values
func NewGetVersionByIDUsingGETUnauthorized() *GetVersionByIDUsingGETUnauthorized {
	return &GetVersionByIDUsingGETUnauthorized{}
}

/*
GetVersionByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVersionByIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get version by Id using g e t unauthorized response has a 2xx status code
func (o *GetVersionByIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get version by Id using g e t unauthorized response has a 3xx status code
func (o *GetVersionByIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version by Id using g e t unauthorized response has a 4xx status code
func (o *GetVersionByIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get version by Id using g e t unauthorized response has a 5xx status code
func (o *GetVersionByIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get version by Id using g e t unauthorized response a status code equal to that given
func (o *GetVersionByIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get version by Id using g e t unauthorized response
func (o *GetVersionByIDUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetVersionByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/versions/id/{id}][%d] getVersionByIdUsingGETUnauthorized", 401)
}

func (o *GetVersionByIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/versions/id/{id}][%d] getVersionByIdUsingGETUnauthorized", 401)
}

func (o *GetVersionByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionByIDUsingGETForbidden creates a GetVersionByIDUsingGETForbidden with default headers values
func NewGetVersionByIDUsingGETForbidden() *GetVersionByIDUsingGETForbidden {
	return &GetVersionByIDUsingGETForbidden{}
}

/*
GetVersionByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVersionByIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get version by Id using g e t forbidden response has a 2xx status code
func (o *GetVersionByIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get version by Id using g e t forbidden response has a 3xx status code
func (o *GetVersionByIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version by Id using g e t forbidden response has a 4xx status code
func (o *GetVersionByIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get version by Id using g e t forbidden response has a 5xx status code
func (o *GetVersionByIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get version by Id using g e t forbidden response a status code equal to that given
func (o *GetVersionByIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get version by Id using g e t forbidden response
func (o *GetVersionByIDUsingGETForbidden) Code() int {
	return 403
}

func (o *GetVersionByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/versions/id/{id}][%d] getVersionByIdUsingGETForbidden", 403)
}

func (o *GetVersionByIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/versions/id/{id}][%d] getVersionByIdUsingGETForbidden", 403)
}

func (o *GetVersionByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionByIDUsingGETNotFound creates a GetVersionByIDUsingGETNotFound with default headers values
func NewGetVersionByIDUsingGETNotFound() *GetVersionByIDUsingGETNotFound {
	return &GetVersionByIDUsingGETNotFound{}
}

/*
GetVersionByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVersionByIDUsingGETNotFound struct {
}

// IsSuccess returns true when this get version by Id using g e t not found response has a 2xx status code
func (o *GetVersionByIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get version by Id using g e t not found response has a 3xx status code
func (o *GetVersionByIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version by Id using g e t not found response has a 4xx status code
func (o *GetVersionByIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get version by Id using g e t not found response has a 5xx status code
func (o *GetVersionByIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get version by Id using g e t not found response a status code equal to that given
func (o *GetVersionByIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get version by Id using g e t not found response
func (o *GetVersionByIDUsingGETNotFound) Code() int {
	return 404
}

func (o *GetVersionByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/versions/id/{id}][%d] getVersionByIdUsingGETNotFound", 404)
}

func (o *GetVersionByIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/versions/id/{id}][%d] getVersionByIdUsingGETNotFound", 404)
}

func (o *GetVersionByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
