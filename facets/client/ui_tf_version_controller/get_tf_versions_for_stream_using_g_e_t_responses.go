// Code generated by go-swagger; DO NOT EDIT.

package ui_tf_version_controller

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

// GetTfVersionsForStreamUsingGETReader is a Reader for the GetTfVersionsForStreamUsingGET structure.
type GetTfVersionsForStreamUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTfVersionsForStreamUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTfVersionsForStreamUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTfVersionsForStreamUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTfVersionsForStreamUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTfVersionsForStreamUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions] getTfVersionsForStreamUsingGET", response, response.Code())
	}
}

// NewGetTfVersionsForStreamUsingGETOK creates a GetTfVersionsForStreamUsingGETOK with default headers values
func NewGetTfVersionsForStreamUsingGETOK() *GetTfVersionsForStreamUsingGETOK {
	return &GetTfVersionsForStreamUsingGETOK{}
}

/*
GetTfVersionsForStreamUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetTfVersionsForStreamUsingGETOK struct {
	Payload []*models.TfVersion
}

// IsSuccess returns true when this get tf versions for stream using g e t o k response has a 2xx status code
func (o *GetTfVersionsForStreamUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tf versions for stream using g e t o k response has a 3xx status code
func (o *GetTfVersionsForStreamUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tf versions for stream using g e t o k response has a 4xx status code
func (o *GetTfVersionsForStreamUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tf versions for stream using g e t o k response has a 5xx status code
func (o *GetTfVersionsForStreamUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tf versions for stream using g e t o k response a status code equal to that given
func (o *GetTfVersionsForStreamUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tf versions for stream using g e t o k response
func (o *GetTfVersionsForStreamUsingGETOK) Code() int {
	return 200
}

func (o *GetTfVersionsForStreamUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions][%d] getTfVersionsForStreamUsingGETOK %s", 200, payload)
}

func (o *GetTfVersionsForStreamUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions][%d] getTfVersionsForStreamUsingGETOK %s", 200, payload)
}

func (o *GetTfVersionsForStreamUsingGETOK) GetPayload() []*models.TfVersion {
	return o.Payload
}

func (o *GetTfVersionsForStreamUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTfVersionsForStreamUsingGETUnauthorized creates a GetTfVersionsForStreamUsingGETUnauthorized with default headers values
func NewGetTfVersionsForStreamUsingGETUnauthorized() *GetTfVersionsForStreamUsingGETUnauthorized {
	return &GetTfVersionsForStreamUsingGETUnauthorized{}
}

/*
GetTfVersionsForStreamUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTfVersionsForStreamUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get tf versions for stream using g e t unauthorized response has a 2xx status code
func (o *GetTfVersionsForStreamUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tf versions for stream using g e t unauthorized response has a 3xx status code
func (o *GetTfVersionsForStreamUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tf versions for stream using g e t unauthorized response has a 4xx status code
func (o *GetTfVersionsForStreamUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tf versions for stream using g e t unauthorized response has a 5xx status code
func (o *GetTfVersionsForStreamUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get tf versions for stream using g e t unauthorized response a status code equal to that given
func (o *GetTfVersionsForStreamUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get tf versions for stream using g e t unauthorized response
func (o *GetTfVersionsForStreamUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetTfVersionsForStreamUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions][%d] getTfVersionsForStreamUsingGETUnauthorized", 401)
}

func (o *GetTfVersionsForStreamUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions][%d] getTfVersionsForStreamUsingGETUnauthorized", 401)
}

func (o *GetTfVersionsForStreamUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTfVersionsForStreamUsingGETForbidden creates a GetTfVersionsForStreamUsingGETForbidden with default headers values
func NewGetTfVersionsForStreamUsingGETForbidden() *GetTfVersionsForStreamUsingGETForbidden {
	return &GetTfVersionsForStreamUsingGETForbidden{}
}

/*
GetTfVersionsForStreamUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTfVersionsForStreamUsingGETForbidden struct {
}

// IsSuccess returns true when this get tf versions for stream using g e t forbidden response has a 2xx status code
func (o *GetTfVersionsForStreamUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tf versions for stream using g e t forbidden response has a 3xx status code
func (o *GetTfVersionsForStreamUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tf versions for stream using g e t forbidden response has a 4xx status code
func (o *GetTfVersionsForStreamUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tf versions for stream using g e t forbidden response has a 5xx status code
func (o *GetTfVersionsForStreamUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get tf versions for stream using g e t forbidden response a status code equal to that given
func (o *GetTfVersionsForStreamUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get tf versions for stream using g e t forbidden response
func (o *GetTfVersionsForStreamUsingGETForbidden) Code() int {
	return 403
}

func (o *GetTfVersionsForStreamUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions][%d] getTfVersionsForStreamUsingGETForbidden", 403)
}

func (o *GetTfVersionsForStreamUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions][%d] getTfVersionsForStreamUsingGETForbidden", 403)
}

func (o *GetTfVersionsForStreamUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTfVersionsForStreamUsingGETNotFound creates a GetTfVersionsForStreamUsingGETNotFound with default headers values
func NewGetTfVersionsForStreamUsingGETNotFound() *GetTfVersionsForStreamUsingGETNotFound {
	return &GetTfVersionsForStreamUsingGETNotFound{}
}

/*
GetTfVersionsForStreamUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetTfVersionsForStreamUsingGETNotFound struct {
}

// IsSuccess returns true when this get tf versions for stream using g e t not found response has a 2xx status code
func (o *GetTfVersionsForStreamUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tf versions for stream using g e t not found response has a 3xx status code
func (o *GetTfVersionsForStreamUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tf versions for stream using g e t not found response has a 4xx status code
func (o *GetTfVersionsForStreamUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tf versions for stream using g e t not found response has a 5xx status code
func (o *GetTfVersionsForStreamUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tf versions for stream using g e t not found response a status code equal to that given
func (o *GetTfVersionsForStreamUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get tf versions for stream using g e t not found response
func (o *GetTfVersionsForStreamUsingGETNotFound) Code() int {
	return 404
}

func (o *GetTfVersionsForStreamUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions][%d] getTfVersionsForStreamUsingGETNotFound", 404)
}

func (o *GetTfVersionsForStreamUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/terraform/stream/{tfStream}/versions][%d] getTfVersionsForStreamUsingGETNotFound", 404)
}

func (o *GetTfVersionsForStreamUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
