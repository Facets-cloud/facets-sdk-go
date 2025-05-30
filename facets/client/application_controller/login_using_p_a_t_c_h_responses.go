// Code generated by go-swagger; DO NOT EDIT.

package application_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LoginUsingPATCHReader is a Reader for the LoginUsingPATCH structure.
type LoginUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewLoginUsingPATCHNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLoginUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoginUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /api/login] loginUsingPATCH", response, response.Code())
	}
}

// NewLoginUsingPATCHOK creates a LoginUsingPATCHOK with default headers values
func NewLoginUsingPATCHOK() *LoginUsingPATCHOK {
	return &LoginUsingPATCHOK{}
}

/*
LoginUsingPATCHOK describes a response with status code 200, with default header values.

OK
*/
type LoginUsingPATCHOK struct {
	Payload string
}

// IsSuccess returns true when this login using p a t c h o k response has a 2xx status code
func (o *LoginUsingPATCHOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this login using p a t c h o k response has a 3xx status code
func (o *LoginUsingPATCHOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login using p a t c h o k response has a 4xx status code
func (o *LoginUsingPATCHOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this login using p a t c h o k response has a 5xx status code
func (o *LoginUsingPATCHOK) IsServerError() bool {
	return false
}

// IsCode returns true when this login using p a t c h o k response a status code equal to that given
func (o *LoginUsingPATCHOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the login using p a t c h o k response
func (o *LoginUsingPATCHOK) Code() int {
	return 200
}

func (o *LoginUsingPATCHOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/login][%d] loginUsingPATCHOK %s", 200, payload)
}

func (o *LoginUsingPATCHOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/login][%d] loginUsingPATCHOK %s", 200, payload)
}

func (o *LoginUsingPATCHOK) GetPayload() string {
	return o.Payload
}

func (o *LoginUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginUsingPATCHNoContent creates a LoginUsingPATCHNoContent with default headers values
func NewLoginUsingPATCHNoContent() *LoginUsingPATCHNoContent {
	return &LoginUsingPATCHNoContent{}
}

/*
LoginUsingPATCHNoContent describes a response with status code 204, with default header values.

No Content
*/
type LoginUsingPATCHNoContent struct {
}

// IsSuccess returns true when this login using p a t c h no content response has a 2xx status code
func (o *LoginUsingPATCHNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this login using p a t c h no content response has a 3xx status code
func (o *LoginUsingPATCHNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login using p a t c h no content response has a 4xx status code
func (o *LoginUsingPATCHNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this login using p a t c h no content response has a 5xx status code
func (o *LoginUsingPATCHNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this login using p a t c h no content response a status code equal to that given
func (o *LoginUsingPATCHNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the login using p a t c h no content response
func (o *LoginUsingPATCHNoContent) Code() int {
	return 204
}

func (o *LoginUsingPATCHNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/login][%d] loginUsingPATCHNoContent", 204)
}

func (o *LoginUsingPATCHNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/login][%d] loginUsingPATCHNoContent", 204)
}

func (o *LoginUsingPATCHNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginUsingPATCHUnauthorized creates a LoginUsingPATCHUnauthorized with default headers values
func NewLoginUsingPATCHUnauthorized() *LoginUsingPATCHUnauthorized {
	return &LoginUsingPATCHUnauthorized{}
}

/*
LoginUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type LoginUsingPATCHUnauthorized struct {
}

// IsSuccess returns true when this login using p a t c h unauthorized response has a 2xx status code
func (o *LoginUsingPATCHUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this login using p a t c h unauthorized response has a 3xx status code
func (o *LoginUsingPATCHUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login using p a t c h unauthorized response has a 4xx status code
func (o *LoginUsingPATCHUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this login using p a t c h unauthorized response has a 5xx status code
func (o *LoginUsingPATCHUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this login using p a t c h unauthorized response a status code equal to that given
func (o *LoginUsingPATCHUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the login using p a t c h unauthorized response
func (o *LoginUsingPATCHUnauthorized) Code() int {
	return 401
}

func (o *LoginUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/login][%d] loginUsingPATCHUnauthorized", 401)
}

func (o *LoginUsingPATCHUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/login][%d] loginUsingPATCHUnauthorized", 401)
}

func (o *LoginUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginUsingPATCHForbidden creates a LoginUsingPATCHForbidden with default headers values
func NewLoginUsingPATCHForbidden() *LoginUsingPATCHForbidden {
	return &LoginUsingPATCHForbidden{}
}

/*
LoginUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type LoginUsingPATCHForbidden struct {
}

// IsSuccess returns true when this login using p a t c h forbidden response has a 2xx status code
func (o *LoginUsingPATCHForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this login using p a t c h forbidden response has a 3xx status code
func (o *LoginUsingPATCHForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login using p a t c h forbidden response has a 4xx status code
func (o *LoginUsingPATCHForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this login using p a t c h forbidden response has a 5xx status code
func (o *LoginUsingPATCHForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this login using p a t c h forbidden response a status code equal to that given
func (o *LoginUsingPATCHForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the login using p a t c h forbidden response
func (o *LoginUsingPATCHForbidden) Code() int {
	return 403
}

func (o *LoginUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/login][%d] loginUsingPATCHForbidden", 403)
}

func (o *LoginUsingPATCHForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/login][%d] loginUsingPATCHForbidden", 403)
}

func (o *LoginUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
