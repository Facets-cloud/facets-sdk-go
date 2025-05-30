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

// GetDeploymentsUsingGET1Reader is a Reader for the GetDeploymentsUsingGET1 structure.
type GetDeploymentsUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentsUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentsUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentsUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeploymentsUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentsUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/clusters/{clusterId}/deployments] getDeploymentsUsingGET_1", response, response.Code())
	}
}

// NewGetDeploymentsUsingGET1OK creates a GetDeploymentsUsingGET1OK with default headers values
func NewGetDeploymentsUsingGET1OK() *GetDeploymentsUsingGET1OK {
	return &GetDeploymentsUsingGET1OK{}
}

/*
GetDeploymentsUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentsUsingGET1OK struct {
	Payload *models.ListDeploymentsWrapper
}

// IsSuccess returns true when this get deployments using g e t1 o k response has a 2xx status code
func (o *GetDeploymentsUsingGET1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployments using g e t1 o k response has a 3xx status code
func (o *GetDeploymentsUsingGET1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments using g e t1 o k response has a 4xx status code
func (o *GetDeploymentsUsingGET1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployments using g e t1 o k response has a 5xx status code
func (o *GetDeploymentsUsingGET1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments using g e t1 o k response a status code equal to that given
func (o *GetDeploymentsUsingGET1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployments using g e t1 o k response
func (o *GetDeploymentsUsingGET1OK) Code() int {
	return 200
}

func (o *GetDeploymentsUsingGET1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments][%d] getDeploymentsUsingGET1OK %s", 200, payload)
}

func (o *GetDeploymentsUsingGET1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments][%d] getDeploymentsUsingGET1OK %s", 200, payload)
}

func (o *GetDeploymentsUsingGET1OK) GetPayload() *models.ListDeploymentsWrapper {
	return o.Payload
}

func (o *GetDeploymentsUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDeploymentsWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsUsingGET1Unauthorized creates a GetDeploymentsUsingGET1Unauthorized with default headers values
func NewGetDeploymentsUsingGET1Unauthorized() *GetDeploymentsUsingGET1Unauthorized {
	return &GetDeploymentsUsingGET1Unauthorized{}
}

/*
GetDeploymentsUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentsUsingGET1Unauthorized struct {
}

// IsSuccess returns true when this get deployments using g e t1 unauthorized response has a 2xx status code
func (o *GetDeploymentsUsingGET1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments using g e t1 unauthorized response has a 3xx status code
func (o *GetDeploymentsUsingGET1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments using g e t1 unauthorized response has a 4xx status code
func (o *GetDeploymentsUsingGET1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments using g e t1 unauthorized response has a 5xx status code
func (o *GetDeploymentsUsingGET1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments using g e t1 unauthorized response a status code equal to that given
func (o *GetDeploymentsUsingGET1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get deployments using g e t1 unauthorized response
func (o *GetDeploymentsUsingGET1Unauthorized) Code() int {
	return 401
}

func (o *GetDeploymentsUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments][%d] getDeploymentsUsingGET1Unauthorized", 401)
}

func (o *GetDeploymentsUsingGET1Unauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments][%d] getDeploymentsUsingGET1Unauthorized", 401)
}

func (o *GetDeploymentsUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentsUsingGET1Forbidden creates a GetDeploymentsUsingGET1Forbidden with default headers values
func NewGetDeploymentsUsingGET1Forbidden() *GetDeploymentsUsingGET1Forbidden {
	return &GetDeploymentsUsingGET1Forbidden{}
}

/*
GetDeploymentsUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDeploymentsUsingGET1Forbidden struct {
}

// IsSuccess returns true when this get deployments using g e t1 forbidden response has a 2xx status code
func (o *GetDeploymentsUsingGET1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments using g e t1 forbidden response has a 3xx status code
func (o *GetDeploymentsUsingGET1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments using g e t1 forbidden response has a 4xx status code
func (o *GetDeploymentsUsingGET1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments using g e t1 forbidden response has a 5xx status code
func (o *GetDeploymentsUsingGET1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments using g e t1 forbidden response a status code equal to that given
func (o *GetDeploymentsUsingGET1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get deployments using g e t1 forbidden response
func (o *GetDeploymentsUsingGET1Forbidden) Code() int {
	return 403
}

func (o *GetDeploymentsUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments][%d] getDeploymentsUsingGET1Forbidden", 403)
}

func (o *GetDeploymentsUsingGET1Forbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments][%d] getDeploymentsUsingGET1Forbidden", 403)
}

func (o *GetDeploymentsUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentsUsingGET1NotFound creates a GetDeploymentsUsingGET1NotFound with default headers values
func NewGetDeploymentsUsingGET1NotFound() *GetDeploymentsUsingGET1NotFound {
	return &GetDeploymentsUsingGET1NotFound{}
}

/*
GetDeploymentsUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeploymentsUsingGET1NotFound struct {
}

// IsSuccess returns true when this get deployments using g e t1 not found response has a 2xx status code
func (o *GetDeploymentsUsingGET1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployments using g e t1 not found response has a 3xx status code
func (o *GetDeploymentsUsingGET1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployments using g e t1 not found response has a 4xx status code
func (o *GetDeploymentsUsingGET1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployments using g e t1 not found response has a 5xx status code
func (o *GetDeploymentsUsingGET1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployments using g e t1 not found response a status code equal to that given
func (o *GetDeploymentsUsingGET1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deployments using g e t1 not found response
func (o *GetDeploymentsUsingGET1NotFound) Code() int {
	return 404
}

func (o *GetDeploymentsUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments][%d] getDeploymentsUsingGET1NotFound", 404)
}

func (o *GetDeploymentsUsingGET1NotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/clusters/{clusterId}/deployments][%d] getDeploymentsUsingGET1NotFound", 404)
}

func (o *GetDeploymentsUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
