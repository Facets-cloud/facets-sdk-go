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
)

// DeleteClusterUsingDELETE1Reader is a Reader for the DeleteClusterUsingDELETE1 structure.
type DeleteClusterUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterUsingDELETE1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteClusterUsingDELETE1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteClusterUsingDELETE1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteClusterUsingDELETE1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cc-ui/v1/clusters/{clusterId}] deleteClusterUsingDELETE_1", response, response.Code())
	}
}

// NewDeleteClusterUsingDELETE1OK creates a DeleteClusterUsingDELETE1OK with default headers values
func NewDeleteClusterUsingDELETE1OK() *DeleteClusterUsingDELETE1OK {
	return &DeleteClusterUsingDELETE1OK{}
}

/*
DeleteClusterUsingDELETE1OK describes a response with status code 200, with default header values.

OK
*/
type DeleteClusterUsingDELETE1OK struct {
	Payload bool
}

// IsSuccess returns true when this delete cluster using d e l e t e1 o k response has a 2xx status code
func (o *DeleteClusterUsingDELETE1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cluster using d e l e t e1 o k response has a 3xx status code
func (o *DeleteClusterUsingDELETE1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster using d e l e t e1 o k response has a 4xx status code
func (o *DeleteClusterUsingDELETE1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cluster using d e l e t e1 o k response has a 5xx status code
func (o *DeleteClusterUsingDELETE1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster using d e l e t e1 o k response a status code equal to that given
func (o *DeleteClusterUsingDELETE1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete cluster using d e l e t e1 o k response
func (o *DeleteClusterUsingDELETE1OK) Code() int {
	return 200
}

func (o *DeleteClusterUsingDELETE1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cc-ui/v1/clusters/{clusterId}][%d] deleteClusterUsingDELETE1OK %s", 200, payload)
}

func (o *DeleteClusterUsingDELETE1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cc-ui/v1/clusters/{clusterId}][%d] deleteClusterUsingDELETE1OK %s", 200, payload)
}

func (o *DeleteClusterUsingDELETE1OK) GetPayload() bool {
	return o.Payload
}

func (o *DeleteClusterUsingDELETE1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterUsingDELETE1NoContent creates a DeleteClusterUsingDELETE1NoContent with default headers values
func NewDeleteClusterUsingDELETE1NoContent() *DeleteClusterUsingDELETE1NoContent {
	return &DeleteClusterUsingDELETE1NoContent{}
}

/*
DeleteClusterUsingDELETE1NoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteClusterUsingDELETE1NoContent struct {
}

// IsSuccess returns true when this delete cluster using d e l e t e1 no content response has a 2xx status code
func (o *DeleteClusterUsingDELETE1NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cluster using d e l e t e1 no content response has a 3xx status code
func (o *DeleteClusterUsingDELETE1NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster using d e l e t e1 no content response has a 4xx status code
func (o *DeleteClusterUsingDELETE1NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cluster using d e l e t e1 no content response has a 5xx status code
func (o *DeleteClusterUsingDELETE1NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster using d e l e t e1 no content response a status code equal to that given
func (o *DeleteClusterUsingDELETE1NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete cluster using d e l e t e1 no content response
func (o *DeleteClusterUsingDELETE1NoContent) Code() int {
	return 204
}

func (o *DeleteClusterUsingDELETE1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/clusters/{clusterId}][%d] deleteClusterUsingDELETE1NoContent", 204)
}

func (o *DeleteClusterUsingDELETE1NoContent) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/clusters/{clusterId}][%d] deleteClusterUsingDELETE1NoContent", 204)
}

func (o *DeleteClusterUsingDELETE1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterUsingDELETE1Unauthorized creates a DeleteClusterUsingDELETE1Unauthorized with default headers values
func NewDeleteClusterUsingDELETE1Unauthorized() *DeleteClusterUsingDELETE1Unauthorized {
	return &DeleteClusterUsingDELETE1Unauthorized{}
}

/*
DeleteClusterUsingDELETE1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteClusterUsingDELETE1Unauthorized struct {
}

// IsSuccess returns true when this delete cluster using d e l e t e1 unauthorized response has a 2xx status code
func (o *DeleteClusterUsingDELETE1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cluster using d e l e t e1 unauthorized response has a 3xx status code
func (o *DeleteClusterUsingDELETE1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster using d e l e t e1 unauthorized response has a 4xx status code
func (o *DeleteClusterUsingDELETE1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cluster using d e l e t e1 unauthorized response has a 5xx status code
func (o *DeleteClusterUsingDELETE1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster using d e l e t e1 unauthorized response a status code equal to that given
func (o *DeleteClusterUsingDELETE1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete cluster using d e l e t e1 unauthorized response
func (o *DeleteClusterUsingDELETE1Unauthorized) Code() int {
	return 401
}

func (o *DeleteClusterUsingDELETE1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/clusters/{clusterId}][%d] deleteClusterUsingDELETE1Unauthorized", 401)
}

func (o *DeleteClusterUsingDELETE1Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/clusters/{clusterId}][%d] deleteClusterUsingDELETE1Unauthorized", 401)
}

func (o *DeleteClusterUsingDELETE1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterUsingDELETE1Forbidden creates a DeleteClusterUsingDELETE1Forbidden with default headers values
func NewDeleteClusterUsingDELETE1Forbidden() *DeleteClusterUsingDELETE1Forbidden {
	return &DeleteClusterUsingDELETE1Forbidden{}
}

/*
DeleteClusterUsingDELETE1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteClusterUsingDELETE1Forbidden struct {
}

// IsSuccess returns true when this delete cluster using d e l e t e1 forbidden response has a 2xx status code
func (o *DeleteClusterUsingDELETE1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete cluster using d e l e t e1 forbidden response has a 3xx status code
func (o *DeleteClusterUsingDELETE1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster using d e l e t e1 forbidden response has a 4xx status code
func (o *DeleteClusterUsingDELETE1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete cluster using d e l e t e1 forbidden response has a 5xx status code
func (o *DeleteClusterUsingDELETE1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster using d e l e t e1 forbidden response a status code equal to that given
func (o *DeleteClusterUsingDELETE1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete cluster using d e l e t e1 forbidden response
func (o *DeleteClusterUsingDELETE1Forbidden) Code() int {
	return 403
}

func (o *DeleteClusterUsingDELETE1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/clusters/{clusterId}][%d] deleteClusterUsingDELETE1Forbidden", 403)
}

func (o *DeleteClusterUsingDELETE1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/clusters/{clusterId}][%d] deleteClusterUsingDELETE1Forbidden", 403)
}

func (o *DeleteClusterUsingDELETE1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
