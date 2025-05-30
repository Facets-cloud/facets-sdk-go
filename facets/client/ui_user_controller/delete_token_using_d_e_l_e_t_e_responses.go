// Code generated by go-swagger; DO NOT EDIT.

package ui_user_controller

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

// DeleteTokenUsingDELETEReader is a Reader for the DeleteTokenUsingDELETE structure.
type DeleteTokenUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTokenUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTokenUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTokenUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTokenUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTokenUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cc-ui/v1/users/tokens/{tokenId}] deleteTokenUsingDELETE", response, response.Code())
	}
}

// NewDeleteTokenUsingDELETEOK creates a DeleteTokenUsingDELETEOK with default headers values
func NewDeleteTokenUsingDELETEOK() *DeleteTokenUsingDELETEOK {
	return &DeleteTokenUsingDELETEOK{}
}

/*
DeleteTokenUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type DeleteTokenUsingDELETEOK struct {
	Payload []*models.UserAccessToken
}

// IsSuccess returns true when this delete token using d e l e t e o k response has a 2xx status code
func (o *DeleteTokenUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete token using d e l e t e o k response has a 3xx status code
func (o *DeleteTokenUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token using d e l e t e o k response has a 4xx status code
func (o *DeleteTokenUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete token using d e l e t e o k response has a 5xx status code
func (o *DeleteTokenUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token using d e l e t e o k response a status code equal to that given
func (o *DeleteTokenUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete token using d e l e t e o k response
func (o *DeleteTokenUsingDELETEOK) Code() int {
	return 200
}

func (o *DeleteTokenUsingDELETEOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cc-ui/v1/users/tokens/{tokenId}][%d] deleteTokenUsingDELETEOK %s", 200, payload)
}

func (o *DeleteTokenUsingDELETEOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cc-ui/v1/users/tokens/{tokenId}][%d] deleteTokenUsingDELETEOK %s", 200, payload)
}

func (o *DeleteTokenUsingDELETEOK) GetPayload() []*models.UserAccessToken {
	return o.Payload
}

func (o *DeleteTokenUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenUsingDELETENoContent creates a DeleteTokenUsingDELETENoContent with default headers values
func NewDeleteTokenUsingDELETENoContent() *DeleteTokenUsingDELETENoContent {
	return &DeleteTokenUsingDELETENoContent{}
}

/*
DeleteTokenUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteTokenUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete token using d e l e t e no content response has a 2xx status code
func (o *DeleteTokenUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete token using d e l e t e no content response has a 3xx status code
func (o *DeleteTokenUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token using d e l e t e no content response has a 4xx status code
func (o *DeleteTokenUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete token using d e l e t e no content response has a 5xx status code
func (o *DeleteTokenUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token using d e l e t e no content response a status code equal to that given
func (o *DeleteTokenUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete token using d e l e t e no content response
func (o *DeleteTokenUsingDELETENoContent) Code() int {
	return 204
}

func (o *DeleteTokenUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/users/tokens/{tokenId}][%d] deleteTokenUsingDELETENoContent", 204)
}

func (o *DeleteTokenUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/users/tokens/{tokenId}][%d] deleteTokenUsingDELETENoContent", 204)
}

func (o *DeleteTokenUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTokenUsingDELETEUnauthorized creates a DeleteTokenUsingDELETEUnauthorized with default headers values
func NewDeleteTokenUsingDELETEUnauthorized() *DeleteTokenUsingDELETEUnauthorized {
	return &DeleteTokenUsingDELETEUnauthorized{}
}

/*
DeleteTokenUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteTokenUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete token using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteTokenUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete token using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteTokenUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteTokenUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete token using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteTokenUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteTokenUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete token using d e l e t e unauthorized response
func (o *DeleteTokenUsingDELETEUnauthorized) Code() int {
	return 401
}

func (o *DeleteTokenUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/users/tokens/{tokenId}][%d] deleteTokenUsingDELETEUnauthorized", 401)
}

func (o *DeleteTokenUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/users/tokens/{tokenId}][%d] deleteTokenUsingDELETEUnauthorized", 401)
}

func (o *DeleteTokenUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTokenUsingDELETEForbidden creates a DeleteTokenUsingDELETEForbidden with default headers values
func NewDeleteTokenUsingDELETEForbidden() *DeleteTokenUsingDELETEForbidden {
	return &DeleteTokenUsingDELETEForbidden{}
}

/*
DeleteTokenUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTokenUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete token using d e l e t e forbidden response has a 2xx status code
func (o *DeleteTokenUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete token using d e l e t e forbidden response has a 3xx status code
func (o *DeleteTokenUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete token using d e l e t e forbidden response has a 4xx status code
func (o *DeleteTokenUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete token using d e l e t e forbidden response has a 5xx status code
func (o *DeleteTokenUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete token using d e l e t e forbidden response a status code equal to that given
func (o *DeleteTokenUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete token using d e l e t e forbidden response
func (o *DeleteTokenUsingDELETEForbidden) Code() int {
	return 403
}

func (o *DeleteTokenUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/users/tokens/{tokenId}][%d] deleteTokenUsingDELETEForbidden", 403)
}

func (o *DeleteTokenUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /cc-ui/v1/users/tokens/{tokenId}][%d] deleteTokenUsingDELETEForbidden", 403)
}

func (o *DeleteTokenUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
