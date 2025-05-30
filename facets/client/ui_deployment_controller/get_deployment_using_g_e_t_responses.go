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

// GetDeploymentUsingGETReader is a Reader for the GetDeploymentUsingGET structure.
type GetDeploymentUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeploymentUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}] getDeploymentUsingGET", response, response.Code())
	}
}

// NewGetDeploymentUsingGETOK creates a GetDeploymentUsingGETOK with default headers values
func NewGetDeploymentUsingGETOK() *GetDeploymentUsingGETOK {
	return &GetDeploymentUsingGETOK{}
}

/*
GetDeploymentUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentUsingGETOK struct {
	Payload *models.DeploymentLog
}

// IsSuccess returns true when this get deployment using g e t o k response has a 2xx status code
func (o *GetDeploymentUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment using g e t o k response has a 3xx status code
func (o *GetDeploymentUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment using g e t o k response has a 4xx status code
func (o *GetDeploymentUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment using g e t o k response has a 5xx status code
func (o *GetDeploymentUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment using g e t o k response a status code equal to that given
func (o *GetDeploymentUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployment using g e t o k response
func (o *GetDeploymentUsingGETOK) Code() int {
	return 200
}

func (o *GetDeploymentUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}][%d] getDeploymentUsingGETOK %s", 200, payload)
}

func (o *GetDeploymentUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}][%d] getDeploymentUsingGETOK %s", 200, payload)
}

func (o *GetDeploymentUsingGETOK) GetPayload() *models.DeploymentLog {
	return o.Payload
}

func (o *GetDeploymentUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentLog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentUsingGETUnauthorized creates a GetDeploymentUsingGETUnauthorized with default headers values
func NewGetDeploymentUsingGETUnauthorized() *GetDeploymentUsingGETUnauthorized {
	return &GetDeploymentUsingGETUnauthorized{}
}

/*
GetDeploymentUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get deployment using g e t unauthorized response has a 2xx status code
func (o *GetDeploymentUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment using g e t unauthorized response has a 3xx status code
func (o *GetDeploymentUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment using g e t unauthorized response has a 4xx status code
func (o *GetDeploymentUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment using g e t unauthorized response has a 5xx status code
func (o *GetDeploymentUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment using g e t unauthorized response a status code equal to that given
func (o *GetDeploymentUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get deployment using g e t unauthorized response
func (o *GetDeploymentUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetDeploymentUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}][%d] getDeploymentUsingGETUnauthorized", 401)
}

func (o *GetDeploymentUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}][%d] getDeploymentUsingGETUnauthorized", 401)
}

func (o *GetDeploymentUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentUsingGETForbidden creates a GetDeploymentUsingGETForbidden with default headers values
func NewGetDeploymentUsingGETForbidden() *GetDeploymentUsingGETForbidden {
	return &GetDeploymentUsingGETForbidden{}
}

/*
GetDeploymentUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDeploymentUsingGETForbidden struct {
}

// IsSuccess returns true when this get deployment using g e t forbidden response has a 2xx status code
func (o *GetDeploymentUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment using g e t forbidden response has a 3xx status code
func (o *GetDeploymentUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment using g e t forbidden response has a 4xx status code
func (o *GetDeploymentUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment using g e t forbidden response has a 5xx status code
func (o *GetDeploymentUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment using g e t forbidden response a status code equal to that given
func (o *GetDeploymentUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get deployment using g e t forbidden response
func (o *GetDeploymentUsingGETForbidden) Code() int {
	return 403
}

func (o *GetDeploymentUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}][%d] getDeploymentUsingGETForbidden", 403)
}

func (o *GetDeploymentUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}][%d] getDeploymentUsingGETForbidden", 403)
}

func (o *GetDeploymentUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentUsingGETNotFound creates a GetDeploymentUsingGETNotFound with default headers values
func NewGetDeploymentUsingGETNotFound() *GetDeploymentUsingGETNotFound {
	return &GetDeploymentUsingGETNotFound{}
}

/*
GetDeploymentUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeploymentUsingGETNotFound struct {
}

// IsSuccess returns true when this get deployment using g e t not found response has a 2xx status code
func (o *GetDeploymentUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment using g e t not found response has a 3xx status code
func (o *GetDeploymentUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment using g e t not found response has a 4xx status code
func (o *GetDeploymentUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment using g e t not found response has a 5xx status code
func (o *GetDeploymentUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment using g e t not found response a status code equal to that given
func (o *GetDeploymentUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deployment using g e t not found response
func (o *GetDeploymentUsingGETNotFound) Code() int {
	return 404
}

func (o *GetDeploymentUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}][%d] getDeploymentUsingGETNotFound", 404)
}

func (o *GetDeploymentUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments/{deploymentId}][%d] getDeploymentUsingGETNotFound", 404)
}

func (o *GetDeploymentUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
