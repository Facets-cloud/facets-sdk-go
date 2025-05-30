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

// GetAvailabilitySchedulesUsingGETReader is a Reader for the GetAvailabilitySchedulesUsingGET structure.
type GetAvailabilitySchedulesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailabilitySchedulesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAvailabilitySchedulesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAvailabilitySchedulesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAvailabilitySchedulesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAvailabilitySchedulesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule] getAvailabilitySchedulesUsingGET", response, response.Code())
	}
}

// NewGetAvailabilitySchedulesUsingGETOK creates a GetAvailabilitySchedulesUsingGETOK with default headers values
func NewGetAvailabilitySchedulesUsingGETOK() *GetAvailabilitySchedulesUsingGETOK {
	return &GetAvailabilitySchedulesUsingGETOK{}
}

/*
GetAvailabilitySchedulesUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetAvailabilitySchedulesUsingGETOK struct {
	Payload *models.AvailabilitySchedule
}

// IsSuccess returns true when this get availability schedules using g e t o k response has a 2xx status code
func (o *GetAvailabilitySchedulesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get availability schedules using g e t o k response has a 3xx status code
func (o *GetAvailabilitySchedulesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get availability schedules using g e t o k response has a 4xx status code
func (o *GetAvailabilitySchedulesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get availability schedules using g e t o k response has a 5xx status code
func (o *GetAvailabilitySchedulesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get availability schedules using g e t o k response a status code equal to that given
func (o *GetAvailabilitySchedulesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get availability schedules using g e t o k response
func (o *GetAvailabilitySchedulesUsingGETOK) Code() int {
	return 200
}

func (o *GetAvailabilitySchedulesUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule][%d] getAvailabilitySchedulesUsingGETOK %s", 200, payload)
}

func (o *GetAvailabilitySchedulesUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule][%d] getAvailabilitySchedulesUsingGETOK %s", 200, payload)
}

func (o *GetAvailabilitySchedulesUsingGETOK) GetPayload() *models.AvailabilitySchedule {
	return o.Payload
}

func (o *GetAvailabilitySchedulesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailabilitySchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailabilitySchedulesUsingGETUnauthorized creates a GetAvailabilitySchedulesUsingGETUnauthorized with default headers values
func NewGetAvailabilitySchedulesUsingGETUnauthorized() *GetAvailabilitySchedulesUsingGETUnauthorized {
	return &GetAvailabilitySchedulesUsingGETUnauthorized{}
}

/*
GetAvailabilitySchedulesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAvailabilitySchedulesUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get availability schedules using g e t unauthorized response has a 2xx status code
func (o *GetAvailabilitySchedulesUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get availability schedules using g e t unauthorized response has a 3xx status code
func (o *GetAvailabilitySchedulesUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get availability schedules using g e t unauthorized response has a 4xx status code
func (o *GetAvailabilitySchedulesUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get availability schedules using g e t unauthorized response has a 5xx status code
func (o *GetAvailabilitySchedulesUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get availability schedules using g e t unauthorized response a status code equal to that given
func (o *GetAvailabilitySchedulesUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get availability schedules using g e t unauthorized response
func (o *GetAvailabilitySchedulesUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetAvailabilitySchedulesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule][%d] getAvailabilitySchedulesUsingGETUnauthorized", 401)
}

func (o *GetAvailabilitySchedulesUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule][%d] getAvailabilitySchedulesUsingGETUnauthorized", 401)
}

func (o *GetAvailabilitySchedulesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAvailabilitySchedulesUsingGETForbidden creates a GetAvailabilitySchedulesUsingGETForbidden with default headers values
func NewGetAvailabilitySchedulesUsingGETForbidden() *GetAvailabilitySchedulesUsingGETForbidden {
	return &GetAvailabilitySchedulesUsingGETForbidden{}
}

/*
GetAvailabilitySchedulesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAvailabilitySchedulesUsingGETForbidden struct {
}

// IsSuccess returns true when this get availability schedules using g e t forbidden response has a 2xx status code
func (o *GetAvailabilitySchedulesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get availability schedules using g e t forbidden response has a 3xx status code
func (o *GetAvailabilitySchedulesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get availability schedules using g e t forbidden response has a 4xx status code
func (o *GetAvailabilitySchedulesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get availability schedules using g e t forbidden response has a 5xx status code
func (o *GetAvailabilitySchedulesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get availability schedules using g e t forbidden response a status code equal to that given
func (o *GetAvailabilitySchedulesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get availability schedules using g e t forbidden response
func (o *GetAvailabilitySchedulesUsingGETForbidden) Code() int {
	return 403
}

func (o *GetAvailabilitySchedulesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule][%d] getAvailabilitySchedulesUsingGETForbidden", 403)
}

func (o *GetAvailabilitySchedulesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule][%d] getAvailabilitySchedulesUsingGETForbidden", 403)
}

func (o *GetAvailabilitySchedulesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAvailabilitySchedulesUsingGETNotFound creates a GetAvailabilitySchedulesUsingGETNotFound with default headers values
func NewGetAvailabilitySchedulesUsingGETNotFound() *GetAvailabilitySchedulesUsingGETNotFound {
	return &GetAvailabilitySchedulesUsingGETNotFound{}
}

/*
GetAvailabilitySchedulesUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAvailabilitySchedulesUsingGETNotFound struct {
}

// IsSuccess returns true when this get availability schedules using g e t not found response has a 2xx status code
func (o *GetAvailabilitySchedulesUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get availability schedules using g e t not found response has a 3xx status code
func (o *GetAvailabilitySchedulesUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get availability schedules using g e t not found response has a 4xx status code
func (o *GetAvailabilitySchedulesUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get availability schedules using g e t not found response has a 5xx status code
func (o *GetAvailabilitySchedulesUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get availability schedules using g e t not found response a status code equal to that given
func (o *GetAvailabilitySchedulesUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get availability schedules using g e t not found response
func (o *GetAvailabilitySchedulesUsingGETNotFound) Code() int {
	return 404
}

func (o *GetAvailabilitySchedulesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule][%d] getAvailabilitySchedulesUsingGETNotFound", 404)
}

func (o *GetAvailabilitySchedulesUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/availability-schedule][%d] getAvailabilitySchedulesUsingGETNotFound", 404)
}

func (o *GetAvailabilitySchedulesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
