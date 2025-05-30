// Code generated by go-swagger; DO NOT EDIT.

package deployment_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TriggerAutomationSuiteUsingPOSTReader is a Reader for the TriggerAutomationSuiteUsingPOST structure.
type TriggerAutomationSuiteUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerAutomationSuiteUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriggerAutomationSuiteUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewTriggerAutomationSuiteUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTriggerAutomationSuiteUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTriggerAutomationSuiteUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTriggerAutomationSuiteUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite] triggerAutomationSuiteUsingPOST", response, response.Code())
	}
}

// NewTriggerAutomationSuiteUsingPOSTOK creates a TriggerAutomationSuiteUsingPOSTOK with default headers values
func NewTriggerAutomationSuiteUsingPOSTOK() *TriggerAutomationSuiteUsingPOSTOK {
	return &TriggerAutomationSuiteUsingPOSTOK{}
}

/*
TriggerAutomationSuiteUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type TriggerAutomationSuiteUsingPOSTOK struct {
	Payload string
}

// IsSuccess returns true when this trigger automation suite using p o s t o k response has a 2xx status code
func (o *TriggerAutomationSuiteUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger automation suite using p o s t o k response has a 3xx status code
func (o *TriggerAutomationSuiteUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger automation suite using p o s t o k response has a 4xx status code
func (o *TriggerAutomationSuiteUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger automation suite using p o s t o k response has a 5xx status code
func (o *TriggerAutomationSuiteUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger automation suite using p o s t o k response a status code equal to that given
func (o *TriggerAutomationSuiteUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the trigger automation suite using p o s t o k response
func (o *TriggerAutomationSuiteUsingPOSTOK) Code() int {
	return 200
}

func (o *TriggerAutomationSuiteUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTOK %s", 200, payload)
}

func (o *TriggerAutomationSuiteUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTOK %s", 200, payload)
}

func (o *TriggerAutomationSuiteUsingPOSTOK) GetPayload() string {
	return o.Payload
}

func (o *TriggerAutomationSuiteUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggerAutomationSuiteUsingPOSTCreated creates a TriggerAutomationSuiteUsingPOSTCreated with default headers values
func NewTriggerAutomationSuiteUsingPOSTCreated() *TriggerAutomationSuiteUsingPOSTCreated {
	return &TriggerAutomationSuiteUsingPOSTCreated{}
}

/*
TriggerAutomationSuiteUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type TriggerAutomationSuiteUsingPOSTCreated struct {
}

// IsSuccess returns true when this trigger automation suite using p o s t created response has a 2xx status code
func (o *TriggerAutomationSuiteUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger automation suite using p o s t created response has a 3xx status code
func (o *TriggerAutomationSuiteUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger automation suite using p o s t created response has a 4xx status code
func (o *TriggerAutomationSuiteUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger automation suite using p o s t created response has a 5xx status code
func (o *TriggerAutomationSuiteUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger automation suite using p o s t created response a status code equal to that given
func (o *TriggerAutomationSuiteUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the trigger automation suite using p o s t created response
func (o *TriggerAutomationSuiteUsingPOSTCreated) Code() int {
	return 201
}

func (o *TriggerAutomationSuiteUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTCreated", 201)
}

func (o *TriggerAutomationSuiteUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTCreated", 201)
}

func (o *TriggerAutomationSuiteUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerAutomationSuiteUsingPOSTUnauthorized creates a TriggerAutomationSuiteUsingPOSTUnauthorized with default headers values
func NewTriggerAutomationSuiteUsingPOSTUnauthorized() *TriggerAutomationSuiteUsingPOSTUnauthorized {
	return &TriggerAutomationSuiteUsingPOSTUnauthorized{}
}

/*
TriggerAutomationSuiteUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TriggerAutomationSuiteUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this trigger automation suite using p o s t unauthorized response has a 2xx status code
func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger automation suite using p o s t unauthorized response has a 3xx status code
func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger automation suite using p o s t unauthorized response has a 4xx status code
func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger automation suite using p o s t unauthorized response has a 5xx status code
func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger automation suite using p o s t unauthorized response a status code equal to that given
func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the trigger automation suite using p o s t unauthorized response
func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTUnauthorized", 401)
}

func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTUnauthorized", 401)
}

func (o *TriggerAutomationSuiteUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerAutomationSuiteUsingPOSTForbidden creates a TriggerAutomationSuiteUsingPOSTForbidden with default headers values
func NewTriggerAutomationSuiteUsingPOSTForbidden() *TriggerAutomationSuiteUsingPOSTForbidden {
	return &TriggerAutomationSuiteUsingPOSTForbidden{}
}

/*
TriggerAutomationSuiteUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TriggerAutomationSuiteUsingPOSTForbidden struct {
}

// IsSuccess returns true when this trigger automation suite using p o s t forbidden response has a 2xx status code
func (o *TriggerAutomationSuiteUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger automation suite using p o s t forbidden response has a 3xx status code
func (o *TriggerAutomationSuiteUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger automation suite using p o s t forbidden response has a 4xx status code
func (o *TriggerAutomationSuiteUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger automation suite using p o s t forbidden response has a 5xx status code
func (o *TriggerAutomationSuiteUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger automation suite using p o s t forbidden response a status code equal to that given
func (o *TriggerAutomationSuiteUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the trigger automation suite using p o s t forbidden response
func (o *TriggerAutomationSuiteUsingPOSTForbidden) Code() int {
	return 403
}

func (o *TriggerAutomationSuiteUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTForbidden", 403)
}

func (o *TriggerAutomationSuiteUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTForbidden", 403)
}

func (o *TriggerAutomationSuiteUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerAutomationSuiteUsingPOSTNotFound creates a TriggerAutomationSuiteUsingPOSTNotFound with default headers values
func NewTriggerAutomationSuiteUsingPOSTNotFound() *TriggerAutomationSuiteUsingPOSTNotFound {
	return &TriggerAutomationSuiteUsingPOSTNotFound{}
}

/*
TriggerAutomationSuiteUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TriggerAutomationSuiteUsingPOSTNotFound struct {
}

// IsSuccess returns true when this trigger automation suite using p o s t not found response has a 2xx status code
func (o *TriggerAutomationSuiteUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger automation suite using p o s t not found response has a 3xx status code
func (o *TriggerAutomationSuiteUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger automation suite using p o s t not found response has a 4xx status code
func (o *TriggerAutomationSuiteUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger automation suite using p o s t not found response has a 5xx status code
func (o *TriggerAutomationSuiteUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger automation suite using p o s t not found response a status code equal to that given
func (o *TriggerAutomationSuiteUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the trigger automation suite using p o s t not found response
func (o *TriggerAutomationSuiteUsingPOSTNotFound) Code() int {
	return 404
}

func (o *TriggerAutomationSuiteUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTNotFound", 404)
}

func (o *TriggerAutomationSuiteUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc/v1/clusters/{clusterId}/deployments/qa/triggerSuite][%d] triggerAutomationSuiteUsingPOSTNotFound", 404)
}

func (o *TriggerAutomationSuiteUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
