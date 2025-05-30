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

	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

// CreateNotificationChannelUsingPOSTReader is a Reader for the CreateNotificationChannelUsingPOST structure.
type CreateNotificationChannelUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNotificationChannelUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNotificationChannelUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateNotificationChannelUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateNotificationChannelUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateNotificationChannelUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateNotificationChannelUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/notification/channels] createNotificationChannelUsingPOST", response, response.Code())
	}
}

// NewCreateNotificationChannelUsingPOSTOK creates a CreateNotificationChannelUsingPOSTOK with default headers values
func NewCreateNotificationChannelUsingPOSTOK() *CreateNotificationChannelUsingPOSTOK {
	return &CreateNotificationChannelUsingPOSTOK{}
}

/*
CreateNotificationChannelUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type CreateNotificationChannelUsingPOSTOK struct {
	Payload []*models.NotificationChannel
}

// IsSuccess returns true when this create notification channel using p o s t o k response has a 2xx status code
func (o *CreateNotificationChannelUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create notification channel using p o s t o k response has a 3xx status code
func (o *CreateNotificationChannelUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification channel using p o s t o k response has a 4xx status code
func (o *CreateNotificationChannelUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create notification channel using p o s t o k response has a 5xx status code
func (o *CreateNotificationChannelUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create notification channel using p o s t o k response a status code equal to that given
func (o *CreateNotificationChannelUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create notification channel using p o s t o k response
func (o *CreateNotificationChannelUsingPOSTOK) Code() int {
	return 200
}

func (o *CreateNotificationChannelUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTOK %s", 200, payload)
}

func (o *CreateNotificationChannelUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTOK %s", 200, payload)
}

func (o *CreateNotificationChannelUsingPOSTOK) GetPayload() []*models.NotificationChannel {
	return o.Payload
}

func (o *CreateNotificationChannelUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotificationChannelUsingPOSTCreated creates a CreateNotificationChannelUsingPOSTCreated with default headers values
func NewCreateNotificationChannelUsingPOSTCreated() *CreateNotificationChannelUsingPOSTCreated {
	return &CreateNotificationChannelUsingPOSTCreated{}
}

/*
CreateNotificationChannelUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type CreateNotificationChannelUsingPOSTCreated struct {
}

// IsSuccess returns true when this create notification channel using p o s t created response has a 2xx status code
func (o *CreateNotificationChannelUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create notification channel using p o s t created response has a 3xx status code
func (o *CreateNotificationChannelUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification channel using p o s t created response has a 4xx status code
func (o *CreateNotificationChannelUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create notification channel using p o s t created response has a 5xx status code
func (o *CreateNotificationChannelUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create notification channel using p o s t created response a status code equal to that given
func (o *CreateNotificationChannelUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create notification channel using p o s t created response
func (o *CreateNotificationChannelUsingPOSTCreated) Code() int {
	return 201
}

func (o *CreateNotificationChannelUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTCreated", 201)
}

func (o *CreateNotificationChannelUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTCreated", 201)
}

func (o *CreateNotificationChannelUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNotificationChannelUsingPOSTUnauthorized creates a CreateNotificationChannelUsingPOSTUnauthorized with default headers values
func NewCreateNotificationChannelUsingPOSTUnauthorized() *CreateNotificationChannelUsingPOSTUnauthorized {
	return &CreateNotificationChannelUsingPOSTUnauthorized{}
}

/*
CreateNotificationChannelUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateNotificationChannelUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this create notification channel using p o s t unauthorized response has a 2xx status code
func (o *CreateNotificationChannelUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notification channel using p o s t unauthorized response has a 3xx status code
func (o *CreateNotificationChannelUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification channel using p o s t unauthorized response has a 4xx status code
func (o *CreateNotificationChannelUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create notification channel using p o s t unauthorized response has a 5xx status code
func (o *CreateNotificationChannelUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create notification channel using p o s t unauthorized response a status code equal to that given
func (o *CreateNotificationChannelUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create notification channel using p o s t unauthorized response
func (o *CreateNotificationChannelUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *CreateNotificationChannelUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTUnauthorized", 401)
}

func (o *CreateNotificationChannelUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTUnauthorized", 401)
}

func (o *CreateNotificationChannelUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNotificationChannelUsingPOSTForbidden creates a CreateNotificationChannelUsingPOSTForbidden with default headers values
func NewCreateNotificationChannelUsingPOSTForbidden() *CreateNotificationChannelUsingPOSTForbidden {
	return &CreateNotificationChannelUsingPOSTForbidden{}
}

/*
CreateNotificationChannelUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateNotificationChannelUsingPOSTForbidden struct {
}

// IsSuccess returns true when this create notification channel using p o s t forbidden response has a 2xx status code
func (o *CreateNotificationChannelUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notification channel using p o s t forbidden response has a 3xx status code
func (o *CreateNotificationChannelUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification channel using p o s t forbidden response has a 4xx status code
func (o *CreateNotificationChannelUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create notification channel using p o s t forbidden response has a 5xx status code
func (o *CreateNotificationChannelUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create notification channel using p o s t forbidden response a status code equal to that given
func (o *CreateNotificationChannelUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create notification channel using p o s t forbidden response
func (o *CreateNotificationChannelUsingPOSTForbidden) Code() int {
	return 403
}

func (o *CreateNotificationChannelUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTForbidden", 403)
}

func (o *CreateNotificationChannelUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTForbidden", 403)
}

func (o *CreateNotificationChannelUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNotificationChannelUsingPOSTNotFound creates a CreateNotificationChannelUsingPOSTNotFound with default headers values
func NewCreateNotificationChannelUsingPOSTNotFound() *CreateNotificationChannelUsingPOSTNotFound {
	return &CreateNotificationChannelUsingPOSTNotFound{}
}

/*
CreateNotificationChannelUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateNotificationChannelUsingPOSTNotFound struct {
}

// IsSuccess returns true when this create notification channel using p o s t not found response has a 2xx status code
func (o *CreateNotificationChannelUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notification channel using p o s t not found response has a 3xx status code
func (o *CreateNotificationChannelUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification channel using p o s t not found response has a 4xx status code
func (o *CreateNotificationChannelUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create notification channel using p o s t not found response has a 5xx status code
func (o *CreateNotificationChannelUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create notification channel using p o s t not found response a status code equal to that given
func (o *CreateNotificationChannelUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create notification channel using p o s t not found response
func (o *CreateNotificationChannelUsingPOSTNotFound) Code() int {
	return 404
}

func (o *CreateNotificationChannelUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTNotFound", 404)
}

func (o *CreateNotificationChannelUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/notification/channels][%d] createNotificationChannelUsingPOSTNotFound", 404)
}

func (o *CreateNotificationChannelUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
