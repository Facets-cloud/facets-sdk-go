// Code generated by go-swagger; DO NOT EDIT.

package ui_artifacts_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteArtifactUsingDELETEReader is a Reader for the DeleteArtifactUsingDELETE structure.
type DeleteArtifactUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArtifactUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteArtifactUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteArtifactUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteArtifactUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteArtifactUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cc-ui/v1/artifacts/{artifactId}] deleteArtifactUsingDELETE", response, response.Code())
	}
}

// NewDeleteArtifactUsingDELETEOK creates a DeleteArtifactUsingDELETEOK with default headers values
func NewDeleteArtifactUsingDELETEOK() *DeleteArtifactUsingDELETEOK {
	return &DeleteArtifactUsingDELETEOK{}
}

/*
DeleteArtifactUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type DeleteArtifactUsingDELETEOK struct {
}

// IsSuccess returns true when this delete artifact using d e l e t e o k response has a 2xx status code
func (o *DeleteArtifactUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete artifact using d e l e t e o k response has a 3xx status code
func (o *DeleteArtifactUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact using d e l e t e o k response has a 4xx status code
func (o *DeleteArtifactUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete artifact using d e l e t e o k response has a 5xx status code
func (o *DeleteArtifactUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact using d e l e t e o k response a status code equal to that given
func (o *DeleteArtifactUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete artifact using d e l e t e o k response
func (o *DeleteArtifactUsingDELETEOK) Code() int {
	return 200
}

func (o *DeleteArtifactUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifacts/{artifactId}][%d] deleteArtifactUsingDELETEOK", 200)
}

func (o *DeleteArtifactUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifacts/{artifactId}][%d] deleteArtifactUsingDELETEOK", 200)
}

func (o *DeleteArtifactUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArtifactUsingDELETENoContent creates a DeleteArtifactUsingDELETENoContent with default headers values
func NewDeleteArtifactUsingDELETENoContent() *DeleteArtifactUsingDELETENoContent {
	return &DeleteArtifactUsingDELETENoContent{}
}

/*
DeleteArtifactUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteArtifactUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete artifact using d e l e t e no content response has a 2xx status code
func (o *DeleteArtifactUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete artifact using d e l e t e no content response has a 3xx status code
func (o *DeleteArtifactUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact using d e l e t e no content response has a 4xx status code
func (o *DeleteArtifactUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete artifact using d e l e t e no content response has a 5xx status code
func (o *DeleteArtifactUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact using d e l e t e no content response a status code equal to that given
func (o *DeleteArtifactUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete artifact using d e l e t e no content response
func (o *DeleteArtifactUsingDELETENoContent) Code() int {
	return 204
}

func (o *DeleteArtifactUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifacts/{artifactId}][%d] deleteArtifactUsingDELETENoContent", 204)
}

func (o *DeleteArtifactUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifacts/{artifactId}][%d] deleteArtifactUsingDELETENoContent", 204)
}

func (o *DeleteArtifactUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArtifactUsingDELETEUnauthorized creates a DeleteArtifactUsingDELETEUnauthorized with default headers values
func NewDeleteArtifactUsingDELETEUnauthorized() *DeleteArtifactUsingDELETEUnauthorized {
	return &DeleteArtifactUsingDELETEUnauthorized{}
}

/*
DeleteArtifactUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteArtifactUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete artifact using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteArtifactUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteArtifactUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteArtifactUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteArtifactUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteArtifactUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete artifact using d e l e t e unauthorized response
func (o *DeleteArtifactUsingDELETEUnauthorized) Code() int {
	return 401
}

func (o *DeleteArtifactUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifacts/{artifactId}][%d] deleteArtifactUsingDELETEUnauthorized", 401)
}

func (o *DeleteArtifactUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifacts/{artifactId}][%d] deleteArtifactUsingDELETEUnauthorized", 401)
}

func (o *DeleteArtifactUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArtifactUsingDELETEForbidden creates a DeleteArtifactUsingDELETEForbidden with default headers values
func NewDeleteArtifactUsingDELETEForbidden() *DeleteArtifactUsingDELETEForbidden {
	return &DeleteArtifactUsingDELETEForbidden{}
}

/*
DeleteArtifactUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteArtifactUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete artifact using d e l e t e forbidden response has a 2xx status code
func (o *DeleteArtifactUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact using d e l e t e forbidden response has a 3xx status code
func (o *DeleteArtifactUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact using d e l e t e forbidden response has a 4xx status code
func (o *DeleteArtifactUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact using d e l e t e forbidden response has a 5xx status code
func (o *DeleteArtifactUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact using d e l e t e forbidden response a status code equal to that given
func (o *DeleteArtifactUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete artifact using d e l e t e forbidden response
func (o *DeleteArtifactUsingDELETEForbidden) Code() int {
	return 403
}

func (o *DeleteArtifactUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifacts/{artifactId}][%d] deleteArtifactUsingDELETEForbidden", 403)
}

func (o *DeleteArtifactUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/artifacts/{artifactId}][%d] deleteArtifactUsingDELETEForbidden", 403)
}

func (o *DeleteArtifactUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
