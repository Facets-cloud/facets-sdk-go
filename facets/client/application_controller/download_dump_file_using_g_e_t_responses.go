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

// DownloadDumpFileUsingGETReader is a Reader for the DownloadDumpFileUsingGET structure.
type DownloadDumpFileUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadDumpFileUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadDumpFileUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDownloadDumpFileUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadDumpFileUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadDumpFileUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download] downloadDumpFileUsingGET", response, response.Code())
	}
}

// NewDownloadDumpFileUsingGETOK creates a DownloadDumpFileUsingGETOK with default headers values
func NewDownloadDumpFileUsingGETOK() *DownloadDumpFileUsingGETOK {
	return &DownloadDumpFileUsingGETOK{}
}

/*
DownloadDumpFileUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type DownloadDumpFileUsingGETOK struct {
	Payload *models.InputStreamResource
}

// IsSuccess returns true when this download dump file using g e t o k response has a 2xx status code
func (o *DownloadDumpFileUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this download dump file using g e t o k response has a 3xx status code
func (o *DownloadDumpFileUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download dump file using g e t o k response has a 4xx status code
func (o *DownloadDumpFileUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this download dump file using g e t o k response has a 5xx status code
func (o *DownloadDumpFileUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this download dump file using g e t o k response a status code equal to that given
func (o *DownloadDumpFileUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the download dump file using g e t o k response
func (o *DownloadDumpFileUsingGETOK) Code() int {
	return 200
}

func (o *DownloadDumpFileUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download][%d] downloadDumpFileUsingGETOK %s", 200, payload)
}

func (o *DownloadDumpFileUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download][%d] downloadDumpFileUsingGETOK %s", 200, payload)
}

func (o *DownloadDumpFileUsingGETOK) GetPayload() *models.InputStreamResource {
	return o.Payload
}

func (o *DownloadDumpFileUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InputStreamResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadDumpFileUsingGETUnauthorized creates a DownloadDumpFileUsingGETUnauthorized with default headers values
func NewDownloadDumpFileUsingGETUnauthorized() *DownloadDumpFileUsingGETUnauthorized {
	return &DownloadDumpFileUsingGETUnauthorized{}
}

/*
DownloadDumpFileUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DownloadDumpFileUsingGETUnauthorized struct {
}

// IsSuccess returns true when this download dump file using g e t unauthorized response has a 2xx status code
func (o *DownloadDumpFileUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download dump file using g e t unauthorized response has a 3xx status code
func (o *DownloadDumpFileUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download dump file using g e t unauthorized response has a 4xx status code
func (o *DownloadDumpFileUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this download dump file using g e t unauthorized response has a 5xx status code
func (o *DownloadDumpFileUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this download dump file using g e t unauthorized response a status code equal to that given
func (o *DownloadDumpFileUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the download dump file using g e t unauthorized response
func (o *DownloadDumpFileUsingGETUnauthorized) Code() int {
	return 401
}

func (o *DownloadDumpFileUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download][%d] downloadDumpFileUsingGETUnauthorized", 401)
}

func (o *DownloadDumpFileUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download][%d] downloadDumpFileUsingGETUnauthorized", 401)
}

func (o *DownloadDumpFileUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadDumpFileUsingGETForbidden creates a DownloadDumpFileUsingGETForbidden with default headers values
func NewDownloadDumpFileUsingGETForbidden() *DownloadDumpFileUsingGETForbidden {
	return &DownloadDumpFileUsingGETForbidden{}
}

/*
DownloadDumpFileUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DownloadDumpFileUsingGETForbidden struct {
}

// IsSuccess returns true when this download dump file using g e t forbidden response has a 2xx status code
func (o *DownloadDumpFileUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download dump file using g e t forbidden response has a 3xx status code
func (o *DownloadDumpFileUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download dump file using g e t forbidden response has a 4xx status code
func (o *DownloadDumpFileUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this download dump file using g e t forbidden response has a 5xx status code
func (o *DownloadDumpFileUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this download dump file using g e t forbidden response a status code equal to that given
func (o *DownloadDumpFileUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the download dump file using g e t forbidden response
func (o *DownloadDumpFileUsingGETForbidden) Code() int {
	return 403
}

func (o *DownloadDumpFileUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download][%d] downloadDumpFileUsingGETForbidden", 403)
}

func (o *DownloadDumpFileUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download][%d] downloadDumpFileUsingGETForbidden", 403)
}

func (o *DownloadDumpFileUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadDumpFileUsingGETNotFound creates a DownloadDumpFileUsingGETNotFound with default headers values
func NewDownloadDumpFileUsingGETNotFound() *DownloadDumpFileUsingGETNotFound {
	return &DownloadDumpFileUsingGETNotFound{}
}

/*
DownloadDumpFileUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DownloadDumpFileUsingGETNotFound struct {
}

// IsSuccess returns true when this download dump file using g e t not found response has a 2xx status code
func (o *DownloadDumpFileUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this download dump file using g e t not found response has a 3xx status code
func (o *DownloadDumpFileUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this download dump file using g e t not found response has a 4xx status code
func (o *DownloadDumpFileUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this download dump file using g e t not found response has a 5xx status code
func (o *DownloadDumpFileUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this download dump file using g e t not found response a status code equal to that given
func (o *DownloadDumpFileUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the download dump file using g e t not found response
func (o *DownloadDumpFileUsingGETNotFound) Code() int {
	return 404
}

func (o *DownloadDumpFileUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download][%d] downloadDumpFileUsingGETNotFound", 404)
}

func (o *DownloadDumpFileUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /api/{applicationFamily}/{environment}/applications/{applicationId}/dumps/download][%d] downloadDumpFileUsingGETNotFound", 404)
}

func (o *DownloadDumpFileUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
