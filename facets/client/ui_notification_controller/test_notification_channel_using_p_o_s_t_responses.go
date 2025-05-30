// Code generated by go-swagger; DO NOT EDIT.

package ui_notification_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TestNotificationChannelUsingPOSTReader is a Reader for the TestNotificationChannelUsingPOST structure.
type TestNotificationChannelUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestNotificationChannelUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestNotificationChannelUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewTestNotificationChannelUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTestNotificationChannelUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTestNotificationChannelUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTestNotificationChannelUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/notification/channels/test] testNotificationChannelUsingPOST", response, response.Code())
	}
}

// NewTestNotificationChannelUsingPOSTOK creates a TestNotificationChannelUsingPOSTOK with default headers values
func NewTestNotificationChannelUsingPOSTOK() *TestNotificationChannelUsingPOSTOK {
	return &TestNotificationChannelUsingPOSTOK{}
}

/*
TestNotificationChannelUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type TestNotificationChannelUsingPOSTOK struct {
	Payload bool
}

// IsSuccess returns true when this test notification channel using p o s t o k response has a 2xx status code
func (o *TestNotificationChannelUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this test notification channel using p o s t o k response has a 3xx status code
func (o *TestNotificationChannelUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test notification channel using p o s t o k response has a 4xx status code
func (o *TestNotificationChannelUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this test notification channel using p o s t o k response has a 5xx status code
func (o *TestNotificationChannelUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this test notification channel using p o s t o k response a status code equal to that given
func (o *TestNotificationChannelUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the test notification channel using p o s t o k response
func (o *TestNotificationChannelUsingPOSTOK) Code() int {
	return 200
}

func (o *TestNotificationChannelUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTOK %s", 200, payload)
}

func (o *TestNotificationChannelUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTOK %s", 200, payload)
}

func (o *TestNotificationChannelUsingPOSTOK) GetPayload() bool {
	return o.Payload
}

func (o *TestNotificationChannelUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestNotificationChannelUsingPOSTCreated creates a TestNotificationChannelUsingPOSTCreated with default headers values
func NewTestNotificationChannelUsingPOSTCreated() *TestNotificationChannelUsingPOSTCreated {
	return &TestNotificationChannelUsingPOSTCreated{}
}

/*
TestNotificationChannelUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type TestNotificationChannelUsingPOSTCreated struct {
}

// IsSuccess returns true when this test notification channel using p o s t created response has a 2xx status code
func (o *TestNotificationChannelUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this test notification channel using p o s t created response has a 3xx status code
func (o *TestNotificationChannelUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test notification channel using p o s t created response has a 4xx status code
func (o *TestNotificationChannelUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this test notification channel using p o s t created response has a 5xx status code
func (o *TestNotificationChannelUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this test notification channel using p o s t created response a status code equal to that given
func (o *TestNotificationChannelUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the test notification channel using p o s t created response
func (o *TestNotificationChannelUsingPOSTCreated) Code() int {
	return 201
}

func (o *TestNotificationChannelUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTCreated", 201)
}

func (o *TestNotificationChannelUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTCreated", 201)
}

func (o *TestNotificationChannelUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestNotificationChannelUsingPOSTUnauthorized creates a TestNotificationChannelUsingPOSTUnauthorized with default headers values
func NewTestNotificationChannelUsingPOSTUnauthorized() *TestNotificationChannelUsingPOSTUnauthorized {
	return &TestNotificationChannelUsingPOSTUnauthorized{}
}

/*
TestNotificationChannelUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TestNotificationChannelUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this test notification channel using p o s t unauthorized response has a 2xx status code
func (o *TestNotificationChannelUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test notification channel using p o s t unauthorized response has a 3xx status code
func (o *TestNotificationChannelUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test notification channel using p o s t unauthorized response has a 4xx status code
func (o *TestNotificationChannelUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this test notification channel using p o s t unauthorized response has a 5xx status code
func (o *TestNotificationChannelUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this test notification channel using p o s t unauthorized response a status code equal to that given
func (o *TestNotificationChannelUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the test notification channel using p o s t unauthorized response
func (o *TestNotificationChannelUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *TestNotificationChannelUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTUnauthorized", 401)
}

func (o *TestNotificationChannelUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTUnauthorized", 401)
}

func (o *TestNotificationChannelUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestNotificationChannelUsingPOSTForbidden creates a TestNotificationChannelUsingPOSTForbidden with default headers values
func NewTestNotificationChannelUsingPOSTForbidden() *TestNotificationChannelUsingPOSTForbidden {
	return &TestNotificationChannelUsingPOSTForbidden{}
}

/*
TestNotificationChannelUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TestNotificationChannelUsingPOSTForbidden struct {
}

// IsSuccess returns true when this test notification channel using p o s t forbidden response has a 2xx status code
func (o *TestNotificationChannelUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test notification channel using p o s t forbidden response has a 3xx status code
func (o *TestNotificationChannelUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test notification channel using p o s t forbidden response has a 4xx status code
func (o *TestNotificationChannelUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this test notification channel using p o s t forbidden response has a 5xx status code
func (o *TestNotificationChannelUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this test notification channel using p o s t forbidden response a status code equal to that given
func (o *TestNotificationChannelUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the test notification channel using p o s t forbidden response
func (o *TestNotificationChannelUsingPOSTForbidden) Code() int {
	return 403
}

func (o *TestNotificationChannelUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTForbidden", 403)
}

func (o *TestNotificationChannelUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTForbidden", 403)
}

func (o *TestNotificationChannelUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTestNotificationChannelUsingPOSTNotFound creates a TestNotificationChannelUsingPOSTNotFound with default headers values
func NewTestNotificationChannelUsingPOSTNotFound() *TestNotificationChannelUsingPOSTNotFound {
	return &TestNotificationChannelUsingPOSTNotFound{}
}

/*
TestNotificationChannelUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TestNotificationChannelUsingPOSTNotFound struct {
}

// IsSuccess returns true when this test notification channel using p o s t not found response has a 2xx status code
func (o *TestNotificationChannelUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this test notification channel using p o s t not found response has a 3xx status code
func (o *TestNotificationChannelUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this test notification channel using p o s t not found response has a 4xx status code
func (o *TestNotificationChannelUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this test notification channel using p o s t not found response has a 5xx status code
func (o *TestNotificationChannelUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this test notification channel using p o s t not found response a status code equal to that given
func (o *TestNotificationChannelUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the test notification channel using p o s t not found response
func (o *TestNotificationChannelUsingPOSTNotFound) Code() int {
	return 404
}

func (o *TestNotificationChannelUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTNotFound", 404)
}

func (o *TestNotificationChannelUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels/test][%d] testNotificationChannelUsingPOSTNotFound", 404)
}

func (o *TestNotificationChannelUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
