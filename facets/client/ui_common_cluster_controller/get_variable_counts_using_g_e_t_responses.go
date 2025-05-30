// Code generated by go-swagger; DO NOT EDIT.

package ui_common_cluster_controller

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

// GetVariableCountsUsingGETReader is a Reader for the GetVariableCountsUsingGET structure.
type GetVariableCountsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVariableCountsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVariableCountsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVariableCountsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVariableCountsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVariableCountsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts] getVariableCountsUsingGET", response, response.Code())
	}
}

// NewGetVariableCountsUsingGETOK creates a GetVariableCountsUsingGETOK with default headers values
func NewGetVariableCountsUsingGETOK() *GetVariableCountsUsingGETOK {
	return &GetVariableCountsUsingGETOK{}
}

/*
GetVariableCountsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetVariableCountsUsingGETOK struct {
	Payload *models.VariableCountDto
}

// IsSuccess returns true when this get variable counts using g e t o k response has a 2xx status code
func (o *GetVariableCountsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get variable counts using g e t o k response has a 3xx status code
func (o *GetVariableCountsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable counts using g e t o k response has a 4xx status code
func (o *GetVariableCountsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get variable counts using g e t o k response has a 5xx status code
func (o *GetVariableCountsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable counts using g e t o k response a status code equal to that given
func (o *GetVariableCountsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get variable counts using g e t o k response
func (o *GetVariableCountsUsingGETOK) Code() int {
	return 200
}

func (o *GetVariableCountsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts][%d] getVariableCountsUsingGETOK %s", 200, payload)
}

func (o *GetVariableCountsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts][%d] getVariableCountsUsingGETOK %s", 200, payload)
}

func (o *GetVariableCountsUsingGETOK) GetPayload() *models.VariableCountDto {
	return o.Payload
}

func (o *GetVariableCountsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableCountDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVariableCountsUsingGETUnauthorized creates a GetVariableCountsUsingGETUnauthorized with default headers values
func NewGetVariableCountsUsingGETUnauthorized() *GetVariableCountsUsingGETUnauthorized {
	return &GetVariableCountsUsingGETUnauthorized{}
}

/*
GetVariableCountsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVariableCountsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get variable counts using g e t unauthorized response has a 2xx status code
func (o *GetVariableCountsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get variable counts using g e t unauthorized response has a 3xx status code
func (o *GetVariableCountsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable counts using g e t unauthorized response has a 4xx status code
func (o *GetVariableCountsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get variable counts using g e t unauthorized response has a 5xx status code
func (o *GetVariableCountsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable counts using g e t unauthorized response a status code equal to that given
func (o *GetVariableCountsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get variable counts using g e t unauthorized response
func (o *GetVariableCountsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetVariableCountsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts][%d] getVariableCountsUsingGETUnauthorized", 401)
}

func (o *GetVariableCountsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts][%d] getVariableCountsUsingGETUnauthorized", 401)
}

func (o *GetVariableCountsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableCountsUsingGETForbidden creates a GetVariableCountsUsingGETForbidden with default headers values
func NewGetVariableCountsUsingGETForbidden() *GetVariableCountsUsingGETForbidden {
	return &GetVariableCountsUsingGETForbidden{}
}

/*
GetVariableCountsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVariableCountsUsingGETForbidden struct {
}

// IsSuccess returns true when this get variable counts using g e t forbidden response has a 2xx status code
func (o *GetVariableCountsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get variable counts using g e t forbidden response has a 3xx status code
func (o *GetVariableCountsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable counts using g e t forbidden response has a 4xx status code
func (o *GetVariableCountsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get variable counts using g e t forbidden response has a 5xx status code
func (o *GetVariableCountsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable counts using g e t forbidden response a status code equal to that given
func (o *GetVariableCountsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get variable counts using g e t forbidden response
func (o *GetVariableCountsUsingGETForbidden) Code() int {
	return 403
}

func (o *GetVariableCountsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts][%d] getVariableCountsUsingGETForbidden", 403)
}

func (o *GetVariableCountsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts][%d] getVariableCountsUsingGETForbidden", 403)
}

func (o *GetVariableCountsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVariableCountsUsingGETNotFound creates a GetVariableCountsUsingGETNotFound with default headers values
func NewGetVariableCountsUsingGETNotFound() *GetVariableCountsUsingGETNotFound {
	return &GetVariableCountsUsingGETNotFound{}
}

/*
GetVariableCountsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVariableCountsUsingGETNotFound struct {
}

// IsSuccess returns true when this get variable counts using g e t not found response has a 2xx status code
func (o *GetVariableCountsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get variable counts using g e t not found response has a 3xx status code
func (o *GetVariableCountsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get variable counts using g e t not found response has a 4xx status code
func (o *GetVariableCountsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get variable counts using g e t not found response has a 5xx status code
func (o *GetVariableCountsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get variable counts using g e t not found response a status code equal to that given
func (o *GetVariableCountsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get variable counts using g e t not found response
func (o *GetVariableCountsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetVariableCountsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts][%d] getVariableCountsUsingGETNotFound", 404)
}

func (o *GetVariableCountsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/variable-counts][%d] getVariableCountsUsingGETNotFound", 404)
}

func (o *GetVariableCountsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
