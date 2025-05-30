// Code generated by go-swagger; DO NOT EDIT.

package ui_artifactory_controller

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

// UpdateECRArtifactoryUsingPOSTReader is a Reader for the UpdateECRArtifactoryUsingPOST structure.
type UpdateECRArtifactoryUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateECRArtifactoryUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateECRArtifactoryUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateECRArtifactoryUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateECRArtifactoryUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateECRArtifactoryUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateECRArtifactoryUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/artifactories/{artifactoryId}] updateECRArtifactoryUsingPOST", response, response.Code())
	}
}

// NewUpdateECRArtifactoryUsingPOSTOK creates a UpdateECRArtifactoryUsingPOSTOK with default headers values
func NewUpdateECRArtifactoryUsingPOSTOK() *UpdateECRArtifactoryUsingPOSTOK {
	return &UpdateECRArtifactoryUsingPOSTOK{}
}

/*
UpdateECRArtifactoryUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type UpdateECRArtifactoryUsingPOSTOK struct {
	Payload *models.ECRArtifactory
}

// IsSuccess returns true when this update e c r artifactory using p o s t o k response has a 2xx status code
func (o *UpdateECRArtifactoryUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update e c r artifactory using p o s t o k response has a 3xx status code
func (o *UpdateECRArtifactoryUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update e c r artifactory using p o s t o k response has a 4xx status code
func (o *UpdateECRArtifactoryUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update e c r artifactory using p o s t o k response has a 5xx status code
func (o *UpdateECRArtifactoryUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update e c r artifactory using p o s t o k response a status code equal to that given
func (o *UpdateECRArtifactoryUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update e c r artifactory using p o s t o k response
func (o *UpdateECRArtifactoryUsingPOSTOK) Code() int {
	return 200
}

func (o *UpdateECRArtifactoryUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTOK %s", 200, payload)
}

func (o *UpdateECRArtifactoryUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTOK %s", 200, payload)
}

func (o *UpdateECRArtifactoryUsingPOSTOK) GetPayload() *models.ECRArtifactory {
	return o.Payload
}

func (o *UpdateECRArtifactoryUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ECRArtifactory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateECRArtifactoryUsingPOSTCreated creates a UpdateECRArtifactoryUsingPOSTCreated with default headers values
func NewUpdateECRArtifactoryUsingPOSTCreated() *UpdateECRArtifactoryUsingPOSTCreated {
	return &UpdateECRArtifactoryUsingPOSTCreated{}
}

/*
UpdateECRArtifactoryUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateECRArtifactoryUsingPOSTCreated struct {
}

// IsSuccess returns true when this update e c r artifactory using p o s t created response has a 2xx status code
func (o *UpdateECRArtifactoryUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update e c r artifactory using p o s t created response has a 3xx status code
func (o *UpdateECRArtifactoryUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update e c r artifactory using p o s t created response has a 4xx status code
func (o *UpdateECRArtifactoryUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update e c r artifactory using p o s t created response has a 5xx status code
func (o *UpdateECRArtifactoryUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update e c r artifactory using p o s t created response a status code equal to that given
func (o *UpdateECRArtifactoryUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update e c r artifactory using p o s t created response
func (o *UpdateECRArtifactoryUsingPOSTCreated) Code() int {
	return 201
}

func (o *UpdateECRArtifactoryUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTCreated", 201)
}

func (o *UpdateECRArtifactoryUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTCreated", 201)
}

func (o *UpdateECRArtifactoryUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateECRArtifactoryUsingPOSTUnauthorized creates a UpdateECRArtifactoryUsingPOSTUnauthorized with default headers values
func NewUpdateECRArtifactoryUsingPOSTUnauthorized() *UpdateECRArtifactoryUsingPOSTUnauthorized {
	return &UpdateECRArtifactoryUsingPOSTUnauthorized{}
}

/*
UpdateECRArtifactoryUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateECRArtifactoryUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this update e c r artifactory using p o s t unauthorized response has a 2xx status code
func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update e c r artifactory using p o s t unauthorized response has a 3xx status code
func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update e c r artifactory using p o s t unauthorized response has a 4xx status code
func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update e c r artifactory using p o s t unauthorized response has a 5xx status code
func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update e c r artifactory using p o s t unauthorized response a status code equal to that given
func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update e c r artifactory using p o s t unauthorized response
func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTUnauthorized", 401)
}

func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTUnauthorized", 401)
}

func (o *UpdateECRArtifactoryUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateECRArtifactoryUsingPOSTForbidden creates a UpdateECRArtifactoryUsingPOSTForbidden with default headers values
func NewUpdateECRArtifactoryUsingPOSTForbidden() *UpdateECRArtifactoryUsingPOSTForbidden {
	return &UpdateECRArtifactoryUsingPOSTForbidden{}
}

/*
UpdateECRArtifactoryUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateECRArtifactoryUsingPOSTForbidden struct {
}

// IsSuccess returns true when this update e c r artifactory using p o s t forbidden response has a 2xx status code
func (o *UpdateECRArtifactoryUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update e c r artifactory using p o s t forbidden response has a 3xx status code
func (o *UpdateECRArtifactoryUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update e c r artifactory using p o s t forbidden response has a 4xx status code
func (o *UpdateECRArtifactoryUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update e c r artifactory using p o s t forbidden response has a 5xx status code
func (o *UpdateECRArtifactoryUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update e c r artifactory using p o s t forbidden response a status code equal to that given
func (o *UpdateECRArtifactoryUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update e c r artifactory using p o s t forbidden response
func (o *UpdateECRArtifactoryUsingPOSTForbidden) Code() int {
	return 403
}

func (o *UpdateECRArtifactoryUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTForbidden", 403)
}

func (o *UpdateECRArtifactoryUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTForbidden", 403)
}

func (o *UpdateECRArtifactoryUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateECRArtifactoryUsingPOSTNotFound creates a UpdateECRArtifactoryUsingPOSTNotFound with default headers values
func NewUpdateECRArtifactoryUsingPOSTNotFound() *UpdateECRArtifactoryUsingPOSTNotFound {
	return &UpdateECRArtifactoryUsingPOSTNotFound{}
}

/*
UpdateECRArtifactoryUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateECRArtifactoryUsingPOSTNotFound struct {
}

// IsSuccess returns true when this update e c r artifactory using p o s t not found response has a 2xx status code
func (o *UpdateECRArtifactoryUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update e c r artifactory using p o s t not found response has a 3xx status code
func (o *UpdateECRArtifactoryUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update e c r artifactory using p o s t not found response has a 4xx status code
func (o *UpdateECRArtifactoryUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update e c r artifactory using p o s t not found response has a 5xx status code
func (o *UpdateECRArtifactoryUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update e c r artifactory using p o s t not found response a status code equal to that given
func (o *UpdateECRArtifactoryUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update e c r artifactory using p o s t not found response
func (o *UpdateECRArtifactoryUsingPOSTNotFound) Code() int {
	return 404
}

func (o *UpdateECRArtifactoryUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTNotFound", 404)
}

func (o *UpdateECRArtifactoryUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/artifactories/{artifactoryId}][%d] updateECRArtifactoryUsingPOSTNotFound", 404)
}

func (o *UpdateECRArtifactoryUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
