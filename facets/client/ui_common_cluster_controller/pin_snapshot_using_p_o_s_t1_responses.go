// Code generated by go-swagger; DO NOT EDIT.

package ui_common_cluster_controller

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

// PinSnapshotUsingPOST1Reader is a Reader for the PinSnapshotUsingPOST1 structure.
type PinSnapshotUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PinSnapshotUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPinSnapshotUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPinSnapshotUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPinSnapshotUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPinSnapshotUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPinSnapshotUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot] pinSnapshotUsingPOST_1", response, response.Code())
	}
}

// NewPinSnapshotUsingPOST1OK creates a PinSnapshotUsingPOST1OK with default headers values
func NewPinSnapshotUsingPOST1OK() *PinSnapshotUsingPOST1OK {
	return &PinSnapshotUsingPOST1OK{}
}

/*
PinSnapshotUsingPOST1OK describes a response with status code 200, with default header values.

OK
*/
type PinSnapshotUsingPOST1OK struct {
	Payload *models.SnapshotInfo
}

// IsSuccess returns true when this pin snapshot using p o s t1 o k response has a 2xx status code
func (o *PinSnapshotUsingPOST1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pin snapshot using p o s t1 o k response has a 3xx status code
func (o *PinSnapshotUsingPOST1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pin snapshot using p o s t1 o k response has a 4xx status code
func (o *PinSnapshotUsingPOST1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pin snapshot using p o s t1 o k response has a 5xx status code
func (o *PinSnapshotUsingPOST1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this pin snapshot using p o s t1 o k response a status code equal to that given
func (o *PinSnapshotUsingPOST1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pin snapshot using p o s t1 o k response
func (o *PinSnapshotUsingPOST1OK) Code() int {
	return 200
}

func (o *PinSnapshotUsingPOST1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1OK %s", 200, payload)
}

func (o *PinSnapshotUsingPOST1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1OK %s", 200, payload)
}

func (o *PinSnapshotUsingPOST1OK) GetPayload() *models.SnapshotInfo {
	return o.Payload
}

func (o *PinSnapshotUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinSnapshotUsingPOST1Created creates a PinSnapshotUsingPOST1Created with default headers values
func NewPinSnapshotUsingPOST1Created() *PinSnapshotUsingPOST1Created {
	return &PinSnapshotUsingPOST1Created{}
}

/*
PinSnapshotUsingPOST1Created describes a response with status code 201, with default header values.

Created
*/
type PinSnapshotUsingPOST1Created struct {
}

// IsSuccess returns true when this pin snapshot using p o s t1 created response has a 2xx status code
func (o *PinSnapshotUsingPOST1Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pin snapshot using p o s t1 created response has a 3xx status code
func (o *PinSnapshotUsingPOST1Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pin snapshot using p o s t1 created response has a 4xx status code
func (o *PinSnapshotUsingPOST1Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this pin snapshot using p o s t1 created response has a 5xx status code
func (o *PinSnapshotUsingPOST1Created) IsServerError() bool {
	return false
}

// IsCode returns true when this pin snapshot using p o s t1 created response a status code equal to that given
func (o *PinSnapshotUsingPOST1Created) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pin snapshot using p o s t1 created response
func (o *PinSnapshotUsingPOST1Created) Code() int {
	return 201
}

func (o *PinSnapshotUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1Created", 201)
}

func (o *PinSnapshotUsingPOST1Created) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1Created", 201)
}

func (o *PinSnapshotUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPinSnapshotUsingPOST1Unauthorized creates a PinSnapshotUsingPOST1Unauthorized with default headers values
func NewPinSnapshotUsingPOST1Unauthorized() *PinSnapshotUsingPOST1Unauthorized {
	return &PinSnapshotUsingPOST1Unauthorized{}
}

/*
PinSnapshotUsingPOST1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PinSnapshotUsingPOST1Unauthorized struct {
}

// IsSuccess returns true when this pin snapshot using p o s t1 unauthorized response has a 2xx status code
func (o *PinSnapshotUsingPOST1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pin snapshot using p o s t1 unauthorized response has a 3xx status code
func (o *PinSnapshotUsingPOST1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pin snapshot using p o s t1 unauthorized response has a 4xx status code
func (o *PinSnapshotUsingPOST1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pin snapshot using p o s t1 unauthorized response has a 5xx status code
func (o *PinSnapshotUsingPOST1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pin snapshot using p o s t1 unauthorized response a status code equal to that given
func (o *PinSnapshotUsingPOST1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pin snapshot using p o s t1 unauthorized response
func (o *PinSnapshotUsingPOST1Unauthorized) Code() int {
	return 401
}

func (o *PinSnapshotUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1Unauthorized", 401)
}

func (o *PinSnapshotUsingPOST1Unauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1Unauthorized", 401)
}

func (o *PinSnapshotUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPinSnapshotUsingPOST1Forbidden creates a PinSnapshotUsingPOST1Forbidden with default headers values
func NewPinSnapshotUsingPOST1Forbidden() *PinSnapshotUsingPOST1Forbidden {
	return &PinSnapshotUsingPOST1Forbidden{}
}

/*
PinSnapshotUsingPOST1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PinSnapshotUsingPOST1Forbidden struct {
}

// IsSuccess returns true when this pin snapshot using p o s t1 forbidden response has a 2xx status code
func (o *PinSnapshotUsingPOST1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pin snapshot using p o s t1 forbidden response has a 3xx status code
func (o *PinSnapshotUsingPOST1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pin snapshot using p o s t1 forbidden response has a 4xx status code
func (o *PinSnapshotUsingPOST1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pin snapshot using p o s t1 forbidden response has a 5xx status code
func (o *PinSnapshotUsingPOST1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pin snapshot using p o s t1 forbidden response a status code equal to that given
func (o *PinSnapshotUsingPOST1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pin snapshot using p o s t1 forbidden response
func (o *PinSnapshotUsingPOST1Forbidden) Code() int {
	return 403
}

func (o *PinSnapshotUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1Forbidden", 403)
}

func (o *PinSnapshotUsingPOST1Forbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1Forbidden", 403)
}

func (o *PinSnapshotUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPinSnapshotUsingPOST1NotFound creates a PinSnapshotUsingPOST1NotFound with default headers values
func NewPinSnapshotUsingPOST1NotFound() *PinSnapshotUsingPOST1NotFound {
	return &PinSnapshotUsingPOST1NotFound{}
}

/*
PinSnapshotUsingPOST1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PinSnapshotUsingPOST1NotFound struct {
}

// IsSuccess returns true when this pin snapshot using p o s t1 not found response has a 2xx status code
func (o *PinSnapshotUsingPOST1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pin snapshot using p o s t1 not found response has a 3xx status code
func (o *PinSnapshotUsingPOST1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pin snapshot using p o s t1 not found response has a 4xx status code
func (o *PinSnapshotUsingPOST1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pin snapshot using p o s t1 not found response has a 5xx status code
func (o *PinSnapshotUsingPOST1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pin snapshot using p o s t1 not found response a status code equal to that given
func (o *PinSnapshotUsingPOST1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pin snapshot using p o s t1 not found response
func (o *PinSnapshotUsingPOST1NotFound) Code() int {
	return 404
}

func (o *PinSnapshotUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1NotFound", 404)
}

func (o *PinSnapshotUsingPOST1NotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/dr/{resourceType}/snapshots/{instanceName}/pinnedSnapshot][%d] pinSnapshotUsingPOST1NotFound", 404)
}

func (o *PinSnapshotUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
