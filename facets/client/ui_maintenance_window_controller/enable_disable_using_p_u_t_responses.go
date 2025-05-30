// Code generated by go-swagger; DO NOT EDIT.

package ui_maintenance_window_controller

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

// EnableDisableUsingPUTReader is a Reader for the EnableDisableUsingPUT structure.
type EnableDisableUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableDisableUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableDisableUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewEnableDisableUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnableDisableUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnableDisableUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnableDisableUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable] enableDisableUsingPUT", response, response.Code())
	}
}

// NewEnableDisableUsingPUTOK creates a EnableDisableUsingPUTOK with default headers values
func NewEnableDisableUsingPUTOK() *EnableDisableUsingPUTOK {
	return &EnableDisableUsingPUTOK{}
}

/*
EnableDisableUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type EnableDisableUsingPUTOK struct {
	Payload *models.MaintenanceWindowDTO
}

// IsSuccess returns true when this enable disable using p u t o k response has a 2xx status code
func (o *EnableDisableUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable disable using p u t o k response has a 3xx status code
func (o *EnableDisableUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable disable using p u t o k response has a 4xx status code
func (o *EnableDisableUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable disable using p u t o k response has a 5xx status code
func (o *EnableDisableUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enable disable using p u t o k response a status code equal to that given
func (o *EnableDisableUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enable disable using p u t o k response
func (o *EnableDisableUsingPUTOK) Code() int {
	return 200
}

func (o *EnableDisableUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTOK %s", 200, payload)
}

func (o *EnableDisableUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTOK %s", 200, payload)
}

func (o *EnableDisableUsingPUTOK) GetPayload() *models.MaintenanceWindowDTO {
	return o.Payload
}

func (o *EnableDisableUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MaintenanceWindowDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableDisableUsingPUTCreated creates a EnableDisableUsingPUTCreated with default headers values
func NewEnableDisableUsingPUTCreated() *EnableDisableUsingPUTCreated {
	return &EnableDisableUsingPUTCreated{}
}

/*
EnableDisableUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type EnableDisableUsingPUTCreated struct {
}

// IsSuccess returns true when this enable disable using p u t created response has a 2xx status code
func (o *EnableDisableUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable disable using p u t created response has a 3xx status code
func (o *EnableDisableUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable disable using p u t created response has a 4xx status code
func (o *EnableDisableUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable disable using p u t created response has a 5xx status code
func (o *EnableDisableUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this enable disable using p u t created response a status code equal to that given
func (o *EnableDisableUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the enable disable using p u t created response
func (o *EnableDisableUsingPUTCreated) Code() int {
	return 201
}

func (o *EnableDisableUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTCreated", 201)
}

func (o *EnableDisableUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTCreated", 201)
}

func (o *EnableDisableUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableDisableUsingPUTUnauthorized creates a EnableDisableUsingPUTUnauthorized with default headers values
func NewEnableDisableUsingPUTUnauthorized() *EnableDisableUsingPUTUnauthorized {
	return &EnableDisableUsingPUTUnauthorized{}
}

/*
EnableDisableUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type EnableDisableUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this enable disable using p u t unauthorized response has a 2xx status code
func (o *EnableDisableUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable disable using p u t unauthorized response has a 3xx status code
func (o *EnableDisableUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable disable using p u t unauthorized response has a 4xx status code
func (o *EnableDisableUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable disable using p u t unauthorized response has a 5xx status code
func (o *EnableDisableUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enable disable using p u t unauthorized response a status code equal to that given
func (o *EnableDisableUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the enable disable using p u t unauthorized response
func (o *EnableDisableUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *EnableDisableUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTUnauthorized", 401)
}

func (o *EnableDisableUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTUnauthorized", 401)
}

func (o *EnableDisableUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableDisableUsingPUTForbidden creates a EnableDisableUsingPUTForbidden with default headers values
func NewEnableDisableUsingPUTForbidden() *EnableDisableUsingPUTForbidden {
	return &EnableDisableUsingPUTForbidden{}
}

/*
EnableDisableUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EnableDisableUsingPUTForbidden struct {
}

// IsSuccess returns true when this enable disable using p u t forbidden response has a 2xx status code
func (o *EnableDisableUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable disable using p u t forbidden response has a 3xx status code
func (o *EnableDisableUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable disable using p u t forbidden response has a 4xx status code
func (o *EnableDisableUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable disable using p u t forbidden response has a 5xx status code
func (o *EnableDisableUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enable disable using p u t forbidden response a status code equal to that given
func (o *EnableDisableUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the enable disable using p u t forbidden response
func (o *EnableDisableUsingPUTForbidden) Code() int {
	return 403
}

func (o *EnableDisableUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTForbidden", 403)
}

func (o *EnableDisableUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTForbidden", 403)
}

func (o *EnableDisableUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableDisableUsingPUTNotFound creates a EnableDisableUsingPUTNotFound with default headers values
func NewEnableDisableUsingPUTNotFound() *EnableDisableUsingPUTNotFound {
	return &EnableDisableUsingPUTNotFound{}
}

/*
EnableDisableUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type EnableDisableUsingPUTNotFound struct {
}

// IsSuccess returns true when this enable disable using p u t not found response has a 2xx status code
func (o *EnableDisableUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable disable using p u t not found response has a 3xx status code
func (o *EnableDisableUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable disable using p u t not found response has a 4xx status code
func (o *EnableDisableUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable disable using p u t not found response has a 5xx status code
func (o *EnableDisableUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enable disable using p u t not found response a status code equal to that given
func (o *EnableDisableUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the enable disable using p u t not found response
func (o *EnableDisableUsingPUTNotFound) Code() int {
	return 404
}

func (o *EnableDisableUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTNotFound", 404)
}

func (o *EnableDisableUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/maintenance-window/{clusterId}/enable-disable][%d] enableDisableUsingPUTNotFound", 404)
}

func (o *EnableDisableUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
