// Code generated by go-swagger; DO NOT EDIT.

package ui_blueprint_designer_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteResourcesUsingDELETEReader is a Reader for the DeleteResourcesUsingDELETE structure.
type DeleteResourcesUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResourcesUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteResourcesUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteResourcesUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteResourcesUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteResourcesUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}] deleteResourcesUsingDELETE", response, response.Code())
	}
}

// NewDeleteResourcesUsingDELETEOK creates a DeleteResourcesUsingDELETEOK with default headers values
func NewDeleteResourcesUsingDELETEOK() *DeleteResourcesUsingDELETEOK {
	return &DeleteResourcesUsingDELETEOK{}
}

/*
DeleteResourcesUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type DeleteResourcesUsingDELETEOK struct {
}

// IsSuccess returns true when this delete resources using d e l e t e o k response has a 2xx status code
func (o *DeleteResourcesUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete resources using d e l e t e o k response has a 3xx status code
func (o *DeleteResourcesUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resources using d e l e t e o k response has a 4xx status code
func (o *DeleteResourcesUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete resources using d e l e t e o k response has a 5xx status code
func (o *DeleteResourcesUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resources using d e l e t e o k response a status code equal to that given
func (o *DeleteResourcesUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete resources using d e l e t e o k response
func (o *DeleteResourcesUsingDELETEOK) Code() int {
	return 200
}

func (o *DeleteResourcesUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}][%d] deleteResourcesUsingDELETEOK", 200)
}

func (o *DeleteResourcesUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}][%d] deleteResourcesUsingDELETEOK", 200)
}

func (o *DeleteResourcesUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourcesUsingDELETENoContent creates a DeleteResourcesUsingDELETENoContent with default headers values
func NewDeleteResourcesUsingDELETENoContent() *DeleteResourcesUsingDELETENoContent {
	return &DeleteResourcesUsingDELETENoContent{}
}

/*
DeleteResourcesUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteResourcesUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete resources using d e l e t e no content response has a 2xx status code
func (o *DeleteResourcesUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete resources using d e l e t e no content response has a 3xx status code
func (o *DeleteResourcesUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resources using d e l e t e no content response has a 4xx status code
func (o *DeleteResourcesUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete resources using d e l e t e no content response has a 5xx status code
func (o *DeleteResourcesUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resources using d e l e t e no content response a status code equal to that given
func (o *DeleteResourcesUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete resources using d e l e t e no content response
func (o *DeleteResourcesUsingDELETENoContent) Code() int {
	return 204
}

func (o *DeleteResourcesUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}][%d] deleteResourcesUsingDELETENoContent", 204)
}

func (o *DeleteResourcesUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}][%d] deleteResourcesUsingDELETENoContent", 204)
}

func (o *DeleteResourcesUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourcesUsingDELETEUnauthorized creates a DeleteResourcesUsingDELETEUnauthorized with default headers values
func NewDeleteResourcesUsingDELETEUnauthorized() *DeleteResourcesUsingDELETEUnauthorized {
	return &DeleteResourcesUsingDELETEUnauthorized{}
}

/*
DeleteResourcesUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteResourcesUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete resources using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteResourcesUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resources using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteResourcesUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resources using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteResourcesUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resources using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteResourcesUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resources using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteResourcesUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete resources using d e l e t e unauthorized response
func (o *DeleteResourcesUsingDELETEUnauthorized) Code() int {
	return 401
}

func (o *DeleteResourcesUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}][%d] deleteResourcesUsingDELETEUnauthorized", 401)
}

func (o *DeleteResourcesUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}][%d] deleteResourcesUsingDELETEUnauthorized", 401)
}

func (o *DeleteResourcesUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourcesUsingDELETEForbidden creates a DeleteResourcesUsingDELETEForbidden with default headers values
func NewDeleteResourcesUsingDELETEForbidden() *DeleteResourcesUsingDELETEForbidden {
	return &DeleteResourcesUsingDELETEForbidden{}
}

/*
DeleteResourcesUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteResourcesUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete resources using d e l e t e forbidden response has a 2xx status code
func (o *DeleteResourcesUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resources using d e l e t e forbidden response has a 3xx status code
func (o *DeleteResourcesUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resources using d e l e t e forbidden response has a 4xx status code
func (o *DeleteResourcesUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resources using d e l e t e forbidden response has a 5xx status code
func (o *DeleteResourcesUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resources using d e l e t e forbidden response a status code equal to that given
func (o *DeleteResourcesUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete resources using d e l e t e forbidden response
func (o *DeleteResourcesUsingDELETEForbidden) Code() int {
	return 403
}

func (o *DeleteResourcesUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}][%d] deleteResourcesUsingDELETEForbidden", 403)
}

func (o *DeleteResourcesUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/designer/{stackName}/branch/{branch}][%d] deleteResourcesUsingDELETEForbidden", 403)
}

func (o *DeleteResourcesUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
