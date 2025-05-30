// Code generated by go-swagger; DO NOT EDIT.

package meta_controller

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

// GetSupportedComponentVersionsUsingGETReader is a Reader for the GetSupportedComponentVersionsUsingGET structure.
type GetSupportedComponentVersionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSupportedComponentVersionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSupportedComponentVersionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSupportedComponentVersionsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSupportedComponentVersionsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSupportedComponentVersionsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc/v1/meta/components/supportedVersion] getSupportedComponentVersionsUsingGET", response, response.Code())
	}
}

// NewGetSupportedComponentVersionsUsingGETOK creates a GetSupportedComponentVersionsUsingGETOK with default headers values
func NewGetSupportedComponentVersionsUsingGETOK() *GetSupportedComponentVersionsUsingGETOK {
	return &GetSupportedComponentVersionsUsingGETOK{}
}

/*
GetSupportedComponentVersionsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetSupportedComponentVersionsUsingGETOK struct {
	Payload []*models.SupportedVersions
}

// IsSuccess returns true when this get supported component versions using g e t o k response has a 2xx status code
func (o *GetSupportedComponentVersionsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get supported component versions using g e t o k response has a 3xx status code
func (o *GetSupportedComponentVersionsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supported component versions using g e t o k response has a 4xx status code
func (o *GetSupportedComponentVersionsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get supported component versions using g e t o k response has a 5xx status code
func (o *GetSupportedComponentVersionsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get supported component versions using g e t o k response a status code equal to that given
func (o *GetSupportedComponentVersionsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get supported component versions using g e t o k response
func (o *GetSupportedComponentVersionsUsingGETOK) Code() int {
	return 200
}

func (o *GetSupportedComponentVersionsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc/v1/meta/components/supportedVersion][%d] getSupportedComponentVersionsUsingGETOK %s", 200, payload)
}

func (o *GetSupportedComponentVersionsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc/v1/meta/components/supportedVersion][%d] getSupportedComponentVersionsUsingGETOK %s", 200, payload)
}

func (o *GetSupportedComponentVersionsUsingGETOK) GetPayload() []*models.SupportedVersions {
	return o.Payload
}

func (o *GetSupportedComponentVersionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedComponentVersionsUsingGETUnauthorized creates a GetSupportedComponentVersionsUsingGETUnauthorized with default headers values
func NewGetSupportedComponentVersionsUsingGETUnauthorized() *GetSupportedComponentVersionsUsingGETUnauthorized {
	return &GetSupportedComponentVersionsUsingGETUnauthorized{}
}

/*
GetSupportedComponentVersionsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSupportedComponentVersionsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get supported component versions using g e t unauthorized response has a 2xx status code
func (o *GetSupportedComponentVersionsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supported component versions using g e t unauthorized response has a 3xx status code
func (o *GetSupportedComponentVersionsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supported component versions using g e t unauthorized response has a 4xx status code
func (o *GetSupportedComponentVersionsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supported component versions using g e t unauthorized response has a 5xx status code
func (o *GetSupportedComponentVersionsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get supported component versions using g e t unauthorized response a status code equal to that given
func (o *GetSupportedComponentVersionsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get supported component versions using g e t unauthorized response
func (o *GetSupportedComponentVersionsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetSupportedComponentVersionsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc/v1/meta/components/supportedVersion][%d] getSupportedComponentVersionsUsingGETUnauthorized", 401)
}

func (o *GetSupportedComponentVersionsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc/v1/meta/components/supportedVersion][%d] getSupportedComponentVersionsUsingGETUnauthorized", 401)
}

func (o *GetSupportedComponentVersionsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSupportedComponentVersionsUsingGETForbidden creates a GetSupportedComponentVersionsUsingGETForbidden with default headers values
func NewGetSupportedComponentVersionsUsingGETForbidden() *GetSupportedComponentVersionsUsingGETForbidden {
	return &GetSupportedComponentVersionsUsingGETForbidden{}
}

/*
GetSupportedComponentVersionsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSupportedComponentVersionsUsingGETForbidden struct {
}

// IsSuccess returns true when this get supported component versions using g e t forbidden response has a 2xx status code
func (o *GetSupportedComponentVersionsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supported component versions using g e t forbidden response has a 3xx status code
func (o *GetSupportedComponentVersionsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supported component versions using g e t forbidden response has a 4xx status code
func (o *GetSupportedComponentVersionsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supported component versions using g e t forbidden response has a 5xx status code
func (o *GetSupportedComponentVersionsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get supported component versions using g e t forbidden response a status code equal to that given
func (o *GetSupportedComponentVersionsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get supported component versions using g e t forbidden response
func (o *GetSupportedComponentVersionsUsingGETForbidden) Code() int {
	return 403
}

func (o *GetSupportedComponentVersionsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc/v1/meta/components/supportedVersion][%d] getSupportedComponentVersionsUsingGETForbidden", 403)
}

func (o *GetSupportedComponentVersionsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc/v1/meta/components/supportedVersion][%d] getSupportedComponentVersionsUsingGETForbidden", 403)
}

func (o *GetSupportedComponentVersionsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSupportedComponentVersionsUsingGETNotFound creates a GetSupportedComponentVersionsUsingGETNotFound with default headers values
func NewGetSupportedComponentVersionsUsingGETNotFound() *GetSupportedComponentVersionsUsingGETNotFound {
	return &GetSupportedComponentVersionsUsingGETNotFound{}
}

/*
GetSupportedComponentVersionsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSupportedComponentVersionsUsingGETNotFound struct {
}

// IsSuccess returns true when this get supported component versions using g e t not found response has a 2xx status code
func (o *GetSupportedComponentVersionsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supported component versions using g e t not found response has a 3xx status code
func (o *GetSupportedComponentVersionsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supported component versions using g e t not found response has a 4xx status code
func (o *GetSupportedComponentVersionsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supported component versions using g e t not found response has a 5xx status code
func (o *GetSupportedComponentVersionsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get supported component versions using g e t not found response a status code equal to that given
func (o *GetSupportedComponentVersionsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get supported component versions using g e t not found response
func (o *GetSupportedComponentVersionsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetSupportedComponentVersionsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc/v1/meta/components/supportedVersion][%d] getSupportedComponentVersionsUsingGETNotFound", 404)
}

func (o *GetSupportedComponentVersionsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc/v1/meta/components/supportedVersion][%d] getSupportedComponentVersionsUsingGETNotFound", 404)
}

func (o *GetSupportedComponentVersionsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
