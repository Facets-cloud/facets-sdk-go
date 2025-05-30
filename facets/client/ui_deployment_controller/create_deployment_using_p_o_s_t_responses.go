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

// CreateDeploymentUsingPOSTReader is a Reader for the CreateDeploymentUsingPOST structure.
type CreateDeploymentUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeploymentUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDeploymentUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateDeploymentUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateDeploymentUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDeploymentUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDeploymentUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/clusters/{clusterId}/deployments] createDeploymentUsingPOST", response, response.Code())
	}
}

// NewCreateDeploymentUsingPOSTOK creates a CreateDeploymentUsingPOSTOK with default headers values
func NewCreateDeploymentUsingPOSTOK() *CreateDeploymentUsingPOSTOK {
	return &CreateDeploymentUsingPOSTOK{}
}

/*
CreateDeploymentUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type CreateDeploymentUsingPOSTOK struct {
	Payload *models.DeploymentLog
}

// IsSuccess returns true when this create deployment using p o s t o k response has a 2xx status code
func (o *CreateDeploymentUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create deployment using p o s t o k response has a 3xx status code
func (o *CreateDeploymentUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment using p o s t o k response has a 4xx status code
func (o *CreateDeploymentUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create deployment using p o s t o k response has a 5xx status code
func (o *CreateDeploymentUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment using p o s t o k response a status code equal to that given
func (o *CreateDeploymentUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create deployment using p o s t o k response
func (o *CreateDeploymentUsingPOSTOK) Code() int {
	return 200
}

func (o *CreateDeploymentUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTOK %s", 200, payload)
}

func (o *CreateDeploymentUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTOK %s", 200, payload)
}

func (o *CreateDeploymentUsingPOSTOK) GetPayload() *models.DeploymentLog {
	return o.Payload
}

func (o *CreateDeploymentUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentLog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentUsingPOSTCreated creates a CreateDeploymentUsingPOSTCreated with default headers values
func NewCreateDeploymentUsingPOSTCreated() *CreateDeploymentUsingPOSTCreated {
	return &CreateDeploymentUsingPOSTCreated{}
}

/*
CreateDeploymentUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type CreateDeploymentUsingPOSTCreated struct {
}

// IsSuccess returns true when this create deployment using p o s t created response has a 2xx status code
func (o *CreateDeploymentUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create deployment using p o s t created response has a 3xx status code
func (o *CreateDeploymentUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment using p o s t created response has a 4xx status code
func (o *CreateDeploymentUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create deployment using p o s t created response has a 5xx status code
func (o *CreateDeploymentUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment using p o s t created response a status code equal to that given
func (o *CreateDeploymentUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create deployment using p o s t created response
func (o *CreateDeploymentUsingPOSTCreated) Code() int {
	return 201
}

func (o *CreateDeploymentUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTCreated", 201)
}

func (o *CreateDeploymentUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTCreated", 201)
}

func (o *CreateDeploymentUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDeploymentUsingPOSTUnauthorized creates a CreateDeploymentUsingPOSTUnauthorized with default headers values
func NewCreateDeploymentUsingPOSTUnauthorized() *CreateDeploymentUsingPOSTUnauthorized {
	return &CreateDeploymentUsingPOSTUnauthorized{}
}

/*
CreateDeploymentUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateDeploymentUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this create deployment using p o s t unauthorized response has a 2xx status code
func (o *CreateDeploymentUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create deployment using p o s t unauthorized response has a 3xx status code
func (o *CreateDeploymentUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment using p o s t unauthorized response has a 4xx status code
func (o *CreateDeploymentUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create deployment using p o s t unauthorized response has a 5xx status code
func (o *CreateDeploymentUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment using p o s t unauthorized response a status code equal to that given
func (o *CreateDeploymentUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create deployment using p o s t unauthorized response
func (o *CreateDeploymentUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *CreateDeploymentUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTUnauthorized", 401)
}

func (o *CreateDeploymentUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTUnauthorized", 401)
}

func (o *CreateDeploymentUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDeploymentUsingPOSTForbidden creates a CreateDeploymentUsingPOSTForbidden with default headers values
func NewCreateDeploymentUsingPOSTForbidden() *CreateDeploymentUsingPOSTForbidden {
	return &CreateDeploymentUsingPOSTForbidden{}
}

/*
CreateDeploymentUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateDeploymentUsingPOSTForbidden struct {
}

// IsSuccess returns true when this create deployment using p o s t forbidden response has a 2xx status code
func (o *CreateDeploymentUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create deployment using p o s t forbidden response has a 3xx status code
func (o *CreateDeploymentUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment using p o s t forbidden response has a 4xx status code
func (o *CreateDeploymentUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create deployment using p o s t forbidden response has a 5xx status code
func (o *CreateDeploymentUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment using p o s t forbidden response a status code equal to that given
func (o *CreateDeploymentUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create deployment using p o s t forbidden response
func (o *CreateDeploymentUsingPOSTForbidden) Code() int {
	return 403
}

func (o *CreateDeploymentUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTForbidden", 403)
}

func (o *CreateDeploymentUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTForbidden", 403)
}

func (o *CreateDeploymentUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDeploymentUsingPOSTNotFound creates a CreateDeploymentUsingPOSTNotFound with default headers values
func NewCreateDeploymentUsingPOSTNotFound() *CreateDeploymentUsingPOSTNotFound {
	return &CreateDeploymentUsingPOSTNotFound{}
}

/*
CreateDeploymentUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateDeploymentUsingPOSTNotFound struct {
}

// IsSuccess returns true when this create deployment using p o s t not found response has a 2xx status code
func (o *CreateDeploymentUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create deployment using p o s t not found response has a 3xx status code
func (o *CreateDeploymentUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment using p o s t not found response has a 4xx status code
func (o *CreateDeploymentUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create deployment using p o s t not found response has a 5xx status code
func (o *CreateDeploymentUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment using p o s t not found response a status code equal to that given
func (o *CreateDeploymentUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create deployment using p o s t not found response
func (o *CreateDeploymentUsingPOSTNotFound) Code() int {
	return 404
}

func (o *CreateDeploymentUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTNotFound", 404)
}

func (o *CreateDeploymentUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/clusters/{clusterId}/deployments][%d] createDeploymentUsingPOSTNotFound", 404)
}

func (o *CreateDeploymentUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
