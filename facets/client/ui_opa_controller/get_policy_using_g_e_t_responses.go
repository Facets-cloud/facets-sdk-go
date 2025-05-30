// Code generated by go-swagger; DO NOT EDIT.

package ui_opa_controller

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

// GetPolicyUsingGETReader is a Reader for the GetPolicyUsingGET structure.
type GetPolicyUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPolicyUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPolicyUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPolicyUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/opa/{policyName}] getPolicyUsingGET", response, response.Code())
	}
}

// NewGetPolicyUsingGETOK creates a GetPolicyUsingGETOK with default headers values
func NewGetPolicyUsingGETOK() *GetPolicyUsingGETOK {
	return &GetPolicyUsingGETOK{}
}

/*
GetPolicyUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetPolicyUsingGETOK struct {
	Payload *models.OpaPolicy
}

// IsSuccess returns true when this get policy using g e t o k response has a 2xx status code
func (o *GetPolicyUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get policy using g e t o k response has a 3xx status code
func (o *GetPolicyUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy using g e t o k response has a 4xx status code
func (o *GetPolicyUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get policy using g e t o k response has a 5xx status code
func (o *GetPolicyUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy using g e t o k response a status code equal to that given
func (o *GetPolicyUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get policy using g e t o k response
func (o *GetPolicyUsingGETOK) Code() int {
	return 200
}

func (o *GetPolicyUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/opa/{policyName}][%d] getPolicyUsingGETOK %s", 200, payload)
}

func (o *GetPolicyUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/opa/{policyName}][%d] getPolicyUsingGETOK %s", 200, payload)
}

func (o *GetPolicyUsingGETOK) GetPayload() *models.OpaPolicy {
	return o.Payload
}

func (o *GetPolicyUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpaPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyUsingGETUnauthorized creates a GetPolicyUsingGETUnauthorized with default headers values
func NewGetPolicyUsingGETUnauthorized() *GetPolicyUsingGETUnauthorized {
	return &GetPolicyUsingGETUnauthorized{}
}

/*
GetPolicyUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPolicyUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get policy using g e t unauthorized response has a 2xx status code
func (o *GetPolicyUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy using g e t unauthorized response has a 3xx status code
func (o *GetPolicyUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy using g e t unauthorized response has a 4xx status code
func (o *GetPolicyUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy using g e t unauthorized response has a 5xx status code
func (o *GetPolicyUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy using g e t unauthorized response a status code equal to that given
func (o *GetPolicyUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get policy using g e t unauthorized response
func (o *GetPolicyUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetPolicyUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/opa/{policyName}][%d] getPolicyUsingGETUnauthorized", 401)
}

func (o *GetPolicyUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/opa/{policyName}][%d] getPolicyUsingGETUnauthorized", 401)
}

func (o *GetPolicyUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyUsingGETForbidden creates a GetPolicyUsingGETForbidden with default headers values
func NewGetPolicyUsingGETForbidden() *GetPolicyUsingGETForbidden {
	return &GetPolicyUsingGETForbidden{}
}

/*
GetPolicyUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPolicyUsingGETForbidden struct {
}

// IsSuccess returns true when this get policy using g e t forbidden response has a 2xx status code
func (o *GetPolicyUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy using g e t forbidden response has a 3xx status code
func (o *GetPolicyUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy using g e t forbidden response has a 4xx status code
func (o *GetPolicyUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy using g e t forbidden response has a 5xx status code
func (o *GetPolicyUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy using g e t forbidden response a status code equal to that given
func (o *GetPolicyUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get policy using g e t forbidden response
func (o *GetPolicyUsingGETForbidden) Code() int {
	return 403
}

func (o *GetPolicyUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/opa/{policyName}][%d] getPolicyUsingGETForbidden", 403)
}

func (o *GetPolicyUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/opa/{policyName}][%d] getPolicyUsingGETForbidden", 403)
}

func (o *GetPolicyUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyUsingGETNotFound creates a GetPolicyUsingGETNotFound with default headers values
func NewGetPolicyUsingGETNotFound() *GetPolicyUsingGETNotFound {
	return &GetPolicyUsingGETNotFound{}
}

/*
GetPolicyUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPolicyUsingGETNotFound struct {
}

// IsSuccess returns true when this get policy using g e t not found response has a 2xx status code
func (o *GetPolicyUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy using g e t not found response has a 3xx status code
func (o *GetPolicyUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy using g e t not found response has a 4xx status code
func (o *GetPolicyUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy using g e t not found response has a 5xx status code
func (o *GetPolicyUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy using g e t not found response a status code equal to that given
func (o *GetPolicyUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get policy using g e t not found response
func (o *GetPolicyUsingGETNotFound) Code() int {
	return 404
}

func (o *GetPolicyUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/opa/{policyName}][%d] getPolicyUsingGETNotFound", 404)
}

func (o *GetPolicyUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/opa/{policyName}][%d] getPolicyUsingGETNotFound", 404)
}

func (o *GetPolicyUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
