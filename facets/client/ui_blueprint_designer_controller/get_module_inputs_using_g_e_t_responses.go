// Code generated by go-swagger; DO NOT EDIT.

package ui_blueprint_designer_controller

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

// GetModuleInputsUsingGETReader is a Reader for the GetModuleInputsUsingGET structure.
type GetModuleInputsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModuleInputsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModuleInputsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetModuleInputsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetModuleInputsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetModuleInputsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input] getModuleInputsUsingGET", response, response.Code())
	}
}

// NewGetModuleInputsUsingGETOK creates a GetModuleInputsUsingGETOK with default headers values
func NewGetModuleInputsUsingGETOK() *GetModuleInputsUsingGETOK {
	return &GetModuleInputsUsingGETOK{}
}

/*
GetModuleInputsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetModuleInputsUsingGETOK struct {
	Payload map[string]models.ModuleInputDTO
}

// IsSuccess returns true when this get module inputs using g e t o k response has a 2xx status code
func (o *GetModuleInputsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get module inputs using g e t o k response has a 3xx status code
func (o *GetModuleInputsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get module inputs using g e t o k response has a 4xx status code
func (o *GetModuleInputsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get module inputs using g e t o k response has a 5xx status code
func (o *GetModuleInputsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get module inputs using g e t o k response a status code equal to that given
func (o *GetModuleInputsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get module inputs using g e t o k response
func (o *GetModuleInputsUsingGETOK) Code() int {
	return 200
}

func (o *GetModuleInputsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input][%d] getModuleInputsUsingGETOK %s", 200, payload)
}

func (o *GetModuleInputsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input][%d] getModuleInputsUsingGETOK %s", 200, payload)
}

func (o *GetModuleInputsUsingGETOK) GetPayload() map[string]models.ModuleInputDTO {
	return o.Payload
}

func (o *GetModuleInputsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModuleInputsUsingGETUnauthorized creates a GetModuleInputsUsingGETUnauthorized with default headers values
func NewGetModuleInputsUsingGETUnauthorized() *GetModuleInputsUsingGETUnauthorized {
	return &GetModuleInputsUsingGETUnauthorized{}
}

/*
GetModuleInputsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetModuleInputsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get module inputs using g e t unauthorized response has a 2xx status code
func (o *GetModuleInputsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get module inputs using g e t unauthorized response has a 3xx status code
func (o *GetModuleInputsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get module inputs using g e t unauthorized response has a 4xx status code
func (o *GetModuleInputsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get module inputs using g e t unauthorized response has a 5xx status code
func (o *GetModuleInputsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get module inputs using g e t unauthorized response a status code equal to that given
func (o *GetModuleInputsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get module inputs using g e t unauthorized response
func (o *GetModuleInputsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetModuleInputsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input][%d] getModuleInputsUsingGETUnauthorized", 401)
}

func (o *GetModuleInputsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input][%d] getModuleInputsUsingGETUnauthorized", 401)
}

func (o *GetModuleInputsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetModuleInputsUsingGETForbidden creates a GetModuleInputsUsingGETForbidden with default headers values
func NewGetModuleInputsUsingGETForbidden() *GetModuleInputsUsingGETForbidden {
	return &GetModuleInputsUsingGETForbidden{}
}

/*
GetModuleInputsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetModuleInputsUsingGETForbidden struct {
}

// IsSuccess returns true when this get module inputs using g e t forbidden response has a 2xx status code
func (o *GetModuleInputsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get module inputs using g e t forbidden response has a 3xx status code
func (o *GetModuleInputsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get module inputs using g e t forbidden response has a 4xx status code
func (o *GetModuleInputsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get module inputs using g e t forbidden response has a 5xx status code
func (o *GetModuleInputsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get module inputs using g e t forbidden response a status code equal to that given
func (o *GetModuleInputsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get module inputs using g e t forbidden response
func (o *GetModuleInputsUsingGETForbidden) Code() int {
	return 403
}

func (o *GetModuleInputsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input][%d] getModuleInputsUsingGETForbidden", 403)
}

func (o *GetModuleInputsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input][%d] getModuleInputsUsingGETForbidden", 403)
}

func (o *GetModuleInputsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetModuleInputsUsingGETNotFound creates a GetModuleInputsUsingGETNotFound with default headers values
func NewGetModuleInputsUsingGETNotFound() *GetModuleInputsUsingGETNotFound {
	return &GetModuleInputsUsingGETNotFound{}
}

/*
GetModuleInputsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetModuleInputsUsingGETNotFound struct {
}

// IsSuccess returns true when this get module inputs using g e t not found response has a 2xx status code
func (o *GetModuleInputsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get module inputs using g e t not found response has a 3xx status code
func (o *GetModuleInputsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get module inputs using g e t not found response has a 4xx status code
func (o *GetModuleInputsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get module inputs using g e t not found response has a 5xx status code
func (o *GetModuleInputsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get module inputs using g e t not found response a status code equal to that given
func (o *GetModuleInputsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get module inputs using g e t not found response
func (o *GetModuleInputsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetModuleInputsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input][%d] getModuleInputsUsingGETNotFound", 404)
}

func (o *GetModuleInputsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/intent/{intent}/flavor/{flavor}/input][%d] getModuleInputsUsingGETNotFound", 404)
}

func (o *GetModuleInputsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
