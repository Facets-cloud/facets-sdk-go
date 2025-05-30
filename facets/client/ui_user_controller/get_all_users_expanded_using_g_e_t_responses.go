// Code generated by go-swagger; DO NOT EDIT.

package ui_user_controller

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

// GetAllUsersExpandedUsingGETReader is a Reader for the GetAllUsersExpandedUsingGET structure.
type GetAllUsersExpandedUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUsersExpandedUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUsersExpandedUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllUsersExpandedUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllUsersExpandedUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllUsersExpandedUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/users/list/users-expanded] getAllUsersExpandedUsingGET", response, response.Code())
	}
}

// NewGetAllUsersExpandedUsingGETOK creates a GetAllUsersExpandedUsingGETOK with default headers values
func NewGetAllUsersExpandedUsingGETOK() *GetAllUsersExpandedUsingGETOK {
	return &GetAllUsersExpandedUsingGETOK{}
}

/*
GetAllUsersExpandedUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAllUsersExpandedUsingGETOK struct {
	Payload []*models.ExpandedUser
}

// IsSuccess returns true when this get all users expanded using g e t o k response has a 2xx status code
func (o *GetAllUsersExpandedUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all users expanded using g e t o k response has a 3xx status code
func (o *GetAllUsersExpandedUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all users expanded using g e t o k response has a 4xx status code
func (o *GetAllUsersExpandedUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all users expanded using g e t o k response has a 5xx status code
func (o *GetAllUsersExpandedUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all users expanded using g e t o k response a status code equal to that given
func (o *GetAllUsersExpandedUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all users expanded using g e t o k response
func (o *GetAllUsersExpandedUsingGETOK) Code() int {
	return 200
}

func (o *GetAllUsersExpandedUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/users/list/users-expanded][%d] getAllUsersExpandedUsingGETOK %s", 200, payload)
}

func (o *GetAllUsersExpandedUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/users/list/users-expanded][%d] getAllUsersExpandedUsingGETOK %s", 200, payload)
}

func (o *GetAllUsersExpandedUsingGETOK) GetPayload() []*models.ExpandedUser {
	return o.Payload
}

func (o *GetAllUsersExpandedUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUsersExpandedUsingGETUnauthorized creates a GetAllUsersExpandedUsingGETUnauthorized with default headers values
func NewGetAllUsersExpandedUsingGETUnauthorized() *GetAllUsersExpandedUsingGETUnauthorized {
	return &GetAllUsersExpandedUsingGETUnauthorized{}
}

/*
GetAllUsersExpandedUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllUsersExpandedUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get all users expanded using g e t unauthorized response has a 2xx status code
func (o *GetAllUsersExpandedUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all users expanded using g e t unauthorized response has a 3xx status code
func (o *GetAllUsersExpandedUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all users expanded using g e t unauthorized response has a 4xx status code
func (o *GetAllUsersExpandedUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all users expanded using g e t unauthorized response has a 5xx status code
func (o *GetAllUsersExpandedUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all users expanded using g e t unauthorized response a status code equal to that given
func (o *GetAllUsersExpandedUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get all users expanded using g e t unauthorized response
func (o *GetAllUsersExpandedUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetAllUsersExpandedUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/users/list/users-expanded][%d] getAllUsersExpandedUsingGETUnauthorized", 401)
}

func (o *GetAllUsersExpandedUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/users/list/users-expanded][%d] getAllUsersExpandedUsingGETUnauthorized", 401)
}

func (o *GetAllUsersExpandedUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsersExpandedUsingGETForbidden creates a GetAllUsersExpandedUsingGETForbidden with default headers values
func NewGetAllUsersExpandedUsingGETForbidden() *GetAllUsersExpandedUsingGETForbidden {
	return &GetAllUsersExpandedUsingGETForbidden{}
}

/*
GetAllUsersExpandedUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllUsersExpandedUsingGETForbidden struct {
}

// IsSuccess returns true when this get all users expanded using g e t forbidden response has a 2xx status code
func (o *GetAllUsersExpandedUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all users expanded using g e t forbidden response has a 3xx status code
func (o *GetAllUsersExpandedUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all users expanded using g e t forbidden response has a 4xx status code
func (o *GetAllUsersExpandedUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all users expanded using g e t forbidden response has a 5xx status code
func (o *GetAllUsersExpandedUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all users expanded using g e t forbidden response a status code equal to that given
func (o *GetAllUsersExpandedUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get all users expanded using g e t forbidden response
func (o *GetAllUsersExpandedUsingGETForbidden) Code() int {
	return 403
}

func (o *GetAllUsersExpandedUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/users/list/users-expanded][%d] getAllUsersExpandedUsingGETForbidden", 403)
}

func (o *GetAllUsersExpandedUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/users/list/users-expanded][%d] getAllUsersExpandedUsingGETForbidden", 403)
}

func (o *GetAllUsersExpandedUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsersExpandedUsingGETNotFound creates a GetAllUsersExpandedUsingGETNotFound with default headers values
func NewGetAllUsersExpandedUsingGETNotFound() *GetAllUsersExpandedUsingGETNotFound {
	return &GetAllUsersExpandedUsingGETNotFound{}
}

/*
GetAllUsersExpandedUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllUsersExpandedUsingGETNotFound struct {
}

// IsSuccess returns true when this get all users expanded using g e t not found response has a 2xx status code
func (o *GetAllUsersExpandedUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all users expanded using g e t not found response has a 3xx status code
func (o *GetAllUsersExpandedUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all users expanded using g e t not found response has a 4xx status code
func (o *GetAllUsersExpandedUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all users expanded using g e t not found response has a 5xx status code
func (o *GetAllUsersExpandedUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all users expanded using g e t not found response a status code equal to that given
func (o *GetAllUsersExpandedUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get all users expanded using g e t not found response
func (o *GetAllUsersExpandedUsingGETNotFound) Code() int {
	return 404
}

func (o *GetAllUsersExpandedUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/users/list/users-expanded][%d] getAllUsersExpandedUsingGETNotFound", 404)
}

func (o *GetAllUsersExpandedUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/users/list/users-expanded][%d] getAllUsersExpandedUsingGETNotFound", 404)
}

func (o *GetAllUsersExpandedUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
