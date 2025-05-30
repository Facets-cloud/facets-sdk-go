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

	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

// GetResourceOverrideObjectUsingGETReader is a Reader for the GetResourceOverrideObjectUsingGET structure.
type GetResourceOverrideObjectUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceOverrideObjectUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceOverrideObjectUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceOverrideObjectUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceOverrideObjectUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceOverrideObjectUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides] getResourceOverrideObjectUsingGET", response, response.Code())
	}
}

// NewGetResourceOverrideObjectUsingGETOK creates a GetResourceOverrideObjectUsingGETOK with default headers values
func NewGetResourceOverrideObjectUsingGETOK() *GetResourceOverrideObjectUsingGETOK {
	return &GetResourceOverrideObjectUsingGETOK{}
}

/*
GetResourceOverrideObjectUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetResourceOverrideObjectUsingGETOK struct {
	Payload *models.OverrideObject
}

// IsSuccess returns true when this get resource override object using g e t o k response has a 2xx status code
func (o *GetResourceOverrideObjectUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource override object using g e t o k response has a 3xx status code
func (o *GetResourceOverrideObjectUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource override object using g e t o k response has a 4xx status code
func (o *GetResourceOverrideObjectUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource override object using g e t o k response has a 5xx status code
func (o *GetResourceOverrideObjectUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource override object using g e t o k response a status code equal to that given
func (o *GetResourceOverrideObjectUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get resource override object using g e t o k response
func (o *GetResourceOverrideObjectUsingGETOK) Code() int {
	return 200
}

func (o *GetResourceOverrideObjectUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides][%d] getResourceOverrideObjectUsingGETOK %s", 200, payload)
}

func (o *GetResourceOverrideObjectUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides][%d] getResourceOverrideObjectUsingGETOK %s", 200, payload)
}

func (o *GetResourceOverrideObjectUsingGETOK) GetPayload() *models.OverrideObject {
	return o.Payload
}

func (o *GetResourceOverrideObjectUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OverrideObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceOverrideObjectUsingGETUnauthorized creates a GetResourceOverrideObjectUsingGETUnauthorized with default headers values
func NewGetResourceOverrideObjectUsingGETUnauthorized() *GetResourceOverrideObjectUsingGETUnauthorized {
	return &GetResourceOverrideObjectUsingGETUnauthorized{}
}

/*
GetResourceOverrideObjectUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetResourceOverrideObjectUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get resource override object using g e t unauthorized response has a 2xx status code
func (o *GetResourceOverrideObjectUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource override object using g e t unauthorized response has a 3xx status code
func (o *GetResourceOverrideObjectUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource override object using g e t unauthorized response has a 4xx status code
func (o *GetResourceOverrideObjectUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource override object using g e t unauthorized response has a 5xx status code
func (o *GetResourceOverrideObjectUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource override object using g e t unauthorized response a status code equal to that given
func (o *GetResourceOverrideObjectUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get resource override object using g e t unauthorized response
func (o *GetResourceOverrideObjectUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetResourceOverrideObjectUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides][%d] getResourceOverrideObjectUsingGETUnauthorized", 401)
}

func (o *GetResourceOverrideObjectUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides][%d] getResourceOverrideObjectUsingGETUnauthorized", 401)
}

func (o *GetResourceOverrideObjectUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceOverrideObjectUsingGETForbidden creates a GetResourceOverrideObjectUsingGETForbidden with default headers values
func NewGetResourceOverrideObjectUsingGETForbidden() *GetResourceOverrideObjectUsingGETForbidden {
	return &GetResourceOverrideObjectUsingGETForbidden{}
}

/*
GetResourceOverrideObjectUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetResourceOverrideObjectUsingGETForbidden struct {
}

// IsSuccess returns true when this get resource override object using g e t forbidden response has a 2xx status code
func (o *GetResourceOverrideObjectUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource override object using g e t forbidden response has a 3xx status code
func (o *GetResourceOverrideObjectUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource override object using g e t forbidden response has a 4xx status code
func (o *GetResourceOverrideObjectUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource override object using g e t forbidden response has a 5xx status code
func (o *GetResourceOverrideObjectUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource override object using g e t forbidden response a status code equal to that given
func (o *GetResourceOverrideObjectUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get resource override object using g e t forbidden response
func (o *GetResourceOverrideObjectUsingGETForbidden) Code() int {
	return 403
}

func (o *GetResourceOverrideObjectUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides][%d] getResourceOverrideObjectUsingGETForbidden", 403)
}

func (o *GetResourceOverrideObjectUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides][%d] getResourceOverrideObjectUsingGETForbidden", 403)
}

func (o *GetResourceOverrideObjectUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceOverrideObjectUsingGETNotFound creates a GetResourceOverrideObjectUsingGETNotFound with default headers values
func NewGetResourceOverrideObjectUsingGETNotFound() *GetResourceOverrideObjectUsingGETNotFound {
	return &GetResourceOverrideObjectUsingGETNotFound{}
}

/*
GetResourceOverrideObjectUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetResourceOverrideObjectUsingGETNotFound struct {
}

// IsSuccess returns true when this get resource override object using g e t not found response has a 2xx status code
func (o *GetResourceOverrideObjectUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource override object using g e t not found response has a 3xx status code
func (o *GetResourceOverrideObjectUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource override object using g e t not found response has a 4xx status code
func (o *GetResourceOverrideObjectUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource override object using g e t not found response has a 5xx status code
func (o *GetResourceOverrideObjectUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource override object using g e t not found response a status code equal to that given
func (o *GetResourceOverrideObjectUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get resource override object using g e t not found response
func (o *GetResourceOverrideObjectUsingGETNotFound) Code() int {
	return 404
}

func (o *GetResourceOverrideObjectUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides][%d] getResourceOverrideObjectUsingGETNotFound", 404)
}

func (o *GetResourceOverrideObjectUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/resourceType/{resourceType}/resourceName/{resourceName}/overrides][%d] getResourceOverrideObjectUsingGETNotFound", 404)
}

func (o *GetResourceOverrideObjectUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
