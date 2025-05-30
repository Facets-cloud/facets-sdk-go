// Code generated by go-swagger; DO NOT EDIT.

package ui_azure_cluster_controller

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

// UpdateAzureClusterUsingPUTReader is a Reader for the UpdateAzureClusterUsingPUT structure.
type UpdateAzureClusterUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAzureClusterUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAzureClusterUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateAzureClusterUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAzureClusterUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAzureClusterUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAzureClusterUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /cc-ui/v1/azure/clusters/{clusterId}] updateAzureClusterUsingPUT", response, response.Code())
	}
}

// NewUpdateAzureClusterUsingPUTOK creates a UpdateAzureClusterUsingPUTOK with default headers values
func NewUpdateAzureClusterUsingPUTOK() *UpdateAzureClusterUsingPUTOK {
	return &UpdateAzureClusterUsingPUTOK{}
}

/*
UpdateAzureClusterUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type UpdateAzureClusterUsingPUTOK struct {
	Payload *models.AzureCluster
}

// IsSuccess returns true when this update azure cluster using p u t o k response has a 2xx status code
func (o *UpdateAzureClusterUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update azure cluster using p u t o k response has a 3xx status code
func (o *UpdateAzureClusterUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update azure cluster using p u t o k response has a 4xx status code
func (o *UpdateAzureClusterUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update azure cluster using p u t o k response has a 5xx status code
func (o *UpdateAzureClusterUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update azure cluster using p u t o k response a status code equal to that given
func (o *UpdateAzureClusterUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update azure cluster using p u t o k response
func (o *UpdateAzureClusterUsingPUTOK) Code() int {
	return 200
}

func (o *UpdateAzureClusterUsingPUTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTOK %s", 200, payload)
}

func (o *UpdateAzureClusterUsingPUTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTOK %s", 200, payload)
}

func (o *UpdateAzureClusterUsingPUTOK) GetPayload() *models.AzureCluster {
	return o.Payload
}

func (o *UpdateAzureClusterUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAzureClusterUsingPUTCreated creates a UpdateAzureClusterUsingPUTCreated with default headers values
func NewUpdateAzureClusterUsingPUTCreated() *UpdateAzureClusterUsingPUTCreated {
	return &UpdateAzureClusterUsingPUTCreated{}
}

/*
UpdateAzureClusterUsingPUTCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateAzureClusterUsingPUTCreated struct {
}

// IsSuccess returns true when this update azure cluster using p u t created response has a 2xx status code
func (o *UpdateAzureClusterUsingPUTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update azure cluster using p u t created response has a 3xx status code
func (o *UpdateAzureClusterUsingPUTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update azure cluster using p u t created response has a 4xx status code
func (o *UpdateAzureClusterUsingPUTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update azure cluster using p u t created response has a 5xx status code
func (o *UpdateAzureClusterUsingPUTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update azure cluster using p u t created response a status code equal to that given
func (o *UpdateAzureClusterUsingPUTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update azure cluster using p u t created response
func (o *UpdateAzureClusterUsingPUTCreated) Code() int {
	return 201
}

func (o *UpdateAzureClusterUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTCreated", 201)
}

func (o *UpdateAzureClusterUsingPUTCreated) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTCreated", 201)
}

func (o *UpdateAzureClusterUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAzureClusterUsingPUTUnauthorized creates a UpdateAzureClusterUsingPUTUnauthorized with default headers values
func NewUpdateAzureClusterUsingPUTUnauthorized() *UpdateAzureClusterUsingPUTUnauthorized {
	return &UpdateAzureClusterUsingPUTUnauthorized{}
}

/*
UpdateAzureClusterUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateAzureClusterUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update azure cluster using p u t unauthorized response has a 2xx status code
func (o *UpdateAzureClusterUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update azure cluster using p u t unauthorized response has a 3xx status code
func (o *UpdateAzureClusterUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update azure cluster using p u t unauthorized response has a 4xx status code
func (o *UpdateAzureClusterUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update azure cluster using p u t unauthorized response has a 5xx status code
func (o *UpdateAzureClusterUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update azure cluster using p u t unauthorized response a status code equal to that given
func (o *UpdateAzureClusterUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update azure cluster using p u t unauthorized response
func (o *UpdateAzureClusterUsingPUTUnauthorized) Code() int {
	return 401
}

func (o *UpdateAzureClusterUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTUnauthorized", 401)
}

func (o *UpdateAzureClusterUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTUnauthorized", 401)
}

func (o *UpdateAzureClusterUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAzureClusterUsingPUTForbidden creates a UpdateAzureClusterUsingPUTForbidden with default headers values
func NewUpdateAzureClusterUsingPUTForbidden() *UpdateAzureClusterUsingPUTForbidden {
	return &UpdateAzureClusterUsingPUTForbidden{}
}

/*
UpdateAzureClusterUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateAzureClusterUsingPUTForbidden struct {
}

// IsSuccess returns true when this update azure cluster using p u t forbidden response has a 2xx status code
func (o *UpdateAzureClusterUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update azure cluster using p u t forbidden response has a 3xx status code
func (o *UpdateAzureClusterUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update azure cluster using p u t forbidden response has a 4xx status code
func (o *UpdateAzureClusterUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update azure cluster using p u t forbidden response has a 5xx status code
func (o *UpdateAzureClusterUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update azure cluster using p u t forbidden response a status code equal to that given
func (o *UpdateAzureClusterUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update azure cluster using p u t forbidden response
func (o *UpdateAzureClusterUsingPUTForbidden) Code() int {
	return 403
}

func (o *UpdateAzureClusterUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTForbidden", 403)
}

func (o *UpdateAzureClusterUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTForbidden", 403)
}

func (o *UpdateAzureClusterUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAzureClusterUsingPUTNotFound creates a UpdateAzureClusterUsingPUTNotFound with default headers values
func NewUpdateAzureClusterUsingPUTNotFound() *UpdateAzureClusterUsingPUTNotFound {
	return &UpdateAzureClusterUsingPUTNotFound{}
}

/*
UpdateAzureClusterUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateAzureClusterUsingPUTNotFound struct {
}

// IsSuccess returns true when this update azure cluster using p u t not found response has a 2xx status code
func (o *UpdateAzureClusterUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update azure cluster using p u t not found response has a 3xx status code
func (o *UpdateAzureClusterUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update azure cluster using p u t not found response has a 4xx status code
func (o *UpdateAzureClusterUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update azure cluster using p u t not found response has a 5xx status code
func (o *UpdateAzureClusterUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update azure cluster using p u t not found response a status code equal to that given
func (o *UpdateAzureClusterUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update azure cluster using p u t not found response
func (o *UpdateAzureClusterUsingPUTNotFound) Code() int {
	return 404
}

func (o *UpdateAzureClusterUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTNotFound", 404)
}

func (o *UpdateAzureClusterUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /cc-ui/v1/azure/clusters/{clusterId}][%d] updateAzureClusterUsingPUTNotFound", 404)
}

func (o *UpdateAzureClusterUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
