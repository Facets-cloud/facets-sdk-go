// Code generated by go-swagger; DO NOT EDIT.

package ui_deployment_controller

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

// StreamDeploymentLogsUsingGETReader is a Reader for the StreamDeploymentLogsUsingGET structure.
type StreamDeploymentLogsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StreamDeploymentLogsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStreamDeploymentLogsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStreamDeploymentLogsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStreamDeploymentLogsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStreamDeploymentLogsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream] streamDeploymentLogsUsingGET", response, response.Code())
	}
}

// NewStreamDeploymentLogsUsingGETOK creates a StreamDeploymentLogsUsingGETOK with default headers values
func NewStreamDeploymentLogsUsingGETOK() *StreamDeploymentLogsUsingGETOK {
	return &StreamDeploymentLogsUsingGETOK{}
}

/*
StreamDeploymentLogsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type StreamDeploymentLogsUsingGETOK struct {
	Payload models.StreamingResponseBody
}

// IsSuccess returns true when this stream deployment logs using g e t o k response has a 2xx status code
func (o *StreamDeploymentLogsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stream deployment logs using g e t o k response has a 3xx status code
func (o *StreamDeploymentLogsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream deployment logs using g e t o k response has a 4xx status code
func (o *StreamDeploymentLogsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stream deployment logs using g e t o k response has a 5xx status code
func (o *StreamDeploymentLogsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stream deployment logs using g e t o k response a status code equal to that given
func (o *StreamDeploymentLogsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stream deployment logs using g e t o k response
func (o *StreamDeploymentLogsUsingGETOK) Code() int {
	return 200
}

func (o *StreamDeploymentLogsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream][%d] streamDeploymentLogsUsingGETOK %s", 200, payload)
}

func (o *StreamDeploymentLogsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream][%d] streamDeploymentLogsUsingGETOK %s", 200, payload)
}

func (o *StreamDeploymentLogsUsingGETOK) GetPayload() models.StreamingResponseBody {
	return o.Payload
}

func (o *StreamDeploymentLogsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamDeploymentLogsUsingGETUnauthorized creates a StreamDeploymentLogsUsingGETUnauthorized with default headers values
func NewStreamDeploymentLogsUsingGETUnauthorized() *StreamDeploymentLogsUsingGETUnauthorized {
	return &StreamDeploymentLogsUsingGETUnauthorized{}
}

/*
StreamDeploymentLogsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StreamDeploymentLogsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this stream deployment logs using g e t unauthorized response has a 2xx status code
func (o *StreamDeploymentLogsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stream deployment logs using g e t unauthorized response has a 3xx status code
func (o *StreamDeploymentLogsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream deployment logs using g e t unauthorized response has a 4xx status code
func (o *StreamDeploymentLogsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stream deployment logs using g e t unauthorized response has a 5xx status code
func (o *StreamDeploymentLogsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stream deployment logs using g e t unauthorized response a status code equal to that given
func (o *StreamDeploymentLogsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stream deployment logs using g e t unauthorized response
func (o *StreamDeploymentLogsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *StreamDeploymentLogsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream][%d] streamDeploymentLogsUsingGETUnauthorized", 401)
}

func (o *StreamDeploymentLogsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream][%d] streamDeploymentLogsUsingGETUnauthorized", 401)
}

func (o *StreamDeploymentLogsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStreamDeploymentLogsUsingGETForbidden creates a StreamDeploymentLogsUsingGETForbidden with default headers values
func NewStreamDeploymentLogsUsingGETForbidden() *StreamDeploymentLogsUsingGETForbidden {
	return &StreamDeploymentLogsUsingGETForbidden{}
}

/*
StreamDeploymentLogsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StreamDeploymentLogsUsingGETForbidden struct {
}

// IsSuccess returns true when this stream deployment logs using g e t forbidden response has a 2xx status code
func (o *StreamDeploymentLogsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stream deployment logs using g e t forbidden response has a 3xx status code
func (o *StreamDeploymentLogsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream deployment logs using g e t forbidden response has a 4xx status code
func (o *StreamDeploymentLogsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stream deployment logs using g e t forbidden response has a 5xx status code
func (o *StreamDeploymentLogsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stream deployment logs using g e t forbidden response a status code equal to that given
func (o *StreamDeploymentLogsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stream deployment logs using g e t forbidden response
func (o *StreamDeploymentLogsUsingGETForbidden) Code() int {
	return 403
}

func (o *StreamDeploymentLogsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream][%d] streamDeploymentLogsUsingGETForbidden", 403)
}

func (o *StreamDeploymentLogsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream][%d] streamDeploymentLogsUsingGETForbidden", 403)
}

func (o *StreamDeploymentLogsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStreamDeploymentLogsUsingGETNotFound creates a StreamDeploymentLogsUsingGETNotFound with default headers values
func NewStreamDeploymentLogsUsingGETNotFound() *StreamDeploymentLogsUsingGETNotFound {
	return &StreamDeploymentLogsUsingGETNotFound{}
}

/*
StreamDeploymentLogsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StreamDeploymentLogsUsingGETNotFound struct {
}

// IsSuccess returns true when this stream deployment logs using g e t not found response has a 2xx status code
func (o *StreamDeploymentLogsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stream deployment logs using g e t not found response has a 3xx status code
func (o *StreamDeploymentLogsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stream deployment logs using g e t not found response has a 4xx status code
func (o *StreamDeploymentLogsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stream deployment logs using g e t not found response has a 5xx status code
func (o *StreamDeploymentLogsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stream deployment logs using g e t not found response a status code equal to that given
func (o *StreamDeploymentLogsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stream deployment logs using g e t not found response
func (o *StreamDeploymentLogsUsingGETNotFound) Code() int {
	return 404
}

func (o *StreamDeploymentLogsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream][%d] streamDeploymentLogsUsingGETNotFound", 404)
}

func (o *StreamDeploymentLogsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}/logs/stream][%d] streamDeploymentLogsUsingGETNotFound", 404)
}

func (o *StreamDeploymentLogsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
