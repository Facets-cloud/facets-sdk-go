// Code generated by go-swagger; DO NOT EDIT.

package ui_resource_group_controller

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

// DeleteUsingDELETE1Reader is a Reader for the DeleteUsingDELETE1 structure.
type DeleteUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUsingDELETE1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteUsingDELETE1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUsingDELETE1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsingDELETE1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}] deleteUsingDELETE_1", response, response.Code())
	}
}

// NewDeleteUsingDELETE1OK creates a DeleteUsingDELETE1OK with default headers values
func NewDeleteUsingDELETE1OK() *DeleteUsingDELETE1OK {
	return &DeleteUsingDELETE1OK{}
}

/*
DeleteUsingDELETE1OK describes a response with status code 200, with default header values.

OK
*/
type DeleteUsingDELETE1OK struct {
	Payload *models.Response
}

// IsSuccess returns true when this delete using d e l e t e1 o k response has a 2xx status code
func (o *DeleteUsingDELETE1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete using d e l e t e1 o k response has a 3xx status code
func (o *DeleteUsingDELETE1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete using d e l e t e1 o k response has a 4xx status code
func (o *DeleteUsingDELETE1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete using d e l e t e1 o k response has a 5xx status code
func (o *DeleteUsingDELETE1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete using d e l e t e1 o k response a status code equal to that given
func (o *DeleteUsingDELETE1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete using d e l e t e1 o k response
func (o *DeleteUsingDELETE1OK) Code() int {
	return 200
}

func (o *DeleteUsingDELETE1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}][%d] deleteUsingDELETE1OK %s", 200, payload)
}

func (o *DeleteUsingDELETE1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}][%d] deleteUsingDELETE1OK %s", 200, payload)
}

func (o *DeleteUsingDELETE1OK) GetPayload() *models.Response {
	return o.Payload
}

func (o *DeleteUsingDELETE1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUsingDELETE1NoContent creates a DeleteUsingDELETE1NoContent with default headers values
func NewDeleteUsingDELETE1NoContent() *DeleteUsingDELETE1NoContent {
	return &DeleteUsingDELETE1NoContent{}
}

/*
DeleteUsingDELETE1NoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteUsingDELETE1NoContent struct {
}

// IsSuccess returns true when this delete using d e l e t e1 no content response has a 2xx status code
func (o *DeleteUsingDELETE1NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete using d e l e t e1 no content response has a 3xx status code
func (o *DeleteUsingDELETE1NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete using d e l e t e1 no content response has a 4xx status code
func (o *DeleteUsingDELETE1NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete using d e l e t e1 no content response has a 5xx status code
func (o *DeleteUsingDELETE1NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete using d e l e t e1 no content response a status code equal to that given
func (o *DeleteUsingDELETE1NoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete using d e l e t e1 no content response
func (o *DeleteUsingDELETE1NoContent) Code() int {
	return 204
}

func (o *DeleteUsingDELETE1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}][%d] deleteUsingDELETE1NoContent", 204)
}

func (o *DeleteUsingDELETE1NoContent) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}][%d] deleteUsingDELETE1NoContent", 204)
}

func (o *DeleteUsingDELETE1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE1Unauthorized creates a DeleteUsingDELETE1Unauthorized with default headers values
func NewDeleteUsingDELETE1Unauthorized() *DeleteUsingDELETE1Unauthorized {
	return &DeleteUsingDELETE1Unauthorized{}
}

/*
DeleteUsingDELETE1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteUsingDELETE1Unauthorized struct {
}

// IsSuccess returns true when this delete using d e l e t e1 unauthorized response has a 2xx status code
func (o *DeleteUsingDELETE1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete using d e l e t e1 unauthorized response has a 3xx status code
func (o *DeleteUsingDELETE1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete using d e l e t e1 unauthorized response has a 4xx status code
func (o *DeleteUsingDELETE1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete using d e l e t e1 unauthorized response has a 5xx status code
func (o *DeleteUsingDELETE1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete using d e l e t e1 unauthorized response a status code equal to that given
func (o *DeleteUsingDELETE1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete using d e l e t e1 unauthorized response
func (o *DeleteUsingDELETE1Unauthorized) Code() int {
	return 401
}

func (o *DeleteUsingDELETE1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}][%d] deleteUsingDELETE1Unauthorized", 401)
}

func (o *DeleteUsingDELETE1Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}][%d] deleteUsingDELETE1Unauthorized", 401)
}

func (o *DeleteUsingDELETE1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE1Forbidden creates a DeleteUsingDELETE1Forbidden with default headers values
func NewDeleteUsingDELETE1Forbidden() *DeleteUsingDELETE1Forbidden {
	return &DeleteUsingDELETE1Forbidden{}
}

/*
DeleteUsingDELETE1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteUsingDELETE1Forbidden struct {
}

// IsSuccess returns true when this delete using d e l e t e1 forbidden response has a 2xx status code
func (o *DeleteUsingDELETE1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete using d e l e t e1 forbidden response has a 3xx status code
func (o *DeleteUsingDELETE1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete using d e l e t e1 forbidden response has a 4xx status code
func (o *DeleteUsingDELETE1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete using d e l e t e1 forbidden response has a 5xx status code
func (o *DeleteUsingDELETE1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete using d e l e t e1 forbidden response a status code equal to that given
func (o *DeleteUsingDELETE1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete using d e l e t e1 forbidden response
func (o *DeleteUsingDELETE1Forbidden) Code() int {
	return 403
}

func (o *DeleteUsingDELETE1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}][%d] deleteUsingDELETE1Forbidden", 403)
}

func (o *DeleteUsingDELETE1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/resource-groups/{resourceGroupId}][%d] deleteUsingDELETE1Forbidden", 403)
}

func (o *DeleteUsingDELETE1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
