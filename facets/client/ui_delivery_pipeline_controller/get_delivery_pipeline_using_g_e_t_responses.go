// Code generated by go-swagger; DO NOT EDIT.

package ui_delivery_pipeline_controller

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

// GetDeliveryPipelineUsingGETReader is a Reader for the GetDeliveryPipelineUsingGET structure.
type GetDeliveryPipelineUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeliveryPipelineUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeliveryPipelineUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeliveryPipelineUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeliveryPipelineUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeliveryPipelineUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/delivery-pipeline/{stackName}] getDeliveryPipelineUsingGET", response, response.Code())
	}
}

// NewGetDeliveryPipelineUsingGETOK creates a GetDeliveryPipelineUsingGETOK with default headers values
func NewGetDeliveryPipelineUsingGETOK() *GetDeliveryPipelineUsingGETOK {
	return &GetDeliveryPipelineUsingGETOK{}
}

/*
GetDeliveryPipelineUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetDeliveryPipelineUsingGETOK struct {
	Payload []*models.PipelineNode
}

// IsSuccess returns true when this get delivery pipeline using g e t o k response has a 2xx status code
func (o *GetDeliveryPipelineUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get delivery pipeline using g e t o k response has a 3xx status code
func (o *GetDeliveryPipelineUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery pipeline using g e t o k response has a 4xx status code
func (o *GetDeliveryPipelineUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get delivery pipeline using g e t o k response has a 5xx status code
func (o *GetDeliveryPipelineUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery pipeline using g e t o k response a status code equal to that given
func (o *GetDeliveryPipelineUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get delivery pipeline using g e t o k response
func (o *GetDeliveryPipelineUsingGETOK) Code() int {
	return 200
}

func (o *GetDeliveryPipelineUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/delivery-pipeline/{stackName}][%d] getDeliveryPipelineUsingGETOK %s", 200, payload)
}

func (o *GetDeliveryPipelineUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/delivery-pipeline/{stackName}][%d] getDeliveryPipelineUsingGETOK %s", 200, payload)
}

func (o *GetDeliveryPipelineUsingGETOK) GetPayload() []*models.PipelineNode {
	return o.Payload
}

func (o *GetDeliveryPipelineUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeliveryPipelineUsingGETUnauthorized creates a GetDeliveryPipelineUsingGETUnauthorized with default headers values
func NewGetDeliveryPipelineUsingGETUnauthorized() *GetDeliveryPipelineUsingGETUnauthorized {
	return &GetDeliveryPipelineUsingGETUnauthorized{}
}

/*
GetDeliveryPipelineUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeliveryPipelineUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get delivery pipeline using g e t unauthorized response has a 2xx status code
func (o *GetDeliveryPipelineUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery pipeline using g e t unauthorized response has a 3xx status code
func (o *GetDeliveryPipelineUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery pipeline using g e t unauthorized response has a 4xx status code
func (o *GetDeliveryPipelineUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery pipeline using g e t unauthorized response has a 5xx status code
func (o *GetDeliveryPipelineUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery pipeline using g e t unauthorized response a status code equal to that given
func (o *GetDeliveryPipelineUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get delivery pipeline using g e t unauthorized response
func (o *GetDeliveryPipelineUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetDeliveryPipelineUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/delivery-pipeline/{stackName}][%d] getDeliveryPipelineUsingGETUnauthorized", 401)
}

func (o *GetDeliveryPipelineUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/delivery-pipeline/{stackName}][%d] getDeliveryPipelineUsingGETUnauthorized", 401)
}

func (o *GetDeliveryPipelineUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeliveryPipelineUsingGETForbidden creates a GetDeliveryPipelineUsingGETForbidden with default headers values
func NewGetDeliveryPipelineUsingGETForbidden() *GetDeliveryPipelineUsingGETForbidden {
	return &GetDeliveryPipelineUsingGETForbidden{}
}

/*
GetDeliveryPipelineUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDeliveryPipelineUsingGETForbidden struct {
}

// IsSuccess returns true when this get delivery pipeline using g e t forbidden response has a 2xx status code
func (o *GetDeliveryPipelineUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery pipeline using g e t forbidden response has a 3xx status code
func (o *GetDeliveryPipelineUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery pipeline using g e t forbidden response has a 4xx status code
func (o *GetDeliveryPipelineUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery pipeline using g e t forbidden response has a 5xx status code
func (o *GetDeliveryPipelineUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery pipeline using g e t forbidden response a status code equal to that given
func (o *GetDeliveryPipelineUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get delivery pipeline using g e t forbidden response
func (o *GetDeliveryPipelineUsingGETForbidden) Code() int {
	return 403
}

func (o *GetDeliveryPipelineUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/delivery-pipeline/{stackName}][%d] getDeliveryPipelineUsingGETForbidden", 403)
}

func (o *GetDeliveryPipelineUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/delivery-pipeline/{stackName}][%d] getDeliveryPipelineUsingGETForbidden", 403)
}

func (o *GetDeliveryPipelineUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeliveryPipelineUsingGETNotFound creates a GetDeliveryPipelineUsingGETNotFound with default headers values
func NewGetDeliveryPipelineUsingGETNotFound() *GetDeliveryPipelineUsingGETNotFound {
	return &GetDeliveryPipelineUsingGETNotFound{}
}

/*
GetDeliveryPipelineUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeliveryPipelineUsingGETNotFound struct {
}

// IsSuccess returns true when this get delivery pipeline using g e t not found response has a 2xx status code
func (o *GetDeliveryPipelineUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get delivery pipeline using g e t not found response has a 3xx status code
func (o *GetDeliveryPipelineUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get delivery pipeline using g e t not found response has a 4xx status code
func (o *GetDeliveryPipelineUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get delivery pipeline using g e t not found response has a 5xx status code
func (o *GetDeliveryPipelineUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get delivery pipeline using g e t not found response a status code equal to that given
func (o *GetDeliveryPipelineUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get delivery pipeline using g e t not found response
func (o *GetDeliveryPipelineUsingGETNotFound) Code() int {
	return 404
}

func (o *GetDeliveryPipelineUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/delivery-pipeline/{stackName}][%d] getDeliveryPipelineUsingGETNotFound", 404)
}

func (o *GetDeliveryPipelineUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/delivery-pipeline/{stackName}][%d] getDeliveryPipelineUsingGETNotFound", 404)
}

func (o *GetDeliveryPipelineUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
