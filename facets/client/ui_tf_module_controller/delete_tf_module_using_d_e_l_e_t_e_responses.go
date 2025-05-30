// Code generated by go-swagger; DO NOT EDIT.

package ui_tf_module_controller

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

// DeleteTfModuleUsingDELETEReader is a Reader for the DeleteTfModuleUsingDELETE structure.
type DeleteTfModuleUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTfModuleUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTfModuleUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTfModuleUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTfModuleUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTfModuleUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTfModuleUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTfModuleUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTfModuleUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cc-ui/v1/modules/{id}] deleteTfModuleUsingDELETE", response, response.Code())
	}
}

// NewDeleteTfModuleUsingDELETEOK creates a DeleteTfModuleUsingDELETEOK with default headers values
func NewDeleteTfModuleUsingDELETEOK() *DeleteTfModuleUsingDELETEOK {
	return &DeleteTfModuleUsingDELETEOK{}
}

/*
DeleteTfModuleUsingDELETEOK describes a response with status code 200, with default header values.

Operation successful
*/
type DeleteTfModuleUsingDELETEOK struct {
	Payload *models.TFModule
}

// IsSuccess returns true when this delete tf module using d e l e t e o k response has a 2xx status code
func (o *DeleteTfModuleUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete tf module using d e l e t e o k response has a 3xx status code
func (o *DeleteTfModuleUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tf module using d e l e t e o k response has a 4xx status code
func (o *DeleteTfModuleUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tf module using d e l e t e o k response has a 5xx status code
func (o *DeleteTfModuleUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tf module using d e l e t e o k response a status code equal to that given
func (o *DeleteTfModuleUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete tf module using d e l e t e o k response
func (o *DeleteTfModuleUsingDELETEOK) Code() int {
	return 200
}

func (o *DeleteTfModuleUsingDELETEOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEOK %s", 200, payload)
}

func (o *DeleteTfModuleUsingDELETEOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEOK %s", 200, payload)
}

func (o *DeleteTfModuleUsingDELETEOK) GetPayload() *models.TFModule {
	return o.Payload
}

func (o *DeleteTfModuleUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TFModule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTfModuleUsingDELETENoContent creates a DeleteTfModuleUsingDELETENoContent with default headers values
func NewDeleteTfModuleUsingDELETENoContent() *DeleteTfModuleUsingDELETENoContent {
	return &DeleteTfModuleUsingDELETENoContent{}
}

/*
DeleteTfModuleUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteTfModuleUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete tf module using d e l e t e no content response has a 2xx status code
func (o *DeleteTfModuleUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete tf module using d e l e t e no content response has a 3xx status code
func (o *DeleteTfModuleUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tf module using d e l e t e no content response has a 4xx status code
func (o *DeleteTfModuleUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tf module using d e l e t e no content response has a 5xx status code
func (o *DeleteTfModuleUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tf module using d e l e t e no content response a status code equal to that given
func (o *DeleteTfModuleUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete tf module using d e l e t e no content response
func (o *DeleteTfModuleUsingDELETENoContent) Code() int {
	return 204
}

func (o *DeleteTfModuleUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETENoContent", 204)
}

func (o *DeleteTfModuleUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETENoContent", 204)
}

func (o *DeleteTfModuleUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTfModuleUsingDELETEBadRequest creates a DeleteTfModuleUsingDELETEBadRequest with default headers values
func NewDeleteTfModuleUsingDELETEBadRequest() *DeleteTfModuleUsingDELETEBadRequest {
	return &DeleteTfModuleUsingDELETEBadRequest{}
}

/*
DeleteTfModuleUsingDELETEBadRequest describes a response with status code 400, with default header values.

Invalid input parameters
*/
type DeleteTfModuleUsingDELETEBadRequest struct {
}

// IsSuccess returns true when this delete tf module using d e l e t e bad request response has a 2xx status code
func (o *DeleteTfModuleUsingDELETEBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tf module using d e l e t e bad request response has a 3xx status code
func (o *DeleteTfModuleUsingDELETEBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tf module using d e l e t e bad request response has a 4xx status code
func (o *DeleteTfModuleUsingDELETEBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tf module using d e l e t e bad request response has a 5xx status code
func (o *DeleteTfModuleUsingDELETEBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tf module using d e l e t e bad request response a status code equal to that given
func (o *DeleteTfModuleUsingDELETEBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete tf module using d e l e t e bad request response
func (o *DeleteTfModuleUsingDELETEBadRequest) Code() int {
	return 400
}

func (o *DeleteTfModuleUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEBadRequest", 400)
}

func (o *DeleteTfModuleUsingDELETEBadRequest) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEBadRequest", 400)
}

func (o *DeleteTfModuleUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTfModuleUsingDELETEUnauthorized creates a DeleteTfModuleUsingDELETEUnauthorized with default headers values
func NewDeleteTfModuleUsingDELETEUnauthorized() *DeleteTfModuleUsingDELETEUnauthorized {
	return &DeleteTfModuleUsingDELETEUnauthorized{}
}

/*
DeleteTfModuleUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteTfModuleUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete tf module using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteTfModuleUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tf module using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteTfModuleUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tf module using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteTfModuleUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tf module using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteTfModuleUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tf module using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteTfModuleUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete tf module using d e l e t e unauthorized response
func (o *DeleteTfModuleUsingDELETEUnauthorized) Code() int {
	return 401
}

func (o *DeleteTfModuleUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEUnauthorized", 401)
}

func (o *DeleteTfModuleUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEUnauthorized", 401)
}

func (o *DeleteTfModuleUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTfModuleUsingDELETEForbidden creates a DeleteTfModuleUsingDELETEForbidden with default headers values
func NewDeleteTfModuleUsingDELETEForbidden() *DeleteTfModuleUsingDELETEForbidden {
	return &DeleteTfModuleUsingDELETEForbidden{}
}

/*
DeleteTfModuleUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTfModuleUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete tf module using d e l e t e forbidden response has a 2xx status code
func (o *DeleteTfModuleUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tf module using d e l e t e forbidden response has a 3xx status code
func (o *DeleteTfModuleUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tf module using d e l e t e forbidden response has a 4xx status code
func (o *DeleteTfModuleUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tf module using d e l e t e forbidden response has a 5xx status code
func (o *DeleteTfModuleUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tf module using d e l e t e forbidden response a status code equal to that given
func (o *DeleteTfModuleUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete tf module using d e l e t e forbidden response
func (o *DeleteTfModuleUsingDELETEForbidden) Code() int {
	return 403
}

func (o *DeleteTfModuleUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEForbidden", 403)
}

func (o *DeleteTfModuleUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEForbidden", 403)
}

func (o *DeleteTfModuleUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTfModuleUsingDELETENotFound creates a DeleteTfModuleUsingDELETENotFound with default headers values
func NewDeleteTfModuleUsingDELETENotFound() *DeleteTfModuleUsingDELETENotFound {
	return &DeleteTfModuleUsingDELETENotFound{}
}

/*
DeleteTfModuleUsingDELETENotFound describes a response with status code 404, with default header values.

Entity not found
*/
type DeleteTfModuleUsingDELETENotFound struct {
}

// IsSuccess returns true when this delete tf module using d e l e t e not found response has a 2xx status code
func (o *DeleteTfModuleUsingDELETENotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tf module using d e l e t e not found response has a 3xx status code
func (o *DeleteTfModuleUsingDELETENotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tf module using d e l e t e not found response has a 4xx status code
func (o *DeleteTfModuleUsingDELETENotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete tf module using d e l e t e not found response has a 5xx status code
func (o *DeleteTfModuleUsingDELETENotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete tf module using d e l e t e not found response a status code equal to that given
func (o *DeleteTfModuleUsingDELETENotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete tf module using d e l e t e not found response
func (o *DeleteTfModuleUsingDELETENotFound) Code() int {
	return 404
}

func (o *DeleteTfModuleUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETENotFound", 404)
}

func (o *DeleteTfModuleUsingDELETENotFound) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETENotFound", 404)
}

func (o *DeleteTfModuleUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTfModuleUsingDELETEInternalServerError creates a DeleteTfModuleUsingDELETEInternalServerError with default headers values
func NewDeleteTfModuleUsingDELETEInternalServerError() *DeleteTfModuleUsingDELETEInternalServerError {
	return &DeleteTfModuleUsingDELETEInternalServerError{}
}

/*
DeleteTfModuleUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteTfModuleUsingDELETEInternalServerError struct {
}

// IsSuccess returns true when this delete tf module using d e l e t e internal server error response has a 2xx status code
func (o *DeleteTfModuleUsingDELETEInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete tf module using d e l e t e internal server error response has a 3xx status code
func (o *DeleteTfModuleUsingDELETEInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete tf module using d e l e t e internal server error response has a 4xx status code
func (o *DeleteTfModuleUsingDELETEInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete tf module using d e l e t e internal server error response has a 5xx status code
func (o *DeleteTfModuleUsingDELETEInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete tf module using d e l e t e internal server error response a status code equal to that given
func (o *DeleteTfModuleUsingDELETEInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete tf module using d e l e t e internal server error response
func (o *DeleteTfModuleUsingDELETEInternalServerError) Code() int {
	return 500
}

func (o *DeleteTfModuleUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEInternalServerError", 500)
}

func (o *DeleteTfModuleUsingDELETEInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/modules/{id}][%d] deleteTfModuleUsingDELETEInternalServerError", 500)
}

func (o *DeleteTfModuleUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
