// Code generated by go-swagger; DO NOT EDIT.

package ui_custom_role_controller

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

// GetCustomRoleUsingGETReader is a Reader for the GetCustomRoleUsingGET structure.
type GetCustomRoleUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomRoleUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomRoleUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCustomRoleUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCustomRoleUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCustomRoleUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/custom-role/{roleName}] getCustomRoleUsingGET", response, response.Code())
	}
}

// NewGetCustomRoleUsingGETOK creates a GetCustomRoleUsingGETOK with default headers values
func NewGetCustomRoleUsingGETOK() *GetCustomRoleUsingGETOK {
	return &GetCustomRoleUsingGETOK{}
}

/*
GetCustomRoleUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetCustomRoleUsingGETOK struct {
	Payload *models.RoleMapping
}

// IsSuccess returns true when this get custom role using g e t o k response has a 2xx status code
func (o *GetCustomRoleUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get custom role using g e t o k response has a 3xx status code
func (o *GetCustomRoleUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom role using g e t o k response has a 4xx status code
func (o *GetCustomRoleUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get custom role using g e t o k response has a 5xx status code
func (o *GetCustomRoleUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom role using g e t o k response a status code equal to that given
func (o *GetCustomRoleUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get custom role using g e t o k response
func (o *GetCustomRoleUsingGETOK) Code() int {
	return 200
}

func (o *GetCustomRoleUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/custom-role/{roleName}][%d] getCustomRoleUsingGETOK %s", 200, payload)
}

func (o *GetCustomRoleUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/custom-role/{roleName}][%d] getCustomRoleUsingGETOK %s", 200, payload)
}

func (o *GetCustomRoleUsingGETOK) GetPayload() *models.RoleMapping {
	return o.Payload
}

func (o *GetCustomRoleUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomRoleUsingGETUnauthorized creates a GetCustomRoleUsingGETUnauthorized with default headers values
func NewGetCustomRoleUsingGETUnauthorized() *GetCustomRoleUsingGETUnauthorized {
	return &GetCustomRoleUsingGETUnauthorized{}
}

/*
GetCustomRoleUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCustomRoleUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get custom role using g e t unauthorized response has a 2xx status code
func (o *GetCustomRoleUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom role using g e t unauthorized response has a 3xx status code
func (o *GetCustomRoleUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom role using g e t unauthorized response has a 4xx status code
func (o *GetCustomRoleUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom role using g e t unauthorized response has a 5xx status code
func (o *GetCustomRoleUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom role using g e t unauthorized response a status code equal to that given
func (o *GetCustomRoleUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get custom role using g e t unauthorized response
func (o *GetCustomRoleUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetCustomRoleUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/custom-role/{roleName}][%d] getCustomRoleUsingGETUnauthorized", 401)
}

func (o *GetCustomRoleUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/custom-role/{roleName}][%d] getCustomRoleUsingGETUnauthorized", 401)
}

func (o *GetCustomRoleUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomRoleUsingGETForbidden creates a GetCustomRoleUsingGETForbidden with default headers values
func NewGetCustomRoleUsingGETForbidden() *GetCustomRoleUsingGETForbidden {
	return &GetCustomRoleUsingGETForbidden{}
}

/*
GetCustomRoleUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCustomRoleUsingGETForbidden struct {
}

// IsSuccess returns true when this get custom role using g e t forbidden response has a 2xx status code
func (o *GetCustomRoleUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom role using g e t forbidden response has a 3xx status code
func (o *GetCustomRoleUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom role using g e t forbidden response has a 4xx status code
func (o *GetCustomRoleUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom role using g e t forbidden response has a 5xx status code
func (o *GetCustomRoleUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom role using g e t forbidden response a status code equal to that given
func (o *GetCustomRoleUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get custom role using g e t forbidden response
func (o *GetCustomRoleUsingGETForbidden) Code() int {
	return 403
}

func (o *GetCustomRoleUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/custom-role/{roleName}][%d] getCustomRoleUsingGETForbidden", 403)
}

func (o *GetCustomRoleUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/custom-role/{roleName}][%d] getCustomRoleUsingGETForbidden", 403)
}

func (o *GetCustomRoleUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomRoleUsingGETNotFound creates a GetCustomRoleUsingGETNotFound with default headers values
func NewGetCustomRoleUsingGETNotFound() *GetCustomRoleUsingGETNotFound {
	return &GetCustomRoleUsingGETNotFound{}
}

/*
GetCustomRoleUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCustomRoleUsingGETNotFound struct {
}

// IsSuccess returns true when this get custom role using g e t not found response has a 2xx status code
func (o *GetCustomRoleUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom role using g e t not found response has a 3xx status code
func (o *GetCustomRoleUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom role using g e t not found response has a 4xx status code
func (o *GetCustomRoleUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom role using g e t not found response has a 5xx status code
func (o *GetCustomRoleUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom role using g e t not found response a status code equal to that given
func (o *GetCustomRoleUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get custom role using g e t not found response
func (o *GetCustomRoleUsingGETNotFound) Code() int {
	return 404
}

func (o *GetCustomRoleUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/custom-role/{roleName}][%d] getCustomRoleUsingGETNotFound", 404)
}

func (o *GetCustomRoleUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/custom-role/{roleName}][%d] getCustomRoleUsingGETNotFound", 404)
}

func (o *GetCustomRoleUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
