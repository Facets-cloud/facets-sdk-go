// Code generated by go-swagger; DO NOT EDIT.

package ui_dropdowns_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SyncSubstackGitHistoryUsingPOSTReader is a Reader for the SyncSubstackGitHistoryUsingPOST structure.
type SyncSubstackGitHistoryUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncSubstackGitHistoryUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncSubstackGitHistoryUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewSyncSubstackGitHistoryUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSyncSubstackGitHistoryUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSyncSubstackGitHistoryUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSyncSubstackGitHistoryUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/dropdown/logs/substack] syncSubstackGitHistoryUsingPOST", response, response.Code())
	}
}

// NewSyncSubstackGitHistoryUsingPOSTOK creates a SyncSubstackGitHistoryUsingPOSTOK with default headers values
func NewSyncSubstackGitHistoryUsingPOSTOK() *SyncSubstackGitHistoryUsingPOSTOK {
	return &SyncSubstackGitHistoryUsingPOSTOK{}
}

/*
SyncSubstackGitHistoryUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type SyncSubstackGitHistoryUsingPOSTOK struct {
}

// IsSuccess returns true when this sync substack git history using p o s t o k response has a 2xx status code
func (o *SyncSubstackGitHistoryUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync substack git history using p o s t o k response has a 3xx status code
func (o *SyncSubstackGitHistoryUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync substack git history using p o s t o k response has a 4xx status code
func (o *SyncSubstackGitHistoryUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync substack git history using p o s t o k response has a 5xx status code
func (o *SyncSubstackGitHistoryUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync substack git history using p o s t o k response a status code equal to that given
func (o *SyncSubstackGitHistoryUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sync substack git history using p o s t o k response
func (o *SyncSubstackGitHistoryUsingPOSTOK) Code() int {
	return 200
}

func (o *SyncSubstackGitHistoryUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTOK", 200)
}

func (o *SyncSubstackGitHistoryUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTOK", 200)
}

func (o *SyncSubstackGitHistoryUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncSubstackGitHistoryUsingPOSTCreated creates a SyncSubstackGitHistoryUsingPOSTCreated with default headers values
func NewSyncSubstackGitHistoryUsingPOSTCreated() *SyncSubstackGitHistoryUsingPOSTCreated {
	return &SyncSubstackGitHistoryUsingPOSTCreated{}
}

/*
SyncSubstackGitHistoryUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type SyncSubstackGitHistoryUsingPOSTCreated struct {
}

// IsSuccess returns true when this sync substack git history using p o s t created response has a 2xx status code
func (o *SyncSubstackGitHistoryUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync substack git history using p o s t created response has a 3xx status code
func (o *SyncSubstackGitHistoryUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync substack git history using p o s t created response has a 4xx status code
func (o *SyncSubstackGitHistoryUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync substack git history using p o s t created response has a 5xx status code
func (o *SyncSubstackGitHistoryUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this sync substack git history using p o s t created response a status code equal to that given
func (o *SyncSubstackGitHistoryUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the sync substack git history using p o s t created response
func (o *SyncSubstackGitHistoryUsingPOSTCreated) Code() int {
	return 201
}

func (o *SyncSubstackGitHistoryUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTCreated", 201)
}

func (o *SyncSubstackGitHistoryUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTCreated", 201)
}

func (o *SyncSubstackGitHistoryUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncSubstackGitHistoryUsingPOSTUnauthorized creates a SyncSubstackGitHistoryUsingPOSTUnauthorized with default headers values
func NewSyncSubstackGitHistoryUsingPOSTUnauthorized() *SyncSubstackGitHistoryUsingPOSTUnauthorized {
	return &SyncSubstackGitHistoryUsingPOSTUnauthorized{}
}

/*
SyncSubstackGitHistoryUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SyncSubstackGitHistoryUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this sync substack git history using p o s t unauthorized response has a 2xx status code
func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync substack git history using p o s t unauthorized response has a 3xx status code
func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync substack git history using p o s t unauthorized response has a 4xx status code
func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync substack git history using p o s t unauthorized response has a 5xx status code
func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this sync substack git history using p o s t unauthorized response a status code equal to that given
func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the sync substack git history using p o s t unauthorized response
func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTUnauthorized", 401)
}

func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTUnauthorized", 401)
}

func (o *SyncSubstackGitHistoryUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncSubstackGitHistoryUsingPOSTForbidden creates a SyncSubstackGitHistoryUsingPOSTForbidden with default headers values
func NewSyncSubstackGitHistoryUsingPOSTForbidden() *SyncSubstackGitHistoryUsingPOSTForbidden {
	return &SyncSubstackGitHistoryUsingPOSTForbidden{}
}

/*
SyncSubstackGitHistoryUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SyncSubstackGitHistoryUsingPOSTForbidden struct {
}

// IsSuccess returns true when this sync substack git history using p o s t forbidden response has a 2xx status code
func (o *SyncSubstackGitHistoryUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync substack git history using p o s t forbidden response has a 3xx status code
func (o *SyncSubstackGitHistoryUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync substack git history using p o s t forbidden response has a 4xx status code
func (o *SyncSubstackGitHistoryUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync substack git history using p o s t forbidden response has a 5xx status code
func (o *SyncSubstackGitHistoryUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this sync substack git history using p o s t forbidden response a status code equal to that given
func (o *SyncSubstackGitHistoryUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the sync substack git history using p o s t forbidden response
func (o *SyncSubstackGitHistoryUsingPOSTForbidden) Code() int {
	return 403
}

func (o *SyncSubstackGitHistoryUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTForbidden", 403)
}

func (o *SyncSubstackGitHistoryUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTForbidden", 403)
}

func (o *SyncSubstackGitHistoryUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncSubstackGitHistoryUsingPOSTNotFound creates a SyncSubstackGitHistoryUsingPOSTNotFound with default headers values
func NewSyncSubstackGitHistoryUsingPOSTNotFound() *SyncSubstackGitHistoryUsingPOSTNotFound {
	return &SyncSubstackGitHistoryUsingPOSTNotFound{}
}

/*
SyncSubstackGitHistoryUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SyncSubstackGitHistoryUsingPOSTNotFound struct {
}

// IsSuccess returns true when this sync substack git history using p o s t not found response has a 2xx status code
func (o *SyncSubstackGitHistoryUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync substack git history using p o s t not found response has a 3xx status code
func (o *SyncSubstackGitHistoryUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync substack git history using p o s t not found response has a 4xx status code
func (o *SyncSubstackGitHistoryUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync substack git history using p o s t not found response has a 5xx status code
func (o *SyncSubstackGitHistoryUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this sync substack git history using p o s t not found response a status code equal to that given
func (o *SyncSubstackGitHistoryUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the sync substack git history using p o s t not found response
func (o *SyncSubstackGitHistoryUsingPOSTNotFound) Code() int {
	return 404
}

func (o *SyncSubstackGitHistoryUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTNotFound", 404)
}

func (o *SyncSubstackGitHistoryUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/dropdown/logs/substack][%d] syncSubstackGitHistoryUsingPOSTNotFound", 404)
}

func (o *SyncSubstackGitHistoryUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
