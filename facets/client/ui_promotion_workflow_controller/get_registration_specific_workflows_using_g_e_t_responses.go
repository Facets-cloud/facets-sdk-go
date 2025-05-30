// Code generated by go-swagger; DO NOT EDIT.

package ui_promotion_workflow_controller

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

// GetRegistrationSpecificWorkflowsUsingGETReader is a Reader for the GetRegistrationSpecificWorkflowsUsingGET structure.
type GetRegistrationSpecificWorkflowsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistrationSpecificWorkflowsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistrationSpecificWorkflowsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegistrationSpecificWorkflowsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegistrationSpecificWorkflowsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegistrationSpecificWorkflowsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/workflow/registration-specific] getRegistrationSpecificWorkflowsUsingGET", response, response.Code())
	}
}

// NewGetRegistrationSpecificWorkflowsUsingGETOK creates a GetRegistrationSpecificWorkflowsUsingGETOK with default headers values
func NewGetRegistrationSpecificWorkflowsUsingGETOK() *GetRegistrationSpecificWorkflowsUsingGETOK {
	return &GetRegistrationSpecificWorkflowsUsingGETOK{}
}

/*
GetRegistrationSpecificWorkflowsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetRegistrationSpecificWorkflowsUsingGETOK struct {
	Payload []*models.PromotionWorkflow
}

// IsSuccess returns true when this get registration specific workflows using g e t o k response has a 2xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get registration specific workflows using g e t o k response has a 3xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registration specific workflows using g e t o k response has a 4xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get registration specific workflows using g e t o k response has a 5xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get registration specific workflows using g e t o k response a status code equal to that given
func (o *GetRegistrationSpecificWorkflowsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get registration specific workflows using g e t o k response
func (o *GetRegistrationSpecificWorkflowsUsingGETOK) Code() int {
	return 200
}

func (o *GetRegistrationSpecificWorkflowsUsingGETOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/workflow/registration-specific][%d] getRegistrationSpecificWorkflowsUsingGETOK %s", 200, payload)
}

func (o *GetRegistrationSpecificWorkflowsUsingGETOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/workflow/registration-specific][%d] getRegistrationSpecificWorkflowsUsingGETOK %s", 200, payload)
}

func (o *GetRegistrationSpecificWorkflowsUsingGETOK) GetPayload() []*models.PromotionWorkflow {
	return o.Payload
}

func (o *GetRegistrationSpecificWorkflowsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistrationSpecificWorkflowsUsingGETUnauthorized creates a GetRegistrationSpecificWorkflowsUsingGETUnauthorized with default headers values
func NewGetRegistrationSpecificWorkflowsUsingGETUnauthorized() *GetRegistrationSpecificWorkflowsUsingGETUnauthorized {
	return &GetRegistrationSpecificWorkflowsUsingGETUnauthorized{}
}

/*
GetRegistrationSpecificWorkflowsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRegistrationSpecificWorkflowsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get registration specific workflows using g e t unauthorized response has a 2xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registration specific workflows using g e t unauthorized response has a 3xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registration specific workflows using g e t unauthorized response has a 4xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registration specific workflows using g e t unauthorized response has a 5xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get registration specific workflows using g e t unauthorized response a status code equal to that given
func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get registration specific workflows using g e t unauthorized response
func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) Code() int {
	return 401
}

func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/workflow/registration-specific][%d] getRegistrationSpecificWorkflowsUsingGETUnauthorized", 401)
}

func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/workflow/registration-specific][%d] getRegistrationSpecificWorkflowsUsingGETUnauthorized", 401)
}

func (o *GetRegistrationSpecificWorkflowsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistrationSpecificWorkflowsUsingGETForbidden creates a GetRegistrationSpecificWorkflowsUsingGETForbidden with default headers values
func NewGetRegistrationSpecificWorkflowsUsingGETForbidden() *GetRegistrationSpecificWorkflowsUsingGETForbidden {
	return &GetRegistrationSpecificWorkflowsUsingGETForbidden{}
}

/*
GetRegistrationSpecificWorkflowsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRegistrationSpecificWorkflowsUsingGETForbidden struct {
}

// IsSuccess returns true when this get registration specific workflows using g e t forbidden response has a 2xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registration specific workflows using g e t forbidden response has a 3xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registration specific workflows using g e t forbidden response has a 4xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registration specific workflows using g e t forbidden response has a 5xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get registration specific workflows using g e t forbidden response a status code equal to that given
func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get registration specific workflows using g e t forbidden response
func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) Code() int {
	return 403
}

func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/workflow/registration-specific][%d] getRegistrationSpecificWorkflowsUsingGETForbidden", 403)
}

func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/workflow/registration-specific][%d] getRegistrationSpecificWorkflowsUsingGETForbidden", 403)
}

func (o *GetRegistrationSpecificWorkflowsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistrationSpecificWorkflowsUsingGETNotFound creates a GetRegistrationSpecificWorkflowsUsingGETNotFound with default headers values
func NewGetRegistrationSpecificWorkflowsUsingGETNotFound() *GetRegistrationSpecificWorkflowsUsingGETNotFound {
	return &GetRegistrationSpecificWorkflowsUsingGETNotFound{}
}

/*
GetRegistrationSpecificWorkflowsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRegistrationSpecificWorkflowsUsingGETNotFound struct {
}

// IsSuccess returns true when this get registration specific workflows using g e t not found response has a 2xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get registration specific workflows using g e t not found response has a 3xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get registration specific workflows using g e t not found response has a 4xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get registration specific workflows using g e t not found response has a 5xx status code
func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get registration specific workflows using g e t not found response a status code equal to that given
func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get registration specific workflows using g e t not found response
func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) Code() int {
	return 404
}

func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/workflow/registration-specific][%d] getRegistrationSpecificWorkflowsUsingGETNotFound", 404)
}

func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/workflow/registration-specific][%d] getRegistrationSpecificWorkflowsUsingGETNotFound", 404)
}

func (o *GetRegistrationSpecificWorkflowsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
