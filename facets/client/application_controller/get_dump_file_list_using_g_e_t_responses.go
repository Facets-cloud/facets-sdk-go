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
)

// GetDumpFileListUsingGETReader is a Reader for the GetDumpFileListUsingGET structure.
type GetDumpFileListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDumpFileListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDumpFileListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDumpFileListUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDumpFileListUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDumpFileListUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps] getDumpFileListUsingGET", response, response.Code())
	}
}

// NewGetDumpFileListUsingGETOK creates a GetDumpFileListUsingGETOK with default headers values
func NewGetDumpFileListUsingGETOK() *GetDumpFileListUsingGETOK {
	return &GetDumpFileListUsingGETOK{}
}

/*
GetDumpFileListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetDumpFileListUsingGETOK struct {
	Payload map[string]string
}

// IsSuccess returns true when this get dump file list using g e t o k response has a 2xx status code
func (o *GetDumpFileListUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dump file list using g e t o k response has a 3xx status code
func (o *GetDumpFileListUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dump file list using g e t o k response has a 4xx status code
func (o *GetDumpFileListUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dump file list using g e t o k response has a 5xx status code
func (o *GetDumpFileListUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dump file list using g e t o k response a status code equal to that given
func (o *GetDumpFileListUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get dump file list using g e t o k response
func (o *GetDumpFileListUsingGETOK) Code() int {
	return 200
}

func (o *GetDumpFileListUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps][%d] getDumpFileListUsingGETOK %s", 200, payload)
}

func (o *GetDumpFileListUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps][%d] getDumpFileListUsingGETOK %s", 200, payload)
}

func (o *GetDumpFileListUsingGETOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *GetDumpFileListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDumpFileListUsingGETUnauthorized creates a GetDumpFileListUsingGETUnauthorized with default headers values
func NewGetDumpFileListUsingGETUnauthorized() *GetDumpFileListUsingGETUnauthorized {
	return &GetDumpFileListUsingGETUnauthorized{}
}

/*
GetDumpFileListUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDumpFileListUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get dump file list using g e t unauthorized response has a 2xx status code
func (o *GetDumpFileListUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dump file list using g e t unauthorized response has a 3xx status code
func (o *GetDumpFileListUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dump file list using g e t unauthorized response has a 4xx status code
func (o *GetDumpFileListUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dump file list using g e t unauthorized response has a 5xx status code
func (o *GetDumpFileListUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get dump file list using g e t unauthorized response a status code equal to that given
func (o *GetDumpFileListUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get dump file list using g e t unauthorized response
func (o *GetDumpFileListUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetDumpFileListUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps][%d] getDumpFileListUsingGETUnauthorized", 401)
}

func (o *GetDumpFileListUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps][%d] getDumpFileListUsingGETUnauthorized", 401)
}

func (o *GetDumpFileListUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDumpFileListUsingGETForbidden creates a GetDumpFileListUsingGETForbidden with default headers values
func NewGetDumpFileListUsingGETForbidden() *GetDumpFileListUsingGETForbidden {
	return &GetDumpFileListUsingGETForbidden{}
}

/*
GetDumpFileListUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDumpFileListUsingGETForbidden struct {
}

// IsSuccess returns true when this get dump file list using g e t forbidden response has a 2xx status code
func (o *GetDumpFileListUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dump file list using g e t forbidden response has a 3xx status code
func (o *GetDumpFileListUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dump file list using g e t forbidden response has a 4xx status code
func (o *GetDumpFileListUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dump file list using g e t forbidden response has a 5xx status code
func (o *GetDumpFileListUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get dump file list using g e t forbidden response a status code equal to that given
func (o *GetDumpFileListUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get dump file list using g e t forbidden response
func (o *GetDumpFileListUsingGETForbidden) Code() int {
	return 403
}

func (o *GetDumpFileListUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps][%d] getDumpFileListUsingGETForbidden", 403)
}

func (o *GetDumpFileListUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps][%d] getDumpFileListUsingGETForbidden", 403)
}

func (o *GetDumpFileListUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDumpFileListUsingGETNotFound creates a GetDumpFileListUsingGETNotFound with default headers values
func NewGetDumpFileListUsingGETNotFound() *GetDumpFileListUsingGETNotFound {
	return &GetDumpFileListUsingGETNotFound{}
}

/*
GetDumpFileListUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDumpFileListUsingGETNotFound struct {
}

// IsSuccess returns true when this get dump file list using g e t not found response has a 2xx status code
func (o *GetDumpFileListUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dump file list using g e t not found response has a 3xx status code
func (o *GetDumpFileListUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dump file list using g e t not found response has a 4xx status code
func (o *GetDumpFileListUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dump file list using g e t not found response has a 5xx status code
func (o *GetDumpFileListUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get dump file list using g e t not found response a status code equal to that given
func (o *GetDumpFileListUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get dump file list using g e t not found response
func (o *GetDumpFileListUsingGETNotFound) Code() int {
	return 404
}

func (o *GetDumpFileListUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps][%d] getDumpFileListUsingGETNotFound", 404)
}

func (o *GetDumpFileListUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps][%d] getDumpFileListUsingGETNotFound", 404)
}

func (o *GetDumpFileListUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
