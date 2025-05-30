// Code generated by go-swagger; DO NOT EDIT.

package ui_accounts_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetVCSOrganisationsUsingPOSTReader is a Reader for the GetVCSOrganisationsUsingPOST structure.
type GetVCSOrganisationsUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVCSOrganisationsUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVCSOrganisationsUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewGetVCSOrganisationsUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVCSOrganisationsUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVCSOrganisationsUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVCSOrganisationsUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cc-ui/v1/accounts/get-organisations] getVCSOrganisationsUsingPOST", response, response.Code())
	}
}

// NewGetVCSOrganisationsUsingPOSTOK creates a GetVCSOrganisationsUsingPOSTOK with default headers values
func NewGetVCSOrganisationsUsingPOSTOK() *GetVCSOrganisationsUsingPOSTOK {
	return &GetVCSOrganisationsUsingPOSTOK{}
}

/*
GetVCSOrganisationsUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type GetVCSOrganisationsUsingPOSTOK struct {
	Payload []string
}

// IsSuccess returns true when this get v c s organisations using p o s t o k response has a 2xx status code
func (o *GetVCSOrganisationsUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v c s organisations using p o s t o k response has a 3xx status code
func (o *GetVCSOrganisationsUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s organisations using p o s t o k response has a 4xx status code
func (o *GetVCSOrganisationsUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v c s organisations using p o s t o k response has a 5xx status code
func (o *GetVCSOrganisationsUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s organisations using p o s t o k response a status code equal to that given
func (o *GetVCSOrganisationsUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v c s organisations using p o s t o k response
func (o *GetVCSOrganisationsUsingPOSTOK) Code() int {
	return 200
}

func (o *GetVCSOrganisationsUsingPOSTOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTOK %s", 200, payload)
}

func (o *GetVCSOrganisationsUsingPOSTOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTOK %s", 200, payload)
}

func (o *GetVCSOrganisationsUsingPOSTOK) GetPayload() []string {
	return o.Payload
}

func (o *GetVCSOrganisationsUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVCSOrganisationsUsingPOSTCreated creates a GetVCSOrganisationsUsingPOSTCreated with default headers values
func NewGetVCSOrganisationsUsingPOSTCreated() *GetVCSOrganisationsUsingPOSTCreated {
	return &GetVCSOrganisationsUsingPOSTCreated{}
}

/*
GetVCSOrganisationsUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type GetVCSOrganisationsUsingPOSTCreated struct {
}

// IsSuccess returns true when this get v c s organisations using p o s t created response has a 2xx status code
func (o *GetVCSOrganisationsUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v c s organisations using p o s t created response has a 3xx status code
func (o *GetVCSOrganisationsUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s organisations using p o s t created response has a 4xx status code
func (o *GetVCSOrganisationsUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v c s organisations using p o s t created response has a 5xx status code
func (o *GetVCSOrganisationsUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s organisations using p o s t created response a status code equal to that given
func (o *GetVCSOrganisationsUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the get v c s organisations using p o s t created response
func (o *GetVCSOrganisationsUsingPOSTCreated) Code() int {
	return 201
}

func (o *GetVCSOrganisationsUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTCreated", 201)
}

func (o *GetVCSOrganisationsUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTCreated", 201)
}

func (o *GetVCSOrganisationsUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVCSOrganisationsUsingPOSTUnauthorized creates a GetVCSOrganisationsUsingPOSTUnauthorized with default headers values
func NewGetVCSOrganisationsUsingPOSTUnauthorized() *GetVCSOrganisationsUsingPOSTUnauthorized {
	return &GetVCSOrganisationsUsingPOSTUnauthorized{}
}

/*
GetVCSOrganisationsUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVCSOrganisationsUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this get v c s organisations using p o s t unauthorized response has a 2xx status code
func (o *GetVCSOrganisationsUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v c s organisations using p o s t unauthorized response has a 3xx status code
func (o *GetVCSOrganisationsUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s organisations using p o s t unauthorized response has a 4xx status code
func (o *GetVCSOrganisationsUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v c s organisations using p o s t unauthorized response has a 5xx status code
func (o *GetVCSOrganisationsUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s organisations using p o s t unauthorized response a status code equal to that given
func (o *GetVCSOrganisationsUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v c s organisations using p o s t unauthorized response
func (o *GetVCSOrganisationsUsingPOSTUnauthorized) Code() int {
	return 401
}

func (o *GetVCSOrganisationsUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTUnauthorized", 401)
}

func (o *GetVCSOrganisationsUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTUnauthorized", 401)
}

func (o *GetVCSOrganisationsUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVCSOrganisationsUsingPOSTForbidden creates a GetVCSOrganisationsUsingPOSTForbidden with default headers values
func NewGetVCSOrganisationsUsingPOSTForbidden() *GetVCSOrganisationsUsingPOSTForbidden {
	return &GetVCSOrganisationsUsingPOSTForbidden{}
}

/*
GetVCSOrganisationsUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVCSOrganisationsUsingPOSTForbidden struct {
}

// IsSuccess returns true when this get v c s organisations using p o s t forbidden response has a 2xx status code
func (o *GetVCSOrganisationsUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v c s organisations using p o s t forbidden response has a 3xx status code
func (o *GetVCSOrganisationsUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s organisations using p o s t forbidden response has a 4xx status code
func (o *GetVCSOrganisationsUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v c s organisations using p o s t forbidden response has a 5xx status code
func (o *GetVCSOrganisationsUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s organisations using p o s t forbidden response a status code equal to that given
func (o *GetVCSOrganisationsUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v c s organisations using p o s t forbidden response
func (o *GetVCSOrganisationsUsingPOSTForbidden) Code() int {
	return 403
}

func (o *GetVCSOrganisationsUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTForbidden", 403)
}

func (o *GetVCSOrganisationsUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTForbidden", 403)
}

func (o *GetVCSOrganisationsUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVCSOrganisationsUsingPOSTNotFound creates a GetVCSOrganisationsUsingPOSTNotFound with default headers values
func NewGetVCSOrganisationsUsingPOSTNotFound() *GetVCSOrganisationsUsingPOSTNotFound {
	return &GetVCSOrganisationsUsingPOSTNotFound{}
}

/*
GetVCSOrganisationsUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVCSOrganisationsUsingPOSTNotFound struct {
}

// IsSuccess returns true when this get v c s organisations using p o s t not found response has a 2xx status code
func (o *GetVCSOrganisationsUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v c s organisations using p o s t not found response has a 3xx status code
func (o *GetVCSOrganisationsUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v c s organisations using p o s t not found response has a 4xx status code
func (o *GetVCSOrganisationsUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v c s organisations using p o s t not found response has a 5xx status code
func (o *GetVCSOrganisationsUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v c s organisations using p o s t not found response a status code equal to that given
func (o *GetVCSOrganisationsUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v c s organisations using p o s t not found response
func (o *GetVCSOrganisationsUsingPOSTNotFound) Code() int {
	return 404
}

func (o *GetVCSOrganisationsUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTNotFound", 404)
}

func (o *GetVCSOrganisationsUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /cc-ui/v1/accounts/get-organisations][%d] getVCSOrganisationsUsingPOSTNotFound", 404)
}

func (o *GetVCSOrganisationsUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
