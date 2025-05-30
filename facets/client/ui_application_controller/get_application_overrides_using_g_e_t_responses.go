// Code generated by go-swagger; DO NOT EDIT.

package ui_application_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetApplicationOverridesUsingGETReader is a Reader for the GetApplicationOverridesUsingGET structure.
type GetApplicationOverridesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationOverridesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationOverridesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationOverridesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationOverridesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationOverridesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides] getApplicationOverridesUsingGET", response, response.Code())
	}
}

// NewGetApplicationOverridesUsingGETOK creates a GetApplicationOverridesUsingGETOK with default headers values
func NewGetApplicationOverridesUsingGETOK() *GetApplicationOverridesUsingGETOK {
	return &GetApplicationOverridesUsingGETOK{}
}

/*
GetApplicationOverridesUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetApplicationOverridesUsingGETOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get application overrides using g e t o k response has a 2xx status code
func (o *GetApplicationOverridesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application overrides using g e t o k response has a 3xx status code
func (o *GetApplicationOverridesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application overrides using g e t o k response has a 4xx status code
func (o *GetApplicationOverridesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application overrides using g e t o k response has a 5xx status code
func (o *GetApplicationOverridesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application overrides using g e t o k response a status code equal to that given
func (o *GetApplicationOverridesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application overrides using g e t o k response
func (o *GetApplicationOverridesUsingGETOK) Code() int {
	return 200
}

func (o *GetApplicationOverridesUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides][%d] getApplicationOverridesUsingGETOK %s", 200, payload)
}

func (o *GetApplicationOverridesUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides][%d] getApplicationOverridesUsingGETOK %s", 200, payload)
}

func (o *GetApplicationOverridesUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetApplicationOverridesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationOverridesUsingGETUnauthorized creates a GetApplicationOverridesUsingGETUnauthorized with default headers values
func NewGetApplicationOverridesUsingGETUnauthorized() *GetApplicationOverridesUsingGETUnauthorized {
	return &GetApplicationOverridesUsingGETUnauthorized{}
}

/*
GetApplicationOverridesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetApplicationOverridesUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get application overrides using g e t unauthorized response has a 2xx status code
func (o *GetApplicationOverridesUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application overrides using g e t unauthorized response has a 3xx status code
func (o *GetApplicationOverridesUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application overrides using g e t unauthorized response has a 4xx status code
func (o *GetApplicationOverridesUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application overrides using g e t unauthorized response has a 5xx status code
func (o *GetApplicationOverridesUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get application overrides using g e t unauthorized response a status code equal to that given
func (o *GetApplicationOverridesUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get application overrides using g e t unauthorized response
func (o *GetApplicationOverridesUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetApplicationOverridesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides][%d] getApplicationOverridesUsingGETUnauthorized", 401)
}

func (o *GetApplicationOverridesUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides][%d] getApplicationOverridesUsingGETUnauthorized", 401)
}

func (o *GetApplicationOverridesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationOverridesUsingGETForbidden creates a GetApplicationOverridesUsingGETForbidden with default headers values
func NewGetApplicationOverridesUsingGETForbidden() *GetApplicationOverridesUsingGETForbidden {
	return &GetApplicationOverridesUsingGETForbidden{}
}

/*
GetApplicationOverridesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetApplicationOverridesUsingGETForbidden struct {
}

// IsSuccess returns true when this get application overrides using g e t forbidden response has a 2xx status code
func (o *GetApplicationOverridesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application overrides using g e t forbidden response has a 3xx status code
func (o *GetApplicationOverridesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application overrides using g e t forbidden response has a 4xx status code
func (o *GetApplicationOverridesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application overrides using g e t forbidden response has a 5xx status code
func (o *GetApplicationOverridesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get application overrides using g e t forbidden response a status code equal to that given
func (o *GetApplicationOverridesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get application overrides using g e t forbidden response
func (o *GetApplicationOverridesUsingGETForbidden) Code() int {
	return 403
}

func (o *GetApplicationOverridesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides][%d] getApplicationOverridesUsingGETForbidden", 403)
}

func (o *GetApplicationOverridesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides][%d] getApplicationOverridesUsingGETForbidden", 403)
}

func (o *GetApplicationOverridesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationOverridesUsingGETNotFound creates a GetApplicationOverridesUsingGETNotFound with default headers values
func NewGetApplicationOverridesUsingGETNotFound() *GetApplicationOverridesUsingGETNotFound {
	return &GetApplicationOverridesUsingGETNotFound{}
}

/*
GetApplicationOverridesUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetApplicationOverridesUsingGETNotFound struct {
}

// IsSuccess returns true when this get application overrides using g e t not found response has a 2xx status code
func (o *GetApplicationOverridesUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application overrides using g e t not found response has a 3xx status code
func (o *GetApplicationOverridesUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application overrides using g e t not found response has a 4xx status code
func (o *GetApplicationOverridesUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application overrides using g e t not found response has a 5xx status code
func (o *GetApplicationOverridesUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application overrides using g e t not found response a status code equal to that given
func (o *GetApplicationOverridesUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application overrides using g e t not found response
func (o *GetApplicationOverridesUsingGETNotFound) Code() int {
	return 404
}

func (o *GetApplicationOverridesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides][%d] getApplicationOverridesUsingGETNotFound", 404)
}

func (o *GetApplicationOverridesUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/{resourceType}/{appName}/overrides][%d] getApplicationOverridesUsingGETNotFound", 404)
}

func (o *GetApplicationOverridesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
