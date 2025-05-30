// Code generated by go-swagger; DO NOT EDIT.

package ui_tf_output_controller

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

// GetAllOutputsUsingGETReader is a Reader for the GetAllOutputsUsingGET structure.
type GetAllOutputsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllOutputsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllOutputsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllOutputsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllOutputsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllOutputsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/tf-outputs] getAllOutputsUsingGET", response, response.Code())
	}
}

// NewGetAllOutputsUsingGETOK creates a GetAllOutputsUsingGETOK with default headers values
func NewGetAllOutputsUsingGETOK() *GetAllOutputsUsingGETOK {
	return &GetAllOutputsUsingGETOK{}
}

/*
GetAllOutputsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAllOutputsUsingGETOK struct {
	Payload []*models.TFOutput
}

// IsSuccess returns true when this get all outputs using g e t o k response has a 2xx status code
func (o *GetAllOutputsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all outputs using g e t o k response has a 3xx status code
func (o *GetAllOutputsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all outputs using g e t o k response has a 4xx status code
func (o *GetAllOutputsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all outputs using g e t o k response has a 5xx status code
func (o *GetAllOutputsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all outputs using g e t o k response a status code equal to that given
func (o *GetAllOutputsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all outputs using g e t o k response
func (o *GetAllOutputsUsingGETOK) Code() int {
	return 200
}

func (o *GetAllOutputsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/tf-outputs][%d] getAllOutputsUsingGETOK %s", 200, payload)
}

func (o *GetAllOutputsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/tf-outputs][%d] getAllOutputsUsingGETOK %s", 200, payload)
}

func (o *GetAllOutputsUsingGETOK) GetPayload() []*models.TFOutput {
	return o.Payload
}

func (o *GetAllOutputsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllOutputsUsingGETUnauthorized creates a GetAllOutputsUsingGETUnauthorized with default headers values
func NewGetAllOutputsUsingGETUnauthorized() *GetAllOutputsUsingGETUnauthorized {
	return &GetAllOutputsUsingGETUnauthorized{}
}

/*
GetAllOutputsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllOutputsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get all outputs using g e t unauthorized response has a 2xx status code
func (o *GetAllOutputsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all outputs using g e t unauthorized response has a 3xx status code
func (o *GetAllOutputsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all outputs using g e t unauthorized response has a 4xx status code
func (o *GetAllOutputsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all outputs using g e t unauthorized response has a 5xx status code
func (o *GetAllOutputsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all outputs using g e t unauthorized response a status code equal to that given
func (o *GetAllOutputsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all outputs using g e t unauthorized response
func (o *GetAllOutputsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetAllOutputsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/tf-outputs][%d] getAllOutputsUsingGETUnauthorized", 401)
}

func (o *GetAllOutputsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/tf-outputs][%d] getAllOutputsUsingGETUnauthorized", 401)
}

func (o *GetAllOutputsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllOutputsUsingGETForbidden creates a GetAllOutputsUsingGETForbidden with default headers values
func NewGetAllOutputsUsingGETForbidden() *GetAllOutputsUsingGETForbidden {
	return &GetAllOutputsUsingGETForbidden{}
}

/*
GetAllOutputsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllOutputsUsingGETForbidden struct {
}

// IsSuccess returns true when this get all outputs using g e t forbidden response has a 2xx status code
func (o *GetAllOutputsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all outputs using g e t forbidden response has a 3xx status code
func (o *GetAllOutputsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all outputs using g e t forbidden response has a 4xx status code
func (o *GetAllOutputsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all outputs using g e t forbidden response has a 5xx status code
func (o *GetAllOutputsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all outputs using g e t forbidden response a status code equal to that given
func (o *GetAllOutputsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get all outputs using g e t forbidden response
func (o *GetAllOutputsUsingGETForbidden) Code() int {
	return 403
}

func (o *GetAllOutputsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/tf-outputs][%d] getAllOutputsUsingGETForbidden", 403)
}

func (o *GetAllOutputsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/tf-outputs][%d] getAllOutputsUsingGETForbidden", 403)
}

func (o *GetAllOutputsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllOutputsUsingGETNotFound creates a GetAllOutputsUsingGETNotFound with default headers values
func NewGetAllOutputsUsingGETNotFound() *GetAllOutputsUsingGETNotFound {
	return &GetAllOutputsUsingGETNotFound{}
}

/*
GetAllOutputsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllOutputsUsingGETNotFound struct {
}

// IsSuccess returns true when this get all outputs using g e t not found response has a 2xx status code
func (o *GetAllOutputsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all outputs using g e t not found response has a 3xx status code
func (o *GetAllOutputsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all outputs using g e t not found response has a 4xx status code
func (o *GetAllOutputsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all outputs using g e t not found response has a 5xx status code
func (o *GetAllOutputsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all outputs using g e t not found response a status code equal to that given
func (o *GetAllOutputsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all outputs using g e t not found response
func (o *GetAllOutputsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetAllOutputsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/tf-outputs][%d] getAllOutputsUsingGETNotFound", 404)
}

func (o *GetAllOutputsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/tf-outputs][%d] getAllOutputsUsingGETNotFound", 404)
}

func (o *GetAllOutputsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
