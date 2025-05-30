// Code generated by go-swagger; DO NOT EDIT.

package ui_kubernetes_explorer_controller

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

// ContainerLogsUsingGETReader is a Reader for the ContainerLogsUsingGET structure.
type ContainerLogsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerLogsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerLogsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewContainerLogsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewContainerLogsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerLogsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs] containerLogsUsingGET", response, response.Code())
	}
}

// NewContainerLogsUsingGETOK creates a ContainerLogsUsingGETOK with default headers values
func NewContainerLogsUsingGETOK() *ContainerLogsUsingGETOK {
	return &ContainerLogsUsingGETOK{}
}

/*
ContainerLogsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type ContainerLogsUsingGETOK struct {
	Payload models.StreamingResponseBody
}

// IsSuccess returns true when this container logs using g e t o k response has a 2xx status code
func (o *ContainerLogsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this container logs using g e t o k response has a 3xx status code
func (o *ContainerLogsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container logs using g e t o k response has a 4xx status code
func (o *ContainerLogsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this container logs using g e t o k response has a 5xx status code
func (o *ContainerLogsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this container logs using g e t o k response a status code equal to that given
func (o *ContainerLogsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the container logs using g e t o k response
func (o *ContainerLogsUsingGETOK) Code() int {
	return 200
}

func (o *ContainerLogsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs][%d] containerLogsUsingGETOK %s", 200, payload)
}

func (o *ContainerLogsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs][%d] containerLogsUsingGETOK %s", 200, payload)
}

func (o *ContainerLogsUsingGETOK) GetPayload() models.StreamingResponseBody {
	return o.Payload
}

func (o *ContainerLogsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerLogsUsingGETUnauthorized creates a ContainerLogsUsingGETUnauthorized with default headers values
func NewContainerLogsUsingGETUnauthorized() *ContainerLogsUsingGETUnauthorized {
	return &ContainerLogsUsingGETUnauthorized{}
}

/*
ContainerLogsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ContainerLogsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this container logs using g e t unauthorized response has a 2xx status code
func (o *ContainerLogsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container logs using g e t unauthorized response has a 3xx status code
func (o *ContainerLogsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container logs using g e t unauthorized response has a 4xx status code
func (o *ContainerLogsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this container logs using g e t unauthorized response has a 5xx status code
func (o *ContainerLogsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this container logs using g e t unauthorized response a status code equal to that given
func (o *ContainerLogsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the container logs using g e t unauthorized response
func (o *ContainerLogsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *ContainerLogsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs][%d] containerLogsUsingGETUnauthorized", 401)
}

func (o *ContainerLogsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs][%d] containerLogsUsingGETUnauthorized", 401)
}

func (o *ContainerLogsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerLogsUsingGETForbidden creates a ContainerLogsUsingGETForbidden with default headers values
func NewContainerLogsUsingGETForbidden() *ContainerLogsUsingGETForbidden {
	return &ContainerLogsUsingGETForbidden{}
}

/*
ContainerLogsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ContainerLogsUsingGETForbidden struct {
}

// IsSuccess returns true when this container logs using g e t forbidden response has a 2xx status code
func (o *ContainerLogsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container logs using g e t forbidden response has a 3xx status code
func (o *ContainerLogsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container logs using g e t forbidden response has a 4xx status code
func (o *ContainerLogsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this container logs using g e t forbidden response has a 5xx status code
func (o *ContainerLogsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this container logs using g e t forbidden response a status code equal to that given
func (o *ContainerLogsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the container logs using g e t forbidden response
func (o *ContainerLogsUsingGETForbidden) Code() int {
	return 403
}

func (o *ContainerLogsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs][%d] containerLogsUsingGETForbidden", 403)
}

func (o *ContainerLogsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs][%d] containerLogsUsingGETForbidden", 403)
}

func (o *ContainerLogsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerLogsUsingGETNotFound creates a ContainerLogsUsingGETNotFound with default headers values
func NewContainerLogsUsingGETNotFound() *ContainerLogsUsingGETNotFound {
	return &ContainerLogsUsingGETNotFound{}
}

/*
ContainerLogsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ContainerLogsUsingGETNotFound struct {
}

// IsSuccess returns true when this container logs using g e t not found response has a 2xx status code
func (o *ContainerLogsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this container logs using g e t not found response has a 3xx status code
func (o *ContainerLogsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this container logs using g e t not found response has a 4xx status code
func (o *ContainerLogsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this container logs using g e t not found response has a 5xx status code
func (o *ContainerLogsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this container logs using g e t not found response a status code equal to that given
func (o *ContainerLogsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the container logs using g e t not found response
func (o *ContainerLogsUsingGETNotFound) Code() int {
	return 404
}

func (o *ContainerLogsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs][%d] containerLogsUsingGETNotFound", 404)
}

func (o *ContainerLogsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/k8s-explorer/pods/{podName}/{containerName}/logs][%d] containerLogsUsingGETNotFound", 404)
}

func (o *ContainerLogsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
