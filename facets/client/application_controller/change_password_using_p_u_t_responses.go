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

	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

// ChangePasswordUsingPUTReader is a Reader for the ChangePasswordUsingPUT structure.
type ChangePasswordUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePasswordUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangePasswordUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewChangePasswordUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewChangePasswordUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewChangePasswordUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChangePasswordUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/users/{userId}/changePassword] changePasswordUsingPUT", response, response.Code())
	}
}

// NewChangePasswordUsingPUTOK creates a ChangePasswordUsingPUTOK with default headers values
func NewChangePasswordUsingPUTOK() *ChangePasswordUsingPUTOK {
	return &ChangePasswordUsingPUTOK{}
}

/*
ChangePasswordUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type ChangePasswordUsingPUTOK struct {
	Payload *models.User
}

// IsSuccess returns true when this change password using p u t o k response has a 2xx status code
func (o *ChangePasswordUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change password using p u t o k response has a 3xx status code
func (o *ChangePasswordUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change password using p u t o k response has a 4xx status code
func (o *ChangePasswordUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this change password using p u t o k response has a 5xx status code
func (o *ChangePasswordUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this change password using p u t o k response a status code equal to that given
func (o *ChangePasswordUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the change password using p u t o k response
func (o *ChangePasswordUsingPUTOK) Code() int {
	return 200
}

func (o *ChangePasswordUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTOK %s", 200, payload)
}

func (o *ChangePasswordUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTOK %s", 200, payload)
}

func (o *ChangePasswordUsingPUTOK) GetPayload() *models.User {
	return o.Payload
}

func (o *ChangePasswordUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangePasswordUsingPUTCreated creates a ChangePasswordUsingPUTCreated with default headers values
func NewChangePasswordUsingPUTCreated() *ChangePasswordUsingPUTCreated {
	return &ChangePasswordUsingPUTCreated{}
}

/*
ChangePasswordUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type ChangePasswordUsingPUTCreated struct {
}

// IsSuccess returns true when this change password using p u t created response has a 2xx status code
func (o *ChangePasswordUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change password using p u t created response has a 3xx status code
func (o *ChangePasswordUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change password using p u t created response has a 4xx status code
func (o *ChangePasswordUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this change password using p u t created response has a 5xx status code
func (o *ChangePasswordUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this change password using p u t created response a status code equal to that given
func (o *ChangePasswordUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the change password using p u t created response
func (o *ChangePasswordUsingPUTCreated) Code() int {
	return 201
}

func (o *ChangePasswordUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTCreated", 201)
}

func (o *ChangePasswordUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTCreated", 201)
}

func (o *ChangePasswordUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangePasswordUsingPUTUnauthorized creates a ChangePasswordUsingPUTUnauthorized with default headers values
func NewChangePasswordUsingPUTUnauthorized() *ChangePasswordUsingPUTUnauthorized {
	return &ChangePasswordUsingPUTUnauthorized{}
}

/*
ChangePasswordUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ChangePasswordUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this change password using p u t unauthorized response has a 2xx status code
func (o *ChangePasswordUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change password using p u t unauthorized response has a 3xx status code
func (o *ChangePasswordUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change password using p u t unauthorized response has a 4xx status code
func (o *ChangePasswordUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this change password using p u t unauthorized response has a 5xx status code
func (o *ChangePasswordUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this change password using p u t unauthorized response a status code equal to that given
func (o *ChangePasswordUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the change password using p u t unauthorized response
func (o *ChangePasswordUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *ChangePasswordUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTUnauthorized", 401)
}

func (o *ChangePasswordUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTUnauthorized", 401)
}

func (o *ChangePasswordUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangePasswordUsingPUTForbidden creates a ChangePasswordUsingPUTForbidden with default headers values
func NewChangePasswordUsingPUTForbidden() *ChangePasswordUsingPUTForbidden {
	return &ChangePasswordUsingPUTForbidden{}
}

/*
ChangePasswordUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ChangePasswordUsingPUTForbidden struct {
}

// IsSuccess returns true when this change password using p u t forbidden response has a 2xx status code
func (o *ChangePasswordUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change password using p u t forbidden response has a 3xx status code
func (o *ChangePasswordUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change password using p u t forbidden response has a 4xx status code
func (o *ChangePasswordUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this change password using p u t forbidden response has a 5xx status code
func (o *ChangePasswordUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this change password using p u t forbidden response a status code equal to that given
func (o *ChangePasswordUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the change password using p u t forbidden response
func (o *ChangePasswordUsingPUTForbidden) Code() int {
	return 403
}

func (o *ChangePasswordUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTForbidden", 403)
}

func (o *ChangePasswordUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTForbidden", 403)
}

func (o *ChangePasswordUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangePasswordUsingPUTNotFound creates a ChangePasswordUsingPUTNotFound with default headers values
func NewChangePasswordUsingPUTNotFound() *ChangePasswordUsingPUTNotFound {
	return &ChangePasswordUsingPUTNotFound{}
}

/*
ChangePasswordUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ChangePasswordUsingPUTNotFound struct {
}

// IsSuccess returns true when this change password using p u t not found response has a 2xx status code
func (o *ChangePasswordUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change password using p u t not found response has a 3xx status code
func (o *ChangePasswordUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change password using p u t not found response has a 4xx status code
func (o *ChangePasswordUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this change password using p u t not found response has a 5xx status code
func (o *ChangePasswordUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this change password using p u t not found response a status code equal to that given
func (o *ChangePasswordUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the change password using p u t not found response
func (o *ChangePasswordUsingPUTNotFound) Code() int {
	return 404
}

func (o *ChangePasswordUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTNotFound", 404)
}

func (o *ChangePasswordUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /api/users/{userId}/changePassword][%d] changePasswordUsingPUTNotFound", 404)
}

func (o *ChangePasswordUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
