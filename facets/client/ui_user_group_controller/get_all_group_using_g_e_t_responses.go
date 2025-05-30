// Code generated by go-swagger; DO NOT EDIT.

package ui_user_group_controller

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

// GetAllGroupUsingGETReader is a Reader for the GetAllGroupUsingGET structure.
type GetAllGroupUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllGroupUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllGroupUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllGroupUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllGroupUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllGroupUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/user-groups/] getAllGroupUsingGET", response, response.Code())
	}
}

// NewGetAllGroupUsingGETOK creates a GetAllGroupUsingGETOK with default headers values
func NewGetAllGroupUsingGETOK() *GetAllGroupUsingGETOK {
	return &GetAllGroupUsingGETOK{}
}

/*
GetAllGroupUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAllGroupUsingGETOK struct {
	Payload []*models.UserGroup
}

// IsSuccess returns true when this get all group using g e t o k response has a 2xx status code
func (o *GetAllGroupUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all group using g e t o k response has a 3xx status code
func (o *GetAllGroupUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all group using g e t o k response has a 4xx status code
func (o *GetAllGroupUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all group using g e t o k response has a 5xx status code
func (o *GetAllGroupUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all group using g e t o k response a status code equal to that given
func (o *GetAllGroupUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all group using g e t o k response
func (o *GetAllGroupUsingGETOK) Code() int {
	return 200
}

func (o *GetAllGroupUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/user-groups/][%d] getAllGroupUsingGETOK %s", 200, payload)
}

func (o *GetAllGroupUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/user-groups/][%d] getAllGroupUsingGETOK %s", 200, payload)
}

func (o *GetAllGroupUsingGETOK) GetPayload() []*models.UserGroup {
	return o.Payload
}

func (o *GetAllGroupUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllGroupUsingGETUnauthorized creates a GetAllGroupUsingGETUnauthorized with default headers values
func NewGetAllGroupUsingGETUnauthorized() *GetAllGroupUsingGETUnauthorized {
	return &GetAllGroupUsingGETUnauthorized{}
}

/*
GetAllGroupUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllGroupUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get all group using g e t unauthorized response has a 2xx status code
func (o *GetAllGroupUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all group using g e t unauthorized response has a 3xx status code
func (o *GetAllGroupUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all group using g e t unauthorized response has a 4xx status code
func (o *GetAllGroupUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all group using g e t unauthorized response has a 5xx status code
func (o *GetAllGroupUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all group using g e t unauthorized response a status code equal to that given
func (o *GetAllGroupUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all group using g e t unauthorized response
func (o *GetAllGroupUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetAllGroupUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/user-groups/][%d] getAllGroupUsingGETUnauthorized", 401)
}

func (o *GetAllGroupUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/user-groups/][%d] getAllGroupUsingGETUnauthorized", 401)
}

func (o *GetAllGroupUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllGroupUsingGETForbidden creates a GetAllGroupUsingGETForbidden with default headers values
func NewGetAllGroupUsingGETForbidden() *GetAllGroupUsingGETForbidden {
	return &GetAllGroupUsingGETForbidden{}
}

/*
GetAllGroupUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllGroupUsingGETForbidden struct {
}

// IsSuccess returns true when this get all group using g e t forbidden response has a 2xx status code
func (o *GetAllGroupUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all group using g e t forbidden response has a 3xx status code
func (o *GetAllGroupUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all group using g e t forbidden response has a 4xx status code
func (o *GetAllGroupUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all group using g e t forbidden response has a 5xx status code
func (o *GetAllGroupUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all group using g e t forbidden response a status code equal to that given
func (o *GetAllGroupUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get all group using g e t forbidden response
func (o *GetAllGroupUsingGETForbidden) Code() int {
	return 403
}

func (o *GetAllGroupUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/user-groups/][%d] getAllGroupUsingGETForbidden", 403)
}

func (o *GetAllGroupUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/user-groups/][%d] getAllGroupUsingGETForbidden", 403)
}

func (o *GetAllGroupUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllGroupUsingGETNotFound creates a GetAllGroupUsingGETNotFound with default headers values
func NewGetAllGroupUsingGETNotFound() *GetAllGroupUsingGETNotFound {
	return &GetAllGroupUsingGETNotFound{}
}

/*
GetAllGroupUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllGroupUsingGETNotFound struct {
}

// IsSuccess returns true when this get all group using g e t not found response has a 2xx status code
func (o *GetAllGroupUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all group using g e t not found response has a 3xx status code
func (o *GetAllGroupUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all group using g e t not found response has a 4xx status code
func (o *GetAllGroupUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all group using g e t not found response has a 5xx status code
func (o *GetAllGroupUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all group using g e t not found response a status code equal to that given
func (o *GetAllGroupUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all group using g e t not found response
func (o *GetAllGroupUsingGETNotFound) Code() int {
	return 404
}

func (o *GetAllGroupUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/user-groups/][%d] getAllGroupUsingGETNotFound", 404)
}

func (o *GetAllGroupUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/user-groups/][%d] getAllGroupUsingGETNotFound", 404)
}

func (o *GetAllGroupUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
