// Code generated by go-swagger; DO NOT EDIT.

package ui_stack_controller

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

// GetRunningBaseClustersUsingGETReader is a Reader for the GetRunningBaseClustersUsingGET structure.
type GetRunningBaseClustersUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunningBaseClustersUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunningBaseClustersUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRunningBaseClustersUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRunningBaseClustersUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunningBaseClustersUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/stacks/running-base-clusters] getRunningBaseClustersUsingGET", response, response.Code())
	}
}

// NewGetRunningBaseClustersUsingGETOK creates a GetRunningBaseClustersUsingGETOK with default headers values
func NewGetRunningBaseClustersUsingGETOK() *GetRunningBaseClustersUsingGETOK {
	return &GetRunningBaseClustersUsingGETOK{}
}

/*
GetRunningBaseClustersUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetRunningBaseClustersUsingGETOK struct {
	Payload []*models.AbstractCluster
}

// IsSuccess returns true when this get running base clusters using g e t o k response has a 2xx status code
func (o *GetRunningBaseClustersUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get running base clusters using g e t o k response has a 3xx status code
func (o *GetRunningBaseClustersUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running base clusters using g e t o k response has a 4xx status code
func (o *GetRunningBaseClustersUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get running base clusters using g e t o k response has a 5xx status code
func (o *GetRunningBaseClustersUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get running base clusters using g e t o k response a status code equal to that given
func (o *GetRunningBaseClustersUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get running base clusters using g e t o k response
func (o *GetRunningBaseClustersUsingGETOK) Code() int {
	return 200
}

func (o *GetRunningBaseClustersUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/running-base-clusters][%d] getRunningBaseClustersUsingGETOK %s", 200, payload)
}

func (o *GetRunningBaseClustersUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/running-base-clusters][%d] getRunningBaseClustersUsingGETOK %s", 200, payload)
}

func (o *GetRunningBaseClustersUsingGETOK) GetPayload() []*models.AbstractCluster {
	return o.Payload
}

func (o *GetRunningBaseClustersUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunningBaseClustersUsingGETUnauthorized creates a GetRunningBaseClustersUsingGETUnauthorized with default headers values
func NewGetRunningBaseClustersUsingGETUnauthorized() *GetRunningBaseClustersUsingGETUnauthorized {
	return &GetRunningBaseClustersUsingGETUnauthorized{}
}

/*
GetRunningBaseClustersUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRunningBaseClustersUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get running base clusters using g e t unauthorized response has a 2xx status code
func (o *GetRunningBaseClustersUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get running base clusters using g e t unauthorized response has a 3xx status code
func (o *GetRunningBaseClustersUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running base clusters using g e t unauthorized response has a 4xx status code
func (o *GetRunningBaseClustersUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get running base clusters using g e t unauthorized response has a 5xx status code
func (o *GetRunningBaseClustersUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get running base clusters using g e t unauthorized response a status code equal to that given
func (o *GetRunningBaseClustersUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get running base clusters using g e t unauthorized response
func (o *GetRunningBaseClustersUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetRunningBaseClustersUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/running-base-clusters][%d] getRunningBaseClustersUsingGETUnauthorized", 401)
}

func (o *GetRunningBaseClustersUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/running-base-clusters][%d] getRunningBaseClustersUsingGETUnauthorized", 401)
}

func (o *GetRunningBaseClustersUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRunningBaseClustersUsingGETForbidden creates a GetRunningBaseClustersUsingGETForbidden with default headers values
func NewGetRunningBaseClustersUsingGETForbidden() *GetRunningBaseClustersUsingGETForbidden {
	return &GetRunningBaseClustersUsingGETForbidden{}
}

/*
GetRunningBaseClustersUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRunningBaseClustersUsingGETForbidden struct {
}

// IsSuccess returns true when this get running base clusters using g e t forbidden response has a 2xx status code
func (o *GetRunningBaseClustersUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get running base clusters using g e t forbidden response has a 3xx status code
func (o *GetRunningBaseClustersUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running base clusters using g e t forbidden response has a 4xx status code
func (o *GetRunningBaseClustersUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get running base clusters using g e t forbidden response has a 5xx status code
func (o *GetRunningBaseClustersUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get running base clusters using g e t forbidden response a status code equal to that given
func (o *GetRunningBaseClustersUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get running base clusters using g e t forbidden response
func (o *GetRunningBaseClustersUsingGETForbidden) Code() int {
	return 403
}

func (o *GetRunningBaseClustersUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/running-base-clusters][%d] getRunningBaseClustersUsingGETForbidden", 403)
}

func (o *GetRunningBaseClustersUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/running-base-clusters][%d] getRunningBaseClustersUsingGETForbidden", 403)
}

func (o *GetRunningBaseClustersUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRunningBaseClustersUsingGETNotFound creates a GetRunningBaseClustersUsingGETNotFound with default headers values
func NewGetRunningBaseClustersUsingGETNotFound() *GetRunningBaseClustersUsingGETNotFound {
	return &GetRunningBaseClustersUsingGETNotFound{}
}

/*
GetRunningBaseClustersUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRunningBaseClustersUsingGETNotFound struct {
}

// IsSuccess returns true when this get running base clusters using g e t not found response has a 2xx status code
func (o *GetRunningBaseClustersUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get running base clusters using g e t not found response has a 3xx status code
func (o *GetRunningBaseClustersUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get running base clusters using g e t not found response has a 4xx status code
func (o *GetRunningBaseClustersUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get running base clusters using g e t not found response has a 5xx status code
func (o *GetRunningBaseClustersUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get running base clusters using g e t not found response a status code equal to that given
func (o *GetRunningBaseClustersUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get running base clusters using g e t not found response
func (o *GetRunningBaseClustersUsingGETNotFound) Code() int {
	return 404
}

func (o *GetRunningBaseClustersUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/running-base-clusters][%d] getRunningBaseClustersUsingGETNotFound", 404)
}

func (o *GetRunningBaseClustersUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/stacks/running-base-clusters][%d] getRunningBaseClustersUsingGETNotFound", 404)
}

func (o *GetRunningBaseClustersUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
