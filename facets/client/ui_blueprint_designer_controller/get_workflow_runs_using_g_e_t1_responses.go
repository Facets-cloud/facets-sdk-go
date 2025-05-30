// Code generated by go-swagger; DO NOT EDIT.

package ui_blueprint_designer_controller

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

// GetWorkflowRunsUsingGET1Reader is a Reader for the GetWorkflowRunsUsingGET1 structure.
type GetWorkflowRunsUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowRunsUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowRunsUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetWorkflowRunsUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkflowRunsUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowRunsUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cc-ui/v1/designer/{stackName}/workflow-runs] getWorkflowRunsUsingGET_1", response, response.Code())
	}
}

// NewGetWorkflowRunsUsingGET1OK creates a GetWorkflowRunsUsingGET1OK with default headers values
func NewGetWorkflowRunsUsingGET1OK() *GetWorkflowRunsUsingGET1OK {
	return &GetWorkflowRunsUsingGET1OK{}
}

/*
GetWorkflowRunsUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetWorkflowRunsUsingGET1OK struct {
	Payload *models.ListWorkflowRunsResponse
}

// IsSuccess returns true when this get workflow runs using g e t1 o k response has a 2xx status code
func (o *GetWorkflowRunsUsingGET1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow runs using g e t1 o k response has a 3xx status code
func (o *GetWorkflowRunsUsingGET1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow runs using g e t1 o k response has a 4xx status code
func (o *GetWorkflowRunsUsingGET1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow runs using g e t1 o k response has a 5xx status code
func (o *GetWorkflowRunsUsingGET1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow runs using g e t1 o k response a status code equal to that given
func (o *GetWorkflowRunsUsingGET1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workflow runs using g e t1 o k response
func (o *GetWorkflowRunsUsingGET1OK) Code() int {
	return 200
}

func (o *GetWorkflowRunsUsingGET1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/workflow-runs][%d] getWorkflowRunsUsingGET1OK %s", 200, payload)
}

func (o *GetWorkflowRunsUsingGET1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/workflow-runs][%d] getWorkflowRunsUsingGET1OK %s", 200, payload)
}

func (o *GetWorkflowRunsUsingGET1OK) GetPayload() *models.ListWorkflowRunsResponse {
	return o.Payload
}

func (o *GetWorkflowRunsUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListWorkflowRunsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowRunsUsingGET1Unauthorized creates a GetWorkflowRunsUsingGET1Unauthorized with default headers values
func NewGetWorkflowRunsUsingGET1Unauthorized() *GetWorkflowRunsUsingGET1Unauthorized {
	return &GetWorkflowRunsUsingGET1Unauthorized{}
}

/*
GetWorkflowRunsUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetWorkflowRunsUsingGET1Unauthorized struct {
}

// IsSuccess returns true when this get workflow runs using g e t1 unauthorized response has a 2xx status code
func (o *GetWorkflowRunsUsingGET1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow runs using g e t1 unauthorized response has a 3xx status code
func (o *GetWorkflowRunsUsingGET1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow runs using g e t1 unauthorized response has a 4xx status code
func (o *GetWorkflowRunsUsingGET1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow runs using g e t1 unauthorized response has a 5xx status code
func (o *GetWorkflowRunsUsingGET1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow runs using g e t1 unauthorized response a status code equal to that given
func (o *GetWorkflowRunsUsingGET1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get workflow runs using g e t1 unauthorized response
func (o *GetWorkflowRunsUsingGET1Unauthorized) Code() int {
	return 401
}

func (o *GetWorkflowRunsUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/workflow-runs][%d] getWorkflowRunsUsingGET1Unauthorized", 401)
}

func (o *GetWorkflowRunsUsingGET1Unauthorized) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/workflow-runs][%d] getWorkflowRunsUsingGET1Unauthorized", 401)
}

func (o *GetWorkflowRunsUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowRunsUsingGET1Forbidden creates a GetWorkflowRunsUsingGET1Forbidden with default headers values
func NewGetWorkflowRunsUsingGET1Forbidden() *GetWorkflowRunsUsingGET1Forbidden {
	return &GetWorkflowRunsUsingGET1Forbidden{}
}

/*
GetWorkflowRunsUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetWorkflowRunsUsingGET1Forbidden struct {
}

// IsSuccess returns true when this get workflow runs using g e t1 forbidden response has a 2xx status code
func (o *GetWorkflowRunsUsingGET1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow runs using g e t1 forbidden response has a 3xx status code
func (o *GetWorkflowRunsUsingGET1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow runs using g e t1 forbidden response has a 4xx status code
func (o *GetWorkflowRunsUsingGET1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow runs using g e t1 forbidden response has a 5xx status code
func (o *GetWorkflowRunsUsingGET1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow runs using g e t1 forbidden response a status code equal to that given
func (o *GetWorkflowRunsUsingGET1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get workflow runs using g e t1 forbidden response
func (o *GetWorkflowRunsUsingGET1Forbidden) Code() int {
	return 403
}

func (o *GetWorkflowRunsUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/workflow-runs][%d] getWorkflowRunsUsingGET1Forbidden", 403)
}

func (o *GetWorkflowRunsUsingGET1Forbidden) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/workflow-runs][%d] getWorkflowRunsUsingGET1Forbidden", 403)
}

func (o *GetWorkflowRunsUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWorkflowRunsUsingGET1NotFound creates a GetWorkflowRunsUsingGET1NotFound with default headers values
func NewGetWorkflowRunsUsingGET1NotFound() *GetWorkflowRunsUsingGET1NotFound {
	return &GetWorkflowRunsUsingGET1NotFound{}
}

/*
GetWorkflowRunsUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetWorkflowRunsUsingGET1NotFound struct {
}

// IsSuccess returns true when this get workflow runs using g e t1 not found response has a 2xx status code
func (o *GetWorkflowRunsUsingGET1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow runs using g e t1 not found response has a 3xx status code
func (o *GetWorkflowRunsUsingGET1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow runs using g e t1 not found response has a 4xx status code
func (o *GetWorkflowRunsUsingGET1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow runs using g e t1 not found response has a 5xx status code
func (o *GetWorkflowRunsUsingGET1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow runs using g e t1 not found response a status code equal to that given
func (o *GetWorkflowRunsUsingGET1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get workflow runs using g e t1 not found response
func (o *GetWorkflowRunsUsingGET1NotFound) Code() int {
	return 404
}

func (o *GetWorkflowRunsUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/workflow-runs][%d] getWorkflowRunsUsingGET1NotFound", 404)
}

func (o *GetWorkflowRunsUsingGET1NotFound) String() string {
	return fmt.Sprintf("[GET /cc-ui/v1/designer/{stackName}/workflow-runs][%d] getWorkflowRunsUsingGET1NotFound", 404)
}

func (o *GetWorkflowRunsUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
