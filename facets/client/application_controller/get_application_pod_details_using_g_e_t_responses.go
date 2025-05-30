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

// GetApplicationPodDetailsUsingGETReader is a Reader for the GetApplicationPodDetailsUsingGET structure.
type GetApplicationPodDetailsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationPodDetailsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationPodDetailsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationPodDetailsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationPodDetailsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationPodDetailsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails] getApplicationPodDetailsUsingGET", response, response.Code())
	}
}

// NewGetApplicationPodDetailsUsingGETOK creates a GetApplicationPodDetailsUsingGETOK with default headers values
func NewGetApplicationPodDetailsUsingGETOK() *GetApplicationPodDetailsUsingGETOK {
	return &GetApplicationPodDetailsUsingGETOK{}
}

/*
GetApplicationPodDetailsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetApplicationPodDetailsUsingGETOK struct {
	Payload []*models.ApplicationPodDetails
}

// IsSuccess returns true when this get application pod details using g e t o k response has a 2xx status code
func (o *GetApplicationPodDetailsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application pod details using g e t o k response has a 3xx status code
func (o *GetApplicationPodDetailsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application pod details using g e t o k response has a 4xx status code
func (o *GetApplicationPodDetailsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application pod details using g e t o k response has a 5xx status code
func (o *GetApplicationPodDetailsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application pod details using g e t o k response a status code equal to that given
func (o *GetApplicationPodDetailsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application pod details using g e t o k response
func (o *GetApplicationPodDetailsUsingGETOK) Code() int {
	return 200
}

func (o *GetApplicationPodDetailsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails][%d] getApplicationPodDetailsUsingGETOK %s", 200, payload)
}

func (o *GetApplicationPodDetailsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails][%d] getApplicationPodDetailsUsingGETOK %s", 200, payload)
}

func (o *GetApplicationPodDetailsUsingGETOK) GetPayload() []*models.ApplicationPodDetails {
	return o.Payload
}

func (o *GetApplicationPodDetailsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationPodDetailsUsingGETUnauthorized creates a GetApplicationPodDetailsUsingGETUnauthorized with default headers values
func NewGetApplicationPodDetailsUsingGETUnauthorized() *GetApplicationPodDetailsUsingGETUnauthorized {
	return &GetApplicationPodDetailsUsingGETUnauthorized{}
}

/*
GetApplicationPodDetailsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetApplicationPodDetailsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get application pod details using g e t unauthorized response has a 2xx status code
func (o *GetApplicationPodDetailsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application pod details using g e t unauthorized response has a 3xx status code
func (o *GetApplicationPodDetailsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application pod details using g e t unauthorized response has a 4xx status code
func (o *GetApplicationPodDetailsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application pod details using g e t unauthorized response has a 5xx status code
func (o *GetApplicationPodDetailsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get application pod details using g e t unauthorized response a status code equal to that given
func (o *GetApplicationPodDetailsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get application pod details using g e t unauthorized response
func (o *GetApplicationPodDetailsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetApplicationPodDetailsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails][%d] getApplicationPodDetailsUsingGETUnauthorized", 401)
}

func (o *GetApplicationPodDetailsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails][%d] getApplicationPodDetailsUsingGETUnauthorized", 401)
}

func (o *GetApplicationPodDetailsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationPodDetailsUsingGETForbidden creates a GetApplicationPodDetailsUsingGETForbidden with default headers values
func NewGetApplicationPodDetailsUsingGETForbidden() *GetApplicationPodDetailsUsingGETForbidden {
	return &GetApplicationPodDetailsUsingGETForbidden{}
}

/*
GetApplicationPodDetailsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetApplicationPodDetailsUsingGETForbidden struct {
}

// IsSuccess returns true when this get application pod details using g e t forbidden response has a 2xx status code
func (o *GetApplicationPodDetailsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application pod details using g e t forbidden response has a 3xx status code
func (o *GetApplicationPodDetailsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application pod details using g e t forbidden response has a 4xx status code
func (o *GetApplicationPodDetailsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application pod details using g e t forbidden response has a 5xx status code
func (o *GetApplicationPodDetailsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get application pod details using g e t forbidden response a status code equal to that given
func (o *GetApplicationPodDetailsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get application pod details using g e t forbidden response
func (o *GetApplicationPodDetailsUsingGETForbidden) Code() int {
	return 403
}

func (o *GetApplicationPodDetailsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails][%d] getApplicationPodDetailsUsingGETForbidden", 403)
}

func (o *GetApplicationPodDetailsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails][%d] getApplicationPodDetailsUsingGETForbidden", 403)
}

func (o *GetApplicationPodDetailsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationPodDetailsUsingGETNotFound creates a GetApplicationPodDetailsUsingGETNotFound with default headers values
func NewGetApplicationPodDetailsUsingGETNotFound() *GetApplicationPodDetailsUsingGETNotFound {
	return &GetApplicationPodDetailsUsingGETNotFound{}
}

/*
GetApplicationPodDetailsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetApplicationPodDetailsUsingGETNotFound struct {
}

// IsSuccess returns true when this get application pod details using g e t not found response has a 2xx status code
func (o *GetApplicationPodDetailsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application pod details using g e t not found response has a 3xx status code
func (o *GetApplicationPodDetailsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application pod details using g e t not found response has a 4xx status code
func (o *GetApplicationPodDetailsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application pod details using g e t not found response has a 5xx status code
func (o *GetApplicationPodDetailsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application pod details using g e t not found response a status code equal to that given
func (o *GetApplicationPodDetailsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application pod details using g e t not found response
func (o *GetApplicationPodDetailsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetApplicationPodDetailsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails][%d] getApplicationPodDetailsUsingGETNotFound", 404)
}

func (o *GetApplicationPodDetailsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/podDetails][%d] getApplicationPodDetailsUsingGETNotFound", 404)
}

func (o *GetApplicationPodDetailsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
