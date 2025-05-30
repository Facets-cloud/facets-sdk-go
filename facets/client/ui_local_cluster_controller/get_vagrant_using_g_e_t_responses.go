// Code generated by go-swagger; DO NOT EDIT.

package ui_local_cluster_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetVagrantUsingGETReader is a Reader for the GetVagrantUsingGET structure.
type GetVagrantUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVagrantUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVagrantUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVagrantUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVagrantUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVagrantUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant] getVagrantUsingGET", response, response.Code())
	}
}

// NewGetVagrantUsingGETOK creates a GetVagrantUsingGETOK with default headers values
func NewGetVagrantUsingGETOK() *GetVagrantUsingGETOK {
	return &GetVagrantUsingGETOK{}
}

/*
GetVagrantUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetVagrantUsingGETOK struct {
	Payload string
}

// IsSuccess returns true when this get vagrant using g e t o k response has a 2xx status code
func (o *GetVagrantUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vagrant using g e t o k response has a 3xx status code
func (o *GetVagrantUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vagrant using g e t o k response has a 4xx status code
func (o *GetVagrantUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vagrant using g e t o k response has a 5xx status code
func (o *GetVagrantUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vagrant using g e t o k response a status code equal to that given
func (o *GetVagrantUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get vagrant using g e t o k response
func (o *GetVagrantUsingGETOK) Code() int {
	return 200
}

func (o *GetVagrantUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant][%d] getVagrantUsingGETOK %s", 200, payload)
}

func (o *GetVagrantUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant][%d] getVagrantUsingGETOK %s", 200, payload)
}

func (o *GetVagrantUsingGETOK) GetPayload() string {
	return o.Payload
}

func (o *GetVagrantUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVagrantUsingGETUnauthorized creates a GetVagrantUsingGETUnauthorized with default headers values
func NewGetVagrantUsingGETUnauthorized() *GetVagrantUsingGETUnauthorized {
	return &GetVagrantUsingGETUnauthorized{}
}

/*
GetVagrantUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVagrantUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get vagrant using g e t unauthorized response has a 2xx status code
func (o *GetVagrantUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vagrant using g e t unauthorized response has a 3xx status code
func (o *GetVagrantUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vagrant using g e t unauthorized response has a 4xx status code
func (o *GetVagrantUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vagrant using g e t unauthorized response has a 5xx status code
func (o *GetVagrantUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get vagrant using g e t unauthorized response a status code equal to that given
func (o *GetVagrantUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get vagrant using g e t unauthorized response
func (o *GetVagrantUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetVagrantUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant][%d] getVagrantUsingGETUnauthorized", 401)
}

func (o *GetVagrantUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant][%d] getVagrantUsingGETUnauthorized", 401)
}

func (o *GetVagrantUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVagrantUsingGETForbidden creates a GetVagrantUsingGETForbidden with default headers values
func NewGetVagrantUsingGETForbidden() *GetVagrantUsingGETForbidden {
	return &GetVagrantUsingGETForbidden{}
}

/*
GetVagrantUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVagrantUsingGETForbidden struct {
}

// IsSuccess returns true when this get vagrant using g e t forbidden response has a 2xx status code
func (o *GetVagrantUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vagrant using g e t forbidden response has a 3xx status code
func (o *GetVagrantUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vagrant using g e t forbidden response has a 4xx status code
func (o *GetVagrantUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vagrant using g e t forbidden response has a 5xx status code
func (o *GetVagrantUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get vagrant using g e t forbidden response a status code equal to that given
func (o *GetVagrantUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get vagrant using g e t forbidden response
func (o *GetVagrantUsingGETForbidden) Code() int {
	return 403
}

func (o *GetVagrantUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant][%d] getVagrantUsingGETForbidden", 403)
}

func (o *GetVagrantUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant][%d] getVagrantUsingGETForbidden", 403)
}

func (o *GetVagrantUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVagrantUsingGETNotFound creates a GetVagrantUsingGETNotFound with default headers values
func NewGetVagrantUsingGETNotFound() *GetVagrantUsingGETNotFound {
	return &GetVagrantUsingGETNotFound{}
}

/*
GetVagrantUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVagrantUsingGETNotFound struct {
}

// IsSuccess returns true when this get vagrant using g e t not found response has a 2xx status code
func (o *GetVagrantUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vagrant using g e t not found response has a 3xx status code
func (o *GetVagrantUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vagrant using g e t not found response has a 4xx status code
func (o *GetVagrantUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vagrant using g e t not found response has a 5xx status code
func (o *GetVagrantUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get vagrant using g e t not found response a status code equal to that given
func (o *GetVagrantUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get vagrant using g e t not found response
func (o *GetVagrantUsingGETNotFound) Code() int {
	return 404
}

func (o *GetVagrantUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant][%d] getVagrantUsingGETNotFound", 404)
}

func (o *GetVagrantUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/local/clusters/{clusterId}/vagrant][%d] getVagrantUsingGETNotFound", 404)
}

func (o *GetVagrantUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
