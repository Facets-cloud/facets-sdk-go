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

// GetApplicationSecretsUsingGETReader is a Reader for the GetApplicationSecretsUsingGET structure.
type GetApplicationSecretsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationSecretsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationSecretsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationSecretsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationSecretsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationSecretsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests] getApplicationSecretsUsingGET", response, response.Code())
	}
}

// NewGetApplicationSecretsUsingGETOK creates a GetApplicationSecretsUsingGETOK with default headers values
func NewGetApplicationSecretsUsingGETOK() *GetApplicationSecretsUsingGETOK {
	return &GetApplicationSecretsUsingGETOK{}
}

/*
GetApplicationSecretsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetApplicationSecretsUsingGETOK struct {
	Payload []*models.ApplicationSecret
}

// IsSuccess returns true when this get application secrets using g e t o k response has a 2xx status code
func (o *GetApplicationSecretsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application secrets using g e t o k response has a 3xx status code
func (o *GetApplicationSecretsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application secrets using g e t o k response has a 4xx status code
func (o *GetApplicationSecretsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application secrets using g e t o k response has a 5xx status code
func (o *GetApplicationSecretsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application secrets using g e t o k response a status code equal to that given
func (o *GetApplicationSecretsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application secrets using g e t o k response
func (o *GetApplicationSecretsUsingGETOK) Code() int {
	return 200
}

func (o *GetApplicationSecretsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests][%d] getApplicationSecretsUsingGETOK %s", 200, payload)
}

func (o *GetApplicationSecretsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests][%d] getApplicationSecretsUsingGETOK %s", 200, payload)
}

func (o *GetApplicationSecretsUsingGETOK) GetPayload() []*models.ApplicationSecret {
	return o.Payload
}

func (o *GetApplicationSecretsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationSecretsUsingGETUnauthorized creates a GetApplicationSecretsUsingGETUnauthorized with default headers values
func NewGetApplicationSecretsUsingGETUnauthorized() *GetApplicationSecretsUsingGETUnauthorized {
	return &GetApplicationSecretsUsingGETUnauthorized{}
}

/*
GetApplicationSecretsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetApplicationSecretsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get application secrets using g e t unauthorized response has a 2xx status code
func (o *GetApplicationSecretsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application secrets using g e t unauthorized response has a 3xx status code
func (o *GetApplicationSecretsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application secrets using g e t unauthorized response has a 4xx status code
func (o *GetApplicationSecretsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application secrets using g e t unauthorized response has a 5xx status code
func (o *GetApplicationSecretsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get application secrets using g e t unauthorized response a status code equal to that given
func (o *GetApplicationSecretsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get application secrets using g e t unauthorized response
func (o *GetApplicationSecretsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetApplicationSecretsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests][%d] getApplicationSecretsUsingGETUnauthorized", 401)
}

func (o *GetApplicationSecretsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests][%d] getApplicationSecretsUsingGETUnauthorized", 401)
}

func (o *GetApplicationSecretsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationSecretsUsingGETForbidden creates a GetApplicationSecretsUsingGETForbidden with default headers values
func NewGetApplicationSecretsUsingGETForbidden() *GetApplicationSecretsUsingGETForbidden {
	return &GetApplicationSecretsUsingGETForbidden{}
}

/*
GetApplicationSecretsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetApplicationSecretsUsingGETForbidden struct {
}

// IsSuccess returns true when this get application secrets using g e t forbidden response has a 2xx status code
func (o *GetApplicationSecretsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application secrets using g e t forbidden response has a 3xx status code
func (o *GetApplicationSecretsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application secrets using g e t forbidden response has a 4xx status code
func (o *GetApplicationSecretsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application secrets using g e t forbidden response has a 5xx status code
func (o *GetApplicationSecretsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get application secrets using g e t forbidden response a status code equal to that given
func (o *GetApplicationSecretsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get application secrets using g e t forbidden response
func (o *GetApplicationSecretsUsingGETForbidden) Code() int {
	return 403
}

func (o *GetApplicationSecretsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests][%d] getApplicationSecretsUsingGETForbidden", 403)
}

func (o *GetApplicationSecretsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests][%d] getApplicationSecretsUsingGETForbidden", 403)
}

func (o *GetApplicationSecretsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationSecretsUsingGETNotFound creates a GetApplicationSecretsUsingGETNotFound with default headers values
func NewGetApplicationSecretsUsingGETNotFound() *GetApplicationSecretsUsingGETNotFound {
	return &GetApplicationSecretsUsingGETNotFound{}
}

/*
GetApplicationSecretsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetApplicationSecretsUsingGETNotFound struct {
}

// IsSuccess returns true when this get application secrets using g e t not found response has a 2xx status code
func (o *GetApplicationSecretsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application secrets using g e t not found response has a 3xx status code
func (o *GetApplicationSecretsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application secrets using g e t not found response has a 4xx status code
func (o *GetApplicationSecretsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application secrets using g e t not found response has a 5xx status code
func (o *GetApplicationSecretsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application secrets using g e t not found response a status code equal to that given
func (o *GetApplicationSecretsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application secrets using g e t not found response
func (o *GetApplicationSecretsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetApplicationSecretsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests][%d] getApplicationSecretsUsingGETNotFound", 404)
}

func (o *GetApplicationSecretsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/secretRequests][%d] getApplicationSecretsUsingGETNotFound", 404)
}

func (o *GetApplicationSecretsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
