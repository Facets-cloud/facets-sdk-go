// Code generated by go-swagger; DO NOT EDIT.

package ui_aws_cluster_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ValidateVpcIDUsingGETReader is a Reader for the ValidateVpcIDUsingGET structure.
type ValidateVpcIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateVpcIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateVpcIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewValidateVpcIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewValidateVpcIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateVpcIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/aws/clusters/validate-vpcId] validateVpcIdUsingGET", response, response.Code())
	}
}

// NewValidateVpcIDUsingGETOK creates a ValidateVpcIDUsingGETOK with default headers values
func NewValidateVpcIDUsingGETOK() *ValidateVpcIDUsingGETOK {
	return &ValidateVpcIDUsingGETOK{}
}

/*
ValidateVpcIDUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type ValidateVpcIDUsingGETOK struct {
	Payload bool
}

// IsSuccess returns true when this validate vpc Id using g e t o k response has a 2xx status code
func (o *ValidateVpcIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate vpc Id using g e t o k response has a 3xx status code
func (o *ValidateVpcIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vpc Id using g e t o k response has a 4xx status code
func (o *ValidateVpcIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate vpc Id using g e t o k response has a 5xx status code
func (o *ValidateVpcIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vpc Id using g e t o k response a status code equal to that given
func (o *ValidateVpcIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the validate vpc Id using g e t o k response
func (o *ValidateVpcIDUsingGETOK) Code() int {
	return 200
}

func (o *ValidateVpcIDUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/aws/clusters/validate-vpcId][%d] validateVpcIdUsingGETOK %s", 200, payload)
}

func (o *ValidateVpcIDUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/aws/clusters/validate-vpcId][%d] validateVpcIdUsingGETOK %s", 200, payload)
}

func (o *ValidateVpcIDUsingGETOK) GetPayload() bool {
	return o.Payload
}

func (o *ValidateVpcIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateVpcIDUsingGETUnauthorized creates a ValidateVpcIDUsingGETUnauthorized with default headers values
func NewValidateVpcIDUsingGETUnauthorized() *ValidateVpcIDUsingGETUnauthorized {
	return &ValidateVpcIDUsingGETUnauthorized{}
}

/*
ValidateVpcIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ValidateVpcIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this validate vpc Id using g e t unauthorized response has a 2xx status code
func (o *ValidateVpcIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate vpc Id using g e t unauthorized response has a 3xx status code
func (o *ValidateVpcIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vpc Id using g e t unauthorized response has a 4xx status code
func (o *ValidateVpcIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate vpc Id using g e t unauthorized response has a 5xx status code
func (o *ValidateVpcIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vpc Id using g e t unauthorized response a status code equal to that given
func (o *ValidateVpcIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the validate vpc Id using g e t unauthorized response
func (o *ValidateVpcIDUsingGETUnauthorized) Code() int {
	return 401
}

func (o *ValidateVpcIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/aws/clusters/validate-vpcId][%d] validateVpcIdUsingGETUnauthorized", 401)
}

func (o *ValidateVpcIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/aws/clusters/validate-vpcId][%d] validateVpcIdUsingGETUnauthorized", 401)
}

func (o *ValidateVpcIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateVpcIDUsingGETForbidden creates a ValidateVpcIDUsingGETForbidden with default headers values
func NewValidateVpcIDUsingGETForbidden() *ValidateVpcIDUsingGETForbidden {
	return &ValidateVpcIDUsingGETForbidden{}
}

/*
ValidateVpcIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ValidateVpcIDUsingGETForbidden struct {
}

// IsSuccess returns true when this validate vpc Id using g e t forbidden response has a 2xx status code
func (o *ValidateVpcIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate vpc Id using g e t forbidden response has a 3xx status code
func (o *ValidateVpcIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vpc Id using g e t forbidden response has a 4xx status code
func (o *ValidateVpcIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate vpc Id using g e t forbidden response has a 5xx status code
func (o *ValidateVpcIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vpc Id using g e t forbidden response a status code equal to that given
func (o *ValidateVpcIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the validate vpc Id using g e t forbidden response
func (o *ValidateVpcIDUsingGETForbidden) Code() int {
	return 403
}

func (o *ValidateVpcIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/aws/clusters/validate-vpcId][%d] validateVpcIdUsingGETForbidden", 403)
}

func (o *ValidateVpcIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/aws/clusters/validate-vpcId][%d] validateVpcIdUsingGETForbidden", 403)
}

func (o *ValidateVpcIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateVpcIDUsingGETNotFound creates a ValidateVpcIDUsingGETNotFound with default headers values
func NewValidateVpcIDUsingGETNotFound() *ValidateVpcIDUsingGETNotFound {
	return &ValidateVpcIDUsingGETNotFound{}
}

/*
ValidateVpcIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ValidateVpcIDUsingGETNotFound struct {
}

// IsSuccess returns true when this validate vpc Id using g e t not found response has a 2xx status code
func (o *ValidateVpcIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate vpc Id using g e t not found response has a 3xx status code
func (o *ValidateVpcIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate vpc Id using g e t not found response has a 4xx status code
func (o *ValidateVpcIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate vpc Id using g e t not found response has a 5xx status code
func (o *ValidateVpcIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this validate vpc Id using g e t not found response a status code equal to that given
func (o *ValidateVpcIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the validate vpc Id using g e t not found response
func (o *ValidateVpcIDUsingGETNotFound) Code() int {
	return 404
}

func (o *ValidateVpcIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/aws/clusters/validate-vpcId][%d] validateVpcIdUsingGETNotFound", 404)
}

func (o *ValidateVpcIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/aws/clusters/validate-vpcId][%d] validateVpcIdUsingGETNotFound", 404)
}

func (o *ValidateVpcIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
